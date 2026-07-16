// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = SetVariablesRequest{}

var schemaSetVariablesRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "setVariableData": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "attributeType": &validation.Schema{Type: "string", Enum: []string{"Actual", "Target", "MinSet", "MaxSet"}}, "attributeValue": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1000, HasMaxLength: true}, "component": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "evse": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "id": &validation.Schema{Type: "integer", AllowAdditional: true}, "connectorId": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"id"}}, "name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}}, Required: []string{"name"}}, "variable": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}}, Required: []string{"name"}}}, Required: []string{"attributeValue", "component", "variable"}}, MinItems: 1, HasMinItems: true}}, Required: []string{"setVariableData"}}

type SetVariablesRequest struct {
	CustomData      *SetVariablesRequestCustomData       `json:"customData,omitempty"`
	SetVariableData []SetVariablesRequestSetVariableData `json:"setVariableData"`
}

type SetVariablesRequestSetVariableData struct {
	CustomData     *SetVariablesRequestCustomData    `json:"customData,omitempty"`
	AttributeType  *SetVariablesRequestAttributeEnum `json:"attributeType,omitempty"`
	AttributeValue string                            `json:"attributeValue"`
	Component      SetVariablesRequestComponent      `json:"component"`
	Variable       SetVariablesRequestVariable       `json:"variable"`
}

type SetVariablesRequestVariable struct {
	CustomData *SetVariablesRequestCustomData `json:"customData,omitempty"`
	Name       string                         `json:"name"`
	Instance   *string                        `json:"instance,omitempty"`
}

type SetVariablesRequestComponent struct {
	CustomData *SetVariablesRequestCustomData `json:"customData,omitempty"`
	EVSE       *SetVariablesRequestEVSE       `json:"evse,omitempty"`
	Name       string                         `json:"name"`
	Instance   *string                        `json:"instance,omitempty"`
}

type SetVariablesRequestEVSE struct {
	CustomData  *SetVariablesRequestCustomData `json:"customData,omitempty"`
	ID          int                            `json:"id"`
	ConnectorID *int                           `json:"connectorId,omitempty"`
}

type SetVariablesRequestAttributeEnum string

const (
	SetVariablesRequestAttributeEnumActual SetVariablesRequestAttributeEnum = "Actual"
	SetVariablesRequestAttributeEnumTarget SetVariablesRequestAttributeEnum = "Target"
	SetVariablesRequestAttributeEnumMinSet SetVariablesRequestAttributeEnum = "MinSet"
	SetVariablesRequestAttributeEnumMaxSet SetVariablesRequestAttributeEnum = "MaxSet"
)

type SetVariablesRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (SetVariablesRequest) ActionName() string { return "SetVariables" }

func (SetVariablesRequest) Version() protocol.Version { return protocol.OCPP201 }

func (SetVariablesRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (SetVariablesRequest) SchemaName() string { return "SetVariablesRequest.json" }

func (message SetVariablesRequest) Validate() error {
	return validation.Validate("SetVariablesRequest.json", schemaSetVariablesRequest, message)
}

func (SetVariablesRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("SetVariablesRequest.json", schemaSetVariablesRequest, data)
}
