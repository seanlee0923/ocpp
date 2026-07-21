# OCPP 1.6

**한국어** | [English](en/ocpp16.md)

WebSocket subprotocol은 `ocpp1.6`이다. 이 프로젝트는 JSON over WebSocket인 OCPP-J만 지원하며
SOAP binding은 지원하지 않는다.

`profiles/ocpp16`은 BootNotification registration gate와 Core에서 자주 쓰는 inbound handler를
제공한다: Heartbeat, StatusNotification, Authorize, StartTransaction, MeterValues,
StopTransaction, DataTransfer. BootNotification이 Accepted가 되기 전 다른 등록 Action은
`ProtocolError`로 거부된다.

BootNotificationConfirmation 전송 뒤 실행되는 callback은 `HandleBootNotificationAfter`로
여러 개 등록할 수 있다. callback은 등록 순서대로 실행되며 각 error/panic은 격리되므로,
accepted Boot 이후 HeartbeatInterval, WebSocket ping 또는 vendor custom configuration을
`CallChangeConfiguration`으로 차례로 설정할 수 있다.

CSMS 방향 convenience call에는 ChangeConfiguration, GetConfiguration, Reset,
RemoteStartTransaction, RemoteStopTransaction, ChangeAvailability, UnlockConnector,
TriggerMessage와 DataTransfer가 있다. 나머지 생성 Action도
`csms.Handle` 또는 `csms.Call`로 사용할 수 있다.

OCPP 1.6 transaction ID는 CSMS 애플리케이션이 할당하는 integer다. idTag 승인, connector와
transaction 저장, reservation과 charging profile 정책은 handler에서 구현한다.
