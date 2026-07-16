// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = ChangeAvailabilityConfirmation{}

var schemaChangeAvailabilityConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "Scheduled"}}}, Required: []string{"status"}}

type ChangeAvailabilityConfirmation struct {
	Status ChangeAvailabilityConfirmationStatus `json:"status"`
}

type ChangeAvailabilityConfirmationStatus string

const (
	ChangeAvailabilityConfirmationStatusAccepted  ChangeAvailabilityConfirmationStatus = "Accepted"
	ChangeAvailabilityConfirmationStatusRejected  ChangeAvailabilityConfirmationStatus = "Rejected"
	ChangeAvailabilityConfirmationStatusScheduled ChangeAvailabilityConfirmationStatus = "Scheduled"
)

func (ChangeAvailabilityConfirmation) ActionName() string { return "ChangeAvailability" }

func (ChangeAvailabilityConfirmation) Version() protocol.Version { return protocol.OCPP16 }

func (ChangeAvailabilityConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (ChangeAvailabilityConfirmation) SchemaName() string { return "ChangeAvailabilityResponse.json" }

func (message ChangeAvailabilityConfirmation) Validate() error {
	return validation.Validate("ChangeAvailabilityResponse.json", schemaChangeAvailabilityConfirmation, message)
}

func (ChangeAvailabilityConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ChangeAvailabilityResponse.json", schemaChangeAvailabilityConfirmation, data)
}
