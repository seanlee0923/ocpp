# 운영 환경 구성

**한국어** | [English](en/production.md)

## HTTP와 종료

`csms.Server`를 `http.Server`의 Handler로 사용하고 프로세스 signal에서 먼저
`csms.Server.Shutdown`, 이어서 `http.Server.Shutdown`을 호출한다. shutdown context에는
업무 handler가 정리될 충분한 deadline을 둔다.

## 제한과 timeout

- `ReadLimit`: 비정상적으로 큰 WebSocket message 제한
- `HandshakeTimeout`: upgrade 제한
- `CallTimeout`: deadline 없는 outbound CALL 제한
- `MaxPendingCalls`: 세션별 outbound 대기 요청 상한
- `MaxConcurrentHandlers`: 세션별 inbound handler 동시 실행 상한
- `WriteTimeout`, `PingInterval`, `PongTimeout`, `IdleTimeout`: 연결 생명주기

0은 문서화된 기본값을 선택하지만 `IdleTimeout: 0`은 idle 감지를 비활성화한다. 음수 duration과
limit은 잘못된 설정으로 거부된다. `Config`는 keyed field로 작성하고 handler는 전달된 context를
DB와 외부 요청에 전파하며 무제한 background goroutine을 만들지 않는다.

## Proxy와 TLS

TLS를 reverse proxy에서 종료한다면 현재 core는 proxy header를 인증된 client certificate로
해석하지 않는다. mTLS Security Profile 3은 CSMS 프로세스가 검증된 Go TLS state를 직접 받는
구성을 기준으로 한다. origin, identity와 인증 정책은 배포 환경에 맞게 제한한다.

## 관측과 민감 정보

`csms.Logger`를 애플리케이션 logger에 연결할 수 있다. 기본 metadata에는 payload, password,
certificate, idToken과 handler error text가 포함되지 않는다. logger callback panic은 서버에서
격리된다.

`csms.Metrics`는 같은 원칙으로 세션 연결/해제, inbound CALL/SEND, outbound `csms.Call`의
각 단계를 관측한다. identity, OCPP version, message type, Action, 소요 시간, error code만
전달하며 payload는 포함하지 않는다. panic은 `Logger`와 동일하게 격리된다. outbound
`csms.Call`은 전송(`MetricOutboundCallSent`)과 최종 결과(완료/실패/timeout/취소/
`MaxPendingCalls` 초과 거부)를 구분해서 관측할 수 있다. `Metrics.Record`는 이벤트마다
별도 goroutine에서 dispatch되어 호출자를 블로킹하지 않는다 — 느리거나 멈춘 구현체가
handler 처리나 outbound `csms.Call`의 취소 응답성을 지연시키지 않는다. 대신 `Record`는
동시 호출에 안전해야 하고 이벤트 도착 순서를 보장하지 않으며, 동시 dispatch 수가 내부
고정 상한을 넘으면 초과분은 버려진다(정상 구현체는 이 상한에 도달하지 않는다).
`Record`에는 항상 `context.Background()`가 전달된다 — timeout/cancel/shutdown을
보고하는 이벤트가 정확히 그 시점엔 원래 context가 이미 취소돼 있을 수 있어서, 원래
context를 그대로 넘기면 `ctx.Err() != nil`일 때 아무 것도 안 하는 방어적 `Record`
구현체가 그 이벤트들을 조용히 놓치게 된다. `Config.Metrics`를 설정하면 inbound CALL
왕복은 약 +6%, outbound `csms.Call` 왕복은 약 +13% 오버헤드가 생긴다(설정하지 않으면
오버헤드 없음, 실측 벤치마크는 루트 [README](../README.md)의 벤치마크 표 참고).

`Server.Snapshot()`과 `Server.Healthy()`로 활성 세션 수, 세션별 상태, shutdown 여부를
조회할 수 있다. 라이브러리는 HTTP endpoint 형식을 강제하지 않으므로 health/readiness
probe에 필요한 경로와 payload 형식은 애플리케이션이 직접 정의한다.

opt-in Prometheus adapter와 OpenTelemetry tracing hook은 아직 제공하지 않는다(루트
[README](../README.md)의 로드맵 참고).

[`examples/structured-logger`](../examples/structured-logger)는 표준 `log/slog` 연결을,
[`examples/metrics-hook`](../examples/metrics-hook)은 in-process metrics 카운터와 상태
endpoint를, [`examples/graceful-shutdown`](../examples/graceful-shutdown)은 signal 기반
종료를 보여준다.

release 전에는 `go test -race ./...`와 opt-in soak를 실행하고 goroutine/heap 추이를 확인한다.
