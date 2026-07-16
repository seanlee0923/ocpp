// Package csms provides an OCPP-J Central System Management System server.
//
// Server owns WebSocket connections and session lifecycles. Router dispatches
// inbound CALL messages to application-provided typed handlers, while Call
// sends typed requests from the CSMS to a connected Charging Station.
package csms
