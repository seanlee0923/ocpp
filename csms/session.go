package csms

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	"github.com/seanlee0923/ocpp/protocol"
)

var (
	ErrSessionClosed       = errors.New("OCPP session is closed")
	ErrSessionReplaced     = errors.New("OCPP session was replaced")
	ErrServerShutdown      = errors.New("OCPP server is shutting down")
	ErrPongTimeout         = errors.New("WebSocket pong timeout")
	ErrIdleTimeout         = errors.New("OCPP session idle timeout")
	ErrTooManyPendingCalls = errors.New("too many pending OCPP calls")
	ErrDuplicateUniqueID   = errors.New("duplicate OCPP unique ID")
)

type SessionState uint32

const (
	SessionActive SessionState = iota + 1
	SessionClosing
	SessionClosed
)

type callOutcome struct {
	result *protocol.CallResult
	err    error
}

// Session represents one server-owned Charging Station WebSocket connection.
// Applications may inspect it, close it, and pass it to Call or typed profile
// methods. Raw frame writes are intentionally not exposed.
type Session struct {
	identity          string
	version           protocol.Version
	principal         Principal
	conn              *websocket.Conn
	writeMu           sync.Mutex
	pendingMu         sync.Mutex
	pending           map[string]chan callOutcome
	maxPendingCalls   int
	callTimeout       time.Duration
	writeTimeout      time.Duration
	pingInterval      time.Duration
	pongTimeout       time.Duration
	idleTimeout       time.Duration
	ctx               context.Context
	cancel            context.CancelCauseFunc
	handlerSlots      chan struct{}
	handlerWG         sync.WaitGroup
	readDone          chan struct{}
	uniqueIDGenerator func() string
	closed            chan struct{}
	close             sync.Once
	errMu             sync.RWMutex
	closeErr          error
	connectedAt       time.Time
	state             atomic.Uint32
	lastReceivedAt    atomic.Int64
	lastSentAt        atomic.Int64
	lastPongAt        atomic.Int64
}

func (s *Session) Identity() string          { return s.identity }
func (s *Session) Version() protocol.Version { return s.version }
func (s *Session) Done() <-chan struct{}     { return s.closed }

func (s *Session) Principal() Principal {
	return Principal{ID: s.principal.ID, Attributes: cloneAttributes(s.principal.Attributes)}
}

func (s *Session) Context() context.Context { return s.ctx }

func (s *Session) Err() error {
	s.errMu.RLock()
	defer s.errMu.RUnlock()
	return s.closeErr
}

type SessionInfo struct {
	Identity       string
	Version        protocol.Version
	PrincipalID    string
	RemoteAddr     string
	ConnectedAt    time.Time
	LastReceivedAt time.Time
	LastSentAt     time.Time
	LastPongAt     time.Time
	State          SessionState
}

func (s *Session) Info() SessionInfo {
	remoteAddr := ""
	if s.conn != nil && s.conn.RemoteAddr() != nil {
		remoteAddr = s.conn.RemoteAddr().String()
	}
	return SessionInfo{
		Identity: s.identity, Version: s.version,
		PrincipalID: s.principal.ID,
		RemoteAddr:  remoteAddr, ConnectedAt: s.connectedAt,
		LastReceivedAt: loadTime(&s.lastReceivedAt),
		LastSentAt:     loadTime(&s.lastSentAt),
		LastPongAt:     loadTime(&s.lastPongAt),
		State:          SessionState(s.state.Load()),
	}
}

func (s *Session) markReceived() { s.lastReceivedAt.Store(time.Now().UnixNano()) }
func (s *Session) markSent()     { s.lastSentAt.Store(time.Now().UnixNano()) }
func (s *Session) markPong()     { s.lastPongAt.Store(time.Now().UnixNano()) }

func loadTime(value *atomic.Int64) time.Time {
	nanoseconds := value.Load()
	if nanoseconds == 0 {
		return time.Time{}
	}
	return time.Unix(0, nanoseconds)
}

func (s *Session) registerCall(id string) (<-chan callOutcome, error) {
	s.pendingMu.Lock()
	defer s.pendingMu.Unlock()
	select {
	case <-s.closed:
		return nil, ErrSessionClosed
	default:
	}
	if s.maxPendingCalls > 0 && len(s.pending) >= s.maxPendingCalls {
		return nil, ErrTooManyPendingCalls
	}
	if _, exists := s.pending[id]; exists {
		return nil, fmt.Errorf("%w %q", ErrDuplicateUniqueID, id)
	}
	response := make(chan callOutcome, 1)
	s.pending[id] = response
	return response, nil
}

func (s *Session) unregisterCall(id string) {
	s.pendingMu.Lock()
	delete(s.pending, id)
	s.pendingMu.Unlock()
}

func (s *Session) resolve(message protocol.Message) bool {
	s.pendingMu.Lock()
	response, ok := s.pending[message.UniqueID()]
	if ok {
		delete(s.pending, message.UniqueID())
	}
	s.pendingMu.Unlock()
	if !ok {
		return false
	}
	switch value := message.(type) {
	case protocol.CallResult:
		response <- callOutcome{result: &value}
	case protocol.CallError:
		response <- callOutcome{err: &RemoteCallError{
			Code: ErrorCode(value.Code), Description: value.Description, Details: value.Details,
		}}
	default:
		response <- callOutcome{err: fmt.Errorf("unexpected OCPP response type %d", message.Type())}
	}
	return true
}

// send is the only raw frame write path. Public callers use Call or typed
// profile methods so outbound CALLs are always registered and correlated.
func (s *Session) send(ctx context.Context, message protocol.Message) error {
	data, err := protocol.Encode(message)
	if err != nil {
		return err
	}
	s.writeMu.Lock()
	defer s.writeMu.Unlock()
	select {
	case <-s.closed:
		return ErrSessionClosed
	default:
	}
	deadline, hasDeadline := ctx.Deadline()
	if s.writeTimeout > 0 {
		configured := time.Now().Add(s.writeTimeout)
		if !hasDeadline || configured.Before(deadline) {
			deadline, hasDeadline = configured, true
		}
	}
	if hasDeadline {
		if err := s.conn.SetWriteDeadline(deadline); err != nil {
			return err
		}
	} else {
		_ = s.conn.SetWriteDeadline(time.Time{})
	}
	if err := s.conn.WriteMessage(websocket.TextMessage, data); err != nil {
		return err
	}
	s.markSent()
	return nil
}

func (s *Session) result(ctx context.Context, id string, payload any) error {
	raw, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	return s.send(ctx, protocol.CallResult{ID: id, Payload: raw})
}

func (s *Session) Close() error {
	return s.closeWithError(ErrSessionClosed)
}

func (s *Session) closeWithError(cause error) error {
	var err error
	s.close.Do(func() {
		s.state.Store(uint32(SessionClosing))
		if cause == nil {
			cause = ErrSessionClosed
		}
		s.errMu.Lock()
		s.closeErr = cause
		s.errMu.Unlock()
		s.cancel(cause)
		s.pendingMu.Lock()
		pending := s.pending
		s.pending = make(map[string]chan callOutcome)
		s.pendingMu.Unlock()
		for _, response := range pending {
			response <- callOutcome{err: cause}
		}
		code, reason := publicCloseReason(cause)
		deadline := time.Now().Add(s.writeTimeout)
		if s.writeTimeout <= 0 {
			deadline = time.Now().Add(5 * time.Second)
		}
		_ = s.conn.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(code, reason), deadline)
		err = s.conn.Close()
		close(s.closed)
		s.state.Store(uint32(SessionClosed))
	})
	return err
}

func publicCloseReason(cause error) (int, string) {
	switch {
	case errors.Is(cause, ErrServerShutdown):
		return websocket.CloseGoingAway, "server shutdown"
	case errors.Is(cause, ErrSessionReplaced):
		return websocket.CloseNormalClosure, "session replaced"
	case errors.Is(cause, ErrIdleTimeout):
		return websocket.CloseNormalClosure, "idle timeout"
	case errors.Is(cause, ErrPongTimeout):
		return websocket.CloseGoingAway, "pong timeout"
	case errors.Is(cause, protocol.ErrInvalidMessage):
		return websocket.CloseProtocolError, "invalid OCPP message"
	default:
		return websocket.CloseNormalClosure, "session closed"
	}
}

func (s *Session) startHandler(run func(context.Context)) bool {
	select {
	case s.handlerSlots <- struct{}{}:
	case <-s.ctx.Done():
		return false
	}
	s.handlerWG.Add(1)
	go func() {
		defer s.handlerWG.Done()
		defer func() { <-s.handlerSlots }()
		run(s.ctx)
	}()
	return true
}

func (s *Session) waitHandlers(ctx context.Context) error {
	select {
	case <-s.readDone:
	case <-ctx.Done():
		return ctx.Err()
	}
	done := make(chan struct{})
	go func() {
		s.handlerWG.Wait()
		close(done)
	}()
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (s *Session) pingLoop() {
	if s.pingInterval <= 0 {
		return
	}
	ticker := time.NewTicker(s.pingInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			deadline := time.Now().Add(s.writeTimeout)
			if s.writeTimeout <= 0 {
				deadline = time.Now().Add(5 * time.Second)
			}
			if err := s.conn.WriteControl(websocket.PingMessage, nil, deadline); err != nil {
				_ = s.closeWithError(err)
				return
			}
		case <-s.closed:
			return
		}
	}
}

func (s *Session) idleLoop() {
	if s.idleTimeout <= 0 {
		return
	}
	interval := s.idleTimeout / 2
	if interval > time.Second {
		interval = time.Second
	}
	if interval <= 0 {
		interval = time.Millisecond
	}
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if time.Since(loadTime(&s.lastReceivedAt)) >= s.idleTimeout {
				_ = s.closeWithError(ErrIdleTimeout)
				return
			}
		case <-s.closed:
			return
		}
	}
}
