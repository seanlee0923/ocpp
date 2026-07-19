// Command prometheus-hook shows how to wire csms.Metrics into Prometheus
// without an official adapter package. The library does not ship one on
// purpose: MetricEvent.Identity (the charging station identity) is unsafe
// to use as a Prometheus label at scale, and whether it's safe depends on
// deployment size, which only the application knows. This example
// deliberately excludes Identity from every label — see
// csms.MetricEvent.Identity's doc comment for the full reasoning.
package main

import (
	"context"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/seanlee0923/ocpp/csms"
)

var (
	eventsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "ocpp_events_total",
		Help: "Total csms.MetricEvent observations, by type/action/version/error_code.",
	}, []string{"type", "action", "version", "error_code"})

	eventDurationSeconds = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "ocpp_event_duration_seconds",
		Help:    "Duration reported by terminal csms.MetricEvent observations.",
		Buckets: prometheus.DefBuckets,
	}, []string{"type", "action", "version"})
)

func recordEvent(_ context.Context, event csms.MetricEvent) {
	eventType := label(event.Type)
	version := string(event.Version)

	// No "identity" label here on purpose: a fleet of thousands of charging
	// stations would create thousands of series per metric. If you need
	// per-charger visibility for a small deployment, branch on
	// event.Identity here and write to a separate, explicitly-scoped
	// metric instead of adding it as a label on these fleet-wide ones.
	eventsTotal.With(prometheus.Labels{
		"type": eventType, "action": event.Action, "version": version,
		"error_code": string(event.ErrorCode),
	}).Inc()

	if event.Duration > 0 {
		eventDurationSeconds.With(prometheus.Labels{
			"type": eventType, "action": event.Action, "version": version,
		}).Observe(event.Duration.Seconds())
	}
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
	server, err := csms.New(csms.Config{Metrics: csms.MetricsFunc(recordEvent)})
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", server)
	mux.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(":8080", mux))
}
