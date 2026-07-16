package csms

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"ocpp-go/protocol"
)

type UniqueIDGenerator func() string

type DuplicateSessionDecision uint8

const (
	RejectNewSession DuplicateSessionDecision = iota + 1
	ReplaceExistingSession
)

type DuplicateSessionAttempt struct {
	Identity string
	Existing SessionInfo
	Request  *http.Request
}

type DuplicateSessionHandler func(context.Context, DuplicateSessionAttempt) (DuplicateSessionDecision, error)

type ErrorCode string

const (
	NotImplemented                ErrorCode = "NotImplemented"
	NotSupported                  ErrorCode = "NotSupported"
	FormationViolation            ErrorCode = "FormationViolation"
	InternalError                 ErrorCode = "InternalError"
	ProtocolError                 ErrorCode = "ProtocolError"
	SecurityError                 ErrorCode = "SecurityError"
	PropertyConstraintViolation   ErrorCode = "PropertyConstraintViolation"
	OccurrenceConstraintViolation ErrorCode = "OccurrenceConstraintViolation"
	TypeConstraintViolation       ErrorCode = "TypeConstraintViolation"
	GenericError                  ErrorCode = "GenericError"
)

type CallError struct {
	Code        ErrorCode
	Description string
	Details     any
}

func (e *CallError) Error() string { return e.Description }

type Config struct {
	Router                *Router
	Versions              []protocol.Version
	ReadLimit             int64
	HandshakeTimeout      time.Duration
	CallTimeout           time.Duration
	MaxPendingCalls       int
	MaxConcurrentHandlers int
	UniqueIDGenerator     UniqueIDGenerator
	OnDuplicateSession    DuplicateSessionHandler
	WriteTimeout          time.Duration
	PingInterval          time.Duration
	PongTimeout           time.Duration
	IdleTimeout           time.Duration
	CheckOrigin           func(*http.Request) bool
	OnConnect             func(*Session)
	OnDisconnect          func(*Session, error)
	Security              SecurityConfig
}

type Server struct {
	config              Config
	upgrader            websocket.Upgrader
	sessionsMu          sync.RWMutex
	sessions            map[string]*sessionEntry
	reservationSequence atomic.Uint64
	shuttingDown        atomic.Bool
}

type sessionEntry struct {
	session     *Session
	reservation uint64
}

func New(config Config) (*Server, error) {
	if err := prepareSecurity(&config.Security); err != nil {
		return nil, err
	}
	if config.Router == nil {
		config.Router = NewRouter()
	}
	if len(config.Versions) == 0 {
		config.Versions = protocol.SupportedVersions()
	}
	for _, version := range config.Versions {
		if !version.Valid() {
			return nil, fmt.Errorf("invalid configured OCPP version %q", version)
		}
	}
	if config.ReadLimit == 0 {
		config.ReadLimit = 1 << 20
	}
	if config.HandshakeTimeout == 0 {
		config.HandshakeTimeout = 10 * time.Second
	}
	if config.CallTimeout == 0 {
		config.CallTimeout = 30 * time.Second
	}
	if config.CallTimeout < 0 {
		return nil, fmt.Errorf("call timeout must not be negative")
	}
	if config.MaxPendingCalls == 0 {
		config.MaxPendingCalls = 100
	}
	if config.MaxPendingCalls < 0 {
		return nil, fmt.Errorf("max pending calls must not be negative")
	}
	if config.MaxConcurrentHandlers == 0 {
		config.MaxConcurrentHandlers = 16
	}
	if config.MaxConcurrentHandlers < 0 {
		return nil, fmt.Errorf("max concurrent handlers must not be negative")
	}
	if config.UniqueIDGenerator == nil {
		config.UniqueIDGenerator = uuid.NewString
	}
	if config.WriteTimeout == 0 {
		config.WriteTimeout = 10 * time.Second
	}
	if config.PingInterval == 0 {
		config.PingInterval = 30 * time.Second
	}
	if config.PongTimeout == 0 {
		config.PongTimeout = 90 * time.Second
	}
	if config.WriteTimeout < 0 || config.PingInterval < 0 || config.PongTimeout < 0 || config.IdleTimeout < 0 {
		return nil, fmt.Errorf("write, ping, pong and idle timeouts must not be negative")
	}
	if config.PingInterval > 0 && config.PongTimeout > 0 && config.PongTimeout <= config.PingInterval {
		return nil, fmt.Errorf("pong timeout must be greater than ping interval")
	}
	protocols := make([]string, len(config.Versions))
	for i, version := range config.Versions {
		protocols[i] = string(version)
	}
	return &Server{
		config:   config,
		sessions: make(map[string]*sessionEntry),
		upgrader: websocket.Upgrader{
			HandshakeTimeout: config.HandshakeTimeout,
			Subprotocols:     protocols,
			CheckOrigin:      config.CheckOrigin,
		},
	}, nil
}

// ServeHTTP accepts /{chargePointIdentity} and negotiates an OCPP subprotocol.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if s.shuttingDown.Load() {
		http.Error(w, "server is shutting down", http.StatusServiceUnavailable)
		return
	}
	identity := strings.Trim(strings.TrimSpace(r.URL.Path), "/")
	if strings.Contains(identity, "/") || s.config.Security.ValidateIdentity(identity) != nil {
		http.Error(w, "invalid charging station identity", http.StatusBadRequest)
		return
	}
	if r.Method != http.MethodGet || !websocket.IsWebSocketUpgrade(r) {
		http.Error(w, "expected a WebSocket upgrade request", http.StatusBadRequest)
		return
	}
	version, ok := negotiatedVersion(r, s.config.Versions)
	if !ok {
		http.Error(w, "unsupported OCPP WebSocket subprotocol", http.StatusBadRequest)
		return
	}
	attempt := HandshakeAttempt{Identity: identity, Version: version, RemoteAddr: r.RemoteAddr}
	if limiter := s.config.Security.HandshakeLimiter; limiter != nil && !limiter.Allow(r.Context(), attempt) {
		s.securityEvent(r.Context(), SecurityEvent{
			Type: SecurityEventHandshakeRateLimited, Identity: identity,
			Version: version, RemoteAddr: r.RemoteAddr, Reason: "handshake rate limit exceeded",
		})
		http.Error(w, "too many connection attempts", http.StatusTooManyRequests)
		return
	}
	principal, status, authErr := s.authenticate(r, identity, version)
	if authErr != nil {
		s.securityEvent(r.Context(), SecurityEvent{
			Type: SecurityEventAuthenticationFailed, Identity: identity,
			Version: version, RemoteAddr: r.RemoteAddr, Reason: authErr.Error(),
		})
		if status == http.StatusUnauthorized {
			w.Header().Set("WWW-Authenticate", `Basic realm="OCPP", charset="UTF-8"`)
		}
		http.Error(w, "connection authentication failed", status)
		return
	}
	s.securityEvent(r.Context(), SecurityEvent{
		Type: SecurityEventAuthenticationSucceeded, Identity: identity,
		Version: version, RemoteAddr: r.RemoteAddr, Reason: "authentication succeeded",
	})
	reservation, previous, ok := s.reserveIdentity(w, r, identity)
	if !ok {
		return
	}
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.releaseReservation(identity, reservation)
		return
	}
	version, err = protocol.ParseVersion(conn.Subprotocol())
	if err != nil {
		s.releaseReservation(identity, reservation)
		_ = conn.Close()
		return
	}
	sessionContext, cancelSession := context.WithCancelCause(r.Context())
	session := &Session{
		identity: identity, version: version, principal: principal, conn: conn,
		pending:           make(map[string]chan callOutcome),
		maxPendingCalls:   s.config.MaxPendingCalls,
		callTimeout:       s.config.CallTimeout,
		writeTimeout:      s.config.WriteTimeout,
		pingInterval:      s.config.PingInterval,
		pongTimeout:       s.config.PongTimeout,
		idleTimeout:       s.config.IdleTimeout,
		ctx:               sessionContext,
		cancel:            cancelSession,
		handlerSlots:      make(chan struct{}, s.config.MaxConcurrentHandlers),
		readDone:          make(chan struct{}),
		uniqueIDGenerator: s.config.UniqueIDGenerator,
		closed:            make(chan struct{}), connectedAt: time.Now(),
	}
	session.state.Store(uint32(SessionActive))
	session.lastReceivedAt.Store(session.connectedAt.UnixNano())
	if session.pongTimeout > 0 {
		_ = conn.SetReadDeadline(time.Now().Add(session.pongTimeout))
		conn.SetPongHandler(func(string) error {
			session.markPong()
			return conn.SetReadDeadline(time.Now().Add(session.pongTimeout))
		})
	}
	if !s.activateReservation(identity, reservation, session) {
		_ = session.Close()
		return
	}
	if previous != nil {
		_ = previous.closeWithError(ErrSessionReplaced)
	}
	conn.SetReadLimit(s.config.ReadLimit)
	go session.pingLoop()
	go session.idleLoop()
	readResult := make(chan error, 1)
	go func() {
		readResult <- s.readLoop(session)
		close(session.readDone)
	}()
	if s.config.OnConnect != nil {
		s.config.OnConnect(session)
	}
	err = <-readResult
	_ = session.closeWithError(err)
	s.unregisterSession(session)
	if s.config.OnDisconnect != nil {
		s.config.OnDisconnect(session, err)
	}
}

func (s *Server) Session(identity string) (*Session, bool) {
	s.sessionsMu.RLock()
	defer s.sessionsMu.RUnlock()
	entry := s.sessions[identity]
	if entry == nil || entry.session == nil {
		return nil, false
	}
	return entry.session, true
}

func (s *Server) Sessions() []*Session {
	s.sessionsMu.RLock()
	defer s.sessionsMu.RUnlock()
	result := make([]*Session, 0, len(s.sessions))
	for _, entry := range s.sessions {
		if entry.session != nil {
			result = append(result, entry.session)
		}
	}
	return result
}

func (s *Server) SessionCount() int { return len(s.Sessions()) }

// Shutdown rejects new connections, closes active sessions and waits until
// they finish or ctx expires.
func (s *Server) Shutdown(ctx context.Context) error {
	s.shuttingDown.Store(true)
	s.sessionsMu.RLock()
	sessions := make([]*Session, 0, len(s.sessions))
	for _, entry := range s.sessions {
		if entry.session != nil {
			sessions = append(sessions, entry.session)
		}
	}
	s.sessionsMu.RUnlock()
	for _, session := range sessions {
		go session.closeWithError(ErrServerShutdown)
	}
	for _, session := range sessions {
		select {
		case <-session.Done():
			s.unregisterSession(session)
		case <-ctx.Done():
			return ctx.Err()
		}
		if err := session.waitHandlers(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) reserveIdentity(w http.ResponseWriter, r *http.Request, identity string) (uint64, *Session, bool) {
	for {
		s.sessionsMu.Lock()
		if s.shuttingDown.Load() {
			s.sessionsMu.Unlock()
			http.Error(w, "server is shutting down", http.StatusServiceUnavailable)
			return 0, nil, false
		}
		entry := s.sessions[identity]
		if entry == nil {
			token := s.nextReservation()
			s.sessions[identity] = &sessionEntry{reservation: token}
			s.sessionsMu.Unlock()
			return token, nil, true
		}
		if entry.reservation != 0 {
			s.sessionsMu.Unlock()
			http.Error(w, "charging station connection is already pending", http.StatusConflict)
			return 0, nil, false
		}
		existing := entry.session
		s.sessionsMu.Unlock()

		decision := RejectNewSession
		if s.config.OnDuplicateSession != nil {
			var err error
			decision, err = s.config.OnDuplicateSession(r.Context(), DuplicateSessionAttempt{
				Identity: identity, Existing: existing.Info(), Request: r,
			})
			if err != nil {
				http.Error(w, "duplicate session policy failed", http.StatusServiceUnavailable)
				return 0, nil, false
			}
		}
		if decision == RejectNewSession {
			s.securityEvent(r.Context(), SecurityEvent{
				Type: SecurityEventDuplicateRejected, Identity: identity,
				Version: existing.version, RemoteAddr: r.RemoteAddr, Reason: "duplicate session rejected",
			})
			http.Error(w, "charging station is already connected", http.StatusConflict)
			return 0, nil, false
		}
		if decision != ReplaceExistingSession {
			http.Error(w, "invalid duplicate session decision", http.StatusInternalServerError)
			return 0, nil, false
		}

		s.sessionsMu.Lock()
		current := s.sessions[identity]
		if current == nil || current.session != existing || current.reservation != 0 {
			s.sessionsMu.Unlock()
			continue
		}
		token := s.nextReservation()
		current.reservation = token
		s.sessionsMu.Unlock()
		s.securityEvent(r.Context(), SecurityEvent{
			Type: SecurityEventSessionReplaced, Identity: identity,
			Version: existing.version, RemoteAddr: r.RemoteAddr, Reason: "existing session replacement approved",
		})
		return token, existing, true
	}
}

func (s *Server) nextReservation() uint64 {
	value := s.reservationSequence.Add(1)
	if value == 0 {
		value = s.reservationSequence.Add(1)
	}
	return value
}

func (s *Server) releaseReservation(identity string, token uint64) {
	s.sessionsMu.Lock()
	defer s.sessionsMu.Unlock()
	entry := s.sessions[identity]
	if entry == nil || entry.reservation != token {
		return
	}
	if entry.session == nil {
		delete(s.sessions, identity)
	} else {
		entry.reservation = 0
	}
}

func (s *Server) activateReservation(identity string, token uint64, session *Session) bool {
	s.sessionsMu.Lock()
	defer s.sessionsMu.Unlock()
	entry := s.sessions[identity]
	if entry == nil || entry.reservation != token {
		return false
	}
	if s.shuttingDown.Load() {
		if entry.session == nil {
			delete(s.sessions, identity)
		} else {
			entry.reservation = 0
		}
		return false
	}
	entry.session = session
	entry.reservation = 0
	return true
}

func (s *Server) unregisterSession(session *Session) {
	s.sessionsMu.Lock()
	defer s.sessionsMu.Unlock()
	entry := s.sessions[session.identity]
	if entry != nil && entry.session == session && entry.reservation == 0 {
		delete(s.sessions, session.identity)
	}
}

func (s *Server) readLoop(session *Session) error {
	for {
		messageType, data, err := session.conn.ReadMessage()
		if err != nil {
			if netErr, ok := err.(interface{ Timeout() bool }); ok && netErr.Timeout() {
				return fmt.Errorf("%w: %v", ErrPongTimeout, err)
			}
			return err
		}
		if messageType != websocket.TextMessage {
			return fmt.Errorf("OCPP only accepts WebSocket text messages")
		}
		session.markReceived()
		message, err := protocol.Decode(data)
		if err != nil {
			return err
		}
		switch value := message.(type) {
		case protocol.Call:
			if !session.startHandler(func(ctx context.Context) {
				s.handleCall(ctx, session, value)
			}) {
				return context.Cause(session.ctx)
			}
		case protocol.Send:
			if session.version != protocol.OCPP21 {
				return fmt.Errorf("SEND is only supported by OCPP 2.1")
			}
			if !session.startHandler(func(ctx context.Context) {
				s.handleSend(ctx, session, value)
			}) {
				return context.Cause(session.ctx)
			}
		case protocol.CallResult, protocol.CallError:
			// Unknown IDs are late or unsolicited responses and are ignored.
			session.resolve(message)
		default:
			return fmt.Errorf("unexpected inbound message type %d", message.Type())
		}
	}
}

func (s *Server) handleSend(ctx context.Context, session *Session, send protocol.Send) {
	handler, ok := s.config.Router.lookup(session.version, send.Action)
	if !ok {
		return
	}
	// OCPP 2.1 SEND is unconfirmed. Handler and validation failures are
	// intentionally not converted to CALLRESULT or CALLERROR.
	_, _ = handler(ctx, session, send.Payload)
}

func (s *Server) handleCall(ctx context.Context, session *Session, call protocol.Call) {
	handler, ok := s.config.Router.lookup(session.version, call.Action)
	if !ok {
		_ = session.Send(ctx, protocol.CallError{ID: call.ID, Code: string(NotImplemented), Description: "action is not registered"})
		return
	}
	response, err := handler(ctx, session, call.Payload)
	if err == nil {
		_ = session.result(ctx, call.ID, response)
		return
	}
	callError := &CallError{Code: InternalError, Description: "internal error", Details: map[string]any{}}
	if !errors.As(err, &callError) {
		callError = &CallError{Code: InternalError, Description: "internal error", Details: map[string]any{}}
	}
	details := callError.Details
	if details == nil {
		details = map[string]any{}
	}
	raw, marshalErr := jsonMarshal(details)
	if marshalErr != nil {
		raw = []byte(`{}`)
	}
	_ = session.Send(ctx, protocol.CallError{ID: call.ID, Code: string(callError.Code), Description: callError.Description, Details: raw})
}

func hasSupportedSubprotocol(r *http.Request, supported []protocol.Version) bool {
	_, ok := negotiatedVersion(r, supported)
	return ok
}

func negotiatedVersion(r *http.Request, supported []protocol.Version) (protocol.Version, bool) {
	requested := websocket.Subprotocols(r)
	for _, candidate := range supported {
		for _, value := range requested {
			if value == string(candidate) {
				return candidate, true
			}
		}
	}
	return "", false
}
