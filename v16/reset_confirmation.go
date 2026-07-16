// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ResetConfirmation{}

var schemaResetConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected"}}}, Required: []string{"status"}}

type ResetConfirmation struct {
	Status ResetConfirmationStatus `json:"status"`
}

type ResetConfirmationStatus string

const (
	ResetConfirmationStatusAccepted ResetConfirmationStatus = "Accepted"
	ResetConfirmationStatusRejected ResetConfirmationStatus = "Rejected"
)

func (ResetConfirmation) ActionName() string { return "Reset" }

func (ResetConfirmation) Version() protocol.Version { return protocol.OCPP16 }

func (ResetConfirmation) Direction() protocol.PayloadDirection { return protocol.ConfirmationPayload }

func (ResetConfirmation) SchemaName() string { return "ResetResponse.json" }

func (message ResetConfirmation) Validate() error {
	return validation.Validate("ResetResponse.json", schemaResetConfirmation, message)
}

func (ResetConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ResetResponse.json", schemaResetConfirmation, data)
}
