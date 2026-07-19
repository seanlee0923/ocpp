# Examples

각 디렉터리는 독립적으로 컴파일되는 작은 예제다.

- `ocpp16-csms`: OCPP 1.6 BootNotification
- `ocpp201-csms`: OCPP 2.0.1 BootNotification
- `ocpp21-csms`: OCPP 2.1 BootNotification과 SEND handler
- `multi-version`: 세 OCPP 버전을 한 서버에서 처리
- `outbound-call`: CSMS에서 시작하는 typed Reset CALL
- `duplicate-session`: 기존 identity 연결 교체 정책
- `basic-auth`: Security Profile 1 인증
- `tls-basic-auth`: Security Profile 2 TLS와 Basic Auth
- `mtls`: Security Profile 3 mutual TLS
- `structured-logger`: `log/slog` JSON logger 연결
- `metrics-hook`: in-process Metrics 카운터와 Snapshot/Healthy 상태 endpoint
- `prometheus-hook`: `csms.Metrics`를 Prometheus counter/histogram으로 직접 연결
  (공식 adapter 없음, cardinality-safe label 예시)
- `otel-hook`: `csms.Router` middleware만으로 OpenTelemetry span 연결(inbound)과
  호출자 ctx로 outbound span 연결(`csms.Call`), 전용 tracing 훅 없이 동작
- `graceful-shutdown`: signal 기반 세션 및 HTTP 종료
- `station-client`: `station` 패키지로 만든 재연결 가능한 Charging Station
  클라이언트 — BootNotification/Heartbeat 전송과 inbound Reset 처리

```sh
go run ./examples/ocpp201-csms
```

TLS 예제는 인증서를 저장소에 포함하지 않는다. `tls-basic-auth`는 `TLS_CERT_FILE`과
`TLS_KEY_FILE`, `mtls`는 여기에 `TLS_CLIENT_CA_FILE`을 추가로 요구한다.

예제는 API 사용법에 집중하기 위해 in-memory 또는 고정 정책만 사용한다. 실제 인증,
transaction, Device Model과 과금 데이터는 애플리케이션 DB와 서비스에 연결한다.
