package csms

import (
	"context"
	"encoding/json"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/seanlee0923/ocpp/protocol"
	"github.com/seanlee0923/ocpp/v16"
)

// TestOnConnectPanicDoesNotLeakSession proves a panicking OnConnect
// callback still lets ServeHTTP finish its cleanup sequence — waiting for
// the read loop, closing the session, unregistering it, and calling
// OnDisconnect — instead of the panic skipping everything after OnConnect
// and leaving the session (and its pingLoop/idleLoop/readLoop goroutines)
// registered forever.
func TestOnConnectPanicDoesNotLeakSession(t *testing.T) {
	disconnected := make(chan struct{}, 1)
	server, err := New(Config{
		Versions:     []protocol.Version{protocol.OCPP16},
		OnConnect:    func(*Session) { panic("boom") },
		OnDisconnect: func(*Session, error) { disconnected <- struct{}{} },
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := newTestHTTPServer(t, server)
	conn := dialTestStation(t, httpServer.URL, protocol.OCPP16)
	waitForSession(t, server, "CP-001")
	if err := conn.Close(); err != nil {
		t.Fatal(err)
	}

	select {
	case <-disconnected:
	case <-time.After(time.Second):
		t.Fatal("OnDisconnect was not called after OnConnect panicked")
	}
	waitForNoSession(t, server, "CP-001")
}

func TestOnConnectCanPerformBlockingCall(t *testing.T) {
	callResult := make(chan error, 1)
	server, err := New(Config{
		Versions: []protocol.Version{protocol.OCPP16},
		OnConnect: func(session *Session) {
			_, err := Call[v16.ResetRequest, v16.ResetConfirmation](
				context.Background(), session, v16.ResetRequest{Type: v16.ResetRequestTypeSoft},
			)
			callResult <- err
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := newTestHTTPServer(t, server)
	conn := dialTestStation(t, httpServer.URL, protocol.OCPP16)
	defer conn.Close()
	_, data, err := conn.ReadMessage()
	if err != nil {
		t.Fatal(err)
	}
	message, err := protocol.Decode(data)
	if err != nil {
		t.Fatal(err)
	}
	call := message.(protocol.Call)
	response := protocol.CallResult{ID: call.ID, Payload: json.RawMessage(`{"status":"Accepted"}`)}
	encoded, err := protocol.Encode(response)
	if err != nil {
		t.Fatal(err)
	}
	if err := conn.WriteMessage(websocket.TextMessage, encoded); err != nil {
		t.Fatal(err)
	}
	select {
	case err := <-callResult:
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(time.Second):
		t.Fatal("OnConnect CALL did not complete")
	}
}

func TestInboundHandlersApplyPerSessionBackpressure(t *testing.T) {
	router := NewRouter()
	started := make(chan struct{}, 2)
	release := make(chan struct{})
	var active atomic.Int32
	var maximum atomic.Int32
	router.Handle(protocol.OCPP16, "Heartbeat", func(context.Context, *Session, json.RawMessage) (any, error) {
		current := active.Add(1)
		for {
			old := maximum.Load()
			if current <= old || maximum.CompareAndSwap(old, current) {
				break
			}
		}
		started <- struct{}{}
		<-release
		active.Add(-1)
		return map[string]any{"currentTime": "2026-07-16T00:00:00Z"}, nil
	})
	server, err := New(Config{
		Router: router, Versions: []protocol.Version{protocol.OCPP16}, MaxConcurrentHandlers: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := newTestHTTPServer(t, server)
	conn := dialTestStation(t, httpServer.URL, protocol.OCPP16)
	defer conn.Close()
	for _, id := range []string{"1", "2"} {
		if err := conn.WriteMessage(websocket.TextMessage, []byte(`[2,"`+id+`","Heartbeat",{}]`)); err != nil {
			t.Fatal(err)
		}
	}
	<-started
	select {
	case <-started:
		t.Fatal("second handler started before the first handler released its slot")
	case <-time.After(30 * time.Millisecond):
	}
	close(release)
	for range 2 {
		if _, _, err := conn.ReadMessage(); err != nil {
			t.Fatal(err)
		}
	}
	if maximum.Load() != 1 {
		t.Fatalf("maximum concurrent handlers = %d", maximum.Load())
	}
}

func TestHandlerBackpressureDoesNotBlockOutboundCallResponse(t *testing.T) {
	router := NewRouter()
	started := make(chan struct{}, 1)
	release := make(chan struct{})
	if err := router.Handle(protocol.OCPP16, "Heartbeat", func(context.Context, *Session, json.RawMessage) (any, error) {
		select {
		case started <- struct{}{}:
		default:
		}
		<-release
		return map[string]any{"currentTime": "2026-07-16T00:00:00Z"}, nil
	}); err != nil {
		t.Fatal(err)
	}
	server, err := New(Config{
		Router: router, Versions: []protocol.Version{protocol.OCPP16},
		MaxConcurrentHandlers: 1, CallTimeout: 250 * time.Millisecond,
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := newTestHTTPServer(t, server)
	conn := dialTestStation(t, httpServer.URL, protocol.OCPP16)
	defer conn.Close()
	defer close(release)
	for _, id := range []string{"blocking-1", "queued-2"} {
		if err := conn.WriteMessage(websocket.TextMessage, []byte(`[2,"`+id+`","Heartbeat",{}]`)); err != nil {
			t.Fatal(err)
		}
	}
	select {
	case <-started:
	case <-time.After(time.Second):
		t.Fatal("first handler did not start")
	}
	session, ok := server.Session("CP-001")
	if !ok {
		t.Fatal("station session was not registered")
	}
	callResult := make(chan error, 1)
	go func() {
		_, err := Call[v16.ResetRequest, v16.ResetConfirmation](
			context.Background(), session, v16.ResetRequest{Type: v16.ResetRequestTypeSoft},
		)
		callResult <- err
	}()
	_, raw, err := conn.ReadMessage()
	if err != nil {
		t.Fatal(err)
	}
	message, err := protocol.Decode(raw)
	if err != nil {
		t.Fatal(err)
	}
	call, ok := message.(protocol.Call)
	if !ok || call.Action != "Reset" {
		t.Fatalf("outbound message = %#v, want Reset CALL", message)
	}
	encoded, err := protocol.Encode(protocol.CallResult{
		ID: call.ID, Payload: json.RawMessage(`{"status":"Accepted"}`),
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := conn.WriteMessage(websocket.TextMessage, encoded); err != nil {
		t.Fatal(err)
	}
	select {
	case err := <-callResult:
		if err != nil {
			t.Fatalf("outbound Call() failed while inbound handlers were saturated: %v", err)
		}
	case <-time.After(time.Second):
		t.Fatal("outbound Call() response was blocked by inbound handler backpressure")
	}
}

func TestHandlerPanicReturnsInternalErrorAndKeepsSessionAlive(t *testing.T) {
	router := NewRouter()
	if err := router.Handle(protocol.OCPP16, "Panic", func(context.Context, *Session, json.RawMessage) (any, error) {
		panic("application failure")
	}); err != nil {
		t.Fatal(err)
	}
	if err := router.Handle(protocol.OCPP16, "Heartbeat", func(context.Context, *Session, json.RawMessage) (any, error) {
		return map[string]any{"currentTime": "2026-07-16T00:00:00Z"}, nil
	}); err != nil {
		t.Fatal(err)
	}
	server, err := New(Config{Router: router, Versions: []protocol.Version{protocol.OCPP16}})
	if err != nil {
		t.Fatal(err)
	}
	conn := dialTestStation(t, newTestHTTPServer(t, server).URL, protocol.OCPP16)
	defer conn.Close()
	if err := conn.WriteMessage(websocket.TextMessage, []byte(`[2,"panic-1","Panic",{}]`)); err != nil {
		t.Fatal(err)
	}
	_, raw, err := conn.ReadMessage()
	if err != nil {
		t.Fatal(err)
	}
	message, err := protocol.Decode(raw)
	if err != nil {
		t.Fatal(err)
	}
	callError, ok := message.(protocol.CallError)
	if !ok || callError.Code != string(InternalError) {
		t.Fatalf("panic response = %#v", message)
	}
	if err := conn.WriteMessage(websocket.TextMessage, []byte(`[2,"alive-1","Heartbeat",{}]`)); err != nil {
		t.Fatal(err)
	}
	if _, _, err := conn.ReadMessage(); err != nil {
		t.Fatalf("session did not survive handler panic: %v", err)
	}
}

func TestShutdownCancelsAndWaitsForHandlers(t *testing.T) {
	router := NewRouter()
	started := make(chan struct{})
	finished := make(chan struct{})
	router.Handle(protocol.OCPP16, "Heartbeat", func(ctx context.Context, _ *Session, _ json.RawMessage) (any, error) {
		close(started)
		<-ctx.Done()
		close(finished)
		return nil, context.Cause(ctx)
	})
	server, err := New(Config{Router: router, Versions: []protocol.Version{protocol.OCPP16}})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := newTestHTTPServer(t, server)
	conn := dialTestStation(t, httpServer.URL, protocol.OCPP16)
	defer conn.Close()
	if err := conn.WriteMessage(websocket.TextMessage, []byte(`[2,"1","Heartbeat",{}]`)); err != nil {
		t.Fatal(err)
	}
	<-started
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		t.Fatal(err)
	}
	select {
	case <-finished:
	default:
		t.Fatal("Shutdown returned before the handler finished")
	}
}

func TestPublicCloseReasonDoesNotExposeInternalError(t *testing.T) {
	_, reason := publicCloseReason(context.DeadlineExceeded)
	if reason != "session closed" {
		t.Fatalf("public reason = %q", reason)
	}
}

func newTestHTTPServer(t *testing.T, server *Server) *httptest.Server {
	t.Helper()
	httpServer := httptest.NewServer(server)
	t.Cleanup(httpServer.Close)
	return httpServer
}
