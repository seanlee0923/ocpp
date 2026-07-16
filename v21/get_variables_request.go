// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = GetVariablesRequest{}

var schemaGetVariablesRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"getVariableData": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"attributeType": &validation.Schema{Type: "string", Enum: []string{"Actual", "Target", "MinSet", "MaxSet"}}, "component": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"evse": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"id": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "connectorId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id"}}, "name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"name"}}, "variable": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"name"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"component", "variable"}}, MinItems: 1, HasMinItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"getVariableData"}}

type GetVariablesRequest struct {
	GetVariableData []GetVariablesRequestGetVariableData `json:"getVariableData"`
	CustomData      *GetVariablesRequestCustomData       `json:"customData,omitempty"`
}

type GetVariablesRequestGetVariableData struct {
	AttributeType *GetVariablesRequestAttributeEnum `json:"attributeType,omitempty"`
	Component     GetVariablesRequestComponent      `json:"component"`
	Variable      GetVariablesRequestVariable       `json:"variable"`
	CustomData    *GetVariablesRequestCustomData    `json:"customData,omitempty"`
}

type GetVariablesRequestVariable struct {
	Name       string                         `json:"name"`
	Instance   *string                        `json:"instance,omitempty"`
	CustomData *GetVariablesRequestCustomData `json:"customData,omitempty"`
}

type GetVariablesRequestComponent struct {
	EVSE       *GetVariablesRequestEVSE       `json:"evse,omitempty"`
	Name       string                         `json:"name"`
	Instance   *string                        `json:"instance,omitempty"`
	CustomData *GetVariablesRequestCustomData `json:"customData,omitempty"`
}

type GetVariablesRequestEVSE struct {
	ID          int                            `json:"id"`
	ConnectorID *int                           `json:"connectorId,omitempty"`
	CustomData  *GetVariablesRequestCustomData `json:"customData,omitempty"`
}

type GetVariablesRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

type GetVariablesRequestAttributeEnum string

const (
	GetVariablesRequestAttributeEnumActual GetVariablesRequestAttributeEnum = "Actual"
	GetVariablesRequestAttributeEnumTarget GetVariablesRequestAttributeEnum = "Target"
	GetVariablesRequestAttributeEnumMinSet GetVariablesRequestAttributeEnum = "MinSet"
	GetVariablesRequestAttributeEnumMaxSet GetVariablesRequestAttributeEnum = "MaxSet"
)

func (GetVariablesRequest) ActionName() string { return "GetVariables" }

func (GetVariablesRequest) Version() protocol.Version { return protocol.OCPP21 }

func (GetVariablesRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (GetVariablesRequest) SchemaName() string { return "GetVariablesRequest.json" }

func (message GetVariablesRequest) Validate() error {
	return validation.Validate("GetVariablesRequest.json", schemaGetVariablesRequest, message)
}

func (GetVariablesRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetVariablesRequest.json", schemaGetVariablesRequest, data)
}
