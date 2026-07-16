// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = MeterValuesRequest{}

var schemaMeterValuesRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "evseId": &validation.Schema{Type: "integer", AllowAdditional: true}, "meterValue": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "sampledValue": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "value": &validation.Schema{Type: "number", AllowAdditional: true}, "context": &validation.Schema{Type: "string", Enum: []string{"Interruption.Begin", "Interruption.End", "Other", "Sample.Clock", "Sample.Periodic", "Transaction.Begin", "Transaction.End", "Trigger"}}, "measurand": &validation.Schema{Type: "string", Enum: []string{"Current.Export", "Current.Import", "Current.Offered", "Energy.Active.Export.Register", "Energy.Active.Import.Register", "Energy.Reactive.Export.Register", "Energy.Reactive.Import.Register", "Energy.Active.Export.Interval", "Energy.Active.Import.Interval", "Energy.Active.Net", "Energy.Reactive.Export.Interval", "Energy.Reactive.Import.Interval", "Energy.Reactive.Net", "Energy.Apparent.Net", "Energy.Apparent.Import", "Energy.Apparent.Export", "Frequency", "Power.Active.Export", "Power.Active.Import", "Power.Factor", "Power.Offered", "Power.Reactive.Export", "Power.Reactive.Import", "SoC", "Voltage"}}, "phase": &validation.Schema{Type: "string", Enum: []string{"L1", "L2", "L3", "N", "L1-N", "L2-N", "L3-N", "L1-L2", "L2-L3", "L3-L1"}}, "location": &validation.Schema{Type: "string", Enum: []string{"Body", "Cable", "EV", "Inlet", "Outlet"}}, "signedMeterValue": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "signedMeterData": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 2500, HasMaxLength: true}, "signingMethod": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "encodingMethod": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "publicKey": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 2500, HasMaxLength: true}}, Required: []string{"signedMeterData", "signingMethod", "encodingMethod", "publicKey"}}, "unitOfMeasure": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "unit": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "multiplier": &validation.Schema{Type: "integer", AllowAdditional: true}}}}, Required: []string{"value"}}, MinItems: 1, HasMinItems: true}, "timestamp": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}}, Required: []string{"timestamp", "sampledValue"}}, MinItems: 1, HasMinItems: true}}, Required: []string{"evseId", "meterValue"}}

type MeterValuesRequest struct {
	CustomData *MeterValuesRequestCustomData  `json:"customData,omitempty"`
	EVSEID     int                            `json:"evseId"`
	MeterValue []MeterValuesRequestMeterValue `json:"meterValue"`
}

type MeterValuesRequestMeterValue struct {
	CustomData   *MeterValuesRequestCustomData    `json:"customData,omitempty"`
	SampledValue []MeterValuesRequestSampledValue `json:"sampledValue"`
	Timestamp    string                           `json:"timestamp"`
}

type MeterValuesRequestSampledValue struct {
	CustomData       *MeterValuesRequestCustomData         `json:"customData,omitempty"`
	Value            float64                               `json:"value"`
	Context          *MeterValuesRequestReadingContextEnum `json:"context,omitempty"`
	Measurand        *MeterValuesRequestMeasurandEnum      `json:"measurand,omitempty"`
	Phase            *MeterValuesRequestPhaseEnum          `json:"phase,omitempty"`
	Location         *MeterValuesRequestLocationEnum       `json:"location,omitempty"`
	SignedMeterValue *MeterValuesRequestSignedMeterValue   `json:"signedMeterValue,omitempty"`
	UnitOfMeasure    *MeterValuesRequestUnitOfMeasure      `json:"unitOfMeasure,omitempty"`
}

type MeterValuesRequestUnitOfMeasure struct {
	CustomData *MeterValuesRequestCustomData `json:"customData,omitempty"`
	Unit       *string                       `json:"unit,omitempty"`
	Multiplier *int                          `json:"multiplier,omitempty"`
}

type MeterValuesRequestSignedMeterValue struct {
	CustomData      *MeterValuesRequestCustomData `json:"customData,omitempty"`
	SignedMeterData string                        `json:"signedMeterData"`
	SigningMethod   string                        `json:"signingMethod"`
	EncodingMethod  string                        `json:"encodingMethod"`
	PublicKey       string                        `json:"publicKey"`
}

type MeterValuesRequestLocationEnum string

const (
	MeterValuesRequestLocationEnumBody   MeterValuesRequestLocationEnum = "Body"
	MeterValuesRequestLocationEnumCable  MeterValuesRequestLocationEnum = "Cable"
	MeterValuesRequestLocationEnumEV     MeterValuesRequestLocationEnum = "EV"
	MeterValuesRequestLocationEnumInlet  MeterValuesRequestLocationEnum = "Inlet"
	MeterValuesRequestLocationEnumOutlet MeterValuesRequestLocationEnum = "Outlet"
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

type MeterValuesRequestMeasurandEnum string

const (
	MeterValuesRequestMeasurandEnumCurrentExport                MeterValuesRequestMeasurandEnum = "Current.Export"
	MeterValuesRequestMeasurandEnumCurrentImport                MeterValuesRequestMeasurandEnum = "Current.Import"
	MeterValuesRequestMeasurandEnumCurrentOffered               MeterValuesRequestMeasurandEnum = "Current.Offered"
	MeterValuesRequestMeasurandEnumEnergyActiveExportRegister   MeterValuesRequestMeasurandEnum = "Energy.Active.Export.Register"
	MeterValuesRequestMeasurandEnumEnergyActiveImportRegister   MeterValuesRequestMeasurandEnum = "Energy.Active.Import.Register"
	MeterValuesRequestMeasurandEnumEnergyReactiveExportRegister MeterValuesRequestMeasurandEnum = "Energy.Reactive.Export.Register"
	MeterValuesRequestMeasurandEnumEnergyReactiveImportRegister MeterValuesRequestMeasurandEnum = "Energy.Reactive.Import.Register"
	MeterValuesRequestMeasurandEnumEnergyActiveExportInterval   MeterValuesRequestMeasurandEnum = "Energy.Active.Export.Interval"
	MeterValuesRequestMeasurandEnumEnergyActiveImportInterval   MeterValuesRequestMeasurandEnum = "Energy.Active.Import.Interval"
	MeterValuesRequestMeasurandEnumEnergyActiveNet              MeterValuesRequestMeasurandEnum = "Energy.Active.Net"
	MeterValuesRequestMeasurandEnumEnergyReactiveExportInterval MeterValuesRequestMeasurandEnum = "Energy.Reactive.Export.Interval"
	MeterValuesRequestMeasurandEnumEnergyReactiveImportInterval MeterValuesRequestMeasurandEnum = "Energy.Reactive.Import.Interval"
	MeterValuesRequestMeasurandEnumEnergyReactiveNet            MeterValuesRequestMeasurandEnum = "Energy.Reactive.Net"
	MeterValuesRequestMeasurandEnumEnergyApparentNet            MeterValuesRequestMeasurandEnum = "Energy.Apparent.Net"
	MeterValuesRequestMeasurandEnumEnergyApparentImport         MeterValuesRequestMeasurandEnum = "Energy.Apparent.Import"
	MeterValuesRequestMeasurandEnumEnergyApparentExport         MeterValuesRequestMeasurandEnum = "Energy.Apparent.Export"
	MeterValuesRequestMeasurandEnumFrequency                    MeterValuesRequestMeasurandEnum = "Frequency"
	MeterValuesRequestMeasurandEnumPowerActiveExport            MeterValuesRequestMeasurandEnum = "Power.Active.Export"
	MeterValuesRequestMeasurandEnumPowerActiveImport            MeterValuesRequestMeasurandEnum = "Power.Active.Import"
	MeterValuesRequestMeasurandEnumPowerFactor                  MeterValuesRequestMeasurandEnum = "Power.Factor"
	MeterValuesRequestMeasurandEnumPowerOffered                 MeterValuesRequestMeasurandEnum = "Power.Offered"
	MeterValuesRequestMeasurandEnumPowerReactiveExport          MeterValuesRequestMeasurandEnum = "Power.Reactive.Export"
	MeterValuesRequestMeasurandEnumPowerReactiveImport          MeterValuesRequestMeasurandEnum = "Power.Reactive.Import"
	MeterValuesRequestMeasurandEnumSoC                          MeterValuesRequestMeasurandEnum = "SoC"
	MeterValuesRequestMeasurandEnumVoltage                      MeterValuesRequestMeasurandEnum = "Voltage"
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

type MeterValuesRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value MeterValuesRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *MeterValuesRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (MeterValuesRequest) ActionName() string { return "MeterValues" }

func (MeterValuesRequest) Version() protocol.Version { return protocol.OCPP201 }

func (MeterValuesRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (MeterValuesRequest) SchemaName() string { return "MeterValuesRequest.json" }

func (message MeterValuesRequest) Validate() error {
	return validation.Validate("MeterValuesRequest.json", schemaMeterValuesRequest, message)
}

func (MeterValuesRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("MeterValuesRequest.json", schemaMeterValuesRequest, data)
}
