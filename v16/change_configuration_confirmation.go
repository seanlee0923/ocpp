// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = ChangeConfigurationConfirmation{}

var schemaChangeConfigurationConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "RebootRequired", "NotSupported"}}}, Required: []string{"status"}}

type ChangeConfigurationConfirmation struct {
	Status ChangeConfigurationConfirmationStatus `json:"status"`
}

type ChangeConfigurationConfirmationStatus string

const (
	ChangeConfigurationConfirmationStatusAccepted       ChangeConfigurationConfirmationStatus = "Accepted"
	ChangeConfigurationConfirmationStatusRejected       ChangeConfigurationConfirmationStatus = "Rejected"
	ChangeConfigurationConfirmationStatusRebootRequired ChangeConfigurationConfirmationStatus = "RebootRequired"
	ChangeConfigurationConfirmationStatusNotSupported   ChangeConfigurationConfirmationStatus = "NotSupported"
)

func (ChangeConfigurationConfirmation) ActionName() string { return "ChangeConfiguration" }

func (ChangeConfigurationConfirmation) Version() protocol.Version { return protocol.OCPP16 }

func (ChangeConfigurationConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (ChangeConfigurationConfirmation) SchemaName() string { return "ChangeConfigurationResponse.json" }

func (message ChangeConfigurationConfirmation) Validate() error {
	return validation.Validate("ChangeConfigurationResponse.json", schemaChangeConfigurationConfirmation, message)
}

func (ChangeConfigurationConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ChangeConfigurationResponse.json", schemaChangeConfigurationConfirmation, data)
}
