// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = NotifyEventConfirmation{}

var schemaNotifyEventConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type NotifyEventConfirmation struct {
	CustomData *NotifyEventConfirmationCustomData `json:"customData,omitempty"`
}

type NotifyEventConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (NotifyEventConfirmation) ActionName() string { return "NotifyEvent" }

func (NotifyEventConfirmation) Version() protocol.Version { return protocol.OCPP201 }

func (NotifyEventConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (NotifyEventConfirmation) SchemaName() string { return "NotifyEventResponse.json" }

func (message NotifyEventConfirmation) Validate() error {
	return validation.Validate("NotifyEventResponse.json", schemaNotifyEventConfirmation, message)
}

func (NotifyEventConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("NotifyEventResponse.json", schemaNotifyEventConfirmation, data)
}
