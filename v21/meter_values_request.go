// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = MeterValuesRequest{}

var schemaMeterValuesRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"evseId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "meterValue": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"sampledValue": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"value": &validation.Schema{Type: "number", AllowAdditional: true}, "measurand": &validation.Schema{Type: "string", Enum: []string{"Current.Export", "Current.Export.Offered", "Current.Export.Minimum", "Current.Import", "Current.Import.Offered", "Current.Import.Minimum", "Current.Offered", "Display.PresentSOC", "Display.MinimumSOC", "Display.TargetSOC", "Display.MaximumSOC", "Display.RemainingTimeToMinimumSOC", "Display.RemainingTimeToTargetSOC", "Display.RemainingTimeToMaximumSOC", "Display.ChargingComplete", "Display.BatteryEnergyCapacity", "Display.InletHot", "Energy.Active.Export.Interval", "Energy.Active.Export.Register", "Energy.Active.Import.Interval", "Energy.Active.Import.Register", "Energy.Active.Import.CableLoss", "Energy.Active.Import.LocalGeneration.Register", "Energy.Active.Net", "Energy.Active.Setpoint.Interval", "Energy.Apparent.Export", "Energy.Apparent.Import", "Energy.Apparent.Net", "Energy.Reactive.Export.Interval", "Energy.Reactive.Export.Register", "Energy.Reactive.Import.Interval", "Energy.Reactive.Import.Register", "Energy.Reactive.Net", "EnergyRequest.Target", "EnergyRequest.Minimum", "EnergyRequest.Maximum", "EnergyRequest.Minimum.V2X", "EnergyRequest.Maximum.V2X", "EnergyRequest.Bulk", "Frequency", "Power.Active.Export", "Power.Active.Import", "Power.Active.Setpoint", "Power.Active.Residual", "Power.Export.Minimum", "Power.Export.Offered", "Power.Factor", "Power.Import.Offered", "Power.Import.Minimum", "Power.Offered", "Power.Reactive.Export", "Power.Reactive.Import", "SoC", "Voltage", "Voltage.Minimum", "Voltage.Maximum"}}, "context": &validation.Schema{Type: "string", Enum: []string{"Interruption.Begin", "Interruption.End", "Other", "Sample.Clock", "Sample.Periodic", "Transaction.Begin", "Transaction.End", "Trigger"}}, "phase": &validation.Schema{Type: "string", Enum: []string{"L1", "L2", "L3", "N", "L1-N", "L2-N", "L3-N", "L1-L2", "L2-L3", "L3-L1"}}, "location": &validation.Schema{Type: "string", Enum: []string{"Body", "Cable", "EV", "Inlet", "Outlet", "Upstream"}}, "signedMeterValue": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"signedMeterData": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 32768, HasMaxLength: true}, "signingMethod": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "encodingMethod": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "publicKey": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 2500, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"signedMeterData", "encodingMethod"}}, "unitOfMeasure": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"unit": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "multiplier": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"value"}}, MinItems: 1, HasMinItems: true}, "timestamp": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"timestamp", "sampledValue"}}, MinItems: 1, HasMinItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"evseId", "meterValue"}}

type MeterValuesRequest struct {
	EVSEID     int                            `json:"evseId"`
	MeterValue []MeterValuesRequestMeterValue `json:"meterValue"`
	CustomData *MeterValuesRequestCustomData  `json:"customData,omitempty"`
}

type MeterValuesRequestMeterValue struct {
	SampledValue []MeterValuesRequestSampledValue `json:"sampledValue"`
	Timestamp    string                           `json:"timestamp"`
	CustomData   *MeterValuesRequestCustomData    `json:"customData,omitempty"`
}

type MeterValuesRequestSampledValue struct {
	Value            float64                               `json:"value"`
	Measurand        *MeterValuesRequestMeasurandEnum      `json:"measurand,omitempty"`
	Context          *MeterValuesRequestReadingContextEnum `json:"context,omitempty"`
	Phase            *MeterValuesRequestPhaseEnum          `json:"phase,omitempty"`
	Location         *MeterValuesRequestLocationEnum       `json:"location,omitempty"`
	SignedMeterValue *MeterValuesRequestSignedMeterValue   `json:"signedMeterValue,omitempty"`
	UnitOfMeasure    *MeterValuesRequestUnitOfMeasure      `json:"unitOfMeasure,omitempty"`
	CustomData       *MeterValuesRequestCustomData         `json:"customData,omitempty"`
}

type MeterValuesRequestUnitOfMeasure struct {
	Unit       *string                       `json:"unit,omitempty"`
	Multiplier *int                          `json:"multiplier,omitempty"`
	CustomData *MeterValuesRequestCustomData `json:"customData,omitempty"`
}

type MeterValuesRequestSignedMeterValue struct {
	SignedMeterData string                        `json:"signedMeterData"`
	SigningMethod   *string                       `json:"signingMethod,omitempty"`
	EncodingMethod  string                        `json:"encodingMethod"`
	PublicKey       *string                       `json:"publicKey,omitempty"`
	CustomData      *MeterValuesRequestCustomData `json:"customData,omitempty"`
}

type MeterValuesRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

type MeterValuesRequestLocationEnum string

const (
	MeterValuesRequestLocationEnumBody     MeterValuesRequestLocationEnum = "Body"
	MeterValuesRequestLocationEnumCable    MeterValuesRequestLocationEnum = "Cable"
	MeterValuesRequestLocationEnumEV       MeterValuesRequestLocationEnum = "EV"
	MeterValuesRequestLocationEnumInlet    MeterValuesRequestLocationEnum = "Inlet"
	MeterValuesRequestLocationEnumOutlet   MeterValuesRequestLocationEnum = "Outlet"
	MeterValuesRequestLocationEnumUpstream MeterValuesRequestLocationEnum = "Upstream"
)

type MeterValuesRequestPhaseEnum string

const (
	MeterValuesRequestPhaseEnumL1   MeterValuesRequestPhaseEnum = "L1"
	MeterValuesRequestPhaseEnumL2   MeterValuesRequestPhaseEnum = "L2"
	MeterValuesRequestPhaseEnumL3   MeterValuesRequestPhaseEnum = "L3"
	MeterValuesRequestPhaseEnumN    MeterValuesRequestPhaseEnum = "N"
	MeterValuesRequestPhaseEnumL1N  MeterValuesRequestPhaseEnum = "L1-N"
	MeterValuesRequestPhaseEnumL2N  MeterValuesRequestPhaseEnum = "L2-N"
	MeterValuesRequestPhaseEnumL3N  MeterValuesRequestPhaseEnum = "L3-N"
	MeterValuesRequestPhaseEnumL1L2 MeterValuesRequestPhaseEnum = "L1-L2"
	MeterValuesRequestPhaseEnumL2L3 MeterValuesRequestPhaseEnum = "L2-L3"
	MeterValuesRequestPhaseEnumL3L1 MeterValuesRequestPhaseEnum = "L3-L1"
)

type MeterValuesRequestReadingContextEnum string

const (
	MeterValuesRequestReadingContextEnumInterruptionBegin MeterValuesRequestReadingContextEnum = "Interruption.Begin"
	MeterValuesRequestReadingContextEnumInterruptionEnd   MeterValuesRequestReadingContextEnum = "Interruption.End"
	MeterValuesRequestReadingContextEnumOther             MeterValuesRequestReadingContextEnum = "Other"
	MeterValuesRequestReadingContextEnumSampleClock       MeterValuesRequestReadingContextEnum = "Sample.Clock"
	MeterValuesRequestReadingContextEnumSamplePeriodic    MeterValuesRequestReadingContextEnum = "Sample.Periodic"
	MeterValuesRequestReadingContextEnumTransactionBegin  MeterValuesRequestReadingContextEnum = "Transaction.Begin"
	MeterValuesRequestReadingContextEnumTransactionEnd    MeterValuesRequestReadingContextEnum = "Transaction.End"
	MeterValuesRequestReadingContextEnumTrigger           MeterValuesRequestReadingContextEnum = "Trigger"
)

type MeterValuesRequestMeasurandEnum string

const (
	MeterValuesRequestMeasurandEnumCurrentExport                             MeterValuesRequestMeasurandEnum = "Current.Export"
	MeterValuesRequestMeasurandEnumCurrentExportOffered                      MeterValuesRequestMeasurandEnum = "Current.Export.Offered"
	MeterValuesRequestMeasurandEnumCurrentExportMinimum                      MeterValuesRequestMeasurandEnum = "Current.Export.Minimum"
	MeterValuesRequestMeasurandEnumCurrentImport                             MeterValuesRequestMeasurandEnum = "Current.Import"
	MeterValuesRequestMeasurandEnumCurrentImportOffered                      MeterValuesRequestMeasurandEnum = "Current.Import.Offered"
	MeterValuesRequestMeasurandEnumCurrentImportMinimum                      MeterValuesRequestMeasurandEnum = "Current.Import.Minimum"
	MeterValuesRequestMeasurandEnumCurrentOffered                            MeterValuesRequestMeasurandEnum = "Current.Offered"
	MeterValuesRequestMeasurandEnumDisplayPresentSOC                         MeterValuesRequestMeasurandEnum = "Display.PresentSOC"
	MeterValuesRequestMeasurandEnumDisplayMinimumSOC                         MeterValuesRequestMeasurandEnum = "Display.MinimumSOC"
	MeterValuesRequestMeasurandEnumDisplayTargetSOC                          MeterValuesRequestMeasurandEnum = "Display.TargetSOC"
	MeterValuesRequestMeasurandEnumDisplayMaximumSOC                         MeterValuesRequestMeasurandEnum = "Display.MaximumSOC"
	MeterValuesRequestMeasurandEnumDisplayRemainingTimeToMinimumSOC          MeterValuesRequestMeasurandEnum = "Display.RemainingTimeToMinimumSOC"
	MeterValuesRequestMeasurandEnumDisplayRemainingTimeToTargetSOC           MeterValuesRequestMeasurandEnum = "Display.RemainingTimeToTargetSOC"
	MeterValuesRequestMeasurandEnumDisplayRemainingTimeToMaximumSOC          MeterValuesRequestMeasurandEnum = "Display.RemainingTimeToMaximumSOC"
	MeterValuesRequestMeasurandEnumDisplayChargingComplete                   MeterValuesRequestMeasurandEnum = "Display.ChargingComplete"
	MeterValuesRequestMeasurandEnumDisplayBatteryEnergyCapacity              MeterValuesRequestMeasurandEnum = "Display.BatteryEnergyCapacity"
	MeterValuesRequestMeasurandEnumDisplayInletHot                           MeterValuesRequestMeasurandEnum = "Display.InletHot"
	MeterValuesRequestMeasurandEnumEnergyActiveExportInterval                MeterValuesRequestMeasurandEnum = "Energy.Active.Export.Interval"
	MeterValuesRequestMeasurandEnumEnergyActiveExportRegister                MeterValuesRequestMeasurandEnum = "Energy.Active.Export.Register"
	MeterValuesRequestMeasurandEnumEnergyActiveImportInterval                MeterValuesRequestMeasurandEnum = "Energy.Active.Import.Interval"
	MeterValuesRequestMeasurandEnumEnergyActiveImportRegister                MeterValuesRequestMeasurandEnum = "Energy.Active.Import.Register"
	MeterValuesRequestMeasurandEnumEnergyActiveImportCableLoss               MeterValuesRequestMeasurandEnum = "Energy.Active.Import.CableLoss"
	MeterValuesRequestMeasurandEnumEnergyActiveImportLocalGenerationRegister MeterValuesRequestMeasurandEnum = "Energy.Active.Import.LocalGeneration.Register"
	MeterValuesRequestMeasurandEnumEnergyActiveNet                           MeterValuesRequestMeasurandEnum = "Energy.Active.Net"
	MeterValuesRequestMeasurandEnumEnergyActiveSetpointInterval              MeterValuesRequestMeasurandEnum = "Energy.Active.Setpoint.Interval"
	MeterValuesRequestMeasurandEnumEnergyApparentExport                      MeterValuesRequestMeasurandEnum = "Energy.Apparent.Export"
	MeterValuesRequestMeasurandEnumEnergyApparentImport                      MeterValuesRequestMeasurandEnum = "Energy.Apparent.Import"
	MeterValuesRequestMeasurandEnumEnergyApparentNet                         MeterValuesRequestMeasurandEnum = "Energy.Apparent.Net"
	MeterValuesRequestMeasurandEnumEnergyReactiveExportInterval              MeterValuesRequestMeasurandEnum = "Energy.Reactive.Export.Interval"
	MeterValuesRequestMeasurandEnumEnergyReactiveExportRegister              MeterValuesRequestMeasurandEnum = "Energy.Reactive.Export.Register"
	MeterValuesRequestMeasurandEnumEnergyReactiveImportInterval              MeterValuesRequestMeasurandEnum = "Energy.Reactive.Import.Interval"
	MeterValuesRequestMeasurandEnumEnergyReactiveImportRegister              MeterValuesRequestMeasurandEnum = "Energy.Reactive.Import.Register"
	MeterValuesRequestMeasurandEnumEnergyReactiveNet                         MeterValuesRequestMeasurandEnum = "Energy.Reactive.Net"
	MeterValuesRequestMeasurandEnumEnergyRequestTarget                       MeterValuesRequestMeasurandEnum = "EnergyRequest.Target"
	MeterValuesRequestMeasurandEnumEnergyRequestMinimum                      MeterValuesRequestMeasurandEnum = "EnergyRequest.Minimum"
	MeterValuesRequestMeasurandEnumEnergyRequestMaximum                      MeterValuesRequestMeasurandEnum = "EnergyRequest.Maximum"
	MeterValuesRequestMeasurandEnumEnergyRequestMinimumV2X                   MeterValuesRequestMeasurandEnum = "EnergyRequest.Minimum.V2X"
	MeterValuesRequestMeasurandEnumEnergyRequestMaximumV2X                   MeterValuesRequestMeasurandEnum = "EnergyRequest.Maximum.V2X"
	MeterValuesRequestMeasurandEnumEnergyRequestBulk                         MeterValuesRequestMeasurandEnum = "EnergyRequest.Bulk"
	MeterValuesRequestMeasurandEnumFrequency                                 MeterValuesRequestMeasurandEnum = "Frequency"
	MeterValuesRequestMeasurandEnumPowerActiveExport                         MeterValuesRequestMeasurandEnum = "Power.Active.Export"
	MeterValuesRequestMeasurandEnumPowerActiveImport                         MeterValuesRequestMeasurandEnum = "Power.Active.Import"
	MeterValuesRequestMeasurandEnumPowerActiveSetpoint                       MeterValuesRequestMeasurandEnum = "Power.Active.Setpoint"
	MeterValuesRequestMeasurandEnumPowerActiveResidual                       MeterValuesRequestMeasurandEnum = "Power.Active.Residual"
	MeterValuesRequestMeasurandEnumPowerExportMinimum                        MeterValuesRequestMeasurandEnum = "Power.Export.Minimum"
	MeterValuesRequestMeasurandEnumPowerExportOffered                        MeterValuesRequestMeasurandEnum = "Power.Export.Offered"
	MeterValuesRequestMeasurandEnumPowerFactor                               MeterValuesRequestMeasurandEnum = "Power.Factor"
	MeterValuesRequestMeasurandEnumPowerImportOffered                        MeterValuesRequestMeasurandEnum = "Power.Import.Offered"
	MeterValuesRequestMeasurandEnumPowerImportMinimum                        MeterValuesRequestMeasurandEnum = "Power.Import.Minimum"
	MeterValuesRequestMeasurandEnumPowerOffered                              MeterValuesRequestMeasurandEnum = "Power.Offered"
	MeterValuesRequestMeasurandEnumPowerReactiveExport                       MeterValuesRequestMeasurandEnum = "Power.Reactive.Export"
	MeterValuesRequestMeasurandEnumPowerReactiveImport                       MeterValuesRequestMeasurandEnum = "Power.Reactive.Import"
	MeterValuesRequestMeasurandEnumSoC                                       MeterValuesRequestMeasurandEnum = "SoC"
	MeterValuesRequestMeasurandEnumVoltage                                   MeterValuesRequestMeasurandEnum = "Voltage"
	MeterValuesRequestMeasurandEnumVoltageMinimum                            MeterValuesRequestMeasurandEnum = "Voltage.Minimum"
	MeterValuesRequestMeasurandEnumVoltageMaximum                            MeterValuesRequestMeasurandEnum = "Voltage.Maximum"
)

func (MeterValuesRequest) ActionName() string { return "MeterValues" }

func (MeterValuesRequest) Version() protocol.Version { return protocol.OCPP21 }

func (MeterValuesRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (MeterValuesRequest) SchemaName() string { return "MeterValuesRequest.json" }

func (message MeterValuesRequest) Validate() error {
	return validation.Validate("MeterValuesRequest.json", schemaMeterValuesRequest, message)
}

func (MeterValuesRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("MeterValuesRequest.json", schemaMeterValuesRequest, data)
}
