// Command metrics-hook shows a minimal in-process Metrics collector and the
// Snapshot/Healthy status API, exposed together over a small status endpoint.
package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/seanlee0923/ocpp/csms"
)

type counters struct {
	mu     sync.Mutex
	counts map[csms.MetricEventType]int64
}

func newCounters() *counters {
	return &counters{counts: make(map[csms.MetricEventType]int64)}
}

func (c *counters) Record(_ context.Context, event csms.MetricEvent) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counts[event.Type]++
}

func (c *counters) snapshot() map[string]int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	result := make(map[string]int64, len(c.counts))
	for eventType, count := range c.counts {
		result[label(eventType)] = count
	}
	return result
}

func label(eventType csms.MetricEventType) string {
	switch eventType {
	case csms.MetricSessionConnected:
		return "session_connected"
	case csms.MetricSessionDisconnected:
		return "session_disconnected"
	case csms.MetricCallReceived:
		return "call_received"
	case csms.MetricCallCompleted:
		return "call_completed"
	case csms.MetricCallRejected:
		return "call_rejected"
	case csms.MetricSendReceived:
		return "send_received"
	case csms.MetricSendCompleted:
		return "send_completed"
	case csms.MetricSendDropped:
		return "send_dropped"
	case csms.MetricOutboundCallRejected:
		return "outbound_call_rejected"
	case csms.MetricOutboundCallSent:
		return "outbound_call_sent"
	case csms.MetricOutboundCallCompleted:
		return "outbound_call_completed"
	case csms.MetricOutboundCallFailed:
		return "outbound_call_failed"
	case csms.MetricOutboundCallTimeout:
		return "outbound_call_timeout"
	case csms.MetricOutboundCallCanceled:
		return "outbound_call_canceled"
	default:
		return "unknown"
	}
}

func main() {
	metrics := newCounters()
	server, err := csms.New(csms.Config{Metrics: metrics})
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", server)
	mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		status := struct {
			Healthy bool                `json:"healthy"`
			Server  csms.ServerSnapshot `json:"server"`
			Metrics map[string]int64    `json:"metrics"`
		}{
			Healthy: server.Healthy(),
			Server:  server.Snapshot(),
			Metrics: metrics.snapshot(),
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(status)
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
