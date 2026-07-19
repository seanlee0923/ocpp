package station

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"unicode/utf8"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/seanlee0923/ocpp/protocol"
)

const defaultMaxConcurrentHandlers = 16

var (
	ErrNotConnected               = errors.New("station: not connected")
	ErrStopped                    = errors.New("station: stopped")
	ErrTooManyPendingCalls        = errors.New("station: too many pending calls")
	ErrDuplicateUniqueID          = errors.New("station: duplicate unique ID")
	ErrUniqueIDGeneration         = errors.New("station: unique ID generation failed")
	ErrHandlerAlreadyRegistered   = errors.New("station: handler already registered for this action")
	ErrInvalidHandlerRegistration = errors.New("station: invalid handler registration")
)

// ConnectionState is the current state of a Station's connection to its
// configured CSMS.
type ConnectionState uint32

const (
	Disconnected ConnectionState = iota
	Connecting
	Connected
)

func (state ConnectionState) String() string {
	switch state {
	case Connecting:
		return "connecting"
	case Connected:
		return "connected"
	default:
		return "disconnected"
	}
}

// BasicCredentials are sent as an HTTP Basic Auth Authorization header on
// the WebSocket handshake request, matching csms's server-side
// SecurityProfileBasicAuth/SecurityProfileTLSBasicAuth handling.
type BasicCredentials struct {
	Username string
	Password string
}

// ReconnectPolicy configures exponential backoff between reconnect
// attempts. A nil ReconnectPolicy on Config disables automatic reconnect:
// Run returns as soon as the connection drops.
type ReconnectPolicy struct {
	InitialDelay time.Duration
	MaxDelay     time.Duration
	Multiplier   float64
}

func (policy *ReconnectPolicy) delay(attempt int) time.Duration {
	if policy.InitialDelay <= 0 {
		return 0
	}
	multiplier := policy.Multiplier
	if multiplier < 1 {
		multiplier = 1
	}
	delay := float64(policy.InitialDelay)
	for range attempt {
		delay *= multiplier
		if policy.MaxDelay > 0 && delay >= float64(policy.MaxDelay) {
			return policy.MaxDelay
		}
	}
	return time.Duration(delay)
}

// Config controls a Station. Zero values select the documented default;
// New rejects negative durations/limits.
type Config struct {
	URL                   string // base CSMS endpoint; Station appends "/"+Identity
	Identity              string
	Version               protocol.Version
	BasicAuth             *BasicCredentials // nil = no Authorization header
	TLSConfig             *tls.Config       // nil = ws://; non-nil enables wss:// (set Certificates here for mTLS)
	ReconnectPolicy       *ReconnectPolicy  // nil = do not auto-reconnect
	CallTimeout           time.Duration
	MaxPendingCalls       int
	MaxConcurrentHandlers int // default 16; bounds concurrent inbound CALL handler goroutines, matching csms.Config.MaxConcurrentHandlers
	HandshakeTimeout      time.Duration
	OnConnect             func(*Station)
	OnDisconnect          func(*Station, error)
	UniqueIDGenerator     func() string // default uuid.NewString
}

func (config Config) validate() error {
	if config.URL == "" {
		return fmt.Errorf("station: URL is required")
	}
	if config.Identity == "" {
		return fmt.Errorf("station: Identity is required")
	}
	if config.Version == "" {
		return fmt.Errorf("station: Version is required")
	}
	if config.CallTimeout < 0 {
		return fmt.Errorf("station: CallTimeout must not be negative")
	}
	if config.MaxPendingCalls < 0 {
		return fmt.Errorf("station: MaxPendingCalls must not be negative")
	}
	if config.MaxConcurrentHandlers < 0 {
		return fmt.Errorf("station: MaxConcurrentHandlers must not be negative")
	}
	if config.HandshakeTimeout < 0 {
		return fmt.Errorf("station: HandshakeTimeout must not be negative")
	}
	if config.ReconnectPolicy != nil && (config.ReconnectPolicy.InitialDelay < 0 || config.ReconnectPolicy.MaxDelay < 0) {
		return fmt.Errorf("station: ReconnectPolicy delays must not be negative")
	}
	return nil
}

type callOutcome struct {
	payload json.RawMessage
	err     error
}

// RemoteCallError is a CALLERROR returned by the CSMS in response to a
// Station-initiated Call.
type RemoteCallError struct {
	Code        string
	Description string
	Details     json.RawMessage
}

func (e *RemoteCallError) Error() string {
	return fmt.Sprintf("station: CALLERROR %s: %s", e.Code, e.Description)
}

// CallError lets a Handle handler control the CALLERROR's code and
// description precisely, instead of the default "InternalError" wrapping
// any other returned error.
type CallError struct {
	Code        string
	Description string
	Details     json.RawMessage
}

func (e *CallError) Error() string { return e.Description }

// stationConn is the state of one underlying WebSocket connection. A fresh
// stationConn replaces the old one on every reconnect; pending calls do not
// survive a reconnect, matching how csms.Session.closeWithError cancels
// pending calls on close. ctx is canceled when the connection closes, so
// inbound-CALL handlers (which run with ctx, not context.Background()) can
// observe disconnection instead of running indefinitely, matching how
// csms.Session runs handlers with its own session-scoped context.
type stationConn struct {
	conn      *websocket.Conn
	ctx       context.Context
	cancel    context.CancelFunc
	writeMu   sync.Mutex
	pendingMu sync.Mutex
	pending   map[string]chan callOutcome
	closed    chan struct{}
	closeOnce sync.Once
}

// newStationConn derives conn's ctx from parent (Run's ctx), not
// context.Background(): handler invocations should inherit any values the
// caller attached to Run's ctx, and observe the caller canceling Run
// directly rather than only through runConnection's watcher goroutine
// forcing the socket closed.
func newStationConn(parent context.Context, conn *websocket.Conn) *stationConn {
	ctx, cancel := context.WithCancel(parent)
	return &stationConn{conn: conn, ctx: ctx, cancel: cancel, pending: make(map[string]chan callOutcome), closed: make(chan struct{})}
}

func (c *stationConn) close(cause error) {
	c.closeOnce.Do(func() {
		close(c.closed)
		c.cancel()
		c.pendingMu.Lock()
		pending := c.pending
		c.pending = nil
		c.pendingMu.Unlock()
		for _, response := range pending {
			response <- callOutcome{err: cause}
		}
		_ = c.conn.Close()
	})
}

func (c *stationConn) registerCall(id string, maxPending int) (<-chan callOutcome, error) {
	c.pendingMu.Lock()
	defer c.pendingMu.Unlock()
	select {
	case <-c.closed:
		return nil, ErrNotConnected
	default:
	}
	if maxPending > 0 && len(c.pending) >= maxPending {
		return nil, ErrTooManyPendingCalls
	}
	if _, exists := c.pending[id]; exists {
		return nil, fmt.Errorf("%w %q", ErrDuplicateUniqueID, id)
	}
	response := make(chan callOutcome, 1)
	c.pending[id] = response
	return response, nil
}

func (c *stationConn) unregisterCall(id string) {
	c.pendingMu.Lock()
	if c.pending != nil {
		delete(c.pending, id)
	}
	c.pendingMu.Unlock()
}

func (c *stationConn) resolve(message protocol.Message) {
	c.pendingMu.Lock()
	var response chan callOutcome
	var ok bool
	if c.pending != nil {
		response, ok = c.pending[message.UniqueID()]
		if ok {
			delete(c.pending, message.UniqueID())
		}
	}
	c.pendingMu.Unlock()
	if !ok {
		return
	}
	switch value := message.(type) {
	case protocol.CallResult:
		response <- callOutcome{payload: value.Payload}
	case protocol.CallError:
		response <- callOutcome{err: &RemoteCallError{Code: value.Code, Description: value.Description, Details: value.Details}}
	default:
		response <- callOutcome{err: fmt.Errorf("station: unexpected response type %d", message.Type())}
	}
}

func (c *stationConn) send(message protocol.Message) error {
	data, err := protocol.Encode(message)
	if err != nil {
		return err
	}
	c.writeMu.Lock()
	defer c.writeMu.Unlock()
	select {
	case <-c.closed:
		return ErrNotConnected
	default:
	}
	return c.conn.WriteMessage(websocket.TextMessage, data)
}

// Station is a long-lived, reconnecting OCPP-J Charging Station client.
// Construct with New, then drive its connect/reconnect loop with Run
// (typically in its own goroutine). Handlers registered via Handle survive
// reconnects; pending outbound Calls do not.
type Station struct {
	config       Config
	uniqueIDGen  func() string
	handlersMu   sync.RWMutex
	handlers     map[string]func(context.Context, json.RawMessage) (json.RawMessage, error)
	handlerSlots chan struct{}
	connMu       sync.RWMutex
	conn         *stationConn
	state        atomic.Uint32
	stopCh       chan struct{}
	stopOnce     sync.Once
}

// New validates config and returns an unconnected Station. Call Run to
// start connecting.
func New(config Config) (*Station, error) {
	if err := config.validate(); err != nil {
		return nil, err
	}
	generator := config.UniqueIDGenerator
	if generator == nil {
		generator = uuid.NewString
	}
	maxConcurrentHandlers := config.MaxConcurrentHandlers
	if maxConcurrentHandlers == 0 {
		maxConcurrentHandlers = defaultMaxConcurrentHandlers
	}
	return &Station{
		config:       config,
		uniqueIDGen:  generator,
		handlers:     make(map[string]func(context.Context, json.RawMessage) (json.RawMessage, error)),
		handlerSlots: make(chan struct{}, maxConcurrentHandlers),
		stopCh:       make(chan struct{}),
	}, nil
}

func (s *Station) Identity() string          { return s.config.Identity }
func (s *Station) Version() protocol.Version { return s.config.Version }
func (s *Station) State() ConnectionState    { return ConnectionState(s.state.Load()) }

func (s *Station) setState(state ConnectionState) { s.state.Store(uint32(state)) }

// Stop ends Run's connect/reconnect loop and closes the current connection,
// if any. Safe to call more than once, including before Run.
func (s *Station) Stop() {
	s.stopOnce.Do(func() { close(s.stopCh) })
	if conn, err := s.currentConn(); err == nil {
		conn.close(ErrStopped)
	}
}

func (s *Station) currentConn() (*stationConn, error) {
	s.connMu.RLock()
	conn := s.conn
	s.connMu.RUnlock()
	if conn == nil {
		return nil, ErrNotConnected
	}
	return conn, nil
}

func (s *Station) setConn(conn *stationConn) {
	s.connMu.Lock()
	s.conn = conn
	s.connMu.Unlock()
}

// Run dials and, if config.ReconnectPolicy is set, keeps reconnecting with
// backoff until ctx is canceled or Stop is called. It blocks; run it in its
// own goroutine, one per Station. Returns ctx.Err(), ErrStopped, or the
// terminal dial/connection error (only when ReconnectPolicy is nil).
func (s *Station) Run(ctx context.Context) error {
	for attempt := 0; ; {
		if stopped, err := s.stoppedOrDone(ctx); stopped {
			return err
		}

		// attemptErr ends up holding whichever error this attempt
		// terminated with — the dial error if dial itself failed, or the
		// read-loop error if a connection was briefly established and then
		// dropped — so the shared "reconnect or return" logic below only
		// needs to look at one variable regardless of which happened.
		s.setState(Connecting)
		conn, attemptErr := s.dial(ctx)
		if attemptErr == nil {
			attempt = 0
			s.setConn(conn)
			s.setState(Connected)
			if s.config.OnConnect != nil {
				s.config.OnConnect(s)
			}

			attemptErr = s.runConnection(ctx, conn)
			conn.close(attemptErr)
			s.setConn(nil)
			s.setState(Disconnected)
			if s.config.OnDisconnect != nil {
				s.config.OnDisconnect(s, attemptErr)
			}

			if stopped, err := s.stoppedOrDone(ctx); stopped {
				return err
			}
		} else {
			s.setState(Disconnected)
		}

		if s.config.ReconnectPolicy == nil {
			return attemptErr
		}
		if err := s.wait(ctx, s.config.ReconnectPolicy.delay(attempt)); err != nil {
			return err
		}
		attempt++
	}
}

func (s *Station) stoppedOrDone(ctx context.Context) (bool, error) {
	select {
	case <-ctx.Done():
		return true, ctx.Err()
	case <-s.stopCh:
		return true, ErrStopped
	default:
		return false, nil
	}
}

// wait blocks for delay, or returns early with the reason if ctx is
// canceled or Stop is called first. A nil return means delay elapsed
// normally and the caller should proceed with another attempt.
func (s *Station) wait(ctx context.Context, delay time.Duration) error {
	timer := time.NewTimer(delay)
	defer timer.Stop()
	select {
	case <-timer.C:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	case <-s.stopCh:
		return ErrStopped
	}
}

// dial connects to the CSMS. Stop()/ctx cancellation must interrupt an
// in-flight handshake even when Config.HandshakeTimeout is left at its
// zero-value default (no timeout) — but gorilla/websocket's DialContext
// only converts ctx into a socket deadline when ctx.Deadline() is set; a
// plain context.WithCancel has no effect on an already-blocked handshake
// read. So dial captures the raw net.Conn as soon as the TCP connection is
// established (via Dialer.NetDialContext) and force-closes it directly if
// ctx/stopCh fires first, the same "close to unblock a blocked read"
// technique runConnection's watcher goroutine uses post-connection.
func (s *Station) dial(ctx context.Context) (*stationConn, error) {
	var mu sync.Mutex
	var rawConn net.Conn
	done := make(chan struct{})
	defer close(done)
	go func() {
		select {
		case <-ctx.Done():
		case <-s.stopCh:
		case <-done:
			return
		}
		mu.Lock()
		if rawConn != nil {
			_ = rawConn.Close()
		}
		mu.Unlock()
	}()

	dialer := websocket.Dialer{
		Subprotocols:     []string{string(s.config.Version)},
		HandshakeTimeout: s.config.HandshakeTimeout,
		TLSClientConfig:  s.config.TLSConfig,
		NetDialContext: func(dialCtx context.Context, network, addr string) (net.Conn, error) {
			conn, err := (&net.Dialer{}).DialContext(dialCtx, network, addr)
			if err != nil {
				return nil, err
			}
			mu.Lock()
			rawConn = conn
			mu.Unlock()
			return conn, nil
		},
	}
	header := http.Header{}
	if s.config.BasicAuth != nil {
		request := &http.Request{Header: header}
		request.SetBasicAuth(s.config.BasicAuth.Username, s.config.BasicAuth.Password)
	}
	target := strings.TrimSuffix(s.config.URL, "/") + "/" + url.PathEscape(s.config.Identity)
	conn, _, err := dialer.DialContext(ctx, target, header)
	if err != nil {
		return nil, err
	}
	return newStationConn(ctx, conn), nil
}

// runConnection runs the read loop until it exits on its own (connection
// error) or ctx/Stop fires, in which case it forces conn closed so the
// blocking ReadMessage call in readLoop unblocks. Returns the read loop's
// error; callers should still check ctx/stopCh afterward, since a forced
// close surfaces as a generic network error here, not ctx.Err().
func (s *Station) runConnection(ctx context.Context, conn *stationConn) error {
	done := make(chan struct{})
	go func() {
		select {
		case <-ctx.Done():
			conn.close(ctx.Err())
		case <-s.stopCh:
			conn.close(ErrStopped)
		case <-done:
		}
	}()
	err := s.readLoop(conn)
	close(done)
	return err
}

func (s *Station) readLoop(conn *stationConn) error {
	for {
		_, data, err := conn.conn.ReadMessage()
		if err != nil {
			return err
		}
		message, err := protocol.Decode(data)
		if err != nil {
			continue
		}
		switch value := message.(type) {
		case protocol.CallResult, protocol.CallError:
			conn.resolve(message)
		case protocol.Call:
			s.dispatch(conn, value)
		}
	}
}

// dispatch runs a registered handler for an inbound CALL. It blocks until a
// handler slot is available (bounded by Config.MaxConcurrentHandlers,
// mirroring csms.Session's handlerSlots backpressure) or the connection
// closes, whichever comes first, then runs the handler on its own
// goroutine. A panicking handler is recovered and reported as an
// InternalError CALLERROR instead of crashing the Station, matching
// csms.Server's handleCall panic recovery.
func (s *Station) dispatch(conn *stationConn, call protocol.Call) {
	s.handlersMu.RLock()
	handler, ok := s.handlers[call.Action]
	s.handlersMu.RUnlock()
	if !ok {
		_ = conn.send(protocol.CallError{
			ID: call.ID, Code: "NotImplemented",
			Description: "no station handler registered for " + call.Action, Details: json.RawMessage(`{}`),
		})
		return
	}
	go func() {
		// The slot semaphore is acquired here, inside the goroutine, not
		// before spawning it: this is the Station's one and only
		// connection, and its read loop also resolves this same
		// connection's outbound Call responses (conn.resolve). Blocking
		// the read loop itself on handler-slot availability would let a
		// burst of inbound CALLs starve/spuriously time out unrelated
		// pending outbound Calls queued behind them in the same byte
		// stream — a cost csms.Session.startHandler doesn't pay, since
		// each csms session's read loop only ever blocks its own single
		// charging station, never the server's calls to other sessions.
		select {
		case s.handlerSlots <- struct{}{}:
		case <-conn.closed:
			return
		}
		// Recover before releasing the slot: a single closure makes that
		// order explicit in the code instead of relying on two separate
		// defers' LIFO registration order, which would silently keep
		// working even if someone later reordered the two statements.
		defer func() {
			if recover() != nil {
				_ = conn.send(protocol.CallError{ID: call.ID, Code: "InternalError", Description: "internal error", Details: json.RawMessage(`{}`)})
			}
			<-s.handlerSlots
		}()
		payload, err := handler(conn.ctx, call.Payload)
		if err != nil {
			var callErr *CallError
			if errors.As(err, &callErr) {
				details := callErr.Details
				if details == nil {
					details = json.RawMessage(`{}`)
				}
				_ = conn.send(protocol.CallError{ID: call.ID, Code: callErr.Code, Description: callErr.Description, Details: details})
				return
			}
			_ = conn.send(protocol.CallError{ID: call.ID, Code: "InternalError", Description: err.Error(), Details: json.RawMessage(`{}`)})
			return
		}
		_ = conn.send(protocol.CallResult{ID: call.ID, Payload: payload})
	}()
}

func (s *Station) handle(action string, handler func(context.Context, json.RawMessage) (json.RawMessage, error)) error {
	s.handlersMu.Lock()
	defer s.handlersMu.Unlock()
	if _, exists := s.handlers[action]; exists {
		return fmt.Errorf("%w: %q", ErrHandlerAlreadyRegistered, action)
	}
	s.handlers[action] = handler
	return nil
}

// isNilType reports whether value is a nil interface or a nil pointer,
// mirroring csms's typed_handler.go isNilType — Handle's type parameters
// must be non-pointer generated payload types, not pointers to them.
func isNilType[T any](value T) bool {
	typeOf := reflect.TypeOf(value)
	return typeOf == nil || typeOf.Kind() == reflect.Pointer
}

// generateUniqueID recovers from a panicking generator and validates the
// resulting ID's length, mirroring csms/call.go's generateUniqueID.
func generateUniqueID(generator func() string) (id string, err error) {
	defer func() {
		if recovered := recover(); recovered != nil {
			id = ""
			err = fmt.Errorf("%w: generator panic: %v", ErrUniqueIDGeneration, recovered)
		}
	}()
	if generator == nil {
		return "", fmt.Errorf("%w: generator is nil", ErrUniqueIDGeneration)
	}
	id = generator()
	if length := utf8.RuneCountInString(id); length == 0 || length > protocol.MaxUniqueIDLength {
		return "", fmt.Errorf("%w: ID must contain 1 to %d characters", ErrUniqueIDGeneration, protocol.MaxUniqueIDLength)
	}
	return id, nil
}
