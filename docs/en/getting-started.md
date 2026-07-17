# Getting started

[한국어](../getting-started.md) | **English**

This library provides an OCPP-J WebSocket CSMS server. OCPP 1.6 SOAP
transport is not supported. Applications implement DB, cache, queue, and
business rules inside handlers.

## Install

```sh
go get github.com/seanlee0923/ocpp
```

The project targets Go 1.26.5 or later.

## Minimal server

```go
router := csms.NewRouter()
profile, err := ocpp16.NewProfile(router)
if err != nil {
    return err
}
err = profile.HandleBootNotification(func(
    ctx context.Context,
    session *csms.Session,
    request v16.BootNotificationRequest,
) (v16.BootNotificationConfirmation, error) {
    return v16.BootNotificationConfirmation{
        CurrentTime: time.Now().UTC().Format(time.RFC3339),
        Interval: 300,
        Status: v16.BootNotificationConfirmationStatusAccepted,
    }, nil
})
if err != nil {
    return err
}
server, err := csms.New(csms.Config{
    Router: router,
    Versions: []protocol.Version{protocol.OCPP16},
})
if err != nil {
    return err
}
return http.ListenAndServe(":8080", server)
```

A Charging Station sends its identity as the URL path and negotiates the
OCPP subprotocol.

```text
GET /CP-001
Upgrade: websocket
Sec-WebSocket-Protocol: ocpp1.6
```

Supported subprotocols are `ocpp1.6`, `ocpp2.0.1`, and `ocpp2.1`. Omitting
`Versions` allows all three. Per-version profiles and shared typed handlers
can both be registered on the same Router.

The full runnable code lives in
[`examples/ocpp16-csms`](../../examples/ocpp16-csms). See
[`examples/multi-version`](../../examples/multi-version) for a setup that
serves all three versions from one process.
