// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = StopTransactionRequest{}

var schemaStopTransactionRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"idTag": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "meterStop": &validation.Schema{Type: "integer", AllowAdditional: true}, "timestamp": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "transactionId": &validation.Schema{Type: "integer", AllowAdditional: true}, "reason": &validation.Schema{Type: "string", Enum: []string{"EmergencyStop", "EVDisconnected", "HardReset", "Local", "Other", "PowerLoss", "Reboot", "Remote", "SoftReset", "UnlockCommand", "DeAuthorized"}}, "transactionData": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"timestamp": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "sampledValue": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"value": &validation.Schema{Type: "string", AllowAdditional: true}, "context": &validation.Schema{Type: "string", Enum: []string{"Interruption.Begin", "Interruption.End", "Sample.Clock", "Sample.Periodic", "Transaction.Begin", "Transaction.End", "Trigger", "Other"}}, "format": &validation.Schema{Type: "string", Enum: []string{"Raw", "SignedData"}}, "measurand": &validation.Schema{Type: "string", Enum: []string{"Energy.Active.Export.Register", "Energy.Active.Import.Register", "Energy.Reactive.Export.Register", "Energy.Reactive.Import.Register", "Energy.Active.Export.Interval", "Energy.Active.Import.Interval", "Energy.Reactive.Export.Interval", "Energy.Reactive.Import.Interval", "Power.Active.Export", "Power.Active.Import", "Power.Offered", "Power.Reactive.Export", "Power.Reactive.Import", "Power.Factor", "Current.Import", "Current.Export", "Current.Offered", "Voltage", "Frequency", "Temperature", "SoC", "RPM"}}, "phase": &validation.Schema{Type: "string", Enum: []string{"L1", "L2", "L3", "N", "L1-N", "L2-N", "L3-N", "L1-L2", "L2-L3", "L3-L1"}}, "location": &validation.Schema{Type: "string", Enum: []string{"Cable", "EV", "Inlet", "Outlet", "Body"}}, "unit": &validation.Schema{Type: "string", Enum: []string{"Wh", "kWh", "varh", "kvarh", "W", "kW", "VA", "kVA", "var", "kvar", "A", "V", "K", "Celcius", "Celsius", "Fahrenheit", "Percent"}}}, Required: []string{"value"}}}}, Required: []string{"timestamp", "sampledValue"}}}}, Required: []string{"transactionId", "timestamp", "meterStop"}}

type StopTransactionRequest struct {
	IDTag           *string                                     `json:"idTag,omitempty"`
	MeterStop       int                                         `json:"meterStop"`
	Timestamp       string                                      `json:"timestamp"`
	TransactionID   int                                         `json:"transactionId"`
	Reason          *StopTransactionRequestReason               `json:"reason,omitempty"`
	TransactionData []StopTransactionRequestTransactionDataItem `json:"transactionData,omitempty"`
}

type StopTransactionRequestTransactionDataItem struct {
	Timestamp    string                                                      `json:"timestamp"`
	SampledValue []StopTransactionRequestTransactionDataItemSampledValueItem `json:"sampledValue"`
}

type StopTransactionRequestTransactionDataItemSampledValueItem struct {
	Value     string                                                              `json:"value"`
	Context   *StopTransactionRequestTransactionDataItemSampledValueItemContext   `json:"context,omitempty"`
	Format    *StopTransactionRequestTransactionDataItemSampledValueItemFormat    `json:"format,omitempty"`
	Measurand *StopTransactionRequestTransactionDataItemSampledValueItemMeasurand `json:"measurand,omitempty"`
	Phase     *StopTransactionRequestTransactionDataItemSampledValueItemPhase     `json:"phase,omitempty"`
	Location  *StopTransactionRequestTransactionDataItemSampledValueItemLocation  `json:"location,omitempty"`
	Unit      *StopTransactionRequestTransactionDataItemSampledValueItemUnit      `json:"unit,omitempty"`
}

type StopTransactionRequestTransactionDataItemSampledValueItemUnit string

const (
	StopTransactionRequestTransactionDataItemSampledValueItemUnitWh         StopTransactionRequestTransactionDataItemSampledValueItemUnit = "Wh"
	StopTransactionRequestTransactionDataItemSampledValueItemUnitKWh        StopTransactionRequestTransactionDataItemSampledValueItemUnit = "kWh"
	StopTransactionRequestTransactionDataItemSampledValueItemUnitVarh       StopTransactionRequestTransactionDataItemSampledValueItemUnit = "varh"
	StopTransactionRequestTransactionDataItemSampledValueItemUnitKvarh      StopTransactionRequestTransactionDataItemSampledValueItemUnit = "kvarh"
	StopTransactionRequestTransactionDataItemSampledValueItemUnitW          StopTransactionRequestTransactionDataItemSampledValueItemUnit = "W"
	StopTransactionRequestTransactionDataItemSampledValueItemUnitKW         StopTransactionRequestTransactionDataItemSampledValueItemUnit = "kW"
	StopTransactionRequestTransactionDataItemSampledValueItemUnitVA         StopTransactionRequestTransactionDataItemSampledValueItemUnit = "VA"
	StopTransactionRequestTransactionDataItemSampledValueItemUnitKVA        StopTransactionRequestTransactionDataItemSampledValueItemUnit = "kVA"
	StopTransactionRequestTransactionDataItemSampledValueItemUnitVar        StopTransactionRequestTransactionDataItemSampledValueItemUnit = "var"
	StopTransactionRequestTransactionDataItemSampledValueItemUnitKvar       StopTransactionRequestTransactionDataItemSampledValueItemUnit = "kvar"
	StopTransactionRequestTransactionDataItemSampledValueItemUnitA          StopTransactionRequestTransactionDataItemSampledValueItemUnit = "A"
	StopTransactionRequestTransactionDataItemSampledValueItemUnitV          StopTransactionRequestTransactionDataItemSampledValueItemUnit = "V"
	StopTransactionRequestTransactionDataItemSampledValueItemUnitK          StopTransactionRequestTransactionDataItemSampledValueItemUnit = "K"
	StopTransactionRequestTransactionDataItemSampledValueItemUnitCelcius    StopTransactionRequestTransactionDataItemSampledValueItemUnit = "Celcius"
	StopTransactionRequestTransactionDataItemSampledValueItemUnitCelsius    StopTransactionRequestTransactionDataItemSampledValueItemUnit = "Celsius"
	StopTransactionRequestTransactionDataItemSampledValueItemUnitFahrenheit StopTransactionRequestTransactionDataItemSampledValueItemUnit = "Fahrenheit"
	StopTransactionRequestTransactionDataItemSampledValueItemUnitPercent    StopTransactionRequestTransactionDataItemSampledValueItemUnit = "Percent"
)

type StopTransactionRequestTransactionDataItemSampledValueItemLocation string

const (
	StopTransactionRequestTransactionDataItemSampledValueItemLocationCable  StopTransactionRequestTransactionDataItemSampledValueItemLocation = "Cable"
	StopTransactionRequestTransactionDataItemSampledValueItemLocationEV     StopTransactionRequestTransactionDataItemSampledValueItemLocation = "EV"
	StopTransactionRequestTransactionDataItemSampledValueItemLocationInlet  StopTransactionRequestTransactionDataItemSampledValueItemLocation = "Inlet"
	StopTransactionRequestTransactionDataItemSampledValueItemLocationOutlet StopTransactionRequestTransactionDataItemSampledValueItemLocation = "Outlet"
	StopTransactionRequestTransactionDataItemSampledValueItemLocationBody   StopTransactionRequestTransactionDataItemSampledValueItemLocation = "Body"
)

type StopTransactionRequestTransactionDataItemSampledValueItemPhase string

const (
	StopTransactionRequestTransactionDataItemSampledValueItemPhaseL1   StopTransactionRequestTransactionDataItemSampledValueItemPhase = "L1"
	StopTransactionRequestTransactionDataItemSampledValueItemPhaseL2   StopTransactionRequestTransactionDataItemSampledValueItemPhase = "L2"
	StopTransactionRequestTransactionDataItemSampledValueItemPhaseL3   StopTransactionRequestTransactionDataItemSampledValueItemPhase = "L3"
	StopTransactionRequestTransactionDataItemSampledValueItemPhaseN    StopTransactionRequestTransactionDataItemSampledValueItemPhase = "N"
	StopTransactionRequestTransactionDataItemSampledValueItemPhaseL1N  StopTransactionRequestTransactionDataItemSampledValueItemPhase = "L1-N"
	StopTransactionRequestTransactionDataItemSampledValueItemPhaseL2N  StopTransactionRequestTransactionDataItemSampledValueItemPhase = "L2-N"
	StopTransactionRequestTransactionDataItemSampledValueItemPhaseL3N  StopTransactionRequestTransactionDataItemSampledValueItemPhase = "L3-N"
	StopTransactionRequestTransactionDataItemSampledValueItemPhaseL1L2 StopTransactionRequestTransactionDataItemSampledValueItemPhase = "L1-L2"
	StopTransactionRequestTransactionDataItemSampledValueItemPhaseL2L3 StopTransactionRequestTransactionDataItemSampledValueItemPhase = "L2-L3"
	StopTransactionRequestTransactionDataItemSampledValueItemPhaseL3L1 StopTransactionRequestTransactionDataItemSampledValueItemPhase = "L3-L1"
)

type StopTransactionRequestTransactionDataItemSampledValueItemMeasurand string

const (
	StopTransactionRequestTransactionDataItemSampledValueItemMeasurandEnergyActiveExportRegister   StopTransactionRequestTransactionDataItemSampledValueItemMeasurand = "Energy.Active.Export.Register"
	StopTransactionRequestTransactionDataItemSampledValueItemMeasurandEnergyActiveImportRegister   StopTransactionRequestTransactionDataItemSampledValueItemMeasurand = "Energy.Active.Import.Register"
	StopTransactionRequestTransactionDataItemSampledValueItemMeasurandEnergyReactiveExportRegister StopTransactionRequestTransactionDataItemSampledValueItemMeasurand = "Energy.Reactive.Export.Register"
	StopTransactionRequestTransactionDataItemSampledValueItemMeasurandEnergyReactiveImportRegister StopTransactionRequestTransactionDataItemSampledValueItemMeasurand = "Energy.Reactive.Import.Register"
	StopTransactionRequestTransactionDataItemSampledValueItemMeasurandEnergyActiveExportInterval   StopTransactionRequestTransactionDataItemSampledValueItemMeasurand = "Energy.Active.Export.Interval"
	StopTransactionRequestTransactionDataItemSampledValueItemMeasurandEnergyActiveImportInterval   StopTransactionRequestTransactionDataItemSampledValueItemMeasurand = "Energy.Active.Import.Interval"
	StopTransactionRequestTransactionDataItemSampledValueItemMeasurandEnergyReactiveExportInterval StopTransactionRequestTransactionDataItemSampledValueItemMeasurand = "Energy.Reactive.Export.Interval"
	StopTransactionRequestTransactionDataItemSampledValueItemMeasurandEnergyReactiveImportInterval StopTransactionRequestTransactionDataItemSampledValueItemMeasurand = "Energy.Reactive.Import.Interval"
	StopTransactionRequestTransactionDataItemSampledValueItemMeasurandPowerActiveExport            StopTransactionRequestTransactionDataItemSampledValueItemMeasurand = "Power.Active.Export"
	StopTransactionRequestTransactionDataItemSampledValueItemMeasurandPowerActiveImport            StopTransactionRequestTransactionDataItemSampledValueItemMeasurand = "Power.Active.Import"
	StopTransactionRequestTransactionDataItemSampledValueItemMeasurandPowerOffered                 StopTransactionRequestTransactionDataItemSampledValueItemMeasurand = "Power.Offered"
	StopTransactionRequestTransactionDataItemSampledValueItemMeasurandPowerReactiveExport          StopTransactionRequestTransactionDataItemSampledValueItemMeasurand = "Power.Reactive.Export"
	StopTransactionRequestTransactionDataItemSampledValueItemMeasurandPowerReactiveImport          StopTransactionRequestTransactionDataItemSampledValueItemMeasurand = "Power.Reactive.Import"
	StopTransactionRequestTransactionDataItemSampledValueItemMeasurandPowerFactor                  StopTransactionRequestTransactionDataItemSampledValueItemMeasurand = "Power.Factor"
	StopTransactionRequestTransactionDataItemSampledValueItemMeasurandCurrentImport                StopTransactionRequestTransactionDataItemSampledValueItemMeasurand = "Current.Import"
	StopTransactionRequestTransactionDataItemSampledValueItemMeasurandCurrentExport                StopTransactionRequestTransactionDataItemSampledValueItemMeasurand = "Current.Export"
	StopTransactionRequestTransactionDataItemSampledValueItemMeasurandCurrentOffered               StopTransactionRequestTransactionDataItemSampledValueItemMeasurand = "Current.Offered"
	StopTransactionRequestTransactionDataItemSampledValueItemMeasurandVoltage                      StopTransactionRequestTransactionDataItemSampledValueItemMeasurand = "Voltage"
	StopTransactionRequestTransactionDataItemSampledValueItemMeasurandFrequency                    StopTransactionRequestTransactionDataItemSampledValueItemMeasurand = "Frequency"
	StopTransactionRequestTransactionDataItemSampledValueItemMeasurandTemperature                  StopTransactionRequestTransactionDataItemSampledValueItemMeasurand = "Temperature"
	StopTransactionRequestTransactionDataItemSampledValueItemMeasurandSoC                          StopTransactionRequestTransactionDataItemSampledValueItemMeasurand = "SoC"
	StopTransactionRequestTransactionDataItemSampledValueItemMeasurandRPM                          StopTransactionRequestTransactionDataItemSampledValueItemMeasurand = "RPM"
)

type StopTransactionRequestTransactionDataItemSampledValueItemFormat string

const (
	StopTransactionRequestTransactionDataItemSampledValueItemFormatRaw        StopTransactionRequestTransactionDataItemSampledValueItemFormat = "Raw"
	StopTransactionRequestTransactionDataItemSampledValueItemFormatSignedData StopTransactionRequestTransactionDataItemSampledValueItemFormat = "SignedData"
)

type StopTransactionRequestTransactionDataItemSampledValueItemContext string

const (
	StopTransactionRequestTransactionDataItemSampledValueItemContextInterruptionBegin StopTransactionRequestTransactionDataItemSampledValueItemContext = "Interruption.Begin"
	StopTransactionRequestTransactionDataItemSampledValueItemContextInterruptionEnd   StopTransactionRequestTransactionDataItemSampledValueItemContext = "Interruption.End"
	StopTransactionRequestTransactionDataItemSampledValueItemContextSampleClock       StopTransactionRequestTransactionDataItemSampledValueItemContext = "Sample.Clock"
	StopTransactionRequestTransactionDataItemSampledValueItemContextSamplePeriodic    StopTransactionRequestTransactionDataItemSampledValueItemContext = "Sample.Periodic"
	StopTransactionRequestTransactionDataItemSampledValueItemContextTransactionBegin  StopTransactionRequestTransactionDataItemSampledValueItemContext = "Transaction.Begin"
	StopTransactionRequestTransactionDataItemSampledValueItemContextTransactionEnd    StopTransactionRequestTransactionDataItemSampledValueItemContext = "Transaction.End"
	StopTransactionRequestTransactionDataItemSampledValueItemContextTrigger           StopTransactionRequestTransactionDataItemSampledValueItemContext = "Trigger"
	StopTransactionRequestTransactionDataItemSampledValueItemContextOther             StopTransactionRequestTransactionDataItemSampledValueItemContext = "Other"
)

type StopTransactionRequestReason string

const (
	StopTransactionRequestReasonEmergencyStop  StopTransactionRequestReason = "EmergencyStop"
	StopTransactionRequestReasonEVDisconnected StopTransactionRequestReason = "EVDisconnected"
	StopTransactionRequestReasonHardReset      StopTransactionRequestReason = "HardReset"
	StopTransactionRequestReasonLocal          StopTransactionRequestReason = "Local"
	StopTransactionRequestReasonOther          StopTransactionRequestReason = "Other"
	StopTransactionRequestReasonPowerLoss      StopTransactionRequestReason = "PowerLoss"
	StopTransactionRequestReasonReboot         StopTransactionRequestReason = "Reboot"
	StopTransactionRequestReasonRemote         StopTransactionRequestReason = "Remote"
	StopTransactionRequestReasonSoftReset      StopTransactionRequestReason = "SoftReset"
	StopTransactionRequestReasonUnlockCommand  StopTransactionRequestReason = "UnlockCommand"
	StopTransactionRequestReasonDeAuthorized   StopTransactionRequestReason = "DeAuthorized"
)

func (StopTransactionRequest) ActionName() string { return "StopTransaction" }

func (StopTransactionRequest) Version() protocol.Version { return protocol.OCPP16 }

func (StopTransactionRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (StopTransactionRequest) SchemaName() string { return "StopTransaction.json" }

func (message StopTransactionRequest) Validate() error {
	return validation.Validate("StopTransaction.json", schemaStopTransactionRequest, message)
}

func (StopTransactionRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("StopTransaction.json", schemaStopTransactionRequest, data)
}
