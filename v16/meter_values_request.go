// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = MeterValuesRequest{}

var schemaMeterValuesRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"connectorId": &validation.Schema{Type: "integer", AllowAdditional: true}, "transactionId": &validation.Schema{Type: "integer", AllowAdditional: true}, "meterValue": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"timestamp": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "sampledValue": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"value": &validation.Schema{Type: "string", AllowAdditional: true}, "context": &validation.Schema{Type: "string", Enum: []string{"Interruption.Begin", "Interruption.End", "Sample.Clock", "Sample.Periodic", "Transaction.Begin", "Transaction.End", "Trigger", "Other"}}, "format": &validation.Schema{Type: "string", Enum: []string{"Raw", "SignedData"}}, "measurand": &validation.Schema{Type: "string", Enum: []string{"Energy.Active.Export.Register", "Energy.Active.Import.Register", "Energy.Reactive.Export.Register", "Energy.Reactive.Import.Register", "Energy.Active.Export.Interval", "Energy.Active.Import.Interval", "Energy.Reactive.Export.Interval", "Energy.Reactive.Import.Interval", "Power.Active.Export", "Power.Active.Import", "Power.Offered", "Power.Reactive.Export", "Power.Reactive.Import", "Power.Factor", "Current.Import", "Current.Export", "Current.Offered", "Voltage", "Frequency", "Temperature", "SoC", "RPM"}}, "phase": &validation.Schema{Type: "string", Enum: []string{"L1", "L2", "L3", "N", "L1-N", "L2-N", "L3-N", "L1-L2", "L2-L3", "L3-L1"}}, "location": &validation.Schema{Type: "string", Enum: []string{"Cable", "EV", "Inlet", "Outlet", "Body"}}, "unit": &validation.Schema{Type: "string", Enum: []string{"Wh", "kWh", "varh", "kvarh", "W", "kW", "VA", "kVA", "var", "kvar", "A", "V", "K", "Celcius", "Celsius", "Fahrenheit", "Percent"}}}, Required: []string{"value"}}}}, Required: []string{"timestamp", "sampledValue"}}}}, Required: []string{"connectorId", "meterValue"}}

type MeterValuesRequest struct {
	ConnectorID   int                                `json:"connectorId"`
	TransactionID *int                               `json:"transactionId,omitempty"`
	MeterValue    []MeterValuesRequestMeterValueItem `json:"meterValue"`
}

type MeterValuesRequestMeterValueItem struct {
	Timestamp    string                                             `json:"timestamp"`
	SampledValue []MeterValuesRequestMeterValueItemSampledValueItem `json:"sampledValue"`
}

type MeterValuesRequestMeterValueItemSampledValueItem struct {
	Value     string                                                     `json:"value"`
	Context   *MeterValuesRequestMeterValueItemSampledValueItemContext   `json:"context,omitempty"`
	Format    *MeterValuesRequestMeterValueItemSampledValueItemFormat    `json:"format,omitempty"`
	Measurand *MeterValuesRequestMeterValueItemSampledValueItemMeasurand `json:"measurand,omitempty"`
	Phase     *MeterValuesRequestMeterValueItemSampledValueItemPhase     `json:"phase,omitempty"`
	Location  *MeterValuesRequestMeterValueItemSampledValueItemLocation  `json:"location,omitempty"`
	Unit      *MeterValuesRequestMeterValueItemSampledValueItemUnit      `json:"unit,omitempty"`
}

type MeterValuesRequestMeterValueItemSampledValueItemUnit string

const (
	MeterValuesRequestMeterValueItemSampledValueItemUnitWh         MeterValuesRequestMeterValueItemSampledValueItemUnit = "Wh"
	MeterValuesRequestMeterValueItemSampledValueItemUnitKWh        MeterValuesRequestMeterValueItemSampledValueItemUnit = "kWh"
	MeterValuesRequestMeterValueItemSampledValueItemUnitVarh       MeterValuesRequestMeterValueItemSampledValueItemUnit = "varh"
	MeterValuesRequestMeterValueItemSampledValueItemUnitKvarh      MeterValuesRequestMeterValueItemSampledValueItemUnit = "kvarh"
	MeterValuesRequestMeterValueItemSampledValueItemUnitW          MeterValuesRequestMeterValueItemSampledValueItemUnit = "W"
	MeterValuesRequestMeterValueItemSampledValueItemUnitKW         MeterValuesRequestMeterValueItemSampledValueItemUnit = "kW"
	MeterValuesRequestMeterValueItemSampledValueItemUnitVA         MeterValuesRequestMeterValueItemSampledValueItemUnit = "VA"
	MeterValuesRequestMeterValueItemSampledValueItemUnitKVA        MeterValuesRequestMeterValueItemSampledValueItemUnit = "kVA"
	MeterValuesRequestMeterValueItemSampledValueItemUnitVar        MeterValuesRequestMeterValueItemSampledValueItemUnit = "var"
	MeterValuesRequestMeterValueItemSampledValueItemUnitKvar       MeterValuesRequestMeterValueItemSampledValueItemUnit = "kvar"
	MeterValuesRequestMeterValueItemSampledValueItemUnitA          MeterValuesRequestMeterValueItemSampledValueItemUnit = "A"
	MeterValuesRequestMeterValueItemSampledValueItemUnitV          MeterValuesRequestMeterValueItemSampledValueItemUnit = "V"
	MeterValuesRequestMeterValueItemSampledValueItemUnitK          MeterValuesRequestMeterValueItemSampledValueItemUnit = "K"
	MeterValuesRequestMeterValueItemSampledValueItemUnitCelcius    MeterValuesRequestMeterValueItemSampledValueItemUnit = "Celcius"
	MeterValuesRequestMeterValueItemSampledValueItemUnitCelsius    MeterValuesRequestMeterValueItemSampledValueItemUnit = "Celsius"
	MeterValuesRequestMeterValueItemSampledValueItemUnitFahrenheit MeterValuesRequestMeterValueItemSampledValueItemUnit = "Fahrenheit"
	MeterValuesRequestMeterValueItemSampledValueItemUnitPercent    MeterValuesRequestMeterValueItemSampledValueItemUnit = "Percent"
)

type MeterValuesRequestMeterValueItemSampledValueItemLocation string

const (
	MeterValuesRequestMeterValueItemSampledValueItemLocationCable  MeterValuesRequestMeterValueItemSampledValueItemLocation = "Cable"
	MeterValuesRequestMeterValueItemSampledValueItemLocationEV     MeterValuesRequestMeterValueItemSampledValueItemLocation = "EV"
	MeterValuesRequestMeterValueItemSampledValueItemLocationInlet  MeterValuesRequestMeterValueItemSampledValueItemLocation = "Inlet"
	MeterValuesRequestMeterValueItemSampledValueItemLocationOutlet MeterValuesRequestMeterValueItemSampledValueItemLocation = "Outlet"
	MeterValuesRequestMeterValueItemSampledValueItemLocationBody   MeterValuesRequestMeterValueItemSampledValueItemLocation = "Body"
)

type MeterValuesRequestMeterValueItemSampledValueItemPhase string

const (
	MeterValuesRequestMeterValueItemSampledValueItemPhaseL1   MeterValuesRequestMeterValueItemSampledValueItemPhase = "L1"
	MeterValuesRequestMeterValueItemSampledValueItemPhaseL2   MeterValuesRequestMeterValueItemSampledValueItemPhase = "L2"
	MeterValuesRequestMeterValueItemSampledValueItemPhaseL3   MeterValuesRequestMeterValueItemSampledValueItemPhase = "L3"
	MeterValuesRequestMeterValueItemSampledValueItemPhaseN    MeterValuesRequestMeterValueItemSampledValueItemPhase = "N"
	MeterValuesRequestMeterValueItemSampledValueItemPhaseL1N  MeterValuesRequestMeterValueItemSampledValueItemPhase = "L1-N"
	MeterValuesRequestMeterValueItemSampledValueItemPhaseL2N  MeterValuesRequestMeterValueItemSampledValueItemPhase = "L2-N"
	MeterValuesRequestMeterValueItemSampledValueItemPhaseL3N  MeterValuesRequestMeterValueItemSampledValueItemPhase = "L3-N"
	MeterValuesRequestMeterValueItemSampledValueItemPhaseL1L2 MeterValuesRequestMeterValueItemSampledValueItemPhase = "L1-L2"
	MeterValuesRequestMeterValueItemSampledValueItemPhaseL2L3 MeterValuesRequestMeterValueItemSampledValueItemPhase = "L2-L3"
	MeterValuesRequestMeterValueItemSampledValueItemPhaseL3L1 MeterValuesRequestMeterValueItemSampledValueItemPhase = "L3-L1"
)

type MeterValuesRequestMeterValueItemSampledValueItemMeasurand string

const (
	MeterValuesRequestMeterValueItemSampledValueItemMeasurandEnergyActiveExportRegister   MeterValuesRequestMeterValueItemSampledValueItemMeasurand = "Energy.Active.Export.Register"
	MeterValuesRequestMeterValueItemSampledValueItemMeasurandEnergyActiveImportRegister   MeterValuesRequestMeterValueItemSampledValueItemMeasurand = "Energy.Active.Import.Register"
	MeterValuesRequestMeterValueItemSampledValueItemMeasurandEnergyReactiveExportRegister MeterValuesRequestMeterValueItemSampledValueItemMeasurand = "Energy.Reactive.Export.Register"
	MeterValuesRequestMeterValueItemSampledValueItemMeasurandEnergyReactiveImportRegister MeterValuesRequestMeterValueItemSampledValueItemMeasurand = "Energy.Reactive.Import.Register"
	MeterValuesRequestMeterValueItemSampledValueItemMeasurandEnergyActiveExportInterval   MeterValuesRequestMeterValueItemSampledValueItemMeasurand = "Energy.Active.Export.Interval"
	MeterValuesRequestMeterValueItemSampledValueItemMeasurandEnergyActiveImportInterval   MeterValuesRequestMeterValueItemSampledValueItemMeasurand = "Energy.Active.Import.Interval"
	MeterValuesRequestMeterValueItemSampledValueItemMeasurandEnergyReactiveExportInterval MeterValuesRequestMeterValueItemSampledValueItemMeasurand = "Energy.Reactive.Export.Interval"
	MeterValuesRequestMeterValueItemSampledValueItemMeasurandEnergyReactiveImportInterval MeterValuesRequestMeterValueItemSampledValueItemMeasurand = "Energy.Reactive.Import.Interval"
	MeterValuesRequestMeterValueItemSampledValueItemMeasurandPowerActiveExport            MeterValuesRequestMeterValueItemSampledValueItemMeasurand = "Power.Active.Export"
	MeterValuesRequestMeterValueItemSampledValueItemMeasurandPowerActiveImport            MeterValuesRequestMeterValueItemSampledValueItemMeasurand = "Power.Active.Import"
	MeterValuesRequestMeterValueItemSampledValueItemMeasurandPowerOffered                 MeterValuesRequestMeterValueItemSampledValueItemMeasurand = "Power.Offered"
	MeterValuesRequestMeterValueItemSampledValueItemMeasurandPowerReactiveExport          MeterValuesRequestMeterValueItemSampledValueItemMeasurand = "Power.Reactive.Export"
	MeterValuesRequestMeterValueItemSampledValueItemMeasurandPowerReactiveImport          MeterValuesRequestMeterValueItemSampledValueItemMeasurand = "Power.Reactive.Import"
	MeterValuesRequestMeterValueItemSampledValueItemMeasurandPowerFactor                  MeterValuesRequestMeterValueItemSampledValueItemMeasurand = "Power.Factor"
	MeterValuesRequestMeterValueItemSampledValueItemMeasurandCurrentImport                MeterValuesRequestMeterValueItemSampledValueItemMeasurand = "Current.Import"
	MeterValuesRequestMeterValueItemSampledValueItemMeasurandCurrentExport                MeterValuesRequestMeterValueItemSampledValueItemMeasurand = "Current.Export"
	MeterValuesRequestMeterValueItemSampledValueItemMeasurandCurrentOffered               MeterValuesRequestMeterValueItemSampledValueItemMeasurand = "Current.Offered"
	MeterValuesRequestMeterValueItemSampledValueItemMeasurandVoltage                      MeterValuesRequestMeterValueItemSampledValueItemMeasurand = "Voltage"
	MeterValuesRequestMeterValueItemSampledValueItemMeasurandFrequency                    MeterValuesRequestMeterValueItemSampledValueItemMeasurand = "Frequency"
	MeterValuesRequestMeterValueItemSampledValueItemMeasurandTemperature                  MeterValuesRequestMeterValueItemSampledValueItemMeasurand = "Temperature"
	MeterValuesRequestMeterValueItemSampledValueItemMeasurandSoC                          MeterValuesRequestMeterValueItemSampledValueItemMeasurand = "SoC"
	MeterValuesRequestMeterValueItemSampledValueItemMeasurandRPM                          MeterValuesRequestMeterValueItemSampledValueItemMeasurand = "RPM"
)

type MeterValuesRequestMeterValueItemSampledValueItemFormat string

const (
	MeterValuesRequestMeterValueItemSampledValueItemFormatRaw        MeterValuesRequestMeterValueItemSampledValueItemFormat = "Raw"
	MeterValuesRequestMeterValueItemSampledValueItemFormatSignedData MeterValuesRequestMeterValueItemSampledValueItemFormat = "SignedData"
)

type MeterValuesRequestMeterValueItemSampledValueItemContext string

const (
	MeterValuesRequestMeterValueItemSampledValueItemContextInterruptionBegin MeterValuesRequestMeterValueItemSampledValueItemContext = "Interruption.Begin"
	MeterValuesRequestMeterValueItemSampledValueItemContextInterruptionEnd   MeterValuesRequestMeterValueItemSampledValueItemContext = "Interruption.End"
	MeterValuesRequestMeterValueItemSampledValueItemContextSampleClock       MeterValuesRequestMeterValueItemSampledValueItemContext = "Sample.Clock"
	MeterValuesRequestMeterValueItemSampledValueItemContextSamplePeriodic    MeterValuesRequestMeterValueItemSampledValueItemContext = "Sample.Periodic"
	MeterValuesRequestMeterValueItemSampledValueItemContextTransactionBegin  MeterValuesRequestMeterValueItemSampledValueItemContext = "Transaction.Begin"
	MeterValuesRequestMeterValueItemSampledValueItemContextTransactionEnd    MeterValuesRequestMeterValueItemSampledValueItemContext = "Transaction.End"
	MeterValuesRequestMeterValueItemSampledValueItemContextTrigger           MeterValuesRequestMeterValueItemSampledValueItemContext = "Trigger"
	MeterValuesRequestMeterValueItemSampledValueItemContextOther             MeterValuesRequestMeterValueItemSampledValueItemContext = "Other"
)

func (MeterValuesRequest) ActionName() string { return "MeterValues" }

func (MeterValuesRequest) Version() protocol.Version { return protocol.OCPP16 }

func (MeterValuesRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (MeterValuesRequest) SchemaName() string { return "MeterValues.json" }

func (message MeterValuesRequest) Validate() error {
	return validation.Validate("MeterValues.json", schemaMeterValuesRequest, message)
}

func (MeterValuesRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("MeterValues.json", schemaMeterValuesRequest, data)
}
