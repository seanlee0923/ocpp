// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = SecurityEventNotificationRequest{}

var schemaSecurityEventNotificationRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "timestamp": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "techInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"type", "timestamp"}}

type SecurityEventNotificationRequest struct {
	Type       string                                      `json:"type"`
	Timestamp  string                                      `json:"timestamp"`
	TechInfo   *string                                     `json:"techInfo,omitempty"`
	CustomData *SecurityEventNotificationRequestCustomData `json:"customData,omitempty"`
}

type SecurityEventNotificationRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (SecurityEventNotificationRequest) ActionName() string { return "SecurityEventNotification" }

func (SecurityEventNotificationRequest) Version() protocol.Version { return protocol.OCPP21 }

func (SecurityEventNotificationRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (SecurityEventNotificationRequest) SchemaName() string {
	return "SecurityEventNotificationRequest.json"
}

func (message SecurityEventNotificationRequest) Validate() error {
	return validation.Validate("SecurityEventNotificationRequest.json", schemaSecurityEventNotificationRequest, message)
}

func (SecurityEventNotificationRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("SecurityEventNotificationRequest.json", schemaSecurityEventNotificationRequest, data)
}
