package csms

import (
	"context"
	"encoding/json"
	"net/http/httptest"
	"strconv"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/seanlee0923/ocpp/protocol"
	"github.com/seanlee0923/ocpp/v16"
)

// benchMetrics is the cheapest possible Metrics implementation: a single
// atomic increment, no I/O. It isolates the library's own dispatch overhead
// (goroutine spawn + semaphore) from any cost an application's real Metrics
// backend might add.
type benchMetrics struct{ count atomic.Int64 }

func (m *benchMetrics) Record(context.Context, MetricEvent) { m.count.Add(1) }

// BenchmarkInboundCallRoundTripWithMetrics is BenchmarkInboundCallRoundTrip
// (protocol/message_bench_test.go-style full loopback round trip) with
// Config.Metrics configured, to measure the overhead of the async metric
// dispatch added on the inbound CALL hot path.
func BenchmarkInboundCallRoundTripWithMetrics(b *testing.B) {
	router := NewRouter()
	if err := router.Handle(protocol.OCPP16, "Heartbeat", func(context.Context, *Session, json.RawMessage) (any, error) {
		return map[string]any{"currentTime": "2026-07-16T00:00:00Z"}, nil
	}); err != nil {
		b.Fatal(err)
	}
	server, err := New(Config{
		Router: router, Versions: []protocol.Version{protocol.OCPP16},
		Metrics: &benchMetrics{},
	})
	if err != nil {
		b.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()

	dialer := websocket.Dialer{Subprotocols: []string{string(protocol.OCPP16)}}
	conn, _, err := dialer.Dial("ws"+strings.TrimPrefix(httpServer.URL, "http")+"/BENCH-METRICS", nil)
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

// BenchmarkOutboundCallRoundTrip and BenchmarkOutboundCallRoundTripWithMetrics
// measure csms.Call's full round trip (request marshal, registerCall,
// session.send, station response, decode) over a real loopback WebSocket
// connection, with and without Config.Metrics configured.
func BenchmarkOutboundCallRoundTrip(b *testing.B) {
	benchmarkOutboundCallRoundTrip(b, nil)
}

func BenchmarkOutboundCallRoundTripWithMetrics(b *testing.B) {
	benchmarkOutboundCallRoundTrip(b, &benchMetrics{})
}

func benchmarkOutboundCallRoundTrip(b *testing.B, metrics Metrics) {
	connected := make(chan *Session, 1)
	var sequence atomic.Int64
	server, err := New(Config{
		Versions:          []protocol.Version{protocol.OCPP16},
		UniqueIDGenerator: func() string { return strconv.FormatInt(sequence.Add(1), 10) },
		OnConnect:         func(session *Session) { connected <- session },
		Metrics:           metrics,
	})
	if err != nil {
		b.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()

	dialer := websocket.Dialer{Subprotocols: []string{string(protocol.OCPP16)}}
	conn, _, err := dialer.Dial("ws"+strings.TrimPrefix(httpServer.URL, "http")+"/BENCH-OUTBOUND", nil)
	if err != nil {
		b.Fatal(err)
	}
	defer conn.Close()
	session := <-connected

	stationDone := make(chan struct{})
	go func() {
		defer close(stationDone)
		for {
			_, data, err := conn.ReadMessage()
			if err != nil {
				return
			}
			message, err := protocol.Decode(data)
			if err != nil {
				return
			}
			call, ok := message.(protocol.Call)
			if !ok {
				return
			}
			if err := conn.WriteMessage(websocket.TextMessage, []byte(`[3,"`+call.ID+`",{"status":"Accepted"}]`)); err != nil {
				return
			}
		}
	}()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := Call[v16.ResetRequest, v16.ResetConfirmation](
			context.Background(), session, v16.ResetRequest{Type: v16.ResetRequestTypeSoft},
		); err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	_ = conn.Close()
	<-stationDone
}
