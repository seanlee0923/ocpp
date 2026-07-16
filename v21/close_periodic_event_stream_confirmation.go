// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = ClosePeriodicEventStreamConfirmation{}

var schemaClosePeriodicEventStreamConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type ClosePeriodicEventStreamConfirmation struct {
	CustomData *ClosePeriodicEventStreamConfirmationCustomData `json:"customData,omitempty"`
}

type ClosePeriodicEventStreamConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (ClosePeriodicEventStreamConfirmation) ActionName() string { return "ClosePeriodicEventStream" }

func (ClosePeriodicEventStreamConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (ClosePeriodicEventStreamConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (ClosePeriodicEventStreamConfirmation) SchemaName() string {
	return "ClosePeriodicEventStreamResponse.json"
}

func (message ClosePeriodicEventStreamConfirmation) Validate() error {
	return validation.Validate("ClosePeriodicEventStreamResponse.json", schemaClosePeriodicEventStreamConfirmation, message)
}

func (ClosePeriodicEventStreamConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ClosePeriodicEventStreamResponse.json", schemaClosePeriodicEventStreamConfirmation, data)
}
