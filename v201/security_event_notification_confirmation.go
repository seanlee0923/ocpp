// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = SecurityEventNotificationConfirmation{}

var schemaSecurityEventNotificationConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type SecurityEventNotificationConfirmation struct {
	CustomData *SecurityEventNotificationConfirmationCustomData `json:"customData,omitempty"`
}

type SecurityEventNotificationConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (SecurityEventNotificationConfirmation) ActionName() string { return "SecurityEventNotification" }

func (SecurityEventNotificationConfirmation) Version() protocol.Version { return protocol.OCPP201 }

func (SecurityEventNotificationConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (SecurityEventNotificationConfirmation) SchemaName() string {
	return "SecurityEventNotificationResponse.json"
}

func (message SecurityEventNotificationConfirmation) Validate() error {
	return validation.Validate("SecurityEventNotificationResponse.json", schemaSecurityEventNotificationConfirmation, message)
}

func (SecurityEventNotificationConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("SecurityEventNotificationResponse.json", schemaSecurityEventNotificationConfirmation, data)
}
