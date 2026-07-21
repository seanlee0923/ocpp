# Changelog

**한국어** | [English](CHANGELOG.en.md)

이 프로젝트는 [Keep a Changelog](https://keepachangelog.com/ko/1.1.0/) 형식과
[시맨틱 버저닝](https://semver.org/lang/ko/)을 따릅니다. `v0.x`에서는 release
note를 통한 API 변경을 허용하며, `v1`부터 같은 major 내 source compatibility를
유지합니다.

## [Unreleased]

### Added

- opt-in `csms.Metrics`/`csms.MetricsFunc` hook — 세션 연결/해제, inbound
  CALL/SEND, outbound `csms.Call`의 각 단계(전송, 완료, 실패, timeout, 취소,
  `MaxPendingCalls` 초과 거부)를 identity, OCPP version, message type,
  Action, 소요 시간, error code로 관측한다. `Logger`와 동일하게 panic이
  프로토콜 서버에 영향을 주지 않는다. `Record`는 `Server`당 고정된 개수의 워커 goroutine이 공유 큐를 소비하며
  dispatch되어 호출자를 절대 블로킹하지 않는다 — 느리거나 멈춘 구현체도
  handler 처리나 outbound `csms.Call`의 취소 응답성을 지연시키지 않는다(대신
  `Record`는 동시 호출 안전성이 필요하고 순서를 보장하지 않으며, 큐가 가득 차면
  이벤트가 드롭된다). `Server.Shutdown`을 호출하면 이 워커들도 함께 정지한다.
  `Record`에는 항상 `context.Background()`가 전달된다 — 원래 컨텍스트를
  전달하면 timeout/cancel/shutdown을 보고하는 이벤트가 정확히 그 컨텍스트가
  이미 취소된 상태라, `ctx.Err() != nil`이면 아무 것도 안 하는 방어적
  `Record` 구현체가 그 이벤트들을 조용히 놓치게 되기 때문이다. `Config.Metrics`를
  설정하면 inbound CALL 왕복은 약 +6%, outbound `csms.Call` 왕복은 약 +13%
  오버헤드가 생긴다(설정하지 않으면 오버헤드 없음, 실측 벤치마크는 README 참고).
- `Server.Snapshot()`/`ServerSnapshot`과 `Server.Healthy()` — 활성 세션 수,
  세션별 상태, shutdown 여부를 조회하는 서버 상태 API. HTTP endpoint 형식은
  강제하지 않는다.
- [`examples/metrics-hook`](examples/metrics-hook) 예제 추가.
- `MetricEvent.Identity`에 cardinality 위험 경고를 doc comment로 명시하고,
  [`examples/prometheus-hook`](examples/prometheus-hook) 예제를 추가했다 —
  공식 Prometheus adapter 패키지는 만들지 않기로 했다(충전기 identity를
  label로 쓰는 게 안전한지는 배포 규모마다 다른 판단이라 라이브러리가 대신
  정할 수 없음). 예제는 `identity`를 label에서 제외하는 안전한 패턴을
  보여준다.
- [`examples/otel-hook`](examples/otel-hook) 예제를 추가했다 — OpenTelemetry
  tracing도 전용 훅 없이 기존 `csms.Router` middleware(inbound)와 호출자가
  직접 넘기는 `ctx`(outbound `csms.Call`)만으로 span 연결이 가능함을 실제로
  확인(inbound handler span과 그 안에서 실행한 outbound `Call` span이 부모-자식
  관계로 정확히 이어지는 것까지 검증)했다. `MetricEvent.Identity`와 달리 span
  attribute는 값마다 time series를 만들지 않아 identity를 붙여도 안전하다.
- README에 GoDoc, CI, Codecov 배지를 추가했다. CI가
  `go test -coverpkg=./... -coverprofile=coverage.out ./...`로 커버리지를
  측정해 Codecov에 업로드한다(`fail_ci_if_error: false`라 Codecov 계정
  연결 전에도 CI 자체는 실패하지 않는다). `codecov.yml`에서 생성 타입
  365개(`v16`/`v201`/`v21`)와 `examples/`를 커버리지 집계에서 제외해,
  배지가 실제 로직 커버리지에 가까운 값을 보여주도록 했다.
- CI에 `lint` job(`golangci-lint`, 기본 linter set)을 추가했다. Go Report
  Card는 2026-07-01부로 서비스 자체가 sunset되어(외부 요인) README 배지를
  제거하고 golangci-lint로 대체했다. 테스트 파일의 `defer conn.Close()`류는
  `.golangci.yml`에서 errcheck 예외 처리했다.
- `station` 패키지 — 재연결/backoff을 갖춘 실제 OCPP-J Charging Station
  클라이언트. `csms.Call`/`csms.Handle`과 동일한 방식으로 생성된
  v16/v201/v21 타입에 대한 typed outbound `station.Call`과 typed inbound
  `station.Handle`을 제공한다. `station.Handle`로 등록한 핸들러는 재연결
  이후에도 유지되며, 연결이 끊긴 pending Call은 즉시 실패 처리된다(오프라인
  큐잉 없음). Basic Auth/TLS(mTLS 포함)는 `Config.BasicAuth`/
  `Config.TLSConfig`로 설정한다. 여러 charger 운영은 호출자 책임(자체
  `map[string]*station.Station` 관리) — 이 패키지는 프로토콜/세션 동작만
  다룬다. [`examples/station-client`](examples/station-client) 예제 추가.
- `ocpp`의 생성 타입 365개가 공식 OCA JSON Schema와 실제로 일치하는지
  누구나 독립적으로 재검증할 수 있도록, from-scratch generator인
  [`ocpp-schema-gen`](https://github.com/seanlee0923/ocpp-schema-gen)을
  README "OCA Schema 출처" 절에 링크했다.
- handler/hook 관련 문서(`docs/handlers.md`, `docs/sessions.md`,
  `docs/security.md`, README)에 "`ctx` 취소는 신호일 뿐 강제 종료가 아니며,
  handler/hook은 반드시 빠르게 반환해야 한다"는 계약을 명시했다. `OnConnect`/
  `OnDisconnect`/`OnDuplicateSession`, `Authenticator`/`HandshakeLimiter`,
  `Metrics.Record`가 각각 멈췄을 때 정확히 어디까지 영향을 주는지도 정리했다.

### Fixed

- 세션 연결 시 `LogSessionConnected`/`MetricSessionConnected` 기록 시점을 read
  loop 시작과 `Config.OnConnect` 호출보다 앞으로 옮겼다. 이전에는 충전기가 업그레이드
  직후 바로 CALL을 보내거나 `OnConnect`에서 blocking `csms.Call`을 사용하면(둘 다
  지원되는 패턴) connected 이벤트보다 다른 이벤트가 먼저 관측될 수 있었다.
- outbound `csms.Call`에서 컨텍스트가 이미 만료된 상태로 전송이 실패하는 경우
  `MetricOutboundCallFailed`가 아니라 `MetricOutboundCallTimeout`/
  `MetricOutboundCallCanceled`로 정확히 분류하도록 수정했다.
- 원격 CALLERROR의 error code 추출에 raw 타입 단언 대신 `errors.As`를 사용하도록
  변경해, 향후 에러가 wrapping되어도 안전하게 동작하도록 했다.
- `golangci-lint` 도입 과정에서 드러난 사용하지 않는 함수(`hasSupportedSubprotocol`)를
  제거하고, `TestIPRateLimiter`의 의도치 않게 동일했던 `||` 조건식을 두 개의
  개별 assertion으로 분리했다(동작 자체는 원래도 올바랐음).
- `csms.Call[Request, *Confirmation]`처럼 confirmation 타입 인자에 포인터를
  넘기면 panic하던 결함을 `isNilType` 가드로 막아 등록 시점 오류로 바꿨다.
- `Router`가 CALL과 SEND로 등록된 handler를 구분하지 않던 결함을 수정했다 —
  충전기가 잘못된 메시지 타입으로 보내면 CALL이 영원히 응답받지 못하거나,
  계산된 confirmation이 조용히 버려질 수 있었다.
- CALL handler의 confirmation이 인코딩에 실패하면 응답 없이 조용히 끊기던
  것을, `InternalError` CALLERROR를 보내도록 수정했다.
- JSON pre-validator의 재귀 깊이에 상한(64)을 둬, 깊게 중첩된 payload로 인한
  스택 소진을 방지했다.
- inbound handler 실행 시간을 `Config.CallTimeout` 기준으로 bound했다 —
  이전에는 세션이 끊기거나 서버가 shutdown될 때만 취소됐다.
- `IPRateLimiter`가 추적하는 client 수에 상한(1024)을 두고 초과 시 가장 오래된
  항목을 evict하도록 했다 — 이전에는 서로 다른 IP가 계속 연결을 시도하면
  메모리가 무제한으로 늘어났다.
- 연결 activate가 경합에서 진 경우 shutdown 원인이 실제 원인 대신 일반 에러로
  기록되던 결함, 그리고 disconnect 사유를 `readLoop`의 raw 에러 대신 먼저
  기록된 진짜 원인(`Session.Err()`)에서 가져오지 않던 결함을 수정했다.
- read loop가 handler 동시 실행 슬롯이 가득 찼을 때 함께 블로킹되던 결함을
  논블로킹 admission(`pendingSlots`)으로 전환해 수정했다 — 가득 차면 CALL을
  즉시 거부한다(`ErrHandlerQueueFull`).
- `station.dispatch`가 inbound CALL마다 무제한으로 goroutine을 spawn하던 것을
  csms와 동일한 bounded pending queue로 전환했다.
- `station`의 outbound 전송에 실제 socket write deadline(`Config.WriteTimeout`,
  기본 10초)을 적용했다 — 이전에는 `ctx` timeout이 블로킹된 write 자체를
  풀어주지 못했다.
- `station`에서 handler가 반환한 `CallError`를 전송 전에 정규화(빈 필드 보정,
  길이 제한)하고, handler 실행 시간을 `Config.CallTimeout`으로 bound했다.
- `station.Server.Shutdown`이 metric dispatch 워커까지 기다렸다가 종료하도록
  수정했다(이전에는 워커가 종료 대상에서 누락돼 있었다).
- `station`이 텍스트가 아닌 WebSocket 프레임(바이너리 등)을 거부하도록 했다 —
  OCPP-J는 텍스트 프레임만 허용한다.
- `station`이 파싱할 수 없는 프레임을 무시하고 계속 읽던 것을, csms와 동일하게
  연결을 끊도록 수정했다.
- `station.New`의 설정 검증에 `Version.Valid()` 체크와
  `ReconnectPolicy.Multiplier`의 NaN/Inf 체크를 추가했다.

### Changed

- CALLERROR 코드 산출 로직을 OCPP 버전별로 흩어져 있던 것에서 공유 헬퍼로
  통합했다(동작 변경 없음, 유지보수성 개선).
- `Metrics.Record` dispatch 방식을 이벤트마다 별도 goroutine을 spawn하던 것에서
  `Server`당 고정된 워커 풀이 공유 큐를 소비하는 방식으로 바꿨다(위 Added 항목의
  설명도 갱신했다). 외부에서 관측 가능한 동작은 동일하다.

## [0.1.0] - 2026-07-18

첫 공개 릴리스입니다. Go와 [gorilla/websocket](https://github.com/gorilla/websocket)으로
작성된 OCPP-J CSMS WebSocket 서버 라이브러리로, OCPP 1.6, OCPP 2.0.1(Edition 4),
OCPP 2.1(Edition 2)을 하나의 전송 계층에서 지원합니다.

### Added

- 공통 OCPP-J 프레이밍(CALL/CALLRESULT/CALLERROR, OCPP 2.1 SEND)과 WebSocket CSMS
  서버, 세션 lifecycle(ping/pong, idle timeout, graceful shutdown, 중복 세션 정책)
- 공식 OCA JSON Schema 기반 생성 Go 타입 365개(OCPP 1.6 request/confirmation
  28/28, 2.0.1 64/64, 2.1 91/90)와 외부 validator 없이 동작하는 순수 Go 검증 코드
- OCPP 1.6 Core Profile: BootNotification/Heartbeat/StatusNotification/Authorize/
  StartTransaction/MeterValues/StopTransaction/DataTransfer 및 Reset 등 outbound 제어
- OCPP 2.0.1 Edition 4 Core Profile: BootNotification/TransactionEvent/
  MeterValues/NotifyReport 및 GetVariables/RequestStartTransaction 등 outbound 제어
- OCPP 2.1 Edition 2 Core Profile 및 확장 기능군: smart charging/에너지 전송,
  우선 충전·동적 스케줄, DER 제어, 요금·비용·결제, 배터리 교환, 주기 이벤트
  스트림(SEND), 인증서/PKI, firmware 업데이트
- Security Profile 0(Unsecured)/1(Basic Auth)/2(TLS+Basic Auth)/3(TLS+mTLS)와
  주입 가능한 `Authenticator`, `HandshakeLimiter`, 동시성 안전 built-in
  `IPRateLimiter`
- 페이로드·자격증명·인증서·idToken을 제외하는 라이브러리 중립 structured logger hook
- 핫패스 벤치마크(frame encode/decode, schema 검증, Router dispatch, WebSocket
  loopback 왕복) 및 30분 race soak 테스트 하네스
- GitHub Actions CI(`gofmt`/`go vet`/`go build`/`go test`/`go test -race`/
  `govulncheck`)
- 한국어 원본 + 영어 번역(`docs/en/`) 상세 문서 11종과 실행 가능한 예제
- MIT 라이선스

### Fixed

- OCPP 2.0.1/2.1 `customData`의 vendor 확장 필드가 marshal/unmarshal 시 유실되던
  결함(`internal/customdata` 도입)
- OCPP 2.1 wrapper 4건의 방향 오류(NotifyAllowedEnergyTransfer, GetTariffs,
  NotifyWebPaymentStarted, GetCertificateChainStatus)
- OCPP 1.6/2.0.1/2.1 전체 wrapper Handle/Call 방향 전수 검증 및 CALLERROR 코드
  버전별(OCPP 1.6 `FormationViolation`/`OccurenceConstraintViolation` vs 2.x
  `FormatViolation`/`OccurrenceConstraintViolation`) 분류
- GitHub Actions Linux 러너 `-race` 환경에서만 드러나던 WebSocket 세션 등록 확인
  테스트의 레이스
- `crypto/tls`의 Encrypted Client Hello 관련 CVE(`GO-2026-5856`)를 반영해 Go
  1.26.5로 툴체인 상향
- soak 테스트 실행 시 `-timeout` 플래그 누락 문서 오류

### Changed

- `profiles/ocpp16|ocpp201|ocpp21`의 `Call<Action>` 메서드 45개 내부 구현을
  공통 제네릭 헬퍼(`callBooted`)로 정리. 공개 메서드 이름·시그니처와 동작은
  변경되지 않았다.

[Unreleased]: https://github.com/seanlee0923/ocpp/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/seanlee0923/ocpp/releases/tag/v0.1.0
