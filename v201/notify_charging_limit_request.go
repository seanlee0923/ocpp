// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = NotifyChargingLimitRequest{}

var schemaNotifyChargingLimitRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "chargingSchedule": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "id": &validation.Schema{Type: "integer", AllowAdditional: true}, "startSchedule": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "duration": &validation.Schema{Type: "integer", AllowAdditional: true}, "chargingRateUnit": &validation.Schema{Type: "string", Enum: []string{"W", "A"}}, "chargingSchedulePeriod": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "startPeriod": &validation.Schema{Type: "integer", AllowAdditional: true}, "limit": &validation.Schema{Type: "number", AllowAdditional: true}, "numberPhases": &validation.Schema{Type: "integer", AllowAdditional: true}, "phaseToUse": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"startPeriod", "limit"}}, MinItems: 1, HasMinItems: true, MaxItems: 1024, HasMaxItems: true}, "minChargingRate": &validation.Schema{Type: "number", AllowAdditional: true}, "salesTariff": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "id": &validation.Schema{Type: "integer", AllowAdditional: true}, "salesTariffDescription": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 32, HasMaxLength: true}, "numEPriceLevels": &validation.Schema{Type: "integer", AllowAdditional: true}, "salesTariffEntry": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "relativeTimeInterval": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "start": &validation.Schema{Type: "integer", AllowAdditional: true}, "duration": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"start"}}, "ePriceLevel": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "consumptionCost": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "startValue": &validation.Schema{Type: "number", AllowAdditional: true}, "cost": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "costKind": &validation.Schema{Type: "string", Enum: []string{"CarbonDioxideEmission", "RelativePricePercentage", "RenewableGenerationPercentage"}}, "amount": &validation.Schema{Type: "integer", AllowAdditional: true}, "amountMultiplier": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"costKind", "amount"}}, MinItems: 1, HasMinItems: true, MaxItems: 3, HasMaxItems: true}}, Required: []string{"startValue", "cost"}}, MinItems: 1, HasMinItems: true, MaxItems: 3, HasMaxItems: true}}, Required: []string{"relativeTimeInterval"}}, MinItems: 1, HasMinItems: true, MaxItems: 1024, HasMaxItems: true}}, Required: []string{"id", "salesTariffEntry"}}}, Required: []string{"id", "chargingRateUnit", "chargingSchedulePeriod"}}, MinItems: 1, HasMinItems: true}, "evseId": &validation.Schema{Type: "integer", AllowAdditional: true}, "chargingLimit": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "chargingLimitSource": &validation.Schema{Type: "string", Enum: []string{"EMS", "Other", "SO", "CSO"}}, "isGridCritical": &validation.Schema{Type: "boolean", AllowAdditional: true}}, Required: []string{"chargingLimitSource"}}}, Required: []string{"chargingLimit"}}

type NotifyChargingLimitRequest struct {
	CustomData       *NotifyChargingLimitRequestCustomData        `json:"customData,omitempty"`
	ChargingSchedule []NotifyChargingLimitRequestChargingSchedule `json:"chargingSchedule,omitempty"`
	EVSEID           *int                                         `json:"evseId,omitempty"`
	ChargingLimit    NotifyChargingLimitRequestChargingLimit      `json:"chargingLimit"`
}

type NotifyChargingLimitRequestChargingLimit struct {
	CustomData          *NotifyChargingLimitRequestCustomData             `json:"customData,omitempty"`
	ChargingLimitSource NotifyChargingLimitRequestChargingLimitSourceEnum `json:"chargingLimitSource"`
	IsGridCritical      *bool                                             `json:"isGridCritical,omitempty"`
}

type NotifyChargingLimitRequestChargingLimitSourceEnum string

const (
	NotifyChargingLimitRequestChargingLimitSourceEnumEMS   NotifyChargingLimitRequestChargingLimitSourceEnum = "EMS"
	NotifyChargingLimitRequestChargingLimitSourceEnumOther NotifyChargingLimitRequestChargingLimitSourceEnum = "Other"
	NotifyChargingLimitRequestChargingLimitSourceEnumSO    NotifyChargingLimitRequestChargingLimitSourceEnum = "SO"
	NotifyChargingLimitRequestChargingLimitSourceEnumCSO   NotifyChargingLimitRequestChargingLimitSourceEnum = "CSO"
)

type NotifyChargingLimitRequestChargingSchedule struct {
	CustomData             *NotifyChargingLimitRequestCustomData              `json:"customData,omitempty"`
	ID                     int                                                `json:"id"`
	StartSchedule          *string                                            `json:"startSchedule,omitempty"`
	Duration               *int                                               `json:"duration,omitempty"`
	ChargingRateUnit       NotifyChargingLimitRequestChargingRateUnitEnum     `json:"chargingRateUnit"`
	ChargingSchedulePeriod []NotifyChargingLimitRequestChargingSchedulePeriod `json:"chargingSchedulePeriod"`
	MinChargingRate        *float64                                           `json:"minChargingRate,omitempty"`
	SalesTariff            *NotifyChargingLimitRequestSalesTariff             `json:"salesTariff,omitempty"`
}

type NotifyChargingLimitRequestSalesTariff struct {
	CustomData             *NotifyChargingLimitRequestCustomData        `json:"customData,omitempty"`
	ID                     int                                          `json:"id"`
	SalesTariffDescription *string                                      `json:"salesTariffDescription,omitempty"`
	NumEPriceLevels        *int                                         `json:"numEPriceLevels,omitempty"`
	SalesTariffEntry       []NotifyChargingLimitRequestSalesTariffEntry `json:"salesTariffEntry"`
}

type NotifyChargingLimitRequestSalesTariffEntry struct {
	CustomData           *NotifyChargingLimitRequestCustomData          `json:"customData,omitempty"`
	RelativeTimeInterval NotifyChargingLimitRequestRelativeTimeInterval `json:"relativeTimeInterval"`
	EPriceLevel          *int                                           `json:"ePriceLevel,omitempty"`
	ConsumptionCost      []NotifyChargingLimitRequestConsumptionCost    `json:"consumptionCost,omitempty"`
}

type NotifyChargingLimitRequestConsumptionCost struct {
	CustomData *NotifyChargingLimitRequestCustomData `json:"customData,omitempty"`
	StartValue float64                               `json:"startValue"`
	Cost       []NotifyChargingLimitRequestCost      `json:"cost"`
}

type NotifyChargingLimitRequestCost struct {
	CustomData       *NotifyChargingLimitRequestCustomData  `json:"customData,omitempty"`
	CostKind         NotifyChargingLimitRequestCostKindEnum `json:"costKind"`
	Amount           int                                    `json:"amount"`
	AmountMultiplier *int                                   `json:"amountMultiplier,omitempty"`
}

type NotifyChargingLimitRequestCostKindEnum string

const (
	NotifyChargingLimitRequestCostKindEnumCarbonDioxideEmission         NotifyChargingLimitRequestCostKindEnum = "CarbonDioxideEmission"
	NotifyChargingLimitRequestCostKindEnumRelativePricePercentage       NotifyChargingLimitRequestCostKindEnum = "RelativePricePercentage"
	NotifyChargingLimitRequestCostKindEnumRenewableGenerationPercentage NotifyChargingLimitRequestCostKindEnum = "RenewableGenerationPercentage"
)

type NotifyChargingLimitRequestRelativeTimeInterval struct {
	CustomData *NotifyChargingLimitRequestCustomData `json:"customData,omitempty"`
	Start      int                                   `json:"start"`
	Duration   *int                                  `json:"duration,omitempty"`
}

type NotifyChargingLimitRequestChargingSchedulePeriod struct {
	CustomData   *NotifyChargingLimitRequestCustomData `json:"customData,omitempty"`
	StartPeriod  int                                   `json:"startPeriod"`
	Limit        float64                               `json:"limit"`
	NumberPhases *int                                  `json:"numberPhases,omitempty"`
	PhaseToUse   *int                                  `json:"phaseToUse,omitempty"`
}

type NotifyChargingLimitRequestChargingRateUnitEnum string

const (
	NotifyChargingLimitRequestChargingRateUnitEnumW NotifyChargingLimitRequestChargingRateUnitEnum = "W"
	NotifyChargingLimitRequestChargingRateUnitEnumA NotifyChargingLimitRequestChargingRateUnitEnum = "A"
)

type NotifyChargingLimitRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (NotifyChargingLimitRequest) ActionName() string { return "NotifyChargingLimit" }

func (NotifyChargingLimitRequest) Version() protocol.Version { return protocol.OCPP201 }

func (NotifyChargingLimitRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (NotifyChargingLimitRequest) SchemaName() string { return "NotifyChargingLimitRequest.json" }

func (message NotifyChargingLimitRequest) Validate() error {
	return validation.Validate("NotifyChargingLimitRequest.json", schemaNotifyChargingLimitRequest, message)
}

func (NotifyChargingLimitRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("NotifyChargingLimitRequest.json", schemaNotifyChargingLimitRequest, data)
}
