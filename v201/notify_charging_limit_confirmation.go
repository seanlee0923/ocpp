// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = NotifyChargingLimitConfirmation{}

var schemaNotifyChargingLimitConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type NotifyChargingLimitConfirmation struct {
	CustomData *NotifyChargingLimitConfirmationCustomData `json:"customData,omitempty"`
}

type NotifyChargingLimitConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (NotifyChargingLimitConfirmation) ActionName() string { return "NotifyChargingLimit" }

func (NotifyChargingLimitConfirmation) Version() protocol.Version { return protocol.OCPP201 }

func (NotifyChargingLimitConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (NotifyChargingLimitConfirmation) SchemaName() string { return "NotifyChargingLimitResponse.json" }

func (message NotifyChargingLimitConfirmation) Validate() error {
	return validation.Validate("NotifyChargingLimitResponse.json", schemaNotifyChargingLimitConfirmation, message)
}

func (NotifyChargingLimitConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("NotifyChargingLimitResponse.json", schemaNotifyChargingLimitConfirmation, data)
}
