// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = NotifyDisplayMessagesConfirmation{}

var schemaNotifyDisplayMessagesConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type NotifyDisplayMessagesConfirmation struct {
	CustomData *NotifyDisplayMessagesConfirmationCustomData `json:"customData,omitempty"`
}

type NotifyDisplayMessagesConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (NotifyDisplayMessagesConfirmation) ActionName() string { return "NotifyDisplayMessages" }

func (NotifyDisplayMessagesConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (NotifyDisplayMessagesConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (NotifyDisplayMessagesConfirmation) SchemaName() string {
	return "NotifyDisplayMessagesResponse.json"
}

func (message NotifyDisplayMessagesConfirmation) Validate() error {
	return validation.Validate("NotifyDisplayMessagesResponse.json", schemaNotifyDisplayMessagesConfirmation, message)
}

func (NotifyDisplayMessagesConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("NotifyDisplayMessagesResponse.json", schemaNotifyDisplayMessagesConfirmation, data)
}
