# Examples

각 디렉터리는 독립적으로 컴파일되는 작은 예제다.

- `ocpp16-csms`: OCPP 1.6 BootNotification
- `ocpp201-csms`: OCPP 2.0.1 BootNotification
- `ocpp21-csms`: OCPP 2.1 BootNotification과 SEND handler
- `basic-auth`: Security Profile 1 인증
- `graceful-shutdown`: signal 기반 세션 및 HTTP 종료

```sh
go run ./examples/ocpp201-csms
```

예제는 API 사용법에 집중하기 위해 in-memory 또는 고정 정책만 사용한다. 실제 인증,
transaction, Device Model과 과금 데이터는 애플리케이션 DB와 서비스에 연결한다.
