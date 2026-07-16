// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = GetCompositeScheduleConfirmation{}

var schemaGetCompositeScheduleConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}, "schedule": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "chargingSchedulePeriod": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "startPeriod": &validation.Schema{Type: "integer", AllowAdditional: true}, "limit": &validation.Schema{Type: "number", AllowAdditional: true}, "numberPhases": &validation.Schema{Type: "integer", AllowAdditional: true}, "phaseToUse": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"startPeriod", "limit"}}, MinItems: 1, HasMinItems: true}, "evseId": &validation.Schema{Type: "integer", AllowAdditional: true}, "duration": &validation.Schema{Type: "integer", AllowAdditional: true}, "scheduleStart": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "chargingRateUnit": &validation.Schema{Type: "string", Enum: []string{"W", "A"}}}, Required: []string{"evseId", "duration", "scheduleStart", "chargingRateUnit", "chargingSchedulePeriod"}}}, Required: []string{"status"}}

type GetCompositeScheduleConfirmation struct {
	CustomData *GetCompositeScheduleConfirmationCustomData        `json:"customData,omitempty"`
	Status     GetCompositeScheduleConfirmationGenericStatusEnum  `json:"status"`
	StatusInfo *GetCompositeScheduleConfirmationStatusInfo        `json:"statusInfo,omitempty"`
	Schedule   *GetCompositeScheduleConfirmationCompositeSchedule `json:"schedule,omitempty"`
}

type GetCompositeScheduleConfirmationCompositeSchedule struct {
	CustomData             *GetCompositeScheduleConfirmationCustomData              `json:"customData,omitempty"`
	ChargingSchedulePeriod []GetCompositeScheduleConfirmationChargingSchedulePeriod `json:"chargingSchedulePeriod"`
	EVSEID                 int                                                      `json:"evseId"`
	Duration               int                                                      `json:"duration"`
	ScheduleStart          string                                                   `json:"scheduleStart"`
	ChargingRateUnit       GetCompositeScheduleConfirmationChargingRateUnitEnum     `json:"chargingRateUnit"`
}

type GetCompositeScheduleConfirmationChargingRateUnitEnum string

const (
	GetCompositeScheduleConfirmationChargingRateUnitEnumW GetCompositeScheduleConfirmationChargingRateUnitEnum = "W"
	GetCompositeScheduleConfirmationChargingRateUnitEnumA GetCompositeScheduleConfirmationChargingRateUnitEnum = "A"
)

type GetCompositeScheduleConfirmationChargingSchedulePeriod struct {
	CustomData   *GetCompositeScheduleConfirmationCustomData `json:"customData,omitempty"`
	StartPeriod  int                                         `json:"startPeriod"`
	Limit        float64                                     `json:"limit"`
	NumberPhases *int                                        `json:"numberPhases,omitempty"`
	PhaseToUse   *int                                        `json:"phaseToUse,omitempty"`
}

type GetCompositeScheduleConfirmationStatusInfo struct {
	CustomData     *GetCompositeScheduleConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                                      `json:"reasonCode"`
	AdditionalInfo *string                                     `json:"additionalInfo,omitempty"`
}

type GetCompositeScheduleConfirmationGenericStatusEnum string

const (
	GetCompositeScheduleConfirmationGenericStatusEnumAccepted GetCompositeScheduleConfirmationGenericStatusEnum = "Accepted"
	GetCompositeScheduleConfirmationGenericStatusEnumRejected GetCompositeScheduleConfirmationGenericStatusEnum = "Rejected"
)

type GetCompositeScheduleConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (GetCompositeScheduleConfirmation) ActionName() string { return "GetCompositeSchedule" }

func (GetCompositeScheduleConfirmation) Version() protocol.Version { return protocol.OCPP201 }

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
