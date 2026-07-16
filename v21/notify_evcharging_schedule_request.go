// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = NotifyEVChargingScheduleRequest{}

var schemaNotifyEVChargingScheduleRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"timeBase": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "chargingSchedule": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"id": &validation.Schema{Type: "integer", AllowAdditional: true}, "limitAtSoC": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"soc": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true, Maximum: 100, HasMaximum: true}, "limit": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"soc", "limit"}}, "startSchedule": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "duration": &validation.Schema{Type: "integer", AllowAdditional: true}, "chargingRateUnit": &validation.Schema{Type: "string", Enum: []string{"W", "A"}}, "minChargingRate": &validation.Schema{Type: "number", AllowAdditional: true}, "powerTolerance": &validation.Schema{Type: "number", AllowAdditional: true}, "signatureId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "digestValue": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 88, HasMaxLength: true}, "useLocalTime": &validation.Schema{Type: "boolean", AllowAdditional: true}, "chargingSchedulePeriod": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"startPeriod": &validation.Schema{Type: "integer", AllowAdditional: true}, "limit": &validation.Schema{Type: "number", AllowAdditional: true}, "limit_L2": &validation.Schema{Type: "number", AllowAdditional: true}, "limit_L3": &validation.Schema{Type: "number", AllowAdditional: true}, "numberPhases": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true, Maximum: 3, HasMaximum: true}, "phaseToUse": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true, Maximum: 3, HasMaximum: true}, "dischargeLimit": &validation.Schema{Type: "number", AllowAdditional: true, Maximum: 0, HasMaximum: true}, "dischargeLimit_L2": &validation.Schema{Type: "number", AllowAdditional: true, Maximum: 0, HasMaximum: true}, "dischargeLimit_L3": &validation.Schema{Type: "number", AllowAdditional: true, Maximum: 0, HasMaximum: true}, "setpoint": &validation.Schema{Type: "number", AllowAdditional: true}, "setpoint_L2": &validation.Schema{Type: "number", AllowAdditional: true}, "setpoint_L3": &validation.Schema{Type: "number", AllowAdditional: true}, "setpointReactive": &validation.Schema{Type: "number", AllowAdditional: true}, "setpointReactive_L2": &validation.Schema{Type: "number", AllowAdditional: true}, "setpointReactive_L3": &validation.Schema{Type: "number", AllowAdditional: true}, "preconditioningRequest": &validation.Schema{Type: "boolean", AllowAdditional: true}, "evseSleep": &validation.Schema{Type: "boolean", AllowAdditional: true}, "v2xBaseline": &validation.Schema{Type: "number", AllowAdditional: true}, "operationMode": &validation.Schema{Type: "string", Enum: []string{"Idle", "ChargingOnly", "CentralSetpoint", "ExternalSetpoint", "ExternalLimits", "CentralFrequency", "LocalFrequency", "LocalLoadBalancing"}}, "v2xFreqWattCurve": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"frequency": &validation.Schema{Type: "number", AllowAdditional: true}, "power": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"frequency", "power"}}, MinItems: 1, HasMinItems: true, MaxItems: 20, HasMaxItems: true}, "v2xSignalWattCurve": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"signal": &validation.Schema{Type: "integer", AllowAdditional: true}, "power": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"signal", "power"}}, MinItems: 1, HasMinItems: true, MaxItems: 20, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"startPeriod"}}, MinItems: 1, HasMinItems: true, MaxItems: 1024, HasMaxItems: true}, "randomizedDelay": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "salesTariff": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"id": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "salesTariffDescription": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 32, HasMaxLength: true}, "numEPriceLevels": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "salesTariffEntry": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"relativeTimeInterval": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"start": &validation.Schema{Type: "integer", AllowAdditional: true}, "duration": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"start"}}, "ePriceLevel": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "consumptionCost": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"startValue": &validation.Schema{Type: "number", AllowAdditional: true}, "cost": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"costKind": &validation.Schema{Type: "string", Enum: []string{"CarbonDioxideEmission", "RelativePricePercentage", "RenewableGenerationPercentage"}}, "amount": &validation.Schema{Type: "integer", AllowAdditional: true}, "amountMultiplier": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"costKind", "amount"}}, MinItems: 1, HasMinItems: true, MaxItems: 3, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"startValue", "cost"}}, MinItems: 1, HasMinItems: true, MaxItems: 3, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"relativeTimeInterval"}}, MinItems: 1, HasMinItems: true, MaxItems: 1024, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id", "salesTariffEntry"}}, "absolutePriceSchedule": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"timeAnchor": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "priceScheduleID": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "priceScheduleDescription": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 160, HasMaxLength: true}, "currency": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 3, HasMaxLength: true}, "language": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 8, HasMaxLength: true}, "priceAlgorithm": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 2000, HasMaxLength: true}, "minimumCost": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exponent": &validation.Schema{Type: "integer", AllowAdditional: true}, "value": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"exponent", "value"}}, "maximumCost": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exponent": &validation.Schema{Type: "integer", AllowAdditional: true}, "value": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"exponent", "value"}}, "priceRuleStacks": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"duration": &validation.Schema{Type: "integer", AllowAdditional: true}, "priceRule": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"parkingFeePeriod": &validation.Schema{Type: "integer", AllowAdditional: true}, "carbonDioxideEmission": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "renewableGenerationPercentage": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true, Maximum: 100, HasMaximum: true}, "energyFee": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exponent": &validation.Schema{Type: "integer", AllowAdditional: true}, "value": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"exponent", "value"}}, "parkingFee": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exponent": &validation.Schema{Type: "integer", AllowAdditional: true}, "value": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"exponent", "value"}}, "powerRangeStart": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exponent": &validation.Schema{Type: "integer", AllowAdditional: true}, "value": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"exponent", "value"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"energyFee", "powerRangeStart"}}, MinItems: 1, HasMinItems: true, MaxItems: 8, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"duration", "priceRule"}}, MinItems: 1, HasMinItems: true, MaxItems: 1024, HasMaxItems: true}, "taxRules": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"taxRuleID": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "taxRuleName": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 100, HasMaxLength: true}, "taxIncludedInPrice": &validation.Schema{Type: "boolean", AllowAdditional: true}, "appliesToEnergyFee": &validation.Schema{Type: "boolean", AllowAdditional: true}, "appliesToParkingFee": &validation.Schema{Type: "boolean", AllowAdditional: true}, "appliesToOverstayFee": &validation.Schema{Type: "boolean", AllowAdditional: true}, "appliesToMinimumMaximumCost": &validation.Schema{Type: "boolean", AllowAdditional: true}, "taxRate": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exponent": &validation.Schema{Type: "integer", AllowAdditional: true}, "value": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"exponent", "value"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"taxRuleID", "appliesToEnergyFee", "appliesToParkingFee", "appliesToOverstayFee", "appliesToMinimumMaximumCost", "taxRate"}}, MinItems: 1, HasMinItems: true, MaxItems: 10, HasMaxItems: true}, "overstayRuleList": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"overstayPowerThreshold": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exponent": &validation.Schema{Type: "integer", AllowAdditional: true}, "value": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"exponent", "value"}}, "overstayRule": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"overstayFee": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exponent": &validation.Schema{Type: "integer", AllowAdditional: true}, "value": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"exponent", "value"}}, "overstayRuleDescription": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 32, HasMaxLength: true}, "startTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "overstayFeePeriod": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"startTime", "overstayFeePeriod", "overstayFee"}}, MinItems: 1, HasMinItems: true, MaxItems: 5, HasMaxItems: true}, "overstayTimeThreshold": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"overstayRule"}}, "additionalSelectedServices": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"serviceFee": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exponent": &validation.Schema{Type: "integer", AllowAdditional: true}, "value": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"exponent", "value"}}, "serviceName": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 80, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"serviceName", "serviceFee"}}, MinItems: 1, HasMinItems: true, MaxItems: 5, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"timeAnchor", "priceScheduleID", "currency", "language", "priceAlgorithm", "priceRuleStacks"}}, "priceLevelSchedule": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"priceLevelScheduleEntries": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"duration": &validation.Schema{Type: "integer", AllowAdditional: true}, "priceLevel": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"duration", "priceLevel"}}, MinItems: 1, HasMinItems: true, MaxItems: 100, HasMaxItems: true}, "timeAnchor": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "priceScheduleId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "priceScheduleDescription": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 32, HasMaxLength: true}, "numberOfPriceLevels": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"timeAnchor", "priceScheduleId", "numberOfPriceLevels", "priceLevelScheduleEntries"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id", "chargingRateUnit", "chargingSchedulePeriod"}}, "evseId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 1, HasMinimum: true}, "selectedChargingScheduleId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "powerToleranceAcceptance": &validation.Schema{Type: "boolean", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"timeBase", "evseId", "chargingSchedule"}}

type NotifyEVChargingScheduleRequest struct {
	TimeBase                   string                                          `json:"timeBase"`
	ChargingSchedule           NotifyEVChargingScheduleRequestChargingSchedule `json:"chargingSchedule"`
	EVSEID                     int                                             `json:"evseId"`
	SelectedChargingScheduleID *int                                            `json:"selectedChargingScheduleId,omitempty"`
	PowerToleranceAcceptance   *bool                                           `json:"powerToleranceAcceptance,omitempty"`
	CustomData                 *NotifyEVChargingScheduleRequestCustomData      `json:"customData,omitempty"`
}

type NotifyEVChargingScheduleRequestChargingSchedule struct {
	ID                     int                                                     `json:"id"`
	LimitAtSoC             *NotifyEVChargingScheduleRequestLimitAtSoC              `json:"limitAtSoC,omitempty"`
	StartSchedule          *string                                                 `json:"startSchedule,omitempty"`
	Duration               *int                                                    `json:"duration,omitempty"`
	ChargingRateUnit       NotifyEVChargingScheduleRequestChargingRateUnitEnum     `json:"chargingRateUnit"`
	MinChargingRate        *float64                                                `json:"minChargingRate,omitempty"`
	PowerTolerance         *float64                                                `json:"powerTolerance,omitempty"`
	SignatureID            *int                                                    `json:"signatureId,omitempty"`
	DigestValue            *string                                                 `json:"digestValue,omitempty"`
	UseLocalTime           *bool                                                   `json:"useLocalTime,omitempty"`
	ChargingSchedulePeriod []NotifyEVChargingScheduleRequestChargingSchedulePeriod `json:"chargingSchedulePeriod"`
	RandomizedDelay        *int                                                    `json:"randomizedDelay,omitempty"`
	SalesTariff            *NotifyEVChargingScheduleRequestSalesTariff             `json:"salesTariff,omitempty"`
	AbsolutePriceSchedule  *NotifyEVChargingScheduleRequestAbsolutePriceSchedule   `json:"absolutePriceSchedule,omitempty"`
	PriceLevelSchedule     *NotifyEVChargingScheduleRequestPriceLevelSchedule      `json:"priceLevelSchedule,omitempty"`
	CustomData             *NotifyEVChargingScheduleRequestCustomData              `json:"customData,omitempty"`
}

type NotifyEVChargingScheduleRequestPriceLevelSchedule struct {
	PriceLevelScheduleEntries []NotifyEVChargingScheduleRequestPriceLevelScheduleEntry `json:"priceLevelScheduleEntries"`
	TimeAnchor                string                                                   `json:"timeAnchor"`
	PriceScheduleID           int                                                      `json:"priceScheduleId"`
	PriceScheduleDescription  *string                                                  `json:"priceScheduleDescription,omitempty"`
	NumberOfPriceLevels       int                                                      `json:"numberOfPriceLevels"`
	CustomData                *NotifyEVChargingScheduleRequestCustomData               `json:"customData,omitempty"`
}

type NotifyEVChargingScheduleRequestPriceLevelScheduleEntry struct {
	Duration   int                                        `json:"duration"`
	PriceLevel int                                        `json:"priceLevel"`
	CustomData *NotifyEVChargingScheduleRequestCustomData `json:"customData,omitempty"`
}

type NotifyEVChargingScheduleRequestAbsolutePriceSchedule struct {
	TimeAnchor                 string                                                      `json:"timeAnchor"`
	PriceScheduleID            int                                                         `json:"priceScheduleID"`
	PriceScheduleDescription   *string                                                     `json:"priceScheduleDescription,omitempty"`
	Currency                   string                                                      `json:"currency"`
	Language                   string                                                      `json:"language"`
	PriceAlgorithm             string                                                      `json:"priceAlgorithm"`
	MinimumCost                *NotifyEVChargingScheduleRequestRationalNumber              `json:"minimumCost,omitempty"`
	MaximumCost                *NotifyEVChargingScheduleRequestRationalNumber              `json:"maximumCost,omitempty"`
	PriceRuleStacks            []NotifyEVChargingScheduleRequestPriceRuleStack             `json:"priceRuleStacks"`
	TaxRules                   []NotifyEVChargingScheduleRequestTaxRule                    `json:"taxRules,omitempty"`
	OverstayRuleList           *NotifyEVChargingScheduleRequestOverstayRuleList            `json:"overstayRuleList,omitempty"`
	AdditionalSelectedServices []NotifyEVChargingScheduleRequestAdditionalSelectedServices `json:"additionalSelectedServices,omitempty"`
	CustomData                 *NotifyEVChargingScheduleRequestCustomData                  `json:"customData,omitempty"`
}

type NotifyEVChargingScheduleRequestAdditionalSelectedServices struct {
	ServiceFee  NotifyEVChargingScheduleRequestRationalNumber `json:"serviceFee"`
	ServiceName string                                        `json:"serviceName"`
	CustomData  *NotifyEVChargingScheduleRequestCustomData    `json:"customData,omitempty"`
}

type NotifyEVChargingScheduleRequestOverstayRuleList struct {
	OverstayPowerThreshold *NotifyEVChargingScheduleRequestRationalNumber `json:"overstayPowerThreshold,omitempty"`
	OverstayRule           []NotifyEVChargingScheduleRequestOverstayRule  `json:"overstayRule"`
	OverstayTimeThreshold  *int                                           `json:"overstayTimeThreshold,omitempty"`
	CustomData             *NotifyEVChargingScheduleRequestCustomData     `json:"customData,omitempty"`
}

type NotifyEVChargingScheduleRequestOverstayRule struct {
	OverstayFee             NotifyEVChargingScheduleRequestRationalNumber `json:"overstayFee"`
	OverstayRuleDescription *string                                       `json:"overstayRuleDescription,omitempty"`
	StartTime               int                                           `json:"startTime"`
	OverstayFeePeriod       int                                           `json:"overstayFeePeriod"`
	CustomData              *NotifyEVChargingScheduleRequestCustomData    `json:"customData,omitempty"`
}

type NotifyEVChargingScheduleRequestTaxRule struct {
	TaxRuleID                   int                                           `json:"taxRuleID"`
	TaxRuleName                 *string                                       `json:"taxRuleName,omitempty"`
	TaxIncludedInPrice          *bool                                         `json:"taxIncludedInPrice,omitempty"`
	AppliesToEnergyFee          bool                                          `json:"appliesToEnergyFee"`
	AppliesToParkingFee         bool                                          `json:"appliesToParkingFee"`
	AppliesToOverstayFee        bool                                          `json:"appliesToOverstayFee"`
	AppliesToMinimumMaximumCost bool                                          `json:"appliesToMinimumMaximumCost"`
	TaxRate                     NotifyEVChargingScheduleRequestRationalNumber `json:"taxRate"`
	CustomData                  *NotifyEVChargingScheduleRequestCustomData    `json:"customData,omitempty"`
}

type NotifyEVChargingScheduleRequestPriceRuleStack struct {
	Duration   int                                        `json:"duration"`
	PriceRule  []NotifyEVChargingScheduleRequestPriceRule `json:"priceRule"`
	CustomData *NotifyEVChargingScheduleRequestCustomData `json:"customData,omitempty"`
}

type NotifyEVChargingScheduleRequestPriceRule struct {
	ParkingFeePeriod              *int                                           `json:"parkingFeePeriod,omitempty"`
	CarbonDioxideEmission         *int                                           `json:"carbonDioxideEmission,omitempty"`
	RenewableGenerationPercentage *int                                           `json:"renewableGenerationPercentage,omitempty"`
	EnergyFee                     NotifyEVChargingScheduleRequestRationalNumber  `json:"energyFee"`
	ParkingFee                    *NotifyEVChargingScheduleRequestRationalNumber `json:"parkingFee,omitempty"`
	PowerRangeStart               NotifyEVChargingScheduleRequestRationalNumber  `json:"powerRangeStart"`
	CustomData                    *NotifyEVChargingScheduleRequestCustomData     `json:"customData,omitempty"`
}

type NotifyEVChargingScheduleRequestRationalNumber struct {
	Exponent   int                                        `json:"exponent"`
	Value      int                                        `json:"value"`
	CustomData *NotifyEVChargingScheduleRequestCustomData `json:"customData,omitempty"`
}

type NotifyEVChargingScheduleRequestSalesTariff struct {
	ID                     int                                               `json:"id"`
	SalesTariffDescription *string                                           `json:"salesTariffDescription,omitempty"`
	NumEPriceLevels        *int                                              `json:"numEPriceLevels,omitempty"`
	SalesTariffEntry       []NotifyEVChargingScheduleRequestSalesTariffEntry `json:"salesTariffEntry"`
	CustomData             *NotifyEVChargingScheduleRequestCustomData        `json:"customData,omitempty"`
}

type NotifyEVChargingScheduleRequestSalesTariffEntry struct {
	RelativeTimeInterval NotifyEVChargingScheduleRequestRelativeTimeInterval `json:"relativeTimeInterval"`
	EPriceLevel          *int                                                `json:"ePriceLevel,omitempty"`
	ConsumptionCost      []NotifyEVChargingScheduleRequestConsumptionCost    `json:"consumptionCost,omitempty"`
	CustomData           *NotifyEVChargingScheduleRequestCustomData          `json:"customData,omitempty"`
}

type NotifyEVChargingScheduleRequestConsumptionCost struct {
	StartValue float64                                    `json:"startValue"`
	Cost       []NotifyEVChargingScheduleRequestCost      `json:"cost"`
	CustomData *NotifyEVChargingScheduleRequestCustomData `json:"customData,omitempty"`
}

type NotifyEVChargingScheduleRequestCost struct {
	CostKind         NotifyEVChargingScheduleRequestCostKindEnum `json:"costKind"`
	Amount           int                                         `json:"amount"`
	AmountMultiplier *int                                        `json:"amountMultiplier,omitempty"`
	CustomData       *NotifyEVChargingScheduleRequestCustomData  `json:"customData,omitempty"`
}

type NotifyEVChargingScheduleRequestCostKindEnum string

const (
	NotifyEVChargingScheduleRequestCostKindEnumCarbonDioxideEmission         NotifyEVChargingScheduleRequestCostKindEnum = "CarbonDioxideEmission"
	NotifyEVChargingScheduleRequestCostKindEnumRelativePricePercentage       NotifyEVChargingScheduleRequestCostKindEnum = "RelativePricePercentage"
	NotifyEVChargingScheduleRequestCostKindEnumRenewableGenerationPercentage NotifyEVChargingScheduleRequestCostKindEnum = "RenewableGenerationPercentage"
)

type NotifyEVChargingScheduleRequestRelativeTimeInterval struct {
	Start      int                                        `json:"start"`
	Duration   *int                                       `json:"duration,omitempty"`
	CustomData *NotifyEVChargingScheduleRequestCustomData `json:"customData,omitempty"`
}

type NotifyEVChargingScheduleRequestChargingSchedulePeriod struct {
	StartPeriod            int                                                 `json:"startPeriod"`
	Limit                  *float64                                            `json:"limit,omitempty"`
	LimitL2                *float64                                            `json:"limit_L2,omitempty"`
	LimitL3                *float64                                            `json:"limit_L3,omitempty"`
	NumberPhases           *int                                                `json:"numberPhases,omitempty"`
	PhaseToUse             *int                                                `json:"phaseToUse,omitempty"`
	DischargeLimit         *float64                                            `json:"dischargeLimit,omitempty"`
	DischargeLimitL2       *float64                                            `json:"dischargeLimit_L2,omitempty"`
	DischargeLimitL3       *float64                                            `json:"dischargeLimit_L3,omitempty"`
	Setpoint               *float64                                            `json:"setpoint,omitempty"`
	SetpointL2             *float64                                            `json:"setpoint_L2,omitempty"`
	SetpointL3             *float64                                            `json:"setpoint_L3,omitempty"`
	SetpointReactive       *float64                                            `json:"setpointReactive,omitempty"`
	SetpointReactiveL2     *float64                                            `json:"setpointReactive_L2,omitempty"`
	SetpointReactiveL3     *float64                                            `json:"setpointReactive_L3,omitempty"`
	PreconditioningRequest *bool                                               `json:"preconditioningRequest,omitempty"`
	EVSESleep              *bool                                               `json:"evseSleep,omitempty"`
	V2XBaseline            *float64                                            `json:"v2xBaseline,omitempty"`
	OperationMode          *NotifyEVChargingScheduleRequestOperationModeEnum   `json:"operationMode,omitempty"`
	V2XFreqWattCurve       []NotifyEVChargingScheduleRequestV2XFreqWattPoint   `json:"v2xFreqWattCurve,omitempty"`
	V2XSignalWattCurve     []NotifyEVChargingScheduleRequestV2XSignalWattPoint `json:"v2xSignalWattCurve,omitempty"`
	CustomData             *NotifyEVChargingScheduleRequestCustomData          `json:"customData,omitempty"`
}

type NotifyEVChargingScheduleRequestV2XSignalWattPoint struct {
	Signal     int                                        `json:"signal"`
	Power      float64                                    `json:"power"`
	CustomData *NotifyEVChargingScheduleRequestCustomData `json:"customData,omitempty"`
}

type NotifyEVChargingScheduleRequestV2XFreqWattPoint struct {
	Frequency  float64                                    `json:"frequency"`
	Power      float64                                    `json:"power"`
	CustomData *NotifyEVChargingScheduleRequestCustomData `json:"customData,omitempty"`
}

type NotifyEVChargingScheduleRequestOperationModeEnum string

const (
	NotifyEVChargingScheduleRequestOperationModeEnumIdle               NotifyEVChargingScheduleRequestOperationModeEnum = "Idle"
	NotifyEVChargingScheduleRequestOperationModeEnumChargingOnly       NotifyEVChargingScheduleRequestOperationModeEnum = "ChargingOnly"
	NotifyEVChargingScheduleRequestOperationModeEnumCentralSetpoint    NotifyEVChargingScheduleRequestOperationModeEnum = "CentralSetpoint"
	NotifyEVChargingScheduleRequestOperationModeEnumExternalSetpoint   NotifyEVChargingScheduleRequestOperationModeEnum = "ExternalSetpoint"
	NotifyEVChargingScheduleRequestOperationModeEnumExternalLimits     NotifyEVChargingScheduleRequestOperationModeEnum = "ExternalLimits"
	NotifyEVChargingScheduleRequestOperationModeEnumCentralFrequency   NotifyEVChargingScheduleRequestOperationModeEnum = "CentralFrequency"
	NotifyEVChargingScheduleRequestOperationModeEnumLocalFrequency     NotifyEVChargingScheduleRequestOperationModeEnum = "LocalFrequency"
	NotifyEVChargingScheduleRequestOperationModeEnumLocalLoadBalancing NotifyEVChargingScheduleRequestOperationModeEnum = "LocalLoadBalancing"
)

type NotifyEVChargingScheduleRequestChargingRateUnitEnum string

const (
	NotifyEVChargingScheduleRequestChargingRateUnitEnumW NotifyEVChargingScheduleRequestChargingRateUnitEnum = "W"
	NotifyEVChargingScheduleRequestChargingRateUnitEnumA NotifyEVChargingScheduleRequestChargingRateUnitEnum = "A"
)

type NotifyEVChargingScheduleRequestLimitAtSoC struct {
	Soc        int                                        `json:"soc"`
	Limit      float64                                    `json:"limit"`
	CustomData *NotifyEVChargingScheduleRequestCustomData `json:"customData,omitempty"`
}

type NotifyEVChargingScheduleRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value NotifyEVChargingScheduleRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *NotifyEVChargingScheduleRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (NotifyEVChargingScheduleRequest) ActionName() string { return "NotifyEVChargingSchedule" }

func (NotifyEVChargingScheduleRequest) Version() protocol.Version { return protocol.OCPP21 }

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
