package csms

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/seanlee0923/ocpp/protocol"
)

var (
	ErrInvalidConfiguration = errors.New("invalid CSMS configuration")
	ErrUniqueIDGeneration   = errors.New("OCPP unique ID generation failed")
)

// UniqueIDGenerator returns an OCPP-J unique ID of 1 to 36 characters.
// Panics, empty IDs, overlong IDs, and duplicate in-flight IDs are converted
// to errors by Call.
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
	FormatViolation               ErrorCode = "FormatViolation"
	InternalError                 ErrorCode = "InternalError"
	ProtocolError                 ErrorCode = "ProtocolError"
	SecurityError                 ErrorCode = "SecurityError"
	PropertyConstraintViolation   ErrorCode = "PropertyConstraintViolation"
	OccurenceConstraintViolation  ErrorCode = "OccurenceConstraintViolation"
	OccurrenceConstraintViolation ErrorCode = "OccurrenceConstraintViolation"
	TypeConstraintViolation       ErrorCode = "TypeConstraintViolation"
	MessageTypeNotSupported       ErrorCode = "MessageTypeNotSupported"
	RpcFrameworkError             ErrorCode = "RpcFrameworkError"
	GenericError                  ErrorCode = "GenericError"
)

type CallError struct {
	Code        ErrorCode
	Description string
	Details     any
}

func (e *CallError) Error() string { return e.Description }

// Config controls a CSMS Server. Construct it with keyed fields: new optional
// fields may be added in backward-compatible releases. Zero values select the
// documented default except IdleTimeout, where zero disables idle detection.
// Negative durations and limits are rejected.
type Config struct {
	Router                *Router
	Versions              []protocol.Version
	ReadLimit             int64
	HandshakeTimeout      time.Duration
	CallTimeout           time.Duration
	MaxPendingCalls       int
	MaxConcurrentHandlers int
	// MaxQueuedHandlers bounds how many inbound CALL/SEND messages may be
	// admitted (running or waiting for a free MaxConcurrentHandlers slot)
	// at once, per session. Default: 4x MaxConcurrentHandlers. Once full, a
	// further inbound CALL gets an immediate CALLERROR (GenericError) and a
	// further inbound SEND is logged and dropped, instead of either
	// blocking the read loop (which would also stall reading this
	// session's own outbound Call responses) or admitting an unbounded
	// number of goroutines waiting for a slot.
	MaxQueuedHandlers  int
	UniqueIDGenerator  UniqueIDGenerator
	OnDuplicateSession DuplicateSessionHandler
	WriteTimeout       time.Duration
	PingInterval       time.Duration
	PongTimeout        time.Duration
	IdleTimeout        time.Duration
	CheckOrigin        func(*http.Request) bool
	OnConnect          func(*Session)
	OnDisconnect       func(*Session, error)
	Logger             Logger
	Metrics            Metrics
	Security           SecurityConfig
}

type Server struct {
	config              Config
	upgrader            websocket.Upgrader
	sessionsMu          sync.RWMutex
	sessions            map[string]*sessionEntry
	reservationSequence atomic.Uint64
	shuttingDown        atomic.Bool
	metricQueue         chan MetricEvent
	metricStop          chan struct{}
	metricStopOnce      sync.Once
	metricWG            sync.WaitGroup
}

type sessionEntry struct {
	session     *Session
	reservation uint64
}

func New(config Config) (*Server, error) {
	if err := prepareSecurity(&config.Security); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidConfiguration, err)
	}
	if config.Router == nil {
		config.Router = NewRouter()
	}
	if len(config.Versions) == 0 {
		config.Versions = protocol.SupportedVersions()
	}
	for _, version := range config.Versions {
		if !version.Valid() {
			return nil, fmt.Errorf("%w: invalid configured OCPP version %q", ErrInvalidConfiguration, version)
		}
	}
	if config.ReadLimit == 0 {
		config.ReadLimit = 1 << 20
	}
	if config.ReadLimit < 0 {
		return nil, fmt.Errorf("%w: read limit must not be negative", ErrInvalidConfiguration)
	}
	if config.HandshakeTimeout == 0 {
		config.HandshakeTimeout = 10 * time.Second
	}
	if config.HandshakeTimeout < 0 {
		return nil, fmt.Errorf("%w: handshake timeout must not be negative", ErrInvalidConfiguration)
	}
	if config.CallTimeout == 0 {
		config.CallTimeout = 30 * time.Second
	}
	if config.CallTimeout < 0 {
		return nil, fmt.Errorf("%w: call timeout must not be negative", ErrInvalidConfiguration)
	}
	if config.MaxPendingCalls == 0 {
		config.MaxPendingCalls = 100
	}
	if config.MaxPendingCalls < 0 {
		return nil, fmt.Errorf("%w: max pending calls must not be negative", ErrInvalidConfiguration)
	}
	if config.MaxConcurrentHandlers == 0 {
		config.MaxConcurrentHandlers = 16
	}
	if config.MaxConcurrentHandlers < 0 {
		return nil, fmt.Errorf("%w: max concurrent handlers must not be negative", ErrInvalidConfiguration)
	}
	if config.MaxQueuedHandlers == 0 {
		config.MaxQueuedHandlers = 4 * config.MaxConcurrentHandlers
	}
	if config.MaxQueuedHandlers < 0 {
		return nil, fmt.Errorf("%w: max queued handlers must not be negative", ErrInvalidConfiguration)
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
		return nil, fmt.Errorf("%w: write, ping, pong and idle timeouts must not be negative", ErrInvalidConfiguration)
	}
	if config.PingInterval > 0 && config.PongTimeout > 0 && config.PongTimeout <= config.PingInterval {
		return nil, fmt.Errorf("%w: pong timeout must be greater than ping interval", ErrInvalidConfiguration)
	}
	protocols := make([]string, len(config.Versions))
	for i, version := range config.Versions {
		protocols[i] = string(version)
	}
	server := &Server{
		config:   config,
		sessions: make(map[string]*sessionEntry),
		upgrader: websocket.Upgrader{
			HandshakeTimeout: config.HandshakeTimeout,
			Subprotocols:     protocols,
			CheckOrigin:      config.CheckOrigin,
		},
	}
	if config.Metrics != nil {
		server.metricQueue = make(chan MetricEvent, maxConcurrentMetricDispatch)
		server.metricStop = make(chan struct{})
		startMetricDispatch(server.metricQueue, server.metricStop, &server.metricWG, config.Metrics)
	}
	return server, nil
}

// disconnectHandlerGrace bounds how long ServeHTTP's ordinary disconnect
// path (a dropped connection, not Shutdown) waits for a session's in-flight
// CALL/SEND handler goroutines to actually exit before unregistering the
// session and calling OnDisconnect. session.ctx is already canceled by the
// time this wait starts, so a well-behaved handler (one that respects ctx,
// as documented) typically returns almost immediately — this grace period
// only matters for narrowing, not eliminating, the window where a fast
// reconnect for the same identity could otherwise start a new session
// while a handler from the just-closed one is still running and touching
// whatever per-identity state the application keeps. Deliberately much
// shorter than CallTimeout: this fires on every ordinary disconnect, not
// just Shutdown, so it must not make routine disconnects noticeably slow.
const disconnectHandlerGrace = 5 * time.Second

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
	// Guarantees the reservation is released on every exit from here on,
	// including a panic — s.upgrader.Upgrade below calls the user-supplied
	// Config.CheckOrigin internally, with no recover anywhere in that call
	// path, so a panicking CheckOrigin would otherwise skip every
	// releaseReservation call below it and permanently reject this
	// identity with "connection is already pending". Safe to defer
	// unconditionally: releaseReservation is a no-op once
	// activateReservation has already zeroed out this same token below.
	defer s.releaseReservation(identity, reservation)
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	version, err = protocol.ParseVersion(conn.Subprotocol())
	if err != nil {
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
		pendingSlots:      make(chan struct{}, s.config.MaxQueuedHandlers),
		readDone:          make(chan struct{}),
		uniqueIDGenerator: s.config.UniqueIDGenerator,
		metricQueue:       s.metricQueue,
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
	if err := s.activateReservation(identity, reservation, session); err != nil {
		_ = session.closeWithError(err)
		return
	}
	if previous != nil {
		_ = previous.closeWithError(ErrSessionReplaced)
	}
	// Record Connected before starting the read loop or calling OnConnect:
	// both can synchronously produce other events for this session (an
	// inbound CALL as soon as the Charging Station writes one, or an
	// outbound csms.Call from within OnConnect, which is a documented
	// supported pattern), and Connected must be observable first.
	s.log(session.Context(), sessionLogRecord(session, LogInfo, LogSessionConnected))
	// context.Background(): a concurrent Server.Shutdown could in principle
	// cancel session.Context() in the narrow window between
	// activateReservation and here, and this event must not be silently
	// dropped by a Metrics implementation that defensively no-ops on an
	// already-canceled ctx (see the same reasoning in call.go's finish).
	session.recordMetric(MetricEvent{
		Type: MetricSessionConnected, Identity: session.identity, Version: session.version,
	})
	conn.SetReadLimit(s.config.ReadLimit)
	go session.pingLoop()
	go session.idleLoop()
	readResult := make(chan error, 1)
	go func() {
		defer close(session.readDone)
		// This goroutine is not covered by net/http's per-request panic
		// recovery: that only protects code running synchronously inside
		// ServeHTTP's own call frame (like callOnConnect below), never a
		// goroutine ServeHTTP spawns. An unrecovered panic here — in
		// ReadMessage, markReceived, or protocol.Decode, none of which run
		// inside handleCall/handleSend's own recover — would crash the
		// entire process, disconnecting every session, not just this one.
		defer func() {
			if r := recover(); r != nil {
				readResult <- fmt.Errorf("csms: session read loop panicked: %v", r)
			}
		}()
		readResult <- s.readLoop(session)
	}()
	s.callOnConnect(session)
	err = <-readResult
	_ = session.closeWithError(err)
	graceCtx, graceCancel := context.WithTimeout(context.Background(), disconnectHandlerGrace)
	_ = session.waitHandlers(graceCtx)
	graceCancel()
	s.unregisterSession(session)
	// session.Err() (not the raw readLoop return above) is the true cause:
	// closeWithError only ever applies the *first* cause it's given
	// (sync.Once). If a concurrent closeWithError elsewhere (idle timeout,
	// Shutdown, a duplicate-session replacement) already fired and force-
	// closed the connection, readLoop's own err here is just a generic
	// "connection closed" error surfaced by the unblocked ReadMessage —
	// session.Err() still holds the real, typed cause that was recorded
	// first.
	cause := session.Err()
	s.callOnDisconnect(session, cause)
	record := sessionLogRecord(session, LogInfo, LogSessionDisconnected)
	record.Reason = disconnectReason(cause)
	s.log(context.Background(), record)
	session.recordMetric(MetricEvent{
		Type: MetricSessionDisconnected, Identity: session.identity, Version: session.version,
		Duration: time.Since(session.connectedAt),
	})
}

// callOnConnect isolates a panicking OnConnect from the rest of
// ServeHTTP's cleanup sequence. Without this, a panic here (net/http
// recovers it per-request, so it doesn't take down the server) still skips
// everything after it in ServeHTTP: readResult is never awaited,
// closeWithError and unregisterSession never run, and OnDisconnect never
// fires — the session and its pingLoop/idleLoop/readLoop goroutines leak
// for the server's remaining lifetime.
func (s *Server) callOnConnect(session *Session) {
	if s.config.OnConnect == nil {
		return
	}
	defer func() {
		if recover() != nil {
			record := sessionLogRecord(session, LogError, LogSessionConnected)
			record.Reason = "on_connect_panic"
			s.log(session.Context(), record)
		}
	}()
	s.config.OnConnect(session)
}

// callOnDisconnect mirrors callOnConnect: a panicking OnDisconnect must not
// skip the disconnect log/metric recorded right after it.
func (s *Server) callOnDisconnect(session *Session, cause error) {
	if s.config.OnDisconnect == nil {
		return
	}
	defer func() {
		if recover() != nil {
			record := sessionLogRecord(session, LogError, LogSessionDisconnected)
			record.Reason = "on_disconnect_panic"
			// context.Background(): session.Context() may already be
			// canceled by the time OnDisconnect runs, for the same reason
			// the disconnect log a few lines below also uses
			// context.Background() instead of session.Context().
			s.log(context.Background(), record)
		}
	}()
	s.config.OnDisconnect(session, cause)
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

// ServerSnapshot is a point-in-time view of server status, suitable for a
// health or readiness endpoint the application wires up itself.
type ServerSnapshot struct {
	ActiveSessions int
	ShuttingDown   bool
	Sessions       []SessionInfo
}

// Snapshot returns the server's current status. It is safe to call from any
// goroutine and does not block on session I/O.
func (s *Server) Snapshot() ServerSnapshot {
	sessions := s.Sessions()
	infos := make([]SessionInfo, len(sessions))
	for i, session := range sessions {
		infos[i] = session.Info()
	}
	return ServerSnapshot{ActiveSessions: len(infos), ShuttingDown: s.shuttingDown.Load(), Sessions: infos}
}

// Healthy reports whether the server is still accepting new connections. It
// becomes false as soon as Shutdown is called.
func (s *Server) Healthy() bool { return !s.shuttingDown.Load() }

// Shutdown rejects new connections, closes active sessions and waits until
// they finish or ctx expires, then stops this Server's metric dispatch
// workers (if Config.Metrics was set) and waits for them to actually exit
// too — Shutdown is terminal, a Server cannot accept new connections
// afterward, so it should honestly release everything it started, not
// leave metricDispatchWorkers goroutines running forever. Repeated calls
// are safe.
func (s *Server) Shutdown(ctx context.Context) error {
	s.shuttingDown.Store(true)
	if s.metricStop != nil {
		// Signals the metric workers to stop on every exit from here on,
		// including an early return from the session-wait loop below when
		// ctx expires first — otherwise that path skips the explicit
		// metricStopOnce.Do call further down entirely, leaking
		// metricDispatchWorkers forever and contradicting this method's
		// own doc comment. metricStopOnce makes a second Do call (the
		// explicit one below, in the normal path) a no-op, not a double
		// close.
		defer s.metricStopOnce.Do(func() { close(s.metricStop) })
	}
	s.sessionsMu.RLock()
	sessions := make([]*Session, 0, len(s.sessions))
	for _, entry := range s.sessions {
		if entry.session != nil {
			sessions = append(sessions, entry.session)
		}
	}
	s.sessionsMu.RUnlock()
	for _, session := range sessions {
		go func() { _ = session.closeWithError(ErrServerShutdown) }()
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
	if s.metricStop != nil {
		s.metricStopOnce.Do(func() { close(s.metricStop) })
		done := make(chan struct{})
		go func() {
			s.metricWG.Wait()
			close(done)
		}()
		select {
		case <-done:
		case <-ctx.Done():
			return ctx.Err()
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

// activateReservation returns nil on success, or the reason the reservation
// could not be activated so the caller can close session with that exact
// cause (e.g. ErrServerShutdown) rather than a generic one.
func (s *Server) activateReservation(identity string, token uint64, session *Session) error {
	s.sessionsMu.Lock()
	defer s.sessionsMu.Unlock()
	entry := s.sessions[identity]
	if entry == nil || entry.reservation != token {
		return ErrSessionClosed
	}
	if s.shuttingDown.Load() {
		if entry.session == nil {
			delete(s.sessions, identity)
		} else {
			entry.reservation = 0
		}
		return ErrServerShutdown
	}
	entry.session = session
	entry.reservation = 0
	return nil
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
			var unsupported *protocol.UnsupportedMessageTypeError
			if session.version != protocol.OCPP16 && errors.As(err, &unsupported) {
				_ = session.send(session.Context(), protocol.CallError{
					ID: unsupported.ID, Code: string(MessageTypeNotSupported),
					Description: "message type is not supported", Details: json.RawMessage(`{}`),
				})
				continue
			}
			return err
		}
		switch value := message.(type) {
		case protocol.Call:
			switch err := session.startHandler(func(ctx context.Context) {
				s.handleCall(ctx, session, value)
			}); {
			case err == nil:
			case errors.Is(err, ErrHandlerQueueFull):
				s.rejectQueuedCall(session, value)
			default:
				return err
			}
		case protocol.Send:
			if session.version != protocol.OCPP21 {
				if session.version == protocol.OCPP201 {
					_ = session.send(session.Context(), protocol.CallError{
						ID: value.ID, Code: string(MessageTypeNotSupported),
						Description: "SEND is only supported by OCPP 2.1", Details: json.RawMessage(`{}`),
					})
					continue
				}
				return fmt.Errorf("SEND is only supported by OCPP 2.1")
			}
			switch err := session.startHandler(func(ctx context.Context) {
				s.handleSend(ctx, session, value)
			}); {
			case err == nil:
			case errors.Is(err, ErrHandlerQueueFull):
				s.dropQueuedSend(session, value)
			default:
				return err
			}
		case protocol.CallResult, protocol.CallError:
			// Unknown IDs are late or unsolicited responses and are ignored.
			session.resolve(message)
		default:
			return fmt.Errorf("unexpected inbound message type %d", message.Type())
		}
	}
}

// rejectQueuedCall answers a CALL that arrived while session's handler
// queue (Config.MaxQueuedHandlers) was already full — an honest CALLERROR
// instead of either blocking the read loop until room freed up (which
// would also stall reading this same session's own outbound Call
// responses) or admitting an unbounded number of goroutines waiting for a
// free handler slot.
func (s *Server) rejectQueuedCall(session *Session, call protocol.Call) {
	record := sessionLogRecord(session, LogWarn, LogCallRejected)
	record.MessageType, record.MessageID, record.Action, record.ErrorCode, record.Reason =
		protocol.CallType, call.ID, call.Action, GenericError, "handler_queue_full"
	s.log(session.Context(), record)
	session.recordMetric(MetricEvent{
		Type: MetricCallRejected, Identity: session.identity, Version: session.version,
		MessageType: protocol.CallType, Action: call.Action, ErrorCode: GenericError,
	})
	_ = session.send(session.Context(), protocol.CallError{
		ID: call.ID, Code: string(GenericError), Description: "too many pending calls", Details: json.RawMessage(`{}`),
	})
}

// dropQueuedSend is rejectQueuedCall's SEND counterpart: SEND is
// unconfirmed, so there's no response to send back — this only logs and
// records the drop, matching how an unregistered SEND action is already
// handled.
func (s *Server) dropQueuedSend(session *Session, send protocol.Send) {
	record := sessionLogRecord(session, LogWarn, LogSendDropped)
	record.MessageType, record.MessageID, record.Action, record.Reason =
		protocol.SendType, send.ID, send.Action, "handler_queue_full"
	s.log(session.Context(), record)
	session.recordMetric(MetricEvent{
		Type: MetricSendDropped, Identity: session.identity, Version: session.version,
		MessageType: protocol.SendType, Action: send.Action,
	})
}

func (s *Server) handleSend(ctx context.Context, session *Session, send protocol.Send) {
	start := time.Now()
	record := sessionLogRecord(session, LogDebug, LogSendReceived)
	record.MessageType, record.MessageID, record.Action = protocol.SendType, send.ID, send.Action
	metric := MetricEvent{Identity: session.identity, Version: session.version, MessageType: protocol.SendType, Action: send.Action}
	defer func() {
		if recover() != nil {
			record.Level, record.Event, record.Reason = LogError, LogSendDropped, "handler_panic"
			s.log(ctx, record)
			metric.Type, metric.Duration = MetricSendDropped, time.Since(start)
			session.recordMetric(metric)
		}
	}()
	s.log(ctx, record)
	metric.Type = MetricSendReceived
	session.recordMetric(metric)
	handler, ok := s.config.Router.lookup(session.version, send.Action, sendKind)
	if !ok {
		record.Level, record.Event, record.Reason = LogWarn, LogSendDropped, "action_not_registered"
		s.log(ctx, record)
		metric.Type, metric.Duration = MetricSendDropped, time.Since(start)
		session.recordMetric(metric)
		return
	}
	// OCPP 2.1 SEND is unconfirmed. Handler and validation failures are
	// intentionally not converted to CALLRESULT or CALLERROR.
	if _, err := handler(ctx, session, send.Payload); err != nil {
		record.Level, record.Event, record.Reason = LogWarn, LogSendDropped, "handler_or_validation_error"
		s.log(ctx, record)
		metric.Type, metric.Duration = MetricSendDropped, time.Since(start)
		session.recordMetric(metric)
		return
	}
	record.Level, record.Event, record.Reason = LogDebug, LogSendCompleted, ""
	s.log(ctx, record)
	metric.Type, metric.Duration = MetricSendCompleted, time.Since(start)
	session.recordMetric(metric)
}

func (s *Server) handleCall(ctx context.Context, session *Session, call protocol.Call) {
	start := time.Now()
	record := sessionLogRecord(session, LogDebug, LogCallReceived)
	record.MessageType, record.MessageID, record.Action = protocol.CallType, call.ID, call.Action
	metric := MetricEvent{Identity: session.identity, Version: session.version, MessageType: protocol.CallType, Action: call.Action}
	defer func() {
		if recover() != nil {
			record.Level, record.Event, record.ErrorCode, record.Reason = LogError, LogCallRejected, InternalError, "handler_panic"
			s.log(ctx, record)
			metric.Type, metric.Duration, metric.ErrorCode = MetricCallRejected, time.Since(start), InternalError
			session.recordMetric(metric)
			_ = session.send(ctx, protocol.CallError{
				ID: call.ID, Code: string(InternalError), Description: "internal error", Details: json.RawMessage(`{}`),
			})
		}
	}()
	s.log(ctx, record)
	metric.Type = MetricCallReceived
	session.recordMetric(metric)
	handler, ok := s.config.Router.lookup(session.version, call.Action, callKind)
	if !ok {
		record.Level, record.Event, record.ErrorCode, record.Reason = LogWarn, LogCallRejected, NotImplemented, "action_not_registered"
		s.log(ctx, record)
		metric.Type, metric.Duration, metric.ErrorCode = MetricCallRejected, time.Since(start), NotImplemented
		session.recordMetric(metric)
		_ = session.send(ctx, protocol.CallError{ID: call.ID, Code: string(NotImplemented), Description: "action is not registered"})
		return
	}
	response, err := handler(ctx, session, call.Payload)
	if err == nil {
		if err := session.result(ctx, call.ID, response); err != nil {
			record.Level, record.Event, record.ErrorCode, record.Reason = LogError, LogCallRejected, InternalError, "response_write_failed"
			s.log(ctx, record)
			metric.Type, metric.Duration, metric.ErrorCode = MetricCallRejected, time.Since(start), InternalError
			session.recordMetric(metric)
			// The CALL was received and routed to a handler, so the
			// Charging Station is waiting on a matching response — leaving
			// it silent (unlike every other rejection path in this
			// function) means it can only ever time out. Best-effort: if
			// the connection itself is what's broken, this send fails too
			// and is ignored, same as the other CALLERROR sends below.
			_ = session.send(ctx, protocol.CallError{
				ID: call.ID, Code: string(InternalError), Description: "internal error", Details: json.RawMessage(`{}`),
			})
			return
		}
		record.Level, record.Event = LogDebug, LogCallCompleted
		s.log(ctx, record)
		metric.Type, metric.Duration = MetricCallCompleted, time.Since(start)
		session.recordMetric(metric)
		return
	}
	callError := &CallError{Code: InternalError, Description: "internal error", Details: map[string]any{}}
	if !errors.As(err, &callError) {
		callError = &CallError{Code: InternalError, Description: "internal error", Details: map[string]any{}}
	}
	if !validErrorCode(session.version, callError.Code) {
		callError = &CallError{Code: InternalError, Description: "internal error", Details: map[string]any{}}
	}
	callError.Description = truncateRunes(callError.Description, protocol.MaxErrorDescriptionLength)
	details := callError.Details
	if details == nil {
		details = map[string]any{}
	}
	raw, marshalErr := jsonMarshal(details)
	if marshalErr != nil {
		raw = []byte(`{}`)
	}
	record.Level, record.Event, record.ErrorCode, record.Reason = LogWarn, LogCallRejected, callError.Code, "handler_error"
	s.log(ctx, record)
	metric.Type, metric.Duration, metric.ErrorCode = MetricCallRejected, time.Since(start), callError.Code
	session.recordMetric(metric)
	_ = session.send(ctx, protocol.CallError{ID: call.ID, Code: string(callError.Code), Description: callError.Description, Details: raw})
}

func validErrorCode(version protocol.Version, code ErrorCode) bool {
	switch code {
	case NotImplemented, NotSupported, InternalError, ProtocolError, SecurityError,
		PropertyConstraintViolation, TypeConstraintViolation, GenericError:
		return true
	case formatViolationCode(version), occurrenceConstraintCode(version):
		return true
	}
	if version == protocol.OCPP16 {
		return false
	}
	return code == MessageTypeNotSupported || code == RpcFrameworkError
}

func truncateRunes(value string, limit int) string {
	runes := []rune(value)
	if len(runes) <= limit {
		return value
	}
	return string(runes[:limit])
}

func disconnectReason(err error) string {
	switch {
	case errors.Is(err, ErrSessionReplaced):
		return "session_replaced"
	case errors.Is(err, ErrServerShutdown):
		return "server_shutdown"
	case errors.Is(err, ErrPongTimeout):
		return "pong_timeout"
	case errors.Is(err, ErrIdleTimeout):
		return "idle_timeout"
	case err == nil:
		return "closed"
	default:
		return "connection_error"
	}
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
