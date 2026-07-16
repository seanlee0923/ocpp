// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = TriggerMessageConfirmation{}

var schemaTriggerMessageConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "NotImplemented"}}}, Required: []string{"status"}}

type TriggerMessageConfirmation struct {
	Status TriggerMessageConfirmationStatus `json:"status"`
}

type TriggerMessageConfirmationStatus string

const (
	TriggerMessageConfirmationStatusAccepted       TriggerMessageConfirmationStatus = "Accepted"
	TriggerMessageConfirmationStatusRejected       TriggerMessageConfirmationStatus = "Rejected"
	TriggerMessageConfirmationStatusNotImplemented TriggerMessageConfirmationStatus = "NotImplemented"
)

func (TriggerMessageConfirmation) ActionName() string { return "TriggerMessage" }

func (TriggerMessageConfirmation) Version() protocol.Version { return protocol.OCPP16 }

func (TriggerMessageConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (TriggerMessageConfirmation) SchemaName() string { return "TriggerMessageResponse.json" }

func (message TriggerMessageConfirmation) Validate() error {
	return validation.Validate("TriggerMessageResponse.json", schemaTriggerMessageConfirmation, message)
}

func (TriggerMessageConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("TriggerMessageResponse.json", schemaTriggerMessageConfirmation, data)
}
