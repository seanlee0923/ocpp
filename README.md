# ocpp-go

Go로 작성된 OCPP-J CSMS WebSocket 서버 라이브러리입니다. 하나의 전송 계층에서
OCPP 1.6, 2.0.1, 2.1을 처리하며 공식 JSON Schema 기반 메시지 타입과 타입 안전
라우팅을 제공합니다.

> 현재 개발 단계의 모듈 경로는 임시 값인 `ocpp-go`입니다. 공개 Go 모듈로 사용하기
> 전에 `github.com/seanlee0923/ocpp`로 변경할 예정입니다.

## 주요 기능

- Gorilla WebSocket 기반 OCPP-J 서버
- `ocpp1.6`, `ocpp2.0.1`, `ocpp2.1` 서브프로토콜 협상
- CALL, CALLRESULT, CALLERROR frame encode/decode
- OCPP 1.6, 2.0.1 Edition 4, 2.1 Edition 2 전체 생성 메시지 타입
- 공식 OCA JSON Schema에서 생성한 Go request/confirmation 검증
- Action별 typed inbound handler와 typed outbound call
- 충전기 identity 기반 세션과 중복 연결 정책
- Security Profile 0, 1, 2, 3 handshake 정책
- Ping/Pong, idle timeout, graceful shutdown, handler backpressure
- 기본 Unique ID generator와 사용자 generator 주입

## 요구 사항

- Go 1.26.4 이상
- OCPP-J WebSocket을 사용하는 Charging Station

주요 의존성:

- `github.com/gorilla/websocket`

## 지원 상태

| 영역 | OCPP 1.6 | OCPP 2.0.1 | OCPP 2.1 |
|---|---:|---:|---:|
| 생성 메시지 타입 | 완료 | 완료 | 완료 |
| 생성 Go 검증 코드 | 완료 | 완료 | 완료 |
| 공통 WebSocket 전송 | 완료 | 완료 | 완료 |
| Core Profile API | 완료 | 예정 | 예정 |

현재 우선순위는 CSMS 서버입니다. Charging Station 클라이언트는 서버 API가 안정화된
후 추가할 예정입니다.

## 빠른 시작

Router와 OCPP 1.6 Profile을 만든 뒤 같은 Router를 CSMS Server에 전달합니다.

```go
package main

import (
    "context"
    "log"
    "net/http"
    "time"

    "ocpp-go/csms"
    "ocpp-go/profiles/ocpp16"
    "ocpp-go/protocol"
    "ocpp-go/v16"
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

## 메시지 타입

버전별 타입은 `req`, `conf` 하위 패키지 없이 단일 패키지에 있습니다.

```go
import (
    "ocpp-go/v16"
    "ocpp-go/v201"
    "ocpp-go/v21"
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
잘못된 request는 `FormationViolation`, 잘못된 confirmation은 `InternalError`로
변환됩니다.

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

## 검증

```sh
GOCACHE=/tmp/ocpp-go-build-cache go vet ./...
GOCACHE=/tmp/ocpp-go-build-cache go test ./...
GOCACHE=/tmp/ocpp-go-build-cache go test -race ./...
```

WebSocket 통합 테스트는 로컬 loopback port를 사용합니다.

## 로드맵

- OCPP 2.0.1 Core Profile
- OCPP 2.1 Core 및 확장 Profile
- 구조화 로그, metrics, tracing
- OCA OCTT 적합성 검증
- Charging Station 클라이언트
