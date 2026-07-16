// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = SetVariablesRequest{}

var schemaSetVariablesRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"setVariableData": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"attributeType": &validation.Schema{Type: "string", Enum: []string{"Actual", "Target", "MinSet", "MaxSet"}}, "attributeValue": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 2500, HasMaxLength: true}, "component": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"evse": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"id": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "connectorId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id"}}, "name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"name"}}, "variable": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"name"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"attributeValue", "component", "variable"}}, MinItems: 1, HasMinItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"setVariableData"}}

type SetVariablesRequest struct {
	SetVariableData []SetVariablesRequestSetVariableData `json:"setVariableData"`
	CustomData      *SetVariablesRequestCustomData       `json:"customData,omitempty"`
}

type SetVariablesRequestSetVariableData struct {
	AttributeType  *SetVariablesRequestAttributeEnum `json:"attributeType,omitempty"`
	AttributeValue string                            `json:"attributeValue"`
	Component      SetVariablesRequestComponent      `json:"component"`
	Variable       SetVariablesRequestVariable       `json:"variable"`
	CustomData     *SetVariablesRequestCustomData    `json:"customData,omitempty"`
}

type SetVariablesRequestVariable struct {
	Name       string                         `json:"name"`
	Instance   *string                        `json:"instance,omitempty"`
	CustomData *SetVariablesRequestCustomData `json:"customData,omitempty"`
}

type SetVariablesRequestComponent struct {
	EVSE       *SetVariablesRequestEVSE       `json:"evse,omitempty"`
	Name       string                         `json:"name"`
	Instance   *string                        `json:"instance,omitempty"`
	CustomData *SetVariablesRequestCustomData `json:"customData,omitempty"`
}

type SetVariablesRequestEVSE struct {
	ID          int                            `json:"id"`
	ConnectorID *int                           `json:"connectorId,omitempty"`
	CustomData  *SetVariablesRequestCustomData `json:"customData,omitempty"`
}

type SetVariablesRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value SetVariablesRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *SetVariablesRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

type SetVariablesRequestAttributeEnum string

const (
	SetVariablesRequestAttributeEnumActual SetVariablesRequestAttributeEnum = "Actual"
	SetVariablesRequestAttributeEnumTarget SetVariablesRequestAttributeEnum = "Target"
	SetVariablesRequestAttributeEnumMinSet SetVariablesRequestAttributeEnum = "MinSet"
	SetVariablesRequestAttributeEnumMaxSet SetVariablesRequestAttributeEnum = "MaxSet"
)

func (SetVariablesRequest) ActionName() string { return "SetVariables" }

func (SetVariablesRequest) Version() protocol.Version { return protocol.OCPP21 }

func (SetVariablesRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (SetVariablesRequest) SchemaName() string { return "SetVariablesRequest.json" }

func (message SetVariablesRequest) Validate() error {
	return validation.Validate("SetVariablesRequest.json", schemaSetVariablesRequest, message)
}

func (SetVariablesRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("SetVariablesRequest.json", schemaSetVariablesRequest, data)
}
