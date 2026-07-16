// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = SetChargingProfileRequest{}

var schemaSetChargingProfileRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"connectorId": &validation.Schema{Type: "integer", AllowAdditional: true}, "csChargingProfiles": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"chargingProfileId": &validation.Schema{Type: "integer", AllowAdditional: true}, "transactionId": &validation.Schema{Type: "integer", AllowAdditional: true}, "stackLevel": &validation.Schema{Type: "integer", AllowAdditional: true}, "chargingProfilePurpose": &validation.Schema{Type: "string", Enum: []string{"ChargePointMaxProfile", "TxDefaultProfile", "TxProfile"}}, "chargingProfileKind": &validation.Schema{Type: "string", Enum: []string{"Absolute", "Recurring", "Relative"}}, "recurrencyKind": &validation.Schema{Type: "string", Enum: []string{"Daily", "Weekly"}}, "validFrom": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "validTo": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "chargingSchedule": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"duration": &validation.Schema{Type: "integer", AllowAdditional: true}, "startSchedule": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "chargingRateUnit": &validation.Schema{Type: "string", Enum: []string{"A", "W"}}, "chargingSchedulePeriod": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"startPeriod": &validation.Schema{Type: "integer", AllowAdditional: true}, "limit": &validation.Schema{Type: "number", AllowAdditional: true, MultipleOf: 0.1}, "numberPhases": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"startPeriod", "limit"}}}, "minChargingRate": &validation.Schema{Type: "number", AllowAdditional: true, MultipleOf: 0.1}}, Required: []string{"chargingRateUnit", "chargingSchedulePeriod"}}}, Required: []string{"chargingProfileId", "stackLevel", "chargingProfilePurpose", "chargingProfileKind", "chargingSchedule"}}}, Required: []string{"connectorId", "csChargingProfiles"}}

type SetChargingProfileRequest struct {
	ConnectorID        int                                         `json:"connectorId"`
	CsChargingProfiles SetChargingProfileRequestCsChargingProfiles `json:"csChargingProfiles"`
}

type SetChargingProfileRequestCsChargingProfiles struct {
	ChargingProfileID      int                                                               `json:"chargingProfileId"`
	TransactionID          *int                                                              `json:"transactionId,omitempty"`
	StackLevel             int                                                               `json:"stackLevel"`
	ChargingProfilePurpose SetChargingProfileRequestCsChargingProfilesChargingProfilePurpose `json:"chargingProfilePurpose"`
	ChargingProfileKind    SetChargingProfileRequestCsChargingProfilesChargingProfileKind    `json:"chargingProfileKind"`
	RecurrencyKind         *SetChargingProfileRequestCsChargingProfilesRecurrencyKind        `json:"recurrencyKind,omitempty"`
	ValidFrom              *string                                                           `json:"validFrom,omitempty"`
	ValidTo                *string                                                           `json:"validTo,omitempty"`
	ChargingSchedule       SetChargingProfileRequestCsChargingProfilesChargingSchedule       `json:"chargingSchedule"`
}

type SetChargingProfileRequestCsChargingProfilesChargingSchedule struct {
	Duration               *int                                                                                    `json:"duration,omitempty"`
	StartSchedule          *string                                                                                 `json:"startSchedule,omitempty"`
	ChargingRateUnit       SetChargingProfileRequestCsChargingProfilesChargingScheduleChargingRateUnit             `json:"chargingRateUnit"`
	ChargingSchedulePeriod []SetChargingProfileRequestCsChargingProfilesChargingScheduleChargingSchedulePeriodItem `json:"chargingSchedulePeriod"`
	MinChargingRate        *float64                                                                                `json:"minChargingRate,omitempty"`
}

type SetChargingProfileRequestCsChargingProfilesChargingScheduleChargingSchedulePeriodItem struct {
	StartPeriod  int     `json:"startPeriod"`
	Limit        float64 `json:"limit"`
	NumberPhases *int    `json:"numberPhases,omitempty"`
}

type SetChargingProfileRequestCsChargingProfilesChargingScheduleChargingRateUnit string

const (
	SetChargingProfileRequestCsChargingProfilesChargingScheduleChargingRateUnitA SetChargingProfileRequestCsChargingProfilesChargingScheduleChargingRateUnit = "A"
	SetChargingProfileRequestCsChargingProfilesChargingScheduleChargingRateUnitW SetChargingProfileRequestCsChargingProfilesChargingScheduleChargingRateUnit = "W"
)

type SetChargingProfileRequestCsChargingProfilesRecurrencyKind string

const (
	SetChargingProfileRequestCsChargingProfilesRecurrencyKindDaily  SetChargingProfileRequestCsChargingProfilesRecurrencyKind = "Daily"
	SetChargingProfileRequestCsChargingProfilesRecurrencyKindWeekly SetChargingProfileRequestCsChargingProfilesRecurrencyKind = "Weekly"
)

type SetChargingProfileRequestCsChargingProfilesChargingProfileKind string

const (
	SetChargingProfileRequestCsChargingProfilesChargingProfileKindAbsolute  SetChargingProfileRequestCsChargingProfilesChargingProfileKind = "Absolute"
	SetChargingProfileRequestCsChargingProfilesChargingProfileKindRecurring SetChargingProfileRequestCsChargingProfilesChargingProfileKind = "Recurring"
	SetChargingProfileRequestCsChargingProfilesChargingProfileKindRelative  SetChargingProfileRequestCsChargingProfilesChargingProfileKind = "Relative"
)

type SetChargingProfileRequestCsChargingProfilesChargingProfilePurpose string

const (
	SetChargingProfileRequestCsChargingProfilesChargingProfilePurposeChargePointMaxProfile SetChargingProfileRequestCsChargingProfilesChargingProfilePurpose = "ChargePointMaxProfile"
	SetChargingProfileRequestCsChargingProfilesChargingProfilePurposeTxDefaultProfile      SetChargingProfileRequestCsChargingProfilesChargingProfilePurpose = "TxDefaultProfile"
	SetChargingProfileRequestCsChargingProfilesChargingProfilePurposeTxProfile             SetChargingProfileRequestCsChargingProfilesChargingProfilePurpose = "TxProfile"
)

func (SetChargingProfileRequest) ActionName() string { return "SetChargingProfile" }

func (SetChargingProfileRequest) Version() protocol.Version { return protocol.OCPP16 }

func (SetChargingProfileRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (SetChargingProfileRequest) SchemaName() string { return "SetChargingProfile.json" }

func (message SetChargingProfileRequest) Validate() error {
	return validation.Validate("SetChargingProfile.json", schemaSetChargingProfileRequest, message)
}

func (SetChargingProfileRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("SetChargingProfile.json", schemaSetChargingProfileRequest, data)
}
