// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = StatusNotificationConfirmation{}

var schemaStatusNotificationConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{}}

type StatusNotificationConfirmation struct {
}

func (StatusNotificationConfirmation) ActionName() string { return "StatusNotification" }

func (StatusNotificationConfirmation) Version() protocol.Version { return protocol.OCPP16 }

func (StatusNotificationConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (StatusNotificationConfirmation) SchemaName() string { return "StatusNotificationResponse.json" }

func (message StatusNotificationConfirmation) Validate() error {
	return validation.Validate("StatusNotificationResponse.json", schemaStatusNotificationConfirmation, message)
}

func (StatusNotificationConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("StatusNotificationResponse.json", schemaStatusNotificationConfirmation, data)
}
