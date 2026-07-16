// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = GetCompositeScheduleConfirmation{}

var schemaGetCompositeScheduleConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "schedule": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"evseId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "duration": &validation.Schema{Type: "integer", AllowAdditional: true}, "scheduleStart": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "chargingRateUnit": &validation.Schema{Type: "string", Enum: []string{"W", "A"}}, "chargingSchedulePeriod": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"startPeriod": &validation.Schema{Type: "integer", AllowAdditional: true}, "limit": &validation.Schema{Type: "number", AllowAdditional: true}, "limit_L2": &validation.Schema{Type: "number", AllowAdditional: true}, "limit_L3": &validation.Schema{Type: "number", AllowAdditional: true}, "numberPhases": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true, Maximum: 3, HasMaximum: true}, "phaseToUse": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true, Maximum: 3, HasMaximum: true}, "dischargeLimit": &validation.Schema{Type: "number", AllowAdditional: true, Maximum: 0, HasMaximum: true}, "dischargeLimit_L2": &validation.Schema{Type: "number", AllowAdditional: true, Maximum: 0, HasMaximum: true}, "dischargeLimit_L3": &validation.Schema{Type: "number", AllowAdditional: true, Maximum: 0, HasMaximum: true}, "setpoint": &validation.Schema{Type: "number", AllowAdditional: true}, "setpoint_L2": &validation.Schema{Type: "number", AllowAdditional: true}, "setpoint_L3": &validation.Schema{Type: "number", AllowAdditional: true}, "setpointReactive": &validation.Schema{Type: "number", AllowAdditional: true}, "setpointReactive_L2": &validation.Schema{Type: "number", AllowAdditional: true}, "setpointReactive_L3": &validation.Schema{Type: "number", AllowAdditional: true}, "preconditioningRequest": &validation.Schema{Type: "boolean", AllowAdditional: true}, "evseSleep": &validation.Schema{Type: "boolean", AllowAdditional: true}, "v2xBaseline": &validation.Schema{Type: "number", AllowAdditional: true}, "operationMode": &validation.Schema{Type: "string", Enum: []string{"Idle", "ChargingOnly", "CentralSetpoint", "ExternalSetpoint", "ExternalLimits", "CentralFrequency", "LocalFrequency", "LocalLoadBalancing"}}, "v2xFreqWattCurve": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"frequency": &validation.Schema{Type: "number", AllowAdditional: true}, "power": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"frequency", "power"}}, MinItems: 1, HasMinItems: true, MaxItems: 20, HasMaxItems: true}, "v2xSignalWattCurve": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"signal": &validation.Schema{Type: "integer", AllowAdditional: true}, "power": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"signal", "power"}}, MinItems: 1, HasMinItems: true, MaxItems: 20, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"startPeriod"}}, MinItems: 1, HasMinItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"evseId", "duration", "scheduleStart", "chargingRateUnit", "chargingSchedulePeriod"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type GetCompositeScheduleConfirmation struct {
	Status     GetCompositeScheduleConfirmationGenericStatusEnum  `json:"status"`
	StatusInfo *GetCompositeScheduleConfirmationStatusInfo        `json:"statusInfo,omitempty"`
	Schedule   *GetCompositeScheduleConfirmationCompositeSchedule `json:"schedule,omitempty"`
	CustomData *GetCompositeScheduleConfirmationCustomData        `json:"customData,omitempty"`
}

type GetCompositeScheduleConfirmationCompositeSchedule struct {
	EVSEID                 int                                                      `json:"evseId"`
	Duration               int                                                      `json:"duration"`
	ScheduleStart          string                                                   `json:"scheduleStart"`
	ChargingRateUnit       GetCompositeScheduleConfirmationChargingRateUnitEnum     `json:"chargingRateUnit"`
	ChargingSchedulePeriod []GetCompositeScheduleConfirmationChargingSchedulePeriod `json:"chargingSchedulePeriod"`
	CustomData             *GetCompositeScheduleConfirmationCustomData              `json:"customData,omitempty"`
}

type GetCompositeScheduleConfirmationChargingSchedulePeriod struct {
	StartPeriod            int                                                  `json:"startPeriod"`
	Limit                  *float64                                             `json:"limit,omitempty"`
	LimitL2                *float64                                             `json:"limit_L2,omitempty"`
	LimitL3                *float64                                             `json:"limit_L3,omitempty"`
	NumberPhases           *int                                                 `json:"numberPhases,omitempty"`
	PhaseToUse             *int                                                 `json:"phaseToUse,omitempty"`
	DischargeLimit         *float64                                             `json:"dischargeLimit,omitempty"`
	DischargeLimitL2       *float64                                             `json:"dischargeLimit_L2,omitempty"`
	DischargeLimitL3       *float64                                             `json:"dischargeLimit_L3,omitempty"`
	Setpoint               *float64                                             `json:"setpoint,omitempty"`
	SetpointL2             *float64                                             `json:"setpoint_L2,omitempty"`
	SetpointL3             *float64                                             `json:"setpoint_L3,omitempty"`
	SetpointReactive       *float64                                             `json:"setpointReactive,omitempty"`
	SetpointReactiveL2     *float64                                             `json:"setpointReactive_L2,omitempty"`
	SetpointReactiveL3     *float64                                             `json:"setpointReactive_L3,omitempty"`
	PreconditioningRequest *bool                                                `json:"preconditioningRequest,omitempty"`
	EVSESleep              *bool                                                `json:"evseSleep,omitempty"`
	V2XBaseline            *float64                                             `json:"v2xBaseline,omitempty"`
	OperationMode          *GetCompositeScheduleConfirmationOperationModeEnum   `json:"operationMode,omitempty"`
	V2XFreqWattCurve       []GetCompositeScheduleConfirmationV2XFreqWattPoint   `json:"v2xFreqWattCurve,omitempty"`
	V2XSignalWattCurve     []GetCompositeScheduleConfirmationV2XSignalWattPoint `json:"v2xSignalWattCurve,omitempty"`
	CustomData             *GetCompositeScheduleConfirmationCustomData          `json:"customData,omitempty"`
}

type GetCompositeScheduleConfirmationV2XSignalWattPoint struct {
	Signal     int                                         `json:"signal"`
	Power      float64                                     `json:"power"`
	CustomData *GetCompositeScheduleConfirmationCustomData `json:"customData,omitempty"`
}

type GetCompositeScheduleConfirmationV2XFreqWattPoint struct {
	Frequency  float64                                     `json:"frequency"`
	Power      float64                                     `json:"power"`
	CustomData *GetCompositeScheduleConfirmationCustomData `json:"customData,omitempty"`
}

type GetCompositeScheduleConfirmationOperationModeEnum string

const (
	GetCompositeScheduleConfirmationOperationModeEnumIdle               GetCompositeScheduleConfirmationOperationModeEnum = "Idle"
	GetCompositeScheduleConfirmationOperationModeEnumChargingOnly       GetCompositeScheduleConfirmationOperationModeEnum = "ChargingOnly"
	GetCompositeScheduleConfirmationOperationModeEnumCentralSetpoint    GetCompositeScheduleConfirmationOperationModeEnum = "CentralSetpoint"
	GetCompositeScheduleConfirmationOperationModeEnumExternalSetpoint   GetCompositeScheduleConfirmationOperationModeEnum = "ExternalSetpoint"
	GetCompositeScheduleConfirmationOperationModeEnumExternalLimits     GetCompositeScheduleConfirmationOperationModeEnum = "ExternalLimits"
	GetCompositeScheduleConfirmationOperationModeEnumCentralFrequency   GetCompositeScheduleConfirmationOperationModeEnum = "CentralFrequency"
	GetCompositeScheduleConfirmationOperationModeEnumLocalFrequency     GetCompositeScheduleConfirmationOperationModeEnum = "LocalFrequency"
	GetCompositeScheduleConfirmationOperationModeEnumLocalLoadBalancing GetCompositeScheduleConfirmationOperationModeEnum = "LocalLoadBalancing"
)

type GetCompositeScheduleConfirmationChargingRateUnitEnum string

const (
	GetCompositeScheduleConfirmationChargingRateUnitEnumW GetCompositeScheduleConfirmationChargingRateUnitEnum = "W"
	GetCompositeScheduleConfirmationChargingRateUnitEnumA GetCompositeScheduleConfirmationChargingRateUnitEnum = "A"
)

type GetCompositeScheduleConfirmationStatusInfo struct {
	ReasonCode     string                                      `json:"reasonCode"`
	AdditionalInfo *string                                     `json:"additionalInfo,omitempty"`
	CustomData     *GetCompositeScheduleConfirmationCustomData `json:"customData,omitempty"`
}

type GetCompositeScheduleConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value GetCompositeScheduleConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *GetCompositeScheduleConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

type GetCompositeScheduleConfirmationGenericStatusEnum string

const (
	GetCompositeScheduleConfirmationGenericStatusEnumAccepted GetCompositeScheduleConfirmationGenericStatusEnum = "Accepted"
	GetCompositeScheduleConfirmationGenericStatusEnumRejected GetCompositeScheduleConfirmationGenericStatusEnum = "Rejected"
)

func (GetCompositeScheduleConfirmation) ActionName() string { return "GetCompositeSchedule" }

func (GetCompositeScheduleConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (GetCompositeScheduleConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (GetCompositeScheduleConfirmation) SchemaName() string {
	return "GetCompositeScheduleResponse.json"
}

func (message GetCompositeScheduleConfirmation) Validate() error {
	return validation.Validate("GetCompositeScheduleResponse.json", schemaGetCompositeScheduleConfirmation, message)
}

func (GetCompositeScheduleConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetCompositeScheduleResponse.json", schemaGetCompositeScheduleConfirmation, data)
}
