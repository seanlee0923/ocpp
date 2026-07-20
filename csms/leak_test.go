package csms

import (
	"context"
	"errors"
	"fmt"
	"net/http/httptest"
	"runtime"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/seanlee0923/ocpp/protocol"
	"github.com/seanlee0923/ocpp/v16"
)

func TestRepeatedDisconnectReleasesSessionsAndGoroutines(t *testing.T) {
	server, err := New(Config{
		Versions:     []protocol.Version{protocol.OCPP16},
		PingInterval: 10 * time.Millisecond,
		PongTimeout:  100 * time.Millisecond,
		IdleTimeout:  time.Second,
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()
	runtime.GC()
	baseline := runtime.NumGoroutine()

	const cycles = 100
	for cycle := 0; cycle < cycles; cycle++ {
		identity := fmt.Sprintf("LEAK-%03d", cycle)
		conn := dialStationIdentity(t, httpServer.URL, protocol.OCPP16, identity)
		session := waitForSession(t, server, identity)
		if err := conn.UnderlyingConn().Close(); err != nil {
			t.Fatal(err)
		}
		select {
		case <-session.Done():
		case <-time.After(time.Second):
			t.Fatalf("cycle %d session did not close", cycle)
		}
		waitForNoSession(t, server, identity)
	}

	if server.SessionCount() != 0 {
		t.Fatalf("session count = %d after %d disconnects", server.SessionCount(), cycles)
	}
	waitForGoroutineBudget(t, baseline, 8)
}

func TestDisconnectReleasesPendingCallsAndReadLoop(t *testing.T) {
	const pendingCalls = 8
	connected := make(chan *Session, 1)
	var sequence atomic.Int64
	server, err := New(Config{
		Versions:        []protocol.Version{protocol.OCPP16},
		MaxPendingCalls: pendingCalls,
		UniqueIDGenerator: func() string {
			return fmt.Sprintf("pending-%d", sequence.Add(1))
		},
		OnConnect: func(session *Session) { connected <- session },
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()
	conn := dialTestStation(t, httpServer.URL, protocol.OCPP16)
	session := <-connected

	results := make(chan error, pendingCalls)
	for range pendingCalls {
		go func() {
			_, err := Call[v16.ResetRequest, v16.ResetConfirmation](context.Background(), session, v16.ResetRequest{Type: v16.ResetRequestTypeSoft})
			results <- err
		}()
	}
	deadline := time.Now().Add(time.Second)
	for pendingCallCount(session) != pendingCalls && time.Now().Before(deadline) {
		time.Sleep(time.Millisecond)
	}
	if got := pendingCallCount(session); got != pendingCalls {
		t.Fatalf("pending calls = %d, want %d", got, pendingCalls)
	}
	if err := conn.UnderlyingConn().Close(); err != nil {
		t.Fatal(err)
	}
	for range pendingCalls {
		if err := <-results; err == nil || errors.Is(err, context.Canceled) {
			t.Fatalf("pending call error after disconnect = %v", err)
		}
	}
	select {
	case <-session.Done():
	case <-time.After(time.Second):
		t.Fatal("session did not close")
	}
	select {
	case <-session.readDone:
	case <-time.After(time.Second):
		t.Fatal("read loop did not stop")
	}
	if got := pendingCallCount(session); got != 0 {
		t.Fatalf("pending calls retained after disconnect = %d", got)
	}
	if err := session.waitHandlers(context.Background()); err != nil {
		t.Fatal(err)
	}
}

func dialStationIdentity(t *testing.T, serverURL string, version protocol.Version, identity string) *websocket.Conn {
	t.Helper()
	dialer := websocket.Dialer{Subprotocols: []string{string(version)}}
	conn, _, err := dialer.Dial("ws"+strings.TrimPrefix(serverURL, "http")+"/"+identity, nil)
	if err != nil {
		t.Fatal(err)
	}
	return conn
}

func waitForSession(t *testing.T, server *Server, identity string) *Session {
	t.Helper()
	deadline := time.Now().Add(time.Second)
	for time.Now().Before(deadline) {
		if session, ok := server.Session(identity); ok {
			return session
		}
		time.Sleep(time.Millisecond)
	}
	t.Fatalf("session %q was not registered", identity)
	return nil
}

func waitForNoSession(t *testing.T, server *Server, identity string) {
	t.Helper()
	deadline := time.Now().Add(time.Second)
	for time.Now().Before(deadline) {
		if _, ok := server.Session(identity); !ok {
			return
		}
		time.Sleep(time.Millisecond)
	}
	t.Fatalf("session %q was retained", identity)
}

// TestShutdownStopsMetricDispatchWorkers proves Server.Shutdown actually
// stops and waits for the metricDispatchWorkers pool (started by New when
// Config.Metrics is set), not just the goroutines started per-connection.
// Without this, a process that repeatedly creates and Shuts down Servers
// (tests, tenant-per-Server setups, config-reload replacing a Server)
// leaks metricDispatchWorkers goroutines per Server, even though Shutdown
// is documented as terminal.
func TestShutdownStopsMetricDispatchWorkers(t *testing.T) {
	runtime.GC()
	baseline := runtime.NumGoroutine()

	server, err := New(Config{
		Versions: []protocol.Version{protocol.OCPP16},
		Metrics:  MetricsFunc(func(context.Context, MetricEvent) {}),
	})
	if err != nil {
		t.Fatal(err)
	}
	runtime.GC()
	if withWorkers := runtime.NumGoroutine(); withWorkers < baseline+metricDispatchWorkers {
		t.Fatalf("goroutines after New() with Metrics set = %d, want at least baseline(%d)+metricDispatchWorkers(%d)",
			withWorkers, baseline, metricDispatchWorkers)
	}

	if err := server.Shutdown(context.Background()); err != nil {
		t.Fatal(err)
	}
	waitForGoroutineBudget(t, baseline, 2)
}

// TestShutdownReturnsCtxErrIfMetricWorkersDoNotExit proves Shutdown gives
// up and returns ctx's error, instead of blocking forever, if a metric
// dispatch worker is stuck inside a hung Metrics.Record call and so never
// notices the stop signal — the complement of
// TestShutdownStopsMetricDispatchWorkers, which only exercises the case
// where every worker exits promptly.
func TestShutdownReturnsCtxErrIfMetricWorkersDoNotExit(t *testing.T) {
	block := make(chan struct{}) // never closed: Record blocks forever
	server, err := New(Config{
		Versions:  []protocol.Version{protocol.OCPP16},
		Metrics:   MetricsFunc(func(context.Context, MetricEvent) { <-block }),
		OnConnect: func(*Session) {},
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()
	defer close(block)

	// Connecting fires MetricSessionConnected, so at least one of the
	// metricDispatchWorkers pool goroutines is guaranteed to pick it up
	// and hang inside Record.
	conn := dialTestStation(t, httpServer.URL, protocol.OCPP16)
	defer conn.Close()
	waitForSession(t, server, "CP-001")

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	if err := server.Shutdown(ctx); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("Shutdown error = %v, want context.DeadlineExceeded", err)
	}
}

func waitForGoroutineBudget(t *testing.T, baseline, allowance int) {
	t.Helper()
	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		runtime.GC()
		current := runtime.NumGoroutine()
		if current <= baseline+allowance {
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
	t.Fatalf("goroutines did not stabilize: baseline=%d current=%d allowance=%d", baseline, runtime.NumGoroutine(), allowance)
}
