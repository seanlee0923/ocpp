# Outbound CALL

**한국어** | [English](en/outbound-calls.md)

CSMS가 Charging Station으로 시작하는 요청은 profile method 또는 `csms.Call`을 사용한다.

```go
confirmation, err := csms.Call[v201.ResetRequest, v201.ResetConfirmation](
    ctx,
    session,
    v201.ResetRequest{Type: v201.ResetRequestResetEnumImmediate},
)
```

`Call`은 request 검증, unique ID 생성, pending 등록, 직렬화, 응답 correlation과 confirmation
검증을 한 번에 수행한다. request와 session의 OCPP version이 다르면 전송하지 않는다.

호출 context에 deadline이 없으면 `Config.CallTimeout`이 적용된다. 동시 대기 요청은
`Config.MaxPendingCalls`로 제한되며 초과 시 `ErrTooManyPendingCalls`가 반환된다. 세션이
종료되면 모든 pending CALL은 해당 종료 원인으로 즉시 해제된다.

Charging Station의 CALLERROR는 `*csms.RemoteCallError`다.

```go
var remote *csms.RemoteCallError
if errors.As(err, &remote) {
    log.Printf("code=%s description=%s", remote.Code, remote.Description)
}
```

기본 unique ID는 `uuid.NewString`이다. 사용자 generator는 `Config.UniqueIDGenerator`로
주입할 수 있다. panic, 빈 값, 36자 초과는 `ErrUniqueIDGeneration`, 같은 세션의 in-flight
중복은 `ErrDuplicateUniqueID`로 반환된다.

`Config.Metrics`를 설정하면 `Call`의 전 단계(전송, 완료, 원격 오류·응답 디코드 실패,
timeout, 취소, `MaxPendingCalls` 초과 거부)가 `MetricOutboundCallSent`/`Completed`/
`Failed`/`Timeout`/`Canceled`/`Rejected` 이벤트로 관측된다. 자세한 내용은
[운영 환경 구성](production.md)의 "관측과 민감 정보"를 참고한다.

`Call`을 부르는 쪽이 이미 `ctx`를 직접 넘기므로, OpenTelemetry tracing도 호출자가
`Call` 전에 자기 span을 만들어 그 `ctx`로 부르기만 하면 된다 — 라이브러리 쪽 지원이
따로 필요 없다. [`examples/otel-hook`](../examples/otel-hook)이 이 패턴을 보여준다.

BootNotification 이후 실제 session으로 Reset을 호출하는 흐름은
[`examples/outbound-call`](../examples/outbound-call)에 있다.

## OCPP handler 밖에서 identity로 호출하기

HTTP 요청, message queue 또는 scheduler처럼 OCPP handler 밖에서 이벤트를 받는 경우
`*csms.Session`을 별도 map이나 DB에 장기 저장하지 않는다. 재연결되면 같은 identity에도
새 `Session`이 등록되므로, CALL 직전에 `Server.Session(identity)`로 현재 연결을 조회한다.

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

`Session` 조회 성공과 CALL 완료 사이에도 연결은 끊기거나 동일 identity의 새 연결로
교체될 수 있다. 따라서 `ErrSessionClosed`, `ErrSessionReplaced`, context timeout과
`*RemoteCallError`는 정상적인 운영 오류로 처리해야 한다. Profile method를 사용하면 해당
profile의 정책도 적용된다. 예를 들어 OCPP 1.6 `Profile.CallDataTransfer`는 accepted
BootNotification 이전에 `ocpp16.ErrNotBooted`를 반환한다.

OCPP WebSocket 서버(`:8080`)와 외부 HTTP API(`:8081`)를 함께 실행하고 JSON body를
OCPP 1.6 `DataTransfer`로 전달하는 전체 예제는
[`examples/external-data-transfer`](../examples/external-data-transfer)에 있다.
