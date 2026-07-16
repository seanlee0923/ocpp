// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = RemoteStartTransactionRequest{}

var schemaRemoteStartTransactionRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"connectorId": &validation.Schema{Type: "integer", AllowAdditional: true}, "idTag": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "chargingProfile": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"chargingProfileId": &validation.Schema{Type: "integer", AllowAdditional: true}, "transactionId": &validation.Schema{Type: "integer", AllowAdditional: true}, "stackLevel": &validation.Schema{Type: "integer", AllowAdditional: true}, "chargingProfilePurpose": &validation.Schema{Type: "string", Enum: []string{"ChargePointMaxProfile", "TxDefaultProfile", "TxProfile"}}, "chargingProfileKind": &validation.Schema{Type: "string", Enum: []string{"Absolute", "Recurring", "Relative"}}, "recurrencyKind": &validation.Schema{Type: "string", Enum: []string{"Daily", "Weekly"}}, "validFrom": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "validTo": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "chargingSchedule": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"duration": &validation.Schema{Type: "integer", AllowAdditional: true}, "startSchedule": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "chargingRateUnit": &validation.Schema{Type: "string", Enum: []string{"A", "W"}}, "chargingSchedulePeriod": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"startPeriod": &validation.Schema{Type: "integer", AllowAdditional: true}, "limit": &validation.Schema{Type: "number", AllowAdditional: true, MultipleOf: 0.1}, "numberPhases": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"startPeriod", "limit"}}}, "minChargingRate": &validation.Schema{Type: "number", AllowAdditional: true, MultipleOf: 0.1}}, Required: []string{"chargingRateUnit", "chargingSchedulePeriod"}}}, Required: []string{"chargingProfileId", "stackLevel", "chargingProfilePurpose", "chargingProfileKind", "chargingSchedule"}}}, Required: []string{"idTag"}}

type RemoteStartTransactionRequest struct {
	ConnectorID     *int                                          `json:"connectorId,omitempty"`
	IDTag           string                                        `json:"idTag"`
	ChargingProfile *RemoteStartTransactionRequestChargingProfile `json:"chargingProfile,omitempty"`
}

type RemoteStartTransactionRequestChargingProfile struct {
	ChargingProfileID      int                                                                `json:"chargingProfileId"`
	TransactionID          *int                                                               `json:"transactionId,omitempty"`
	StackLevel             int                                                                `json:"stackLevel"`
	ChargingProfilePurpose RemoteStartTransactionRequestChargingProfileChargingProfilePurpose `json:"chargingProfilePurpose"`
	ChargingProfileKind    RemoteStartTransactionRequestChargingProfileChargingProfileKind    `json:"chargingProfileKind"`
	RecurrencyKind         *RemoteStartTransactionRequestChargingProfileRecurrencyKind        `json:"recurrencyKind,omitempty"`
	ValidFrom              *string                                                            `json:"validFrom,omitempty"`
	ValidTo                *string                                                            `json:"validTo,omitempty"`
	ChargingSchedule       RemoteStartTransactionRequestChargingProfileChargingSchedule       `json:"chargingSchedule"`
}

type RemoteStartTransactionRequestChargingProfileChargingSchedule struct {
	Duration               *int                                                                                     `json:"duration,omitempty"`
	StartSchedule          *string                                                                                  `json:"startSchedule,omitempty"`
	ChargingRateUnit       RemoteStartTransactionRequestChargingProfileChargingScheduleChargingRateUnit             `json:"chargingRateUnit"`
	ChargingSchedulePeriod []RemoteStartTransactionRequestChargingProfileChargingScheduleChargingSchedulePeriodItem `json:"chargingSchedulePeriod"`
	MinChargingRate        *float64                                                                                 `json:"minChargingRate,omitempty"`
}

type RemoteStartTransactionRequestChargingProfileChargingScheduleChargingSchedulePeriodItem struct {
	StartPeriod  int     `json:"startPeriod"`
	Limit        float64 `json:"limit"`
	NumberPhases *int    `json:"numberPhases,omitempty"`
}

type RemoteStartTransactionRequestChargingProfileChargingScheduleChargingRateUnit string

const (
	RemoteStartTransactionRequestChargingProfileChargingScheduleChargingRateUnitA RemoteStartTransactionRequestChargingProfileChargingScheduleChargingRateUnit = "A"
	RemoteStartTransactionRequestChargingProfileChargingScheduleChargingRateUnitW RemoteStartTransactionRequestChargingProfileChargingScheduleChargingRateUnit = "W"
)

type RemoteStartTransactionRequestChargingProfileRecurrencyKind string

const (
	RemoteStartTransactionRequestChargingProfileRecurrencyKindDaily  RemoteStartTransactionRequestChargingProfileRecurrencyKind = "Daily"
	RemoteStartTransactionRequestChargingProfileRecurrencyKindWeekly RemoteStartTransactionRequestChargingProfileRecurrencyKind = "Weekly"
)

type RemoteStartTransactionRequestChargingProfileChargingProfileKind string

const (
	RemoteStartTransactionRequestChargingProfileChargingProfileKindAbsolute  RemoteStartTransactionRequestChargingProfileChargingProfileKind = "Absolute"
	RemoteStartTransactionRequestChargingProfileChargingProfileKindRecurring RemoteStartTransactionRequestChargingProfileChargingProfileKind = "Recurring"
	RemoteStartTransactionRequestChargingProfileChargingProfileKindRelative  RemoteStartTransactionRequestChargingProfileChargingProfileKind = "Relative"
)

type RemoteStartTransactionRequestChargingProfileChargingProfilePurpose string

const (
	RemoteStartTransactionRequestChargingProfileChargingProfilePurposeChargePointMaxProfile RemoteStartTransactionRequestChargingProfileChargingProfilePurpose = "ChargePointMaxProfile"
	RemoteStartTransactionRequestChargingProfileChargingProfilePurposeTxDefaultProfile      RemoteStartTransactionRequestChargingProfileChargingProfilePurpose = "TxDefaultProfile"
	RemoteStartTransactionRequestChargingProfileChargingProfilePurposeTxProfile             RemoteStartTransactionRequestChargingProfileChargingProfilePurpose = "TxProfile"
)

func (RemoteStartTransactionRequest) ActionName() string { return "RemoteStartTransaction" }

func (RemoteStartTransactionRequest) Version() protocol.Version { return protocol.OCPP16 }

func (RemoteStartTransactionRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (RemoteStartTransactionRequest) SchemaName() string { return "RemoteStartTransaction.json" }

func (message RemoteStartTransactionRequest) Validate() error {
	return validation.Validate("RemoteStartTransaction.json", schemaRemoteStartTransactionRequest, message)
}

func (RemoteStartTransactionRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("RemoteStartTransaction.json", schemaRemoteStartTransactionRequest, data)
}
