// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = GetConfigurationConfirmation{}

var schemaGetConfigurationConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"configurationKey": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"key": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "readonly": &validation.Schema{Type: "boolean", AllowAdditional: true}, "value": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 500, HasMaxLength: true}}, Required: []string{"key", "readonly"}}}, "unknownKey": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}}}}

type GetConfigurationConfirmation struct {
	ConfigurationKey []GetConfigurationConfirmationConfigurationKeyItem `json:"configurationKey,omitempty"`
	UnknownKey       []string                                           `json:"unknownKey,omitempty"`
}

type GetConfigurationConfirmationConfigurationKeyItem struct {
	Key      string  `json:"key"`
	Readonly bool    `json:"readonly"`
	Value    *string `json:"value,omitempty"`
}

func (GetConfigurationConfirmation) ActionName() string { return "GetConfiguration" }

func (GetConfigurationConfirmation) Version() protocol.Version { return protocol.OCPP16 }

func (GetConfigurationConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (GetConfigurationConfirmation) SchemaName() string { return "GetConfigurationResponse.json" }

func (message GetConfigurationConfirmation) Validate() error {
	return validation.Validate("GetConfigurationResponse.json", schemaGetConfigurationConfirmation, message)
}

func (GetConfigurationConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetConfigurationResponse.json", schemaGetConfigurationConfirmation, data)
}
