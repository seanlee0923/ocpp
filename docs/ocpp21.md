# OCPP 2.1

**한국어** | [English](en/ocpp21.md)

WebSocket subprotocol은 `ocpp2.1`이며 생성 타입은 Edition 2 Schema를 기준으로 한다.
2.0.1 Core API 외에 smart charging, dynamic schedule, DER, tariff/payment, battery swap,
periodic event stream, certificate와 firmware 관련 typed wrapper를 제공한다.

OCPP 2.1의 SEND frame은 응답 없는 메시지다. 현재 서버는 Charging Station에서 CSMS로 오는
`NotifyPeriodicEventStream`을 `HandleNotifyPeriodicEventStream` 또는 `csms.HandleSend`로
처리한다. SEND handler 오류에는 CALLRESULT/CALLERROR를 보낼 수 없다.

새로운 Action도 업무 서비스를 내장하지 않는다. DER 제어 정책, tariff 계산, payment provider,
인증서 CA와 firmware 검증은 handler가 애플리케이션 인프라에 연결한다. 생성 request와
confirmation을 사용하면 profile convenience wrapper가 없는 Action도 공통 typed API로 처리할
수 있다.
