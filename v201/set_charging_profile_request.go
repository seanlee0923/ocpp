// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = SetChargingProfileRequest{}

var schemaSetChargingProfileRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "evseId": &validation.Schema{Type: "integer", AllowAdditional: true}, "chargingProfile": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "id": &validation.Schema{Type: "integer", AllowAdditional: true}, "stackLevel": &validation.Schema{Type: "integer", AllowAdditional: true}, "chargingProfilePurpose": &validation.Schema{Type: "string", Enum: []string{"ChargingStationExternalConstraints", "ChargingStationMaxProfile", "TxDefaultProfile", "TxProfile"}}, "chargingProfileKind": &validation.Schema{Type: "string", Enum: []string{"Absolute", "Recurring", "Relative"}}, "recurrencyKind": &validation.Schema{Type: "string", Enum: []string{"Daily", "Weekly"}}, "validFrom": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "validTo": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "chargingSchedule": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "id": &validation.Schema{Type: "integer", AllowAdditional: true}, "startSchedule": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "duration": &validation.Schema{Type: "integer", AllowAdditional: true}, "chargingRateUnit": &validation.Schema{Type: "string", Enum: []string{"W", "A"}}, "chargingSchedulePeriod": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "startPeriod": &validation.Schema{Type: "integer", AllowAdditional: true}, "limit": &validation.Schema{Type: "number", AllowAdditional: true}, "numberPhases": &validation.Schema{Type: "integer", AllowAdditional: true}, "phaseToUse": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"startPeriod", "limit"}}, MinItems: 1, HasMinItems: true, MaxItems: 1024, HasMaxItems: true}, "minChargingRate": &validation.Schema{Type: "number", AllowAdditional: true}, "salesTariff": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "id": &validation.Schema{Type: "integer", AllowAdditional: true}, "salesTariffDescription": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 32, HasMaxLength: true}, "numEPriceLevels": &validation.Schema{Type: "integer", AllowAdditional: true}, "salesTariffEntry": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "relativeTimeInterval": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "start": &validation.Schema{Type: "integer", AllowAdditional: true}, "duration": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"start"}}, "ePriceLevel": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "consumptionCost": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "startValue": &validation.Schema{Type: "number", AllowAdditional: true}, "cost": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "costKind": &validation.Schema{Type: "string", Enum: []string{"CarbonDioxideEmission", "RelativePricePercentage", "RenewableGenerationPercentage"}}, "amount": &validation.Schema{Type: "integer", AllowAdditional: true}, "amountMultiplier": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"costKind", "amount"}}, MinItems: 1, HasMinItems: true, MaxItems: 3, HasMaxItems: true}}, Required: []string{"startValue", "cost"}}, MinItems: 1, HasMinItems: true, MaxItems: 3, HasMaxItems: true}}, Required: []string{"relativeTimeInterval"}}, MinItems: 1, HasMinItems: true, MaxItems: 1024, HasMaxItems: true}}, Required: []string{"id", "salesTariffEntry"}}}, Required: []string{"id", "chargingRateUnit", "chargingSchedulePeriod"}}, MinItems: 1, HasMinItems: true, MaxItems: 3, HasMaxItems: true}, "transactionId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}}, Required: []string{"id", "stackLevel", "chargingProfilePurpose", "chargingProfileKind", "chargingSchedule"}}}, Required: []string{"evseId", "chargingProfile"}}

type SetChargingProfileRequest struct {
	CustomData      *SetChargingProfileRequestCustomData     `json:"customData,omitempty"`
	EVSEID          int                                      `json:"evseId"`
	ChargingProfile SetChargingProfileRequestChargingProfile `json:"chargingProfile"`
}

type SetChargingProfileRequestChargingProfile struct {
	CustomData             *SetChargingProfileRequestCustomData                `json:"customData,omitempty"`
	ID                     int                                                 `json:"id"`
	StackLevel             int                                                 `json:"stackLevel"`
	ChargingProfilePurpose SetChargingProfileRequestChargingProfilePurposeEnum `json:"chargingProfilePurpose"`
	ChargingProfileKind    SetChargingProfileRequestChargingProfileKindEnum    `json:"chargingProfileKind"`
	RecurrencyKind         *SetChargingProfileRequestRecurrencyKindEnum        `json:"recurrencyKind,omitempty"`
	ValidFrom              *string                                             `json:"validFrom,omitempty"`
	ValidTo                *string                                             `json:"validTo,omitempty"`
	ChargingSchedule       []SetChargingProfileRequestChargingSchedule         `json:"chargingSchedule"`
	TransactionID          *string                                             `json:"transactionId,omitempty"`
}

type SetChargingProfileRequestChargingSchedule struct {
	CustomData             *SetChargingProfileRequestCustomData              `json:"customData,omitempty"`
	ID                     int                                               `json:"id"`
	StartSchedule          *string                                           `json:"startSchedule,omitempty"`
	Duration               *int                                              `json:"duration,omitempty"`
	ChargingRateUnit       SetChargingProfileRequestChargingRateUnitEnum     `json:"chargingRateUnit"`
	ChargingSchedulePeriod []SetChargingProfileRequestChargingSchedulePeriod `json:"chargingSchedulePeriod"`
	MinChargingRate        *float64                                          `json:"minChargingRate,omitempty"`
	SalesTariff            *SetChargingProfileRequestSalesTariff             `json:"salesTariff,omitempty"`
}

type SetChargingProfileRequestSalesTariff struct {
	CustomData             *SetChargingProfileRequestCustomData        `json:"customData,omitempty"`
	ID                     int                                         `json:"id"`
	SalesTariffDescription *string                                     `json:"salesTariffDescription,omitempty"`
	NumEPriceLevels        *int                                        `json:"numEPriceLevels,omitempty"`
	SalesTariffEntry       []SetChargingProfileRequestSalesTariffEntry `json:"salesTariffEntry"`
}

type SetChargingProfileRequestSalesTariffEntry struct {
	CustomData           *SetChargingProfileRequestCustomData          `json:"customData,omitempty"`
	RelativeTimeInterval SetChargingProfileRequestRelativeTimeInterval `json:"relativeTimeInterval"`
	EPriceLevel          *int                                          `json:"ePriceLevel,omitempty"`
	ConsumptionCost      []SetChargingProfileRequestConsumptionCost    `json:"consumptionCost,omitempty"`
}

type SetChargingProfileRequestConsumptionCost struct {
	CustomData *SetChargingProfileRequestCustomData `json:"customData,omitempty"`
	StartValue float64                              `json:"startValue"`
	Cost       []SetChargingProfileRequestCost      `json:"cost"`
}

type SetChargingProfileRequestCost struct {
	CustomData       *SetChargingProfileRequestCustomData  `json:"customData,omitempty"`
	CostKind         SetChargingProfileRequestCostKindEnum `json:"costKind"`
	Amount           int                                   `json:"amount"`
	AmountMultiplier *int                                  `json:"amountMultiplier,omitempty"`
}

type SetChargingProfileRequestCostKindEnum string

const (
	SetChargingProfileRequestCostKindEnumCarbonDioxideEmission         SetChargingProfileRequestCostKindEnum = "CarbonDioxideEmission"
	SetChargingProfileRequestCostKindEnumRelativePricePercentage       SetChargingProfileRequestCostKindEnum = "RelativePricePercentage"
	SetChargingProfileRequestCostKindEnumRenewableGenerationPercentage SetChargingProfileRequestCostKindEnum = "RenewableGenerationPercentage"
)

type SetChargingProfileRequestRelativeTimeInterval struct {
	CustomData *SetChargingProfileRequestCustomData `json:"customData,omitempty"`
	Start      int                                  `json:"start"`
	Duration   *int                                 `json:"duration,omitempty"`
}

type SetChargingProfileRequestChargingSchedulePeriod struct {
	CustomData   *SetChargingProfileRequestCustomData `json:"customData,omitempty"`
	StartPeriod  int                                  `json:"startPeriod"`
	Limit        float64                              `json:"limit"`
	NumberPhases *int                                 `json:"numberPhases,omitempty"`
	PhaseToUse   *int                                 `json:"phaseToUse,omitempty"`
}

type SetChargingProfileRequestChargingRateUnitEnum string

const (
	SetChargingProfileRequestChargingRateUnitEnumW SetChargingProfileRequestChargingRateUnitEnum = "W"
	SetChargingProfileRequestChargingRateUnitEnumA SetChargingProfileRequestChargingRateUnitEnum = "A"
)

type SetChargingProfileRequestRecurrencyKindEnum string

const (
	SetChargingProfileRequestRecurrencyKindEnumDaily  SetChargingProfileRequestRecurrencyKindEnum = "Daily"
	SetChargingProfileRequestRecurrencyKindEnumWeekly SetChargingProfileRequestRecurrencyKindEnum = "Weekly"
)

type SetChargingProfileRequestChargingProfileKindEnum string

const (
	SetChargingProfileRequestChargingProfileKindEnumAbsolute  SetChargingProfileRequestChargingProfileKindEnum = "Absolute"
	SetChargingProfileRequestChargingProfileKindEnumRecurring SetChargingProfileRequestChargingProfileKindEnum = "Recurring"
	SetChargingProfileRequestChargingProfileKindEnumRelative  SetChargingProfileRequestChargingProfileKindEnum = "Relative"
)

type SetChargingProfileRequestChargingProfilePurposeEnum string

const (
	SetChargingProfileRequestChargingProfilePurposeEnumChargingStationExternalConstraints SetChargingProfileRequestChargingProfilePurposeEnum = "ChargingStationExternalConstraints"
	SetChargingProfileRequestChargingProfilePurposeEnumChargingStationMaxProfile          SetChargingProfileRequestChargingProfilePurposeEnum = "ChargingStationMaxProfile"
	SetChargingProfileRequestChargingProfilePurposeEnumTxDefaultProfile                   SetChargingProfileRequestChargingProfilePurposeEnum = "TxDefaultProfile"
	SetChargingProfileRequestChargingProfilePurposeEnumTxProfile                          SetChargingProfileRequestChargingProfilePurposeEnum = "TxProfile"
)

type SetChargingProfileRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (SetChargingProfileRequest) ActionName() string { return "SetChargingProfile" }

func (SetChargingProfileRequest) Version() protocol.Version { return protocol.OCPP201 }

func (SetChargingProfileRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (SetChargingProfileRequest) SchemaName() string { return "SetChargingProfileRequest.json" }

func (message SetChargingProfileRequest) Validate() error {
	return validation.Validate("SetChargingProfileRequest.json", schemaSetChargingProfileRequest, message)
}

func (SetChargingProfileRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("SetChargingProfileRequest.json", schemaSetChargingProfileRequest, data)
}
