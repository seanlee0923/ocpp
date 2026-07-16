// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = GetConfigurationRequest{}

var schemaGetConfigurationRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"key": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}}}}

type GetConfigurationRequest struct {
	Key []string `json:"key,omitempty"`
}

func (GetConfigurationRequest) ActionName() string { return "GetConfiguration" }

func (GetConfigurationRequest) Version() protocol.Version { return protocol.OCPP16 }

func (GetConfigurationRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (GetConfigurationRequest) SchemaName() string { return "GetConfiguration.json" }

func (message GetConfigurationRequest) Validate() error {
	return validation.Validate("GetConfiguration.json", schemaGetConfigurationRequest, message)
}

func (GetConfigurationRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetConfiguration.json", schemaGetConfigurationRequest, data)
}
