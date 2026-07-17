package csms

import (
	"context"
	"encoding/json"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/seanlee0923/ocpp/protocol"
)

func TestStructuredLoggerReceivesProtocolMetadata(t *testing.T) {
	router := NewRouter()
	router.Handle(protocol.OCPP201, "Heartbeat", func(context.Context, *Session, json.RawMessage) (any, error) {
		return map[string]any{"currentTime": "2026-07-16T00:00:00Z"}, nil
	})
	var mu sync.Mutex
	var records []LogRecord
	logged := make(chan struct{}, 8)
	server, err := New(Config{
		Router: router, Versions: []protocol.Version{protocol.OCPP201},
		Logger: LoggerFunc(func(_ context.Context, record LogRecord) {
			mu.Lock()
			records = append(records, record)
			mu.Unlock()
			logged <- struct{}{}
		}),
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()
	conn := dialTestStation(t, httpServer.URL, protocol.OCPP201)
	if err := conn.WriteMessage(websocket.TextMessage, []byte(`[2,"log-1","Heartbeat",{}]`)); err != nil {
		t.Fatal(err)
	}
	if _, _, err := conn.ReadMessage(); err != nil {
		t.Fatal(err)
	}
	_ = conn.Close()

	deadline := time.After(time.Second)
	for {
		mu.Lock()
		found := false
		for _, record := range records {
			if record.Event == LogCallCompleted && record.Identity == "CP-001" && record.Version == protocol.OCPP201 && record.MessageType == protocol.CallType && record.MessageID == "log-1" && record.Action == "Heartbeat" {
				found = true
			}
		}
		mu.Unlock()
		if found {
			return
		}
		select {
		case <-logged:
		case <-deadline:
			t.Fatalf("records = %#v", records)
		}
	}
}

func TestSendSuccessLogsCompletionWithoutWireResponse(t *testing.T) {
	router := NewRouter()
	handled := make(chan struct{}, 1)
	router.Handle(protocol.OCPP21, "NotifyDisplayMessages", func(context.Context, *Session, json.RawMessage) (any, error) {
		handled <- struct{}{}
		return nil, nil
	})
	var mu sync.Mutex
	var records []LogRecord
	logged := make(chan struct{}, 8)
	server, err := New(Config{
		Router: router, Versions: []protocol.Version{protocol.OCPP21},
		Logger: LoggerFunc(func(_ context.Context, record LogRecord) {
			mu.Lock()
			records = append(records, record)
			mu.Unlock()
			logged <- struct{}{}
		}),
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()
	conn := dialTestStation(t, httpServer.URL, protocol.OCPP21)
	defer conn.Close()
	if err := conn.WriteMessage(websocket.TextMessage, []byte(`[6,"send-1","NotifyDisplayMessages",{}]`)); err != nil {
		t.Fatal(err)
	}
	select {
	case <-handled:
	case <-time.After(time.Second):
		t.Fatal("SEND handler did not run")
	}

	deadline := time.After(time.Second)
	for {
		mu.Lock()
		found := false
		for _, record := range records {
			if record.Event == LogSendCompleted && record.Identity == "CP-001" && record.Version == protocol.OCPP21 && record.MessageType == protocol.SendType && record.MessageID == "send-1" && record.Action == "NotifyDisplayMessages" {
				found = true
			}
		}
		mu.Unlock()
		if found {
			break
		}
		select {
		case <-logged:
		case <-deadline:
			t.Fatalf("records = %#v", records)
		}
	}

	// A successful SEND must not write anything back on the wire. A
	// subsequent CALL proves the connection is still usable and that no
	// leftover SEND response is sitting ahead of it.
	if err := conn.WriteMessage(websocket.TextMessage, []byte(`[2,"after-send","UnknownAction",{}]`)); err != nil {
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
	if message.Type() != protocol.CallErrorType {
		t.Fatalf("response type = %d, want CALLERROR for the CALL after SEND", message.Type())
	}
	if id := message.(protocol.CallError).ID; id != "after-send" {
		t.Fatalf("response ID = %q, want %q (leftover SEND response on the wire?)", id, "after-send")
	}
}

func TestLoggerPanicDoesNotStopServer(t *testing.T) {
	server, err := New(Config{Versions: []protocol.Version{protocol.OCPP16}, Logger: LoggerFunc(func(context.Context, LogRecord) { panic("logger failure") })})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()
	conn := dialTestStation(t, httpServer.URL, protocol.OCPP16)
	defer conn.Close()
	// The client-side dial returns as soon as it reads the HTTP 101
	// response; the server registers the session a few lines later. Poll
	// instead of asserting SessionCount() immediately.
	waitForSession(t, server, "CP-001")
}
