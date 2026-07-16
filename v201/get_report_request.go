// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = GetReportRequest{}

var schemaGetReportRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "componentVariable": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "component": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "evse": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "id": &validation.Schema{Type: "integer", AllowAdditional: true}, "connectorId": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"id"}}, "name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}}, Required: []string{"name"}}, "variable": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}}, Required: []string{"name"}}}, Required: []string{"component"}}, MinItems: 1, HasMinItems: true}, "requestId": &validation.Schema{Type: "integer", AllowAdditional: true}, "componentCriteria": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", Enum: []string{"Active", "Available", "Enabled", "Problem"}}, MinItems: 1, HasMinItems: true, MaxItems: 4, HasMaxItems: true}}, Required: []string{"requestId"}}

type GetReportRequest struct {
	CustomData        *GetReportRequestCustomData              `json:"customData,omitempty"`
	ComponentVariable []GetReportRequestComponentVariable      `json:"componentVariable,omitempty"`
	RequestID         int                                      `json:"requestId"`
	ComponentCriteria []GetReportRequestComponentCriterionEnum `json:"componentCriteria,omitempty"`
}

type GetReportRequestComponentCriterionEnum string

const (
	GetReportRequestComponentCriterionEnumActive    GetReportRequestComponentCriterionEnum = "Active"
	GetReportRequestComponentCriterionEnumAvailable GetReportRequestComponentCriterionEnum = "Available"
	GetReportRequestComponentCriterionEnumEnabled   GetReportRequestComponentCriterionEnum = "Enabled"
	GetReportRequestComponentCriterionEnumProblem   GetReportRequestComponentCriterionEnum = "Problem"
)

type GetReportRequestComponentVariable struct {
	CustomData *GetReportRequestCustomData `json:"customData,omitempty"`
	Component  GetReportRequestComponent   `json:"component"`
	Variable   *GetReportRequestVariable   `json:"variable,omitempty"`
}

type GetReportRequestVariable struct {
	CustomData *GetReportRequestCustomData `json:"customData,omitempty"`
	Name       string                      `json:"name"`
	Instance   *string                     `json:"instance,omitempty"`
}

type GetReportRequestComponent struct {
	CustomData *GetReportRequestCustomData `json:"customData,omitempty"`
	EVSE       *GetReportRequestEVSE       `json:"evse,omitempty"`
	Name       string                      `json:"name"`
	Instance   *string                     `json:"instance,omitempty"`
}

type GetReportRequestEVSE struct {
	CustomData  *GetReportRequestCustomData `json:"customData,omitempty"`
	ID          int                         `json:"id"`
	ConnectorID *int                        `json:"connectorId,omitempty"`
}

type GetReportRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (GetReportRequest) ActionName() string { return "GetReport" }

func (GetReportRequest) Version() protocol.Version { return protocol.OCPP201 }

func (GetReportRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (GetReportRequest) SchemaName() string { return "GetReportRequest.json" }

func (message GetReportRequest) Validate() error {
	return validation.Validate("GetReportRequest.json", schemaGetReportRequest, message)
}

func (GetReportRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetReportRequest.json", schemaGetReportRequest, data)
}
