// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = GetMonitoringReportRequest{}

var schemaGetMonitoringReportRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "componentVariable": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "component": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "evse": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "id": &validation.Schema{Type: "integer", AllowAdditional: true}, "connectorId": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"id"}}, "name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}}, Required: []string{"name"}}, "variable": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}}, Required: []string{"name"}}}, Required: []string{"component"}}, MinItems: 1, HasMinItems: true}, "requestId": &validation.Schema{Type: "integer", AllowAdditional: true}, "monitoringCriteria": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", Enum: []string{"ThresholdMonitoring", "DeltaMonitoring", "PeriodicMonitoring"}}, MinItems: 1, HasMinItems: true, MaxItems: 3, HasMaxItems: true}}, Required: []string{"requestId"}}

type GetMonitoringReportRequest struct {
	CustomData         *GetMonitoringReportRequestCustomData               `json:"customData,omitempty"`
	ComponentVariable  []GetMonitoringReportRequestComponentVariable       `json:"componentVariable,omitempty"`
	RequestID          int                                                 `json:"requestId"`
	MonitoringCriteria []GetMonitoringReportRequestMonitoringCriterionEnum `json:"monitoringCriteria,omitempty"`
}

type GetMonitoringReportRequestMonitoringCriterionEnum string

const (
	GetMonitoringReportRequestMonitoringCriterionEnumThresholdMonitoring GetMonitoringReportRequestMonitoringCriterionEnum = "ThresholdMonitoring"
	GetMonitoringReportRequestMonitoringCriterionEnumDeltaMonitoring     GetMonitoringReportRequestMonitoringCriterionEnum = "DeltaMonitoring"
	GetMonitoringReportRequestMonitoringCriterionEnumPeriodicMonitoring  GetMonitoringReportRequestMonitoringCriterionEnum = "PeriodicMonitoring"
)

type GetMonitoringReportRequestComponentVariable struct {
	CustomData *GetMonitoringReportRequestCustomData `json:"customData,omitempty"`
	Component  GetMonitoringReportRequestComponent   `json:"component"`
	Variable   *GetMonitoringReportRequestVariable   `json:"variable,omitempty"`
}

type GetMonitoringReportRequestVariable struct {
	CustomData *GetMonitoringReportRequestCustomData `json:"customData,omitempty"`
	Name       string                                `json:"name"`
	Instance   *string                               `json:"instance,omitempty"`
}

type GetMonitoringReportRequestComponent struct {
	CustomData *GetMonitoringReportRequestCustomData `json:"customData,omitempty"`
	EVSE       *GetMonitoringReportRequestEVSE       `json:"evse,omitempty"`
	Name       string                                `json:"name"`
	Instance   *string                               `json:"instance,omitempty"`
}

type GetMonitoringReportRequestEVSE struct {
	CustomData  *GetMonitoringReportRequestCustomData `json:"customData,omitempty"`
	ID          int                                   `json:"id"`
	ConnectorID *int                                  `json:"connectorId,omitempty"`
}

type GetMonitoringReportRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value GetMonitoringReportRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *GetMonitoringReportRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (GetMonitoringReportRequest) ActionName() string { return "GetMonitoringReport" }

func (GetMonitoringReportRequest) Version() protocol.Version { return protocol.OCPP201 }

func (GetMonitoringReportRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (GetMonitoringReportRequest) SchemaName() string { return "GetMonitoringReportRequest.json" }

func (message GetMonitoringReportRequest) Validate() error {
	return validation.Validate("GetMonitoringReportRequest.json", schemaGetMonitoringReportRequest, message)
}

func (GetMonitoringReportRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetMonitoringReportRequest.json", schemaGetMonitoringReportRequest, data)
}
