# OCPP 2.0.1

[한국어](../ocpp201.md) | **English**

The WebSocket subprotocol is `ocpp2.0.1`, and generated types are based on
the Edition 4 Schemas. `profiles/ocpp201` provides the BootNotification
registration gate along with Heartbeat, StatusNotification, Authorize,
TransactionEvent, MeterValues, and NotifyReport handlers.

CSMS-directed convenience calls include GetVariables, SetVariables,
GetBaseReport, RequestStartTransaction, RequestStopTransaction, and Reset.
Actions without a wrapper — Device Model, monitoring, display, reservation,
smart charging, certificates, firmware, and more — are still reachable
through the full set of generated types and the shared typed API.

In 2.0.1, TransactionEvent is the center of the transaction lifecycle,
replacing 1.6's Start/StopTransaction. Business-level consistency and
persistence of transactionId, seqNo, and eventType are the application's
responsibility.
