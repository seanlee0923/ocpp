# Inbound handler

**한국어** | [English](en/handlers.md)

Charging Station이 시작하는 CALL은 generic typed handler 또는 profile convenience method로
등록한다.

```go
err := csms.Handle(router, func(
    ctx context.Context,
    session *csms.Session,
    request v201.AuthorizeRequest,
) (v201.AuthorizeConfirmation, error) {
    return authorize(ctx, session.Identity(), request)
})
```

제네릭 타입에서 version, Action, request/confirmation 방향을 추론한다. 서로 맞지 않는
타입, nil handler와 동일 `version + Action` 중복 등록은 등록 시 오류가 된다. 동적 handler
교체는 지원하지 않는다.

입력 payload는 handler 실행 전에 공식 Schema 제약으로 검증된다. handler가 반환한
confirmation도 전송 전에 검증된다. handler는 `ctx` 취소를 존중해야 하며 세션이 종료되거나
서버가 shutdown되면, 또는 `Config.CallTimeout`이 지나면 취소된다. panic은 CALL의
`InternalError`로 격리되어 서버를 종료하지 않지만 정상적인 오류 처리를 대신하지 않는다.

`ctx` 취소는 신호일 뿐 강제 종료가 아니다 — Go에는 goroutine을 외부에서 강제로 멈추는
수단이 없으므로, handler가 취소를 무시하고 계속 실행되면 라이브러리도 이를 막을 수 없다.
이 경우 handler는 이미 시간이 초과된 요청에 대해 계속 실행되며, 그동안 handler 동시 실행
슬롯 하나를 계속 점유한다 — 슬롯이 고갈되면 이후 CALL은 `pendingSlots`/handler 큐가 가득
찬 것으로 취급되어 거부된다. `station.Handle`로 등록하는 Charging Station 쪽 handler도
동일한 계약을 따르며, `station.Config.CallTimeout`이 같은 역할을 한다.

예상 가능한 OCPP 오류는 `*csms.CallError`로 반환한다.

```go
return confirmation, &csms.CallError{
    Code: csms.NotSupported,
    Description: "operation is disabled",
    Details: map[string]any{"policy": "maintenance"},
}
```

성공한 CALLRESULT가 WebSocket에 기록된 뒤 후속 작업이 필요하면 `csms.HandleAfter`를
사용한다. 동일 Action에 여러 callback을 등록할 수 있고 등록 순서대로 실행된다. 각
callback에는 세션과 `Config.CallTimeout`에 연결된 새 context가 주어진다. error 또는
panic은 `LogCallAfterFailed`로 기록되고 다음 callback은 계속 실행된다. 이미 응답을 보낸
뒤이므로 After 오류를 CALLERROR로 바꿀 수는 없다. main handler 오류나 response write
실패 시에는 After가 실행되지 않는다.

After는 main handler를 실행한 것과 같은 goroutine에서 동기·순차 실행되어 해당 세션의
`MaxConcurrentHandlers` 슬롯 하나를 계속 점유한다. callback마다 새 `CallTimeout` 예산을
받으므로 여러 callback 또는 callback 내부의 여러 blocking CALL은 timeout이 합산될 수 있다.
각 CALL에 더 짧은 deadline을 설정하더라도 누적 구조는 그대로이므로 callback과 순차 CALL
개수를 작게 유지한다. 긴 작업은 애플리케이션 소유의 bounded worker queue로 전달한다.

OCPP 2.1 SEND는 `csms.HandleSend`로 등록한다. SEND에는 프로토콜 응답이 없으므로 validation
또는 handler 오류는 `csms.Logger`와 `csms.Metrics`(`MetricSendDropped`)로만 관찰할 수 있다.

Router middleware는 등록 순서로 handler를 감싸며 여러 profile이 Router를 공유할 때는
전달받은 version을 반드시 확인해야 한다. middleware는 `ctx`를 새로 만들어 다음 handler에
넘길 수 있으므로, OpenTelemetry 같은 tracing도 전용 훅 없이 middleware만으로 연결할 수
있다 — [`examples/otel-hook`](../examples/otel-hook)을 참고한다.
