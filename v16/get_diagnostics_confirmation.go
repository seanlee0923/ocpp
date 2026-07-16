// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = GetDiagnosticsConfirmation{}

var schemaGetDiagnosticsConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"fileName": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}}

type GetDiagnosticsConfirmation struct {
	FileName *string `json:"fileName,omitempty"`
}

func (GetDiagnosticsConfirmation) ActionName() string { return "GetDiagnostics" }

func (GetDiagnosticsConfirmation) Version() protocol.Version { return protocol.OCPP16 }

func (GetDiagnosticsConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (GetDiagnosticsConfirmation) SchemaName() string { return "GetDiagnosticsResponse.json" }

func (message GetDiagnosticsConfirmation) Validate() error {
	return validation.Validate("GetDiagnosticsResponse.json", schemaGetDiagnosticsConfirmation, message)
}

func (GetDiagnosticsConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetDiagnosticsResponse.json", schemaGetDiagnosticsConfirmation, data)
}
