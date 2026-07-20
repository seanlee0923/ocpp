package csms

import (
	"context"
	"sync"
	"time"

	"github.com/seanlee0923/ocpp/protocol"
)

type MetricEventType uint8

const (
	// MetricSessionConnected is recorded as early as possible in the
	// connection lifecycle, before the read loop starts and before
	// Config.OnConnect runs. It is still not guaranteed to be the first
	// event Metrics.Record observes for a session: dispatch is async and
	// unordered (see the Metrics doc comment), so a Metrics implementation
	// must not assume Connected always precedes every other event for the
	// same identity.
	MetricSessionConnected MetricEventType = iota + 1
	// MetricSessionDisconnected carries Duration equal to the session's
	// total lifetime (time since connect).
	MetricSessionDisconnected
	MetricCallReceived
	// MetricCallCompleted carries Duration for validation, handler
	// execution, and the CALLRESULT write.
	MetricCallCompleted
	// MetricCallRejected carries Duration and ErrorCode.
	MetricCallRejected
	MetricSendReceived
	MetricSendCompleted
	MetricSendDropped
	// MetricOutboundCallRejected fires when Call is rejected before a CALL
	// frame is ever sent because MaxPendingCalls was reached. Duration is
	// always zero.
	MetricOutboundCallRejected
	// MetricOutboundCallSent fires once the CALL frame has been written.
	// Duration is always zero.
	MetricOutboundCallSent
	// MetricOutboundCallCompleted carries Duration measured from just after
	// the call was admitted into the pending-call table (i.e. it includes
	// CALL marshal/write time, not strictly the interval after
	// MetricOutboundCallSent) to a valid CALLRESULT.
	MetricOutboundCallCompleted
	// MetricOutboundCallFailed covers a remote CALLERROR, an undecodable or
	// invalid confirmation, a transport write failure unrelated to context
	// cancellation, and the session closing while the call was pending.
	// ErrorCode is set only for a remote CALLERROR. Note: this also fires
	// for every pending call during an ordinary graceful Server.Shutdown
	// (the session closes with ErrServerShutdown), which is expected and
	// not a regression.
	MetricOutboundCallFailed
	// MetricOutboundCallTimeout fires when the call's context deadline
	// expires, including when that deadline has already passed at the
	// moment the CALL frame is written. It cannot distinguish
	// Config.CallTimeout from a deadline the caller supplied on its own
	// context.
	MetricOutboundCallTimeout
	// MetricOutboundCallCanceled fires when the caller cancels its own
	// context while the call is pending (including if it was already
	// canceled at the moment the CALL frame is written). This is not a
	// failure. Caveat: context.Context only exposes context.Canceled, not
	// why a context was canceled, so if a caller ties Call's ctx to
	// Session.Context() (or another context whose cancellation is driven by
	// the session's own lifecycle rather than the caller's own decision),
	// the session closing for a reason unrelated to the call itself
	// (Server.Shutdown, an idle/pong timeout, being replaced) is also
	// reported here rather than as MetricOutboundCallFailed — this hook
	// cannot distinguish the two cases from ctx.Err() alone.
	MetricOutboundCallCanceled
)

// MetricEvent describes one measurable protocol event. Unlike LogRecord, it
// is not meant for human-readable diagnostics: it carries Duration for
// latency observation and has no free-form reason text, so implementations
// can use Type/Action/Version/ErrorCode directly as low-cardinality metric
// labels. Identity is deliberately not in that list — see its own comment.
type MetricEvent struct {
	Type MetricEventType
	// Identity is the charging station's identity string. It is NOT safe to
	// use directly as a label in a cardinality-sensitive metrics backend
	// (for example a Prometheus label) in a large deployment: every
	// distinct Identity creates a new time series, and a fleet of
	// thousands of charging stations produces thousands of series per
	// metric — a well-known Prometheus operational failure mode. This is a
	// deployment-scale judgment call the library cannot make correctly for
	// every user (fine for a 10-charger depot, wrong for a 50,000-charger
	// public network), so it is left entirely to the Metrics
	// implementation: use Identity for filtering/branching logic (e.g.
	// route a subset of chargers to more detailed stats) rather than
	// passing it straight through as a label, unless your deployment is
	// small enough that per-charger series are acceptable. See
	// examples/prometheus-hook for a wiring example that deliberately
	// excludes Identity from its labels.
	Identity    string
	Version     protocol.Version
	MessageType protocol.MessageTypeID
	Action      string
	Duration    time.Duration
	ErrorCode   ErrorCode
}

// Metrics receives MetricEvent notifications for connection, CALL, and SEND
// lifecycle events. Implementations are expected to aggregate events (for
// example into counters and histograms) rather than perform I/O per event.
//
// Record must be safe for concurrent use: it is dispatched from its own
// goroutine for every event, across every session, so it may run
// concurrently with itself and carries no ordering guarantee relative to
// when events logically occurred. A slow or blocking Record implementation
// never delays the protocol server — recordMetric never waits for Record to
// return — but if Record is persistently slower than events are produced,
// the bounded number of concurrent dispatches (maxConcurrentMetricDispatch)
// is exhausted and further events are silently dropped rather than queued
// or blocked.
//
// The context passed to Record is always context.Background(), never a
// request- or session-scoped context: every MetricEvent can be produced at a
// point where the natural originating context (an inbound handler's context,
// an outbound Call's context, or the session's own context) is already
// canceled — most obviously for the very events that report a timeout,
// cancellation, or shutdown — and propagating it would let a defensively
// written Record implementation (one that no-ops when ctx.Err() != nil)
// silently drop exactly those events. A future tracing integration that
// needs span correlation will need its own mechanism rather than relying on
// this ctx.
type Metrics interface {
	Record(context.Context, MetricEvent)
}

type MetricsFunc func(context.Context, MetricEvent)

func (f MetricsFunc) Record(ctx context.Context, event MetricEvent) { f(ctx, event) }

// maxConcurrentMetricDispatch bounds how many MetricEvents may be queued for
// dispatch at once across an entire Server, shared by every session on the
// Server rather than allocated per session. It exists to cap memory/backlog
// growth if an application's Metrics implementation hangs or is persistently
// slower than events are produced; healthy implementations never approach
// it. If enough sessions simultaneously accumulate hung (not panicking,
// merely never returning) Record calls to exhaust the workers draining this
// queue, event dispatch for unrelated, healthy sessions starts silently
// dropping too. A hung (as opposed to slow) Record implementation is already
// a bug in the application; this is a last-resort ceiling on its blast
// radius, not a per-session fairness guarantee.
const maxConcurrentMetricDispatch = 1024

// metricDispatchWorkers is how many long-lived goroutines drain a Server's
// metric queue. It is deliberately much smaller than
// maxConcurrentMetricDispatch: those goroutines are started once (in New)
// and reused for the Server's whole lifetime, so — unlike the queue capacity
// above, sized for a worst-case backlog — this only needs to be big enough
// to keep up with the steady-state rate of genuinely concurrent Record
// calls, which the Metrics doc comment already asks implementations to keep
// fast (aggregate into counters/histograms, not perform I/O per event).
const metricDispatchWorkers = 32

// startMetricDispatch launches the fixed pool of goroutines that drain
// queue and call metrics.Record for every event, replacing what used to be
// a new goroutine spawned per MetricEvent (unbounded goroutine *creation*
// churn under sustained message rate, even though concurrent goroutines
// were already bounded by the old metricSlots semaphore). These goroutines
// run until stop is closed (by Server.Shutdown) — queue itself is never
// closed, since a concurrent recordMetric could still be sending to it,
// and closing a channel a sender might still write to panics. wg lets
// Shutdown wait for every worker to actually exit, not just signal them,
// so it can honestly call metric dispatch finished rather than merely
// asked-to-stop. If metrics is nil, no workers are started and
// recordMetric's nil-queue check makes dispatch a no-op.
//
// Because stop and queue are both ready once Shutdown starts closing
// things down, a worker's select between them can still pick up and
// deliver an event that was already queued at that moment (Go's select
// among ready cases is pseudo-random) — this is intentional, not a bug to
// fix: it's the same "may still fire once, may not" guarantee recordMetric
// already documents for a nearly-full queue, just at shutdown instead of
// under load.
func startMetricDispatch(queue chan MetricEvent, stop <-chan struct{}, wg *sync.WaitGroup, metrics Metrics) {
	for range metricDispatchWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case event := <-queue:
					func() {
						// A diagnostic hook must not be able to take down
						// the protocol server.
						defer func() { _ = recover() }()
						metrics.Record(context.Background(), event)
					}()
				case <-stop:
					return
				}
			}
		}()
	}
}

// recordMetric is the only dispatch path to the Server's metric queue. It is
// used both by the inbound handlers in server.go and by the outbound Call in
// call.go, since Call only has access to a Session, never a Server. It never
// blocks the caller: this only enqueues the event for one of the pool
// goroutines started by startMetricDispatch to pick up, so a slow or hung
// Metrics implementation cannot delay handler completion (and the
// handlerSlots it occupies) or Call's prompt return on context cancellation.
// Record is always called with context.Background() (see the Metrics doc
// comment for why), so recordMetric takes no context.Context parameter
// itself — there is nothing meaningful for a caller to supply.
func (s *Session) recordMetric(event MetricEvent) {
	if s.metricQueue == nil {
		return
	}
	select {
	case s.metricQueue <- event:
	default:
		// The queue is full — every dispatch worker is busy long enough
		// that maxConcurrentMetricDispatch events are already pending;
		// drop this one rather than block or grow the queue unboundedly.
	}
}
