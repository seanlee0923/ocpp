// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ReportChargingProfilesRequest{}

var schemaReportChargingProfilesRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"requestId": &validation.Schema{Type: "integer", AllowAdditional: true}, "chargingLimitSource": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "chargingProfile": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"id": &validation.Schema{Type: "integer", AllowAdditional: true}, "stackLevel": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "chargingProfilePurpose": &validation.Schema{Type: "string", Enum: []string{"ChargingStationExternalConstraints", "ChargingStationMaxProfile", "TxDefaultProfile", "TxProfile", "PriorityCharging", "LocalGeneration"}}, "chargingProfileKind": &validation.Schema{Type: "string", Enum: []string{"Absolute", "Recurring", "Relative", "Dynamic"}}, "recurrencyKind": &validation.Schema{Type: "string", Enum: []string{"Daily", "Weekly"}}, "validFrom": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "validTo": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "transactionId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "maxOfflineDuration": &validation.Schema{Type: "integer", AllowAdditional: true}, "chargingSchedule": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"id": &validation.Schema{Type: "integer", AllowAdditional: true}, "limitAtSoC": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"soc": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true, Maximum: 100, HasMaximum: true}, "limit": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"soc", "limit"}}, "startSchedule": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "duration": &validation.Schema{Type: "integer", AllowAdditional: true}, "chargingRateUnit": &validation.Schema{Type: "string", Enum: []string{"W", "A"}}, "minChargingRate": &validation.Schema{Type: "number", AllowAdditional: true}, "powerTolerance": &validation.Schema{Type: "number", AllowAdditional: true}, "signatureId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "digestValue": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 88, HasMaxLength: true}, "useLocalTime": &validation.Schema{Type: "boolean", AllowAdditional: true}, "chargingSchedulePeriod": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"startPeriod": &validation.Schema{Type: "integer", AllowAdditional: true}, "limit": &validation.Schema{Type: "number", AllowAdditional: true}, "limit_L2": &validation.Schema{Type: "number", AllowAdditional: true}, "limit_L3": &validation.Schema{Type: "number", AllowAdditional: true}, "numberPhases": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true, Maximum: 3, HasMaximum: true}, "phaseToUse": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true, Maximum: 3, HasMaximum: true}, "dischargeLimit": &validation.Schema{Type: "number", AllowAdditional: true, Maximum: 0, HasMaximum: true}, "dischargeLimit_L2": &validation.Schema{Type: "number", AllowAdditional: true, Maximum: 0, HasMaximum: true}, "dischargeLimit_L3": &validation.Schema{Type: "number", AllowAdditional: true, Maximum: 0, HasMaximum: true}, "setpoint": &validation.Schema{Type: "number", AllowAdditional: true}, "setpoint_L2": &validation.Schema{Type: "number", AllowAdditional: true}, "setpoint_L3": &validation.Schema{Type: "number", AllowAdditional: true}, "setpointReactive": &validation.Schema{Type: "number", AllowAdditional: true}, "setpointReactive_L2": &validation.Schema{Type: "number", AllowAdditional: true}, "setpointReactive_L3": &validation.Schema{Type: "number", AllowAdditional: true}, "preconditioningRequest": &validation.Schema{Type: "boolean", AllowAdditional: true}, "evseSleep": &validation.Schema{Type: "boolean", AllowAdditional: true}, "v2xBaseline": &validation.Schema{Type: "number", AllowAdditional: true}, "operationMode": &validation.Schema{Type: "string", Enum: []string{"Idle", "ChargingOnly", "CentralSetpoint", "ExternalSetpoint", "ExternalLimits", "CentralFrequency", "LocalFrequency", "LocalLoadBalancing"}}, "v2xFreqWattCurve": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"frequency": &validation.Schema{Type: "number", AllowAdditional: true}, "power": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"frequency", "power"}}, MinItems: 1, HasMinItems: true, MaxItems: 20, HasMaxItems: true}, "v2xSignalWattCurve": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"signal": &validation.Schema{Type: "integer", AllowAdditional: true}, "power": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"signal", "power"}}, MinItems: 1, HasMinItems: true, MaxItems: 20, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"startPeriod"}}, MinItems: 1, HasMinItems: true, MaxItems: 1024, HasMaxItems: true}, "randomizedDelay": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "salesTariff": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"id": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "salesTariffDescription": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 32, HasMaxLength: true}, "numEPriceLevels": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "salesTariffEntry": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"relativeTimeInterval": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"start": &validation.Schema{Type: "integer", AllowAdditional: true}, "duration": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"start"}}, "ePriceLevel": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "consumptionCost": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"startValue": &validation.Schema{Type: "number", AllowAdditional: true}, "cost": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"costKind": &validation.Schema{Type: "string", Enum: []string{"CarbonDioxideEmission", "RelativePricePercentage", "RenewableGenerationPercentage"}}, "amount": &validation.Schema{Type: "integer", AllowAdditional: true}, "amountMultiplier": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"costKind", "amount"}}, MinItems: 1, HasMinItems: true, MaxItems: 3, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"startValue", "cost"}}, MinItems: 1, HasMinItems: true, MaxItems: 3, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"relativeTimeInterval"}}, MinItems: 1, HasMinItems: true, MaxItems: 1024, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id", "salesTariffEntry"}}, "absolutePriceSchedule": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"timeAnchor": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "priceScheduleID": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "priceScheduleDescription": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 160, HasMaxLength: true}, "currency": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 3, HasMaxLength: true}, "language": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 8, HasMaxLength: true}, "priceAlgorithm": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 2000, HasMaxLength: true}, "minimumCost": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exponent": &validation.Schema{Type: "integer", AllowAdditional: true}, "value": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"exponent", "value"}}, "maximumCost": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exponent": &validation.Schema{Type: "integer", AllowAdditional: true}, "value": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"exponent", "value"}}, "priceRuleStacks": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"duration": &validation.Schema{Type: "integer", AllowAdditional: true}, "priceRule": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"parkingFeePeriod": &validation.Schema{Type: "integer", AllowAdditional: true}, "carbonDioxideEmission": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "renewableGenerationPercentage": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true, Maximum: 100, HasMaximum: true}, "energyFee": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exponent": &validation.Schema{Type: "integer", AllowAdditional: true}, "value": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"exponent", "value"}}, "parkingFee": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exponent": &validation.Schema{Type: "integer", AllowAdditional: true}, "value": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"exponent", "value"}}, "powerRangeStart": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exponent": &validation.Schema{Type: "integer", AllowAdditional: true}, "value": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"exponent", "value"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"energyFee", "powerRangeStart"}}, MinItems: 1, HasMinItems: true, MaxItems: 8, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"duration", "priceRule"}}, MinItems: 1, HasMinItems: true, MaxItems: 1024, HasMaxItems: true}, "taxRules": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"taxRuleID": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "taxRuleName": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 100, HasMaxLength: true}, "taxIncludedInPrice": &validation.Schema{Type: "boolean", AllowAdditional: true}, "appliesToEnergyFee": &validation.Schema{Type: "boolean", AllowAdditional: true}, "appliesToParkingFee": &validation.Schema{Type: "boolean", AllowAdditional: true}, "appliesToOverstayFee": &validation.Schema{Type: "boolean", AllowAdditional: true}, "appliesToMinimumMaximumCost": &validation.Schema{Type: "boolean", AllowAdditional: true}, "taxRate": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exponent": &validation.Schema{Type: "integer", AllowAdditional: true}, "value": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"exponent", "value"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"taxRuleID", "appliesToEnergyFee", "appliesToParkingFee", "appliesToOverstayFee", "appliesToMinimumMaximumCost", "taxRate"}}, MinItems: 1, HasMinItems: true, MaxItems: 10, HasMaxItems: true}, "overstayRuleList": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"overstayPowerThreshold": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exponent": &validation.Schema{Type: "integer", AllowAdditional: true}, "value": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"exponent", "value"}}, "overstayRule": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"overstayFee": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exponent": &validation.Schema{Type: "integer", AllowAdditional: true}, "value": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"exponent", "value"}}, "overstayRuleDescription": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 32, HasMaxLength: true}, "startTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "overstayFeePeriod": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"startTime", "overstayFeePeriod", "overstayFee"}}, MinItems: 1, HasMinItems: true, MaxItems: 5, HasMaxItems: true}, "overstayTimeThreshold": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"overstayRule"}}, "additionalSelectedServices": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"serviceFee": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exponent": &validation.Schema{Type: "integer", AllowAdditional: true}, "value": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"exponent", "value"}}, "serviceName": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 80, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"serviceName", "serviceFee"}}, MinItems: 1, HasMinItems: true, MaxItems: 5, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"timeAnchor", "priceScheduleID", "currency", "language", "priceAlgorithm", "priceRuleStacks"}}, "priceLevelSchedule": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"priceLevelScheduleEntries": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"duration": &validation.Schema{Type: "integer", AllowAdditional: true}, "priceLevel": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"duration", "priceLevel"}}, MinItems: 1, HasMinItems: true, MaxItems: 100, HasMaxItems: true}, "timeAnchor": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "priceScheduleId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "priceScheduleDescription": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 32, HasMaxLength: true}, "numberOfPriceLevels": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"timeAnchor", "priceScheduleId", "numberOfPriceLevels", "priceLevelScheduleEntries"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id", "chargingRateUnit", "chargingSchedulePeriod"}}, MinItems: 1, HasMinItems: true, MaxItems: 3, HasMaxItems: true}, "invalidAfterOfflineDuration": &validation.Schema{Type: "boolean", AllowAdditional: true}, "dynUpdateInterval": &validation.Schema{Type: "integer", AllowAdditional: true}, "dynUpdateTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "priceScheduleSignature": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 256, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id", "stackLevel", "chargingProfilePurpose", "chargingProfileKind", "chargingSchedule"}}, MinItems: 1, HasMinItems: true}, "tbc": &validation.Schema{Type: "boolean", AllowAdditional: true}, "evseId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"requestId", "chargingLimitSource", "evseId", "chargingProfile"}}

type ReportChargingProfilesRequest struct {
	RequestID           int                                            `json:"requestId"`
	ChargingLimitSource string                                         `json:"chargingLimitSource"`
	ChargingProfile     []ReportChargingProfilesRequestChargingProfile `json:"chargingProfile"`
	Tbc                 *bool                                          `json:"tbc,omitempty"`
	EVSEID              int                                            `json:"evseId"`
	CustomData          *ReportChargingProfilesRequestCustomData       `json:"customData,omitempty"`
}

type ReportChargingProfilesRequestChargingProfile struct {
	ID                          int                                                     `json:"id"`
	StackLevel                  int                                                     `json:"stackLevel"`
	ChargingProfilePurpose      ReportChargingProfilesRequestChargingProfilePurposeEnum `json:"chargingProfilePurpose"`
	ChargingProfileKind         ReportChargingProfilesRequestChargingProfileKindEnum    `json:"chargingProfileKind"`
	RecurrencyKind              *ReportChargingProfilesRequestRecurrencyKindEnum        `json:"recurrencyKind,omitempty"`
	ValidFrom                   *string                                                 `json:"validFrom,omitempty"`
	ValidTo                     *string                                                 `json:"validTo,omitempty"`
	TransactionID               *string                                                 `json:"transactionId,omitempty"`
	MaxOfflineDuration          *int                                                    `json:"maxOfflineDuration,omitempty"`
	ChargingSchedule            []ReportChargingProfilesRequestChargingSchedule         `json:"chargingSchedule"`
	InvalidAfterOfflineDuration *bool                                                   `json:"invalidAfterOfflineDuration,omitempty"`
	DynUpdateInterval           *int                                                    `json:"dynUpdateInterval,omitempty"`
	DynUpdateTime               *string                                                 `json:"dynUpdateTime,omitempty"`
	PriceScheduleSignature      *string                                                 `json:"priceScheduleSignature,omitempty"`
	CustomData                  *ReportChargingProfilesRequestCustomData                `json:"customData,omitempty"`
}

type ReportChargingProfilesRequestChargingSchedule struct {
	ID                     int                                                   `json:"id"`
	LimitAtSoC             *ReportChargingProfilesRequestLimitAtSoC              `json:"limitAtSoC,omitempty"`
	StartSchedule          *string                                               `json:"startSchedule,omitempty"`
	Duration               *int                                                  `json:"duration,omitempty"`
	ChargingRateUnit       ReportChargingProfilesRequestChargingRateUnitEnum     `json:"chargingRateUnit"`
	MinChargingRate        *float64                                              `json:"minChargingRate,omitempty"`
	PowerTolerance         *float64                                              `json:"powerTolerance,omitempty"`
	SignatureID            *int                                                  `json:"signatureId,omitempty"`
	DigestValue            *string                                               `json:"digestValue,omitempty"`
	UseLocalTime           *bool                                                 `json:"useLocalTime,omitempty"`
	ChargingSchedulePeriod []ReportChargingProfilesRequestChargingSchedulePeriod `json:"chargingSchedulePeriod"`
	RandomizedDelay        *int                                                  `json:"randomizedDelay,omitempty"`
	SalesTariff            *ReportChargingProfilesRequestSalesTariff             `json:"salesTariff,omitempty"`
	AbsolutePriceSchedule  *ReportChargingProfilesRequestAbsolutePriceSchedule   `json:"absolutePriceSchedule,omitempty"`
	PriceLevelSchedule     *ReportChargingProfilesRequestPriceLevelSchedule      `json:"priceLevelSchedule,omitempty"`
	CustomData             *ReportChargingProfilesRequestCustomData              `json:"customData,omitempty"`
}

type ReportChargingProfilesRequestPriceLevelSchedule struct {
	PriceLevelScheduleEntries []ReportChargingProfilesRequestPriceLevelScheduleEntry `json:"priceLevelScheduleEntries"`
	TimeAnchor                string                                                 `json:"timeAnchor"`
	PriceScheduleID           int                                                    `json:"priceScheduleId"`
	PriceScheduleDescription  *string                                                `json:"priceScheduleDescription,omitempty"`
	NumberOfPriceLevels       int                                                    `json:"numberOfPriceLevels"`
	CustomData                *ReportChargingProfilesRequestCustomData               `json:"customData,omitempty"`
}

type ReportChargingProfilesRequestPriceLevelScheduleEntry struct {
	Duration   int                                      `json:"duration"`
	PriceLevel int                                      `json:"priceLevel"`
	CustomData *ReportChargingProfilesRequestCustomData `json:"customData,omitempty"`
}

type ReportChargingProfilesRequestAbsolutePriceSchedule struct {
	TimeAnchor                 string                                                    `json:"timeAnchor"`
	PriceScheduleID            int                                                       `json:"priceScheduleID"`
	PriceScheduleDescription   *string                                                   `json:"priceScheduleDescription,omitempty"`
	Currency                   string                                                    `json:"currency"`
	Language                   string                                                    `json:"language"`
	PriceAlgorithm             string                                                    `json:"priceAlgorithm"`
	MinimumCost                *ReportChargingProfilesRequestRationalNumber              `json:"minimumCost,omitempty"`
	MaximumCost                *ReportChargingProfilesRequestRationalNumber              `json:"maximumCost,omitempty"`
	PriceRuleStacks            []ReportChargingProfilesRequestPriceRuleStack             `json:"priceRuleStacks"`
	TaxRules                   []ReportChargingProfilesRequestTaxRule                    `json:"taxRules,omitempty"`
	OverstayRuleList           *ReportChargingProfilesRequestOverstayRuleList            `json:"overstayRuleList,omitempty"`
	AdditionalSelectedServices []ReportChargingProfilesRequestAdditionalSelectedServices `json:"additionalSelectedServices,omitempty"`
	CustomData                 *ReportChargingProfilesRequestCustomData                  `json:"customData,omitempty"`
}

type ReportChargingProfilesRequestAdditionalSelectedServices struct {
	ServiceFee  ReportChargingProfilesRequestRationalNumber `json:"serviceFee"`
	ServiceName string                                      `json:"serviceName"`
	CustomData  *ReportChargingProfilesRequestCustomData    `json:"customData,omitempty"`
}

type ReportChargingProfilesRequestOverstayRuleList struct {
	OverstayPowerThreshold *ReportChargingProfilesRequestRationalNumber `json:"overstayPowerThreshold,omitempty"`
	OverstayRule           []ReportChargingProfilesRequestOverstayRule  `json:"overstayRule"`
	OverstayTimeThreshold  *int                                         `json:"overstayTimeThreshold,omitempty"`
	CustomData             *ReportChargingProfilesRequestCustomData     `json:"customData,omitempty"`
}

type ReportChargingProfilesRequestOverstayRule struct {
	OverstayFee             ReportChargingProfilesRequestRationalNumber `json:"overstayFee"`
	OverstayRuleDescription *string                                     `json:"overstayRuleDescription,omitempty"`
	StartTime               int                                         `json:"startTime"`
	OverstayFeePeriod       int                                         `json:"overstayFeePeriod"`
	CustomData              *ReportChargingProfilesRequestCustomData    `json:"customData,omitempty"`
}

type ReportChargingProfilesRequestTaxRule struct {
	TaxRuleID                   int                                         `json:"taxRuleID"`
	TaxRuleName                 *string                                     `json:"taxRuleName,omitempty"`
	TaxIncludedInPrice          *bool                                       `json:"taxIncludedInPrice,omitempty"`
	AppliesToEnergyFee          bool                                        `json:"appliesToEnergyFee"`
	AppliesToParkingFee         bool                                        `json:"appliesToParkingFee"`
	AppliesToOverstayFee        bool                                        `json:"appliesToOverstayFee"`
	AppliesToMinimumMaximumCost bool                                        `json:"appliesToMinimumMaximumCost"`
	TaxRate                     ReportChargingProfilesRequestRationalNumber `json:"taxRate"`
	CustomData                  *ReportChargingProfilesRequestCustomData    `json:"customData,omitempty"`
}

type ReportChargingProfilesRequestPriceRuleStack struct {
	Duration   int                                      `json:"duration"`
	PriceRule  []ReportChargingProfilesRequestPriceRule `json:"priceRule"`
	CustomData *ReportChargingProfilesRequestCustomData `json:"customData,omitempty"`
}

type ReportChargingProfilesRequestPriceRule struct {
	ParkingFeePeriod              *int                                         `json:"parkingFeePeriod,omitempty"`
	CarbonDioxideEmission         *int                                         `json:"carbonDioxideEmission,omitempty"`
	RenewableGenerationPercentage *int                                         `json:"renewableGenerationPercentage,omitempty"`
	EnergyFee                     ReportChargingProfilesRequestRationalNumber  `json:"energyFee"`
	ParkingFee                    *ReportChargingProfilesRequestRationalNumber `json:"parkingFee,omitempty"`
	PowerRangeStart               ReportChargingProfilesRequestRationalNumber  `json:"powerRangeStart"`
	CustomData                    *ReportChargingProfilesRequestCustomData     `json:"customData,omitempty"`
}

type ReportChargingProfilesRequestRationalNumber struct {
	Exponent   int                                      `json:"exponent"`
	Value      int                                      `json:"value"`
	CustomData *ReportChargingProfilesRequestCustomData `json:"customData,omitempty"`
}

type ReportChargingProfilesRequestSalesTariff struct {
	ID                     int                                             `json:"id"`
	SalesTariffDescription *string                                         `json:"salesTariffDescription,omitempty"`
	NumEPriceLevels        *int                                            `json:"numEPriceLevels,omitempty"`
	SalesTariffEntry       []ReportChargingProfilesRequestSalesTariffEntry `json:"salesTariffEntry"`
	CustomData             *ReportChargingProfilesRequestCustomData        `json:"customData,omitempty"`
}

type ReportChargingProfilesRequestSalesTariffEntry struct {
	RelativeTimeInterval ReportChargingProfilesRequestRelativeTimeInterval `json:"relativeTimeInterval"`
	EPriceLevel          *int                                              `json:"ePriceLevel,omitempty"`
	ConsumptionCost      []ReportChargingProfilesRequestConsumptionCost    `json:"consumptionCost,omitempty"`
	CustomData           *ReportChargingProfilesRequestCustomData          `json:"customData,omitempty"`
}

type ReportChargingProfilesRequestConsumptionCost struct {
	StartValue float64                                  `json:"startValue"`
	Cost       []ReportChargingProfilesRequestCost      `json:"cost"`
	CustomData *ReportChargingProfilesRequestCustomData `json:"customData,omitempty"`
}

type ReportChargingProfilesRequestCost struct {
	CostKind         ReportChargingProfilesRequestCostKindEnum `json:"costKind"`
	Amount           int                                       `json:"amount"`
	AmountMultiplier *int                                      `json:"amountMultiplier,omitempty"`
	CustomData       *ReportChargingProfilesRequestCustomData  `json:"customData,omitempty"`
}

type ReportChargingProfilesRequestCostKindEnum string

const (
	ReportChargingProfilesRequestCostKindEnumCarbonDioxideEmission         ReportChargingProfilesRequestCostKindEnum = "CarbonDioxideEmission"
	ReportChargingProfilesRequestCostKindEnumRelativePricePercentage       ReportChargingProfilesRequestCostKindEnum = "RelativePricePercentage"
	ReportChargingProfilesRequestCostKindEnumRenewableGenerationPercentage ReportChargingProfilesRequestCostKindEnum = "RenewableGenerationPercentage"
)

type ReportChargingProfilesRequestRelativeTimeInterval struct {
	Start      int                                      `json:"start"`
	Duration   *int                                     `json:"duration,omitempty"`
	CustomData *ReportChargingProfilesRequestCustomData `json:"customData,omitempty"`
}

type ReportChargingProfilesRequestChargingSchedulePeriod struct {
	StartPeriod            int                                               `json:"startPeriod"`
	Limit                  *float64                                          `json:"limit,omitempty"`
	LimitL2                *float64                                          `json:"limit_L2,omitempty"`
	LimitL3                *float64                                          `json:"limit_L3,omitempty"`
	NumberPhases           *int                                              `json:"numberPhases,omitempty"`
	PhaseToUse             *int                                              `json:"phaseToUse,omitempty"`
	DischargeLimit         *float64                                          `json:"dischargeLimit,omitempty"`
	DischargeLimitL2       *float64                                          `json:"dischargeLimit_L2,omitempty"`
	DischargeLimitL3       *float64                                          `json:"dischargeLimit_L3,omitempty"`
	Setpoint               *float64                                          `json:"setpoint,omitempty"`
	SetpointL2             *float64                                          `json:"setpoint_L2,omitempty"`
	SetpointL3             *float64                                          `json:"setpoint_L3,omitempty"`
	SetpointReactive       *float64                                          `json:"setpointReactive,omitempty"`
	SetpointReactiveL2     *float64                                          `json:"setpointReactive_L2,omitempty"`
	SetpointReactiveL3     *float64                                          `json:"setpointReactive_L3,omitempty"`
	PreconditioningRequest *bool                                             `json:"preconditioningRequest,omitempty"`
	EVSESleep              *bool                                             `json:"evseSleep,omitempty"`
	V2XBaseline            *float64                                          `json:"v2xBaseline,omitempty"`
	OperationMode          *ReportChargingProfilesRequestOperationModeEnum   `json:"operationMode,omitempty"`
	V2XFreqWattCurve       []ReportChargingProfilesRequestV2XFreqWattPoint   `json:"v2xFreqWattCurve,omitempty"`
	V2XSignalWattCurve     []ReportChargingProfilesRequestV2XSignalWattPoint `json:"v2xSignalWattCurve,omitempty"`
	CustomData             *ReportChargingProfilesRequestCustomData          `json:"customData,omitempty"`
}

type ReportChargingProfilesRequestV2XSignalWattPoint struct {
	Signal     int                                      `json:"signal"`
	Power      float64                                  `json:"power"`
	CustomData *ReportChargingProfilesRequestCustomData `json:"customData,omitempty"`
}

type ReportChargingProfilesRequestV2XFreqWattPoint struct {
	Frequency  float64                                  `json:"frequency"`
	Power      float64                                  `json:"power"`
	CustomData *ReportChargingProfilesRequestCustomData `json:"customData,omitempty"`
}

type ReportChargingProfilesRequestOperationModeEnum string

const (
	ReportChargingProfilesRequestOperationModeEnumIdle               ReportChargingProfilesRequestOperationModeEnum = "Idle"
	ReportChargingProfilesRequestOperationModeEnumChargingOnly       ReportChargingProfilesRequestOperationModeEnum = "ChargingOnly"
	ReportChargingProfilesRequestOperationModeEnumCentralSetpoint    ReportChargingProfilesRequestOperationModeEnum = "CentralSetpoint"
	ReportChargingProfilesRequestOperationModeEnumExternalSetpoint   ReportChargingProfilesRequestOperationModeEnum = "ExternalSetpoint"
	ReportChargingProfilesRequestOperationModeEnumExternalLimits     ReportChargingProfilesRequestOperationModeEnum = "ExternalLimits"
	ReportChargingProfilesRequestOperationModeEnumCentralFrequency   ReportChargingProfilesRequestOperationModeEnum = "CentralFrequency"
	ReportChargingProfilesRequestOperationModeEnumLocalFrequency     ReportChargingProfilesRequestOperationModeEnum = "LocalFrequency"
	ReportChargingProfilesRequestOperationModeEnumLocalLoadBalancing ReportChargingProfilesRequestOperationModeEnum = "LocalLoadBalancing"
)

type ReportChargingProfilesRequestChargingRateUnitEnum string

const (
	ReportChargingProfilesRequestChargingRateUnitEnumW ReportChargingProfilesRequestChargingRateUnitEnum = "W"
	ReportChargingProfilesRequestChargingRateUnitEnumA ReportChargingProfilesRequestChargingRateUnitEnum = "A"
)

type ReportChargingProfilesRequestLimitAtSoC struct {
	Soc        int                                      `json:"soc"`
	Limit      float64                                  `json:"limit"`
	CustomData *ReportChargingProfilesRequestCustomData `json:"customData,omitempty"`
}

type ReportChargingProfilesRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value ReportChargingProfilesRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *ReportChargingProfilesRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

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
	ReportChargingProfilesRequestChargingProfileKindEnumDynamic   ReportChargingProfilesRequestChargingProfileKindEnum = "Dynamic"
)

type ReportChargingProfilesRequestChargingProfilePurposeEnum string

const (
	ReportChargingProfilesRequestChargingProfilePurposeEnumChargingStationExternalConstraints ReportChargingProfilesRequestChargingProfilePurposeEnum = "ChargingStationExternalConstraints"
	ReportChargingProfilesRequestChargingProfilePurposeEnumChargingStationMaxProfile          ReportChargingProfilesRequestChargingProfilePurposeEnum = "ChargingStationMaxProfile"
	ReportChargingProfilesRequestChargingProfilePurposeEnumTxDefaultProfile                   ReportChargingProfilesRequestChargingProfilePurposeEnum = "TxDefaultProfile"
	ReportChargingProfilesRequestChargingProfilePurposeEnumTxProfile                          ReportChargingProfilesRequestChargingProfilePurposeEnum = "TxProfile"
	ReportChargingProfilesRequestChargingProfilePurposeEnumPriorityCharging                   ReportChargingProfilesRequestChargingProfilePurposeEnum = "PriorityCharging"
	ReportChargingProfilesRequestChargingProfilePurposeEnumLocalGeneration                    ReportChargingProfilesRequestChargingProfilePurposeEnum = "LocalGeneration"
)

func (ReportChargingProfilesRequest) ActionName() string { return "ReportChargingProfiles" }

func (ReportChargingProfilesRequest) Version() protocol.Version { return protocol.OCPP21 }

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
