package station

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"math"
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
const defaultWriteTimeout = 10 * time.Second
const defaultReadLimit = 1 << 20 // 1MB, matching csms.Config.ReadLimit's default

// disconnectHandlerGrace bounds how long Run waits, after a connection
// ends, for that connection's in-flight dispatch-spawned handler
// goroutines to actually exit before calling OnDisconnect and moving on to
// the next reconnect attempt. conn.ctx is already canceled by then, so a
// well-behaved handler (one that respects ctx, as documented) typically
// returns almost immediately — this grace period only narrows, not
// eliminates, the window where a fast reconnect could otherwise fire
// OnConnect (and a fresh handler invocation) while a handler from the
// just-closed connection is still running, matching csms.Server's
// equivalent disconnectHandlerGrace. Deliberately short: this fires on
// every disconnect, not just a graceful shutdown.
const disconnectHandlerGrace = 5 * time.Second

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
	// MaxQueuedHandlers bounds how many inbound CALLs may be admitted
	// (running or waiting for a free MaxConcurrentHandlers slot) at once.
	// Default: 4x MaxConcurrentHandlers, matching csms.Config.MaxQueuedHandlers.
	// Once full, a further inbound CALL gets an immediate CALLERROR
	// instead of an unbounded number of goroutines piling up waiting for a
	// slot.
	MaxQueuedHandlers int
	HandshakeTimeout  time.Duration
	// WriteTimeout bounds how long a single WebSocket write (an outbound
	// Call's CALL frame, or an inbound CALL's CALLRESULT/CALLERROR
	// response) may take, matching csms.Config.WriteTimeout. Default: 10s.
	// Without this, a stalled write (e.g. the CSMS not reading) could block
	// indefinitely regardless of any context deadline the caller passed to
	// Call, since a context deadline alone does nothing to bound a
	// blocking WriteMessage call — this is the actual socket-level
	// deadline that does.
	WriteTimeout time.Duration
	// ReadLimit bounds the size in bytes of a single inbound WebSocket
	// message, matching csms.Config.ReadLimit. Default: 1MB. Without this,
	// a malicious or misbehaving CSMS sending an oversized message could
	// make the Station buffer an unbounded amount of memory reading it.
	ReadLimit         int64
	OnConnect         func(*Station)
	OnDisconnect      func(*Station, error)
	UniqueIDGenerator func() string // default uuid.NewString
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
	if !config.Version.Valid() {
		return fmt.Errorf("station: unsupported OCPP version %q", config.Version)
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
	if config.MaxQueuedHandlers < 0 {
		return fmt.Errorf("station: MaxQueuedHandlers must not be negative")
	}
	if config.HandshakeTimeout < 0 {
		return fmt.Errorf("station: HandshakeTimeout must not be negative")
	}
	if config.WriteTimeout < 0 {
		return fmt.Errorf("station: WriteTimeout must not be negative")
	}
	if config.ReadLimit < 0 {
		return fmt.Errorf("station: ReadLimit must not be negative")
	}
	if config.ReconnectPolicy != nil {
		if config.ReconnectPolicy.InitialDelay < 0 || config.ReconnectPolicy.MaxDelay < 0 {
			return fmt.Errorf("station: ReconnectPolicy delays must not be negative")
		}
		if multiplier := config.ReconnectPolicy.Multiplier; math.IsNaN(multiplier) || math.IsInf(multiplier, 0) {
			return fmt.Errorf("station: ReconnectPolicy.Multiplier must not be NaN or infinite")
		}
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

// normalizeCallError makes a handler-returned *CallError's fields safe to
// encode as a wire CALLERROR before ever attempting to send it, mirroring
// how csms's handleCall validates/truncates a handler-returned CallError.
// Without this, a handler that returns an invalid code (empty or over
// protocol.MaxErrorCodeLength), an overlong description, or Details that
// isn't a JSON object makes protocol.Encode fail — and since dispatch's
// conn.send call ignores that error (there is nothing else useful to do
// with it at that point), the CSMS would be left with no response at all
// for a CALL it's waiting on, rather than a well-formed one.
func normalizeCallError(version protocol.Version, code, description string, details json.RawMessage) (string, string, json.RawMessage) {
	if !validCallErrorCode(version, code) {
		code = "InternalError"
	}
	description = truncateRunes(description, protocol.MaxErrorDescriptionLength)
	if len(details) == 0 {
		return code, description, json.RawMessage(`{}`)
	}
	var object map[string]json.RawMessage
	if json.Unmarshal(details, &object) != nil || object == nil {
		details = json.RawMessage(`{}`)
	}
	return code, description, details
}

func validCallErrorCode(version protocol.Version, code string) bool {
	switch code {
	case "NotImplemented", "NotSupported", "InternalError", "ProtocolError", "SecurityError",
		"PropertyConstraintViolation", "TypeConstraintViolation", "GenericError":
		return true
	}
	if version == protocol.OCPP16 {
		return code == "FormationViolation" || code == "OccurenceConstraintViolation"
	}
	return code == "FormatViolation" || code == "OccurrenceConstraintViolation" ||
		code == "MessageTypeNotSupported" || code == "RpcFrameworkError"
}

func truncateRunes(value string, limit int) string {
	runes := []rune(value)
	if len(runes) <= limit {
		return value
	}
	return string(runes[:limit])
}

// stationConn is the state of one underlying WebSocket connection. A fresh
// stationConn replaces the old one on every reconnect; pending calls do not
// survive a reconnect, matching how csms.Session.closeWithError cancels
// pending calls on close. ctx is canceled when the connection closes, so
// inbound-CALL handlers (which run with ctx, not context.Background()) can
// observe disconnection instead of running indefinitely, matching how
// csms.Session runs handlers with its own session-scoped context.
type stationConn struct {
	conn         *websocket.Conn
	ctx          context.Context
	cancel       context.CancelFunc
	writeMu      sync.Mutex
	writeTimeout time.Duration
	pendingMu    sync.Mutex
	pending      map[string]chan callOutcome
	closed       chan struct{}
	closeOnce    sync.Once
	handlerWG    sync.WaitGroup
}

// newStationConn derives conn's ctx from parent (Run's ctx), not
// context.Background(): handler invocations should inherit any values the
// caller attached to Run's ctx, and observe the caller canceling Run
// directly rather than only through runConnection's watcher goroutine
// forcing the socket closed.
func newStationConn(parent context.Context, conn *websocket.Conn, writeTimeout time.Duration) *stationConn {
	ctx, cancel := context.WithCancel(parent)
	return &stationConn{conn: conn, ctx: ctx, cancel: cancel, writeTimeout: writeTimeout, pending: make(map[string]chan callOutcome), closed: make(chan struct{})}
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

// waitHandlers blocks until every dispatch-spawned handler goroutine for
// this connection has exited, or ctx expires first — mirroring
// csms.Session.waitHandlers.
func (c *stationConn) waitHandlers(ctx context.Context) error {
	done := make(chan struct{})
	go func() {
		c.handlerWG.Wait()
		close(done)
	}()
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
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

// send writes message to the wire, bounding the actual WriteMessage call
// via SetWriteDeadline (ctx's deadline if it has one and it's sooner, else
// c.writeTimeout) — matching csms.Session.send. A context deadline alone
// does nothing to bound a blocking WriteMessage call; this is the real
// socket-level deadline that does, so a stalled write (e.g. the CSMS not
// reading) can't block ctx's caller indefinitely regardless of what
// deadline ctx carries.
func (c *stationConn) send(ctx context.Context, message protocol.Message) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	data, err := protocol.Encode(message)
	if err != nil {
		return err
	}
	c.writeMu.Lock()
	defer c.writeMu.Unlock()
	if err := ctx.Err(); err != nil {
		return err
	}
	select {
	case <-c.closed:
		return ErrNotConnected
	default:
	}
	deadline, hasDeadline := ctx.Deadline()
	if c.writeTimeout > 0 {
		configured := time.Now().Add(c.writeTimeout)
		if !hasDeadline || configured.Before(deadline) {
			deadline, hasDeadline = configured, true
		}
	}
	if hasDeadline {
		if err := c.conn.SetWriteDeadline(deadline); err != nil {
			return err
		}
	} else {
		_ = c.conn.SetWriteDeadline(time.Time{})
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
	writeTimeout time.Duration
	readLimit    int64
	handlersMu   sync.RWMutex
	handlers     map[string]func(context.Context, json.RawMessage) (json.RawMessage, error)
	handlerSlots chan struct{}
	pendingSlots chan struct{}
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
	maxQueuedHandlers := config.MaxQueuedHandlers
	if maxQueuedHandlers == 0 {
		maxQueuedHandlers = 4 * maxConcurrentHandlers
	}
	writeTimeout := config.WriteTimeout
	if writeTimeout == 0 {
		writeTimeout = defaultWriteTimeout
	}
	readLimit := config.ReadLimit
	if readLimit == 0 {
		readLimit = defaultReadLimit
	}
	return &Station{
		config:       config,
		uniqueIDGen:  generator,
		writeTimeout: writeTimeout,
		readLimit:    readLimit,
		handlers:     make(map[string]func(context.Context, json.RawMessage) (json.RawMessage, error)),
		handlerSlots: make(chan struct{}, maxConcurrentHandlers),
		pendingSlots: make(chan struct{}, maxQueuedHandlers),
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

// callOnConnect isolates a panicking OnConnect from the rest of Run's loop.
// Without this, a panic here has nowhere to be recovered (Run is typically
// launched by the caller as `go st.Run(ctx)`, and recover only works within
// the panicking goroutine's own call stack) and crashes the whole process,
// not just this Station — worse than merely leaving state stuck at
// Connected, since setState(Connected) above has already run and nothing
// downstream would otherwise get a chance to correct it.
func (s *Station) callOnConnect() {
	if s.config.OnConnect == nil {
		return
	}
	defer func() { _ = recover() }()
	s.config.OnConnect(s)
}

// callOnDisconnect mirrors callOnConnect for the disconnect hook.
func (s *Station) callOnDisconnect(cause error) {
	if s.config.OnDisconnect == nil {
		return
	}
	defer func() { _ = recover() }()
	s.config.OnDisconnect(s, cause)
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
			// Start the read loop before OnConnect. A connection hook commonly
			// sends BootNotification synchronously; its CALLRESULT can only be
			// resolved while the read loop is running.
			connectionResult := make(chan error, 1)
			go func() { connectionResult <- s.runConnection(ctx, conn) }()
			s.callOnConnect()

			attemptErr = <-connectionResult
			conn.close(attemptErr)
			graceCtx, graceCancel := context.WithTimeout(context.Background(), disconnectHandlerGrace)
			_ = conn.waitHandlers(graceCtx)
			graceCancel()
			s.setConn(nil)
			s.setState(Disconnected)
			s.callOnDisconnect(attemptErr)

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
	conn.SetReadLimit(s.readLimit)
	return newStationConn(ctx, conn, s.writeTimeout), nil
}

// runConnection runs the read loop until it exits on its own (connection
// error) or ctx/Stop fires, in which case it forces conn closed so the
// blocking ReadMessage call in readLoop unblocks. Returns the read loop's
// error; callers should still check ctx/stopCh afterward, since a forced
// close surfaces as a generic network error here, not ctx.Err().
func (s *Station) runConnection(ctx context.Context, conn *stationConn) (err error) {
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
	defer close(done)
	// Run is typically launched by the caller as `go st.Run(ctx)`, with no
	// recover of its own anywhere up that goroutine's stack — an
	// unrecovered panic in readLoop (ReadMessage, Decode, ...) would crash
	// the whole process instead of just ending this one connection.
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("station: read loop panicked: %v", r)
		}
	}()
	err = s.readLoop(conn)
	return
}

func (s *Station) readLoop(conn *stationConn) error {
	for {
		messageType, data, err := conn.conn.ReadMessage()
		if err != nil {
			return err
		}
		if messageType != websocket.TextMessage {
			return fmt.Errorf("OCPP only accepts WebSocket text messages")
		}
		message, err := protocol.Decode(data)
		if err != nil {
			// A well-formed envelope naming an unsupported message type ID
			// is the one Decode failure that isn't a fatal framing error:
			// the ID survived parsing, so the station can name it in a
			// proper CALLERROR and keep the connection open, mirroring
			// csms.Server's own read loop. Like csms, this is only valid
			// from OCPP 2.0.1 onward — 1.6 has no MessageTypeNotSupported
			// CALLERROR code to send.
			var unsupported *protocol.UnsupportedMessageTypeError
			if s.config.Version != protocol.OCPP16 && errors.As(err, &unsupported) {
				_ = conn.send(conn.ctx, protocol.CallError{
					ID: unsupported.ID, Code: "MessageTypeNotSupported",
					Description: "message type is not supported", Details: json.RawMessage(`{}`),
				})
				continue
			}
			// Anything else (malformed JSON, wrong array shape, invalid
			// UTF-8, ...) is a genuinely unrecoverable framing failure, not
			// a message the station simply disagrees with — matching
			// csms.Server's own read loop, which always closes the
			// connection for this case too, independent of OCPP version.
			// Silently ignoring it instead would leave the CSMS waiting on
			// a response that's never coming, with no signal anything went
			// wrong.
			return err
		}
		switch value := message.(type) {
		case protocol.CallResult, protocol.CallError:
			conn.resolve(message)
		case protocol.Call:
			s.dispatch(conn, value)
		}
	}
}

// dispatch admits and runs a registered handler for an inbound CALL. Like
// csms.Session.startHandler, admission (the pendingSlots check just below)
// never blocks and only bounds how many CALLs may be running or waiting
// for a free Config.MaxConcurrentHandlers slot at once
// (Config.MaxQueuedHandlers) — dispatch itself must never block, since
// this is the Station's one and only connection, and its read loop also
// resolves this same connection's outbound Call responses (conn.resolve).
// Blocking here on handler-slot availability would let a burst of inbound
// CALLs starve/spuriously time out unrelated pending outbound Calls queued
// behind them in the same byte stream. A panicking handler is recovered
// and reported as an InternalError CALLERROR instead of crashing the
// Station, matching csms.Server's handleCall panic recovery.
func (s *Station) dispatch(conn *stationConn, call protocol.Call) {
	s.handlersMu.RLock()
	handler, ok := s.handlers[call.Action]
	s.handlersMu.RUnlock()
	if !ok {
		_ = conn.send(conn.ctx, protocol.CallError{
			ID: call.ID, Code: "NotImplemented",
			Description: "no station handler registered for " + call.Action, Details: json.RawMessage(`{}`),
		})
		return
	}
	select {
	case s.pendingSlots <- struct{}{}:
	default:
		// Every pendingSlots slot is already running or waiting for a free
		// handlerSlots slot: reject this one immediately instead of
		// piling up an unbounded number of goroutines (each holding its
		// own stack and the CALL's payload) waiting behind them.
		_ = conn.send(conn.ctx, protocol.CallError{
			ID: call.ID, Code: "GenericError",
			Description: "too many pending calls", Details: json.RawMessage(`{}`),
		})
		return
	}
	conn.handlerWG.Add(1)
	go func() {
		defer conn.handlerWG.Done()
		defer func() { <-s.pendingSlots }()
		select {
		case s.handlerSlots <- struct{}{}:
		case <-conn.closed:
			return
		}
		// A per-call bound, not just conn.ctx directly: conn.ctx only ends
		// when the whole connection does, so without this a handler that
		// never returns permanently occupies its handlerSlots slot,
		// matching why csms.Session.startHandler bounds its own handler
		// invocations by Config.CallTimeout instead of only the session's
		// ctx. This ctx also covers the response sends below (not just the
		// handler call), matching csms.Server.handleCall's own scoping.
		ctx := conn.ctx
		if s.config.CallTimeout > 0 {
			var cancel context.CancelFunc
			ctx, cancel = context.WithTimeout(conn.ctx, s.config.CallTimeout)
			defer cancel()
		}
		// Recover before releasing the slot: a single closure makes that
		// order explicit in the code instead of relying on two separate
		// defers' LIFO registration order, which would silently keep
		// working even if someone later reordered the two statements.
		defer func() {
			if recover() != nil {
				_ = conn.send(ctx, protocol.CallError{ID: call.ID, Code: "InternalError", Description: "internal error", Details: json.RawMessage(`{}`)})
			}
			<-s.handlerSlots
		}()
		payload, err := handler(ctx, call.Payload)
		if err != nil {
			var callErr *CallError
			if errors.As(err, &callErr) {
				code, description, details := normalizeCallError(s.config.Version, callErr.Code, callErr.Description, callErr.Details)
				_ = conn.send(ctx, protocol.CallError{ID: call.ID, Code: code, Description: description, Details: details})
				return
			}
			// A plain (non-*CallError) handler error is never sent as-is:
			// its Error() string is application-internal detail (DB
			// errors, file paths, internal hostnames, ...) with no reason
			// to cross the wire to the CSMS, matching how csms's handleCall
			// normalizes an unrecognized handler error to a fixed
			// "internal error" instead of relaying err.Error(). A handler
			// that wants to disclose a specific code/description to the
			// CSMS should return *CallError, not rely on this fallback.
			_ = conn.send(ctx, protocol.CallError{ID: call.ID, Code: "InternalError", Description: "internal error", Details: json.RawMessage(`{}`)})
			return
		}
		_ = conn.send(ctx, protocol.CallResult{ID: call.ID, Payload: payload})
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
