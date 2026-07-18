# Changelog

[한국어](CHANGELOG.md) | **English**

This project follows [Keep a Changelog](https://keepachangelog.com/en/1.1.0/)
and [Semantic Versioning](https://semver.org/). During `v0.x`, API changes may
land via release notes; from `v1` on, source compatibility is preserved within
the same major version.

## [Unreleased]

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
