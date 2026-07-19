# Outbound CALL

[한국어](../outbound-calls.md) | **English**

Requests initiated by the CSMS toward a Charging Station use a profile
method or `csms.Call`.

```go
confirmation, err := csms.Call[v201.ResetRequest, v201.ResetConfirmation](
    ctx,
    session,
    v201.ResetRequest{Type: v201.ResetRequestResetEnumImmediate},
)
```

`Call` performs request validation, unique ID generation, pending
registration, serialization, response correlation, and confirmation
validation in a single step. It does not send if the request's OCPP version
does not match the session's.

If the call context has no deadline, `Config.CallTimeout` applies. Concurrent
pending requests are capped by `Config.MaxPendingCalls`; exceeding it returns
`ErrTooManyPendingCalls`. When a session closes, every pending CALL is
released immediately with that closure's cause.

A Charging Station's CALLERROR is a `*csms.RemoteCallError`.

```go
var remote *csms.RemoteCallError
if errors.As(err, &remote) {
    log.Printf("code=%s description=%s", remote.Code, remote.Description)
}
```

The default unique ID is `uuid.NewString`. A custom generator can be
injected via `Config.UniqueIDGenerator`. A panic, an empty value, or a value
longer than 36 characters returns `ErrUniqueIDGeneration`; an in-flight
duplicate within the same session returns `ErrDuplicateUniqueID`.

If `Config.Metrics` is set, every stage of `Call` (sent, completed, a remote
error or undecodable response, timed out, canceled, or rejected because
`MaxPendingCalls` was reached) is observed as a `MetricOutboundCallSent`/
`Completed`/`Failed`/`Timeout`/`Canceled`/`Rejected` event. See "Observability
and sensitive data" in [Production configuration](production.md) for details.

Since the caller already supplies `ctx` to `Call`, OpenTelemetry tracing
needs no library support either: start your own span before calling `Call`
and pass its `ctx` in. [`examples/otel-hook`](../../examples/otel-hook) shows
this pattern.

A flow that calls Reset on a real session after BootNotification is in
[`examples/outbound-call`](../../examples/outbound-call).
