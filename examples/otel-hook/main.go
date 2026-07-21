// Command otel-hook shows how to trace OCPP CALL/SEND handling with
// OpenTelemetry using nothing but the existing csms.Router middleware
// mechanism — csms has no dedicated tracing hook, and none is needed.
//
// Inbound tracing is a Middleware that starts a span and hands the
// resulting span-bearing context down to the registered handler, so any
// instrumented work the handler does becomes a child span. Outbound tracing
// (csms.Call) needs even less: the caller already owns the ctx it passes to
// Call, so it starts its own span before calling in — exactly like tracing
// any other function call, with no library involvement at all.
package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/trace"
	oteltrace "go.opentelemetry.io/otel/trace"

	"github.com/seanlee0923/ocpp/csms"
	"github.com/seanlee0923/ocpp/protocol"
	"github.com/seanlee0923/ocpp/v16"
)

// tracingMiddleware starts a span for every inbound CALL/SEND action and
// passes the span-bearing context to the wrapped handler. Unlike a
// Prometheus label (see examples/prometheus-hook), a span attribute does
// not create a persistent per-value time series, so it's fine to attach the
// charging station identity here.
func tracingMiddleware(tracer oteltrace.Tracer) csms.Middleware {
	return func(version protocol.Version, action string, next csms.Handler) csms.Handler {
		return func(ctx context.Context, session *csms.Session, payload json.RawMessage) (any, error) {
			ctx, span := tracer.Start(ctx, action, oteltrace.WithAttributes(
				attribute.String("ocpp.version", string(version)),
				attribute.String("ocpp.identity", session.Identity()),
			))
			defer span.End()

			result, err := next(ctx, session, payload)
			if err != nil {
				span.RecordError(err)
			}
			return result, err
		}
	}
}

// callResetTraced demonstrates outbound tracing: start a span using
// whatever ctx is on hand (here, the inbound handler's own traced ctx, so
// this becomes a child span) and pass the resulting ctx into csms.Call.
func callResetTraced(tracer oteltrace.Tracer, parent context.Context, session *csms.Session) {
	ctx, span := tracer.Start(parent, "csms.Call/Reset")
	defer span.End()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	if _, err := csms.Call[v16.ResetRequest, v16.ResetConfirmation](
		ctx, session, v16.ResetRequest{Type: v16.ResetRequestTypeSoft},
	); err != nil {
		span.RecordError(err)
	}
}

func main() {
	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		log.Fatal(err)
	}
	tracerProvider := trace.NewTracerProvider(
		trace.WithBatcher(exporter, trace.WithBatchTimeout(time.Second)),
	)
	defer func() { _ = tracerProvider.Shutdown(context.Background()) }()
	tracer := tracerProvider.Tracer("ocpp/examples/otel-hook")

	router := csms.NewRouter()
	router.Use(tracingMiddleware(tracer))
	if err := csms.Handle(router, func(
		ctx context.Context, session *csms.Session, _ v16.HeartbeatRequest,
	) (v16.HeartbeatConfirmation, error) {
		// Fire-and-forget outbound call, parented to this Heartbeat's span.
		go callResetTraced(tracer, ctx, session)
		return v16.HeartbeatConfirmation{CurrentTime: time.Now().UTC().Format(time.RFC3339)}, nil
	}); err != nil {
		log.Fatal(err)
	}

	server, err := csms.New(csms.Config{Router: router, Versions: []protocol.Version{protocol.OCPP16}})
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":8080", server))
}
