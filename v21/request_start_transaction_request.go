// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = RequestStartTransactionRequest{}

var schemaRequestStartTransactionRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"evseId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 1, HasMinimum: true}, "groupIdToken": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"additionalInfo": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"additionalIdToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}, "type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"additionalIdToken", "type"}}, MinItems: 1, HasMinItems: true}, "idToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}, "type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"idToken", "type"}}, "idToken": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"additionalInfo": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"additionalIdToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}, "type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"additionalIdToken", "type"}}, MinItems: 1, HasMinItems: true}, "idToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}, "type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"idToken", "type"}}, "remoteStartId": &validation.Schema{Type: "integer", AllowAdditional: true}, "chargingProfile": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"id": &validation.Schema{Type: "integer", AllowAdditional: true}, "stackLevel": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "chargingProfilePurpose": &validation.Schema{Type: "string", Enum: []string{"ChargingStationExternalConstraints", "ChargingStationMaxProfile", "TxDefaultProfile", "TxProfile", "PriorityCharging", "LocalGeneration"}}, "chargingProfileKind": &validation.Schema{Type: "string", Enum: []string{"Absolute", "Recurring", "Relative", "Dynamic"}}, "recurrencyKind": &validation.Schema{Type: "string", Enum: []string{"Daily", "Weekly"}}, "validFrom": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "validTo": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "transactionId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "maxOfflineDuration": &validation.Schema{Type: "integer", AllowAdditional: true}, "chargingSchedule": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"id": &validation.Schema{Type: "integer", AllowAdditional: true}, "limitAtSoC": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"soc": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true, Maximum: 100, HasMaximum: true}, "limit": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"soc", "limit"}}, "startSchedule": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "duration": &validation.Schema{Type: "integer", AllowAdditional: true}, "chargingRateUnit": &validation.Schema{Type: "string", Enum: []string{"W", "A"}}, "minChargingRate": &validation.Schema{Type: "number", AllowAdditional: true}, "powerTolerance": &validation.Schema{Type: "number", AllowAdditional: true}, "signatureId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "digestValue": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 88, HasMaxLength: true}, "useLocalTime": &validation.Schema{Type: "boolean", AllowAdditional: true}, "chargingSchedulePeriod": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"startPeriod": &validation.Schema{Type: "integer", AllowAdditional: true}, "limit": &validation.Schema{Type: "number", AllowAdditional: true}, "limit_L2": &validation.Schema{Type: "number", AllowAdditional: true}, "limit_L3": &validation.Schema{Type: "number", AllowAdditional: true}, "numberPhases": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true, Maximum: 3, HasMaximum: true}, "phaseToUse": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true, Maximum: 3, HasMaximum: true}, "dischargeLimit": &validation.Schema{Type: "number", AllowAdditional: true, Maximum: 0, HasMaximum: true}, "dischargeLimit_L2": &validation.Schema{Type: "number", AllowAdditional: true, Maximum: 0, HasMaximum: true}, "dischargeLimit_L3": &validation.Schema{Type: "number", AllowAdditional: true, Maximum: 0, HasMaximum: true}, "setpoint": &validation.Schema{Type: "number", AllowAdditional: true}, "setpoint_L2": &validation.Schema{Type: "number", AllowAdditional: true}, "setpoint_L3": &validation.Schema{Type: "number", AllowAdditional: true}, "setpointReactive": &validation.Schema{Type: "number", AllowAdditional: true}, "setpointReactive_L2": &validation.Schema{Type: "number", AllowAdditional: true}, "setpointReactive_L3": &validation.Schema{Type: "number", AllowAdditional: true}, "preconditioningRequest": &validation.Schema{Type: "boolean", AllowAdditional: true}, "evseSleep": &validation.Schema{Type: "boolean", AllowAdditional: true}, "v2xBaseline": &validation.Schema{Type: "number", AllowAdditional: true}, "operationMode": &validation.Schema{Type: "string", Enum: []string{"Idle", "ChargingOnly", "CentralSetpoint", "ExternalSetpoint", "ExternalLimits", "CentralFrequency", "LocalFrequency", "LocalLoadBalancing"}}, "v2xFreqWattCurve": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"frequency": &validation.Schema{Type: "number", AllowAdditional: true}, "power": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"frequency", "power"}}, MinItems: 1, HasMinItems: true, MaxItems: 20, HasMaxItems: true}, "v2xSignalWattCurve": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"signal": &validation.Schema{Type: "integer", AllowAdditional: true}, "power": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"signal", "power"}}, MinItems: 1, HasMinItems: true, MaxItems: 20, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"startPeriod"}}, MinItems: 1, HasMinItems: true, MaxItems: 1024, HasMaxItems: true}, "randomizedDelay": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "salesTariff": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"id": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "salesTariffDescription": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 32, HasMaxLength: true}, "numEPriceLevels": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "salesTariffEntry": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"relativeTimeInterval": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"start": &validation.Schema{Type: "integer", AllowAdditional: true}, "duration": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"start"}}, "ePriceLevel": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "consumptionCost": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"startValue": &validation.Schema{Type: "number", AllowAdditional: true}, "cost": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"costKind": &validation.Schema{Type: "string", Enum: []string{"CarbonDioxideEmission", "RelativePricePercentage", "RenewableGenerationPercentage"}}, "amount": &validation.Schema{Type: "integer", AllowAdditional: true}, "amountMultiplier": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"costKind", "amount"}}, MinItems: 1, HasMinItems: true, MaxItems: 3, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"startValue", "cost"}}, MinItems: 1, HasMinItems: true, MaxItems: 3, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"relativeTimeInterval"}}, MinItems: 1, HasMinItems: true, MaxItems: 1024, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id", "salesTariffEntry"}}, "absolutePriceSchedule": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"timeAnchor": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "priceScheduleID": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "priceScheduleDescription": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 160, HasMaxLength: true}, "currency": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 3, HasMaxLength: true}, "language": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 8, HasMaxLength: true}, "priceAlgorithm": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 2000, HasMaxLength: true}, "minimumCost": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exponent": &validation.Schema{Type: "integer", AllowAdditional: true}, "value": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"exponent", "value"}}, "maximumCost": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exponent": &validation.Schema{Type: "integer", AllowAdditional: true}, "value": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"exponent", "value"}}, "priceRuleStacks": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"duration": &validation.Schema{Type: "integer", AllowAdditional: true}, "priceRule": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"parkingFeePeriod": &validation.Schema{Type: "integer", AllowAdditional: true}, "carbonDioxideEmission": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "renewableGenerationPercentage": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true, Maximum: 100, HasMaximum: true}, "energyFee": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exponent": &validation.Schema{Type: "integer", AllowAdditional: true}, "value": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"exponent", "value"}}, "parkingFee": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exponent": &validation.Schema{Type: "integer", AllowAdditional: true}, "value": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"exponent", "value"}}, "powerRangeStart": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exponent": &validation.Schema{Type: "integer", AllowAdditional: true}, "value": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"exponent", "value"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"energyFee", "powerRangeStart"}}, MinItems: 1, HasMinItems: true, MaxItems: 8, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"duration", "priceRule"}}, MinItems: 1, HasMinItems: true, MaxItems: 1024, HasMaxItems: true}, "taxRules": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"taxRuleID": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "taxRuleName": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 100, HasMaxLength: true}, "taxIncludedInPrice": &validation.Schema{Type: "boolean", AllowAdditional: true}, "appliesToEnergyFee": &validation.Schema{Type: "boolean", AllowAdditional: true}, "appliesToParkingFee": &validation.Schema{Type: "boolean", AllowAdditional: true}, "appliesToOverstayFee": &validation.Schema{Type: "boolean", AllowAdditional: true}, "appliesToMinimumMaximumCost": &validation.Schema{Type: "boolean", AllowAdditional: true}, "taxRate": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exponent": &validation.Schema{Type: "integer", AllowAdditional: true}, "value": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"exponent", "value"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"taxRuleID", "appliesToEnergyFee", "appliesToParkingFee", "appliesToOverstayFee", "appliesToMinimumMaximumCost", "taxRate"}}, MinItems: 1, HasMinItems: true, MaxItems: 10, HasMaxItems: true}, "overstayRuleList": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"overstayPowerThreshold": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exponent": &validation.Schema{Type: "integer", AllowAdditional: true}, "value": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"exponent", "value"}}, "overstayRule": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"overstayFee": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exponent": &validation.Schema{Type: "integer", AllowAdditional: true}, "value": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"exponent", "value"}}, "overstayRuleDescription": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 32, HasMaxLength: true}, "startTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "overstayFeePeriod": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"startTime", "overstayFeePeriod", "overstayFee"}}, MinItems: 1, HasMinItems: true, MaxItems: 5, HasMaxItems: true}, "overstayTimeThreshold": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"overstayRule"}}, "additionalSelectedServices": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"serviceFee": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exponent": &validation.Schema{Type: "integer", AllowAdditional: true}, "value": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"exponent", "value"}}, "serviceName": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 80, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"serviceName", "serviceFee"}}, MinItems: 1, HasMinItems: true, MaxItems: 5, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"timeAnchor", "priceScheduleID", "currency", "language", "priceAlgorithm", "priceRuleStacks"}}, "priceLevelSchedule": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"priceLevelScheduleEntries": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"duration": &validation.Schema{Type: "integer", AllowAdditional: true}, "priceLevel": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"duration", "priceLevel"}}, MinItems: 1, HasMinItems: true, MaxItems: 100, HasMaxItems: true}, "timeAnchor": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "priceScheduleId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "priceScheduleDescription": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 32, HasMaxLength: true}, "numberOfPriceLevels": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"timeAnchor", "priceScheduleId", "numberOfPriceLevels", "priceLevelScheduleEntries"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id", "chargingRateUnit", "chargingSchedulePeriod"}}, MinItems: 1, HasMinItems: true, MaxItems: 3, HasMaxItems: true}, "invalidAfterOfflineDuration": &validation.Schema{Type: "boolean", AllowAdditional: true}, "dynUpdateInterval": &validation.Schema{Type: "integer", AllowAdditional: true}, "dynUpdateTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "priceScheduleSignature": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 256, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id", "stackLevel", "chargingProfilePurpose", "chargingProfileKind", "chargingSchedule"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"remoteStartId", "idToken"}}

type RequestStartTransactionRequest struct {
	EVSEID          *int                                           `json:"evseId,omitempty"`
	GroupIDToken    *RequestStartTransactionRequestIDToken         `json:"groupIdToken,omitempty"`
	IDToken         RequestStartTransactionRequestIDToken          `json:"idToken"`
	RemoteStartID   int                                            `json:"remoteStartId"`
	ChargingProfile *RequestStartTransactionRequestChargingProfile `json:"chargingProfile,omitempty"`
	CustomData      *RequestStartTransactionRequestCustomData      `json:"customData,omitempty"`
}

type RequestStartTransactionRequestChargingProfile struct {
	ID                          int                                                      `json:"id"`
	StackLevel                  int                                                      `json:"stackLevel"`
	ChargingProfilePurpose      RequestStartTransactionRequestChargingProfilePurposeEnum `json:"chargingProfilePurpose"`
	ChargingProfileKind         RequestStartTransactionRequestChargingProfileKindEnum    `json:"chargingProfileKind"`
	RecurrencyKind              *RequestStartTransactionRequestRecurrencyKindEnum        `json:"recurrencyKind,omitempty"`
	ValidFrom                   *string                                                  `json:"validFrom,omitempty"`
	ValidTo                     *string                                                  `json:"validTo,omitempty"`
	TransactionID               *string                                                  `json:"transactionId,omitempty"`
	MaxOfflineDuration          *int                                                     `json:"maxOfflineDuration,omitempty"`
	ChargingSchedule            []RequestStartTransactionRequestChargingSchedule         `json:"chargingSchedule"`
	InvalidAfterOfflineDuration *bool                                                    `json:"invalidAfterOfflineDuration,omitempty"`
	DynUpdateInterval           *int                                                     `json:"dynUpdateInterval,omitempty"`
	DynUpdateTime               *string                                                  `json:"dynUpdateTime,omitempty"`
	PriceScheduleSignature      *string                                                  `json:"priceScheduleSignature,omitempty"`
	CustomData                  *RequestStartTransactionRequestCustomData                `json:"customData,omitempty"`
}

type RequestStartTransactionRequestChargingSchedule struct {
	ID                     int                                                    `json:"id"`
	LimitAtSoC             *RequestStartTransactionRequestLimitAtSoC              `json:"limitAtSoC,omitempty"`
	StartSchedule          *string                                                `json:"startSchedule,omitempty"`
	Duration               *int                                                   `json:"duration,omitempty"`
	ChargingRateUnit       RequestStartTransactionRequestChargingRateUnitEnum     `json:"chargingRateUnit"`
	MinChargingRate        *float64                                               `json:"minChargingRate,omitempty"`
	PowerTolerance         *float64                                               `json:"powerTolerance,omitempty"`
	SignatureID            *int                                                   `json:"signatureId,omitempty"`
	DigestValue            *string                                                `json:"digestValue,omitempty"`
	UseLocalTime           *bool                                                  `json:"useLocalTime,omitempty"`
	ChargingSchedulePeriod []RequestStartTransactionRequestChargingSchedulePeriod `json:"chargingSchedulePeriod"`
	RandomizedDelay        *int                                                   `json:"randomizedDelay,omitempty"`
	SalesTariff            *RequestStartTransactionRequestSalesTariff             `json:"salesTariff,omitempty"`
	AbsolutePriceSchedule  *RequestStartTransactionRequestAbsolutePriceSchedule   `json:"absolutePriceSchedule,omitempty"`
	PriceLevelSchedule     *RequestStartTransactionRequestPriceLevelSchedule      `json:"priceLevelSchedule,omitempty"`
	CustomData             *RequestStartTransactionRequestCustomData              `json:"customData,omitempty"`
}

type RequestStartTransactionRequestPriceLevelSchedule struct {
	PriceLevelScheduleEntries []RequestStartTransactionRequestPriceLevelScheduleEntry `json:"priceLevelScheduleEntries"`
	TimeAnchor                string                                                  `json:"timeAnchor"`
	PriceScheduleID           int                                                     `json:"priceScheduleId"`
	PriceScheduleDescription  *string                                                 `json:"priceScheduleDescription,omitempty"`
	NumberOfPriceLevels       int                                                     `json:"numberOfPriceLevels"`
	CustomData                *RequestStartTransactionRequestCustomData               `json:"customData,omitempty"`
}

type RequestStartTransactionRequestPriceLevelScheduleEntry struct {
	Duration   int                                       `json:"duration"`
	PriceLevel int                                       `json:"priceLevel"`
	CustomData *RequestStartTransactionRequestCustomData `json:"customData,omitempty"`
}

type RequestStartTransactionRequestAbsolutePriceSchedule struct {
	TimeAnchor                 string                                                     `json:"timeAnchor"`
	PriceScheduleID            int                                                        `json:"priceScheduleID"`
	PriceScheduleDescription   *string                                                    `json:"priceScheduleDescription,omitempty"`
	Currency                   string                                                     `json:"currency"`
	Language                   string                                                     `json:"language"`
	PriceAlgorithm             string                                                     `json:"priceAlgorithm"`
	MinimumCost                *RequestStartTransactionRequestRationalNumber              `json:"minimumCost,omitempty"`
	MaximumCost                *RequestStartTransactionRequestRationalNumber              `json:"maximumCost,omitempty"`
	PriceRuleStacks            []RequestStartTransactionRequestPriceRuleStack             `json:"priceRuleStacks"`
	TaxRules                   []RequestStartTransactionRequestTaxRule                    `json:"taxRules,omitempty"`
	OverstayRuleList           *RequestStartTransactionRequestOverstayRuleList            `json:"overstayRuleList,omitempty"`
	AdditionalSelectedServices []RequestStartTransactionRequestAdditionalSelectedServices `json:"additionalSelectedServices,omitempty"`
	CustomData                 *RequestStartTransactionRequestCustomData                  `json:"customData,omitempty"`
}

type RequestStartTransactionRequestAdditionalSelectedServices struct {
	ServiceFee  RequestStartTransactionRequestRationalNumber `json:"serviceFee"`
	ServiceName string                                       `json:"serviceName"`
	CustomData  *RequestStartTransactionRequestCustomData    `json:"customData,omitempty"`
}

type RequestStartTransactionRequestOverstayRuleList struct {
	OverstayPowerThreshold *RequestStartTransactionRequestRationalNumber `json:"overstayPowerThreshold,omitempty"`
	OverstayRule           []RequestStartTransactionRequestOverstayRule  `json:"overstayRule"`
	OverstayTimeThreshold  *int                                          `json:"overstayTimeThreshold,omitempty"`
	CustomData             *RequestStartTransactionRequestCustomData     `json:"customData,omitempty"`
}

type RequestStartTransactionRequestOverstayRule struct {
	OverstayFee             RequestStartTransactionRequestRationalNumber `json:"overstayFee"`
	OverstayRuleDescription *string                                      `json:"overstayRuleDescription,omitempty"`
	StartTime               int                                          `json:"startTime"`
	OverstayFeePeriod       int                                          `json:"overstayFeePeriod"`
	CustomData              *RequestStartTransactionRequestCustomData    `json:"customData,omitempty"`
}

type RequestStartTransactionRequestTaxRule struct {
	TaxRuleID                   int                                          `json:"taxRuleID"`
	TaxRuleName                 *string                                      `json:"taxRuleName,omitempty"`
	TaxIncludedInPrice          *bool                                        `json:"taxIncludedInPrice,omitempty"`
	AppliesToEnergyFee          bool                                         `json:"appliesToEnergyFee"`
	AppliesToParkingFee         bool                                         `json:"appliesToParkingFee"`
	AppliesToOverstayFee        bool                                         `json:"appliesToOverstayFee"`
	AppliesToMinimumMaximumCost bool                                         `json:"appliesToMinimumMaximumCost"`
	TaxRate                     RequestStartTransactionRequestRationalNumber `json:"taxRate"`
	CustomData                  *RequestStartTransactionRequestCustomData    `json:"customData,omitempty"`
}

type RequestStartTransactionRequestPriceRuleStack struct {
	Duration   int                                       `json:"duration"`
	PriceRule  []RequestStartTransactionRequestPriceRule `json:"priceRule"`
	CustomData *RequestStartTransactionRequestCustomData `json:"customData,omitempty"`
}

type RequestStartTransactionRequestPriceRule struct {
	ParkingFeePeriod              *int                                          `json:"parkingFeePeriod,omitempty"`
	CarbonDioxideEmission         *int                                          `json:"carbonDioxideEmission,omitempty"`
	RenewableGenerationPercentage *int                                          `json:"renewableGenerationPercentage,omitempty"`
	EnergyFee                     RequestStartTransactionRequestRationalNumber  `json:"energyFee"`
	ParkingFee                    *RequestStartTransactionRequestRationalNumber `json:"parkingFee,omitempty"`
	PowerRangeStart               RequestStartTransactionRequestRationalNumber  `json:"powerRangeStart"`
	CustomData                    *RequestStartTransactionRequestCustomData     `json:"customData,omitempty"`
}

type RequestStartTransactionRequestRationalNumber struct {
	Exponent   int                                       `json:"exponent"`
	Value      int                                       `json:"value"`
	CustomData *RequestStartTransactionRequestCustomData `json:"customData,omitempty"`
}

type RequestStartTransactionRequestSalesTariff struct {
	ID                     int                                              `json:"id"`
	SalesTariffDescription *string                                          `json:"salesTariffDescription,omitempty"`
	NumEPriceLevels        *int                                             `json:"numEPriceLevels,omitempty"`
	SalesTariffEntry       []RequestStartTransactionRequestSalesTariffEntry `json:"salesTariffEntry"`
	CustomData             *RequestStartTransactionRequestCustomData        `json:"customData,omitempty"`
}

type RequestStartTransactionRequestSalesTariffEntry struct {
	RelativeTimeInterval RequestStartTransactionRequestRelativeTimeInterval `json:"relativeTimeInterval"`
	EPriceLevel          *int                                               `json:"ePriceLevel,omitempty"`
	ConsumptionCost      []RequestStartTransactionRequestConsumptionCost    `json:"consumptionCost,omitempty"`
	CustomData           *RequestStartTransactionRequestCustomData          `json:"customData,omitempty"`
}

type RequestStartTransactionRequestConsumptionCost struct {
	StartValue float64                                   `json:"startValue"`
	Cost       []RequestStartTransactionRequestCost      `json:"cost"`
	CustomData *RequestStartTransactionRequestCustomData `json:"customData,omitempty"`
}

type RequestStartTransactionRequestCost struct {
	CostKind         RequestStartTransactionRequestCostKindEnum `json:"costKind"`
	Amount           int                                        `json:"amount"`
	AmountMultiplier *int                                       `json:"amountMultiplier,omitempty"`
	CustomData       *RequestStartTransactionRequestCustomData  `json:"customData,omitempty"`
}

type RequestStartTransactionRequestCostKindEnum string

const (
	RequestStartTransactionRequestCostKindEnumCarbonDioxideEmission         RequestStartTransactionRequestCostKindEnum = "CarbonDioxideEmission"
	RequestStartTransactionRequestCostKindEnumRelativePricePercentage       RequestStartTransactionRequestCostKindEnum = "RelativePricePercentage"
	RequestStartTransactionRequestCostKindEnumRenewableGenerationPercentage RequestStartTransactionRequestCostKindEnum = "RenewableGenerationPercentage"
)

type RequestStartTransactionRequestRelativeTimeInterval struct {
	Start      int                                       `json:"start"`
	Duration   *int                                      `json:"duration,omitempty"`
	CustomData *RequestStartTransactionRequestCustomData `json:"customData,omitempty"`
}

type RequestStartTransactionRequestChargingSchedulePeriod struct {
	StartPeriod            int                                                `json:"startPeriod"`
	Limit                  *float64                                           `json:"limit,omitempty"`
	LimitL2                *float64                                           `json:"limit_L2,omitempty"`
	LimitL3                *float64                                           `json:"limit_L3,omitempty"`
	NumberPhases           *int                                               `json:"numberPhases,omitempty"`
	PhaseToUse             *int                                               `json:"phaseToUse,omitempty"`
	DischargeLimit         *float64                                           `json:"dischargeLimit,omitempty"`
	DischargeLimitL2       *float64                                           `json:"dischargeLimit_L2,omitempty"`
	DischargeLimitL3       *float64                                           `json:"dischargeLimit_L3,omitempty"`
	Setpoint               *float64                                           `json:"setpoint,omitempty"`
	SetpointL2             *float64                                           `json:"setpoint_L2,omitempty"`
	SetpointL3             *float64                                           `json:"setpoint_L3,omitempty"`
	SetpointReactive       *float64                                           `json:"setpointReactive,omitempty"`
	SetpointReactiveL2     *float64                                           `json:"setpointReactive_L2,omitempty"`
	SetpointReactiveL3     *float64                                           `json:"setpointReactive_L3,omitempty"`
	PreconditioningRequest *bool                                              `json:"preconditioningRequest,omitempty"`
	EVSESleep              *bool                                              `json:"evseSleep,omitempty"`
	V2XBaseline            *float64                                           `json:"v2xBaseline,omitempty"`
	OperationMode          *RequestStartTransactionRequestOperationModeEnum   `json:"operationMode,omitempty"`
	V2XFreqWattCurve       []RequestStartTransactionRequestV2XFreqWattPoint   `json:"v2xFreqWattCurve,omitempty"`
	V2XSignalWattCurve     []RequestStartTransactionRequestV2XSignalWattPoint `json:"v2xSignalWattCurve,omitempty"`
	CustomData             *RequestStartTransactionRequestCustomData          `json:"customData,omitempty"`
}

type RequestStartTransactionRequestV2XSignalWattPoint struct {
	Signal     int                                       `json:"signal"`
	Power      float64                                   `json:"power"`
	CustomData *RequestStartTransactionRequestCustomData `json:"customData,omitempty"`
}

type RequestStartTransactionRequestV2XFreqWattPoint struct {
	Frequency  float64                                   `json:"frequency"`
	Power      float64                                   `json:"power"`
	CustomData *RequestStartTransactionRequestCustomData `json:"customData,omitempty"`
}

type RequestStartTransactionRequestOperationModeEnum string

const (
	RequestStartTransactionRequestOperationModeEnumIdle               RequestStartTransactionRequestOperationModeEnum = "Idle"
	RequestStartTransactionRequestOperationModeEnumChargingOnly       RequestStartTransactionRequestOperationModeEnum = "ChargingOnly"
	RequestStartTransactionRequestOperationModeEnumCentralSetpoint    RequestStartTransactionRequestOperationModeEnum = "CentralSetpoint"
	RequestStartTransactionRequestOperationModeEnumExternalSetpoint   RequestStartTransactionRequestOperationModeEnum = "ExternalSetpoint"
	RequestStartTransactionRequestOperationModeEnumExternalLimits     RequestStartTransactionRequestOperationModeEnum = "ExternalLimits"
	RequestStartTransactionRequestOperationModeEnumCentralFrequency   RequestStartTransactionRequestOperationModeEnum = "CentralFrequency"
	RequestStartTransactionRequestOperationModeEnumLocalFrequency     RequestStartTransactionRequestOperationModeEnum = "LocalFrequency"
	RequestStartTransactionRequestOperationModeEnumLocalLoadBalancing RequestStartTransactionRequestOperationModeEnum = "LocalLoadBalancing"
)

type RequestStartTransactionRequestChargingRateUnitEnum string

const (
	RequestStartTransactionRequestChargingRateUnitEnumW RequestStartTransactionRequestChargingRateUnitEnum = "W"
	RequestStartTransactionRequestChargingRateUnitEnumA RequestStartTransactionRequestChargingRateUnitEnum = "A"
)

type RequestStartTransactionRequestLimitAtSoC struct {
	Soc        int                                       `json:"soc"`
	Limit      float64                                   `json:"limit"`
	CustomData *RequestStartTransactionRequestCustomData `json:"customData,omitempty"`
}

type RequestStartTransactionRequestRecurrencyKindEnum string

const (
	RequestStartTransactionRequestRecurrencyKindEnumDaily  RequestStartTransactionRequestRecurrencyKindEnum = "Daily"
	RequestStartTransactionRequestRecurrencyKindEnumWeekly RequestStartTransactionRequestRecurrencyKindEnum = "Weekly"
)

type RequestStartTransactionRequestChargingProfileKindEnum string

const (
	RequestStartTransactionRequestChargingProfileKindEnumAbsolute  RequestStartTransactionRequestChargingProfileKindEnum = "Absolute"
	RequestStartTransactionRequestChargingProfileKindEnumRecurring RequestStartTransactionRequestChargingProfileKindEnum = "Recurring"
	RequestStartTransactionRequestChargingProfileKindEnumRelative  RequestStartTransactionRequestChargingProfileKindEnum = "Relative"
	RequestStartTransactionRequestChargingProfileKindEnumDynamic   RequestStartTransactionRequestChargingProfileKindEnum = "Dynamic"
)

type RequestStartTransactionRequestChargingProfilePurposeEnum string

const (
	RequestStartTransactionRequestChargingProfilePurposeEnumChargingStationExternalConstraints RequestStartTransactionRequestChargingProfilePurposeEnum = "ChargingStationExternalConstraints"
	RequestStartTransactionRequestChargingProfilePurposeEnumChargingStationMaxProfile          RequestStartTransactionRequestChargingProfilePurposeEnum = "ChargingStationMaxProfile"
	RequestStartTransactionRequestChargingProfilePurposeEnumTxDefaultProfile                   RequestStartTransactionRequestChargingProfilePurposeEnum = "TxDefaultProfile"
	RequestStartTransactionRequestChargingProfilePurposeEnumTxProfile                          RequestStartTransactionRequestChargingProfilePurposeEnum = "TxProfile"
	RequestStartTransactionRequestChargingProfilePurposeEnumPriorityCharging                   RequestStartTransactionRequestChargingProfilePurposeEnum = "PriorityCharging"
	RequestStartTransactionRequestChargingProfilePurposeEnumLocalGeneration                    RequestStartTransactionRequestChargingProfilePurposeEnum = "LocalGeneration"
)

type RequestStartTransactionRequestIDToken struct {
	AdditionalInfo []RequestStartTransactionRequestAdditionalInfo `json:"additionalInfo,omitempty"`
	IDToken        string                                         `json:"idToken"`
	Type           string                                         `json:"type"`
	CustomData     *RequestStartTransactionRequestCustomData      `json:"customData,omitempty"`
}

type RequestStartTransactionRequestAdditionalInfo struct {
	AdditionalIDToken string                                    `json:"additionalIdToken"`
	Type              string                                    `json:"type"`
	CustomData        *RequestStartTransactionRequestCustomData `json:"customData,omitempty"`
}

type RequestStartTransactionRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (RequestStartTransactionRequest) ActionName() string { return "RequestStartTransaction" }

func (RequestStartTransactionRequest) Version() protocol.Version { return protocol.OCPP21 }

func (RequestStartTransactionRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (RequestStartTransactionRequest) SchemaName() string {
	return "RequestStartTransactionRequest.json"
}

func (message RequestStartTransactionRequest) Validate() error {
	return validation.Validate("RequestStartTransactionRequest.json", schemaRequestStartTransactionRequest, message)
}

func (RequestStartTransactionRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("RequestStartTransactionRequest.json", schemaRequestStartTransactionRequest, data)
}
