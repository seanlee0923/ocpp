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
  프로토콜 서버에 영향을 주지 않는다. `Record`는 이벤트마다 별도 goroutine에서
  dispatch되어 느리거나 멈춘 구현체도 handler 처리나 outbound `csms.Call`의
  취소 응답성을 지연시키지 않는다(대신 `Record`는 동시 호출 안전성이 필요하고
  순서를 보장하지 않으며, 동시 dispatch 상한 초과 시 이벤트가 드롭된다).
- `Server.Snapshot()`/`ServerSnapshot`과 `Server.Healthy()` — 활성 세션 수,
  세션별 상태, shutdown 여부를 조회하는 서버 상태 API. HTTP endpoint 형식은
  강제하지 않는다.
- [`examples/metrics-hook`](examples/metrics-hook) 예제 추가.

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
