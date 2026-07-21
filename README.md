# ocpp-go

**한국어** | [English](README.en.md)

[![Go Reference](https://pkg.go.dev/badge/github.com/seanlee0923/ocpp.svg)](https://pkg.go.dev/github.com/seanlee0923/ocpp)
[![CI](https://github.com/seanlee0923/ocpp/actions/workflows/ci.yml/badge.svg)](https://github.com/seanlee0923/ocpp/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/seanlee0923/ocpp/branch/main/graph/badge.svg)](https://codecov.io/gh/seanlee0923/ocpp)

Go로 작성된 OCPP-J CSMS WebSocket 서버 라이브러리입니다. 하나의 전송 계층에서
OCPP 1.6, 2.0.1, 2.1을 처리하며 공식 JSON Schema 기반 메시지 타입과 타입 안전
라우팅을 제공합니다.

모듈 경로는 `github.com/seanlee0923/ocpp`입니다.

## 주요 기능

- Gorilla WebSocket 기반 OCPP-J 서버
- `ocpp1.6`, `ocpp2.0.1`, `ocpp2.1` 서브프로토콜 협상
- CALL, CALLRESULT, CALLERROR와 OCPP 2.1 SEND frame encode/decode
- OCPP 1.6, 2.0.1 Edition 4, 2.1 Edition 2 전체 생성 메시지 타입
- 공식 OCA JSON Schema에서 생성한 Go request/confirmation 검증
- Action별 typed inbound handler와 typed outbound call
- 충전기 identity 기반 세션과 중복 연결 정책
- Security Profile 0, 1, 2, 3 handshake 정책
- Ping/Pong, idle timeout, graceful shutdown, handler backpressure
- 기본 Unique ID generator와 사용자 generator 주입
- `station` 패키지: 재연결/backoff을 갖춘 typed Charging Station 클라이언트

## 요구 사항

- Go 1.26.5 이상
- OCPP-J WebSocket을 사용하는 Charging Station

주요 의존성:

- `github.com/gorilla/websocket`

## 지원 상태

| 영역 | OCPP 1.6 | OCPP 2.0.1 | OCPP 2.1 |
|---|---:|---:|---:|
| 생성 메시지 타입 | 완료 | 완료 | 완료 |
| 생성 Go 검증 코드 | 완료 | 완료 | 완료 |
| 공통 WebSocket 전송 | 완료 | 완료 | 완료 |
| Core Profile API | 완료 | 완료 | 완료 |

CSMS 서버가 우선순위였고 이제 API가 안정화됐습니다. `station` 패키지로 Charging
Station 클라이언트도 제공합니다 — 아래 "Charging Station 클라이언트" 참고.

## 빠른 시작

Router와 OCPP 1.6 Profile을 만든 뒤 같은 Router를 CSMS Server에 전달합니다.

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

Charging Station은 identity를 WebSocket 경로로 사용합니다.

```text
ws://localhost:8080/CP-001
Sec-WebSocket-Protocol: ocpp1.6
```

같은 코드는 [`examples/ocpp16-csms`](examples/ocpp16-csms)에서 실행 가능한 형태로
제공합니다.

OCPP handler 밖의 HTTP API나 queue event에서 identity로 현재 세션을 조회해 CALL을
보내는 예제는 [`examples/external-data-transfer`](examples/external-data-transfer)에
있습니다. 예제를 실행하면 충전기는 `ws://localhost:8080/{cpid}`에 연결하고, 외부에서는
다음처럼 OCPP 1.6 `DataTransfer`를 요청합니다.

```sh
curl -X POST http://localhost:8081/stations/CP-001/data-transfer \
  -H 'Content-Type: application/json' \
  -d '{"vendorId":"com.example","messageId":"Sync","data":"{\"enabled\":true}"}'
```

## 메시지 타입

버전별 타입은 `req`, `conf` 하위 패키지 없이 단일 패키지에 있습니다.

```go
import (
    "github.com/seanlee0923/ocpp/v16"
    "github.com/seanlee0923/ocpp/v201"
    "github.com/seanlee0923/ocpp/v21"
)
```

모든 생성 메시지는 `protocol.Payload`를 구현합니다.

```go
var payload protocol.Payload = v201.BootNotificationRequest{}

payload.ActionName()
payload.Version()
payload.Direction()
payload.SchemaName()
err := payload.Validate()
```

## 문서

- [시작하기](docs/getting-started.md)
- [구조와 책임 범위](docs/architecture.md)
- [Inbound handler](docs/handlers.md)
- [Outbound CALL](docs/outbound-calls.md)
- [세션과 연결 생명주기](docs/sessions.md)
- [인증과 TLS](docs/security.md)
- [검증과 오류 처리](docs/validation-and-errors.md)
- [OCPP 1.6](docs/ocpp16.md)
- [OCPP 2.0.1](docs/ocpp201.md)
- [OCPP 2.1](docs/ocpp21.md)
- [운영 환경 구성](docs/production.md)

컴파일 가능한 작은 예제는 [`examples`](examples)에 있습니다. DB, cache, queue와 업무
규칙은 애플리케이션 handler에서 연결하며 이 라이브러리는 특정 인프라를 강제하지 않습니다.

네트워크 입력은 Go 타입으로 decode하기 전에 생성된 Go 검증 코드를 거칩니다. 따라서
unknown property, required, enum, 길이, 범위와 format 제약도 검사합니다. 공식 JSON
Schema 원본이나 코드 생성 도구는 라이브러리 배포 파일에 포함하지 않습니다.

## OCA Schema 출처

메시지 타입과 검증 코드는 2026-07-16에 OCA 공식 다운로드 페이지에서 받은 Part 3
JSON Schema를 기준으로 생성했습니다. 생성 후 원본 JSON과 생성 도구는 저장소에서
제외했습니다.

- 공식 다운로드: https://openchargealliance.org/my-oca/ocpp/
- OCPP 1.6 all files and errata ZIP SHA-256:
  `2a1d80284ca60449e85951fc55bc0538b5d45bd6f97c6d9228745acacb52c11b`
- OCPP 2.0.1 Edition 4 all files ZIP SHA-256:
  `192482c82a5e27a2319d2142be2d8c074b68a22851ff5a12d0541efc1eda775a`
- OCPP 2.1 Edition 2 all files ZIP SHA-256:
  `a306ddcca06936ba9de3c1c436d728fc74289367aa1c393f73b0deb2180c217e`

생성 당시 파일 수는 OCPP 1.6 56개, OCPP 2.0.1 128개, OCPP 2.1 181개였으며,
각 파일은 공식 ZIP과 byte 단위로 일치함을 확인했습니다. OCPP와 원본 Schema의
저작권 및 라이선스는 Open Charge Alliance에 있습니다.

checksum만으로는 원본 스키마가 변조되지 않았다는 것만 확인할 수 있을 뿐, 그
스키마로부터 생성된 Go 코드가 실제로 스키마를 충실히 반영하는지는 증명하지
못합니다. 이 부분을 직접 확인하고 싶다면
[seanlee0923/ocpp-schema-gen](https://github.com/seanlee0923/ocpp-schema-gen)을
참고하세요 — 위 세 버전 전체 365개 생성 파일을 처음부터 재생성해서 이
저장소의 `v16`/`v201`/`v21` 파일과 byte 단위로 동일함을 독립적으로 재현할 수
있는 별도의 생성기입니다.

## Typed inbound handler

공식 OCPP Profile 외의 Action도 공통 typed API로 등록할 수 있습니다.

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

Request와 Confirmation의 버전, Action, 방향이 다르면 등록 단계에서 거부됩니다.
같은 `version + action`을 두 번 등록해도 `csms.ErrHandlerAlreadyRegistered` 오류가
반환되며 기존 handler는 유지됩니다. 실행 중 handler 교체 API는 제공하지 않습니다.

잘못된 request는 원인과 OCPP-J 버전에 따라 constraint CALLERROR로 변환됩니다.

- unknown property, enum, 길이·범위·format: `PropertyConstraintViolation`
- required 누락, 배열 항목 수: OCPP 1.6 `OccurenceConstraintViolation`,
  OCPP 2.0.1/2.1 `OccurrenceConstraintViolation`
- JSON field type 불일치: `TypeConstraintViolation`
- payload 문법 오류: OCPP 1.6 `FormationViolation`, OCPP 2.0.1/2.1
  `FormatViolation`

잘못된 confirmation은 `InternalError`로 변환됩니다. OCPP 2.0.1/2.1에서 지원하지 않는
MessageType은 `MessageTypeNotSupported` CALLERROR로 응답합니다.

## OCPP 1.6 Core Profile

충전기에서 시작하는 Action은 독립적인 handler로 등록합니다. DB 조회, 승인, 과금,
transaction 저장 등의 업무 로직은 라이브러리 사용자가 구현합니다.

```go
profile.HandleHeartbeat(heartbeatHandler)
profile.HandleStatusNotification(statusNotificationHandler)
profile.HandleAuthorize(authorizeHandler)
profile.HandleStartTransaction(startTransactionHandler)
profile.HandleMeterValues(meterValuesHandler)
profile.HandleStopTransaction(stopTransactionHandler)
profile.HandleDataTransfer(dataTransferHandler)
```

BootNotification 결과가 `Accepted`가 되기 전에는 다른 OCPP 1.6 inbound Action이
`ProtocolError`로 거부됩니다.

```go
state := profile.State(session)
booted := profile.IsBooted(session)
```

CSMS에서 시작하는 Action은 typed call로 호출합니다.

```go
confirmation, err := profile.CallReset(ctx, session, v16.ResetRequest{
    Type: v16.ResetRequestTypeSoft,
})
```

제공되는 outbound API:

- `CallDataTransfer`
- `CallReset`
- `CallRemoteStartTransaction`
- `CallRemoteStopTransaction`
- `CallChangeAvailability`
- `CallUnlockConnector`
- `CallTriggerMessage`

DataTransfer는 양방향이므로 `HandleDataTransfer`와 `CallDataTransfer`를 모두
제공합니다. Profile outbound call은 Boot가 완료되지 않은 세션에 `ErrNotBooted`를
반환합니다.

## OCPP 2.0.1 Core Profile

`profiles/ocpp201`은 OCPP 2.0.1의 Boot 등록 상태와 Core typed API를 제공합니다.
업무 데이터와 TransactionEvent 순서 정책은 애플리케이션 handler가 관리합니다.

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

Charging Station에서 시작하는 `Heartbeat`, `StatusNotification`, `Authorize`,
`TransactionEvent`, `MeterValues`, `NotifyReport` handler를 독립적으로 등록할 수
있습니다. CSMS에서는 `GetVariables`, `SetVariables`, `GetBaseReport`,
`RequestStartTransaction`, `RequestStopTransaction`, `Reset`을 typed call로 호출합니다.

Boot가 `Accepted` 상태가 되기 전에는 다른 inbound Action이 `ProtocolError`로
거부되고 outbound call은 `ocpp201.ErrNotBooted`를 반환합니다.

OCPP 2.1의 `NotifyPeriodicEventStream`은 응답 없는 SEND 메시지입니다. 일반 typed
handler와 분리된 API로 등록하며 CSMS는 CALLRESULT나 CALLERROR를 보내지 않습니다.

```go
err := profile21.HandleNotifyPeriodicEventStream(func(
    ctx context.Context,
    session *csms.Session,
    request v21.NotifyPeriodicEventStreamRequest,
) error {
    // stream 값을 저장하거나 전달합니다.
    return nil
})
```

OCPP 2.1 Profile은 Core 외에 다음 서버 API를 제공합니다.

- 스마트 충전과 에너지 전송
- 우선 충전과 동적 스케줄, aFRR 신호
- DER 제어와 상태 보고
- tariff, 비용, 정산, 웹 결제, VAT 검증
- 배터리 교환
- 주기 이벤트 스트림
- 인증서, ISO 15118 인증서와 firmware lifecycle

주기 이벤트 스트림의 방향은 다음과 같습니다.

- Charging Station → CSMS CALL: `OpenPeriodicEventStream`, `ClosePeriodicEventStream`
- CSMS → Charging Station CALL: `GetPeriodicEventStream`, `AdjustPeriodicEventStream`
- Charging Station → CSMS SEND: `NotifyPeriodicEventStream`

Profile은 typed 통신 경계와 Boot 상태만 관리합니다. 스케줄 계산, DER 상태, tariff와
정산 저장소, CA/OCSP/CRL, private key, firmware 파일 및 서명 검증은 라이브러리
사용자가 구현합니다.

## 공통 outbound call

Profile wrapper가 없는 Action은 `csms.Call`을 직접 사용할 수 있습니다.

```go
confirmation, err := csms.Call[v16.ResetRequest, v16.ResetConfirmation](
    ctx,
    session,
    v16.ResetRequest{Type: v16.ResetRequestTypeSoft},
)
```

원격 CALLERROR는 `*csms.RemoteCallError`로 반환됩니다. 호출 timeout과 최대 pending
수는 Server 설정으로 제어합니다.

```go
server, err := csms.New(csms.Config{
    CallTimeout:     30 * time.Second,
    MaxPendingCalls: 100,
})
```

사용자 Unique ID generator는 concurrency-safe해야 하며 빈 ID 또는 현재 pending
중인 중복 ID는 거부됩니다.

## 세션과 중복 연결

동일한 Charging Station identity가 이미 연결되어 있으면 기본적으로 새 연결을 HTTP
`409 Conflict`로 거부합니다.

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

교체 정책에서도 새 WebSocket upgrade가 성공한 후 기존 세션을 종료합니다.

```go
session, ok := server.Session("CP-001")
sessions := server.Sessions()
count := server.SessionCount()
```

## 연결 생명주기

```go
server, err := csms.New(csms.Config{
    WriteTimeout:          10 * time.Second,
    PingInterval:          30 * time.Second,
    PongTimeout:           90 * time.Second,
    IdleTimeout:           0,
    MaxConcurrentHandlers: 16,
})
```

`Session.Info()`에서 연결 시각, 최근 송수신, 최근 Pong과 상태를 확인할 수 있습니다.

```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

err := server.Shutdown(ctx)
```

Shutdown 중에는 새 연결을 HTTP `503 Service Unavailable`로 거부하고 활성 세션과
handler context를 종료합니다.

## 보안 프로필

지원하는 OCPP-J handshake 정책:

- Profile 0: 인증 없는 WebSocket
- Profile 1: Basic Authentication
- Profile 2: TLS 1.2 이상과 Basic Authentication
- Profile 3: TLS 1.2 이상과 검증된 client certificate

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

TLS handshake와 인증서 chain 검증은 외부 `http.Server`의 `tls.Config`가 담당합니다.
라이브러리는 검증된 TLS 상태를 확인하고 credentials를 Authenticator에 전달하며,
Session에는 비밀번호 대신 `Principal`만 저장합니다.

## API 호환성

이 프로젝트는 Semantic Versioning을 따릅니다. `v0.x`에서는 첫 안정 버전을 준비하는
과정에서 공개 API가 변경될 수 있으며, 변경 사항은 release note에 기록합니다. `v1.0.0`
이후에는 같은 major version 안에서 기존 공개 API의 source compatibility를 유지합니다.

`csms.Config`는 keyed field literal로 생성하는 것을 권장합니다. duration과 limit의 0 값은
문서화된 기본값을 선택하지만 `IdleTimeout: 0`은 idle 감지를 비활성화합니다. 음수 duration과
limit은 `csms.ErrInvalidConfiguration`으로 거부됩니다.
`Server.Shutdown`은 여러 번 호출해도 안전하지만 종료 후 서버를 다시 시작할 수는 없습니다.

애플리케이션은 raw WebSocket frame을 직접 전송하지 않고 `csms.Call` 또는 버전별 profile의
typed `Call...` 메서드를 사용합니다. 이 경로에서 pending CALL 등록, unique ID 검증과 응답
상관관계가 함께 관리됩니다.

분기 가능한 주요 오류:

- `csms.ErrInvalidConfiguration`
- `csms.ErrInvalidHandlerRegistration`
- `csms.ErrHandlerAlreadyRegistered`
- `csms.ErrUniqueIDGeneration`
- `csms.ErrDuplicateUniqueID`
- `csms.ErrTooManyPendingCalls`
- `csms.ErrSessionClosed`, `csms.ErrSessionReplaced`, `csms.ErrServerShutdown`

오류 비교에는 문자열 대신 `errors.Is`와 `errors.As`를 사용합니다. Charging Station이 보낸
CALLERROR는 `*csms.RemoteCallError`로 반환되며 `Code`는 `csms.ErrorCode`입니다.

## 검증

```sh
GOCACHE=/tmp/ocpp-go-build-cache go vet ./...
GOCACHE=/tmp/ocpp-go-build-cache go test ./...
GOCACHE=/tmp/ocpp-go-build-cache go test -race ./...
golangci-lint run ./...
```

WebSocket 통합 테스트는 로컬 loopback port를 사용합니다.

패키지 경계를 넘는 실제 실행 커버리지(생성 타입 365개 포함)는 다음으로 측정합니다.
2026-07-19 기준 66.5%이며, `csms`/`internal` 등 손으로 작성한 핵심 패키지는 개별적으로
80%대입니다. 생성 타입 365개(`v16`/`v201`/`v21`, `DO NOT EDIT` 헤더)와 `examples/`는
분기 로직이 사실상 없거나 테스트 대상이 아니라 로컬 측정치를 낮춰 보이게 만드므로,
Codecov 배지는 `codecov.yml`에서 이 경로들을 제외하고 집계한 값을 보여줍니다.

```sh
GOCACHE=/tmp/ocpp-go-build-cache go test -coverpkg=./... -coverprofile=coverage.out ./...
go tool cover -func=coverage.out | tail -1
```

일반 suite에는 16개 세션에서 총 800개 CALL을 처리하는 bounded load와 concurrent
outbound pending-call 상한 테스트가 포함됩니다. 반복적인 비정상 연결 종료 뒤 세션,
pending call, read loop, profile 상태와 goroutine이 해제되는 회귀 테스트도 포함됩니다.
장시간 soak는 opt-in이며 종료 전후 heap/goroutine 통계를 로그로 남깁니다.

```sh
OCPP_SOAK_DURATION=30m GOCACHE=/tmp/ocpp-go-build-cache go test -race \
  -timeout 35m -run '^TestSessionSoak$' ./csms
```

`go test`의 기본 `-timeout`은 10분이므로 `OCPP_SOAK_DURATION`보다 긴 `-timeout`을
반드시 함께 지정해야 합니다.

frame encode/decode, schema 검증, Router dispatch와 WebSocket loopback 왕복 전체
경로에 대한 벤치마크도 있습니다.

```sh
GOCACHE=/tmp/ocpp-go-build-cache go test -run '^$' -bench . -benchmem ./...
```

참고용 측정치입니다(Apple M2, `-benchtime=1s -count=3`의 중앙값; 실제 수치는
하드웨어와 시스템 부하에 따라 달라집니다).

| 벤치마크 | 시간/op | 메모리/op | 할당/op |
|---|---:|---:|---:|
| `DecodeCall` | 5.85 µs | 4.8 KB | 115 |
| `EncodeCallResult` | 1.55 µs | 969 B | 21 |
| `ValidateJSONSmall` (BootNotification) | 1.37 µs | 1.8 KB | 27 |
| `ValidateJSONLarge` (배열 100개) | 57.0 µs | 55.6 KB | 1026 |
| `ValidateThenUnmarshalSmall` | 2.41 µs | 2.2 KB | 37 |
| `ValidateThenUnmarshalLarge` | 95.1 µs | 65.9 KB | 1143 |
| `RouterLookup`, 미들웨어 0개 | 16.1 ns | 0 B | 0 |
| `RouterLookup`, 미들웨어 5개 | 122.2 ns | 128 B | 6 |
| `InboundCallRoundTrip` (실제 WebSocket loopback 전체 왕복) | 30.8 µs | 6.0 KB | 86 |
| `InboundCallRoundTrip`, `Config.Metrics` 설정 (약 0%, 측정 잡음 범위) | 30.8 µs | 6.0 KB | 86 |
| `OutboundCallRoundTrip` (`csms.Call` 실제 왕복) | 35.4 µs | 12.9 KB | 202 |
| `OutboundCallRoundTrip`, `Config.Metrics` 설정 (+2.9%) | 36.5 µs | 12.9 KB | 202 |

CALL 하나를 실제 네트워크까지 왕복시켜도 30µs대이므로, OCPP처럼 충전기 한 대가
몇 분에 한 번 메시지를 보내는 저빈도 프로토콜에서는 처리량이 병목이 될 가능성이
낮습니다. 배열이 큰 payload(예: `NotifyPeriodicEventStream`)를 아주 많은 세션이
동시에 보내는 배포에서는 `ValidateJSONLarge`/`ValidateThenUnmarshalLarge` 수치를
참고하세요.

`Config.Metrics`를 설정하면 고정 워커 풀로 이벤트를 dispatch한다. 같은 실행에서
측정한 결과 inbound CALL 차이는 측정 잡음 범위였고, outbound `csms.Call`은 약 2.9%
늘었으며 두 경로 모두 allocation 수는 변하지 않았다. 실제 비용은 `Metrics.Record`
구현에 따라 달라진다. `Config.Metrics`를 설정하지 않으면 dispatch 비용이 발생하지 않는다.

## 구조화 로그

특정 로깅 라이브러리 의존 없이 `csms.Logger`를 주입할 수 있습니다.

```go
server, err := csms.New(csms.Config{
    Router: router,
    Logger: csms.LoggerFunc(func(ctx context.Context, record csms.LogRecord) {
        // slog, zap, zerolog 등 애플리케이션 logger로 전달합니다.
    }),
})
```

기본 로그 레코드는 identity, OCPP version, message type, message ID, Action, error code와
정규화된 reason만 제공합니다. payload, Authorization header, 인증서 내용, idToken 및
handler 오류 문자열은 전달하지 않습니다. logger가 panic해도 프로토콜 서버는 계속
동작합니다.

## 메트릭 hook과 서버 상태

특정 metrics 라이브러리 의존 없이 `csms.Metrics`를 주입할 수 있습니다. 세션 연결/해제,
inbound CALL/SEND, outbound `csms.Call`의 각 단계에서 identity, OCPP version, message
type, Action, 소요 시간(`Duration`), error code를 담은 이벤트가 전달됩니다.

```go
server, err := csms.New(csms.Config{
    Router: router,
    Metrics: csms.MetricsFunc(func(ctx context.Context, event csms.MetricEvent) {
        // Prometheus counter/histogram 등 애플리케이션 metrics로 집계합니다.
    }),
})
```

`Metrics`도 `Logger`와 동일하게 panic이 프로토콜 서버에 영향을 주지 않습니다.
outbound `csms.Call`은 전송 성공(`MetricOutboundCallSent`)과 최종 결과(완료/실패/
timeout/취소/`MaxPendingCalls` 초과 거부)를 구분해서 관측할 수 있습니다.

`Metrics.Record`는 `Server` 하나당 고정된 개수의 워커 goroutine이 공유 큐를 소비하며
dispatch되어 호출자를 절대 블로킹하지 않습니다. 느리거나 멈춘 `Record` 구현체가 handler
처리나 outbound `csms.Call`의 취소 응답성을 지연시키지 않는다는 뜻입니다. 대신 `Record`는
동시 호출에 안전해야 하며, 이벤트 발생 순서를 보장하지 않습니다. 큐가 가득 차면(내부적으로
고정된 용량) 초과분 이벤트는 대기하지 않고 버려집니다 — 정상적인 `Record` 구현체라면 이
상한에 도달할 일이 없습니다. `Server.Shutdown`을 호출하면 이 워커들도 함께 정지합니다.

단, `Record` 구현체가 스스로 반환하지 않고 영원히 block되면 그 워커는 영구히 점유되어
pool의 실효 동시성이 하나 줄어든다 — handler ctx와 달리 이 워커 goroutine에는 별도의
취소/timeout 메커니즘이 없으므로, `Record`는 반드시 자체적으로 빠르게 반환하거나 내부
timeout을 둬야 한다(`Server.Shutdown`은 이런 워커를 기다리다 `ctx`가 만료되면
`Shutdown`도 timeout 오류로 끝난다).

서버 상태는 `Server.Snapshot()`과 `Server.Healthy()`로 조회합니다. 라이브러리는 HTTP
endpoint를 강제하지 않으므로 애플리케이션이 원하는 경로에 직접 연결합니다.

```go
mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(struct {
        Healthy bool               `json:"healthy"`
        Server  csms.ServerSnapshot `json:"server"`
    }{server.Healthy(), server.Snapshot()})
})
```

[`examples/metrics-hook`](examples/metrics-hook)는 in-process 카운터와 상태 endpoint를
함께 보여줍니다.

## Charging Station 클라이언트

`station` 패키지는 재연결/backoff을 갖춘 실제 OCPP-J Charging Station 클라이언트입니다.
`csms.Call`/`csms.Handle`과 동일한 방식으로 생성된 v16/v201/v21 타입에 대해 typed
outbound `station.Call`과 typed inbound `station.Handle`을 제공합니다.

```go
st, err := station.New(station.Config{
    URL: "wss://csms.example.com", Identity: "STATION-01", Version: protocol.OCPP16,
    ReconnectPolicy: &station.ReconnectPolicy{InitialDelay: time.Second, MaxDelay: 30 * time.Second, Multiplier: 2},
})
go st.Run(ctx) // 연결이 끊기면 ReconnectPolicy에 따라 자동 재연결

confirmation, err := station.Call[v16.BootNotificationRequest, v16.BootNotificationConfirmation](
    ctx, st, v16.BootNotificationRequest{ChargePointVendor: "Acme", ChargePointModel: "X1"},
)
```

`station.Handle`로 등록한 핸들러는 재연결 이후에도 유지됩니다. 재연결 중 끊긴 연결의
pending Call은 즉시 실패 처리되며(오프라인 큐잉 없음), Basic Auth/TLS(mTLS 포함)는
`Config.BasicAuth`/`Config.TLSConfig`로 설정합니다. 여러 charger를 동시에 운영하려면
호출자가 직접 `map[string]*station.Station`을 관리합니다 — 이 라이브러리는 프로토콜/세션
동작만 책임집니다. [`examples/station-client`](examples/station-client) 참고.

## 로드맵

- 유료 OCA OCTT와 공식 인증은 release 이후 선택적으로 수행

Prometheus/OpenTelemetry는 전용 훅 없이 기존 `csms.Metrics`/`csms.Router`
middleware/호출자 `ctx`만으로 연동 가능하다고 확인해 로드맵에서 뺐다 —
[`examples/prometheus-hook`](examples/prometheus-hook)과
[`examples/otel-hook`](examples/otel-hook) 참고.

## 변경 이력

[CHANGELOG.md](CHANGELOG.md)

## 라이선스

[MIT](LICENSE)
