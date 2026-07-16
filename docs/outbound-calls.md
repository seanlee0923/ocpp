# Outbound CALL

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

BootNotification 이후 실제 session으로 Reset을 호출하는 흐름은
[`examples/outbound-call`](../examples/outbound-call)에 있다.
