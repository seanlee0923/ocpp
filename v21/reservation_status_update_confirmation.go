// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = ReservationStatusUpdateConfirmation{}

var schemaReservationStatusUpdateConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type ReservationStatusUpdateConfirmation struct {
	CustomData *ReservationStatusUpdateConfirmationCustomData `json:"customData,omitempty"`
}

type ReservationStatusUpdateConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (ReservationStatusUpdateConfirmation) ActionName() string { return "ReservationStatusUpdate" }

func (ReservationStatusUpdateConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (ReservationStatusUpdateConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (ReservationStatusUpdateConfirmation) SchemaName() string {
	return "ReservationStatusUpdateResponse.json"
}

func (message ReservationStatusUpdateConfirmation) Validate() error {
	return validation.Validate("ReservationStatusUpdateResponse.json", schemaReservationStatusUpdateConfirmation, message)
}

func (ReservationStatusUpdateConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ReservationStatusUpdateResponse.json", schemaReservationStatusUpdateConfirmation, data)
}
