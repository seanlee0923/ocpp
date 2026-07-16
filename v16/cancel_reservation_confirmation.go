// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = CancelReservationConfirmation{}

var schemaCancelReservationConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected"}}}, Required: []string{"status"}}

type CancelReservationConfirmation struct {
	Status CancelReservationConfirmationStatus `json:"status"`
}

type CancelReservationConfirmationStatus string

const (
	CancelReservationConfirmationStatusAccepted CancelReservationConfirmationStatus = "Accepted"
	CancelReservationConfirmationStatusRejected CancelReservationConfirmationStatus = "Rejected"
)

func (CancelReservationConfirmation) ActionName() string { return "CancelReservation" }

func (CancelReservationConfirmation) Version() protocol.Version { return protocol.OCPP16 }

func (CancelReservationConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (CancelReservationConfirmation) SchemaName() string { return "CancelReservationResponse.json" }

func (message CancelReservationConfirmation) Validate() error {
	return validation.Validate("CancelReservationResponse.json", schemaCancelReservationConfirmation, message)
}

func (CancelReservationConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("CancelReservationResponse.json", schemaCancelReservationConfirmation, data)
}
