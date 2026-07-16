// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = ClearDERControlRequest{}

var schemaClearDERControlRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"isDefault": &validation.Schema{Type: "boolean", AllowAdditional: true}, "controlType": &validation.Schema{Type: "string", Enum: []string{"EnterService", "FreqDroop", "FreqWatt", "FixedPFAbsorb", "FixedPFInject", "FixedVar", "Gradients", "HFMustTrip", "HFMayTrip", "HVMustTrip", "HVMomCess", "HVMayTrip", "LimitMaxDischarge", "LFMustTrip", "LVMustTrip", "LVMomCess", "LVMayTrip", "PowerMonitoringMustTrip", "VoltVar", "VoltWatt", "WattPF", "WattVar"}}, "controlId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"isDefault"}}

type ClearDERControlRequest struct {
	IsDefault   bool                                  `json:"isDefault"`
	ControlType *ClearDERControlRequestDERControlEnum `json:"controlType,omitempty"`
	ControlID   *string                               `json:"controlId,omitempty"`
	CustomData  *ClearDERControlRequestCustomData     `json:"customData,omitempty"`
}

type ClearDERControlRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

type ClearDERControlRequestDERControlEnum string

const (
	ClearDERControlRequestDERControlEnumEnterService            ClearDERControlRequestDERControlEnum = "EnterService"
	ClearDERControlRequestDERControlEnumFreqDroop               ClearDERControlRequestDERControlEnum = "FreqDroop"
	ClearDERControlRequestDERControlEnumFreqWatt                ClearDERControlRequestDERControlEnum = "FreqWatt"
	ClearDERControlRequestDERControlEnumFixedPFAbsorb           ClearDERControlRequestDERControlEnum = "FixedPFAbsorb"
	ClearDERControlRequestDERControlEnumFixedPFInject           ClearDERControlRequestDERControlEnum = "FixedPFInject"
	ClearDERControlRequestDERControlEnumFixedVar                ClearDERControlRequestDERControlEnum = "FixedVar"
	ClearDERControlRequestDERControlEnumGradients               ClearDERControlRequestDERControlEnum = "Gradients"
	ClearDERControlRequestDERControlEnumHFMustTrip              ClearDERControlRequestDERControlEnum = "HFMustTrip"
	ClearDERControlRequestDERControlEnumHFMayTrip               ClearDERControlRequestDERControlEnum = "HFMayTrip"
	ClearDERControlRequestDERControlEnumHVMustTrip              ClearDERControlRequestDERControlEnum = "HVMustTrip"
	ClearDERControlRequestDERControlEnumHVMomCess               ClearDERControlRequestDERControlEnum = "HVMomCess"
	ClearDERControlRequestDERControlEnumHVMayTrip               ClearDERControlRequestDERControlEnum = "HVMayTrip"
	ClearDERControlRequestDERControlEnumLimitMaxDischarge       ClearDERControlRequestDERControlEnum = "LimitMaxDischarge"
	ClearDERControlRequestDERControlEnumLFMustTrip              ClearDERControlRequestDERControlEnum = "LFMustTrip"
	ClearDERControlRequestDERControlEnumLVMustTrip              ClearDERControlRequestDERControlEnum = "LVMustTrip"
	ClearDERControlRequestDERControlEnumLVMomCess               ClearDERControlRequestDERControlEnum = "LVMomCess"
	ClearDERControlRequestDERControlEnumLVMayTrip               ClearDERControlRequestDERControlEnum = "LVMayTrip"
	ClearDERControlRequestDERControlEnumPowerMonitoringMustTrip ClearDERControlRequestDERControlEnum = "PowerMonitoringMustTrip"
	ClearDERControlRequestDERControlEnumVoltVar                 ClearDERControlRequestDERControlEnum = "VoltVar"
	ClearDERControlRequestDERControlEnumVoltWatt                ClearDERControlRequestDERControlEnum = "VoltWatt"
	ClearDERControlRequestDERControlEnumWattPF                  ClearDERControlRequestDERControlEnum = "WattPF"
	ClearDERControlRequestDERControlEnumWattVar                 ClearDERControlRequestDERControlEnum = "WattVar"
)

func (ClearDERControlRequest) ActionName() string { return "ClearDERControl" }

func (ClearDERControlRequest) Version() protocol.Version { return protocol.OCPP21 }

func (ClearDERControlRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (ClearDERControlRequest) SchemaName() string { return "ClearDERControlRequest.json" }

func (message ClearDERControlRequest) Validate() error {
	return validation.Validate("ClearDERControlRequest.json", schemaClearDERControlRequest, message)
}

func (ClearDERControlRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ClearDERControlRequest.json", schemaClearDERControlRequest, data)
}
