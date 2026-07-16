// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = SetVariablesConfirmation{}

var schemaSetVariablesConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "setVariableResult": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "attributeType": &validation.Schema{Type: "string", Enum: []string{"Actual", "Target", "MinSet", "MaxSet"}}, "attributeStatus": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "UnknownComponent", "UnknownVariable", "NotSupportedAttributeType", "RebootRequired"}}, "attributeStatusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}, "component": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "evse": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "id": &validation.Schema{Type: "integer", AllowAdditional: true}, "connectorId": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"id"}}, "name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}}, Required: []string{"name"}}, "variable": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}}, Required: []string{"name"}}}, Required: []string{"attributeStatus", "component", "variable"}}, MinItems: 1, HasMinItems: true}}, Required: []string{"setVariableResult"}}

type SetVariablesConfirmation struct {
	CustomData        *SetVariablesConfirmationCustomData         `json:"customData,omitempty"`
	SetVariableResult []SetVariablesConfirmationSetVariableResult `json:"setVariableResult"`
}

type SetVariablesConfirmationSetVariableResult struct {
	CustomData          *SetVariablesConfirmationCustomData           `json:"customData,omitempty"`
	AttributeType       *SetVariablesConfirmationAttributeEnum        `json:"attributeType,omitempty"`
	AttributeStatus     SetVariablesConfirmationSetVariableStatusEnum `json:"attributeStatus"`
	AttributeStatusInfo *SetVariablesConfirmationStatusInfo           `json:"attributeStatusInfo,omitempty"`
	Component           SetVariablesConfirmationComponent             `json:"component"`
	Variable            SetVariablesConfirmationVariable              `json:"variable"`
}

type SetVariablesConfirmationVariable struct {
	CustomData *SetVariablesConfirmationCustomData `json:"customData,omitempty"`
	Name       string                              `json:"name"`
	Instance   *string                             `json:"instance,omitempty"`
}

type SetVariablesConfirmationComponent struct {
	CustomData *SetVariablesConfirmationCustomData `json:"customData,omitempty"`
	EVSE       *SetVariablesConfirmationEVSE       `json:"evse,omitempty"`
	Name       string                              `json:"name"`
	Instance   *string                             `json:"instance,omitempty"`
}

type SetVariablesConfirmationEVSE struct {
	CustomData  *SetVariablesConfirmationCustomData `json:"customData,omitempty"`
	ID          int                                 `json:"id"`
	ConnectorID *int                                `json:"connectorId,omitempty"`
}

type SetVariablesConfirmationStatusInfo struct {
	CustomData     *SetVariablesConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                              `json:"reasonCode"`
	AdditionalInfo *string                             `json:"additionalInfo,omitempty"`
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

type SetVariablesConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (SetVariablesConfirmation) ActionName() string { return "SetVariables" }

func (SetVariablesConfirmation) Version() protocol.Version { return protocol.OCPP201 }

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
