// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = GetVariablesConfirmation{}

var schemaGetVariablesConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"getVariableResult": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"attributeStatus": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "UnknownComponent", "UnknownVariable", "NotSupportedAttributeType"}}, "attributeStatusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "attributeType": &validation.Schema{Type: "string", Enum: []string{"Actual", "Target", "MinSet", "MaxSet"}}, "attributeValue": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 2500, HasMaxLength: true}, "component": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"evse": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"id": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "connectorId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id"}}, "name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"name"}}, "variable": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"name"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"attributeStatus", "component", "variable"}}, MinItems: 1, HasMinItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"getVariableResult"}}

type GetVariablesConfirmation struct {
	GetVariableResult []GetVariablesConfirmationGetVariableResult `json:"getVariableResult"`
	CustomData        *GetVariablesConfirmationCustomData         `json:"customData,omitempty"`
}

type GetVariablesConfirmationGetVariableResult struct {
	AttributeStatus     GetVariablesConfirmationGetVariableStatusEnum `json:"attributeStatus"`
	AttributeStatusInfo *GetVariablesConfirmationStatusInfo           `json:"attributeStatusInfo,omitempty"`
	AttributeType       *GetVariablesConfirmationAttributeEnum        `json:"attributeType,omitempty"`
	AttributeValue      *string                                       `json:"attributeValue,omitempty"`
	Component           GetVariablesConfirmationComponent             `json:"component"`
	Variable            GetVariablesConfirmationVariable              `json:"variable"`
	CustomData          *GetVariablesConfirmationCustomData           `json:"customData,omitempty"`
}

type GetVariablesConfirmationVariable struct {
	Name       string                              `json:"name"`
	Instance   *string                             `json:"instance,omitempty"`
	CustomData *GetVariablesConfirmationCustomData `json:"customData,omitempty"`
}

type GetVariablesConfirmationComponent struct {
	EVSE       *GetVariablesConfirmationEVSE       `json:"evse,omitempty"`
	Name       string                              `json:"name"`
	Instance   *string                             `json:"instance,omitempty"`
	CustomData *GetVariablesConfirmationCustomData `json:"customData,omitempty"`
}

type GetVariablesConfirmationEVSE struct {
	ID          int                                 `json:"id"`
	ConnectorID *int                                `json:"connectorId,omitempty"`
	CustomData  *GetVariablesConfirmationCustomData `json:"customData,omitempty"`
}

type GetVariablesConfirmationAttributeEnum string

const (
	GetVariablesConfirmationAttributeEnumActual GetVariablesConfirmationAttributeEnum = "Actual"
	GetVariablesConfirmationAttributeEnumTarget GetVariablesConfirmationAttributeEnum = "Target"
	GetVariablesConfirmationAttributeEnumMinSet GetVariablesConfirmationAttributeEnum = "MinSet"
	GetVariablesConfirmationAttributeEnumMaxSet GetVariablesConfirmationAttributeEnum = "MaxSet"
)

type GetVariablesConfirmationStatusInfo struct {
	ReasonCode     string                              `json:"reasonCode"`
	AdditionalInfo *string                             `json:"additionalInfo,omitempty"`
	CustomData     *GetVariablesConfirmationCustomData `json:"customData,omitempty"`
}

type GetVariablesConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

type GetVariablesConfirmationGetVariableStatusEnum string

const (
	GetVariablesConfirmationGetVariableStatusEnumAccepted                  GetVariablesConfirmationGetVariableStatusEnum = "Accepted"
	GetVariablesConfirmationGetVariableStatusEnumRejected                  GetVariablesConfirmationGetVariableStatusEnum = "Rejected"
	GetVariablesConfirmationGetVariableStatusEnumUnknownComponent          GetVariablesConfirmationGetVariableStatusEnum = "UnknownComponent"
	GetVariablesConfirmationGetVariableStatusEnumUnknownVariable           GetVariablesConfirmationGetVariableStatusEnum = "UnknownVariable"
	GetVariablesConfirmationGetVariableStatusEnumNotSupportedAttributeType GetVariablesConfirmationGetVariableStatusEnum = "NotSupportedAttributeType"
)

func (GetVariablesConfirmation) ActionName() string { return "GetVariables" }

func (GetVariablesConfirmation) Version() protocol.Version { return protocol.OCPP21 }

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
