// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = NotifyDERAlarmRequest{}

var schemaNotifyDERAlarmRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"controlType": &validation.Schema{Type: "string", Enum: []string{"EnterService", "FreqDroop", "FreqWatt", "FixedPFAbsorb", "FixedPFInject", "FixedVar", "Gradients", "HFMustTrip", "HFMayTrip", "HVMustTrip", "HVMomCess", "HVMayTrip", "LimitMaxDischarge", "LFMustTrip", "LVMustTrip", "LVMomCess", "LVMayTrip", "PowerMonitoringMustTrip", "VoltVar", "VoltWatt", "WattPF", "WattVar"}}, "gridEventFault": &validation.Schema{Type: "string", Enum: []string{"CurrentImbalance", "LocalEmergency", "LowInputPower", "OverCurrent", "OverFrequency", "OverVoltage", "PhaseRotation", "RemoteEmergency", "UnderFrequency", "UnderVoltage", "VoltageImbalance"}}, "alarmEnded": &validation.Schema{Type: "boolean", AllowAdditional: true}, "timestamp": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "extraInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 200, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"controlType", "timestamp"}}

type NotifyDERAlarmRequest struct {
	ControlType    NotifyDERAlarmRequestDERControlEnum      `json:"controlType"`
	GridEventFault *NotifyDERAlarmRequestGridEventFaultEnum `json:"gridEventFault,omitempty"`
	AlarmEnded     *bool                                    `json:"alarmEnded,omitempty"`
	Timestamp      string                                   `json:"timestamp"`
	ExtraInfo      *string                                  `json:"extraInfo,omitempty"`
	CustomData     *NotifyDERAlarmRequestCustomData         `json:"customData,omitempty"`
}

type NotifyDERAlarmRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

type NotifyDERAlarmRequestGridEventFaultEnum string

const (
	NotifyDERAlarmRequestGridEventFaultEnumCurrentImbalance NotifyDERAlarmRequestGridEventFaultEnum = "CurrentImbalance"
	NotifyDERAlarmRequestGridEventFaultEnumLocalEmergency   NotifyDERAlarmRequestGridEventFaultEnum = "LocalEmergency"
	NotifyDERAlarmRequestGridEventFaultEnumLowInputPower    NotifyDERAlarmRequestGridEventFaultEnum = "LowInputPower"
	NotifyDERAlarmRequestGridEventFaultEnumOverCurrent      NotifyDERAlarmRequestGridEventFaultEnum = "OverCurrent"
	NotifyDERAlarmRequestGridEventFaultEnumOverFrequency    NotifyDERAlarmRequestGridEventFaultEnum = "OverFrequency"
	NotifyDERAlarmRequestGridEventFaultEnumOverVoltage      NotifyDERAlarmRequestGridEventFaultEnum = "OverVoltage"
	NotifyDERAlarmRequestGridEventFaultEnumPhaseRotation    NotifyDERAlarmRequestGridEventFaultEnum = "PhaseRotation"
	NotifyDERAlarmRequestGridEventFaultEnumRemoteEmergency  NotifyDERAlarmRequestGridEventFaultEnum = "RemoteEmergency"
	NotifyDERAlarmRequestGridEventFaultEnumUnderFrequency   NotifyDERAlarmRequestGridEventFaultEnum = "UnderFrequency"
	NotifyDERAlarmRequestGridEventFaultEnumUnderVoltage     NotifyDERAlarmRequestGridEventFaultEnum = "UnderVoltage"
	NotifyDERAlarmRequestGridEventFaultEnumVoltageImbalance NotifyDERAlarmRequestGridEventFaultEnum = "VoltageImbalance"
)

type NotifyDERAlarmRequestDERControlEnum string

const (
	NotifyDERAlarmRequestDERControlEnumEnterService            NotifyDERAlarmRequestDERControlEnum = "EnterService"
	NotifyDERAlarmRequestDERControlEnumFreqDroop               NotifyDERAlarmRequestDERControlEnum = "FreqDroop"
	NotifyDERAlarmRequestDERControlEnumFreqWatt                NotifyDERAlarmRequestDERControlEnum = "FreqWatt"
	NotifyDERAlarmRequestDERControlEnumFixedPFAbsorb           NotifyDERAlarmRequestDERControlEnum = "FixedPFAbsorb"
	NotifyDERAlarmRequestDERControlEnumFixedPFInject           NotifyDERAlarmRequestDERControlEnum = "FixedPFInject"
	NotifyDERAlarmRequestDERControlEnumFixedVar                NotifyDERAlarmRequestDERControlEnum = "FixedVar"
	NotifyDERAlarmRequestDERControlEnumGradients               NotifyDERAlarmRequestDERControlEnum = "Gradients"
	NotifyDERAlarmRequestDERControlEnumHFMustTrip              NotifyDERAlarmRequestDERControlEnum = "HFMustTrip"
	NotifyDERAlarmRequestDERControlEnumHFMayTrip               NotifyDERAlarmRequestDERControlEnum = "HFMayTrip"
	NotifyDERAlarmRequestDERControlEnumHVMustTrip              NotifyDERAlarmRequestDERControlEnum = "HVMustTrip"
	NotifyDERAlarmRequestDERControlEnumHVMomCess               NotifyDERAlarmRequestDERControlEnum = "HVMomCess"
	NotifyDERAlarmRequestDERControlEnumHVMayTrip               NotifyDERAlarmRequestDERControlEnum = "HVMayTrip"
	NotifyDERAlarmRequestDERControlEnumLimitMaxDischarge       NotifyDERAlarmRequestDERControlEnum = "LimitMaxDischarge"
	NotifyDERAlarmRequestDERControlEnumLFMustTrip              NotifyDERAlarmRequestDERControlEnum = "LFMustTrip"
	NotifyDERAlarmRequestDERControlEnumLVMustTrip              NotifyDERAlarmRequestDERControlEnum = "LVMustTrip"
	NotifyDERAlarmRequestDERControlEnumLVMomCess               NotifyDERAlarmRequestDERControlEnum = "LVMomCess"
	NotifyDERAlarmRequestDERControlEnumLVMayTrip               NotifyDERAlarmRequestDERControlEnum = "LVMayTrip"
	NotifyDERAlarmRequestDERControlEnumPowerMonitoringMustTrip NotifyDERAlarmRequestDERControlEnum = "PowerMonitoringMustTrip"
	NotifyDERAlarmRequestDERControlEnumVoltVar                 NotifyDERAlarmRequestDERControlEnum = "VoltVar"
	NotifyDERAlarmRequestDERControlEnumVoltWatt                NotifyDERAlarmRequestDERControlEnum = "VoltWatt"
	NotifyDERAlarmRequestDERControlEnumWattPF                  NotifyDERAlarmRequestDERControlEnum = "WattPF"
	NotifyDERAlarmRequestDERControlEnumWattVar                 NotifyDERAlarmRequestDERControlEnum = "WattVar"
)

func (NotifyDERAlarmRequest) ActionName() string { return "NotifyDERAlarm" }

func (NotifyDERAlarmRequest) Version() protocol.Version { return protocol.OCPP21 }

func (NotifyDERAlarmRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (NotifyDERAlarmRequest) SchemaName() string { return "NotifyDERAlarmRequest.json" }

func (message NotifyDERAlarmRequest) Validate() error {
	return validation.Validate("NotifyDERAlarmRequest.json", schemaNotifyDERAlarmRequest, message)
}

func (NotifyDERAlarmRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("NotifyDERAlarmRequest.json", schemaNotifyDERAlarmRequest, data)
}
