// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = CancelReservationRequest{}

var schemaCancelReservationRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reservationId": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"reservationId"}}

type CancelReservationRequest struct {
	CustomData    *CancelReservationRequestCustomData `json:"customData,omitempty"`
	ReservationID int                                 `json:"reservationId"`
}

type CancelReservationRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (CancelReservationRequest) ActionName() string { return "CancelReservation" }

func (CancelReservationRequest) Version() protocol.Version { return protocol.OCPP201 }

func (CancelReservationRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (CancelReservationRequest) SchemaName() string { return "CancelReservationRequest.json" }

func (message CancelReservationRequest) Validate() error {
	return validation.Validate("CancelReservationRequest.json", schemaCancelReservationRequest, message)
}

func (CancelReservationRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("CancelReservationRequest.json", schemaCancelReservationRequest, data)
}
