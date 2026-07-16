// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = TransactionEventRequest{}

var schemaTransactionEventRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "eventType": &validation.Schema{Type: "string", Enum: []string{"Ended", "Started", "Updated"}}, "meterValue": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "sampledValue": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "value": &validation.Schema{Type: "number", AllowAdditional: true}, "context": &validation.Schema{Type: "string", Enum: []string{"Interruption.Begin", "Interruption.End", "Other", "Sample.Clock", "Sample.Periodic", "Transaction.Begin", "Transaction.End", "Trigger"}}, "measurand": &validation.Schema{Type: "string", Enum: []string{"Current.Export", "Current.Import", "Current.Offered", "Energy.Active.Export.Register", "Energy.Active.Import.Register", "Energy.Reactive.Export.Register", "Energy.Reactive.Import.Register", "Energy.Active.Export.Interval", "Energy.Active.Import.Interval", "Energy.Active.Net", "Energy.Reactive.Export.Interval", "Energy.Reactive.Import.Interval", "Energy.Reactive.Net", "Energy.Apparent.Net", "Energy.Apparent.Import", "Energy.Apparent.Export", "Frequency", "Power.Active.Export", "Power.Active.Import", "Power.Factor", "Power.Offered", "Power.Reactive.Export", "Power.Reactive.Import", "SoC", "Voltage"}}, "phase": &validation.Schema{Type: "string", Enum: []string{"L1", "L2", "L3", "N", "L1-N", "L2-N", "L3-N", "L1-L2", "L2-L3", "L3-L1"}}, "location": &validation.Schema{Type: "string", Enum: []string{"Body", "Cable", "EV", "Inlet", "Outlet"}}, "signedMeterValue": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "signedMeterData": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 2500, HasMaxLength: true}, "signingMethod": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "encodingMethod": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "publicKey": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 2500, HasMaxLength: true}}, Required: []string{"signedMeterData", "signingMethod", "encodingMethod", "publicKey"}}, "unitOfMeasure": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "unit": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "multiplier": &validation.Schema{Type: "integer", AllowAdditional: true}}}}, Required: []string{"value"}}, MinItems: 1, HasMinItems: true}, "timestamp": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}}, Required: []string{"timestamp", "sampledValue"}}, MinItems: 1, HasMinItems: true}, "timestamp": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "triggerReason": &validation.Schema{Type: "string", Enum: []string{"Authorized", "CablePluggedIn", "ChargingRateChanged", "ChargingStateChanged", "Deauthorized", "EnergyLimitReached", "EVCommunicationLost", "EVConnectTimeout", "MeterValueClock", "MeterValuePeriodic", "TimeLimitReached", "Trigger", "UnlockCommand", "StopAuthorized", "EVDeparted", "EVDetected", "RemoteStop", "RemoteStart", "AbnormalCondition", "SignedDataReceived", "ResetCommand"}}, "seqNo": &validation.Schema{Type: "integer", AllowAdditional: true}, "offline": &validation.Schema{Type: "boolean", AllowAdditional: true}, "numberOfPhasesUsed": &validation.Schema{Type: "integer", AllowAdditional: true}, "cableMaxCurrent": &validation.Schema{Type: "integer", AllowAdditional: true}, "reservationId": &validation.Schema{Type: "integer", AllowAdditional: true}, "transactionInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "transactionId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "chargingState": &validation.Schema{Type: "string", Enum: []string{"Charging", "EVConnected", "SuspendedEV", "SuspendedEVSE", "Idle"}}, "timeSpentCharging": &validation.Schema{Type: "integer", AllowAdditional: true}, "stoppedReason": &validation.Schema{Type: "string", Enum: []string{"DeAuthorized", "EmergencyStop", "EnergyLimitReached", "EVDisconnected", "GroundFault", "ImmediateReset", "Local", "LocalOutOfCredit", "MasterPass", "Other", "OvercurrentFault", "PowerLoss", "PowerQuality", "Reboot", "Remote", "SOCLimitReached", "StoppedByEV", "TimeLimitReached", "Timeout"}}, "remoteStartId": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"transactionId"}}, "evse": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "id": &validation.Schema{Type: "integer", AllowAdditional: true}, "connectorId": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"id"}}, "idToken": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "additionalInfo": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "additionalIdToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}}, Required: []string{"additionalIdToken", "type"}}, MinItems: 1, HasMinItems: true}, "idToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "type": &validation.Schema{Type: "string", Enum: []string{"Central", "eMAID", "ISO14443", "ISO15693", "KeyCode", "Local", "MacAddress", "NoAuthorization"}}}, Required: []string{"idToken", "type"}}}, Required: []string{"eventType", "timestamp", "triggerReason", "seqNo", "transactionInfo"}}

type TransactionEventRequest struct {
	CustomData         *TransactionEventRequestCustomData          `json:"customData,omitempty"`
	EventType          TransactionEventRequestTransactionEventEnum `json:"eventType"`
	MeterValue         []TransactionEventRequestMeterValue         `json:"meterValue,omitempty"`
	Timestamp          string                                      `json:"timestamp"`
	TriggerReason      TransactionEventRequestTriggerReasonEnum    `json:"triggerReason"`
	SeqNo              int                                         `json:"seqNo"`
	Offline            *bool                                       `json:"offline,omitempty"`
	NumberOfPhasesUsed *int                                        `json:"numberOfPhasesUsed,omitempty"`
	CableMaxCurrent    *int                                        `json:"cableMaxCurrent,omitempty"`
	ReservationID      *int                                        `json:"reservationId,omitempty"`
	TransactionInfo    TransactionEventRequestTransaction          `json:"transactionInfo"`
	EVSE               *TransactionEventRequestEVSE                `json:"evse,omitempty"`
	IDToken            *TransactionEventRequestIDToken             `json:"idToken,omitempty"`
}

type TransactionEventRequestIDToken struct {
	CustomData     *TransactionEventRequestCustomData      `json:"customData,omitempty"`
	AdditionalInfo []TransactionEventRequestAdditionalInfo `json:"additionalInfo,omitempty"`
	IDToken        string                                  `json:"idToken"`
	Type           TransactionEventRequestIDTokenEnum      `json:"type"`
}

type TransactionEventRequestIDTokenEnum string

const (
	TransactionEventRequestIDTokenEnumCentral         TransactionEventRequestIDTokenEnum = "Central"
	TransactionEventRequestIDTokenEnumEMAID           TransactionEventRequestIDTokenEnum = "eMAID"
	TransactionEventRequestIDTokenEnumISO14443        TransactionEventRequestIDTokenEnum = "ISO14443"
	TransactionEventRequestIDTokenEnumISO15693        TransactionEventRequestIDTokenEnum = "ISO15693"
	TransactionEventRequestIDTokenEnumKeyCode         TransactionEventRequestIDTokenEnum = "KeyCode"
	TransactionEventRequestIDTokenEnumLocal           TransactionEventRequestIDTokenEnum = "Local"
	TransactionEventRequestIDTokenEnumMacAddress      TransactionEventRequestIDTokenEnum = "MacAddress"
	TransactionEventRequestIDTokenEnumNoAuthorization TransactionEventRequestIDTokenEnum = "NoAuthorization"
)

type TransactionEventRequestAdditionalInfo struct {
	CustomData        *TransactionEventRequestCustomData `json:"customData,omitempty"`
	AdditionalIDToken string                             `json:"additionalIdToken"`
	Type              string                             `json:"type"`
}

type TransactionEventRequestEVSE struct {
	CustomData  *TransactionEventRequestCustomData `json:"customData,omitempty"`
	ID          int                                `json:"id"`
	ConnectorID *int                               `json:"connectorId,omitempty"`
}

type TransactionEventRequestTransaction struct {
	CustomData        *TransactionEventRequestCustomData        `json:"customData,omitempty"`
	TransactionID     string                                    `json:"transactionId"`
	ChargingState     *TransactionEventRequestChargingStateEnum `json:"chargingState,omitempty"`
	TimeSpentCharging *int                                      `json:"timeSpentCharging,omitempty"`
	StoppedReason     *TransactionEventRequestReasonEnum        `json:"stoppedReason,omitempty"`
	RemoteStartID     *int                                      `json:"remoteStartId,omitempty"`
}

type TransactionEventRequestReasonEnum string

const (
	TransactionEventRequestReasonEnumDeAuthorized       TransactionEventRequestReasonEnum = "DeAuthorized"
	TransactionEventRequestReasonEnumEmergencyStop      TransactionEventRequestReasonEnum = "EmergencyStop"
	TransactionEventRequestReasonEnumEnergyLimitReached TransactionEventRequestReasonEnum = "EnergyLimitReached"
	TransactionEventRequestReasonEnumEVDisconnected     TransactionEventRequestReasonEnum = "EVDisconnected"
	TransactionEventRequestReasonEnumGroundFault        TransactionEventRequestReasonEnum = "GroundFault"
	TransactionEventRequestReasonEnumImmediateReset     TransactionEventRequestReasonEnum = "ImmediateReset"
	TransactionEventRequestReasonEnumLocal              TransactionEventRequestReasonEnum = "Local"
	TransactionEventRequestReasonEnumLocalOutOfCredit   TransactionEventRequestReasonEnum = "LocalOutOfCredit"
	TransactionEventRequestReasonEnumMasterPass         TransactionEventRequestReasonEnum = "MasterPass"
	TransactionEventRequestReasonEnumOther              TransactionEventRequestReasonEnum = "Other"
	TransactionEventRequestReasonEnumOvercurrentFault   TransactionEventRequestReasonEnum = "OvercurrentFault"
	TransactionEventRequestReasonEnumPowerLoss          TransactionEventRequestReasonEnum = "PowerLoss"
	TransactionEventRequestReasonEnumPowerQuality       TransactionEventRequestReasonEnum = "PowerQuality"
	TransactionEventRequestReasonEnumReboot             TransactionEventRequestReasonEnum = "Reboot"
	TransactionEventRequestReasonEnumRemote             TransactionEventRequestReasonEnum = "Remote"
	TransactionEventRequestReasonEnumSOCLimitReached    TransactionEventRequestReasonEnum = "SOCLimitReached"
	TransactionEventRequestReasonEnumStoppedByEV        TransactionEventRequestReasonEnum = "StoppedByEV"
	TransactionEventRequestReasonEnumTimeLimitReached   TransactionEventRequestReasonEnum = "TimeLimitReached"
	TransactionEventRequestReasonEnumTimeout            TransactionEventRequestReasonEnum = "Timeout"
)

type TransactionEventRequestChargingStateEnum string

const (
	TransactionEventRequestChargingStateEnumCharging      TransactionEventRequestChargingStateEnum = "Charging"
	TransactionEventRequestChargingStateEnumEVConnected   TransactionEventRequestChargingStateEnum = "EVConnected"
	TransactionEventRequestChargingStateEnumSuspendedEV   TransactionEventRequestChargingStateEnum = "SuspendedEV"
	TransactionEventRequestChargingStateEnumSuspendedEVSE TransactionEventRequestChargingStateEnum = "SuspendedEVSE"
	TransactionEventRequestChargingStateEnumIdle          TransactionEventRequestChargingStateEnum = "Idle"
)

type TransactionEventRequestTriggerReasonEnum string

const (
	TransactionEventRequestTriggerReasonEnumAuthorized           TransactionEventRequestTriggerReasonEnum = "Authorized"
	TransactionEventRequestTriggerReasonEnumCablePluggedIn       TransactionEventRequestTriggerReasonEnum = "CablePluggedIn"
	TransactionEventRequestTriggerReasonEnumChargingRateChanged  TransactionEventRequestTriggerReasonEnum = "ChargingRateChanged"
	TransactionEventRequestTriggerReasonEnumChargingStateChanged TransactionEventRequestTriggerReasonEnum = "ChargingStateChanged"
	TransactionEventRequestTriggerReasonEnumDeauthorized         TransactionEventRequestTriggerReasonEnum = "Deauthorized"
	TransactionEventRequestTriggerReasonEnumEnergyLimitReached   TransactionEventRequestTriggerReasonEnum = "EnergyLimitReached"
	TransactionEventRequestTriggerReasonEnumEVCommunicationLost  TransactionEventRequestTriggerReasonEnum = "EVCommunicationLost"
	TransactionEventRequestTriggerReasonEnumEVConnectTimeout     TransactionEventRequestTriggerReasonEnum = "EVConnectTimeout"
	TransactionEventRequestTriggerReasonEnumMeterValueClock      TransactionEventRequestTriggerReasonEnum = "MeterValueClock"
	TransactionEventRequestTriggerReasonEnumMeterValuePeriodic   TransactionEventRequestTriggerReasonEnum = "MeterValuePeriodic"
	TransactionEventRequestTriggerReasonEnumTimeLimitReached     TransactionEventRequestTriggerReasonEnum = "TimeLimitReached"
	TransactionEventRequestTriggerReasonEnumTrigger              TransactionEventRequestTriggerReasonEnum = "Trigger"
	TransactionEventRequestTriggerReasonEnumUnlockCommand        TransactionEventRequestTriggerReasonEnum = "UnlockCommand"
	TransactionEventRequestTriggerReasonEnumStopAuthorized       TransactionEventRequestTriggerReasonEnum = "StopAuthorized"
	TransactionEventRequestTriggerReasonEnumEVDeparted           TransactionEventRequestTriggerReasonEnum = "EVDeparted"
	TransactionEventRequestTriggerReasonEnumEVDetected           TransactionEventRequestTriggerReasonEnum = "EVDetected"
	TransactionEventRequestTriggerReasonEnumRemoteStop           TransactionEventRequestTriggerReasonEnum = "RemoteStop"
	TransactionEventRequestTriggerReasonEnumRemoteStart          TransactionEventRequestTriggerReasonEnum = "RemoteStart"
	TransactionEventRequestTriggerReasonEnumAbnormalCondition    TransactionEventRequestTriggerReasonEnum = "AbnormalCondition"
	TransactionEventRequestTriggerReasonEnumSignedDataReceived   TransactionEventRequestTriggerReasonEnum = "SignedDataReceived"
	TransactionEventRequestTriggerReasonEnumResetCommand         TransactionEventRequestTriggerReasonEnum = "ResetCommand"
)

type TransactionEventRequestMeterValue struct {
	CustomData   *TransactionEventRequestCustomData    `json:"customData,omitempty"`
	SampledValue []TransactionEventRequestSampledValue `json:"sampledValue"`
	Timestamp    string                                `json:"timestamp"`
}

type TransactionEventRequestSampledValue struct {
	CustomData       *TransactionEventRequestCustomData         `json:"customData,omitempty"`
	Value            float64                                    `json:"value"`
	Context          *TransactionEventRequestReadingContextEnum `json:"context,omitempty"`
	Measurand        *TransactionEventRequestMeasurandEnum      `json:"measurand,omitempty"`
	Phase            *TransactionEventRequestPhaseEnum          `json:"phase,omitempty"`
	Location         *TransactionEventRequestLocationEnum       `json:"location,omitempty"`
	SignedMeterValue *TransactionEventRequestSignedMeterValue   `json:"signedMeterValue,omitempty"`
	UnitOfMeasure    *TransactionEventRequestUnitOfMeasure      `json:"unitOfMeasure,omitempty"`
}

type TransactionEventRequestUnitOfMeasure struct {
	CustomData *TransactionEventRequestCustomData `json:"customData,omitempty"`
	Unit       *string                            `json:"unit,omitempty"`
	Multiplier *int                               `json:"multiplier,omitempty"`
}

type TransactionEventRequestSignedMeterValue struct {
	CustomData      *TransactionEventRequestCustomData `json:"customData,omitempty"`
	SignedMeterData string                             `json:"signedMeterData"`
	SigningMethod   string                             `json:"signingMethod"`
	EncodingMethod  string                             `json:"encodingMethod"`
	PublicKey       string                             `json:"publicKey"`
}

type TransactionEventRequestLocationEnum string

const (
	TransactionEventRequestLocationEnumBody   TransactionEventRequestLocationEnum = "Body"
	TransactionEventRequestLocationEnumCable  TransactionEventRequestLocationEnum = "Cable"
	TransactionEventRequestLocationEnumEV     TransactionEventRequestLocationEnum = "EV"
	TransactionEventRequestLocationEnumInlet  TransactionEventRequestLocationEnum = "Inlet"
	TransactionEventRequestLocationEnumOutlet TransactionEventRequestLocationEnum = "Outlet"
)

type TransactionEventRequestPhaseEnum string

const (
	TransactionEventRequestPhaseEnumL1   TransactionEventRequestPhaseEnum = "L1"
	TransactionEventRequestPhaseEnumL2   TransactionEventRequestPhaseEnum = "L2"
	TransactionEventRequestPhaseEnumL3   TransactionEventRequestPhaseEnum = "L3"
	TransactionEventRequestPhaseEnumN    TransactionEventRequestPhaseEnum = "N"
	TransactionEventRequestPhaseEnumL1N  TransactionEventRequestPhaseEnum = "L1-N"
	TransactionEventRequestPhaseEnumL2N  TransactionEventRequestPhaseEnum = "L2-N"
	TransactionEventRequestPhaseEnumL3N  TransactionEventRequestPhaseEnum = "L3-N"
	TransactionEventRequestPhaseEnumL1L2 TransactionEventRequestPhaseEnum = "L1-L2"
	TransactionEventRequestPhaseEnumL2L3 TransactionEventRequestPhaseEnum = "L2-L3"
	TransactionEventRequestPhaseEnumL3L1 TransactionEventRequestPhaseEnum = "L3-L1"
)

type TransactionEventRequestMeasurandEnum string

const (
	TransactionEventRequestMeasurandEnumCurrentExport                TransactionEventRequestMeasurandEnum = "Current.Export"
	TransactionEventRequestMeasurandEnumCurrentImport                TransactionEventRequestMeasurandEnum = "Current.Import"
	TransactionEventRequestMeasurandEnumCurrentOffered               TransactionEventRequestMeasurandEnum = "Current.Offered"
	TransactionEventRequestMeasurandEnumEnergyActiveExportRegister   TransactionEventRequestMeasurandEnum = "Energy.Active.Export.Register"
	TransactionEventRequestMeasurandEnumEnergyActiveImportRegister   TransactionEventRequestMeasurandEnum = "Energy.Active.Import.Register"
	TransactionEventRequestMeasurandEnumEnergyReactiveExportRegister TransactionEventRequestMeasurandEnum = "Energy.Reactive.Export.Register"
	TransactionEventRequestMeasurandEnumEnergyReactiveImportRegister TransactionEventRequestMeasurandEnum = "Energy.Reactive.Import.Register"
	TransactionEventRequestMeasurandEnumEnergyActiveExportInterval   TransactionEventRequestMeasurandEnum = "Energy.Active.Export.Interval"
	TransactionEventRequestMeasurandEnumEnergyActiveImportInterval   TransactionEventRequestMeasurandEnum = "Energy.Active.Import.Interval"
	TransactionEventRequestMeasurandEnumEnergyActiveNet              TransactionEventRequestMeasurandEnum = "Energy.Active.Net"
	TransactionEventRequestMeasurandEnumEnergyReactiveExportInterval TransactionEventRequestMeasurandEnum = "Energy.Reactive.Export.Interval"
	TransactionEventRequestMeasurandEnumEnergyReactiveImportInterval TransactionEventRequestMeasurandEnum = "Energy.Reactive.Import.Interval"
	TransactionEventRequestMeasurandEnumEnergyReactiveNet            TransactionEventRequestMeasurandEnum = "Energy.Reactive.Net"
	TransactionEventRequestMeasurandEnumEnergyApparentNet            TransactionEventRequestMeasurandEnum = "Energy.Apparent.Net"
	TransactionEventRequestMeasurandEnumEnergyApparentImport         TransactionEventRequestMeasurandEnum = "Energy.Apparent.Import"
	TransactionEventRequestMeasurandEnumEnergyApparentExport         TransactionEventRequestMeasurandEnum = "Energy.Apparent.Export"
	TransactionEventRequestMeasurandEnumFrequency                    TransactionEventRequestMeasurandEnum = "Frequency"
	TransactionEventRequestMeasurandEnumPowerActiveExport            TransactionEventRequestMeasurandEnum = "Power.Active.Export"
	TransactionEventRequestMeasurandEnumPowerActiveImport            TransactionEventRequestMeasurandEnum = "Power.Active.Import"
	TransactionEventRequestMeasurandEnumPowerFactor                  TransactionEventRequestMeasurandEnum = "Power.Factor"
	TransactionEventRequestMeasurandEnumPowerOffered                 TransactionEventRequestMeasurandEnum = "Power.Offered"
	TransactionEventRequestMeasurandEnumPowerReactiveExport          TransactionEventRequestMeasurandEnum = "Power.Reactive.Export"
	TransactionEventRequestMeasurandEnumPowerReactiveImport          TransactionEventRequestMeasurandEnum = "Power.Reactive.Import"
	TransactionEventRequestMeasurandEnumSoC                          TransactionEventRequestMeasurandEnum = "SoC"
	TransactionEventRequestMeasurandEnumVoltage                      TransactionEventRequestMeasurandEnum = "Voltage"
)

type TransactionEventRequestReadingContextEnum string

const (
	TransactionEventRequestReadingContextEnumInterruptionBegin TransactionEventRequestReadingContextEnum = "Interruption.Begin"
	TransactionEventRequestReadingContextEnumInterruptionEnd   TransactionEventRequestReadingContextEnum = "Interruption.End"
	TransactionEventRequestReadingContextEnumOther             TransactionEventRequestReadingContextEnum = "Other"
	TransactionEventRequestReadingContextEnumSampleClock       TransactionEventRequestReadingContextEnum = "Sample.Clock"
	TransactionEventRequestReadingContextEnumSamplePeriodic    TransactionEventRequestReadingContextEnum = "Sample.Periodic"
	TransactionEventRequestReadingContextEnumTransactionBegin  TransactionEventRequestReadingContextEnum = "Transaction.Begin"
	TransactionEventRequestReadingContextEnumTransactionEnd    TransactionEventRequestReadingContextEnum = "Transaction.End"
	TransactionEventRequestReadingContextEnumTrigger           TransactionEventRequestReadingContextEnum = "Trigger"
)

type TransactionEventRequestTransactionEventEnum string

const (
	TransactionEventRequestTransactionEventEnumEnded   TransactionEventRequestTransactionEventEnum = "Ended"
	TransactionEventRequestTransactionEventEnumStarted TransactionEventRequestTransactionEventEnum = "Started"
	TransactionEventRequestTransactionEventEnumUpdated TransactionEventRequestTransactionEventEnum = "Updated"
)

type TransactionEventRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value TransactionEventRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *TransactionEventRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (TransactionEventRequest) ActionName() string { return "TransactionEvent" }

func (TransactionEventRequest) Version() protocol.Version { return protocol.OCPP201 }

func (TransactionEventRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (TransactionEventRequest) SchemaName() string { return "TransactionEventRequest.json" }

func (message TransactionEventRequest) Validate() error {
	return validation.Validate("TransactionEventRequest.json", schemaTransactionEventRequest, message)
}

func (TransactionEventRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("TransactionEventRequest.json", schemaTransactionEventRequest, data)
}
