package csms

import (
	"context"
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/seanlee0923/ocpp/protocol"
)

// BenchmarkInboundCallRoundTrip measures a full CALL dispatch over a real
// loopback WebSocket connection: read, protocol.Decode, route, schema
// validate, unmarshal, handler, marshal, protocol.Encode, write. This is the
// realistic end-to-end number rather than a single hot-path function.
func BenchmarkInboundCallRoundTrip(b *testing.B) {
	router := NewRouter()
	if err := router.Handle(protocol.OCPP16, "Heartbeat", func(context.Context, *Session, json.RawMessage) (any, error) {
		return map[string]any{"currentTime": "2026-07-16T00:00:00Z"}, nil
	}); err != nil {
		b.Fatal(err)
	}
	server, err := New(Config{Router: router, Versions: []protocol.Version{protocol.OCPP16}})
	if err != nil {
		b.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()

	dialer := websocket.Dialer{Subprotocols: []string{string(protocol.OCPP16)}}
	conn, _, err := dialer.Dial("ws"+strings.TrimPrefix(httpServer.URL, "http")+"/BENCH-001", nil)
	if err != nil {
		b.Fatal(err)
	}
	defer conn.Close()

	frame := []byte(`[2,"bench","Heartbeat",{}]`)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := conn.WriteMessage(websocket.TextMessage, frame); err != nil {
			b.Fatal(err)
		}
		if _, _, err := conn.ReadMessage(); err != nil {
			b.Fatal(err)
		}
	}
}
