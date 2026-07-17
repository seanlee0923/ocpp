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
Metrics, snapshots, and tracing are planned as future opt-in hooks.

[`examples/structured-logger`](../../examples/structured-logger) shows
wiring up the standard `log/slog`, and
[`examples/graceful-shutdown`](../../examples/graceful-shutdown) shows
signal-based shutdown.

Before release, run `go test -race ./...` and an opt-in soak, and check the
goroutine/heap trend.
