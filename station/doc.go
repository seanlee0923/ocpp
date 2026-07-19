// Package station is a real, long-lived OCPP-J Charging Station client: it
// dials a CSMS, reconnects with backoff on disconnect, and provides typed
// outbound Call and typed inbound Handle built on the same generated
// v16/v201/v21 request/confirmation types and protocol codec csms uses.
//
// It intentionally does not manage multiple charging stations, queue
// messages while disconnected, or implement the OCPP-spec-precise
// BootNotification retry-interval backoff — those are either
// application-specific (which stations exist, what "offline" should mean
// for a given deployment) or out of scope for this pass. A caller running
// many stations holds its own map of *Station.
package station
