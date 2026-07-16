// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = GetVariablesRequest{}

var schemaGetVariablesRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "getVariableData": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "attributeType": &validation.Schema{Type: "string", Enum: []string{"Actual", "Target", "MinSet", "MaxSet"}}, "component": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "evse": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "id": &validation.Schema{Type: "integer", AllowAdditional: true}, "connectorId": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"id"}}, "name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}}, Required: []string{"name"}}, "variable": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}}, Required: []string{"name"}}}, Required: []string{"component", "variable"}}, MinItems: 1, HasMinItems: true}}, Required: []string{"getVariableData"}}

type GetVariablesRequest struct {
	CustomData      *GetVariablesRequestCustomData       `json:"customData,omitempty"`
	GetVariableData []GetVariablesRequestGetVariableData `json:"getVariableData"`
}

type GetVariablesRequestGetVariableData struct {
	CustomData    *GetVariablesRequestCustomData    `json:"customData,omitempty"`
	AttributeType *GetVariablesRequestAttributeEnum `json:"attributeType,omitempty"`
	Component     GetVariablesRequestComponent      `json:"component"`
	Variable      GetVariablesRequestVariable       `json:"variable"`
}

type GetVariablesRequestVariable struct {
	CustomData *GetVariablesRequestCustomData `json:"customData,omitempty"`
	Name       string                         `json:"name"`
	Instance   *string                        `json:"instance,omitempty"`
}

type GetVariablesRequestComponent struct {
	CustomData *GetVariablesRequestCustomData `json:"customData,omitempty"`
	EVSE       *GetVariablesRequestEVSE       `json:"evse,omitempty"`
	Name       string                         `json:"name"`
	Instance   *string                        `json:"instance,omitempty"`
}

type GetVariablesRequestEVSE struct {
	CustomData  *GetVariablesRequestCustomData `json:"customData,omitempty"`
	ID          int                            `json:"id"`
	ConnectorID *int                           `json:"connectorId,omitempty"`
}

type GetVariablesRequestAttributeEnum string

const (
	GetVariablesRequestAttributeEnumActual GetVariablesRequestAttributeEnum = "Actual"
	GetVariablesRequestAttributeEnumTarget GetVariablesRequestAttributeEnum = "Target"
	GetVariablesRequestAttributeEnumMinSet GetVariablesRequestAttributeEnum = "MinSet"
	GetVariablesRequestAttributeEnumMaxSet GetVariablesRequestAttributeEnum = "MaxSet"
)

type GetVariablesRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value GetVariablesRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *GetVariablesRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (GetVariablesRequest) ActionName() string { return "GetVariables" }

func (GetVariablesRequest) Version() protocol.Version { return protocol.OCPP201 }

func (GetVariablesRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (GetVariablesRequest) SchemaName() string { return "GetVariablesRequest.json" }

func (message GetVariablesRequest) Validate() error {
	return validation.Validate("GetVariablesRequest.json", schemaGetVariablesRequest, message)
}

func (GetVariablesRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetVariablesRequest.json", schemaGetVariablesRequest, data)
}
