// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = UnlockConnectorConfirmation{}

var schemaUnlockConnectorConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Unlocked", "UnlockFailed", "NotSupported"}}}, Required: []string{"status"}}

type UnlockConnectorConfirmation struct {
	Status UnlockConnectorConfirmationStatus `json:"status"`
}

type UnlockConnectorConfirmationStatus string

const (
	UnlockConnectorConfirmationStatusUnlocked     UnlockConnectorConfirmationStatus = "Unlocked"
	UnlockConnectorConfirmationStatusUnlockFailed UnlockConnectorConfirmationStatus = "UnlockFailed"
	UnlockConnectorConfirmationStatusNotSupported UnlockConnectorConfirmationStatus = "NotSupported"
)

func (UnlockConnectorConfirmation) ActionName() string { return "UnlockConnector" }

func (UnlockConnectorConfirmation) Version() protocol.Version { return protocol.OCPP16 }

func (UnlockConnectorConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (UnlockConnectorConfirmation) SchemaName() string { return "UnlockConnectorResponse.json" }

func (message UnlockConnectorConfirmation) Validate() error {
	return validation.Validate("UnlockConnectorResponse.json", schemaUnlockConnectorConfirmation, message)
}

func (UnlockConnectorConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("UnlockConnectorResponse.json", schemaUnlockConnectorConfirmation, data)
}
