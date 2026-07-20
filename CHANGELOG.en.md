# Changelog

[한국어](CHANGELOG.md) | **English**

This project follows [Keep a Changelog](https://keepachangelog.com/en/1.1.0/)
and [Semantic Versioning](https://semver.org/). During `v0.x`, API changes may
land via release notes; from `v1` on, source compatibility is preserved within
the same major version.

## [Unreleased]

### Added

- Opt-in `csms.Metrics`/`csms.MetricsFunc` hook — observes session connect/
  disconnect, inbound CALL/SEND, and each stage of outbound `csms.Call`
  (sent, completed, failed, timed out, canceled, rejected because
  `MaxPendingCalls` was reached), carrying identity, OCPP version, message
  type, Action, elapsed time, and error code. A panicking implementation
  never affects the protocol server, matching `Logger`. `Record` is
  dispatched by a fixed pool of worker goroutines per `Server`, shared
  through a queue, so it never blocks the caller — a slow or hung
  implementation never delays handler processing or an outbound
  `csms.Call`'s cancellation responsiveness (in exchange, `Record` must be
  concurrency-safe, is not strictly ordered, and events are dropped once the
  queue is full). `Server.Shutdown` stops these workers too. `Record` always receives
  `context.Background()` — propagating the originating context would let a
  Metrics implementation that defensively no-ops on an already-canceled ctx
  silently drop exactly the events that report a timeout, cancellation, or
  shutdown. Configuring `Config.Metrics` adds roughly +6% to an inbound CALL
  round trip and +13% to an outbound `csms.Call` round trip (no overhead if
  left unset; see README for measured benchmarks).
- `Server.Snapshot()`/`ServerSnapshot` and `Server.Healthy()` — a server
  status API exposing active session count, per-session status, and
  shutdown state. No HTTP endpoint shape is imposed.
- Added the [`examples/metrics-hook`](examples/metrics-hook) example.
- Documented the cardinality risk of `MetricEvent.Identity` in its doc
  comment and added the [`examples/prometheus-hook`](examples/prometheus-hook)
  example. No official Prometheus adapter package is shipped — whether
  identity-scoped labels are safe depends on deployment size, a call the
  library cannot make correctly for every user. The example deliberately
  excludes `identity` from its labels.
- Added the [`examples/otel-hook`](examples/otel-hook) example. OpenTelemetry
  tracing also needs no dedicated hook: verified that `csms.Router`
  middleware (inbound) and the caller's own `ctx` (outbound `csms.Call`) are
  enough, including confirming end-to-end that an inbound handler span and
  an outbound `Call` span started inside it form a correct parent-child
  relationship. Unlike `MetricEvent.Identity`, a span attribute does not
  create a persistent time series per value, so attaching identity is safe.
- Added GoDoc, CI, and Codecov badges to the README. CI now measures coverage
  with `go test -coverpkg=./... -coverprofile=coverage.out ./...` and uploads
  it to Codecov (`fail_ci_if_error: false`, so CI itself doesn't fail before
  the Codecov account is connected). `codecov.yml` excludes the 365 generated
  types (`v16`/`v201`/`v21`) and `examples/` from the reported coverage, so
  the badge reflects real logic coverage rather than being dragged down by
  generated code and non-tested example programs.
- Added a `lint` job to CI (`golangci-lint`, default linter set). Go Report
  Card was sunset as a service on 2026-07-01 (an external event, not
  repo-specific), so its README badge was removed and replaced with
  golangci-lint. `defer conn.Close()`-style patterns in test files are
  excluded from `errcheck` via `.golangci.yml`.
- `station` package — a real OCPP-J Charging Station client with
  reconnect/backoff. Offers typed outbound `station.Call` and typed inbound
  `station.Handle` over the same generated v16/v201/v21 types, the same way
  `csms.Call`/`csms.Handle` do. Handlers registered via `station.Handle`
  survive reconnects; a pending Call on a connection that drops fails
  immediately (no offline queueing). Basic Auth/TLS (including mTLS) are
  configured via `Config.BasicAuth`/`Config.TLSConfig`. Running many
  chargers is the caller's responsibility (keep your own
  `map[string]*station.Station`) — this package only owns protocol/session
  mechanics. Added the [`examples/station-client`](examples/station-client)
  example.
- Linked [`ocpp-schema-gen`](https://github.com/seanlee0923/ocpp-schema-gen),
  a from-scratch generator, from the README's "OCA Schema source" section,
  so anyone can independently re-verify that `ocpp-go`'s 365 generated types
  actually match the official OCA JSON Schema.
- Documented, across `docs/handlers.md`, `docs/sessions.md`,
  `docs/security.md`, and the README, the contract that "`ctx` cancellation
  is a signal, not a forced termination, and handlers/hooks must return
  promptly." Spells out exactly how far the impact reaches when
  `OnConnect`/`OnDisconnect`/`OnDuplicateSession`,
  `Authenticator`/`HandshakeLimiter`, or `Metrics.Record` each hang.

### Fixed

- Moved `LogSessionConnected`/`MetricSessionConnected` recording earlier,
  before the read loop starts and before `Config.OnConnect` runs. Previously,
  a Charging Station writing a frame immediately after the upgrade, or an
  `OnConnect` callback using a blocking `csms.Call` (both supported
  patterns), could produce another event before the connected event.
- Outbound `csms.Call` now classifies a send failure caused by an
  already-expired context as `MetricOutboundCallTimeout`/
  `MetricOutboundCallCanceled` instead of `MetricOutboundCallFailed`.
- Remote CALLERROR error-code extraction now uses `errors.As` instead of a
  raw type assertion, so it stays correct if the error is ever wrapped.
- Removed an unused function (`hasSupportedSubprotocol`) surfaced while
  adopting `golangci-lint`, and split an unintentionally identical `||`
  condition in `TestIPRateLimiter` into two separate assertions (the
  behavior itself was already correct).
- Guarded `csms.Call[Request, *Confirmation]` against a pointer-typed
  confirmation type argument, which used to panic; it is now a
  registration-time error (`isNilType` guard).
- Fixed `Router` not distinguishing between handlers registered for CALL vs.
  SEND — a Charging Station sending the wrong message type for an action
  could leave a CALL permanently unanswered, or silently discard a computed
  confirmation.
- A CALL handler's confirmation failing to encode used to close the
  connection silently with no response; it now sends an `InternalError`
  CALLERROR.
- Bounded the JSON pre-validator's recursion depth (64), preventing stack
  exhaustion from a deeply nested payload.
- Bounded inbound handler execution time to `Config.CallTimeout` — previously
  it was only canceled when the session closed or the server shut down.
- Capped the number of clients `IPRateLimiter` tracks (1024), evicting the
  oldest entry once the cap is reached — previously an unbounded number of
  distinct IPs attempting to connect grew memory without limit.
- Fixed the shutdown cause being recorded as a generic error instead of the
  real one when connection activation lost a race, and fixed the disconnect
  reason being read from `readLoop`'s raw error instead of the first cause
  actually recorded (`Session.Err()`).
- Fixed the read loop blocking whenever the handler concurrency slots were
  full; admission is now non-blocking (`pendingSlots`) and rejects the CALL
  immediately once full (`ErrHandlerQueueFull`).
- Switched `station.dispatch` from spawning an unbounded goroutine per
  inbound CALL to the same bounded pending queue csms uses.
- Applied a real socket write deadline to `station`'s outbound sends
  (`Config.WriteTimeout`, default 10s) — previously a `ctx` timeout alone
  could not unblock an already-blocked write.
- `station` now normalizes a handler-returned `CallError` before sending it
  (filling empty fields, enforcing length limits) and bounds handler
  execution time to `Config.CallTimeout`.
- Fixed `station.Server.Shutdown` not waiting for metric dispatch workers to
  exit before returning.
- `station` now rejects non-text WebSocket frames (e.g. binary) — OCPP-J
  only allows text frames.
- `station` now disconnects on an unparseable frame instead of silently
  ignoring it and continuing to read, matching csms's behavior.
- Added `Version.Valid()` and `ReconnectPolicy.Multiplier` NaN/Inf checks to
  `station.New`'s config validation.

### Changed

- Consolidated CALLERROR code derivation, previously scattered per OCPP
  version, into shared helpers (no behavior change, easier to maintain).
- Changed `Metrics.Record` dispatch from spawning a goroutine per event to a
  fixed worker pool per `Server` consuming a shared queue (the Added entry
  above reflects the updated description). Externally observable behavior
  is unchanged.

## [0.1.0] - 2026-07-18

Initial public release. A Go library providing an OCPP-J CSMS WebSocket
server built on [gorilla/websocket](https://github.com/gorilla/websocket),
supporting OCPP 1.6, 2.0.1 Edition 4, and 2.1 Edition 2 on a single transport
layer.

### Added

- Common OCPP-J framing (CALL/CALLRESULT/CALLERROR, OCPP 2.1 SEND), a
  WebSocket CSMS server, and session lifecycle (ping/pong, idle timeout,
  graceful shutdown, duplicate-session policy)
- 365 generated Go types from the official OCA JSON Schemas (OCPP 1.6
  request/confirmation 28/28, 2.0.1 64/64, 2.1 91/90) with pure-Go validation
  and no external validator dependency
- OCPP 1.6 Core Profile: BootNotification/Heartbeat/StatusNotification/
  Authorize/StartTransaction/MeterValues/StopTransaction/DataTransfer and
  outbound control (Reset, etc.)
- OCPP 2.0.1 Edition 4 Core Profile: BootNotification/TransactionEvent/
  MeterValues/NotifyReport and outbound control (GetVariables,
  RequestStartTransaction, etc.)
- OCPP 2.1 Edition 2 Core Profile and extension groups: smart charging /
  energy transfer, priority charging & dynamic schedule, DER control,
  tariff/cost/payment, battery swap, periodic event streams (SEND),
  certificate/PKI, firmware update
- Security Profile 0 (Unsecured) / 1 (Basic Auth) / 2 (TLS + Basic Auth) / 3
  (TLS + mTLS), with a pluggable `Authenticator`, `HandshakeLimiter`, and a
  concurrency-safe built-in `IPRateLimiter`
- A library-neutral structured logger hook that excludes payloads,
  credentials, certificates, and idTokens
- Hot-path benchmarks (frame encode/decode, schema validation, Router
  dispatch, full WebSocket loopback round trip) and a 30-minute race soak
  test harness
- GitHub Actions CI (`gofmt`, `go vet`, `go build`, `go test`,
  `go test -race`, `govulncheck`)
- 11 detailed docs in Korean plus English translations (`docs/en/`), and
  runnable examples
- MIT license

### Fixed

- `customData` vendor extension fields on OCPP 2.0.1/2.1 types being silently
  dropped on marshal/unmarshal (introduced `internal/customdata`)
- 4 direction errors in OCPP 2.1 wrappers (NotifyAllowedEnergyTransfer,
  GetTariffs, NotifyWebPaymentStarted, GetCertificateChainStatus)
- Full Handle/Call direction audit across OCPP 1.6/2.0.1/2.1 wrappers, and
  version-aware CALLERROR code classification (OCPP 1.6
  `FormationViolation`/`OccurenceConstraintViolation` vs 2.x
  `FormatViolation`/`OccurrenceConstraintViolation`)
- A WebSocket session-registration race in tests that only surfaced under
  `-race` on GitHub Actions Linux runners
- Go toolchain bumped to 1.26.5 to address a `crypto/tls` Encrypted Client
  Hello CVE (`GO-2026-5856`)
- Missing `-timeout` flag in the documented soak test command

### Changed

- Internal implementation of the 45 `Call<Action>` methods across
  `profiles/ocpp16|ocpp201|ocpp21` was consolidated into a shared generic
  helper (`callBooted`). Public method names, signatures, and behavior are
  unchanged.

[Unreleased]: https://github.com/seanlee0923/ocpp/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/seanlee0923/ocpp/releases/tag/v0.1.0
