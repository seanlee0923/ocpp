# Sessions and connection lifecycle

[한국어](../sessions.md) | **English**

`Server` keeps at most one active `Session` per identity. By default an
identity is limited to RFC 3986 unreserved characters and at most 255 UTF-8
bytes.

```go
session, ok := server.Session("CP-001")
sessions := server.Sessions()
info := session.Info()
```

`SessionInfo` includes the version, principal, remote address, connection
time, last send/receive time, last pong time, and state. The returned
principal attributes and session slice are copies, so they never mutate the
internal registry.

A duplicate identity is rejected with HTTP 409 by default. Only when needed
does a policy callback approve replacing the existing connection.

```go
OnDuplicateSession: func(ctx context.Context, attempt csms.DuplicateSessionAttempt) (
    csms.DuplicateSessionDecision, error,
) {
    return csms.ReplaceExistingSession, nil
},
```

The existing session is not closed until the new WebSocket upgrade
succeeds. A replaced session's error is `ErrSessionReplaced`.

`PingInterval`, `PongTimeout`, and `IdleTimeout` manage connection health.
`Session.Done()` closes on termination and `Session.Err()` holds the cause.
`Session.Context()` can be used as the cancellation basis for handlers and
background work.

`Server.Shutdown(ctx)` rejects new connections and waits for active
sessions and handlers to finish. It is safe to call repeatedly, but a
shut-down Server cannot be reused.

`Server.Snapshot()` returns a `ServerSnapshot` with the active session count,
each session's `SessionInfo`, and whether the server is shutting down.
`Server.Healthy()` is a simple bool that turns `false` as soon as `Shutdown`
is called. Neither imposes an HTTP endpoint — wire them into whatever
health/readiness route your application needs.

Runnable code is available in
[`examples/duplicate-session`](../../examples/duplicate-session),
[`examples/graceful-shutdown`](../../examples/graceful-shutdown), and
[`examples/metrics-hook`](../../examples/metrics-hook) (a status endpoint
example).
