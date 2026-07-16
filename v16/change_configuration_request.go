// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = ChangeConfigurationRequest{}

var schemaChangeConfigurationRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"key": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "value": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 500, HasMaxLength: true}}, Required: []string{"key", "value"}}

type ChangeConfigurationRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (ChangeConfigurationRequest) ActionName() string { return "ChangeConfiguration" }

func (ChangeConfigurationRequest) Version() protocol.Version { return protocol.OCPP16 }

func (ChangeConfigurationRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (ChangeConfigurationRequest) SchemaName() string { return "ChangeConfiguration.json" }

func (message ChangeConfigurationRequest) Validate() error {
	return validation.Validate("ChangeConfiguration.json", schemaChangeConfigurationRequest, message)
}

func (ChangeConfigurationRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ChangeConfiguration.json", schemaChangeConfigurationRequest, data)
}
