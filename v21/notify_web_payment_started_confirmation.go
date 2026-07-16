// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = NotifyWebPaymentStartedConfirmation{}

var schemaNotifyWebPaymentStartedConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type NotifyWebPaymentStartedConfirmation struct {
	CustomData *NotifyWebPaymentStartedConfirmationCustomData `json:"customData,omitempty"`
}

type NotifyWebPaymentStartedConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (NotifyWebPaymentStartedConfirmation) ActionName() string { return "NotifyWebPaymentStarted" }

func (NotifyWebPaymentStartedConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (NotifyWebPaymentStartedConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (NotifyWebPaymentStartedConfirmation) SchemaName() string {
	return "NotifyWebPaymentStartedResponse.json"
}

func (message NotifyWebPaymentStartedConfirmation) Validate() error {
	return validation.Validate("NotifyWebPaymentStartedResponse.json", schemaNotifyWebPaymentStartedConfirmation, message)
}

func (NotifyWebPaymentStartedConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("NotifyWebPaymentStartedResponse.json", schemaNotifyWebPaymentStartedConfirmation, data)
}
