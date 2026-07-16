// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = GetDERControlRequest{}

var schemaGetDERControlRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"requestId": &validation.Schema{Type: "integer", AllowAdditional: true}, "isDefault": &validation.Schema{Type: "boolean", AllowAdditional: true}, "controlType": &validation.Schema{Type: "string", Enum: []string{"EnterService", "FreqDroop", "FreqWatt", "FixedPFAbsorb", "FixedPFInject", "FixedVar", "Gradients", "HFMustTrip", "HFMayTrip", "HVMustTrip", "HVMomCess", "HVMayTrip", "LimitMaxDischarge", "LFMustTrip", "LVMustTrip", "LVMomCess", "LVMayTrip", "PowerMonitoringMustTrip", "VoltVar", "VoltWatt", "WattPF", "WattVar"}}, "controlId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"requestId"}}

type GetDERControlRequest struct {
	RequestID   int                                 `json:"requestId"`
	IsDefault   *bool                               `json:"isDefault,omitempty"`
	ControlType *GetDERControlRequestDERControlEnum `json:"controlType,omitempty"`
	ControlID   *string                             `json:"controlId,omitempty"`
	CustomData  *GetDERControlRequestCustomData     `json:"customData,omitempty"`
}

type GetDERControlRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

type GetDERControlRequestDERControlEnum string

const (
	GetDERControlRequestDERControlEnumEnterService            GetDERControlRequestDERControlEnum = "EnterService"
	GetDERControlRequestDERControlEnumFreqDroop               GetDERControlRequestDERControlEnum = "FreqDroop"
	GetDERControlRequestDERControlEnumFreqWatt                GetDERControlRequestDERControlEnum = "FreqWatt"
	GetDERControlRequestDERControlEnumFixedPFAbsorb           GetDERControlRequestDERControlEnum = "FixedPFAbsorb"
	GetDERControlRequestDERControlEnumFixedPFInject           GetDERControlRequestDERControlEnum = "FixedPFInject"
	GetDERControlRequestDERControlEnumFixedVar                GetDERControlRequestDERControlEnum = "FixedVar"
	GetDERControlRequestDERControlEnumGradients               GetDERControlRequestDERControlEnum = "Gradients"
	GetDERControlRequestDERControlEnumHFMustTrip              GetDERControlRequestDERControlEnum = "HFMustTrip"
	GetDERControlRequestDERControlEnumHFMayTrip               GetDERControlRequestDERControlEnum = "HFMayTrip"
	GetDERControlRequestDERControlEnumHVMustTrip              GetDERControlRequestDERControlEnum = "HVMustTrip"
	GetDERControlRequestDERControlEnumHVMomCess               GetDERControlRequestDERControlEnum = "HVMomCess"
	GetDERControlRequestDERControlEnumHVMayTrip               GetDERControlRequestDERControlEnum = "HVMayTrip"
	GetDERControlRequestDERControlEnumLimitMaxDischarge       GetDERControlRequestDERControlEnum = "LimitMaxDischarge"
	GetDERControlRequestDERControlEnumLFMustTrip              GetDERControlRequestDERControlEnum = "LFMustTrip"
	GetDERControlRequestDERControlEnumLVMustTrip              GetDERControlRequestDERControlEnum = "LVMustTrip"
	GetDERControlRequestDERControlEnumLVMomCess               GetDERControlRequestDERControlEnum = "LVMomCess"
	GetDERControlRequestDERControlEnumLVMayTrip               GetDERControlRequestDERControlEnum = "LVMayTrip"
	GetDERControlRequestDERControlEnumPowerMonitoringMustTrip GetDERControlRequestDERControlEnum = "PowerMonitoringMustTrip"
	GetDERControlRequestDERControlEnumVoltVar                 GetDERControlRequestDERControlEnum = "VoltVar"
	GetDERControlRequestDERControlEnumVoltWatt                GetDERControlRequestDERControlEnum = "VoltWatt"
	GetDERControlRequestDERControlEnumWattPF                  GetDERControlRequestDERControlEnum = "WattPF"
	GetDERControlRequestDERControlEnumWattVar                 GetDERControlRequestDERControlEnum = "WattVar"
)

func (GetDERControlRequest) ActionName() string { return "GetDERControl" }

func (GetDERControlRequest) Version() protocol.Version { return protocol.OCPP21 }

func (GetDERControlRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (GetDERControlRequest) SchemaName() string { return "GetDERControlRequest.json" }

func (message GetDERControlRequest) Validate() error {
	return validation.Validate("GetDERControlRequest.json", schemaGetDERControlRequest, message)
}

func (GetDERControlRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetDERControlRequest.json", schemaGetDERControlRequest, data)
}
