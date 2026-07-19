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
서버가 shutdown되면 취소된다. panic은 CALL의 `InternalError`로 격리되어 서버를 종료하지
않지만 정상적인 오류 처리를 대신하지 않는다.

예상 가능한 OCPP 오류는 `*csms.CallError`로 반환한다.

```go
return confirmation, &csms.CallError{
    Code: csms.NotSupported,
    Description: "operation is disabled",
    Details: map[string]any{"policy": "maintenance"},
}
```

OCPP 2.1 SEND는 `csms.HandleSend`로 등록한다. SEND에는 프로토콜 응답이 없으므로 validation
또는 handler 오류는 `csms.Logger`와 `csms.Metrics`(`MetricSendDropped`)로만 관찰할 수 있다.

Router middleware는 등록 순서로 handler를 감싸며 여러 profile이 Router를 공유할 때는
전달받은 version을 반드시 확인해야 한다.
