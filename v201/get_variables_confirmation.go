// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = GetVariablesConfirmation{}

var schemaGetVariablesConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "getVariableResult": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "attributeStatusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}, "attributeStatus": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "UnknownComponent", "UnknownVariable", "NotSupportedAttributeType"}}, "attributeType": &validation.Schema{Type: "string", Enum: []string{"Actual", "Target", "MinSet", "MaxSet"}}, "attributeValue": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 2500, HasMaxLength: true}, "component": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "evse": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "id": &validation.Schema{Type: "integer", AllowAdditional: true}, "connectorId": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"id"}}, "name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}}, Required: []string{"name"}}, "variable": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}}, Required: []string{"name"}}}, Required: []string{"attributeStatus", "component", "variable"}}, MinItems: 1, HasMinItems: true}}, Required: []string{"getVariableResult"}}

type GetVariablesConfirmation struct {
	CustomData        *GetVariablesConfirmationCustomData         `json:"customData,omitempty"`
	GetVariableResult []GetVariablesConfirmationGetVariableResult `json:"getVariableResult"`
}

type GetVariablesConfirmationGetVariableResult struct {
	CustomData          *GetVariablesConfirmationCustomData           `json:"customData,omitempty"`
	AttributeStatusInfo *GetVariablesConfirmationStatusInfo           `json:"attributeStatusInfo,omitempty"`
	AttributeStatus     GetVariablesConfirmationGetVariableStatusEnum `json:"attributeStatus"`
	AttributeType       *GetVariablesConfirmationAttributeEnum        `json:"attributeType,omitempty"`
	AttributeValue      *string                                       `json:"attributeValue,omitempty"`
	Component           GetVariablesConfirmationComponent             `json:"component"`
	Variable            GetVariablesConfirmationVariable              `json:"variable"`
}

type GetVariablesConfirmationVariable struct {
	CustomData *GetVariablesConfirmationCustomData `json:"customData,omitempty"`
	Name       string                              `json:"name"`
	Instance   *string                             `json:"instance,omitempty"`
}

type GetVariablesConfirmationComponent struct {
	CustomData *GetVariablesConfirmationCustomData `json:"customData,omitempty"`
	EVSE       *GetVariablesConfirmationEVSE       `json:"evse,omitempty"`
	Name       string                              `json:"name"`
	Instance   *string                             `json:"instance,omitempty"`
}

type GetVariablesConfirmationEVSE struct {
	CustomData  *GetVariablesConfirmationCustomData `json:"customData,omitempty"`
	ID          int                                 `json:"id"`
	ConnectorID *int                                `json:"connectorId,omitempty"`
}

type GetVariablesConfirmationAttributeEnum string

const (
	GetVariablesConfirmationAttributeEnumActual GetVariablesConfirmationAttributeEnum = "Actual"
	GetVariablesConfirmationAttributeEnumTarget GetVariablesConfirmationAttributeEnum = "Target"
	GetVariablesConfirmationAttributeEnumMinSet GetVariablesConfirmationAttributeEnum = "MinSet"
	GetVariablesConfirmationAttributeEnumMaxSet GetVariablesConfirmationAttributeEnum = "MaxSet"
)

type GetVariablesConfirmationGetVariableStatusEnum string

const (
	GetVariablesConfirmationGetVariableStatusEnumAccepted                  GetVariablesConfirmationGetVariableStatusEnum = "Accepted"
	GetVariablesConfirmationGetVariableStatusEnumRejected                  GetVariablesConfirmationGetVariableStatusEnum = "Rejected"
	GetVariablesConfirmationGetVariableStatusEnumUnknownComponent          GetVariablesConfirmationGetVariableStatusEnum = "UnknownComponent"
	GetVariablesConfirmationGetVariableStatusEnumUnknownVariable           GetVariablesConfirmationGetVariableStatusEnum = "UnknownVariable"
	GetVariablesConfirmationGetVariableStatusEnumNotSupportedAttributeType GetVariablesConfirmationGetVariableStatusEnum = "NotSupportedAttributeType"
)

type GetVariablesConfirmationStatusInfo struct {
	CustomData     *GetVariablesConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                              `json:"reasonCode"`
	AdditionalInfo *string                             `json:"additionalInfo,omitempty"`
}

type GetVariablesConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (GetVariablesConfirmation) ActionName() string { return "GetVariables" }

func (GetVariablesConfirmation) Version() protocol.Version { return protocol.OCPP201 }

func (GetVariablesConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (GetVariablesConfirmation) SchemaName() string { return "GetVariablesResponse.json" }

func (message GetVariablesConfirmation) Validate() error {
	return validation.Validate("GetVariablesResponse.json", schemaGetVariablesConfirmation, message)
}

func (GetVariablesConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetVariablesResponse.json", schemaGetVariablesConfirmation, data)
}
