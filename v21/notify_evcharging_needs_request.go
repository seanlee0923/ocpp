// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = NotifyEVChargingNeedsRequest{}

var schemaNotifyEVChargingNeedsRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"evseId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 1, HasMinimum: true}, "maxScheduleTuples": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "chargingNeeds": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"acChargingParameters": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"energyAmount": &validation.Schema{Type: "number", AllowAdditional: true}, "evMinCurrent": &validation.Schema{Type: "number", AllowAdditional: true}, "evMaxCurrent": &validation.Schema{Type: "number", AllowAdditional: true}, "evMaxVoltage": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"energyAmount", "evMinCurrent", "evMaxCurrent", "evMaxVoltage"}}, "derChargingParameters": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"evSupportedDERControl": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", Enum: []string{"EnterService", "FreqDroop", "FreqWatt", "FixedPFAbsorb", "FixedPFInject", "FixedVar", "Gradients", "HFMustTrip", "HFMayTrip", "HVMustTrip", "HVMomCess", "HVMayTrip", "LimitMaxDischarge", "LFMustTrip", "LVMustTrip", "LVMomCess", "LVMayTrip", "PowerMonitoringMustTrip", "VoltVar", "VoltWatt", "WattPF", "WattVar"}}, MinItems: 1, HasMinItems: true}, "evOverExcitedMaxDischargePower": &validation.Schema{Type: "number", AllowAdditional: true}, "evOverExcitedPowerFactor": &validation.Schema{Type: "number", AllowAdditional: true}, "evUnderExcitedMaxDischargePower": &validation.Schema{Type: "number", AllowAdditional: true}, "evUnderExcitedPowerFactor": &validation.Schema{Type: "number", AllowAdditional: true}, "maxApparentPower": &validation.Schema{Type: "number", AllowAdditional: true}, "maxChargeApparentPower": &validation.Schema{Type: "number", AllowAdditional: true}, "maxChargeApparentPower_L2": &validation.Schema{Type: "number", AllowAdditional: true}, "maxChargeApparentPower_L3": &validation.Schema{Type: "number", AllowAdditional: true}, "maxDischargeApparentPower": &validation.Schema{Type: "number", AllowAdditional: true}, "maxDischargeApparentPower_L2": &validation.Schema{Type: "number", AllowAdditional: true}, "maxDischargeApparentPower_L3": &validation.Schema{Type: "number", AllowAdditional: true}, "maxChargeReactivePower": &validation.Schema{Type: "number", AllowAdditional: true}, "maxChargeReactivePower_L2": &validation.Schema{Type: "number", AllowAdditional: true}, "maxChargeReactivePower_L3": &validation.Schema{Type: "number", AllowAdditional: true}, "minChargeReactivePower": &validation.Schema{Type: "number", AllowAdditional: true}, "minChargeReactivePower_L2": &validation.Schema{Type: "number", AllowAdditional: true}, "minChargeReactivePower_L3": &validation.Schema{Type: "number", AllowAdditional: true}, "maxDischargeReactivePower": &validation.Schema{Type: "number", AllowAdditional: true}, "maxDischargeReactivePower_L2": &validation.Schema{Type: "number", AllowAdditional: true}, "maxDischargeReactivePower_L3": &validation.Schema{Type: "number", AllowAdditional: true}, "minDischargeReactivePower": &validation.Schema{Type: "number", AllowAdditional: true}, "minDischargeReactivePower_L2": &validation.Schema{Type: "number", AllowAdditional: true}, "minDischargeReactivePower_L3": &validation.Schema{Type: "number", AllowAdditional: true}, "nominalVoltage": &validation.Schema{Type: "number", AllowAdditional: true}, "nominalVoltageOffset": &validation.Schema{Type: "number", AllowAdditional: true}, "maxNominalVoltage": &validation.Schema{Type: "number", AllowAdditional: true}, "minNominalVoltage": &validation.Schema{Type: "number", AllowAdditional: true}, "evInverterManufacturer": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "evInverterModel": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "evInverterSerialNumber": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "evInverterSwVersion": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "evInverterHwVersion": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "evIslandingDetectionMethod": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", Enum: []string{"NoAntiIslandingSupport", "RoCoF", "UVP_OVP", "UFP_OFP", "VoltageVectorShift", "ZeroCrossingDetection", "OtherPassive", "ImpedanceMeasurement", "ImpedanceAtFrequency", "SlipModeFrequencyShift", "SandiaFrequencyShift", "SandiaVoltageShift", "FrequencyJump", "RCLQFactor", "OtherActive"}}, MinItems: 1, HasMinItems: true}, "evIslandingTripTime": &validation.Schema{Type: "number", AllowAdditional: true}, "evMaximumLevel1DCInjection": &validation.Schema{Type: "number", AllowAdditional: true}, "evDurationLevel1DCInjection": &validation.Schema{Type: "number", AllowAdditional: true}, "evMaximumLevel2DCInjection": &validation.Schema{Type: "number", AllowAdditional: true}, "evDurationLevel2DCInjection": &validation.Schema{Type: "number", AllowAdditional: true}, "evReactiveSusceptance": &validation.Schema{Type: "number", AllowAdditional: true}, "evSessionTotalDischargeEnergyAvailable": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "evEnergyOffer": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"evAbsolutePriceSchedule": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"timeAnchor": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "currency": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 3, HasMaxLength: true}, "evAbsolutePriceScheduleEntries": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"duration": &validation.Schema{Type: "integer", AllowAdditional: true}, "evPriceRule": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"energyFee": &validation.Schema{Type: "number", AllowAdditional: true}, "powerRangeStart": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"energyFee", "powerRangeStart"}}, MinItems: 1, HasMinItems: true, MaxItems: 8, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"duration", "evPriceRule"}}, MinItems: 1, HasMinItems: true, MaxItems: 1024, HasMaxItems: true}, "priceAlgorithm": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 2000, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"timeAnchor", "currency", "priceAlgorithm", "evAbsolutePriceScheduleEntries"}}, "evPowerSchedule": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"evPowerScheduleEntries": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"duration": &validation.Schema{Type: "integer", AllowAdditional: true}, "power": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"duration", "power"}}, MinItems: 1, HasMinItems: true, MaxItems: 1024, HasMaxItems: true}, "timeAnchor": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"timeAnchor", "evPowerScheduleEntries"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"evPowerSchedule"}}, "requestedEnergyTransfer": &validation.Schema{Type: "string", Enum: []string{"AC_single_phase", "AC_two_phase", "AC_three_phase", "DC", "AC_BPT", "AC_BPT_DER", "AC_DER", "DC_BPT", "DC_ACDP", "DC_ACDP_BPT", "WPT"}}, "dcChargingParameters": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"evMaxCurrent": &validation.Schema{Type: "number", AllowAdditional: true}, "evMaxVoltage": &validation.Schema{Type: "number", AllowAdditional: true}, "evMaxPower": &validation.Schema{Type: "number", AllowAdditional: true}, "evEnergyCapacity": &validation.Schema{Type: "number", AllowAdditional: true}, "energyAmount": &validation.Schema{Type: "number", AllowAdditional: true}, "stateOfCharge": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true, Maximum: 100, HasMaximum: true}, "fullSoC": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true, Maximum: 100, HasMaximum: true}, "bulkSoC": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true, Maximum: 100, HasMaximum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"evMaxCurrent", "evMaxVoltage"}}, "v2xChargingParameters": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"minChargePower": &validation.Schema{Type: "number", AllowAdditional: true}, "minChargePower_L2": &validation.Schema{Type: "number", AllowAdditional: true}, "minChargePower_L3": &validation.Schema{Type: "number", AllowAdditional: true}, "maxChargePower": &validation.Schema{Type: "number", AllowAdditional: true}, "maxChargePower_L2": &validation.Schema{Type: "number", AllowAdditional: true}, "maxChargePower_L3": &validation.Schema{Type: "number", AllowAdditional: true}, "minDischargePower": &validation.Schema{Type: "number", AllowAdditional: true}, "minDischargePower_L2": &validation.Schema{Type: "number", AllowAdditional: true}, "minDischargePower_L3": &validation.Schema{Type: "number", AllowAdditional: true}, "maxDischargePower": &validation.Schema{Type: "number", AllowAdditional: true}, "maxDischargePower_L2": &validation.Schema{Type: "number", AllowAdditional: true}, "maxDischargePower_L3": &validation.Schema{Type: "number", AllowAdditional: true}, "minChargeCurrent": &validation.Schema{Type: "number", AllowAdditional: true}, "maxChargeCurrent": &validation.Schema{Type: "number", AllowAdditional: true}, "minDischargeCurrent": &validation.Schema{Type: "number", AllowAdditional: true}, "maxDischargeCurrent": &validation.Schema{Type: "number", AllowAdditional: true}, "minVoltage": &validation.Schema{Type: "number", AllowAdditional: true}, "maxVoltage": &validation.Schema{Type: "number", AllowAdditional: true}, "evTargetEnergyRequest": &validation.Schema{Type: "number", AllowAdditional: true}, "evMinEnergyRequest": &validation.Schema{Type: "number", AllowAdditional: true}, "evMaxEnergyRequest": &validation.Schema{Type: "number", AllowAdditional: true}, "evMinV2XEnergyRequest": &validation.Schema{Type: "number", AllowAdditional: true}, "evMaxV2XEnergyRequest": &validation.Schema{Type: "number", AllowAdditional: true}, "targetSoC": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true, Maximum: 100, HasMaximum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "availableEnergyTransfer": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", Enum: []string{"AC_single_phase", "AC_two_phase", "AC_three_phase", "DC", "AC_BPT", "AC_BPT_DER", "AC_DER", "DC_BPT", "DC_ACDP", "DC_ACDP_BPT", "WPT"}}, MinItems: 1, HasMinItems: true}, "controlMode": &validation.Schema{Type: "string", Enum: []string{"ScheduledControl", "DynamicControl"}}, "mobilityNeedsMode": &validation.Schema{Type: "string", Enum: []string{"EVCC", "EVCC_SECC"}}, "departureTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"requestedEnergyTransfer"}}, "timestamp": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"evseId", "chargingNeeds"}}

type NotifyEVChargingNeedsRequest struct {
	EVSEID            int                                       `json:"evseId"`
	MaxScheduleTuples *int                                      `json:"maxScheduleTuples,omitempty"`
	ChargingNeeds     NotifyEVChargingNeedsRequestChargingNeeds `json:"chargingNeeds"`
	Timestamp         *string                                   `json:"timestamp,omitempty"`
	CustomData        *NotifyEVChargingNeedsRequestCustomData   `json:"customData,omitempty"`
}

type NotifyEVChargingNeedsRequestChargingNeeds struct {
	AcChargingParameters    *NotifyEVChargingNeedsRequestACChargingParameters    `json:"acChargingParameters,omitempty"`
	DerChargingParameters   *NotifyEVChargingNeedsRequestDERChargingParameters   `json:"derChargingParameters,omitempty"`
	EVEnergyOffer           *NotifyEVChargingNeedsRequestEVEnergyOffer           `json:"evEnergyOffer,omitempty"`
	RequestedEnergyTransfer NotifyEVChargingNeedsRequestEnergyTransferModeEnum   `json:"requestedEnergyTransfer"`
	DcChargingParameters    *NotifyEVChargingNeedsRequestDCChargingParameters    `json:"dcChargingParameters,omitempty"`
	V2XChargingParameters   *NotifyEVChargingNeedsRequestV2XChargingParameters   `json:"v2xChargingParameters,omitempty"`
	AvailableEnergyTransfer []NotifyEVChargingNeedsRequestEnergyTransferModeEnum `json:"availableEnergyTransfer,omitempty"`
	ControlMode             *NotifyEVChargingNeedsRequestControlModeEnum         `json:"controlMode,omitempty"`
	MobilityNeedsMode       *NotifyEVChargingNeedsRequestMobilityNeedsModeEnum   `json:"mobilityNeedsMode,omitempty"`
	DepartureTime           *string                                              `json:"departureTime,omitempty"`
	CustomData              *NotifyEVChargingNeedsRequestCustomData              `json:"customData,omitempty"`
}

type NotifyEVChargingNeedsRequestMobilityNeedsModeEnum string

const (
	NotifyEVChargingNeedsRequestMobilityNeedsModeEnumEVCC     NotifyEVChargingNeedsRequestMobilityNeedsModeEnum = "EVCC"
	NotifyEVChargingNeedsRequestMobilityNeedsModeEnumEVCCSECC NotifyEVChargingNeedsRequestMobilityNeedsModeEnum = "EVCC_SECC"
)

type NotifyEVChargingNeedsRequestControlModeEnum string

const (
	NotifyEVChargingNeedsRequestControlModeEnumScheduledControl NotifyEVChargingNeedsRequestControlModeEnum = "ScheduledControl"
	NotifyEVChargingNeedsRequestControlModeEnumDynamicControl   NotifyEVChargingNeedsRequestControlModeEnum = "DynamicControl"
)

type NotifyEVChargingNeedsRequestV2XChargingParameters struct {
	MinChargePower        *float64                                `json:"minChargePower,omitempty"`
	MinChargePowerL2      *float64                                `json:"minChargePower_L2,omitempty"`
	MinChargePowerL3      *float64                                `json:"minChargePower_L3,omitempty"`
	MaxChargePower        *float64                                `json:"maxChargePower,omitempty"`
	MaxChargePowerL2      *float64                                `json:"maxChargePower_L2,omitempty"`
	MaxChargePowerL3      *float64                                `json:"maxChargePower_L3,omitempty"`
	MinDischargePower     *float64                                `json:"minDischargePower,omitempty"`
	MinDischargePowerL2   *float64                                `json:"minDischargePower_L2,omitempty"`
	MinDischargePowerL3   *float64                                `json:"minDischargePower_L3,omitempty"`
	MaxDischargePower     *float64                                `json:"maxDischargePower,omitempty"`
	MaxDischargePowerL2   *float64                                `json:"maxDischargePower_L2,omitempty"`
	MaxDischargePowerL3   *float64                                `json:"maxDischargePower_L3,omitempty"`
	MinChargeCurrent      *float64                                `json:"minChargeCurrent,omitempty"`
	MaxChargeCurrent      *float64                                `json:"maxChargeCurrent,omitempty"`
	MinDischargeCurrent   *float64                                `json:"minDischargeCurrent,omitempty"`
	MaxDischargeCurrent   *float64                                `json:"maxDischargeCurrent,omitempty"`
	MinVoltage            *float64                                `json:"minVoltage,omitempty"`
	MaxVoltage            *float64                                `json:"maxVoltage,omitempty"`
	EVTargetEnergyRequest *float64                                `json:"evTargetEnergyRequest,omitempty"`
	EVMinEnergyRequest    *float64                                `json:"evMinEnergyRequest,omitempty"`
	EVMaxEnergyRequest    *float64                                `json:"evMaxEnergyRequest,omitempty"`
	EVMinV2XEnergyRequest *float64                                `json:"evMinV2XEnergyRequest,omitempty"`
	EVMaxV2XEnergyRequest *float64                                `json:"evMaxV2XEnergyRequest,omitempty"`
	TargetSoC             *int                                    `json:"targetSoC,omitempty"`
	CustomData            *NotifyEVChargingNeedsRequestCustomData `json:"customData,omitempty"`
}

type NotifyEVChargingNeedsRequestDCChargingParameters struct {
	EVMaxCurrent     float64                                 `json:"evMaxCurrent"`
	EVMaxVoltage     float64                                 `json:"evMaxVoltage"`
	EVMaxPower       *float64                                `json:"evMaxPower,omitempty"`
	EVEnergyCapacity *float64                                `json:"evEnergyCapacity,omitempty"`
	EnergyAmount     *float64                                `json:"energyAmount,omitempty"`
	StateOfCharge    *int                                    `json:"stateOfCharge,omitempty"`
	FullSoC          *int                                    `json:"fullSoC,omitempty"`
	BulkSoC          *int                                    `json:"bulkSoC,omitempty"`
	CustomData       *NotifyEVChargingNeedsRequestCustomData `json:"customData,omitempty"`
}

type NotifyEVChargingNeedsRequestEnergyTransferModeEnum string

const (
	NotifyEVChargingNeedsRequestEnergyTransferModeEnumACSinglePhase NotifyEVChargingNeedsRequestEnergyTransferModeEnum = "AC_single_phase"
	NotifyEVChargingNeedsRequestEnergyTransferModeEnumACTwoPhase    NotifyEVChargingNeedsRequestEnergyTransferModeEnum = "AC_two_phase"
	NotifyEVChargingNeedsRequestEnergyTransferModeEnumACThreePhase  NotifyEVChargingNeedsRequestEnergyTransferModeEnum = "AC_three_phase"
	NotifyEVChargingNeedsRequestEnergyTransferModeEnumDC            NotifyEVChargingNeedsRequestEnergyTransferModeEnum = "DC"
	NotifyEVChargingNeedsRequestEnergyTransferModeEnumACBPT         NotifyEVChargingNeedsRequestEnergyTransferModeEnum = "AC_BPT"
	NotifyEVChargingNeedsRequestEnergyTransferModeEnumACBPTDER      NotifyEVChargingNeedsRequestEnergyTransferModeEnum = "AC_BPT_DER"
	NotifyEVChargingNeedsRequestEnergyTransferModeEnumACDER         NotifyEVChargingNeedsRequestEnergyTransferModeEnum = "AC_DER"
	NotifyEVChargingNeedsRequestEnergyTransferModeEnumDCBPT         NotifyEVChargingNeedsRequestEnergyTransferModeEnum = "DC_BPT"
	NotifyEVChargingNeedsRequestEnergyTransferModeEnumDCACDP        NotifyEVChargingNeedsRequestEnergyTransferModeEnum = "DC_ACDP"
	NotifyEVChargingNeedsRequestEnergyTransferModeEnumDCACDPBPT     NotifyEVChargingNeedsRequestEnergyTransferModeEnum = "DC_ACDP_BPT"
	NotifyEVChargingNeedsRequestEnergyTransferModeEnumWPT           NotifyEVChargingNeedsRequestEnergyTransferModeEnum = "WPT"
)

type NotifyEVChargingNeedsRequestEVEnergyOffer struct {
	EVAbsolutePriceSchedule *NotifyEVChargingNeedsRequestEVAbsolutePriceSchedule `json:"evAbsolutePriceSchedule,omitempty"`
	EVPowerSchedule         NotifyEVChargingNeedsRequestEVPowerSchedule          `json:"evPowerSchedule"`
	CustomData              *NotifyEVChargingNeedsRequestCustomData              `json:"customData,omitempty"`
}

type NotifyEVChargingNeedsRequestEVPowerSchedule struct {
	EVPowerScheduleEntries []NotifyEVChargingNeedsRequestEVPowerScheduleEntry `json:"evPowerScheduleEntries"`
	TimeAnchor             string                                             `json:"timeAnchor"`
	CustomData             *NotifyEVChargingNeedsRequestCustomData            `json:"customData,omitempty"`
}

type NotifyEVChargingNeedsRequestEVPowerScheduleEntry struct {
	Duration   int                                     `json:"duration"`
	Power      float64                                 `json:"power"`
	CustomData *NotifyEVChargingNeedsRequestCustomData `json:"customData,omitempty"`
}

type NotifyEVChargingNeedsRequestEVAbsolutePriceSchedule struct {
	TimeAnchor                     string                                                     `json:"timeAnchor"`
	Currency                       string                                                     `json:"currency"`
	EVAbsolutePriceScheduleEntries []NotifyEVChargingNeedsRequestEVAbsolutePriceScheduleEntry `json:"evAbsolutePriceScheduleEntries"`
	PriceAlgorithm                 string                                                     `json:"priceAlgorithm"`
	CustomData                     *NotifyEVChargingNeedsRequestCustomData                    `json:"customData,omitempty"`
}

type NotifyEVChargingNeedsRequestEVAbsolutePriceScheduleEntry struct {
	Duration    int                                       `json:"duration"`
	EVPriceRule []NotifyEVChargingNeedsRequestEVPriceRule `json:"evPriceRule"`
	CustomData  *NotifyEVChargingNeedsRequestCustomData   `json:"customData,omitempty"`
}

type NotifyEVChargingNeedsRequestEVPriceRule struct {
	EnergyFee       float64                                 `json:"energyFee"`
	PowerRangeStart float64                                 `json:"powerRangeStart"`
	CustomData      *NotifyEVChargingNeedsRequestCustomData `json:"customData,omitempty"`
}

type NotifyEVChargingNeedsRequestDERChargingParameters struct {
	EVSupportedDERControl                  []NotifyEVChargingNeedsRequestDERControlEnum         `json:"evSupportedDERControl,omitempty"`
	EVOverExcitedMaxDischargePower         *float64                                             `json:"evOverExcitedMaxDischargePower,omitempty"`
	EVOverExcitedPowerFactor               *float64                                             `json:"evOverExcitedPowerFactor,omitempty"`
	EVUnderExcitedMaxDischargePower        *float64                                             `json:"evUnderExcitedMaxDischargePower,omitempty"`
	EVUnderExcitedPowerFactor              *float64                                             `json:"evUnderExcitedPowerFactor,omitempty"`
	MaxApparentPower                       *float64                                             `json:"maxApparentPower,omitempty"`
	MaxChargeApparentPower                 *float64                                             `json:"maxChargeApparentPower,omitempty"`
	MaxChargeApparentPowerL2               *float64                                             `json:"maxChargeApparentPower_L2,omitempty"`
	MaxChargeApparentPowerL3               *float64                                             `json:"maxChargeApparentPower_L3,omitempty"`
	MaxDischargeApparentPower              *float64                                             `json:"maxDischargeApparentPower,omitempty"`
	MaxDischargeApparentPowerL2            *float64                                             `json:"maxDischargeApparentPower_L2,omitempty"`
	MaxDischargeApparentPowerL3            *float64                                             `json:"maxDischargeApparentPower_L3,omitempty"`
	MaxChargeReactivePower                 *float64                                             `json:"maxChargeReactivePower,omitempty"`
	MaxChargeReactivePowerL2               *float64                                             `json:"maxChargeReactivePower_L2,omitempty"`
	MaxChargeReactivePowerL3               *float64                                             `json:"maxChargeReactivePower_L3,omitempty"`
	MinChargeReactivePower                 *float64                                             `json:"minChargeReactivePower,omitempty"`
	MinChargeReactivePowerL2               *float64                                             `json:"minChargeReactivePower_L2,omitempty"`
	MinChargeReactivePowerL3               *float64                                             `json:"minChargeReactivePower_L3,omitempty"`
	MaxDischargeReactivePower              *float64                                             `json:"maxDischargeReactivePower,omitempty"`
	MaxDischargeReactivePowerL2            *float64                                             `json:"maxDischargeReactivePower_L2,omitempty"`
	MaxDischargeReactivePowerL3            *float64                                             `json:"maxDischargeReactivePower_L3,omitempty"`
	MinDischargeReactivePower              *float64                                             `json:"minDischargeReactivePower,omitempty"`
	MinDischargeReactivePowerL2            *float64                                             `json:"minDischargeReactivePower_L2,omitempty"`
	MinDischargeReactivePowerL3            *float64                                             `json:"minDischargeReactivePower_L3,omitempty"`
	NominalVoltage                         *float64                                             `json:"nominalVoltage,omitempty"`
	NominalVoltageOffset                   *float64                                             `json:"nominalVoltageOffset,omitempty"`
	MaxNominalVoltage                      *float64                                             `json:"maxNominalVoltage,omitempty"`
	MinNominalVoltage                      *float64                                             `json:"minNominalVoltage,omitempty"`
	EVInverterManufacturer                 *string                                              `json:"evInverterManufacturer,omitempty"`
	EVInverterModel                        *string                                              `json:"evInverterModel,omitempty"`
	EVInverterSerialNumber                 *string                                              `json:"evInverterSerialNumber,omitempty"`
	EVInverterSwVersion                    *string                                              `json:"evInverterSwVersion,omitempty"`
	EVInverterHwVersion                    *string                                              `json:"evInverterHwVersion,omitempty"`
	EVIslandingDetectionMethod             []NotifyEVChargingNeedsRequestIslandingDetectionEnum `json:"evIslandingDetectionMethod,omitempty"`
	EVIslandingTripTime                    *float64                                             `json:"evIslandingTripTime,omitempty"`
	EVMaximumLevel1DCInjection             *float64                                             `json:"evMaximumLevel1DCInjection,omitempty"`
	EVDurationLevel1DCInjection            *float64                                             `json:"evDurationLevel1DCInjection,omitempty"`
	EVMaximumLevel2DCInjection             *float64                                             `json:"evMaximumLevel2DCInjection,omitempty"`
	EVDurationLevel2DCInjection            *float64                                             `json:"evDurationLevel2DCInjection,omitempty"`
	EVReactiveSusceptance                  *float64                                             `json:"evReactiveSusceptance,omitempty"`
	EVSessionTotalDischargeEnergyAvailable *float64                                             `json:"evSessionTotalDischargeEnergyAvailable,omitempty"`
	CustomData                             *NotifyEVChargingNeedsRequestCustomData              `json:"customData,omitempty"`
}

type NotifyEVChargingNeedsRequestIslandingDetectionEnum string

const (
	NotifyEVChargingNeedsRequestIslandingDetectionEnumNoAntiIslandingSupport NotifyEVChargingNeedsRequestIslandingDetectionEnum = "NoAntiIslandingSupport"
	NotifyEVChargingNeedsRequestIslandingDetectionEnumRoCoF                  NotifyEVChargingNeedsRequestIslandingDetectionEnum = "RoCoF"
	NotifyEVChargingNeedsRequestIslandingDetectionEnumUVPOVP                 NotifyEVChargingNeedsRequestIslandingDetectionEnum = "UVP_OVP"
	NotifyEVChargingNeedsRequestIslandingDetectionEnumUFPOFP                 NotifyEVChargingNeedsRequestIslandingDetectionEnum = "UFP_OFP"
	NotifyEVChargingNeedsRequestIslandingDetectionEnumVoltageVectorShift     NotifyEVChargingNeedsRequestIslandingDetectionEnum = "VoltageVectorShift"
	NotifyEVChargingNeedsRequestIslandingDetectionEnumZeroCrossingDetection  NotifyEVChargingNeedsRequestIslandingDetectionEnum = "ZeroCrossingDetection"
	NotifyEVChargingNeedsRequestIslandingDetectionEnumOtherPassive           NotifyEVChargingNeedsRequestIslandingDetectionEnum = "OtherPassive"
	NotifyEVChargingNeedsRequestIslandingDetectionEnumImpedanceMeasurement   NotifyEVChargingNeedsRequestIslandingDetectionEnum = "ImpedanceMeasurement"
	NotifyEVChargingNeedsRequestIslandingDetectionEnumImpedanceAtFrequency   NotifyEVChargingNeedsRequestIslandingDetectionEnum = "ImpedanceAtFrequency"
	NotifyEVChargingNeedsRequestIslandingDetectionEnumSlipModeFrequencyShift NotifyEVChargingNeedsRequestIslandingDetectionEnum = "SlipModeFrequencyShift"
	NotifyEVChargingNeedsRequestIslandingDetectionEnumSandiaFrequencyShift   NotifyEVChargingNeedsRequestIslandingDetectionEnum = "SandiaFrequencyShift"
	NotifyEVChargingNeedsRequestIslandingDetectionEnumSandiaVoltageShift     NotifyEVChargingNeedsRequestIslandingDetectionEnum = "SandiaVoltageShift"
	NotifyEVChargingNeedsRequestIslandingDetectionEnumFrequencyJump          NotifyEVChargingNeedsRequestIslandingDetectionEnum = "FrequencyJump"
	NotifyEVChargingNeedsRequestIslandingDetectionEnumRCLQFactor             NotifyEVChargingNeedsRequestIslandingDetectionEnum = "RCLQFactor"
	NotifyEVChargingNeedsRequestIslandingDetectionEnumOtherActive            NotifyEVChargingNeedsRequestIslandingDetectionEnum = "OtherActive"
)

type NotifyEVChargingNeedsRequestDERControlEnum string

const (
	NotifyEVChargingNeedsRequestDERControlEnumEnterService            NotifyEVChargingNeedsRequestDERControlEnum = "EnterService"
	NotifyEVChargingNeedsRequestDERControlEnumFreqDroop               NotifyEVChargingNeedsRequestDERControlEnum = "FreqDroop"
	NotifyEVChargingNeedsRequestDERControlEnumFreqWatt                NotifyEVChargingNeedsRequestDERControlEnum = "FreqWatt"
	NotifyEVChargingNeedsRequestDERControlEnumFixedPFAbsorb           NotifyEVChargingNeedsRequestDERControlEnum = "FixedPFAbsorb"
	NotifyEVChargingNeedsRequestDERControlEnumFixedPFInject           NotifyEVChargingNeedsRequestDERControlEnum = "FixedPFInject"
	NotifyEVChargingNeedsRequestDERControlEnumFixedVar                NotifyEVChargingNeedsRequestDERControlEnum = "FixedVar"
	NotifyEVChargingNeedsRequestDERControlEnumGradients               NotifyEVChargingNeedsRequestDERControlEnum = "Gradients"
	NotifyEVChargingNeedsRequestDERControlEnumHFMustTrip              NotifyEVChargingNeedsRequestDERControlEnum = "HFMustTrip"
	NotifyEVChargingNeedsRequestDERControlEnumHFMayTrip               NotifyEVChargingNeedsRequestDERControlEnum = "HFMayTrip"
	NotifyEVChargingNeedsRequestDERControlEnumHVMustTrip              NotifyEVChargingNeedsRequestDERControlEnum = "HVMustTrip"
	NotifyEVChargingNeedsRequestDERControlEnumHVMomCess               NotifyEVChargingNeedsRequestDERControlEnum = "HVMomCess"
	NotifyEVChargingNeedsRequestDERControlEnumHVMayTrip               NotifyEVChargingNeedsRequestDERControlEnum = "HVMayTrip"
	NotifyEVChargingNeedsRequestDERControlEnumLimitMaxDischarge       NotifyEVChargingNeedsRequestDERControlEnum = "LimitMaxDischarge"
	NotifyEVChargingNeedsRequestDERControlEnumLFMustTrip              NotifyEVChargingNeedsRequestDERControlEnum = "LFMustTrip"
	NotifyEVChargingNeedsRequestDERControlEnumLVMustTrip              NotifyEVChargingNeedsRequestDERControlEnum = "LVMustTrip"
	NotifyEVChargingNeedsRequestDERControlEnumLVMomCess               NotifyEVChargingNeedsRequestDERControlEnum = "LVMomCess"
	NotifyEVChargingNeedsRequestDERControlEnumLVMayTrip               NotifyEVChargingNeedsRequestDERControlEnum = "LVMayTrip"
	NotifyEVChargingNeedsRequestDERControlEnumPowerMonitoringMustTrip NotifyEVChargingNeedsRequestDERControlEnum = "PowerMonitoringMustTrip"
	NotifyEVChargingNeedsRequestDERControlEnumVoltVar                 NotifyEVChargingNeedsRequestDERControlEnum = "VoltVar"
	NotifyEVChargingNeedsRequestDERControlEnumVoltWatt                NotifyEVChargingNeedsRequestDERControlEnum = "VoltWatt"
	NotifyEVChargingNeedsRequestDERControlEnumWattPF                  NotifyEVChargingNeedsRequestDERControlEnum = "WattPF"
	NotifyEVChargingNeedsRequestDERControlEnumWattVar                 NotifyEVChargingNeedsRequestDERControlEnum = "WattVar"
)

type NotifyEVChargingNeedsRequestACChargingParameters struct {
	EnergyAmount float64                                 `json:"energyAmount"`
	EVMinCurrent float64                                 `json:"evMinCurrent"`
	EVMaxCurrent float64                                 `json:"evMaxCurrent"`
	EVMaxVoltage float64                                 `json:"evMaxVoltage"`
	CustomData   *NotifyEVChargingNeedsRequestCustomData `json:"customData,omitempty"`
}

type NotifyEVChargingNeedsRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (NotifyEVChargingNeedsRequest) ActionName() string { return "NotifyEVChargingNeeds" }

func (NotifyEVChargingNeedsRequest) Version() protocol.Version { return protocol.OCPP21 }

func (NotifyEVChargingNeedsRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (NotifyEVChargingNeedsRequest) SchemaName() string { return "NotifyEVChargingNeedsRequest.json" }

func (message NotifyEVChargingNeedsRequest) Validate() error {
	return validation.Validate("NotifyEVChargingNeedsRequest.json", schemaNotifyEVChargingNeedsRequest, message)
}

func (NotifyEVChargingNeedsRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("NotifyEVChargingNeedsRequest.json", schemaNotifyEVChargingNeedsRequest, data)
}
