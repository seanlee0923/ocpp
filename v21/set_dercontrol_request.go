// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = SetDERControlRequest{}

var schemaSetDERControlRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"isDefault": &validation.Schema{Type: "boolean", AllowAdditional: true}, "controlId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "controlType": &validation.Schema{Type: "string", Enum: []string{"EnterService", "FreqDroop", "FreqWatt", "FixedPFAbsorb", "FixedPFInject", "FixedVar", "Gradients", "HFMustTrip", "HFMayTrip", "HVMustTrip", "HVMomCess", "HVMayTrip", "LimitMaxDischarge", "LFMustTrip", "LVMustTrip", "LVMomCess", "LVMayTrip", "PowerMonitoringMustTrip", "VoltVar", "VoltWatt", "WattPF", "WattVar"}}, "curve": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"curveData": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"x": &validation.Schema{Type: "number", AllowAdditional: true}, "y": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"x", "y"}}, MinItems: 1, HasMinItems: true, MaxItems: 10, HasMaxItems: true}, "hysteresis": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"hysteresisHigh": &validation.Schema{Type: "number", AllowAdditional: true}, "hysteresisLow": &validation.Schema{Type: "number", AllowAdditional: true}, "hysteresisDelay": &validation.Schema{Type: "number", AllowAdditional: true}, "hysteresisGradient": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "priority": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "reactivePowerParams": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vRef": &validation.Schema{Type: "number", AllowAdditional: true}, "autonomousVRefEnable": &validation.Schema{Type: "boolean", AllowAdditional: true}, "autonomousVRefTimeConstant": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "voltageParams": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"hv10MinMeanValue": &validation.Schema{Type: "number", AllowAdditional: true}, "hv10MinMeanTripDelay": &validation.Schema{Type: "number", AllowAdditional: true}, "powerDuringCessation": &validation.Schema{Type: "string", Enum: []string{"Active", "Reactive"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "yUnit": &validation.Schema{Type: "string", Enum: []string{"Not_Applicable", "PctMaxW", "PctMaxVar", "PctWAvail", "PctVarAvail", "PctEffectiveV"}}, "responseTime": &validation.Schema{Type: "number", AllowAdditional: true}, "startTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "duration": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"priority", "yUnit", "curveData"}}, "enterService": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"priority": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "highVoltage": &validation.Schema{Type: "number", AllowAdditional: true}, "lowVoltage": &validation.Schema{Type: "number", AllowAdditional: true}, "highFreq": &validation.Schema{Type: "number", AllowAdditional: true}, "lowFreq": &validation.Schema{Type: "number", AllowAdditional: true}, "delay": &validation.Schema{Type: "number", AllowAdditional: true}, "randomDelay": &validation.Schema{Type: "number", AllowAdditional: true}, "rampRate": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"priority", "highVoltage", "lowVoltage", "highFreq", "lowFreq"}}, "fixedPFAbsorb": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"priority": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "displacement": &validation.Schema{Type: "number", AllowAdditional: true}, "excitation": &validation.Schema{Type: "boolean", AllowAdditional: true}, "startTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "duration": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"priority", "displacement", "excitation"}}, "fixedPFInject": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"priority": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "displacement": &validation.Schema{Type: "number", AllowAdditional: true}, "excitation": &validation.Schema{Type: "boolean", AllowAdditional: true}, "startTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "duration": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"priority", "displacement", "excitation"}}, "fixedVar": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"priority": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "setpoint": &validation.Schema{Type: "number", AllowAdditional: true}, "unit": &validation.Schema{Type: "string", Enum: []string{"Not_Applicable", "PctMaxW", "PctMaxVar", "PctWAvail", "PctVarAvail", "PctEffectiveV"}}, "startTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "duration": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"priority", "setpoint", "unit"}}, "freqDroop": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"priority": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "overFreq": &validation.Schema{Type: "number", AllowAdditional: true}, "underFreq": &validation.Schema{Type: "number", AllowAdditional: true}, "overDroop": &validation.Schema{Type: "number", AllowAdditional: true}, "underDroop": &validation.Schema{Type: "number", AllowAdditional: true}, "responseTime": &validation.Schema{Type: "number", AllowAdditional: true}, "startTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "duration": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"priority", "overFreq", "underFreq", "overDroop", "underDroop", "responseTime"}}, "gradient": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"priority": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "gradient": &validation.Schema{Type: "number", AllowAdditional: true}, "softGradient": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"priority", "gradient", "softGradient"}}, "limitMaxDischarge": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"priority": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "pctMaxDischargePower": &validation.Schema{Type: "number", AllowAdditional: true}, "powerMonitoringMustTrip": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"curveData": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"x": &validation.Schema{Type: "number", AllowAdditional: true}, "y": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"x", "y"}}, MinItems: 1, HasMinItems: true, MaxItems: 10, HasMaxItems: true}, "hysteresis": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"hysteresisHigh": &validation.Schema{Type: "number", AllowAdditional: true}, "hysteresisLow": &validation.Schema{Type: "number", AllowAdditional: true}, "hysteresisDelay": &validation.Schema{Type: "number", AllowAdditional: true}, "hysteresisGradient": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "priority": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "reactivePowerParams": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vRef": &validation.Schema{Type: "number", AllowAdditional: true}, "autonomousVRefEnable": &validation.Schema{Type: "boolean", AllowAdditional: true}, "autonomousVRefTimeConstant": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "voltageParams": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"hv10MinMeanValue": &validation.Schema{Type: "number", AllowAdditional: true}, "hv10MinMeanTripDelay": &validation.Schema{Type: "number", AllowAdditional: true}, "powerDuringCessation": &validation.Schema{Type: "string", Enum: []string{"Active", "Reactive"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "yUnit": &validation.Schema{Type: "string", Enum: []string{"Not_Applicable", "PctMaxW", "PctMaxVar", "PctWAvail", "PctVarAvail", "PctEffectiveV"}}, "responseTime": &validation.Schema{Type: "number", AllowAdditional: true}, "startTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "duration": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"priority", "yUnit", "curveData"}}, "startTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "duration": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"priority"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"isDefault", "controlId", "controlType"}}

type SetDERControlRequest struct {
	IsDefault         bool                                   `json:"isDefault"`
	ControlID         string                                 `json:"controlId"`
	ControlType       SetDERControlRequestDERControlEnum     `json:"controlType"`
	Curve             *SetDERControlRequestDERCurve          `json:"curve,omitempty"`
	EnterService      *SetDERControlRequestEnterService      `json:"enterService,omitempty"`
	FixedPFAbsorb     *SetDERControlRequestFixedPF           `json:"fixedPFAbsorb,omitempty"`
	FixedPFInject     *SetDERControlRequestFixedPF           `json:"fixedPFInject,omitempty"`
	FixedVar          *SetDERControlRequestFixedVar          `json:"fixedVar,omitempty"`
	FreqDroop         *SetDERControlRequestFreqDroop         `json:"freqDroop,omitempty"`
	Gradient          *SetDERControlRequestGradient          `json:"gradient,omitempty"`
	LimitMaxDischarge *SetDERControlRequestLimitMaxDischarge `json:"limitMaxDischarge,omitempty"`
	CustomData        *SetDERControlRequestCustomData        `json:"customData,omitempty"`
}

type SetDERControlRequestLimitMaxDischarge struct {
	Priority                int                             `json:"priority"`
	PctMaxDischargePower    *float64                        `json:"pctMaxDischargePower,omitempty"`
	PowerMonitoringMustTrip *SetDERControlRequestDERCurve   `json:"powerMonitoringMustTrip,omitempty"`
	StartTime               *string                         `json:"startTime,omitempty"`
	Duration                *float64                        `json:"duration,omitempty"`
	CustomData              *SetDERControlRequestCustomData `json:"customData,omitempty"`
}

type SetDERControlRequestGradient struct {
	Priority     int                             `json:"priority"`
	Gradient     float64                         `json:"gradient"`
	SoftGradient float64                         `json:"softGradient"`
	CustomData   *SetDERControlRequestCustomData `json:"customData,omitempty"`
}

type SetDERControlRequestFreqDroop struct {
	Priority     int                             `json:"priority"`
	OverFreq     float64                         `json:"overFreq"`
	UnderFreq    float64                         `json:"underFreq"`
	OverDroop    float64                         `json:"overDroop"`
	UnderDroop   float64                         `json:"underDroop"`
	ResponseTime float64                         `json:"responseTime"`
	StartTime    *string                         `json:"startTime,omitempty"`
	Duration     *float64                        `json:"duration,omitempty"`
	CustomData   *SetDERControlRequestCustomData `json:"customData,omitempty"`
}

type SetDERControlRequestFixedVar struct {
	Priority   int                             `json:"priority"`
	Setpoint   float64                         `json:"setpoint"`
	Unit       SetDERControlRequestDERUnitEnum `json:"unit"`
	StartTime  *string                         `json:"startTime,omitempty"`
	Duration   *float64                        `json:"duration,omitempty"`
	CustomData *SetDERControlRequestCustomData `json:"customData,omitempty"`
}

type SetDERControlRequestFixedPF struct {
	Priority     int                             `json:"priority"`
	Displacement float64                         `json:"displacement"`
	Excitation   bool                            `json:"excitation"`
	StartTime    *string                         `json:"startTime,omitempty"`
	Duration     *float64                        `json:"duration,omitempty"`
	CustomData   *SetDERControlRequestCustomData `json:"customData,omitempty"`
}

type SetDERControlRequestEnterService struct {
	Priority    int                             `json:"priority"`
	HighVoltage float64                         `json:"highVoltage"`
	LowVoltage  float64                         `json:"lowVoltage"`
	HighFreq    float64                         `json:"highFreq"`
	LowFreq     float64                         `json:"lowFreq"`
	Delay       *float64                        `json:"delay,omitempty"`
	RandomDelay *float64                        `json:"randomDelay,omitempty"`
	RampRate    *float64                        `json:"rampRate,omitempty"`
	CustomData  *SetDERControlRequestCustomData `json:"customData,omitempty"`
}

type SetDERControlRequestDERCurve struct {
	CurveData           []SetDERControlRequestDERCurvePoints     `json:"curveData"`
	Hysteresis          *SetDERControlRequestHysteresis          `json:"hysteresis,omitempty"`
	Priority            int                                      `json:"priority"`
	ReactivePowerParams *SetDERControlRequestReactivePowerParams `json:"reactivePowerParams,omitempty"`
	VoltageParams       *SetDERControlRequestVoltageParams       `json:"voltageParams,omitempty"`
	YUnit               SetDERControlRequestDERUnitEnum          `json:"yUnit"`
	ResponseTime        *float64                                 `json:"responseTime,omitempty"`
	StartTime           *string                                  `json:"startTime,omitempty"`
	Duration            *float64                                 `json:"duration,omitempty"`
	CustomData          *SetDERControlRequestCustomData          `json:"customData,omitempty"`
}

type SetDERControlRequestDERUnitEnum string

const (
	SetDERControlRequestDERUnitEnumNotApplicable SetDERControlRequestDERUnitEnum = "Not_Applicable"
	SetDERControlRequestDERUnitEnumPctMaxW       SetDERControlRequestDERUnitEnum = "PctMaxW"
	SetDERControlRequestDERUnitEnumPctMaxVar     SetDERControlRequestDERUnitEnum = "PctMaxVar"
	SetDERControlRequestDERUnitEnumPctWAvail     SetDERControlRequestDERUnitEnum = "PctWAvail"
	SetDERControlRequestDERUnitEnumPctVarAvail   SetDERControlRequestDERUnitEnum = "PctVarAvail"
	SetDERControlRequestDERUnitEnumPctEffectiveV SetDERControlRequestDERUnitEnum = "PctEffectiveV"
)

type SetDERControlRequestVoltageParams struct {
	Hv10MinMeanValue     *float64                                      `json:"hv10MinMeanValue,omitempty"`
	Hv10MinMeanTripDelay *float64                                      `json:"hv10MinMeanTripDelay,omitempty"`
	PowerDuringCessation *SetDERControlRequestPowerDuringCessationEnum `json:"powerDuringCessation,omitempty"`
	CustomData           *SetDERControlRequestCustomData               `json:"customData,omitempty"`
}

type SetDERControlRequestPowerDuringCessationEnum string

const (
	SetDERControlRequestPowerDuringCessationEnumActive   SetDERControlRequestPowerDuringCessationEnum = "Active"
	SetDERControlRequestPowerDuringCessationEnumReactive SetDERControlRequestPowerDuringCessationEnum = "Reactive"
)

type SetDERControlRequestReactivePowerParams struct {
	VRef                       *float64                        `json:"vRef,omitempty"`
	AutonomousVRefEnable       *bool                           `json:"autonomousVRefEnable,omitempty"`
	AutonomousVRefTimeConstant *float64                        `json:"autonomousVRefTimeConstant,omitempty"`
	CustomData                 *SetDERControlRequestCustomData `json:"customData,omitempty"`
}

type SetDERControlRequestHysteresis struct {
	HysteresisHigh     *float64                        `json:"hysteresisHigh,omitempty"`
	HysteresisLow      *float64                        `json:"hysteresisLow,omitempty"`
	HysteresisDelay    *float64                        `json:"hysteresisDelay,omitempty"`
	HysteresisGradient *float64                        `json:"hysteresisGradient,omitempty"`
	CustomData         *SetDERControlRequestCustomData `json:"customData,omitempty"`
}

type SetDERControlRequestDERCurvePoints struct {
	X          float64                         `json:"x"`
	Y          float64                         `json:"y"`
	CustomData *SetDERControlRequestCustomData `json:"customData,omitempty"`
}

type SetDERControlRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value SetDERControlRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *SetDERControlRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

type SetDERControlRequestDERControlEnum string

const (
	SetDERControlRequestDERControlEnumEnterService            SetDERControlRequestDERControlEnum = "EnterService"
	SetDERControlRequestDERControlEnumFreqDroop               SetDERControlRequestDERControlEnum = "FreqDroop"
	SetDERControlRequestDERControlEnumFreqWatt                SetDERControlRequestDERControlEnum = "FreqWatt"
	SetDERControlRequestDERControlEnumFixedPFAbsorb           SetDERControlRequestDERControlEnum = "FixedPFAbsorb"
	SetDERControlRequestDERControlEnumFixedPFInject           SetDERControlRequestDERControlEnum = "FixedPFInject"
	SetDERControlRequestDERControlEnumFixedVar                SetDERControlRequestDERControlEnum = "FixedVar"
	SetDERControlRequestDERControlEnumGradients               SetDERControlRequestDERControlEnum = "Gradients"
	SetDERControlRequestDERControlEnumHFMustTrip              SetDERControlRequestDERControlEnum = "HFMustTrip"
	SetDERControlRequestDERControlEnumHFMayTrip               SetDERControlRequestDERControlEnum = "HFMayTrip"
	SetDERControlRequestDERControlEnumHVMustTrip              SetDERControlRequestDERControlEnum = "HVMustTrip"
	SetDERControlRequestDERControlEnumHVMomCess               SetDERControlRequestDERControlEnum = "HVMomCess"
	SetDERControlRequestDERControlEnumHVMayTrip               SetDERControlRequestDERControlEnum = "HVMayTrip"
	SetDERControlRequestDERControlEnumLimitMaxDischarge       SetDERControlRequestDERControlEnum = "LimitMaxDischarge"
	SetDERControlRequestDERControlEnumLFMustTrip              SetDERControlRequestDERControlEnum = "LFMustTrip"
	SetDERControlRequestDERControlEnumLVMustTrip              SetDERControlRequestDERControlEnum = "LVMustTrip"
	SetDERControlRequestDERControlEnumLVMomCess               SetDERControlRequestDERControlEnum = "LVMomCess"
	SetDERControlRequestDERControlEnumLVMayTrip               SetDERControlRequestDERControlEnum = "LVMayTrip"
	SetDERControlRequestDERControlEnumPowerMonitoringMustTrip SetDERControlRequestDERControlEnum = "PowerMonitoringMustTrip"
	SetDERControlRequestDERControlEnumVoltVar                 SetDERControlRequestDERControlEnum = "VoltVar"
	SetDERControlRequestDERControlEnumVoltWatt                SetDERControlRequestDERControlEnum = "VoltWatt"
	SetDERControlRequestDERControlEnumWattPF                  SetDERControlRequestDERControlEnum = "WattPF"
	SetDERControlRequestDERControlEnumWattVar                 SetDERControlRequestDERControlEnum = "WattVar"
)

func (SetDERControlRequest) ActionName() string { return "SetDERControl" }

func (SetDERControlRequest) Version() protocol.Version { return protocol.OCPP21 }

func (SetDERControlRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (SetDERControlRequest) SchemaName() string { return "SetDERControlRequest.json" }

func (message SetDERControlRequest) Validate() error {
	return validation.Validate("SetDERControlRequest.json", schemaSetDERControlRequest, message)
}

func (SetDERControlRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("SetDERControlRequest.json", schemaSetDERControlRequest, data)
}
