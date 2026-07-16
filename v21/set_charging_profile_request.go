// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = SetChargingProfileRequest{}

var schemaSetChargingProfileRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"evseId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "chargingProfile": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"id": &validation.Schema{Type: "integer", AllowAdditional: true}, "stackLevel": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "chargingProfilePurpose": &validation.Schema{Type: "string", Enum: []string{"ChargingStationExternalConstraints", "ChargingStationMaxProfile", "TxDefaultProfile", "TxProfile", "PriorityCharging", "LocalGeneration"}}, "chargingProfileKind": &validation.Schema{Type: "string", Enum: []string{"Absolute", "Recurring", "Relative", "Dynamic"}}, "recurrencyKind": &validation.Schema{Type: "string", Enum: []string{"Daily", "Weekly"}}, "validFrom": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "validTo": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "transactionId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "maxOfflineDuration": &validation.Schema{Type: "integer", AllowAdditional: true}, "chargingSchedule": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"id": &validation.Schema{Type: "integer", AllowAdditional: true}, "limitAtSoC": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"soc": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true, Maximum: 100, HasMaximum: true}, "limit": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"soc", "limit"}}, "startSchedule": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "duration": &validation.Schema{Type: "integer", AllowAdditional: true}, "chargingRateUnit": &validation.Schema{Type: "string", Enum: []string{"W", "A"}}, "minChargingRate": &validation.Schema{Type: "number", AllowAdditional: true}, "powerTolerance": &validation.Schema{Type: "number", AllowAdditional: true}, "signatureId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "digestValue": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 88, HasMaxLength: true}, "useLocalTime": &validation.Schema{Type: "boolean", AllowAdditional: true}, "chargingSchedulePeriod": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"startPeriod": &validation.Schema{Type: "integer", AllowAdditional: true}, "limit": &validation.Schema{Type: "number", AllowAdditional: true}, "limit_L2": &validation.Schema{Type: "number", AllowAdditional: true}, "limit_L3": &validation.Schema{Type: "number", AllowAdditional: true}, "numberPhases": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true, Maximum: 3, HasMaximum: true}, "phaseToUse": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true, Maximum: 3, HasMaximum: true}, "dischargeLimit": &validation.Schema{Type: "number", AllowAdditional: true, Maximum: 0, HasMaximum: true}, "dischargeLimit_L2": &validation.Schema{Type: "number", AllowAdditional: true, Maximum: 0, HasMaximum: true}, "dischargeLimit_L3": &validation.Schema{Type: "number", AllowAdditional: true, Maximum: 0, HasMaximum: true}, "setpoint": &validation.Schema{Type: "number", AllowAdditional: true}, "setpoint_L2": &validation.Schema{Type: "number", AllowAdditional: true}, "setpoint_L3": &validation.Schema{Type: "number", AllowAdditional: true}, "setpointReactive": &validation.Schema{Type: "number", AllowAdditional: true}, "setpointReactive_L2": &validation.Schema{Type: "number", AllowAdditional: true}, "setpointReactive_L3": &validation.Schema{Type: "number", AllowAdditional: true}, "preconditioningRequest": &validation.Schema{Type: "boolean", AllowAdditional: true}, "evseSleep": &validation.Schema{Type: "boolean", AllowAdditional: true}, "v2xBaseline": &validation.Schema{Type: "number", AllowAdditional: true}, "operationMode": &validation.Schema{Type: "string", Enum: []string{"Idle", "ChargingOnly", "CentralSetpoint", "ExternalSetpoint", "ExternalLimits", "CentralFrequency", "LocalFrequency", "LocalLoadBalancing"}}, "v2xFreqWattCurve": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"frequency": &validation.Schema{Type: "number", AllowAdditional: true}, "power": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"frequency", "power"}}, MinItems: 1, HasMinItems: true, MaxItems: 20, HasMaxItems: true}, "v2xSignalWattCurve": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"signal": &validation.Schema{Type: "integer", AllowAdditional: true}, "power": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"signal", "power"}}, MinItems: 1, HasMinItems: true, MaxItems: 20, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"startPeriod"}}, MinItems: 1, HasMinItems: true, MaxItems: 1024, HasMaxItems: true}, "randomizedDelay": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "salesTariff": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"id": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "salesTariffDescription": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 32, HasMaxLength: true}, "numEPriceLevels": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "salesTariffEntry": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"relativeTimeInterval": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"start": &validation.Schema{Type: "integer", AllowAdditional: true}, "duration": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"start"}}, "ePriceLevel": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "consumptionCost": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"startValue": &validation.Schema{Type: "number", AllowAdditional: true}, "cost": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"costKind": &validation.Schema{Type: "string", Enum: []string{"CarbonDioxideEmission", "RelativePricePercentage", "RenewableGenerationPercentage"}}, "amount": &validation.Schema{Type: "integer", AllowAdditional: true}, "amountMultiplier": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"costKind", "amount"}}, MinItems: 1, HasMinItems: true, MaxItems: 3, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"startValue", "cost"}}, MinItems: 1, HasMinItems: true, MaxItems: 3, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"relativeTimeInterval"}}, MinItems: 1, HasMinItems: true, MaxItems: 1024, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id", "salesTariffEntry"}}, "absolutePriceSchedule": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"timeAnchor": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "priceScheduleID": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "priceScheduleDescription": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 160, HasMaxLength: true}, "currency": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 3, HasMaxLength: true}, "language": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 8, HasMaxLength: true}, "priceAlgorithm": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 2000, HasMaxLength: true}, "minimumCost": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exponent": &validation.Schema{Type: "integer", AllowAdditional: true}, "value": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"exponent", "value"}}, "maximumCost": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exponent": &validation.Schema{Type: "integer", AllowAdditional: true}, "value": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"exponent", "value"}}, "priceRuleStacks": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"duration": &validation.Schema{Type: "integer", AllowAdditional: true}, "priceRule": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"parkingFeePeriod": &validation.Schema{Type: "integer", AllowAdditional: true}, "carbonDioxideEmission": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "renewableGenerationPercentage": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true, Maximum: 100, HasMaximum: true}, "energyFee": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exponent": &validation.Schema{Type: "integer", AllowAdditional: true}, "value": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"exponent", "value"}}, "parkingFee": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exponent": &validation.Schema{Type: "integer", AllowAdditional: true}, "value": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"exponent", "value"}}, "powerRangeStart": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exponent": &validation.Schema{Type: "integer", AllowAdditional: true}, "value": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"exponent", "value"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"energyFee", "powerRangeStart"}}, MinItems: 1, HasMinItems: true, MaxItems: 8, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"duration", "priceRule"}}, MinItems: 1, HasMinItems: true, MaxItems: 1024, HasMaxItems: true}, "taxRules": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"taxRuleID": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "taxRuleName": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 100, HasMaxLength: true}, "taxIncludedInPrice": &validation.Schema{Type: "boolean", AllowAdditional: true}, "appliesToEnergyFee": &validation.Schema{Type: "boolean", AllowAdditional: true}, "appliesToParkingFee": &validation.Schema{Type: "boolean", AllowAdditional: true}, "appliesToOverstayFee": &validation.Schema{Type: "boolean", AllowAdditional: true}, "appliesToMinimumMaximumCost": &validation.Schema{Type: "boolean", AllowAdditional: true}, "taxRate": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exponent": &validation.Schema{Type: "integer", AllowAdditional: true}, "value": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"exponent", "value"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"taxRuleID", "appliesToEnergyFee", "appliesToParkingFee", "appliesToOverstayFee", "appliesToMinimumMaximumCost", "taxRate"}}, MinItems: 1, HasMinItems: true, MaxItems: 10, HasMaxItems: true}, "overstayRuleList": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"overstayPowerThreshold": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exponent": &validation.Schema{Type: "integer", AllowAdditional: true}, "value": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"exponent", "value"}}, "overstayRule": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"overstayFee": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exponent": &validation.Schema{Type: "integer", AllowAdditional: true}, "value": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"exponent", "value"}}, "overstayRuleDescription": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 32, HasMaxLength: true}, "startTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "overstayFeePeriod": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"startTime", "overstayFeePeriod", "overstayFee"}}, MinItems: 1, HasMinItems: true, MaxItems: 5, HasMaxItems: true}, "overstayTimeThreshold": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"overstayRule"}}, "additionalSelectedServices": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"serviceFee": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exponent": &validation.Schema{Type: "integer", AllowAdditional: true}, "value": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"exponent", "value"}}, "serviceName": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 80, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"serviceName", "serviceFee"}}, MinItems: 1, HasMinItems: true, MaxItems: 5, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"timeAnchor", "priceScheduleID", "currency", "language", "priceAlgorithm", "priceRuleStacks"}}, "priceLevelSchedule": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"priceLevelScheduleEntries": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"duration": &validation.Schema{Type: "integer", AllowAdditional: true}, "priceLevel": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"duration", "priceLevel"}}, MinItems: 1, HasMinItems: true, MaxItems: 100, HasMaxItems: true}, "timeAnchor": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "priceScheduleId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "priceScheduleDescription": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 32, HasMaxLength: true}, "numberOfPriceLevels": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"timeAnchor", "priceScheduleId", "numberOfPriceLevels", "priceLevelScheduleEntries"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id", "chargingRateUnit", "chargingSchedulePeriod"}}, MinItems: 1, HasMinItems: true, MaxItems: 3, HasMaxItems: true}, "invalidAfterOfflineDuration": &validation.Schema{Type: "boolean", AllowAdditional: true}, "dynUpdateInterval": &validation.Schema{Type: "integer", AllowAdditional: true}, "dynUpdateTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "priceScheduleSignature": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 256, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id", "stackLevel", "chargingProfilePurpose", "chargingProfileKind", "chargingSchedule"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"evseId", "chargingProfile"}}

type SetChargingProfileRequest struct {
	EVSEID          int                                      `json:"evseId"`
	ChargingProfile SetChargingProfileRequestChargingProfile `json:"chargingProfile"`
	CustomData      *SetChargingProfileRequestCustomData     `json:"customData,omitempty"`
}

type SetChargingProfileRequestChargingProfile struct {
	ID                          int                                                 `json:"id"`
	StackLevel                  int                                                 `json:"stackLevel"`
	ChargingProfilePurpose      SetChargingProfileRequestChargingProfilePurposeEnum `json:"chargingProfilePurpose"`
	ChargingProfileKind         SetChargingProfileRequestChargingProfileKindEnum    `json:"chargingProfileKind"`
	RecurrencyKind              *SetChargingProfileRequestRecurrencyKindEnum        `json:"recurrencyKind,omitempty"`
	ValidFrom                   *string                                             `json:"validFrom,omitempty"`
	ValidTo                     *string                                             `json:"validTo,omitempty"`
	TransactionID               *string                                             `json:"transactionId,omitempty"`
	MaxOfflineDuration          *int                                                `json:"maxOfflineDuration,omitempty"`
	ChargingSchedule            []SetChargingProfileRequestChargingSchedule         `json:"chargingSchedule"`
	InvalidAfterOfflineDuration *bool                                               `json:"invalidAfterOfflineDuration,omitempty"`
	DynUpdateInterval           *int                                                `json:"dynUpdateInterval,omitempty"`
	DynUpdateTime               *string                                             `json:"dynUpdateTime,omitempty"`
	PriceScheduleSignature      *string                                             `json:"priceScheduleSignature,omitempty"`
	CustomData                  *SetChargingProfileRequestCustomData                `json:"customData,omitempty"`
}

type SetChargingProfileRequestChargingSchedule struct {
	ID                     int                                               `json:"id"`
	LimitAtSoC             *SetChargingProfileRequestLimitAtSoC              `json:"limitAtSoC,omitempty"`
	StartSchedule          *string                                           `json:"startSchedule,omitempty"`
	Duration               *int                                              `json:"duration,omitempty"`
	ChargingRateUnit       SetChargingProfileRequestChargingRateUnitEnum     `json:"chargingRateUnit"`
	MinChargingRate        *float64                                          `json:"minChargingRate,omitempty"`
	PowerTolerance         *float64                                          `json:"powerTolerance,omitempty"`
	SignatureID            *int                                              `json:"signatureId,omitempty"`
	DigestValue            *string                                           `json:"digestValue,omitempty"`
	UseLocalTime           *bool                                             `json:"useLocalTime,omitempty"`
	ChargingSchedulePeriod []SetChargingProfileRequestChargingSchedulePeriod `json:"chargingSchedulePeriod"`
	RandomizedDelay        *int                                              `json:"randomizedDelay,omitempty"`
	SalesTariff            *SetChargingProfileRequestSalesTariff             `json:"salesTariff,omitempty"`
	AbsolutePriceSchedule  *SetChargingProfileRequestAbsolutePriceSchedule   `json:"absolutePriceSchedule,omitempty"`
	PriceLevelSchedule     *SetChargingProfileRequestPriceLevelSchedule      `json:"priceLevelSchedule,omitempty"`
	CustomData             *SetChargingProfileRequestCustomData              `json:"customData,omitempty"`
}

type SetChargingProfileRequestPriceLevelSchedule struct {
	PriceLevelScheduleEntries []SetChargingProfileRequestPriceLevelScheduleEntry `json:"priceLevelScheduleEntries"`
	TimeAnchor                string                                             `json:"timeAnchor"`
	PriceScheduleID           int                                                `json:"priceScheduleId"`
	PriceScheduleDescription  *string                                            `json:"priceScheduleDescription,omitempty"`
	NumberOfPriceLevels       int                                                `json:"numberOfPriceLevels"`
	CustomData                *SetChargingProfileRequestCustomData               `json:"customData,omitempty"`
}

type SetChargingProfileRequestPriceLevelScheduleEntry struct {
	Duration   int                                  `json:"duration"`
	PriceLevel int                                  `json:"priceLevel"`
	CustomData *SetChargingProfileRequestCustomData `json:"customData,omitempty"`
}

type SetChargingProfileRequestAbsolutePriceSchedule struct {
	TimeAnchor                 string                                                `json:"timeAnchor"`
	PriceScheduleID            int                                                   `json:"priceScheduleID"`
	PriceScheduleDescription   *string                                               `json:"priceScheduleDescription,omitempty"`
	Currency                   string                                                `json:"currency"`
	Language                   string                                                `json:"language"`
	PriceAlgorithm             string                                                `json:"priceAlgorithm"`
	MinimumCost                *SetChargingProfileRequestRationalNumber              `json:"minimumCost,omitempty"`
	MaximumCost                *SetChargingProfileRequestRationalNumber              `json:"maximumCost,omitempty"`
	PriceRuleStacks            []SetChargingProfileRequestPriceRuleStack             `json:"priceRuleStacks"`
	TaxRules                   []SetChargingProfileRequestTaxRule                    `json:"taxRules,omitempty"`
	OverstayRuleList           *SetChargingProfileRequestOverstayRuleList            `json:"overstayRuleList,omitempty"`
	AdditionalSelectedServices []SetChargingProfileRequestAdditionalSelectedServices `json:"additionalSelectedServices,omitempty"`
	CustomData                 *SetChargingProfileRequestCustomData                  `json:"customData,omitempty"`
}

type SetChargingProfileRequestAdditionalSelectedServices struct {
	ServiceFee  SetChargingProfileRequestRationalNumber `json:"serviceFee"`
	ServiceName string                                  `json:"serviceName"`
	CustomData  *SetChargingProfileRequestCustomData    `json:"customData,omitempty"`
}

type SetChargingProfileRequestOverstayRuleList struct {
	OverstayPowerThreshold *SetChargingProfileRequestRationalNumber `json:"overstayPowerThreshold,omitempty"`
	OverstayRule           []SetChargingProfileRequestOverstayRule  `json:"overstayRule"`
	OverstayTimeThreshold  *int                                     `json:"overstayTimeThreshold,omitempty"`
	CustomData             *SetChargingProfileRequestCustomData     `json:"customData,omitempty"`
}

type SetChargingProfileRequestOverstayRule struct {
	OverstayFee             SetChargingProfileRequestRationalNumber `json:"overstayFee"`
	OverstayRuleDescription *string                                 `json:"overstayRuleDescription,omitempty"`
	StartTime               int                                     `json:"startTime"`
	OverstayFeePeriod       int                                     `json:"overstayFeePeriod"`
	CustomData              *SetChargingProfileRequestCustomData    `json:"customData,omitempty"`
}

type SetChargingProfileRequestTaxRule struct {
	TaxRuleID                   int                                     `json:"taxRuleID"`
	TaxRuleName                 *string                                 `json:"taxRuleName,omitempty"`
	TaxIncludedInPrice          *bool                                   `json:"taxIncludedInPrice,omitempty"`
	AppliesToEnergyFee          bool                                    `json:"appliesToEnergyFee"`
	AppliesToParkingFee         bool                                    `json:"appliesToParkingFee"`
	AppliesToOverstayFee        bool                                    `json:"appliesToOverstayFee"`
	AppliesToMinimumMaximumCost bool                                    `json:"appliesToMinimumMaximumCost"`
	TaxRate                     SetChargingProfileRequestRationalNumber `json:"taxRate"`
	CustomData                  *SetChargingProfileRequestCustomData    `json:"customData,omitempty"`
}

type SetChargingProfileRequestPriceRuleStack struct {
	Duration   int                                  `json:"duration"`
	PriceRule  []SetChargingProfileRequestPriceRule `json:"priceRule"`
	CustomData *SetChargingProfileRequestCustomData `json:"customData,omitempty"`
}

type SetChargingProfileRequestPriceRule struct {
	ParkingFeePeriod              *int                                     `json:"parkingFeePeriod,omitempty"`
	CarbonDioxideEmission         *int                                     `json:"carbonDioxideEmission,omitempty"`
	RenewableGenerationPercentage *int                                     `json:"renewableGenerationPercentage,omitempty"`
	EnergyFee                     SetChargingProfileRequestRationalNumber  `json:"energyFee"`
	ParkingFee                    *SetChargingProfileRequestRationalNumber `json:"parkingFee,omitempty"`
	PowerRangeStart               SetChargingProfileRequestRationalNumber  `json:"powerRangeStart"`
	CustomData                    *SetChargingProfileRequestCustomData     `json:"customData,omitempty"`
}

type SetChargingProfileRequestRationalNumber struct {
	Exponent   int                                  `json:"exponent"`
	Value      int                                  `json:"value"`
	CustomData *SetChargingProfileRequestCustomData `json:"customData,omitempty"`
}

type SetChargingProfileRequestSalesTariff struct {
	ID                     int                                         `json:"id"`
	SalesTariffDescription *string                                     `json:"salesTariffDescription,omitempty"`
	NumEPriceLevels        *int                                        `json:"numEPriceLevels,omitempty"`
	SalesTariffEntry       []SetChargingProfileRequestSalesTariffEntry `json:"salesTariffEntry"`
	CustomData             *SetChargingProfileRequestCustomData        `json:"customData,omitempty"`
}

type SetChargingProfileRequestSalesTariffEntry struct {
	RelativeTimeInterval SetChargingProfileRequestRelativeTimeInterval `json:"relativeTimeInterval"`
	EPriceLevel          *int                                          `json:"ePriceLevel,omitempty"`
	ConsumptionCost      []SetChargingProfileRequestConsumptionCost    `json:"consumptionCost,omitempty"`
	CustomData           *SetChargingProfileRequestCustomData          `json:"customData,omitempty"`
}

type SetChargingProfileRequestConsumptionCost struct {
	StartValue float64                              `json:"startValue"`
	Cost       []SetChargingProfileRequestCost      `json:"cost"`
	CustomData *SetChargingProfileRequestCustomData `json:"customData,omitempty"`
}

type SetChargingProfileRequestCost struct {
	CostKind         SetChargingProfileRequestCostKindEnum `json:"costKind"`
	Amount           int                                   `json:"amount"`
	AmountMultiplier *int                                  `json:"amountMultiplier,omitempty"`
	CustomData       *SetChargingProfileRequestCustomData  `json:"customData,omitempty"`
}

type SetChargingProfileRequestCostKindEnum string

const (
	SetChargingProfileRequestCostKindEnumCarbonDioxideEmission         SetChargingProfileRequestCostKindEnum = "CarbonDioxideEmission"
	SetChargingProfileRequestCostKindEnumRelativePricePercentage       SetChargingProfileRequestCostKindEnum = "RelativePricePercentage"
	SetChargingProfileRequestCostKindEnumRenewableGenerationPercentage SetChargingProfileRequestCostKindEnum = "RenewableGenerationPercentage"
)

type SetChargingProfileRequestRelativeTimeInterval struct {
	Start      int                                  `json:"start"`
	Duration   *int                                 `json:"duration,omitempty"`
	CustomData *SetChargingProfileRequestCustomData `json:"customData,omitempty"`
}

type SetChargingProfileRequestChargingSchedulePeriod struct {
	StartPeriod            int                                           `json:"startPeriod"`
	Limit                  *float64                                      `json:"limit,omitempty"`
	LimitL2                *float64                                      `json:"limit_L2,omitempty"`
	LimitL3                *float64                                      `json:"limit_L3,omitempty"`
	NumberPhases           *int                                          `json:"numberPhases,omitempty"`
	PhaseToUse             *int                                          `json:"phaseToUse,omitempty"`
	DischargeLimit         *float64                                      `json:"dischargeLimit,omitempty"`
	DischargeLimitL2       *float64                                      `json:"dischargeLimit_L2,omitempty"`
	DischargeLimitL3       *float64                                      `json:"dischargeLimit_L3,omitempty"`
	Setpoint               *float64                                      `json:"setpoint,omitempty"`
	SetpointL2             *float64                                      `json:"setpoint_L2,omitempty"`
	SetpointL3             *float64                                      `json:"setpoint_L3,omitempty"`
	SetpointReactive       *float64                                      `json:"setpointReactive,omitempty"`
	SetpointReactiveL2     *float64                                      `json:"setpointReactive_L2,omitempty"`
	SetpointReactiveL3     *float64                                      `json:"setpointReactive_L3,omitempty"`
	PreconditioningRequest *bool                                         `json:"preconditioningRequest,omitempty"`
	EVSESleep              *bool                                         `json:"evseSleep,omitempty"`
	V2XBaseline            *float64                                      `json:"v2xBaseline,omitempty"`
	OperationMode          *SetChargingProfileRequestOperationModeEnum   `json:"operationMode,omitempty"`
	V2XFreqWattCurve       []SetChargingProfileRequestV2XFreqWattPoint   `json:"v2xFreqWattCurve,omitempty"`
	V2XSignalWattCurve     []SetChargingProfileRequestV2XSignalWattPoint `json:"v2xSignalWattCurve,omitempty"`
	CustomData             *SetChargingProfileRequestCustomData          `json:"customData,omitempty"`
}

type SetChargingProfileRequestV2XSignalWattPoint struct {
	Signal     int                                  `json:"signal"`
	Power      float64                              `json:"power"`
	CustomData *SetChargingProfileRequestCustomData `json:"customData,omitempty"`
}

type SetChargingProfileRequestV2XFreqWattPoint struct {
	Frequency  float64                              `json:"frequency"`
	Power      float64                              `json:"power"`
	CustomData *SetChargingProfileRequestCustomData `json:"customData,omitempty"`
}

type SetChargingProfileRequestOperationModeEnum string

const (
	SetChargingProfileRequestOperationModeEnumIdle               SetChargingProfileRequestOperationModeEnum = "Idle"
	SetChargingProfileRequestOperationModeEnumChargingOnly       SetChargingProfileRequestOperationModeEnum = "ChargingOnly"
	SetChargingProfileRequestOperationModeEnumCentralSetpoint    SetChargingProfileRequestOperationModeEnum = "CentralSetpoint"
	SetChargingProfileRequestOperationModeEnumExternalSetpoint   SetChargingProfileRequestOperationModeEnum = "ExternalSetpoint"
	SetChargingProfileRequestOperationModeEnumExternalLimits     SetChargingProfileRequestOperationModeEnum = "ExternalLimits"
	SetChargingProfileRequestOperationModeEnumCentralFrequency   SetChargingProfileRequestOperationModeEnum = "CentralFrequency"
	SetChargingProfileRequestOperationModeEnumLocalFrequency     SetChargingProfileRequestOperationModeEnum = "LocalFrequency"
	SetChargingProfileRequestOperationModeEnumLocalLoadBalancing SetChargingProfileRequestOperationModeEnum = "LocalLoadBalancing"
)

type SetChargingProfileRequestChargingRateUnitEnum string

const (
	SetChargingProfileRequestChargingRateUnitEnumW SetChargingProfileRequestChargingRateUnitEnum = "W"
	SetChargingProfileRequestChargingRateUnitEnumA SetChargingProfileRequestChargingRateUnitEnum = "A"
)

type SetChargingProfileRequestLimitAtSoC struct {
	Soc        int                                  `json:"soc"`
	Limit      float64                              `json:"limit"`
	CustomData *SetChargingProfileRequestCustomData `json:"customData,omitempty"`
}

type SetChargingProfileRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

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
	SetChargingProfileRequestChargingProfileKindEnumDynamic   SetChargingProfileRequestChargingProfileKindEnum = "Dynamic"
)

type SetChargingProfileRequestChargingProfilePurposeEnum string

const (
	SetChargingProfileRequestChargingProfilePurposeEnumChargingStationExternalConstraints SetChargingProfileRequestChargingProfilePurposeEnum = "ChargingStationExternalConstraints"
	SetChargingProfileRequestChargingProfilePurposeEnumChargingStationMaxProfile          SetChargingProfileRequestChargingProfilePurposeEnum = "ChargingStationMaxProfile"
	SetChargingProfileRequestChargingProfilePurposeEnumTxDefaultProfile                   SetChargingProfileRequestChargingProfilePurposeEnum = "TxDefaultProfile"
	SetChargingProfileRequestChargingProfilePurposeEnumTxProfile                          SetChargingProfileRequestChargingProfilePurposeEnum = "TxProfile"
	SetChargingProfileRequestChargingProfilePurposeEnumPriorityCharging                   SetChargingProfileRequestChargingProfilePurposeEnum = "PriorityCharging"
	SetChargingProfileRequestChargingProfilePurposeEnumLocalGeneration                    SetChargingProfileRequestChargingProfilePurposeEnum = "LocalGeneration"
)

func (SetChargingProfileRequest) ActionName() string { return "SetChargingProfile" }

func (SetChargingProfileRequest) Version() protocol.Version { return protocol.OCPP21 }

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
