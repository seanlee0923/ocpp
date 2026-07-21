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

## Calling by identity outside an OCPP handler

When an event arrives outside the OCPP handler flow—from an HTTP request, a
message queue, or a scheduler—do not retain a `*csms.Session` long-term in a
separate map or database. A reconnect registers a new `Session` for the same
identity, so look up the current connection with `Server.Session(identity)`
immediately before the CALL.

```go
type application struct {
    server  *csms.Server
    profile *ocpp16.Profile
}

func (app *application) dataTransfer(w http.ResponseWriter, r *http.Request) {
    cpid := r.PathValue("cpid")
    session, ok := app.server.Session(cpid)
    if !ok {
        http.Error(w, "charging station is not connected", http.StatusNotFound)
        return
    }

    ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
    defer cancel()
    confirmation, err := app.profile.CallDataTransfer(
        ctx,
        session,
        v16.DataTransferRequest{VendorID: "com.example"},
    )
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadGateway)
        return
    }
    json.NewEncoder(w).Encode(confirmation)
}

mux.HandleFunc("POST /stations/{cpid}/data-transfer", app.dataTransfer)
```

The connection may still close or be replaced by a new connection for the
same identity between a successful lookup and CALL completion. Treat
`ErrSessionClosed`, `ErrSessionReplaced`, context timeouts, and
`*RemoteCallError` as normal operational outcomes. Profile methods also
enforce their profile policy; for example, OCPP 1.6
`Profile.CallDataTransfer` returns `ocpp16.ErrNotBooted` before an accepted
BootNotification.

The complete runnable example starts the OCPP WebSocket server on `:8080`
and an external HTTP API on `:8081`, then forwards a JSON body as an OCPP 1.6
`DataTransfer`: [`examples/external-data-transfer`](../../examples/external-data-transfer).
