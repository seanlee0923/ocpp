# Production configuration

[한국어](../production.md) | **English**

## HTTP and shutdown

Use `csms.Server` as the Handler of an `http.Server`. On a process signal,
call `csms.Server.Shutdown` first, then `http.Server.Shutdown`. Give the
shutdown context enough deadline for business handlers to finish cleanly.

## Limits and timeouts

- `ReadLimit`: caps abnormally large WebSocket messages
- `HandshakeTimeout`: caps the upgrade
- `CallTimeout`: caps an outbound CALL with no deadline
- `MaxPendingCalls`: per-session cap on pending outbound requests
- `MaxConcurrentHandlers`: per-session cap on concurrently running inbound handlers
- `WriteTimeout`, `PingInterval`, `PongTimeout`, `IdleTimeout`: connection lifecycle

Zero selects the documented default, except `IdleTimeout: 0`, which
disables idle detection. Negative durations and limits are rejected as
invalid configuration. Construct `Config` with keyed fields, propagate the
handler's context to DB and outbound calls, and avoid unbounded background
goroutines.

After handlers run synchronously in registration order after the CALLRESULT
write while continuing to occupy one `MaxConcurrentHandlers` slot for that
session. Each callback receives a fresh `CallTimeout`, so N callbacks making
blocking CALLs can accumulate up to roughly N times that bound. Use a short
individual deadline for each CALL and keep both callback count and sequential
CALL count small; send long work to a bounded worker queue. Holding the slot
too long can delay the same station's post-Boot StatusNotification/Heartbeat
and eventually fill `MaxQueuedHandlers`.

## Proxy and TLS

If TLS is terminated at a reverse proxy, the current core does not
interpret proxy headers as a verified client certificate. mTLS Security
Profile 3 assumes the CSMS process receives the verified Go TLS state
directly. Restrict origin, identity, and authentication policy to match
your deployment.

## Observability and sensitive data

A `csms.Logger` can be wired into an application logger. Default metadata
never includes payloads, passwords, certificates, idTokens, or handler
error text. A panic in the logger callback is isolated from the server.

`csms.Metrics` follows the same principle, observing session connect/
disconnect, inbound CALL/SEND, and each stage of outbound `csms.Call`. It
only carries identity, OCPP version, message type, Action, elapsed time, and
an error code — never a payload. Panics are isolated the same way as
`Logger`. Outbound `csms.Call` distinguishes the send itself
(`MetricOutboundCallSent`) from its final outcome — completed, failed, timed
out, canceled, or rejected because `MaxPendingCalls` was reached.
`Metrics.Record` is dispatched on its own goroutine per event, so it never
blocks the caller — a slow or hung implementation cannot delay handler
processing or an outbound `csms.Call`'s cancellation responsiveness. In
exchange, `Record` must be safe for concurrent use and carries no ordering
guarantee; if concurrent dispatches exceed a fixed internal bound, further
events are dropped rather than queued (a healthy implementation never
reaches that bound). `Record` always receives `context.Background()`: the
events that report a timeout, cancellation, or shutdown are produced exactly
when the original context may already be canceled, so propagating it would
let a Metrics implementation that defensively no-ops when `ctx.Err() != nil`
silently drop those events. Configuring `Config.Metrics` adds roughly +6% to
an inbound CALL round trip and +13% to an outbound `csms.Call` round trip
(no overhead if left unset; see the benchmark table in the root
[README](../../README.en.md) for measured numbers).

`Server.Snapshot()` and `Server.Healthy()` expose the active session count,
per-session status, and shutdown state. The library does not impose an HTTP
endpoint shape — define whatever path and payload your health/readiness
probes need.

Do not use `MetricEvent.Identity` (the charging station identity) directly
as a label. In a deployment with thousands or tens of thousands of charging
stations, every distinct identity creates a new time series, which is a
well-known operational failure mode for cardinality-sensitive backends like
Prometheus. Whether identity-scoped labels are safe depends on deployment
size — a judgment call the library cannot make correctly for every user (fine
for a small depot, dangerous for a large public network) — so no official
Prometheus adapter package is provided. Use `Identity` for filtering or
branching logic instead of passing it through as a label.

OpenTelemetry tracing has no dedicated hook either — `csms.Router` middleware
can hand handlers a new `ctx`, so inbound tracing is just middleware, and
outbound `csms.Call` needs nothing more than the caller starting its own span
and passing that `ctx` in. Unlike `MetricEvent.Identity`, a span attribute
does not create a persistent time series per value, so attaching identity as
an attribute is fine.

[`examples/structured-logger`](../../examples/structured-logger) shows
wiring up the standard `log/slog`,
[`examples/metrics-hook`](../../examples/metrics-hook) shows an in-process
metrics counter with a status endpoint,
[`examples/prometheus-hook`](../../examples/prometheus-hook) shows wiring
Prometheus counters/histograms directly, with no official adapter, using
cardinality-safe labels,
[`examples/otel-hook`](../../examples/otel-hook) shows wiring OpenTelemetry
spans with no dedicated tracing hook, using middleware and the caller's own
`ctx`, and
[`examples/graceful-shutdown`](../../examples/graceful-shutdown) shows
signal-based shutdown.

Before release, run `go test -race ./...` and an opt-in soak, and check the
goroutine/heap trend.
