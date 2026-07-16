// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = GetMonitoringReportRequest{}

var schemaGetMonitoringReportRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"componentVariable": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"component": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"evse": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"id": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "connectorId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id"}}, "name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"name"}}, "variable": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"name"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"component"}}, MinItems: 1, HasMinItems: true}, "requestId": &validation.Schema{Type: "integer", AllowAdditional: true}, "monitoringCriteria": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", Enum: []string{"ThresholdMonitoring", "DeltaMonitoring", "PeriodicMonitoring"}}, MinItems: 1, HasMinItems: true, MaxItems: 3, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"requestId"}}

type GetMonitoringReportRequest struct {
	ComponentVariable  []GetMonitoringReportRequestComponentVariable       `json:"componentVariable,omitempty"`
	RequestID          int                                                 `json:"requestId"`
	MonitoringCriteria []GetMonitoringReportRequestMonitoringCriterionEnum `json:"monitoringCriteria,omitempty"`
	CustomData         *GetMonitoringReportRequestCustomData               `json:"customData,omitempty"`
}

type GetMonitoringReportRequestMonitoringCriterionEnum string

const (
	GetMonitoringReportRequestMonitoringCriterionEnumThresholdMonitoring GetMonitoringReportRequestMonitoringCriterionEnum = "ThresholdMonitoring"
	GetMonitoringReportRequestMonitoringCriterionEnumDeltaMonitoring     GetMonitoringReportRequestMonitoringCriterionEnum = "DeltaMonitoring"
	GetMonitoringReportRequestMonitoringCriterionEnumPeriodicMonitoring  GetMonitoringReportRequestMonitoringCriterionEnum = "PeriodicMonitoring"
)

type GetMonitoringReportRequestComponentVariable struct {
	Component  GetMonitoringReportRequestComponent   `json:"component"`
	Variable   *GetMonitoringReportRequestVariable   `json:"variable,omitempty"`
	CustomData *GetMonitoringReportRequestCustomData `json:"customData,omitempty"`
}

type GetMonitoringReportRequestVariable struct {
	Name       string                                `json:"name"`
	Instance   *string                               `json:"instance,omitempty"`
	CustomData *GetMonitoringReportRequestCustomData `json:"customData,omitempty"`
}

type GetMonitoringReportRequestComponent struct {
	EVSE       *GetMonitoringReportRequestEVSE       `json:"evse,omitempty"`
	Name       string                                `json:"name"`
	Instance   *string                               `json:"instance,omitempty"`
	CustomData *GetMonitoringReportRequestCustomData `json:"customData,omitempty"`
}

type GetMonitoringReportRequestEVSE struct {
	ID          int                                   `json:"id"`
	ConnectorID *int                                  `json:"connectorId,omitempty"`
	CustomData  *GetMonitoringReportRequestCustomData `json:"customData,omitempty"`
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

func (GetMonitoringReportRequest) Version() protocol.Version { return protocol.OCPP21 }

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
