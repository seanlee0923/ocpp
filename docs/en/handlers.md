# Inbound handlers

[한국어](../handlers.md) | **English**

CALLs initiated by a Charging Station are registered either through a
generic typed handler or a profile convenience method.

```go
err := csms.Handle(router, func(
    ctx context.Context,
    session *csms.Session,
    request v201.AuthorizeRequest,
) (v201.AuthorizeConfirmation, error) {
    return authorize(ctx, session.Identity(), request)
})
```

The version, Action, and request/confirmation direction are inferred from
the generic type parameters. Mismatched types, a nil handler, and
registering the same `version + action` twice are all registration-time
errors. Replacing a handler dynamically is not supported.

Inbound payloads are validated against the official Schema constraints
before the handler runs. The confirmation returned by the handler is
validated again before it is sent. Handlers must respect `ctx` cancellation,
which fires when the session closes, the server shuts down, or
`Config.CallTimeout` elapses. A panic is isolated into an `InternalError`
CALLERROR and does not crash the server, but it is not a substitute for
proper error handling.

`ctx` cancellation is a signal, not a forced termination — Go has no
mechanism to kill a goroutine from the outside, so if a handler ignores
cancellation and keeps running, the library cannot stop it. In that case the
handler keeps executing an already-timed-out request and continues to hold
one of the concurrent handler execution slots for as long as it runs; once
those slots are exhausted, subsequent CALLs are rejected as the
`pendingSlots`/handler queue being full. Handlers registered on the
Charging Station side via `station.Handle` follow the same contract, with
`station.Config.CallTimeout` playing the equivalent role.

Expected OCPP errors are returned as `*csms.CallError`.

```go
return confirmation, &csms.CallError{
    Code: csms.NotSupported,
    Description: "operation is disabled",
    Details: map[string]any{"policy": "maintenance"},
}
```

OCPP 2.1 SEND is registered with `csms.HandleSend`. Since SEND has no
protocol response, validation or handler errors can only be observed through
`csms.Logger` and `csms.Metrics` (`MetricSendDropped`).

Router middleware wraps handlers in registration order; when multiple
profiles share a Router, middleware must check the version it receives.
Middleware can construct a new `ctx` and pass it to the next handler, so
tracing (e.g. OpenTelemetry) can be wired up as middleware with no dedicated
hook needed — see [`examples/otel-hook`](../../examples/otel-hook).
