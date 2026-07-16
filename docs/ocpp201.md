# OCPP 2.0.1

WebSocket subprotocol은 `ocpp2.0.1`이며 생성 타입은 Edition 4 Schema를 기준으로 한다.
`profiles/ocpp201`은 BootNotification registration gate와 Heartbeat, StatusNotification,
Authorize, TransactionEvent, MeterValues, NotifyReport handler를 제공한다.

CSMS 방향 convenience call에는 GetVariables, SetVariables, GetBaseReport,
RequestStartTransaction, RequestStopTransaction, Reset이 있다. Device Model, monitoring,
display, reservation, smart charging, certificate와 firmware 등 wrapper가 없는 Action도 전체
생성 타입과 공통 typed API로 접근할 수 있다.

2.0.1에서는 1.6의 Start/StopTransaction 대신 TransactionEvent가 transaction lifecycle의
중심이다. transactionId, seqNo, eventType의 업무 일관성과 영속화는 애플리케이션 책임이다.
