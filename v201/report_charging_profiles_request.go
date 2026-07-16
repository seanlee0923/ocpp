// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ReportChargingProfilesRequest{}

var schemaReportChargingProfilesRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "requestId": &validation.Schema{Type: "integer", AllowAdditional: true}, "chargingLimitSource": &validation.Schema{Type: "string", Enum: []string{"EMS", "Other", "SO", "CSO"}}, "chargingProfile": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "id": &validation.Schema{Type: "integer", AllowAdditional: true}, "stackLevel": &validation.Schema{Type: "integer", AllowAdditional: true}, "chargingProfilePurpose": &validation.Schema{Type: "string", Enum: []string{"ChargingStationExternalConstraints", "ChargingStationMaxProfile", "TxDefaultProfile", "TxProfile"}}, "chargingProfileKind": &validation.Schema{Type: "string", Enum: []string{"Absolute", "Recurring", "Relative"}}, "recurrencyKind": &validation.Schema{Type: "string", Enum: []string{"Daily", "Weekly"}}, "validFrom": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "validTo": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "chargingSchedule": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "id": &validation.Schema{Type: "integer", AllowAdditional: true}, "startSchedule": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "duration": &validation.Schema{Type: "integer", AllowAdditional: true}, "chargingRateUnit": &validation.Schema{Type: "string", Enum: []string{"W", "A"}}, "chargingSchedulePeriod": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "startPeriod": &validation.Schema{Type: "integer", AllowAdditional: true}, "limit": &validation.Schema{Type: "number", AllowAdditional: true}, "numberPhases": &validation.Schema{Type: "integer", AllowAdditional: true}, "phaseToUse": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"startPeriod", "limit"}}, MinItems: 1, HasMinItems: true, MaxItems: 1024, HasMaxItems: true}, "minChargingRate": &validation.Schema{Type: "number", AllowAdditional: true}, "salesTariff": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "id": &validation.Schema{Type: "integer", AllowAdditional: true}, "salesTariffDescription": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 32, HasMaxLength: true}, "numEPriceLevels": &validation.Schema{Type: "integer", AllowAdditional: true}, "salesTariffEntry": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "relativeTimeInterval": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "start": &validation.Schema{Type: "integer", AllowAdditional: true}, "duration": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"start"}}, "ePriceLevel": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "consumptionCost": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "startValue": &validation.Schema{Type: "number", AllowAdditional: true}, "cost": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "costKind": &validation.Schema{Type: "string", Enum: []string{"CarbonDioxideEmission", "RelativePricePercentage", "RenewableGenerationPercentage"}}, "amount": &validation.Schema{Type: "integer", AllowAdditional: true}, "amountMultiplier": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"costKind", "amount"}}, MinItems: 1, HasMinItems: true, MaxItems: 3, HasMaxItems: true}}, Required: []string{"startValue", "cost"}}, MinItems: 1, HasMinItems: true, MaxItems: 3, HasMaxItems: true}}, Required: []string{"relativeTimeInterval"}}, MinItems: 1, HasMinItems: true, MaxItems: 1024, HasMaxItems: true}}, Required: []string{"id", "salesTariffEntry"}}}, Required: []string{"id", "chargingRateUnit", "chargingSchedulePeriod"}}, MinItems: 1, HasMinItems: true, MaxItems: 3, HasMaxItems: true}, "transactionId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}}, Required: []string{"id", "stackLevel", "chargingProfilePurpose", "chargingProfileKind", "chargingSchedule"}}, MinItems: 1, HasMinItems: true}, "tbc": &validation.Schema{Type: "boolean", AllowAdditional: true}, "evseId": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"requestId", "chargingLimitSource", "evseId", "chargingProfile"}}

type ReportChargingProfilesRequest struct {
	CustomData          *ReportChargingProfilesRequestCustomData             `json:"customData,omitempty"`
	RequestID           int                                                  `json:"requestId"`
	ChargingLimitSource ReportChargingProfilesRequestChargingLimitSourceEnum `json:"chargingLimitSource"`
	ChargingProfile     []ReportChargingProfilesRequestChargingProfile       `json:"chargingProfile"`
	Tbc                 *bool                                                `json:"tbc,omitempty"`
	EVSEID              int                                                  `json:"evseId"`
}

type ReportChargingProfilesRequestChargingProfile struct {
	CustomData             *ReportChargingProfilesRequestCustomData                `json:"customData,omitempty"`
	ID                     int                                                     `json:"id"`
	StackLevel             int                                                     `json:"stackLevel"`
	ChargingProfilePurpose ReportChargingProfilesRequestChargingProfilePurposeEnum `json:"chargingProfilePurpose"`
	ChargingProfileKind    ReportChargingProfilesRequestChargingProfileKindEnum    `json:"chargingProfileKind"`
	RecurrencyKind         *ReportChargingProfilesRequestRecurrencyKindEnum        `json:"recurrencyKind,omitempty"`
	ValidFrom              *string                                                 `json:"validFrom,omitempty"`
	ValidTo                *string                                                 `json:"validTo,omitempty"`
	ChargingSchedule       []ReportChargingProfilesRequestChargingSchedule         `json:"chargingSchedule"`
	TransactionID          *string                                                 `json:"transactionId,omitempty"`
}

type ReportChargingProfilesRequestChargingSchedule struct {
	CustomData             *ReportChargingProfilesRequestCustomData              `json:"customData,omitempty"`
	ID                     int                                                   `json:"id"`
	StartSchedule          *string                                               `json:"startSchedule,omitempty"`
	Duration               *int                                                  `json:"duration,omitempty"`
	ChargingRateUnit       ReportChargingProfilesRequestChargingRateUnitEnum     `json:"chargingRateUnit"`
	ChargingSchedulePeriod []ReportChargingProfilesRequestChargingSchedulePeriod `json:"chargingSchedulePeriod"`
	MinChargingRate        *float64                                              `json:"minChargingRate,omitempty"`
	SalesTariff            *ReportChargingProfilesRequestSalesTariff             `json:"salesTariff,omitempty"`
}

type ReportChargingProfilesRequestSalesTariff struct {
	CustomData             *ReportChargingProfilesRequestCustomData        `json:"customData,omitempty"`
	ID                     int                                             `json:"id"`
	SalesTariffDescription *string                                         `json:"salesTariffDescription,omitempty"`
	NumEPriceLevels        *int                                            `json:"numEPriceLevels,omitempty"`
	SalesTariffEntry       []ReportChargingProfilesRequestSalesTariffEntry `json:"salesTariffEntry"`
}

type ReportChargingProfilesRequestSalesTariffEntry struct {
	CustomData           *ReportChargingProfilesRequestCustomData          `json:"customData,omitempty"`
	RelativeTimeInterval ReportChargingProfilesRequestRelativeTimeInterval `json:"relativeTimeInterval"`
	EPriceLevel          *int                                              `json:"ePriceLevel,omitempty"`
	ConsumptionCost      []ReportChargingProfilesRequestConsumptionCost    `json:"consumptionCost,omitempty"`
}

type ReportChargingProfilesRequestConsumptionCost struct {
	CustomData *ReportChargingProfilesRequestCustomData `json:"customData,omitempty"`
	StartValue float64                                  `json:"startValue"`
	Cost       []ReportChargingProfilesRequestCost      `json:"cost"`
}

type ReportChargingProfilesRequestCost struct {
	CustomData       *ReportChargingProfilesRequestCustomData  `json:"customData,omitempty"`
	CostKind         ReportChargingProfilesRequestCostKindEnum `json:"costKind"`
	Amount           int                                       `json:"amount"`
	AmountMultiplier *int                                      `json:"amountMultiplier,omitempty"`
}

type ReportChargingProfilesRequestCostKindEnum string

const (
	ReportChargingProfilesRequestCostKindEnumCarbonDioxideEmission         ReportChargingProfilesRequestCostKindEnum = "CarbonDioxideEmission"
	ReportChargingProfilesRequestCostKindEnumRelativePricePercentage       ReportChargingProfilesRequestCostKindEnum = "RelativePricePercentage"
	ReportChargingProfilesRequestCostKindEnumRenewableGenerationPercentage ReportChargingProfilesRequestCostKindEnum = "RenewableGenerationPercentage"
)

type ReportChargingProfilesRequestRelativeTimeInterval struct {
	CustomData *ReportChargingProfilesRequestCustomData `json:"customData,omitempty"`
	Start      int                                      `json:"start"`
	Duration   *int                                     `json:"duration,omitempty"`
}

type ReportChargingProfilesRequestChargingSchedulePeriod struct {
	CustomData   *ReportChargingProfilesRequestCustomData `json:"customData,omitempty"`
	StartPeriod  int                                      `json:"startPeriod"`
	Limit        float64                                  `json:"limit"`
	NumberPhases *int                                     `json:"numberPhases,omitempty"`
	PhaseToUse   *int                                     `json:"phaseToUse,omitempty"`
}

type ReportChargingProfilesRequestChargingRateUnitEnum string

const (
	ReportChargingProfilesRequestChargingRateUnitEnumW ReportChargingProfilesRequestChargingRateUnitEnum = "W"
	ReportChargingProfilesRequestChargingRateUnitEnumA ReportChargingProfilesRequestChargingRateUnitEnum = "A"
)

type ReportChargingProfilesRequestRecurrencyKindEnum string

const (
	ReportChargingProfilesRequestRecurrencyKindEnumDaily  ReportChargingProfilesRequestRecurrencyKindEnum = "Daily"
	ReportChargingProfilesRequestRecurrencyKindEnumWeekly ReportChargingProfilesRequestRecurrencyKindEnum = "Weekly"
)

type ReportChargingProfilesRequestChargingProfileKindEnum string

const (
	ReportChargingProfilesRequestChargingProfileKindEnumAbsolute  ReportChargingProfilesRequestChargingProfileKindEnum = "Absolute"
	ReportChargingProfilesRequestChargingProfileKindEnumRecurring ReportChargingProfilesRequestChargingProfileKindEnum = "Recurring"
	ReportChargingProfilesRequestChargingProfileKindEnumRelative  ReportChargingProfilesRequestChargingProfileKindEnum = "Relative"
)

type ReportChargingProfilesRequestChargingProfilePurposeEnum string

const (
	ReportChargingProfilesRequestChargingProfilePurposeEnumChargingStationExternalConstraints ReportChargingProfilesRequestChargingProfilePurposeEnum = "ChargingStationExternalConstraints"
	ReportChargingProfilesRequestChargingProfilePurposeEnumChargingStationMaxProfile          ReportChargingProfilesRequestChargingProfilePurposeEnum = "ChargingStationMaxProfile"
	ReportChargingProfilesRequestChargingProfilePurposeEnumTxDefaultProfile                   ReportChargingProfilesRequestChargingProfilePurposeEnum = "TxDefaultProfile"
	ReportChargingProfilesRequestChargingProfilePurposeEnumTxProfile                          ReportChargingProfilesRequestChargingProfilePurposeEnum = "TxProfile"
)

type ReportChargingProfilesRequestChargingLimitSourceEnum string

const (
	ReportChargingProfilesRequestChargingLimitSourceEnumEMS   ReportChargingProfilesRequestChargingLimitSourceEnum = "EMS"
	ReportChargingProfilesRequestChargingLimitSourceEnumOther ReportChargingProfilesRequestChargingLimitSourceEnum = "Other"
	ReportChargingProfilesRequestChargingLimitSourceEnumSO    ReportChargingProfilesRequestChargingLimitSourceEnum = "SO"
	ReportChargingProfilesRequestChargingLimitSourceEnumCSO   ReportChargingProfilesRequestChargingLimitSourceEnum = "CSO"
)

type ReportChargingProfilesRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (ReportChargingProfilesRequest) ActionName() string { return "ReportChargingProfiles" }

func (ReportChargingProfilesRequest) Version() protocol.Version { return protocol.OCPP201 }

func (ReportChargingProfilesRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (ReportChargingProfilesRequest) SchemaName() string { return "ReportChargingProfilesRequest.json" }

func (message ReportChargingProfilesRequest) Validate() error {
	return validation.Validate("ReportChargingProfilesRequest.json", schemaReportChargingProfilesRequest, message)
}

func (ReportChargingProfilesRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ReportChargingProfilesRequest.json", schemaReportChargingProfilesRequest, data)
}
