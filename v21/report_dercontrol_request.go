// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = ReportDERControlRequest{}

var schemaReportDERControlRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"curve": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"curve": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"curveData": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"x": &validation.Schema{Type: "number", AllowAdditional: true}, "y": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"x", "y"}}, MinItems: 1, HasMinItems: true, MaxItems: 10, HasMaxItems: true}, "hysteresis": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"hysteresisHigh": &validation.Schema{Type: "number", AllowAdditional: true}, "hysteresisLow": &validation.Schema{Type: "number", AllowAdditional: true}, "hysteresisDelay": &validation.Schema{Type: "number", AllowAdditional: true}, "hysteresisGradient": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "priority": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "reactivePowerParams": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vRef": &validation.Schema{Type: "number", AllowAdditional: true}, "autonomousVRefEnable": &validation.Schema{Type: "boolean", AllowAdditional: true}, "autonomousVRefTimeConstant": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "voltageParams": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"hv10MinMeanValue": &validation.Schema{Type: "number", AllowAdditional: true}, "hv10MinMeanTripDelay": &validation.Schema{Type: "number", AllowAdditional: true}, "powerDuringCessation": &validation.Schema{Type: "string", Enum: []string{"Active", "Reactive"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "yUnit": &validation.Schema{Type: "string", Enum: []string{"Not_Applicable", "PctMaxW", "PctMaxVar", "PctWAvail", "PctVarAvail", "PctEffectiveV"}}, "responseTime": &validation.Schema{Type: "number", AllowAdditional: true}, "startTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "duration": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"priority", "yUnit", "curveData"}}, "id": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "curveType": &validation.Schema{Type: "string", Enum: []string{"EnterService", "FreqDroop", "FreqWatt", "FixedPFAbsorb", "FixedPFInject", "FixedVar", "Gradients", "HFMustTrip", "HFMayTrip", "HVMustTrip", "HVMomCess", "HVMayTrip", "LimitMaxDischarge", "LFMustTrip", "LVMustTrip", "LVMomCess", "LVMayTrip", "PowerMonitoringMustTrip", "VoltVar", "VoltWatt", "WattPF", "WattVar"}}, "isDefault": &validation.Schema{Type: "boolean", AllowAdditional: true}, "isSuperseded": &validation.Schema{Type: "boolean", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id", "curveType", "isDefault", "isSuperseded", "curve"}}, MinItems: 1, HasMinItems: true, MaxItems: 24, HasMaxItems: true}, "enterService": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"enterService": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"priority": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "highVoltage": &validation.Schema{Type: "number", AllowAdditional: true}, "lowVoltage": &validation.Schema{Type: "number", AllowAdditional: true}, "highFreq": &validation.Schema{Type: "number", AllowAdditional: true}, "lowFreq": &validation.Schema{Type: "number", AllowAdditional: true}, "delay": &validation.Schema{Type: "number", AllowAdditional: true}, "randomDelay": &validation.Schema{Type: "number", AllowAdditional: true}, "rampRate": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"priority", "highVoltage", "lowVoltage", "highFreq", "lowFreq"}}, "id": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id", "enterService"}}, MinItems: 1, HasMinItems: true, MaxItems: 24, HasMaxItems: true}, "fixedPFAbsorb": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"fixedPF": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"priority": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "displacement": &validation.Schema{Type: "number", AllowAdditional: true}, "excitation": &validation.Schema{Type: "boolean", AllowAdditional: true}, "startTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "duration": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"priority", "displacement", "excitation"}}, "id": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "isDefault": &validation.Schema{Type: "boolean", AllowAdditional: true}, "isSuperseded": &validation.Schema{Type: "boolean", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id", "isDefault", "isSuperseded", "fixedPF"}}, MinItems: 1, HasMinItems: true, MaxItems: 24, HasMaxItems: true}, "fixedPFInject": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"fixedPF": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"priority": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "displacement": &validation.Schema{Type: "number", AllowAdditional: true}, "excitation": &validation.Schema{Type: "boolean", AllowAdditional: true}, "startTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "duration": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"priority", "displacement", "excitation"}}, "id": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "isDefault": &validation.Schema{Type: "boolean", AllowAdditional: true}, "isSuperseded": &validation.Schema{Type: "boolean", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id", "isDefault", "isSuperseded", "fixedPF"}}, MinItems: 1, HasMinItems: true, MaxItems: 24, HasMaxItems: true}, "fixedVar": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"fixedVar": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"priority": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "setpoint": &validation.Schema{Type: "number", AllowAdditional: true}, "unit": &validation.Schema{Type: "string", Enum: []string{"Not_Applicable", "PctMaxW", "PctMaxVar", "PctWAvail", "PctVarAvail", "PctEffectiveV"}}, "startTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "duration": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"priority", "setpoint", "unit"}}, "id": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "isDefault": &validation.Schema{Type: "boolean", AllowAdditional: true}, "isSuperseded": &validation.Schema{Type: "boolean", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id", "isDefault", "isSuperseded", "fixedVar"}}, MinItems: 1, HasMinItems: true, MaxItems: 24, HasMaxItems: true}, "freqDroop": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"freqDroop": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"priority": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "overFreq": &validation.Schema{Type: "number", AllowAdditional: true}, "underFreq": &validation.Schema{Type: "number", AllowAdditional: true}, "overDroop": &validation.Schema{Type: "number", AllowAdditional: true}, "underDroop": &validation.Schema{Type: "number", AllowAdditional: true}, "responseTime": &validation.Schema{Type: "number", AllowAdditional: true}, "startTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "duration": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"priority", "overFreq", "underFreq", "overDroop", "underDroop", "responseTime"}}, "id": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "isDefault": &validation.Schema{Type: "boolean", AllowAdditional: true}, "isSuperseded": &validation.Schema{Type: "boolean", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id", "isDefault", "isSuperseded", "freqDroop"}}, MinItems: 1, HasMinItems: true, MaxItems: 24, HasMaxItems: true}, "gradient": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"gradient": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"priority": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "gradient": &validation.Schema{Type: "number", AllowAdditional: true}, "softGradient": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"priority", "gradient", "softGradient"}}, "id": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id", "gradient"}}, MinItems: 1, HasMinItems: true, MaxItems: 24, HasMaxItems: true}, "limitMaxDischarge": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"id": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "isDefault": &validation.Schema{Type: "boolean", AllowAdditional: true}, "isSuperseded": &validation.Schema{Type: "boolean", AllowAdditional: true}, "limitMaxDischarge": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"priority": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "pctMaxDischargePower": &validation.Schema{Type: "number", AllowAdditional: true}, "powerMonitoringMustTrip": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"curveData": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"x": &validation.Schema{Type: "number", AllowAdditional: true}, "y": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"x", "y"}}, MinItems: 1, HasMinItems: true, MaxItems: 10, HasMaxItems: true}, "hysteresis": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"hysteresisHigh": &validation.Schema{Type: "number", AllowAdditional: true}, "hysteresisLow": &validation.Schema{Type: "number", AllowAdditional: true}, "hysteresisDelay": &validation.Schema{Type: "number", AllowAdditional: true}, "hysteresisGradient": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "priority": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "reactivePowerParams": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vRef": &validation.Schema{Type: "number", AllowAdditional: true}, "autonomousVRefEnable": &validation.Schema{Type: "boolean", AllowAdditional: true}, "autonomousVRefTimeConstant": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "voltageParams": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"hv10MinMeanValue": &validation.Schema{Type: "number", AllowAdditional: true}, "hv10MinMeanTripDelay": &validation.Schema{Type: "number", AllowAdditional: true}, "powerDuringCessation": &validation.Schema{Type: "string", Enum: []string{"Active", "Reactive"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "yUnit": &validation.Schema{Type: "string", Enum: []string{"Not_Applicable", "PctMaxW", "PctMaxVar", "PctWAvail", "PctVarAvail", "PctEffectiveV"}}, "responseTime": &validation.Schema{Type: "number", AllowAdditional: true}, "startTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "duration": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"priority", "yUnit", "curveData"}}, "startTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "duration": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"priority"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id", "isDefault", "isSuperseded", "limitMaxDischarge"}}, MinItems: 1, HasMinItems: true, MaxItems: 24, HasMaxItems: true}, "requestId": &validation.Schema{Type: "integer", AllowAdditional: true}, "tbc": &validation.Schema{Type: "boolean", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"requestId"}}

type ReportDERControlRequest struct {
	Curve             []ReportDERControlRequestDERCurveGet          `json:"curve,omitempty"`
	EnterService      []ReportDERControlRequestEnterServiceGet      `json:"enterService,omitempty"`
	FixedPFAbsorb     []ReportDERControlRequestFixedPFGet           `json:"fixedPFAbsorb,omitempty"`
	FixedPFInject     []ReportDERControlRequestFixedPFGet           `json:"fixedPFInject,omitempty"`
	FixedVar          []ReportDERControlRequestFixedVarGet          `json:"fixedVar,omitempty"`
	FreqDroop         []ReportDERControlRequestFreqDroopGet         `json:"freqDroop,omitempty"`
	Gradient          []ReportDERControlRequestGradientGet          `json:"gradient,omitempty"`
	LimitMaxDischarge []ReportDERControlRequestLimitMaxDischargeGet `json:"limitMaxDischarge,omitempty"`
	RequestID         int                                           `json:"requestId"`
	Tbc               *bool                                         `json:"tbc,omitempty"`
	CustomData        *ReportDERControlRequestCustomData            `json:"customData,omitempty"`
}

type ReportDERControlRequestLimitMaxDischargeGet struct {
	ID                string                                   `json:"id"`
	IsDefault         bool                                     `json:"isDefault"`
	IsSuperseded      bool                                     `json:"isSuperseded"`
	LimitMaxDischarge ReportDERControlRequestLimitMaxDischarge `json:"limitMaxDischarge"`
	CustomData        *ReportDERControlRequestCustomData       `json:"customData,omitempty"`
}

type ReportDERControlRequestLimitMaxDischarge struct {
	Priority                int                                `json:"priority"`
	PctMaxDischargePower    *float64                           `json:"pctMaxDischargePower,omitempty"`
	PowerMonitoringMustTrip *ReportDERControlRequestDERCurve   `json:"powerMonitoringMustTrip,omitempty"`
	StartTime               *string                            `json:"startTime,omitempty"`
	Duration                *float64                           `json:"duration,omitempty"`
	CustomData              *ReportDERControlRequestCustomData `json:"customData,omitempty"`
}

type ReportDERControlRequestGradientGet struct {
	Gradient   ReportDERControlRequestGradient    `json:"gradient"`
	ID         string                             `json:"id"`
	CustomData *ReportDERControlRequestCustomData `json:"customData,omitempty"`
}

type ReportDERControlRequestGradient struct {
	Priority     int                                `json:"priority"`
	Gradient     float64                            `json:"gradient"`
	SoftGradient float64                            `json:"softGradient"`
	CustomData   *ReportDERControlRequestCustomData `json:"customData,omitempty"`
}

type ReportDERControlRequestFreqDroopGet struct {
	FreqDroop    ReportDERControlRequestFreqDroop   `json:"freqDroop"`
	ID           string                             `json:"id"`
	IsDefault    bool                               `json:"isDefault"`
	IsSuperseded bool                               `json:"isSuperseded"`
	CustomData   *ReportDERControlRequestCustomData `json:"customData,omitempty"`
}

type ReportDERControlRequestFreqDroop struct {
	Priority     int                                `json:"priority"`
	OverFreq     float64                            `json:"overFreq"`
	UnderFreq    float64                            `json:"underFreq"`
	OverDroop    float64                            `json:"overDroop"`
	UnderDroop   float64                            `json:"underDroop"`
	ResponseTime float64                            `json:"responseTime"`
	StartTime    *string                            `json:"startTime,omitempty"`
	Duration     *float64                           `json:"duration,omitempty"`
	CustomData   *ReportDERControlRequestCustomData `json:"customData,omitempty"`
}

type ReportDERControlRequestFixedVarGet struct {
	FixedVar     ReportDERControlRequestFixedVar    `json:"fixedVar"`
	ID           string                             `json:"id"`
	IsDefault    bool                               `json:"isDefault"`
	IsSuperseded bool                               `json:"isSuperseded"`
	CustomData   *ReportDERControlRequestCustomData `json:"customData,omitempty"`
}

type ReportDERControlRequestFixedVar struct {
	Priority   int                                `json:"priority"`
	Setpoint   float64                            `json:"setpoint"`
	Unit       ReportDERControlRequestDERUnitEnum `json:"unit"`
	StartTime  *string                            `json:"startTime,omitempty"`
	Duration   *float64                           `json:"duration,omitempty"`
	CustomData *ReportDERControlRequestCustomData `json:"customData,omitempty"`
}

type ReportDERControlRequestFixedPFGet struct {
	FixedPF      ReportDERControlRequestFixedPF     `json:"fixedPF"`
	ID           string                             `json:"id"`
	IsDefault    bool                               `json:"isDefault"`
	IsSuperseded bool                               `json:"isSuperseded"`
	CustomData   *ReportDERControlRequestCustomData `json:"customData,omitempty"`
}

type ReportDERControlRequestFixedPF struct {
	Priority     int                                `json:"priority"`
	Displacement float64                            `json:"displacement"`
	Excitation   bool                               `json:"excitation"`
	StartTime    *string                            `json:"startTime,omitempty"`
	Duration     *float64                           `json:"duration,omitempty"`
	CustomData   *ReportDERControlRequestCustomData `json:"customData,omitempty"`
}

type ReportDERControlRequestEnterServiceGet struct {
	EnterService ReportDERControlRequestEnterService `json:"enterService"`
	ID           string                              `json:"id"`
	CustomData   *ReportDERControlRequestCustomData  `json:"customData,omitempty"`
}

type ReportDERControlRequestEnterService struct {
	Priority    int                                `json:"priority"`
	HighVoltage float64                            `json:"highVoltage"`
	LowVoltage  float64                            `json:"lowVoltage"`
	HighFreq    float64                            `json:"highFreq"`
	LowFreq     float64                            `json:"lowFreq"`
	Delay       *float64                           `json:"delay,omitempty"`
	RandomDelay *float64                           `json:"randomDelay,omitempty"`
	RampRate    *float64                           `json:"rampRate,omitempty"`
	CustomData  *ReportDERControlRequestCustomData `json:"customData,omitempty"`
}

type ReportDERControlRequestDERCurveGet struct {
	Curve        ReportDERControlRequestDERCurve       `json:"curve"`
	ID           string                                `json:"id"`
	CurveType    ReportDERControlRequestDERControlEnum `json:"curveType"`
	IsDefault    bool                                  `json:"isDefault"`
	IsSuperseded bool                                  `json:"isSuperseded"`
	CustomData   *ReportDERControlRequestCustomData    `json:"customData,omitempty"`
}

type ReportDERControlRequestDERControlEnum string

const (
	ReportDERControlRequestDERControlEnumEnterService            ReportDERControlRequestDERControlEnum = "EnterService"
	ReportDERControlRequestDERControlEnumFreqDroop               ReportDERControlRequestDERControlEnum = "FreqDroop"
	ReportDERControlRequestDERControlEnumFreqWatt                ReportDERControlRequestDERControlEnum = "FreqWatt"
	ReportDERControlRequestDERControlEnumFixedPFAbsorb           ReportDERControlRequestDERControlEnum = "FixedPFAbsorb"
	ReportDERControlRequestDERControlEnumFixedPFInject           ReportDERControlRequestDERControlEnum = "FixedPFInject"
	ReportDERControlRequestDERControlEnumFixedVar                ReportDERControlRequestDERControlEnum = "FixedVar"
	ReportDERControlRequestDERControlEnumGradients               ReportDERControlRequestDERControlEnum = "Gradients"
	ReportDERControlRequestDERControlEnumHFMustTrip              ReportDERControlRequestDERControlEnum = "HFMustTrip"
	ReportDERControlRequestDERControlEnumHFMayTrip               ReportDERControlRequestDERControlEnum = "HFMayTrip"
	ReportDERControlRequestDERControlEnumHVMustTrip              ReportDERControlRequestDERControlEnum = "HVMustTrip"
	ReportDERControlRequestDERControlEnumHVMomCess               ReportDERControlRequestDERControlEnum = "HVMomCess"
	ReportDERControlRequestDERControlEnumHVMayTrip               ReportDERControlRequestDERControlEnum = "HVMayTrip"
	ReportDERControlRequestDERControlEnumLimitMaxDischarge       ReportDERControlRequestDERControlEnum = "LimitMaxDischarge"
	ReportDERControlRequestDERControlEnumLFMustTrip              ReportDERControlRequestDERControlEnum = "LFMustTrip"
	ReportDERControlRequestDERControlEnumLVMustTrip              ReportDERControlRequestDERControlEnum = "LVMustTrip"
	ReportDERControlRequestDERControlEnumLVMomCess               ReportDERControlRequestDERControlEnum = "LVMomCess"
	ReportDERControlRequestDERControlEnumLVMayTrip               ReportDERControlRequestDERControlEnum = "LVMayTrip"
	ReportDERControlRequestDERControlEnumPowerMonitoringMustTrip ReportDERControlRequestDERControlEnum = "PowerMonitoringMustTrip"
	ReportDERControlRequestDERControlEnumVoltVar                 ReportDERControlRequestDERControlEnum = "VoltVar"
	ReportDERControlRequestDERControlEnumVoltWatt                ReportDERControlRequestDERControlEnum = "VoltWatt"
	ReportDERControlRequestDERControlEnumWattPF                  ReportDERControlRequestDERControlEnum = "WattPF"
	ReportDERControlRequestDERControlEnumWattVar                 ReportDERControlRequestDERControlEnum = "WattVar"
)

type ReportDERControlRequestDERCurve struct {
	CurveData           []ReportDERControlRequestDERCurvePoints     `json:"curveData"`
	Hysteresis          *ReportDERControlRequestHysteresis          `json:"hysteresis,omitempty"`
	Priority            int                                         `json:"priority"`
	ReactivePowerParams *ReportDERControlRequestReactivePowerParams `json:"reactivePowerParams,omitempty"`
	VoltageParams       *ReportDERControlRequestVoltageParams       `json:"voltageParams,omitempty"`
	YUnit               ReportDERControlRequestDERUnitEnum          `json:"yUnit"`
	ResponseTime        *float64                                    `json:"responseTime,omitempty"`
	StartTime           *string                                     `json:"startTime,omitempty"`
	Duration            *float64                                    `json:"duration,omitempty"`
	CustomData          *ReportDERControlRequestCustomData          `json:"customData,omitempty"`
}

type ReportDERControlRequestDERUnitEnum string

const (
	ReportDERControlRequestDERUnitEnumNotApplicable ReportDERControlRequestDERUnitEnum = "Not_Applicable"
	ReportDERControlRequestDERUnitEnumPctMaxW       ReportDERControlRequestDERUnitEnum = "PctMaxW"
	ReportDERControlRequestDERUnitEnumPctMaxVar     ReportDERControlRequestDERUnitEnum = "PctMaxVar"
	ReportDERControlRequestDERUnitEnumPctWAvail     ReportDERControlRequestDERUnitEnum = "PctWAvail"
	ReportDERControlRequestDERUnitEnumPctVarAvail   ReportDERControlRequestDERUnitEnum = "PctVarAvail"
	ReportDERControlRequestDERUnitEnumPctEffectiveV ReportDERControlRequestDERUnitEnum = "PctEffectiveV"
)

type ReportDERControlRequestVoltageParams struct {
	Hv10MinMeanValue     *float64                                         `json:"hv10MinMeanValue,omitempty"`
	Hv10MinMeanTripDelay *float64                                         `json:"hv10MinMeanTripDelay,omitempty"`
	PowerDuringCessation *ReportDERControlRequestPowerDuringCessationEnum `json:"powerDuringCessation,omitempty"`
	CustomData           *ReportDERControlRequestCustomData               `json:"customData,omitempty"`
}

type ReportDERControlRequestPowerDuringCessationEnum string

const (
	ReportDERControlRequestPowerDuringCessationEnumActive   ReportDERControlRequestPowerDuringCessationEnum = "Active"
	ReportDERControlRequestPowerDuringCessationEnumReactive ReportDERControlRequestPowerDuringCessationEnum = "Reactive"
)

type ReportDERControlRequestReactivePowerParams struct {
	VRef                       *float64                           `json:"vRef,omitempty"`
	AutonomousVRefEnable       *bool                              `json:"autonomousVRefEnable,omitempty"`
	AutonomousVRefTimeConstant *float64                           `json:"autonomousVRefTimeConstant,omitempty"`
	CustomData                 *ReportDERControlRequestCustomData `json:"customData,omitempty"`
}

type ReportDERControlRequestHysteresis struct {
	HysteresisHigh     *float64                           `json:"hysteresisHigh,omitempty"`
	HysteresisLow      *float64                           `json:"hysteresisLow,omitempty"`
	HysteresisDelay    *float64                           `json:"hysteresisDelay,omitempty"`
	HysteresisGradient *float64                           `json:"hysteresisGradient,omitempty"`
	CustomData         *ReportDERControlRequestCustomData `json:"customData,omitempty"`
}

type ReportDERControlRequestDERCurvePoints struct {
	X          float64                            `json:"x"`
	Y          float64                            `json:"y"`
	CustomData *ReportDERControlRequestCustomData `json:"customData,omitempty"`
}

type ReportDERControlRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (ReportDERControlRequest) ActionName() string { return "ReportDERControl" }

func (ReportDERControlRequest) Version() protocol.Version { return protocol.OCPP21 }

func (ReportDERControlRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (ReportDERControlRequest) SchemaName() string { return "ReportDERControlRequest.json" }

func (message ReportDERControlRequest) Validate() error {
	return validation.Validate("ReportDERControlRequest.json", schemaReportDERControlRequest, message)
}

func (ReportDERControlRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ReportDERControlRequest.json", schemaReportDERControlRequest, data)
}
