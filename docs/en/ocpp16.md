# OCPP 1.6

[한국어](../ocpp16.md) | **English**

The WebSocket subprotocol is `ocpp1.6`. This project supports only OCPP-J
(JSON over WebSocket); SOAP binding is not supported.

`profiles/ocpp16` provides the BootNotification registration gate and
commonly used Core inbound handlers: Heartbeat, StatusNotification,
Authorize, StartTransaction, MeterValues, StopTransaction, and
DataTransfer. Before BootNotification is Accepted, other registered Actions
are rejected with `ProtocolError`.

Multiple callbacks that run after the BootNotificationConfirmation write can
be registered with `HandleBootNotificationAfter`. They run in registration
order and isolate each error or panic, allowing accepted Boot to be followed
by HeartbeatInterval, WebSocket ping, or vendor-specific configuration calls.

CSMS-directed convenience calls include ChangeConfiguration,
GetConfiguration, Reset, RemoteStartTransaction, RemoteStopTransaction,
ChangeAvailability, UnlockConnector, TriggerMessage, and DataTransfer. Any other generated Action can also be
used through `csms.Handle` or `csms.Call`.

The OCPP 1.6 transaction ID is an integer assigned by the CSMS application.
idTag authorization, connector and transaction storage, and reservation and
charging profile policy are implemented in handlers.
