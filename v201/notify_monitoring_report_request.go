// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = NotifyMonitoringReportRequest{}

var schemaNotifyMonitoringReportRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "monitor": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "component": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "evse": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "id": &validation.Schema{Type: "integer", AllowAdditional: true}, "connectorId": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"id"}}, "name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}}, Required: []string{"name"}}, "variable": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}}, Required: []string{"name"}}, "variableMonitoring": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "id": &validation.Schema{Type: "integer", AllowAdditional: true}, "transaction": &validation.Schema{Type: "boolean", AllowAdditional: true}, "value": &validation.Schema{Type: "number", AllowAdditional: true}, "type": &validation.Schema{Type: "string", Enum: []string{"UpperThreshold", "LowerThreshold", "Delta", "Periodic", "PeriodicClockAligned"}}, "severity": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"id", "transaction", "value", "type", "severity"}}, MinItems: 1, HasMinItems: true}}, Required: []string{"component", "variable", "variableMonitoring"}}, MinItems: 1, HasMinItems: true}, "requestId": &validation.Schema{Type: "integer", AllowAdditional: true}, "tbc": &validation.Schema{Type: "boolean", AllowAdditional: true}, "seqNo": &validation.Schema{Type: "integer", AllowAdditional: true}, "generatedAt": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}}, Required: []string{"requestId", "seqNo", "generatedAt"}}

type NotifyMonitoringReportRequest struct {
	CustomData  *NotifyMonitoringReportRequestCustomData      `json:"customData,omitempty"`
	Monitor     []NotifyMonitoringReportRequestMonitoringData `json:"monitor,omitempty"`
	RequestID   int                                           `json:"requestId"`
	Tbc         *bool                                         `json:"tbc,omitempty"`
	SeqNo       int                                           `json:"seqNo"`
	GeneratedAt string                                        `json:"generatedAt"`
}

type NotifyMonitoringReportRequestMonitoringData struct {
	CustomData         *NotifyMonitoringReportRequestCustomData          `json:"customData,omitempty"`
	Component          NotifyMonitoringReportRequestComponent            `json:"component"`
	Variable           NotifyMonitoringReportRequestVariable             `json:"variable"`
	VariableMonitoring []NotifyMonitoringReportRequestVariableMonitoring `json:"variableMonitoring"`
}

type NotifyMonitoringReportRequestVariableMonitoring struct {
	CustomData  *NotifyMonitoringReportRequestCustomData `json:"customData,omitempty"`
	ID          int                                      `json:"id"`
	Transaction bool                                     `json:"transaction"`
	Value       float64                                  `json:"value"`
	Type        NotifyMonitoringReportRequestMonitorEnum `json:"type"`
	Severity    int                                      `json:"severity"`
}

type NotifyMonitoringReportRequestMonitorEnum string

const (
	NotifyMonitoringReportRequestMonitorEnumUpperThreshold       NotifyMonitoringReportRequestMonitorEnum = "UpperThreshold"
	NotifyMonitoringReportRequestMonitorEnumLowerThreshold       NotifyMonitoringReportRequestMonitorEnum = "LowerThreshold"
	NotifyMonitoringReportRequestMonitorEnumDelta                NotifyMonitoringReportRequestMonitorEnum = "Delta"
	NotifyMonitoringReportRequestMonitorEnumPeriodic             NotifyMonitoringReportRequestMonitorEnum = "Periodic"
	NotifyMonitoringReportRequestMonitorEnumPeriodicClockAligned NotifyMonitoringReportRequestMonitorEnum = "PeriodicClockAligned"
)

type NotifyMonitoringReportRequestVariable struct {
	CustomData *NotifyMonitoringReportRequestCustomData `json:"customData,omitempty"`
	Name       string                                   `json:"name"`
	Instance   *string                                  `json:"instance,omitempty"`
}

type NotifyMonitoringReportRequestComponent struct {
	CustomData *NotifyMonitoringReportRequestCustomData `json:"customData,omitempty"`
	EVSE       *NotifyMonitoringReportRequestEVSE       `json:"evse,omitempty"`
	Name       string                                   `json:"name"`
	Instance   *string                                  `json:"instance,omitempty"`
}

type NotifyMonitoringReportRequestEVSE struct {
	CustomData  *NotifyMonitoringReportRequestCustomData `json:"customData,omitempty"`
	ID          int                                      `json:"id"`
	ConnectorID *int                                     `json:"connectorId,omitempty"`
}

type NotifyMonitoringReportRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (NotifyMonitoringReportRequest) ActionName() string { return "NotifyMonitoringReport" }

func (NotifyMonitoringReportRequest) Version() protocol.Version { return protocol.OCPP201 }

func (NotifyMonitoringReportRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (NotifyMonitoringReportRequest) SchemaName() string { return "NotifyMonitoringReportRequest.json" }

func (message NotifyMonitoringReportRequest) Validate() error {
	return validation.Validate("NotifyMonitoringReportRequest.json", schemaNotifyMonitoringReportRequest, message)
}

func (NotifyMonitoringReportRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("NotifyMonitoringReportRequest.json", schemaNotifyMonitoringReportRequest, data)
}
