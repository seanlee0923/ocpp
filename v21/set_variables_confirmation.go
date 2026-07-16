// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = SetVariablesConfirmation{}

var schemaSetVariablesConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"setVariableResult": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"attributeType": &validation.Schema{Type: "string", Enum: []string{"Actual", "Target", "MinSet", "MaxSet"}}, "attributeStatus": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "UnknownComponent", "UnknownVariable", "NotSupportedAttributeType", "RebootRequired"}}, "attributeStatusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "component": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"evse": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"id": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "connectorId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id"}}, "name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"name"}}, "variable": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"name"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"attributeStatus", "component", "variable"}}, MinItems: 1, HasMinItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"setVariableResult"}}

type SetVariablesConfirmation struct {
	SetVariableResult []SetVariablesConfirmationSetVariableResult `json:"setVariableResult"`
	CustomData        *SetVariablesConfirmationCustomData         `json:"customData,omitempty"`
}

type SetVariablesConfirmationSetVariableResult struct {
	AttributeType       *SetVariablesConfirmationAttributeEnum        `json:"attributeType,omitempty"`
	AttributeStatus     SetVariablesConfirmationSetVariableStatusEnum `json:"attributeStatus"`
	AttributeStatusInfo *SetVariablesConfirmationStatusInfo           `json:"attributeStatusInfo,omitempty"`
	Component           SetVariablesConfirmationComponent             `json:"component"`
	Variable            SetVariablesConfirmationVariable              `json:"variable"`
	CustomData          *SetVariablesConfirmationCustomData           `json:"customData,omitempty"`
}

type SetVariablesConfirmationVariable struct {
	Name       string                              `json:"name"`
	Instance   *string                             `json:"instance,omitempty"`
	CustomData *SetVariablesConfirmationCustomData `json:"customData,omitempty"`
}

type SetVariablesConfirmationComponent struct {
	EVSE       *SetVariablesConfirmationEVSE       `json:"evse,omitempty"`
	Name       string                              `json:"name"`
	Instance   *string                             `json:"instance,omitempty"`
	CustomData *SetVariablesConfirmationCustomData `json:"customData,omitempty"`
}

type SetVariablesConfirmationEVSE struct {
	ID          int                                 `json:"id"`
	ConnectorID *int                                `json:"connectorId,omitempty"`
	CustomData  *SetVariablesConfirmationCustomData `json:"customData,omitempty"`
}

type SetVariablesConfirmationStatusInfo struct {
	ReasonCode     string                              `json:"reasonCode"`
	AdditionalInfo *string                             `json:"additionalInfo,omitempty"`
	CustomData     *SetVariablesConfirmationCustomData `json:"customData,omitempty"`
}

type SetVariablesConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value SetVariablesConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *SetVariablesConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

type SetVariablesConfirmationSetVariableStatusEnum string

const (
	SetVariablesConfirmationSetVariableStatusEnumAccepted                  SetVariablesConfirmationSetVariableStatusEnum = "Accepted"
	SetVariablesConfirmationSetVariableStatusEnumRejected                  SetVariablesConfirmationSetVariableStatusEnum = "Rejected"
	SetVariablesConfirmationSetVariableStatusEnumUnknownComponent          SetVariablesConfirmationSetVariableStatusEnum = "UnknownComponent"
	SetVariablesConfirmationSetVariableStatusEnumUnknownVariable           SetVariablesConfirmationSetVariableStatusEnum = "UnknownVariable"
	SetVariablesConfirmationSetVariableStatusEnumNotSupportedAttributeType SetVariablesConfirmationSetVariableStatusEnum = "NotSupportedAttributeType"
	SetVariablesConfirmationSetVariableStatusEnumRebootRequired            SetVariablesConfirmationSetVariableStatusEnum = "RebootRequired"
)

type SetVariablesConfirmationAttributeEnum string

const (
	SetVariablesConfirmationAttributeEnumActual SetVariablesConfirmationAttributeEnum = "Actual"
	SetVariablesConfirmationAttributeEnumTarget SetVariablesConfirmationAttributeEnum = "Target"
	SetVariablesConfirmationAttributeEnumMinSet SetVariablesConfirmationAttributeEnum = "MinSet"
	SetVariablesConfirmationAttributeEnumMaxSet SetVariablesConfirmationAttributeEnum = "MaxSet"
)

func (SetVariablesConfirmation) ActionName() string { return "SetVariables" }

func (SetVariablesConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (SetVariablesConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (SetVariablesConfirmation) SchemaName() string { return "SetVariablesResponse.json" }

func (message SetVariablesConfirmation) Validate() error {
	return validation.Validate("SetVariablesResponse.json", schemaSetVariablesConfirmation, message)
}

func (SetVariablesConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("SetVariablesResponse.json", schemaSetVariablesConfirmation, data)
}
