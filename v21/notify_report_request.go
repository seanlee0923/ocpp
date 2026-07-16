// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = NotifyReportRequest{}

var schemaNotifyReportRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"requestId": &validation.Schema{Type: "integer", AllowAdditional: true}, "generatedAt": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "reportData": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"component": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"evse": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"id": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "connectorId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id"}}, "name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"name"}}, "variable": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"name"}}, "variableAttribute": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"type": &validation.Schema{Type: "string", Enum: []string{"Actual", "Target", "MinSet", "MaxSet"}}, "value": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 2500, HasMaxLength: true}, "mutability": &validation.Schema{Type: "string", Enum: []string{"ReadOnly", "WriteOnly", "ReadWrite"}}, "persistent": &validation.Schema{Type: "boolean", AllowAdditional: true}, "constant": &validation.Schema{Type: "boolean", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, MinItems: 1, HasMinItems: true, MaxItems: 4, HasMaxItems: true}, "variableCharacteristics": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"unit": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 16, HasMaxLength: true}, "dataType": &validation.Schema{Type: "string", Enum: []string{"string", "decimal", "integer", "dateTime", "boolean", "OptionList", "SequenceList", "MemberList"}}, "minLimit": &validation.Schema{Type: "number", AllowAdditional: true}, "maxLimit": &validation.Schema{Type: "number", AllowAdditional: true}, "maxElements": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 1, HasMinimum: true}, "valuesList": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1000, HasMaxLength: true}, "supportsMonitoring": &validation.Schema{Type: "boolean", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"dataType", "supportsMonitoring"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"component", "variable", "variableAttribute"}}, MinItems: 1, HasMinItems: true}, "tbc": &validation.Schema{Type: "boolean", AllowAdditional: true}, "seqNo": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"requestId", "generatedAt", "seqNo"}}

type NotifyReportRequest struct {
	RequestID   int                             `json:"requestId"`
	GeneratedAt string                          `json:"generatedAt"`
	ReportData  []NotifyReportRequestReportData `json:"reportData,omitempty"`
	Tbc         *bool                           `json:"tbc,omitempty"`
	SeqNo       int                             `json:"seqNo"`
	CustomData  *NotifyReportRequestCustomData  `json:"customData,omitempty"`
}

type NotifyReportRequestReportData struct {
	Component               NotifyReportRequestComponent                `json:"component"`
	Variable                NotifyReportRequestVariable                 `json:"variable"`
	VariableAttribute       []NotifyReportRequestVariableAttribute      `json:"variableAttribute"`
	VariableCharacteristics *NotifyReportRequestVariableCharacteristics `json:"variableCharacteristics,omitempty"`
	CustomData              *NotifyReportRequestCustomData              `json:"customData,omitempty"`
}

type NotifyReportRequestVariableCharacteristics struct {
	Unit               *string                        `json:"unit,omitempty"`
	DataType           NotifyReportRequestDataEnum    `json:"dataType"`
	MinLimit           *float64                       `json:"minLimit,omitempty"`
	MaxLimit           *float64                       `json:"maxLimit,omitempty"`
	MaxElements        *int                           `json:"maxElements,omitempty"`
	ValuesList         *string                        `json:"valuesList,omitempty"`
	SupportsMonitoring bool                           `json:"supportsMonitoring"`
	CustomData         *NotifyReportRequestCustomData `json:"customData,omitempty"`
}

type NotifyReportRequestDataEnum string

const (
	NotifyReportRequestDataEnumString       NotifyReportRequestDataEnum = "string"
	NotifyReportRequestDataEnumDecimal      NotifyReportRequestDataEnum = "decimal"
	NotifyReportRequestDataEnumInteger      NotifyReportRequestDataEnum = "integer"
	NotifyReportRequestDataEnumDateTime     NotifyReportRequestDataEnum = "dateTime"
	NotifyReportRequestDataEnumBoolean      NotifyReportRequestDataEnum = "boolean"
	NotifyReportRequestDataEnumOptionList   NotifyReportRequestDataEnum = "OptionList"
	NotifyReportRequestDataEnumSequenceList NotifyReportRequestDataEnum = "SequenceList"
	NotifyReportRequestDataEnumMemberList   NotifyReportRequestDataEnum = "MemberList"
)

type NotifyReportRequestVariableAttribute struct {
	Type       *NotifyReportRequestAttributeEnum  `json:"type,omitempty"`
	Value      *string                            `json:"value,omitempty"`
	Mutability *NotifyReportRequestMutabilityEnum `json:"mutability,omitempty"`
	Persistent *bool                              `json:"persistent,omitempty"`
	Constant   *bool                              `json:"constant,omitempty"`
	CustomData *NotifyReportRequestCustomData     `json:"customData,omitempty"`
}

type NotifyReportRequestMutabilityEnum string

const (
	NotifyReportRequestMutabilityEnumReadOnly  NotifyReportRequestMutabilityEnum = "ReadOnly"
	NotifyReportRequestMutabilityEnumWriteOnly NotifyReportRequestMutabilityEnum = "WriteOnly"
	NotifyReportRequestMutabilityEnumReadWrite NotifyReportRequestMutabilityEnum = "ReadWrite"
)

type NotifyReportRequestAttributeEnum string

const (
	NotifyReportRequestAttributeEnumActual NotifyReportRequestAttributeEnum = "Actual"
	NotifyReportRequestAttributeEnumTarget NotifyReportRequestAttributeEnum = "Target"
	NotifyReportRequestAttributeEnumMinSet NotifyReportRequestAttributeEnum = "MinSet"
	NotifyReportRequestAttributeEnumMaxSet NotifyReportRequestAttributeEnum = "MaxSet"
)

type NotifyReportRequestVariable struct {
	Name       string                         `json:"name"`
	Instance   *string                        `json:"instance,omitempty"`
	CustomData *NotifyReportRequestCustomData `json:"customData,omitempty"`
}

type NotifyReportRequestComponent struct {
	EVSE       *NotifyReportRequestEVSE       `json:"evse,omitempty"`
	Name       string                         `json:"name"`
	Instance   *string                        `json:"instance,omitempty"`
	CustomData *NotifyReportRequestCustomData `json:"customData,omitempty"`
}

type NotifyReportRequestEVSE struct {
	ID          int                            `json:"id"`
	ConnectorID *int                           `json:"connectorId,omitempty"`
	CustomData  *NotifyReportRequestCustomData `json:"customData,omitempty"`
}

type NotifyReportRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (NotifyReportRequest) ActionName() string { return "NotifyReport" }

func (NotifyReportRequest) Version() protocol.Version { return protocol.OCPP21 }

func (NotifyReportRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (NotifyReportRequest) SchemaName() string { return "NotifyReportRequest.json" }

func (message NotifyReportRequest) Validate() error {
	return validation.Validate("NotifyReportRequest.json", schemaNotifyReportRequest, message)
}

func (NotifyReportRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("NotifyReportRequest.json", schemaNotifyReportRequest, data)
}
