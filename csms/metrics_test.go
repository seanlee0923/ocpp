package csms

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/seanlee0923/ocpp/protocol"
	"github.com/seanlee0923/ocpp/v16"
)

type metricRecorder struct {
	mu     sync.Mutex
	events []MetricEvent
	signal chan struct{}
}

func newMetricRecorder() *metricRecorder {
	return &metricRecorder{signal: make(chan struct{}, 64)}
}

func (r *metricRecorder) Record(_ context.Context, event MetricEvent) {
	r.mu.Lock()
	r.events = append(r.events, event)
	r.mu.Unlock()
	select {
	case r.signal <- struct{}{}:
	default:
	}
}

func (r *metricRecorder) snapshot() []MetricEvent {
	r.mu.Lock()
	defer r.mu.Unlock()
	return append([]MetricEvent(nil), r.events...)
}

// waitForMetric polls until match matches some recorded event, or fails the
// test after one second.
func waitForMetric(t *testing.T, r *metricRecorder, match func(MetricEvent) bool) MetricEvent {
	t.Helper()
	deadline := time.After(time.Second)
	for {
		for _, event := range r.snapshot() {
			if match(event) {
				return event
			}
		}
		select {
		case <-r.signal:
		case <-deadline:
			t.Fatalf("events = %#v", r.snapshot())
		}
	}
}

func TestMetricsPanicDoesNotStopServer(t *testing.T) {
	server, err := New(Config{
		Versions: []protocol.Version{protocol.OCPP16},
		Metrics:  MetricsFunc(func(context.Context, MetricEvent) { panic("metrics failure") }),
	})
	if err != nil {
		t.Fatal(err)
	}
	conn := dialTestStation(t, newTestHTTPServer(t, server).URL, protocol.OCPP16)
	defer conn.Close()
	waitForSession(t, server, "CP-001")
}

func TestCallReceivedAndCompletedEmitMetrics(t *testing.T) {
	router := NewRouter()
	if err := router.Handle(protocol.OCPP201, "Heartbeat", func(context.Context, *Session, json.RawMessage) (any, error) {
		return map[string]any{"currentTime": "2026-07-16T00:00:00Z"}, nil
	}); err != nil {
		t.Fatal(err)
	}
	recorder := newMetricRecorder()
	server, err := New(Config{Router: router, Versions: []protocol.Version{protocol.OCPP201}, Metrics: recorder})
	if err != nil {
		t.Fatal(err)
	}
	conn := dialTestStation(t, newTestHTTPServer(t, server).URL, protocol.OCPP201)
	defer conn.Close()
	if err := conn.WriteMessage(websocket.TextMessage, []byte(`[2,"metric-1","Heartbeat",{}]`)); err != nil {
		t.Fatal(err)
	}
	if _, _, err := conn.ReadMessage(); err != nil {
		t.Fatal(err)
	}

	waitForMetric(t, recorder, func(e MetricEvent) bool {
		return e.Type == MetricCallReceived && e.Identity == "CP-001" && e.Version == protocol.OCPP201 &&
			e.MessageType == protocol.CallType && e.Action == "Heartbeat"
	})
	completed := waitForMetric(t, recorder, func(e MetricEvent) bool {
		return e.Type == MetricCallCompleted && e.Action == "Heartbeat"
	})
	if completed.Duration <= 0 {
		t.Fatalf("Duration = %v, want > 0", completed.Duration)
	}
}

func TestCallRejectedEmitsMetricWithErrorCode(t *testing.T) {
	recorder := newMetricRecorder()
	server, err := New(Config{Versions: []protocol.Version{protocol.OCPP16}, Metrics: recorder})
	if err != nil {
		t.Fatal(err)
	}
	conn := dialTestStation(t, newTestHTTPServer(t, server).URL, protocol.OCPP16)
	defer conn.Close()
	if err := conn.WriteMessage(websocket.TextMessage, []byte(`[2,"metric-2","UnknownAction",{}]`)); err != nil {
		t.Fatal(err)
	}
	if _, _, err := conn.ReadMessage(); err != nil {
		t.Fatal(err)
	}
	waitForMetric(t, recorder, func(e MetricEvent) bool {
		return e.Type == MetricCallRejected && e.ErrorCode == NotImplemented && e.Action == "UnknownAction"
	})
}

func TestHandlerPanicEmitsCallRejectedMetric(t *testing.T) {
	router := NewRouter()
	if err := router.Handle(protocol.OCPP16, "Panic", func(context.Context, *Session, json.RawMessage) (any, error) {
		panic("application failure")
	}); err != nil {
		t.Fatal(err)
	}
	recorder := newMetricRecorder()
	server, err := New(Config{Router: router, Versions: []protocol.Version{protocol.OCPP16}, Metrics: recorder})
	if err != nil {
		t.Fatal(err)
	}
	conn := dialTestStation(t, newTestHTTPServer(t, server).URL, protocol.OCPP16)
	defer conn.Close()
	if err := conn.WriteMessage(websocket.TextMessage, []byte(`[2,"panic-metric","Panic",{}]`)); err != nil {
		t.Fatal(err)
	}
	if _, _, err := conn.ReadMessage(); err != nil {
		t.Fatal(err)
	}
	waitForMetric(t, recorder, func(e MetricEvent) bool {
		return e.Type == MetricCallRejected && e.ErrorCode == InternalError && e.Action == "Panic"
	})
}

func TestSendDroppedEmitsMetric(t *testing.T) {
	router := NewRouter()
	if err := router.Handle(protocol.OCPP21, "NotifyDisplayMessages", func(context.Context, *Session, json.RawMessage) (any, error) {
		return nil, errors.New("boom")
	}); err != nil {
		t.Fatal(err)
	}
	recorder := newMetricRecorder()
	server, err := New(Config{Router: router, Versions: []protocol.Version{protocol.OCPP21}, Metrics: recorder})
	if err != nil {
		t.Fatal(err)
	}
	conn := dialTestStation(t, newTestHTTPServer(t, server).URL, protocol.OCPP21)
	defer conn.Close()
	if err := conn.WriteMessage(websocket.TextMessage, []byte(`[6,"send-metric","NotifyDisplayMessages",{}]`)); err != nil {
		t.Fatal(err)
	}
	waitForMetric(t, recorder, func(e MetricEvent) bool {
		return e.Type == MetricSendDropped && e.MessageType == protocol.SendType && e.Action == "NotifyDisplayMessages"
	})
}

func TestOutboundCallSuccessEmitsMetrics(t *testing.T) {
	connected := make(chan *Session, 1)
	recorder := newMetricRecorder()
	server, err := New(Config{
		Versions:          []protocol.Version{protocol.OCPP16},
		UniqueIDGenerator: func() string { return "metric-out-1" },
		OnConnect:         func(session *Session) { connected <- session },
		Metrics:           recorder,
	})
	if err != nil {
		t.Fatal(err)
	}
	conn := dialTestStation(t, newTestHTTPServer(t, server).URL, protocol.OCPP16)
	defer conn.Close()
	session := <-connected
	go func() {
		_, _, _ = conn.ReadMessage()
		_ = conn.WriteMessage(websocket.TextMessage, []byte(`[3,"metric-out-1",{"status":"Accepted"}]`))
	}()
	if _, err := Call[v16.ResetRequest, v16.ResetConfirmation](
		context.Background(), session, v16.ResetRequest{Type: v16.ResetRequestTypeSoft},
	); err != nil {
		t.Fatal(err)
	}
	waitForMetric(t, recorder, func(e MetricEvent) bool { return e.Type == MetricOutboundCallSent && e.Action == "Reset" })
	completed := waitForMetric(t, recorder, func(e MetricEvent) bool {
		return e.Type == MetricOutboundCallCompleted && e.Action == "Reset"
	})
	if completed.Duration <= 0 {
		t.Fatalf("Duration = %v, want > 0", completed.Duration)
	}
}

func TestOutboundCallRemoteErrorEmitsMetric(t *testing.T) {
	connected := make(chan *Session, 1)
	recorder := newMetricRecorder()
	server, err := New(Config{
		Versions:          []protocol.Version{protocol.OCPP16},
		UniqueIDGenerator: func() string { return "metric-remote-err" },
		OnConnect:         func(session *Session) { connected <- session },
		Metrics:           recorder,
	})
	if err != nil {
		t.Fatal(err)
	}
	conn := dialTestStation(t, newTestHTTPServer(t, server).URL, protocol.OCPP16)
	defer conn.Close()
	session := <-connected
	go func() {
		_, _, _ = conn.ReadMessage()
		_ = conn.WriteMessage(websocket.TextMessage, []byte(`[4,"metric-remote-err","NotSupported","cannot reset",{"reason":"busy"}]`))
	}()
	if _, err := Call[v16.ResetRequest, v16.ResetConfirmation](
		context.Background(), session, v16.ResetRequest{Type: v16.ResetRequestTypeSoft},
	); err == nil {
		t.Fatal("expected error")
	}
	waitForMetric(t, recorder, func(e MetricEvent) bool {
		return e.Type == MetricOutboundCallFailed && e.ErrorCode == "NotSupported" && e.Action == "Reset"
	})
}

func TestOutboundCallTimeoutEmitsMetric(t *testing.T) {
	connected := make(chan *Session, 1)
	recorder := newMetricRecorder()
	server, err := New(Config{
		Versions:  []protocol.Version{protocol.OCPP16},
		OnConnect: func(session *Session) { connected <- session },
		Metrics:   recorder,
	})
	if err != nil {
		t.Fatal(err)
	}
	conn := dialTestStation(t, newTestHTTPServer(t, server).URL, protocol.OCPP16)
	defer conn.Close()
	session := <-connected
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()
	_, err = Call[v16.ResetRequest, v16.ResetConfirmation](ctx, session, v16.ResetRequest{Type: v16.ResetRequestTypeSoft})
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("error = %v, want context deadline exceeded", err)
	}
	waitForMetric(t, recorder, func(e MetricEvent) bool { return e.Type == MetricOutboundCallTimeout && e.Action == "Reset" })
}

func TestOutboundCallCanceledEmitsMetric(t *testing.T) {
	connected := make(chan *Session, 1)
	recorder := newMetricRecorder()
	server, err := New(Config{
		Versions:  []protocol.Version{protocol.OCPP16},
		OnConnect: func(session *Session) { connected <- session },
		Metrics:   recorder,
	})
	if err != nil {
		t.Fatal(err)
	}
	conn := dialTestStation(t, newTestHTTPServer(t, server).URL, protocol.OCPP16)
	defer conn.Close()
	session := <-connected
	// The station never responds; the caller cancels its own context
	// instead of letting the call time out.
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(20 * time.Millisecond)
		cancel()
	}()
	_, err = Call[v16.ResetRequest, v16.ResetConfirmation](ctx, session, v16.ResetRequest{Type: v16.ResetRequestTypeSoft})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("error = %v, want context canceled", err)
	}
	waitForMetric(t, recorder, func(e MetricEvent) bool { return e.Type == MetricOutboundCallCanceled && e.Action == "Reset" })
}

func TestOutboundCallOverflowEmitsRejectedMetric(t *testing.T) {
	const pendingLimit = 4
	connected := make(chan *Session, 1)
	var sequence atomic.Int64
	recorder := newMetricRecorder()
	server, err := New(Config{
		Versions:          []protocol.Version{protocol.OCPP16},
		MaxPendingCalls:   pendingLimit,
		UniqueIDGenerator: func() string { return fmt.Sprintf("overflow-%d", sequence.Add(1)) },
		OnConnect:         func(session *Session) { connected <- session },
		Metrics:           recorder,
	})
	if err != nil {
		t.Fatal(err)
	}
	conn := dialTestStation(t, newTestHTTPServer(t, server).URL, protocol.OCPP16)
	defer conn.Close()
	session := <-connected

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	results := make(chan error, pendingLimit)
	for range pendingLimit {
		go func() {
			_, err := Call[v16.ResetRequest, v16.ResetConfirmation](ctx, session, v16.ResetRequest{Type: v16.ResetRequestTypeSoft})
			results <- err
		}()
	}
	deadline := time.Now().Add(time.Second)
	for pendingCallCount(session) != pendingLimit && time.Now().Before(deadline) {
		time.Sleep(time.Millisecond)
	}
	if got := pendingCallCount(session); got != pendingLimit {
		t.Fatalf("pending calls = %d, want %d", got, pendingLimit)
	}

	_, err = Call[v16.ResetRequest, v16.ResetConfirmation](context.Background(), session, v16.ResetRequest{Type: v16.ResetRequestTypeSoft})
	if !errors.Is(err, ErrTooManyPendingCalls) {
		t.Fatalf("overflow error = %v, want ErrTooManyPendingCalls", err)
	}
	waitForMetric(t, recorder, func(e MetricEvent) bool { return e.Type == MetricOutboundCallRejected && e.Action == "Reset" })

	// Metric dispatch is asynchronous (see Session.recordMetric), so poll for
	// the expected count instead of assuming all Sent events already landed.
	countSent := func() int {
		count := 0
		for _, event := range recorder.snapshot() {
			if event.Type == MetricOutboundCallSent {
				count++
			}
		}
		return count
	}
	deadline = time.Now().Add(time.Second)
	for countSent() != pendingLimit && time.Now().Before(deadline) {
		time.Sleep(time.Millisecond)
	}
	if got := countSent(); got != pendingLimit {
		t.Fatalf("MetricOutboundCallSent count = %d, want %d (the rejected overflow call must not be counted as sent)", got, pendingLimit)
	}
	// Give any stray extra dispatch a brief window to arrive, then confirm
	// the count didn't grow further.
	time.Sleep(50 * time.Millisecond)
	if got := countSent(); got != pendingLimit {
		t.Fatalf("MetricOutboundCallSent count = %d after settling, want %d", got, pendingLimit)
	}

	cancel()
	for range pendingLimit {
		if err := <-results; !errors.Is(err, context.Canceled) {
			t.Fatalf("pending call error = %v, want context canceled", err)
		}
	}
}

func TestSnapshotReflectsRealSessions(t *testing.T) {
	server, err := New(Config{Versions: []protocol.Version{protocol.OCPP16}})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := newTestHTTPServer(t, server)
	conn1 := dialStationIdentity(t, httpServer.URL, protocol.OCPP16, "CP-snap-1")
	defer conn1.Close()
	conn2 := dialStationIdentity(t, httpServer.URL, protocol.OCPP16, "CP-snap-2")
	defer conn2.Close()
	waitForSession(t, server, "CP-snap-1")
	waitForSession(t, server, "CP-snap-2")

	snapshot := server.Snapshot()
	if snapshot.ActiveSessions != 2 {
		t.Fatalf("ActiveSessions = %d, want 2", snapshot.ActiveSessions)
	}
	if snapshot.ShuttingDown {
		t.Fatal("ShuttingDown = true, want false")
	}
	seen := map[string]bool{}
	for _, info := range snapshot.Sessions {
		seen[info.Identity] = true
		if info.State != SessionActive {
			t.Fatalf("session %s state = %v, want SessionActive", info.Identity, info.State)
		}
	}
	if !seen["CP-snap-1"] || !seen["CP-snap-2"] {
		t.Fatalf("Sessions = %#v, want CP-snap-1 and CP-snap-2", snapshot.Sessions)
	}

	_ = conn1.Close()
	waitForNoSession(t, server, "CP-snap-1")
	snapshot = server.Snapshot()
	if snapshot.ActiveSessions != 1 {
		t.Fatalf("ActiveSessions after close = %d, want 1", snapshot.ActiveSessions)
	}
}

func TestHealthyDuringShutdown(t *testing.T) {
	server, err := New(Config{Versions: []protocol.Version{protocol.OCPP16}})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := newTestHTTPServer(t, server)
	conn := dialTestStation(t, httpServer.URL, protocol.OCPP16)
	defer conn.Close()
	waitForSession(t, server, "CP-001")

	if !server.Healthy() {
		t.Fatal("Healthy() = false before Shutdown, want true")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		t.Fatal(err)
	}
	if server.Healthy() {
		t.Fatal("Healthy() = true after Shutdown, want false")
	}
}

// TestSlowMetricsDoesNotDelayCallCancellation guards against Session.recordMetric
// blocking Call()'s return. Before recordMetric dispatched Metrics.Record on
// its own goroutine, a slow implementation delayed Call() past its own
// context deadline, defeating the point of the callCtx.Done() branch.
func TestSlowMetricsDoesNotDelayCallCancellation(t *testing.T) {
	connected := make(chan *Session, 1)
	server, err := New(Config{
		Versions:  []protocol.Version{protocol.OCPP16},
		OnConnect: func(session *Session) { connected <- session },
		Metrics:   MetricsFunc(func(context.Context, MetricEvent) { time.Sleep(200 * time.Millisecond) }),
	})
	if err != nil {
		t.Fatal(err)
	}
	conn := dialTestStation(t, newTestHTTPServer(t, server).URL, protocol.OCPP16)
	defer conn.Close()
	session := <-connected

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()
	start := time.Now()
	_, err = Call[v16.ResetRequest, v16.ResetConfirmation](ctx, session, v16.ResetRequest{Type: v16.ResetRequestTypeSoft})
	elapsed := time.Since(start)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("error = %v, want context deadline exceeded", err)
	}
	if elapsed > 100*time.Millisecond {
		t.Fatalf("Call() took %v to return after a 20ms context timeout; a slow Metrics.Record (200ms/event) must not delay it", elapsed)
	}
}

// TestSlowMetricsDoesNotStallHandlerSlots guards against Session.recordMetric
// extending how long a handler goroutine occupies its handlerSlots slot.
// Before the async dispatch fix, a slow Metrics.Record kept handleCall from
// returning, which kept the sole handler slot (MaxConcurrentHandlers=1)
// occupied and stalled the read loop's next dispatch.
func TestSlowMetricsDoesNotStallHandlerSlots(t *testing.T) {
	router := NewRouter()
	if err := router.Handle(protocol.OCPP16, "Heartbeat", func(context.Context, *Session, json.RawMessage) (any, error) {
		return map[string]any{"currentTime": "2026-07-16T00:00:00Z"}, nil
	}); err != nil {
		t.Fatal(err)
	}
	server, err := New(Config{
		Router:                router,
		Versions:              []protocol.Version{protocol.OCPP16},
		MaxConcurrentHandlers: 1,
		Metrics:               MetricsFunc(func(context.Context, MetricEvent) { time.Sleep(150 * time.Millisecond) }),
	})
	if err != nil {
		t.Fatal(err)
	}
	conn := dialTestStation(t, newTestHTTPServer(t, server).URL, protocol.OCPP16)
	defer conn.Close()

	start := time.Now()
	for i := range 3 {
		if err := conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(`[2,"slow-metrics-%d","Heartbeat",{}]`, i))); err != nil {
			t.Fatal(err)
		}
		if _, _, err := conn.ReadMessage(); err != nil {
			t.Fatal(err)
		}
	}
	if elapsed := time.Since(start); elapsed > 300*time.Millisecond {
		t.Fatalf("3 sequential Heartbeats with MaxConcurrentHandlers=1 took %v; a slow Metrics.Record (150ms/event) must not serialize handler-slot release", elapsed)
	}
}

// TestOutboundCallExpiredDeadlineClassifiedAsTimeout guards against an
// already-expired context causing session.send to fail and Call classifying
// that as MetricOutboundCallFailed instead of MetricOutboundCallTimeout.
func TestOutboundCallExpiredDeadlineClassifiedAsTimeout(t *testing.T) {
	connected := make(chan *Session, 1)
	recorder := newMetricRecorder()
	server, err := New(Config{
		Versions:  []protocol.Version{protocol.OCPP16},
		OnConnect: func(session *Session) { connected <- session },
		Metrics:   recorder,
	})
	if err != nil {
		t.Fatal(err)
	}
	conn := dialTestStation(t, newTestHTTPServer(t, server).URL, protocol.OCPP16)
	defer conn.Close()
	session := <-connected

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(-time.Second))
	defer cancel()
	if _, err := Call[v16.ResetRequest, v16.ResetConfirmation](ctx, session, v16.ResetRequest{Type: v16.ResetRequestTypeSoft}); err == nil {
		t.Fatal("expected an error for an already-expired context")
	}
	waitForMetric(t, recorder, func(e MetricEvent) bool { return e.Type == MetricOutboundCallTimeout && e.Action == "Reset" })
}

// TestMetricSessionConnectedPrecedesOnConnectOutboundCall guards the fix that
// moved MetricSessionConnected recording before the read loop starts and
// before OnConnect runs, and confirms the documented "blocking csms.Call from
// OnConnect" pattern still works with that reordering.
func TestMetricSessionConnectedPrecedesOnConnectOutboundCall(t *testing.T) {
	recorder := newMetricRecorder()
	server, err := New(Config{
		Versions:          []protocol.Version{protocol.OCPP16},
		UniqueIDGenerator: func() string { return "onconnect-call" },
		Metrics:           recorder,
		OnConnect: func(session *Session) {
			_, _ = Call[v16.ResetRequest, v16.ResetConfirmation](context.Background(), session, v16.ResetRequest{Type: v16.ResetRequestTypeSoft})
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	conn := dialTestStation(t, newTestHTTPServer(t, server).URL, protocol.OCPP16)
	defer conn.Close()
	go func() {
		_, _, _ = conn.ReadMessage()
		_ = conn.WriteMessage(websocket.TextMessage, []byte(`[3,"onconnect-call",{"status":"Accepted"}]`))
	}()

	waitForMetric(t, recorder, func(e MetricEvent) bool { return e.Type == MetricOutboundCallCompleted && e.Action == "Reset" })

	found := false
	for _, event := range recorder.snapshot() {
		if event.Type == MetricSessionConnected {
			found = true
			break
		}
	}
	if !found {
		t.Fatal("MetricSessionConnected was never recorded even though OnConnect's blocking Call completed")
	}
}
