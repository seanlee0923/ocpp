// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = NotifyEVChargingScheduleRequest{}

var schemaNotifyEVChargingScheduleRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "timeBase": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "chargingSchedule": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "id": &validation.Schema{Type: "integer", AllowAdditional: true}, "startSchedule": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "duration": &validation.Schema{Type: "integer", AllowAdditional: true}, "chargingRateUnit": &validation.Schema{Type: "string", Enum: []string{"W", "A"}}, "chargingSchedulePeriod": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "startPeriod": &validation.Schema{Type: "integer", AllowAdditional: true}, "limit": &validation.Schema{Type: "number", AllowAdditional: true}, "numberPhases": &validation.Schema{Type: "integer", AllowAdditional: true}, "phaseToUse": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"startPeriod", "limit"}}, MinItems: 1, HasMinItems: true, MaxItems: 1024, HasMaxItems: true}, "minChargingRate": &validation.Schema{Type: "number", AllowAdditional: true}, "salesTariff": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "id": &validation.Schema{Type: "integer", AllowAdditional: true}, "salesTariffDescription": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 32, HasMaxLength: true}, "numEPriceLevels": &validation.Schema{Type: "integer", AllowAdditional: true}, "salesTariffEntry": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "relativeTimeInterval": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "start": &validation.Schema{Type: "integer", AllowAdditional: true}, "duration": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"start"}}, "ePriceLevel": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "consumptionCost": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "startValue": &validation.Schema{Type: "number", AllowAdditional: true}, "cost": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "costKind": &validation.Schema{Type: "string", Enum: []string{"CarbonDioxideEmission", "RelativePricePercentage", "RenewableGenerationPercentage"}}, "amount": &validation.Schema{Type: "integer", AllowAdditional: true}, "amountMultiplier": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"costKind", "amount"}}, MinItems: 1, HasMinItems: true, MaxItems: 3, HasMaxItems: true}}, Required: []string{"startValue", "cost"}}, MinItems: 1, HasMinItems: true, MaxItems: 3, HasMaxItems: true}}, Required: []string{"relativeTimeInterval"}}, MinItems: 1, HasMinItems: true, MaxItems: 1024, HasMaxItems: true}}, Required: []string{"id", "salesTariffEntry"}}}, Required: []string{"id", "chargingRateUnit", "chargingSchedulePeriod"}}, "evseId": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"timeBase", "evseId", "chargingSchedule"}}

type NotifyEVChargingScheduleRequest struct {
	CustomData       *NotifyEVChargingScheduleRequestCustomData      `json:"customData,omitempty"`
	TimeBase         string                                          `json:"timeBase"`
	ChargingSchedule NotifyEVChargingScheduleRequestChargingSchedule `json:"chargingSchedule"`
	EVSEID           int                                             `json:"evseId"`
}

type NotifyEVChargingScheduleRequestChargingSchedule struct {
	CustomData             *NotifyEVChargingScheduleRequestCustomData              `json:"customData,omitempty"`
	ID                     int                                                     `json:"id"`
	StartSchedule          *string                                                 `json:"startSchedule,omitempty"`
	Duration               *int                                                    `json:"duration,omitempty"`
	ChargingRateUnit       NotifyEVChargingScheduleRequestChargingRateUnitEnum     `json:"chargingRateUnit"`
	ChargingSchedulePeriod []NotifyEVChargingScheduleRequestChargingSchedulePeriod `json:"chargingSchedulePeriod"`
	MinChargingRate        *float64                                                `json:"minChargingRate,omitempty"`
	SalesTariff            *NotifyEVChargingScheduleRequestSalesTariff             `json:"salesTariff,omitempty"`
}

type NotifyEVChargingScheduleRequestSalesTariff struct {
	CustomData             *NotifyEVChargingScheduleRequestCustomData        `json:"customData,omitempty"`
	ID                     int                                               `json:"id"`
	SalesTariffDescription *string                                           `json:"salesTariffDescription,omitempty"`
	NumEPriceLevels        *int                                              `json:"numEPriceLevels,omitempty"`
	SalesTariffEntry       []NotifyEVChargingScheduleRequestSalesTariffEntry `json:"salesTariffEntry"`
}

type NotifyEVChargingScheduleRequestSalesTariffEntry struct {
	CustomData           *NotifyEVChargingScheduleRequestCustomData          `json:"customData,omitempty"`
	RelativeTimeInterval NotifyEVChargingScheduleRequestRelativeTimeInterval `json:"relativeTimeInterval"`
	EPriceLevel          *int                                                `json:"ePriceLevel,omitempty"`
	ConsumptionCost      []NotifyEVChargingScheduleRequestConsumptionCost    `json:"consumptionCost,omitempty"`
}

type NotifyEVChargingScheduleRequestConsumptionCost struct {
	CustomData *NotifyEVChargingScheduleRequestCustomData `json:"customData,omitempty"`
	StartValue float64                                    `json:"startValue"`
	Cost       []NotifyEVChargingScheduleRequestCost      `json:"cost"`
}

type NotifyEVChargingScheduleRequestCost struct {
	CustomData       *NotifyEVChargingScheduleRequestCustomData  `json:"customData,omitempty"`
	CostKind         NotifyEVChargingScheduleRequestCostKindEnum `json:"costKind"`
	Amount           int                                         `json:"amount"`
	AmountMultiplier *int                                        `json:"amountMultiplier,omitempty"`
}

type NotifyEVChargingScheduleRequestCostKindEnum string

const (
	NotifyEVChargingScheduleRequestCostKindEnumCarbonDioxideEmission         NotifyEVChargingScheduleRequestCostKindEnum = "CarbonDioxideEmission"
	NotifyEVChargingScheduleRequestCostKindEnumRelativePricePercentage       NotifyEVChargingScheduleRequestCostKindEnum = "RelativePricePercentage"
	NotifyEVChargingScheduleRequestCostKindEnumRenewableGenerationPercentage NotifyEVChargingScheduleRequestCostKindEnum = "RenewableGenerationPercentage"
)

type NotifyEVChargingScheduleRequestRelativeTimeInterval struct {
	CustomData *NotifyEVChargingScheduleRequestCustomData `json:"customData,omitempty"`
	Start      int                                        `json:"start"`
	Duration   *int                                       `json:"duration,omitempty"`
}

type NotifyEVChargingScheduleRequestChargingSchedulePeriod struct {
	CustomData   *NotifyEVChargingScheduleRequestCustomData `json:"customData,omitempty"`
	StartPeriod  int                                        `json:"startPeriod"`
	Limit        float64                                    `json:"limit"`
	NumberPhases *int                                       `json:"numberPhases,omitempty"`
	PhaseToUse   *int                                       `json:"phaseToUse,omitempty"`
}

type NotifyEVChargingScheduleRequestChargingRateUnitEnum string

const (
	NotifyEVChargingScheduleRequestChargingRateUnitEnumW NotifyEVChargingScheduleRequestChargingRateUnitEnum = "W"
	NotifyEVChargingScheduleRequestChargingRateUnitEnumA NotifyEVChargingScheduleRequestChargingRateUnitEnum = "A"
)

type NotifyEVChargingScheduleRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (NotifyEVChargingScheduleRequest) ActionName() string { return "NotifyEVChargingSchedule" }

func (NotifyEVChargingScheduleRequest) Version() protocol.Version { return protocol.OCPP201 }

func (NotifyEVChargingScheduleRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (NotifyEVChargingScheduleRequest) SchemaName() string {
	return "NotifyEVChargingScheduleRequest.json"
}

func (message NotifyEVChargingScheduleRequest) Validate() error {
	return validation.Validate("NotifyEVChargingScheduleRequest.json", schemaNotifyEVChargingScheduleRequest, message)
}

func (NotifyEVChargingScheduleRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("NotifyEVChargingScheduleRequest.json", schemaNotifyEVChargingScheduleRequest, data)
}
