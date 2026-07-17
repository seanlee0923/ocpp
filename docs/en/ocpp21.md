# OCPP 2.1

[한국어](../ocpp21.md) | **English**

The WebSocket subprotocol is `ocpp2.1`, and generated types are based on
the Edition 2 Schemas. Beyond the 2.0.1 Core API, it provides typed
wrappers for smart charging, dynamic schedules, DER, tariff/payment,
battery swap, periodic event streams, certificates, and firmware.

OCPP 2.1's SEND frame carries no response. The server currently handles the
Charging Station-to-CSMS `NotifyPeriodicEventStream` through either
`HandleNotifyPeriodicEventStream` or `csms.HandleSend`. A SEND handler
error cannot be turned into a CALLRESULT or CALLERROR.

The new Actions do not embed any business service either. DER control
policy, tariff calculation, payment providers, certificate CAs, and
firmware verification are wired to application infrastructure by the
handler. Using the generated requests and confirmations lets Actions
without a profile convenience wrapper be handled through the shared typed
API as well.
