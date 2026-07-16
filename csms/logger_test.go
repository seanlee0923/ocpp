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

func TestLoggerPanicDoesNotStopServer(t *testing.T) {
	server, err := New(Config{Versions: []protocol.Version{protocol.OCPP16}, Logger: LoggerFunc(func(context.Context, LogRecord) { panic("logger failure") })})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()
	conn := dialTestStation(t, httpServer.URL, protocol.OCPP16)
	defer conn.Close()
	if server.SessionCount() != 1 {
		t.Fatalf("session count = %d", server.SessionCount())
	}
}
