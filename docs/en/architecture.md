# Architecture and responsibilities

[한국어](../architecture.md) | **English**

```text
Charging Station
      │ OCPP-J WebSocket
      ▼
  csms.Server ── session, timeout, CALL correlation, shutdown
      │
      ▼
  csms.Router ── version + Action dispatch, middleware
      │
      ▼
 Typed handler ── generated request / confirmation
      │
      ▼
Application ── DB, cache, queue, authorization, billing, PKI
```

The library owns WebSocket transport, OCPP-J framing, schema validation,
typed routing, session lifecycle, and concurrency control. The application
owns storing charge points, EVSEs, connectors, and transactions; idToken
authorization; reservations; smart charging policy; billing; certificate
issuance; and firmware distribution.

`v16`, `v201`, and `v21` are single packages generated from the official
Schemas. Requests and confirmations implement `protocol.Payload`.
`profiles/ocpp16`, `profiles/ocpp201`, and `profiles/ocpp21` are optional
convenience layers that provide commonly used Actions and the
BootNotification registration gate. Actions without a wrapper can still be
used through `csms.Handle` and `csms.Call`.

`Session` and the underlying WebSocket are owned by the server.
Applications should use the typed API instead of writing raw frames so that
pending CALLs, unique IDs, and response correlation stay consistent.
