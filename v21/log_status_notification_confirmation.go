// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = LogStatusNotificationConfirmation{}

var schemaLogStatusNotificationConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type LogStatusNotificationConfirmation struct {
	CustomData *LogStatusNotificationConfirmationCustomData `json:"customData,omitempty"`
}

type LogStatusNotificationConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (LogStatusNotificationConfirmation) ActionName() string { return "LogStatusNotification" }

func (LogStatusNotificationConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (LogStatusNotificationConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (LogStatusNotificationConfirmation) SchemaName() string {
	return "LogStatusNotificationResponse.json"
}

func (message LogStatusNotificationConfirmation) Validate() error {
	return validation.Validate("LogStatusNotificationResponse.json", schemaLogStatusNotificationConfirmation, message)
}

func (LogStatusNotificationConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("LogStatusNotificationResponse.json", schemaLogStatusNotificationConfirmation, data)
}
