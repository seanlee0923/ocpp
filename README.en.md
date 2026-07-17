# ocpp-go

[한국어](README.md) | **English**

A Go library providing an OCPP-J CSMS WebSocket server. It handles OCPP 1.6,
2.0.1, and 2.1 on a single transport layer, with message types generated from
the official JSON Schemas and type-safe routing.

The module path is `github.com/seanlee0923/ocpp`.

## Key features

- OCPP-J server built on Gorilla WebSocket
- `ocpp1.6`, `ocpp2.0.1`, `ocpp2.1` subprotocol negotiation
- CALL, CALLRESULT, CALLERROR, and OCPP 2.1 SEND frame encode/decode
- Fully generated message types for OCPP 1.6, 2.0.1 Edition 4, and 2.1 Edition 2
- Go request/confirmation validation generated from the official OCA JSON Schemas
- Typed inbound handlers and typed outbound calls per Action
- Charging Station identity based sessions and duplicate-connection policy
- Security Profile 0, 1, 2, 3 handshake policies
- Ping/Pong, idle timeout, graceful shutdown, handler backpressure
- Default unique ID generator with support for injecting a custom one

## Requirements

- Go 1.26.5 or later
- A Charging Station that speaks OCPP-J WebSocket

Main dependency:

- `github.com/gorilla/websocket`

## Support status

| Area | OCPP 1.6 | OCPP 2.0.1 | OCPP 2.1 |
|---|---:|---:|---:|
| Generated message types | Done | Done | Done |
| Generated Go validation code | Done | Done | Done |
| Shared WebSocket transport | Done | Done | Done |
| Core Profile API | Done | Done | Done |

The current priority is the CSMS server. A Charging Station client will be
added once the server API has stabilized.

## Quick start

Create a Router and an OCPP 1.6 Profile, then pass the same Router to the
CSMS Server.

```go
package main

import (
    "context"
    "log"
    "net/http"
    "time"

    "github.com/seanlee0923/ocpp/csms"
    "github.com/seanlee0923/ocpp/profiles/ocpp16"
    "github.com/seanlee0923/ocpp/protocol"
    "github.com/seanlee0923/ocpp/v16"
)

func main() {
    router := csms.NewRouter()

    profile, err := ocpp16.NewProfile(router)
    if err != nil {
        log.Fatal(err)
    }

    err = profile.HandleBootNotification(func(
        ctx context.Context,
        session *csms.Session,
        request v16.BootNotificationRequest,
    ) (v16.BootNotificationConfirmation, error) {
        return v16.BootNotificationConfirmation{
            CurrentTime: time.Now().UTC().Format(time.RFC3339),
            Interval:    300,
            Status:      v16.BootNotificationConfirmationStatusAccepted,
        }, nil
    })
    if err != nil {
        log.Fatal(err)
    }

    server, err := csms.New(csms.Config{
        Router:   router,
        Versions: []protocol.Version{protocol.OCPP16},
    })
    if err != nil {
        log.Fatal(err)
    }

    log.Fatal(http.ListenAndServe(":8080", server))
}
```

A Charging Station uses its identity as the WebSocket path.

```text
ws://localhost:8080/CP-001
Sec-WebSocket-Protocol: ocpp1.6
```

The same code is available as a runnable example in
[`examples/ocpp16-csms`](examples/ocpp16-csms).

## Message types

Types per version live in a single package each, with no `req`/`conf`
subpackages.

```go
import (
    "github.com/seanlee0923/ocpp/v16"
    "github.com/seanlee0923/ocpp/v201"
    "github.com/seanlee0923/ocpp/v21"
)
```

Every generated message implements `protocol.Payload`.

```go
var payload protocol.Payload = v201.BootNotificationRequest{}

payload.ActionName()
payload.Version()
payload.Direction()
payload.SchemaName()
err := payload.Validate()
```

## Documentation

- [Getting started](docs/en/getting-started.md)
- [Architecture and responsibilities](docs/en/architecture.md)
- [Inbound handlers](docs/en/handlers.md)
- [Outbound CALL](docs/en/outbound-calls.md)
- [Sessions and connection lifecycle](docs/en/sessions.md)
- [Authentication and TLS](docs/en/security.md)
- [Validation and error handling](docs/en/validation-and-errors.md)
- [OCPP 1.6](docs/en/ocpp16.md)
- [OCPP 2.0.1](docs/en/ocpp201.md)
- [OCPP 2.1](docs/en/ocpp21.md)
- [Production configuration](docs/en/production.md)

Small, compilable examples live in [`examples`](examples). DB, cache, queue,
and business rules are wired up in application handlers; this library does
not impose any particular infrastructure.

Network input goes through generated Go validation code before it is decoded
into a Go type, so unknown properties, required fields, enums, length,
range, and format constraints are all checked. The official JSON Schema
sources and the code generation tooling are not included in the library
distribution.

## OCA Schema provenance

Message types and validation code were generated on 2026-07-16 from the
Part 3 JSON Schemas downloaded from the official OCA download page. After
generation, the original JSON files and the generator tooling were removed
from the repository.

- Official download: https://openchargealliance.org/my-oca/ocpp/
- OCPP 1.6 all files and errata ZIP SHA-256:
  `2a1d80284ca60449e85951fc55bc0538b5d45bd6f97c6d9228745acacb52c11b`
- OCPP 2.0.1 Edition 4 all files ZIP SHA-256:
  `192482c82a5e27a2319d2142be2d8c074b68a22851ff5a12d0541efc1eda775a`
- OCPP 2.1 Edition 2 all files ZIP SHA-256:
  `a306ddcca06936ba9de3c1c436d728fc74289367aa1c393f73b0deb2180c217e`

At generation time there were 56 OCPP 1.6 files, 128 OCPP 2.0.1 files, and
181 OCPP 2.1 files, each verified byte-for-byte identical to the official
ZIP contents. Copyright and licensing of OCPP and the original schemas
belong to the Open Charge Alliance.

## Typed inbound handlers

Actions outside the official OCPP Profiles can also be registered through
the shared typed API.

```go
err := csms.Handle(router, func(
    ctx context.Context,
    session *csms.Session,
    request v16.AuthorizeRequest,
) (v16.AuthorizeConfirmation, error) {
    return v16.AuthorizeConfirmation{
        IDTagInfo: v16.AuthorizeConfirmationIDTagInfo{
            Status: v16.AuthorizeConfirmationIDTagInfoStatusAccepted,
        },
    }, nil
})
```

Registration is rejected if the request and confirmation disagree on
version, Action, or direction. Registering the same `version + action` twice
returns `csms.ErrHandlerAlreadyRegistered` and leaves the existing handler in
place. There is no API for replacing a handler at runtime.

Invalid requests are converted into a constraint CALLERROR based on the
cause and the OCPP-J version.

- Unknown property, enum, length/range/format: `PropertyConstraintViolation`
- Missing required field, array item count: OCPP 1.6
  `OccurenceConstraintViolation`, OCPP 2.0.1/2.1
  `OccurrenceConstraintViolation`
- JSON field type mismatch: `TypeConstraintViolation`
- Payload syntax error: OCPP 1.6 `FormationViolation`, OCPP 2.0.1/2.1
  `FormatViolation`

An invalid confirmation is converted into `InternalError`. A MessageType
unsupported on OCPP 2.0.1/2.1 is answered with a `MessageTypeNotSupported`
CALLERROR.

## OCPP 1.6 Core Profile

Actions initiated by the Charging Station are registered as independent
handlers. Business logic such as DB lookups, authorization, billing, and
transaction storage is implemented by the library user.

```go
profile.HandleHeartbeat(heartbeatHandler)
profile.HandleStatusNotification(statusNotificationHandler)
profile.HandleAuthorize(authorizeHandler)
profile.HandleStartTransaction(startTransactionHandler)
profile.HandleMeterValues(meterValuesHandler)
profile.HandleStopTransaction(stopTransactionHandler)
profile.HandleDataTransfer(dataTransferHandler)
```

Before BootNotification results in `Accepted`, any other OCPP 1.6 inbound
Action is rejected with `ProtocolError`.

```go
state := profile.State(session)
booted := profile.IsBooted(session)
```

Actions initiated by the CSMS are invoked as typed calls.

```go
confirmation, err := profile.CallReset(ctx, session, v16.ResetRequest{
    Type: v16.ResetRequestTypeSoft,
})
```

Provided outbound API:

- `CallDataTransfer`
- `CallReset`
- `CallRemoteStartTransaction`
- `CallRemoteStopTransaction`
- `CallChangeAvailability`
- `CallUnlockConnector`
- `CallTriggerMessage`

DataTransfer is bidirectional, so both `HandleDataTransfer` and
`CallDataTransfer` are provided. Profile outbound calls return
`ErrNotBooted` for a session that has not completed Boot.

## OCPP 2.0.1 Core Profile

`profiles/ocpp201` provides the OCPP 2.0.1 Boot registration state and the
Core typed API. Business data and TransactionEvent ordering policy are
managed by the application handler.

```go
profile, err := ocpp201.NewProfile(router)
if err != nil {
    return err
}

err = profile.HandleBootNotification(func(
    ctx context.Context,
    session *csms.Session,
    request v201.BootNotificationRequest,
) (v201.BootNotificationConfirmation, error) {
    return v201.BootNotificationConfirmation{
        CurrentTime: time.Now().UTC().Format(time.RFC3339),
        Interval:    300,
        Status: v201.BootNotificationConfirmationRegistrationStatusEnumAccepted,
    }, nil
})
```

Handlers for `Heartbeat`, `StatusNotification`, `Authorize`,
`TransactionEvent`, `MeterValues`, and `NotifyReport`, initiated by the
Charging Station, can be registered independently. On the CSMS side,
`GetVariables`, `SetVariables`, `GetBaseReport`, `RequestStartTransaction`,
`RequestStopTransaction`, and `Reset` are invoked as typed calls.

Before Boot reaches `Accepted`, other inbound Actions are rejected with
`ProtocolError` and outbound calls return `ocpp201.ErrNotBooted`.

OCPP 2.1's `NotifyPeriodicEventStream` is a SEND message with no response.
It is registered through an API separate from regular typed handlers, and
the CSMS never sends a CALLRESULT or CALLERROR for it.

```go
err := profile21.HandleNotifyPeriodicEventStream(func(
    ctx context.Context,
    session *csms.Session,
    request v21.NotifyPeriodicEventStreamRequest,
) error {
    // Store or forward the stream values.
    return nil
})
```

Beyond Core, the OCPP 2.1 Profile provides the following server API:

- Smart charging and energy transfer
- Priority charging, dynamic schedules, and aFRR signals
- DER control and status reporting
- Tariffs, cost, settlement, web payment, and VAT validation
- Battery swap
- Periodic event streams
- Certificates, ISO 15118 certificates, and firmware lifecycle

Periodic event stream directions:

- Charging Station → CSMS CALL: `OpenPeriodicEventStream`,
  `ClosePeriodicEventStream`
- CSMS → Charging Station CALL: `GetPeriodicEventStream`,
  `AdjustPeriodicEventStream`
- Charging Station → CSMS SEND: `NotifyPeriodicEventStream`

The Profile only manages the typed transport boundary and Boot state.
Schedule computation, DER state, tariff and settlement storage, CA/OCSP/CRL,
private keys, and firmware file and signature verification are implemented
by the library user.

## Shared outbound call

Actions without a Profile wrapper can use `csms.Call` directly.

```go
confirmation, err := csms.Call[v16.ResetRequest, v16.ResetConfirmation](
    ctx,
    session,
    v16.ResetRequest{Type: v16.ResetRequestTypeSoft},
)
```

A remote CALLERROR is returned as `*csms.RemoteCallError`. Call timeout and
the maximum number of pending calls are controlled by the Server
configuration.

```go
server, err := csms.New(csms.Config{
    CallTimeout:     30 * time.Second,
    MaxPendingCalls: 100,
})
```

A custom unique ID generator must be concurrency-safe; empty IDs and IDs
that duplicate a currently pending call are rejected.

## Sessions and duplicate connections

If the same Charging Station identity is already connected, a new
connection is rejected with HTTP `409 Conflict` by default.

```go
server, err := csms.New(csms.Config{
    OnDuplicateSession: func(
        ctx context.Context,
        attempt csms.DuplicateSessionAttempt,
    ) (csms.DuplicateSessionDecision, error) {
        return csms.ReplaceExistingSession, nil
    },
})
```

Even under the replacement policy, the existing session is closed only
after the new WebSocket upgrade succeeds.

```go
session, ok := server.Session("CP-001")
sessions := server.Sessions()
count := server.SessionCount()
```

## Connection lifecycle

```go
server, err := csms.New(csms.Config{
    WriteTimeout:          10 * time.Second,
    PingInterval:          30 * time.Second,
    PongTimeout:           90 * time.Second,
    IdleTimeout:           0,
    MaxConcurrentHandlers: 16,
})
```

`Session.Info()` exposes connection time, last send/receive, last pong, and
state.

```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

err := server.Shutdown(ctx)
```

During shutdown, new connections are rejected with HTTP
`503 Service Unavailable`, and active sessions and handler contexts are
terminated.

## Security profiles

Supported OCPP-J handshake policies:

- Profile 0: unauthenticated WebSocket
- Profile 1: Basic Authentication
- Profile 2: TLS 1.2+ with Basic Authentication
- Profile 3: TLS 1.2+ with a verified client certificate

```go
limiter, err := csms.NewIPRateLimiter(20, time.Minute)

server, err := csms.New(csms.Config{
    Security: csms.SecurityConfig{
        Profile:          csms.SecurityProfileTLSBasicAuth,
        HandshakeLimiter: limiter,
        Authenticator: csms.AuthenticatorFunc(func(
            ctx context.Context,
            request csms.AuthenticationRequest,
        ) (csms.Principal, error) {
            if !verify(request.Identity, request.Basic.Password) {
                return csms.Principal{}, csms.ErrAuthenticationFailed
            }
            return csms.Principal{ID: request.Identity}, nil
        }),
    },
})
```

TLS handshake and certificate chain verification are handled by the outer
`http.Server`'s `tls.Config`. The library checks the verified TLS state and
passes credentials to the Authenticator; the Session stores only a
`Principal`, never a password.

## API compatibility

This project follows Semantic Versioning. During `v0.x`, the public API may
change while working toward a first stable release, and changes are
recorded in release notes. From `v1.0.0` onward, source compatibility of
the existing public API is preserved within the same major version.

`csms.Config` should be constructed with keyed field literals. A zero value
for a duration or limit selects the documented default, except
`IdleTimeout: 0`, which disables idle detection. Negative durations and
limits are rejected with `csms.ErrInvalidConfiguration`.
`Server.Shutdown` is safe to call multiple times, but the server cannot be
restarted after shutdown.

Applications never write raw WebSocket frames directly; they use
`csms.Call` or a version-specific profile's typed `Call...` methods. This
path manages pending CALL registration, unique ID validation, and response
correlation together.

Errors that can be branched on:

- `csms.ErrInvalidConfiguration`
- `csms.ErrInvalidHandlerRegistration`
- `csms.ErrHandlerAlreadyRegistered`
- `csms.ErrUniqueIDGeneration`
- `csms.ErrDuplicateUniqueID`
- `csms.ErrTooManyPendingCalls`
- `csms.ErrSessionClosed`, `csms.ErrSessionReplaced`, `csms.ErrServerShutdown`

Use `errors.Is` and `errors.As` instead of string comparison. A CALLERROR
sent by a Charging Station is returned as `*csms.RemoteCallError`, whose
`Code` is a `csms.ErrorCode`.

## Verification

```sh
GOCACHE=/tmp/ocpp-go-build-cache go vet ./...
GOCACHE=/tmp/ocpp-go-build-cache go test ./...
GOCACHE=/tmp/ocpp-go-build-cache go test -race ./...
```

WebSocket integration tests use a local loopback port.

The regular suite includes a bounded load test that processes 800 CALLs
across 16 sessions, and a concurrent outbound pending-call limit test. It
also includes regression tests confirming that sessions, pending calls, the
read loop, profile state, and goroutines are released after repeated
abrupt disconnects. Long-running soak tests are opt-in and log heap/goroutine
statistics before and after the run.

```sh
OCPP_SOAK_DURATION=30m GOCACHE=/tmp/ocpp-go-build-cache go test -race \
  -timeout 35m -run '^TestSessionSoak$' ./csms
```

`go test`'s default `-timeout` is 10 minutes, so you must pass a `-timeout`
longer than `OCPP_SOAK_DURATION`.

Benchmarks cover frame encode/decode, schema validation, Router dispatch, and
a full WebSocket loopback round trip.

```sh
GOCACHE=/tmp/ocpp-go-build-cache go test -run '^$' -bench . -benchmem ./...
```

Reference numbers below (Apple M2, local run; actual results vary by hardware).

| Benchmark | Time/op | Bytes/op | Allocs/op |
|---|---:|---:|---:|
| `DecodeCall` | 9.6 µs | 4.8 KB | 115 |
| `EncodeCallResult` | 2.6 µs | 969 B | 21 |
| `ValidateJSONSmall` (BootNotification) | 2.2 µs | 1.8 KB | 27 |
| `ValidateJSONLarge` (100-item array) | 93.4 µs | 55.6 KB | 1026 |
| `ValidateThenUnmarshalSmall` | 3.8 µs | 2.2 KB | 37 |
| `ValidateThenUnmarshalLarge` | 156.9 µs | 65.9 KB | 1143 |
| `RouterLookup`, 0 middleware | 15.5 ns | 0 B | 0 |
| `RouterLookup`, 5 middleware | 113.6 ns | 128 B | 6 |
| `InboundCallRoundTrip` (full WebSocket loopback round trip) | 30.4 µs | 5.7 KB | 82 |

Even a full round trip over a real network stays in the tens of
microseconds, so throughput is unlikely to be a bottleneck for OCPP, where a
single Charging Station sends a message every few minutes at most. For
deployments where many sessions send large-array payloads (such as
`NotifyPeriodicEventStream`) concurrently, see the
`ValidateJSONLarge`/`ValidateThenUnmarshalLarge` numbers.

## Structured logging

A `csms.Logger` can be injected without depending on any specific logging
library.

```go
server, err := csms.New(csms.Config{
    Router: router,
    Logger: csms.LoggerFunc(func(ctx context.Context, record csms.LogRecord) {
        // Forward to an application logger such as slog, zap, or zerolog.
    }),
})
```

The default log record exposes only identity, OCPP version, message type,
message ID, Action, error code, and a normalized reason. It never carries
payloads, Authorization headers, certificate contents, idTokens, or handler
error strings. A panicking logger does not stop the protocol server.

## Roadmap

- Finish documentation and per-version/operational examples
- Prepare for release
- Opt-in metrics, snapshots, and tracing as a final step
- Paid OCA OCTT and official certification are optional, post-release steps
- Charging Station client

## License

[MIT](LICENSE)
