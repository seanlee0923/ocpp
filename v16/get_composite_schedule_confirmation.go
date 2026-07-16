// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = GetCompositeScheduleConfirmation{}

var schemaGetCompositeScheduleConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected"}}, "connectorId": &validation.Schema{Type: "integer", AllowAdditional: true}, "scheduleStart": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "chargingSchedule": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"duration": &validation.Schema{Type: "integer", AllowAdditional: true}, "startSchedule": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "chargingRateUnit": &validation.Schema{Type: "string", Enum: []string{"A", "W"}}, "chargingSchedulePeriod": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"startPeriod": &validation.Schema{Type: "integer", AllowAdditional: true}, "limit": &validation.Schema{Type: "number", AllowAdditional: true, MultipleOf: 0.1}, "numberPhases": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"startPeriod", "limit"}}}, "minChargingRate": &validation.Schema{Type: "number", AllowAdditional: true, MultipleOf: 0.1}}, Required: []string{"chargingRateUnit", "chargingSchedulePeriod"}}}, Required: []string{"status"}}

type GetCompositeScheduleConfirmation struct {
	Status           GetCompositeScheduleConfirmationStatus            `json:"status"`
	ConnectorID      *int                                              `json:"connectorId,omitempty"`
	ScheduleStart    *string                                           `json:"scheduleStart,omitempty"`
	ChargingSchedule *GetCompositeScheduleConfirmationChargingSchedule `json:"chargingSchedule,omitempty"`
}

type GetCompositeScheduleConfirmationChargingSchedule struct {
	Duration               *int                                                                         `json:"duration,omitempty"`
	StartSchedule          *string                                                                      `json:"startSchedule,omitempty"`
	ChargingRateUnit       GetCompositeScheduleConfirmationChargingScheduleChargingRateUnit             `json:"chargingRateUnit"`
	ChargingSchedulePeriod []GetCompositeScheduleConfirmationChargingScheduleChargingSchedulePeriodItem `json:"chargingSchedulePeriod"`
	MinChargingRate        *float64                                                                     `json:"minChargingRate,omitempty"`
}

type GetCompositeScheduleConfirmationChargingScheduleChargingSchedulePeriodItem struct {
	StartPeriod  int     `json:"startPeriod"`
	Limit        float64 `json:"limit"`
	NumberPhases *int    `json:"numberPhases,omitempty"`
}

type GetCompositeScheduleConfirmationChargingScheduleChargingRateUnit string

const (
	GetCompositeScheduleConfirmationChargingScheduleChargingRateUnitA GetCompositeScheduleConfirmationChargingScheduleChargingRateUnit = "A"
	GetCompositeScheduleConfirmationChargingScheduleChargingRateUnitW GetCompositeScheduleConfirmationChargingScheduleChargingRateUnit = "W"
)

type GetCompositeScheduleConfirmationStatus string

const (
	GetCompositeScheduleConfirmationStatusAccepted GetCompositeScheduleConfirmationStatus = "Accepted"
	GetCompositeScheduleConfirmationStatusRejected GetCompositeScheduleConfirmationStatus = "Rejected"
)

func (GetCompositeScheduleConfirmation) ActionName() string { return "GetCompositeSchedule" }

func (GetCompositeScheduleConfirmation) Version() protocol.Version { return protocol.OCPP16 }

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
