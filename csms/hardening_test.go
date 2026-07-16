package csms

import (
	"context"
	"encoding/json"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"ocpp-go/protocol"
	"ocpp-go/v16"
)

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
