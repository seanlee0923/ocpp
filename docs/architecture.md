# 구조와 책임 범위

```text
Charging Station
      │ OCPP-J WebSocket
      ▼
  csms.Server ── session, timeout, CALL correlation, shutdown
      │
      ▼
  csms.Router ── version + Action dispatch, middleware
      │
      ▼
 Typed handler ── generated request / confirmation
      │
      ▼
Application ── DB, cache, queue, authorization, billing, PKI
```

라이브러리는 WebSocket, OCPP-J frame, schema 검증, typed routing, session lifecycle과
동시성 제어를 담당한다. 애플리케이션은 충전기·EVSE·connector·transaction 저장, idToken
승인, 예약, smart charging 정책, 과금, 인증서 발급과 firmware 배포를 담당한다.

`v16`, `v201`, `v21`은 공식 Schema에서 생성된 단일 패키지다. request와 confirmation은
`protocol.Payload`를 구현한다. `profiles/ocpp16`, `profiles/ocpp201`, `profiles/ocpp21`은
자주 쓰는 Action과 BootNotification registration gate를 제공하는 선택적 convenience
계층이다. wrapper가 없는 Action도 `csms.Handle`과 `csms.Call`로 사용할 수 있다.

`Session`과 내부 WebSocket은 서버가 소유한다. 애플리케이션은 raw frame을 쓰지 않고 typed
API를 사용해야 pending CALL, unique ID와 응답 상관관계가 유지된다.
