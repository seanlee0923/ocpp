# 운영 환경 구성

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

0은 기본값을 선택하므로 `Config`는 keyed field로 작성한다. handler는 전달된 context를 DB와
외부 요청에 전파하고 무제한 background goroutine을 만들지 않는다.

## Proxy와 TLS

TLS를 reverse proxy에서 종료한다면 현재 core는 proxy header를 인증된 client certificate로
해석하지 않는다. mTLS Security Profile 3은 CSMS 프로세스가 검증된 Go TLS state를 직접 받는
구성을 기준으로 한다. origin, identity와 인증 정책은 배포 환경에 맞게 제한한다.

## 관측과 민감 정보

`csms.Logger`를 애플리케이션 logger에 연결할 수 있다. 기본 metadata에는 payload, password,
certificate, idToken과 handler error text가 포함되지 않는다. logger callback panic은 서버에서
격리된다. metrics, snapshot과 tracing은 향후 opt-in hook으로 제공할 계획이다.

release 전에는 `go test -race ./...`와 opt-in soak를 실행하고 goroutine/heap 추이를 확인한다.
