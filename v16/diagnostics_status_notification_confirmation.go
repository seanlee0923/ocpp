// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = DiagnosticsStatusNotificationConfirmation{}

var schemaDiagnosticsStatusNotificationConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{}}

type DiagnosticsStatusNotificationConfirmation struct {
}

func (DiagnosticsStatusNotificationConfirmation) ActionName() string {
	return "DiagnosticsStatusNotification"
}

func (DiagnosticsStatusNotificationConfirmation) Version() protocol.Version { return protocol.OCPP16 }

func (DiagnosticsStatusNotificationConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (DiagnosticsStatusNotificationConfirmation) SchemaName() string {
	return "DiagnosticsStatusNotificationResponse.json"
}

func (message DiagnosticsStatusNotificationConfirmation) Validate() error {
	return validation.Validate("DiagnosticsStatusNotificationResponse.json", schemaDiagnosticsStatusNotificationConfirmation, message)
}

func (DiagnosticsStatusNotificationConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("DiagnosticsStatusNotificationResponse.json", schemaDiagnosticsStatusNotificationConfirmation, data)
}
