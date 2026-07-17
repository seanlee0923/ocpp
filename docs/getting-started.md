# 시작하기

**한국어** | [English](en/getting-started.md)

이 라이브러리는 OCPP-J WebSocket CSMS 서버를 제공한다. OCPP 1.6 SOAP 전송은 지원하지
않는다. 애플리케이션은 handler 안에서 DB, cache, queue와 업무 규칙을 구현한다.

## 설치

```sh
go get github.com/seanlee0923/ocpp
```

프로젝트는 Go 1.26.5 이상을 대상으로 한다.

## 최소 서버

```go
router := csms.NewRouter()
profile, err := ocpp16.NewProfile(router)
if err != nil {
    return err
}
err = profile.HandleBootNotification(func(
    ctx context.Context,
    session *csms.Session,
    request v16.BootNotificationRequest,
) (v16.BootNotificationConfirmation, error) {
    return v16.BootNotificationConfirmation{
        CurrentTime: time.Now().UTC().Format(time.RFC3339),
        Interval: 300,
        Status: v16.BootNotificationConfirmationStatusAccepted,
    }, nil
})
if err != nil {
    return err
}
server, err := csms.New(csms.Config{
    Router: router,
    Versions: []protocol.Version{protocol.OCPP16},
})
if err != nil {
    return err
}
return http.ListenAndServe(":8080", server)
```

Charging Station은 identity를 URL path로 보내고 OCPP subprotocol을 협상한다.

```text
GET /CP-001
Upgrade: websocket
Sec-WebSocket-Protocol: ocpp1.6
```

지원 subprotocol은 `ocpp1.6`, `ocpp2.0.1`, `ocpp2.1`이다. `Versions`를 생략하면 세
버전을 모두 허용한다. 동일 Router에 버전별 profile과 공통 typed handler를 함께 등록할
수 있다.

전체 실행 코드는 [`examples/ocpp16-csms`](../examples/ocpp16-csms)에 있다.
세 버전을 한 process에서 받는 구성은 [`examples/multi-version`](../examples/multi-version)을
참고한다.
