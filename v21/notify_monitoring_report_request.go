// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = NotifyMonitoringReportRequest{}

var schemaNotifyMonitoringReportRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"monitor": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"component": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"evse": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"id": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "connectorId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id"}}, "name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"name"}}, "variable": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"name"}}, "variableMonitoring": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"id": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "transaction": &validation.Schema{Type: "boolean", AllowAdditional: true}, "value": &validation.Schema{Type: "number", AllowAdditional: true}, "type": &validation.Schema{Type: "string", Enum: []string{"UpperThreshold", "LowerThreshold", "Delta", "Periodic", "PeriodicClockAligned", "TargetDelta", "TargetDeltaRelative"}}, "severity": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "eventNotificationType": &validation.Schema{Type: "string", Enum: []string{"HardWiredNotification", "HardWiredMonitor", "PreconfiguredMonitor", "CustomMonitor"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id", "transaction", "value", "type", "severity", "eventNotificationType"}}, MinItems: 1, HasMinItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"component", "variable", "variableMonitoring"}}, MinItems: 1, HasMinItems: true}, "requestId": &validation.Schema{Type: "integer", AllowAdditional: true}, "tbc": &validation.Schema{Type: "boolean", AllowAdditional: true}, "seqNo": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "generatedAt": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"requestId", "seqNo", "generatedAt"}}

type NotifyMonitoringReportRequest struct {
	Monitor     []NotifyMonitoringReportRequestMonitoringData `json:"monitor,omitempty"`
	RequestID   int                                           `json:"requestId"`
	Tbc         *bool                                         `json:"tbc,omitempty"`
	SeqNo       int                                           `json:"seqNo"`
	GeneratedAt string                                        `json:"generatedAt"`
	CustomData  *NotifyMonitoringReportRequestCustomData      `json:"customData,omitempty"`
}

type NotifyMonitoringReportRequestMonitoringData struct {
	Component          NotifyMonitoringReportRequestComponent            `json:"component"`
	Variable           NotifyMonitoringReportRequestVariable             `json:"variable"`
	VariableMonitoring []NotifyMonitoringReportRequestVariableMonitoring `json:"variableMonitoring"`
	CustomData         *NotifyMonitoringReportRequestCustomData          `json:"customData,omitempty"`
}

type NotifyMonitoringReportRequestVariableMonitoring struct {
	ID                    int                                                `json:"id"`
	Transaction           bool                                               `json:"transaction"`
	Value                 float64                                            `json:"value"`
	Type                  NotifyMonitoringReportRequestMonitorEnum           `json:"type"`
	Severity              int                                                `json:"severity"`
	EventNotificationType NotifyMonitoringReportRequestEventNotificationEnum `json:"eventNotificationType"`
	CustomData            *NotifyMonitoringReportRequestCustomData           `json:"customData,omitempty"`
}

type NotifyMonitoringReportRequestEventNotificationEnum string

const (
	NotifyMonitoringReportRequestEventNotificationEnumHardWiredNotification NotifyMonitoringReportRequestEventNotificationEnum = "HardWiredNotification"
	NotifyMonitoringReportRequestEventNotificationEnumHardWiredMonitor      NotifyMonitoringReportRequestEventNotificationEnum = "HardWiredMonitor"
	NotifyMonitoringReportRequestEventNotificationEnumPreconfiguredMonitor  NotifyMonitoringReportRequestEventNotificationEnum = "PreconfiguredMonitor"
	NotifyMonitoringReportRequestEventNotificationEnumCustomMonitor         NotifyMonitoringReportRequestEventNotificationEnum = "CustomMonitor"
)

type NotifyMonitoringReportRequestMonitorEnum string

const (
	NotifyMonitoringReportRequestMonitorEnumUpperThreshold       NotifyMonitoringReportRequestMonitorEnum = "UpperThreshold"
	NotifyMonitoringReportRequestMonitorEnumLowerThreshold       NotifyMonitoringReportRequestMonitorEnum = "LowerThreshold"
	NotifyMonitoringReportRequestMonitorEnumDelta                NotifyMonitoringReportRequestMonitorEnum = "Delta"
	NotifyMonitoringReportRequestMonitorEnumPeriodic             NotifyMonitoringReportRequestMonitorEnum = "Periodic"
	NotifyMonitoringReportRequestMonitorEnumPeriodicClockAligned NotifyMonitoringReportRequestMonitorEnum = "PeriodicClockAligned"
	NotifyMonitoringReportRequestMonitorEnumTargetDelta          NotifyMonitoringReportRequestMonitorEnum = "TargetDelta"
	NotifyMonitoringReportRequestMonitorEnumTargetDeltaRelative  NotifyMonitoringReportRequestMonitorEnum = "TargetDeltaRelative"
)

type NotifyMonitoringReportRequestVariable struct {
	Name       string                                   `json:"name"`
	Instance   *string                                  `json:"instance,omitempty"`
	CustomData *NotifyMonitoringReportRequestCustomData `json:"customData,omitempty"`
}

type NotifyMonitoringReportRequestComponent struct {
	EVSE       *NotifyMonitoringReportRequestEVSE       `json:"evse,omitempty"`
	Name       string                                   `json:"name"`
	Instance   *string                                  `json:"instance,omitempty"`
	CustomData *NotifyMonitoringReportRequestCustomData `json:"customData,omitempty"`
}

type NotifyMonitoringReportRequestEVSE struct {
	ID          int                                      `json:"id"`
	ConnectorID *int                                     `json:"connectorId,omitempty"`
	CustomData  *NotifyMonitoringReportRequestCustomData `json:"customData,omitempty"`
}

type NotifyMonitoringReportRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (NotifyMonitoringReportRequest) ActionName() string { return "NotifyMonitoringReport" }

func (NotifyMonitoringReportRequest) Version() protocol.Version { return protocol.OCPP21 }

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
