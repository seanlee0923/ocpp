// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = TransactionEventRequest{}

var schemaTransactionEventRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"costDetails": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"chargingPeriods": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"dimensions": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"type": &validation.Schema{Type: "string", Enum: []string{"Energy", "MaxCurrent", "MinCurrent", "MaxPower", "MinPower", "IdleTIme", "ChargingTime"}}, "volume": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"type", "volume"}}, MinItems: 1, HasMinItems: true}, "tariffId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 60, HasMaxLength: true}, "startPeriod": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"startPeriod"}}, MinItems: 1, HasMinItems: true}, "totalCost": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"currency": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 3, HasMaxLength: true}, "typeOfCost": &validation.Schema{Type: "string", Enum: []string{"NormalCost", "MinCost", "MaxCost"}}, "fixed": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exclTax": &validation.Schema{Type: "number", AllowAdditional: true}, "inclTax": &validation.Schema{Type: "number", AllowAdditional: true}, "taxRates": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "tax": &validation.Schema{Type: "number", AllowAdditional: true}, "stack": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"type", "tax"}}, MinItems: 1, HasMinItems: true, MaxItems: 5, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "energy": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exclTax": &validation.Schema{Type: "number", AllowAdditional: true}, "inclTax": &validation.Schema{Type: "number", AllowAdditional: true}, "taxRates": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "tax": &validation.Schema{Type: "number", AllowAdditional: true}, "stack": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"type", "tax"}}, MinItems: 1, HasMinItems: true, MaxItems: 5, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "chargingTime": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exclTax": &validation.Schema{Type: "number", AllowAdditional: true}, "inclTax": &validation.Schema{Type: "number", AllowAdditional: true}, "taxRates": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "tax": &validation.Schema{Type: "number", AllowAdditional: true}, "stack": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"type", "tax"}}, MinItems: 1, HasMinItems: true, MaxItems: 5, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "idleTime": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exclTax": &validation.Schema{Type: "number", AllowAdditional: true}, "inclTax": &validation.Schema{Type: "number", AllowAdditional: true}, "taxRates": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "tax": &validation.Schema{Type: "number", AllowAdditional: true}, "stack": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"type", "tax"}}, MinItems: 1, HasMinItems: true, MaxItems: 5, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "reservationTime": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exclTax": &validation.Schema{Type: "number", AllowAdditional: true}, "inclTax": &validation.Schema{Type: "number", AllowAdditional: true}, "taxRates": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "tax": &validation.Schema{Type: "number", AllowAdditional: true}, "stack": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"type", "tax"}}, MinItems: 1, HasMinItems: true, MaxItems: 5, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "reservationFixed": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exclTax": &validation.Schema{Type: "number", AllowAdditional: true}, "inclTax": &validation.Schema{Type: "number", AllowAdditional: true}, "taxRates": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "tax": &validation.Schema{Type: "number", AllowAdditional: true}, "stack": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"type", "tax"}}, MinItems: 1, HasMinItems: true, MaxItems: 5, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "total": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exclTax": &validation.Schema{Type: "number", AllowAdditional: true}, "inclTax": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"currency", "typeOfCost", "total"}}, "totalUsage": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"energy": &validation.Schema{Type: "number", AllowAdditional: true}, "chargingTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "idleTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "reservationTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"energy", "chargingTime", "idleTime"}}, "failureToCalculate": &validation.Schema{Type: "boolean", AllowAdditional: true}, "failureReason": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 500, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"totalCost", "totalUsage"}}, "eventType": &validation.Schema{Type: "string", Enum: []string{"Ended", "Started", "Updated"}}, "meterValue": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"sampledValue": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"value": &validation.Schema{Type: "number", AllowAdditional: true}, "measurand": &validation.Schema{Type: "string", Enum: []string{"Current.Export", "Current.Export.Offered", "Current.Export.Minimum", "Current.Import", "Current.Import.Offered", "Current.Import.Minimum", "Current.Offered", "Display.PresentSOC", "Display.MinimumSOC", "Display.TargetSOC", "Display.MaximumSOC", "Display.RemainingTimeToMinimumSOC", "Display.RemainingTimeToTargetSOC", "Display.RemainingTimeToMaximumSOC", "Display.ChargingComplete", "Display.BatteryEnergyCapacity", "Display.InletHot", "Energy.Active.Export.Interval", "Energy.Active.Export.Register", "Energy.Active.Import.Interval", "Energy.Active.Import.Register", "Energy.Active.Import.CableLoss", "Energy.Active.Import.LocalGeneration.Register", "Energy.Active.Net", "Energy.Active.Setpoint.Interval", "Energy.Apparent.Export", "Energy.Apparent.Import", "Energy.Apparent.Net", "Energy.Reactive.Export.Interval", "Energy.Reactive.Export.Register", "Energy.Reactive.Import.Interval", "Energy.Reactive.Import.Register", "Energy.Reactive.Net", "EnergyRequest.Target", "EnergyRequest.Minimum", "EnergyRequest.Maximum", "EnergyRequest.Minimum.V2X", "EnergyRequest.Maximum.V2X", "EnergyRequest.Bulk", "Frequency", "Power.Active.Export", "Power.Active.Import", "Power.Active.Setpoint", "Power.Active.Residual", "Power.Export.Minimum", "Power.Export.Offered", "Power.Factor", "Power.Import.Offered", "Power.Import.Minimum", "Power.Offered", "Power.Reactive.Export", "Power.Reactive.Import", "SoC", "Voltage", "Voltage.Minimum", "Voltage.Maximum"}}, "context": &validation.Schema{Type: "string", Enum: []string{"Interruption.Begin", "Interruption.End", "Other", "Sample.Clock", "Sample.Periodic", "Transaction.Begin", "Transaction.End", "Trigger"}}, "phase": &validation.Schema{Type: "string", Enum: []string{"L1", "L2", "L3", "N", "L1-N", "L2-N", "L3-N", "L1-L2", "L2-L3", "L3-L1"}}, "location": &validation.Schema{Type: "string", Enum: []string{"Body", "Cable", "EV", "Inlet", "Outlet", "Upstream"}}, "signedMeterValue": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"signedMeterData": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 32768, HasMaxLength: true}, "signingMethod": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "encodingMethod": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "publicKey": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 2500, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"signedMeterData", "encodingMethod"}}, "unitOfMeasure": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"unit": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "multiplier": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"value"}}, MinItems: 1, HasMinItems: true}, "timestamp": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"timestamp", "sampledValue"}}, MinItems: 1, HasMinItems: true}, "timestamp": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "triggerReason": &validation.Schema{Type: "string", Enum: []string{"AbnormalCondition", "Authorized", "CablePluggedIn", "ChargingRateChanged", "ChargingStateChanged", "CostLimitReached", "Deauthorized", "EnergyLimitReached", "EVCommunicationLost", "EVConnectTimeout", "EVDeparted", "EVDetected", "LimitSet", "MeterValueClock", "MeterValuePeriodic", "OperationModeChanged", "RemoteStart", "RemoteStop", "ResetCommand", "RunningCost", "SignedDataReceived", "SoCLimitReached", "StopAuthorized", "TariffChanged", "TariffNotAccepted", "TimeLimitReached", "Trigger", "TxResumed", "UnlockCommand"}}, "seqNo": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "offline": &validation.Schema{Type: "boolean", AllowAdditional: true}, "numberOfPhasesUsed": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true, Maximum: 3, HasMaximum: true}, "cableMaxCurrent": &validation.Schema{Type: "integer", AllowAdditional: true}, "reservationId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "preconditioningStatus": &validation.Schema{Type: "string", Enum: []string{"Unknown", "Ready", "NotReady", "Preconditioning"}}, "evseSleep": &validation.Schema{Type: "boolean", AllowAdditional: true}, "transactionInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"transactionId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "chargingState": &validation.Schema{Type: "string", Enum: []string{"EVConnected", "Charging", "SuspendedEV", "SuspendedEVSE", "Idle"}}, "timeSpentCharging": &validation.Schema{Type: "integer", AllowAdditional: true}, "stoppedReason": &validation.Schema{Type: "string", Enum: []string{"DeAuthorized", "EmergencyStop", "EnergyLimitReached", "EVDisconnected", "GroundFault", "ImmediateReset", "MasterPass", "Local", "LocalOutOfCredit", "Other", "OvercurrentFault", "PowerLoss", "PowerQuality", "Reboot", "Remote", "SOCLimitReached", "StoppedByEV", "TimeLimitReached", "Timeout", "ReqEnergyTransferRejected"}}, "remoteStartId": &validation.Schema{Type: "integer", AllowAdditional: true}, "operationMode": &validation.Schema{Type: "string", Enum: []string{"Idle", "ChargingOnly", "CentralSetpoint", "ExternalSetpoint", "ExternalLimits", "CentralFrequency", "LocalFrequency", "LocalLoadBalancing"}}, "tariffId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 60, HasMaxLength: true}, "transactionLimit": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"maxCost": &validation.Schema{Type: "number", AllowAdditional: true}, "maxEnergy": &validation.Schema{Type: "number", AllowAdditional: true}, "maxTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "maxSoC": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true, Maximum: 100, HasMaximum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"transactionId"}}, "evse": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"id": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "connectorId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id"}}, "idToken": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"additionalInfo": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"additionalIdToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}, "type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"additionalIdToken", "type"}}, MinItems: 1, HasMinItems: true}, "idToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}, "type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"idToken", "type"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"eventType", "timestamp", "triggerReason", "seqNo", "transactionInfo"}}

type TransactionEventRequest struct {
	CostDetails           *TransactionEventRequestCostDetails               `json:"costDetails,omitempty"`
	EventType             TransactionEventRequestTransactionEventEnum       `json:"eventType"`
	MeterValue            []TransactionEventRequestMeterValue               `json:"meterValue,omitempty"`
	Timestamp             string                                            `json:"timestamp"`
	TriggerReason         TransactionEventRequestTriggerReasonEnum          `json:"triggerReason"`
	SeqNo                 int                                               `json:"seqNo"`
	Offline               *bool                                             `json:"offline,omitempty"`
	NumberOfPhasesUsed    *int                                              `json:"numberOfPhasesUsed,omitempty"`
	CableMaxCurrent       *int                                              `json:"cableMaxCurrent,omitempty"`
	ReservationID         *int                                              `json:"reservationId,omitempty"`
	PreconditioningStatus *TransactionEventRequestPreconditioningStatusEnum `json:"preconditioningStatus,omitempty"`
	EVSESleep             *bool                                             `json:"evseSleep,omitempty"`
	TransactionInfo       TransactionEventRequestTransaction                `json:"transactionInfo"`
	EVSE                  *TransactionEventRequestEVSE                      `json:"evse,omitempty"`
	IDToken               *TransactionEventRequestIDToken                   `json:"idToken,omitempty"`
	CustomData            *TransactionEventRequestCustomData                `json:"customData,omitempty"`
}

type TransactionEventRequestIDToken struct {
	AdditionalInfo []TransactionEventRequestAdditionalInfo `json:"additionalInfo,omitempty"`
	IDToken        string                                  `json:"idToken"`
	Type           string                                  `json:"type"`
	CustomData     *TransactionEventRequestCustomData      `json:"customData,omitempty"`
}

type TransactionEventRequestAdditionalInfo struct {
	AdditionalIDToken string                             `json:"additionalIdToken"`
	Type              string                             `json:"type"`
	CustomData        *TransactionEventRequestCustomData `json:"customData,omitempty"`
}

type TransactionEventRequestEVSE struct {
	ID          int                                `json:"id"`
	ConnectorID *int                               `json:"connectorId,omitempty"`
	CustomData  *TransactionEventRequestCustomData `json:"customData,omitempty"`
}

type TransactionEventRequestTransaction struct {
	TransactionID     string                                    `json:"transactionId"`
	ChargingState     *TransactionEventRequestChargingStateEnum `json:"chargingState,omitempty"`
	TimeSpentCharging *int                                      `json:"timeSpentCharging,omitempty"`
	StoppedReason     *TransactionEventRequestReasonEnum        `json:"stoppedReason,omitempty"`
	RemoteStartID     *int                                      `json:"remoteStartId,omitempty"`
	OperationMode     *TransactionEventRequestOperationModeEnum `json:"operationMode,omitempty"`
	TariffID          *string                                   `json:"tariffId,omitempty"`
	TransactionLimit  *TransactionEventRequestTransactionLimit  `json:"transactionLimit,omitempty"`
	CustomData        *TransactionEventRequestCustomData        `json:"customData,omitempty"`
}

type TransactionEventRequestTransactionLimit struct {
	MaxCost    *float64                           `json:"maxCost,omitempty"`
	MaxEnergy  *float64                           `json:"maxEnergy,omitempty"`
	MaxTime    *int                               `json:"maxTime,omitempty"`
	MaxSoC     *int                               `json:"maxSoC,omitempty"`
	CustomData *TransactionEventRequestCustomData `json:"customData,omitempty"`
}

type TransactionEventRequestOperationModeEnum string

const (
	TransactionEventRequestOperationModeEnumIdle               TransactionEventRequestOperationModeEnum = "Idle"
	TransactionEventRequestOperationModeEnumChargingOnly       TransactionEventRequestOperationModeEnum = "ChargingOnly"
	TransactionEventRequestOperationModeEnumCentralSetpoint    TransactionEventRequestOperationModeEnum = "CentralSetpoint"
	TransactionEventRequestOperationModeEnumExternalSetpoint   TransactionEventRequestOperationModeEnum = "ExternalSetpoint"
	TransactionEventRequestOperationModeEnumExternalLimits     TransactionEventRequestOperationModeEnum = "ExternalLimits"
	TransactionEventRequestOperationModeEnumCentralFrequency   TransactionEventRequestOperationModeEnum = "CentralFrequency"
	TransactionEventRequestOperationModeEnumLocalFrequency     TransactionEventRequestOperationModeEnum = "LocalFrequency"
	TransactionEventRequestOperationModeEnumLocalLoadBalancing TransactionEventRequestOperationModeEnum = "LocalLoadBalancing"
)

type TransactionEventRequestReasonEnum string

const (
	TransactionEventRequestReasonEnumDeAuthorized              TransactionEventRequestReasonEnum = "DeAuthorized"
	TransactionEventRequestReasonEnumEmergencyStop             TransactionEventRequestReasonEnum = "EmergencyStop"
	TransactionEventRequestReasonEnumEnergyLimitReached        TransactionEventRequestReasonEnum = "EnergyLimitReached"
	TransactionEventRequestReasonEnumEVDisconnected            TransactionEventRequestReasonEnum = "EVDisconnected"
	TransactionEventRequestReasonEnumGroundFault               TransactionEventRequestReasonEnum = "GroundFault"
	TransactionEventRequestReasonEnumImmediateReset            TransactionEventRequestReasonEnum = "ImmediateReset"
	TransactionEventRequestReasonEnumMasterPass                TransactionEventRequestReasonEnum = "MasterPass"
	TransactionEventRequestReasonEnumLocal                     TransactionEventRequestReasonEnum = "Local"
	TransactionEventRequestReasonEnumLocalOutOfCredit          TransactionEventRequestReasonEnum = "LocalOutOfCredit"
	TransactionEventRequestReasonEnumOther                     TransactionEventRequestReasonEnum = "Other"
	TransactionEventRequestReasonEnumOvercurrentFault          TransactionEventRequestReasonEnum = "OvercurrentFault"
	TransactionEventRequestReasonEnumPowerLoss                 TransactionEventRequestReasonEnum = "PowerLoss"
	TransactionEventRequestReasonEnumPowerQuality              TransactionEventRequestReasonEnum = "PowerQuality"
	TransactionEventRequestReasonEnumReboot                    TransactionEventRequestReasonEnum = "Reboot"
	TransactionEventRequestReasonEnumRemote                    TransactionEventRequestReasonEnum = "Remote"
	TransactionEventRequestReasonEnumSOCLimitReached           TransactionEventRequestReasonEnum = "SOCLimitReached"
	TransactionEventRequestReasonEnumStoppedByEV               TransactionEventRequestReasonEnum = "StoppedByEV"
	TransactionEventRequestReasonEnumTimeLimitReached          TransactionEventRequestReasonEnum = "TimeLimitReached"
	TransactionEventRequestReasonEnumTimeout                   TransactionEventRequestReasonEnum = "Timeout"
	TransactionEventRequestReasonEnumReqEnergyTransferRejected TransactionEventRequestReasonEnum = "ReqEnergyTransferRejected"
)

type TransactionEventRequestChargingStateEnum string

const (
	TransactionEventRequestChargingStateEnumEVConnected   TransactionEventRequestChargingStateEnum = "EVConnected"
	TransactionEventRequestChargingStateEnumCharging      TransactionEventRequestChargingStateEnum = "Charging"
	TransactionEventRequestChargingStateEnumSuspendedEV   TransactionEventRequestChargingStateEnum = "SuspendedEV"
	TransactionEventRequestChargingStateEnumSuspendedEVSE TransactionEventRequestChargingStateEnum = "SuspendedEVSE"
	TransactionEventRequestChargingStateEnumIdle          TransactionEventRequestChargingStateEnum = "Idle"
)

type TransactionEventRequestPreconditioningStatusEnum string

const (
	TransactionEventRequestPreconditioningStatusEnumUnknown         TransactionEventRequestPreconditioningStatusEnum = "Unknown"
	TransactionEventRequestPreconditioningStatusEnumReady           TransactionEventRequestPreconditioningStatusEnum = "Ready"
	TransactionEventRequestPreconditioningStatusEnumNotReady        TransactionEventRequestPreconditioningStatusEnum = "NotReady"
	TransactionEventRequestPreconditioningStatusEnumPreconditioning TransactionEventRequestPreconditioningStatusEnum = "Preconditioning"
)

type TransactionEventRequestTriggerReasonEnum string

const (
	TransactionEventRequestTriggerReasonEnumAbnormalCondition    TransactionEventRequestTriggerReasonEnum = "AbnormalCondition"
	TransactionEventRequestTriggerReasonEnumAuthorized           TransactionEventRequestTriggerReasonEnum = "Authorized"
	TransactionEventRequestTriggerReasonEnumCablePluggedIn       TransactionEventRequestTriggerReasonEnum = "CablePluggedIn"
	TransactionEventRequestTriggerReasonEnumChargingRateChanged  TransactionEventRequestTriggerReasonEnum = "ChargingRateChanged"
	TransactionEventRequestTriggerReasonEnumChargingStateChanged TransactionEventRequestTriggerReasonEnum = "ChargingStateChanged"
	TransactionEventRequestTriggerReasonEnumCostLimitReached     TransactionEventRequestTriggerReasonEnum = "CostLimitReached"
	TransactionEventRequestTriggerReasonEnumDeauthorized         TransactionEventRequestTriggerReasonEnum = "Deauthorized"
	TransactionEventRequestTriggerReasonEnumEnergyLimitReached   TransactionEventRequestTriggerReasonEnum = "EnergyLimitReached"
	TransactionEventRequestTriggerReasonEnumEVCommunicationLost  TransactionEventRequestTriggerReasonEnum = "EVCommunicationLost"
	TransactionEventRequestTriggerReasonEnumEVConnectTimeout     TransactionEventRequestTriggerReasonEnum = "EVConnectTimeout"
	TransactionEventRequestTriggerReasonEnumEVDeparted           TransactionEventRequestTriggerReasonEnum = "EVDeparted"
	TransactionEventRequestTriggerReasonEnumEVDetected           TransactionEventRequestTriggerReasonEnum = "EVDetected"
	TransactionEventRequestTriggerReasonEnumLimitSet             TransactionEventRequestTriggerReasonEnum = "LimitSet"
	TransactionEventRequestTriggerReasonEnumMeterValueClock      TransactionEventRequestTriggerReasonEnum = "MeterValueClock"
	TransactionEventRequestTriggerReasonEnumMeterValuePeriodic   TransactionEventRequestTriggerReasonEnum = "MeterValuePeriodic"
	TransactionEventRequestTriggerReasonEnumOperationModeChanged TransactionEventRequestTriggerReasonEnum = "OperationModeChanged"
	TransactionEventRequestTriggerReasonEnumRemoteStart          TransactionEventRequestTriggerReasonEnum = "RemoteStart"
	TransactionEventRequestTriggerReasonEnumRemoteStop           TransactionEventRequestTriggerReasonEnum = "RemoteStop"
	TransactionEventRequestTriggerReasonEnumResetCommand         TransactionEventRequestTriggerReasonEnum = "ResetCommand"
	TransactionEventRequestTriggerReasonEnumRunningCost          TransactionEventRequestTriggerReasonEnum = "RunningCost"
	TransactionEventRequestTriggerReasonEnumSignedDataReceived   TransactionEventRequestTriggerReasonEnum = "SignedDataReceived"
	TransactionEventRequestTriggerReasonEnumSoCLimitReached      TransactionEventRequestTriggerReasonEnum = "SoCLimitReached"
	TransactionEventRequestTriggerReasonEnumStopAuthorized       TransactionEventRequestTriggerReasonEnum = "StopAuthorized"
	TransactionEventRequestTriggerReasonEnumTariffChanged        TransactionEventRequestTriggerReasonEnum = "TariffChanged"
	TransactionEventRequestTriggerReasonEnumTariffNotAccepted    TransactionEventRequestTriggerReasonEnum = "TariffNotAccepted"
	TransactionEventRequestTriggerReasonEnumTimeLimitReached     TransactionEventRequestTriggerReasonEnum = "TimeLimitReached"
	TransactionEventRequestTriggerReasonEnumTrigger              TransactionEventRequestTriggerReasonEnum = "Trigger"
	TransactionEventRequestTriggerReasonEnumTxResumed            TransactionEventRequestTriggerReasonEnum = "TxResumed"
	TransactionEventRequestTriggerReasonEnumUnlockCommand        TransactionEventRequestTriggerReasonEnum = "UnlockCommand"
)

type TransactionEventRequestMeterValue struct {
	SampledValue []TransactionEventRequestSampledValue `json:"sampledValue"`
	Timestamp    string                                `json:"timestamp"`
	CustomData   *TransactionEventRequestCustomData    `json:"customData,omitempty"`
}

type TransactionEventRequestSampledValue struct {
	Value            float64                                    `json:"value"`
	Measurand        *TransactionEventRequestMeasurandEnum      `json:"measurand,omitempty"`
	Context          *TransactionEventRequestReadingContextEnum `json:"context,omitempty"`
	Phase            *TransactionEventRequestPhaseEnum          `json:"phase,omitempty"`
	Location         *TransactionEventRequestLocationEnum       `json:"location,omitempty"`
	SignedMeterValue *TransactionEventRequestSignedMeterValue   `json:"signedMeterValue,omitempty"`
	UnitOfMeasure    *TransactionEventRequestUnitOfMeasure      `json:"unitOfMeasure,omitempty"`
	CustomData       *TransactionEventRequestCustomData         `json:"customData,omitempty"`
}

type TransactionEventRequestUnitOfMeasure struct {
	Unit       *string                            `json:"unit,omitempty"`
	Multiplier *int                               `json:"multiplier,omitempty"`
	CustomData *TransactionEventRequestCustomData `json:"customData,omitempty"`
}

type TransactionEventRequestSignedMeterValue struct {
	SignedMeterData string                             `json:"signedMeterData"`
	SigningMethod   *string                            `json:"signingMethod,omitempty"`
	EncodingMethod  string                             `json:"encodingMethod"`
	PublicKey       *string                            `json:"publicKey,omitempty"`
	CustomData      *TransactionEventRequestCustomData `json:"customData,omitempty"`
}

type TransactionEventRequestLocationEnum string

const (
	TransactionEventRequestLocationEnumBody     TransactionEventRequestLocationEnum = "Body"
	TransactionEventRequestLocationEnumCable    TransactionEventRequestLocationEnum = "Cable"
	TransactionEventRequestLocationEnumEV       TransactionEventRequestLocationEnum = "EV"
	TransactionEventRequestLocationEnumInlet    TransactionEventRequestLocationEnum = "Inlet"
	TransactionEventRequestLocationEnumOutlet   TransactionEventRequestLocationEnum = "Outlet"
	TransactionEventRequestLocationEnumUpstream TransactionEventRequestLocationEnum = "Upstream"
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

type TransactionEventRequestMeasurandEnum string

const (
	TransactionEventRequestMeasurandEnumCurrentExport                             TransactionEventRequestMeasurandEnum = "Current.Export"
	TransactionEventRequestMeasurandEnumCurrentExportOffered                      TransactionEventRequestMeasurandEnum = "Current.Export.Offered"
	TransactionEventRequestMeasurandEnumCurrentExportMinimum                      TransactionEventRequestMeasurandEnum = "Current.Export.Minimum"
	TransactionEventRequestMeasurandEnumCurrentImport                             TransactionEventRequestMeasurandEnum = "Current.Import"
	TransactionEventRequestMeasurandEnumCurrentImportOffered                      TransactionEventRequestMeasurandEnum = "Current.Import.Offered"
	TransactionEventRequestMeasurandEnumCurrentImportMinimum                      TransactionEventRequestMeasurandEnum = "Current.Import.Minimum"
	TransactionEventRequestMeasurandEnumCurrentOffered                            TransactionEventRequestMeasurandEnum = "Current.Offered"
	TransactionEventRequestMeasurandEnumDisplayPresentSOC                         TransactionEventRequestMeasurandEnum = "Display.PresentSOC"
	TransactionEventRequestMeasurandEnumDisplayMinimumSOC                         TransactionEventRequestMeasurandEnum = "Display.MinimumSOC"
	TransactionEventRequestMeasurandEnumDisplayTargetSOC                          TransactionEventRequestMeasurandEnum = "Display.TargetSOC"
	TransactionEventRequestMeasurandEnumDisplayMaximumSOC                         TransactionEventRequestMeasurandEnum = "Display.MaximumSOC"
	TransactionEventRequestMeasurandEnumDisplayRemainingTimeToMinimumSOC          TransactionEventRequestMeasurandEnum = "Display.RemainingTimeToMinimumSOC"
	TransactionEventRequestMeasurandEnumDisplayRemainingTimeToTargetSOC           TransactionEventRequestMeasurandEnum = "Display.RemainingTimeToTargetSOC"
	TransactionEventRequestMeasurandEnumDisplayRemainingTimeToMaximumSOC          TransactionEventRequestMeasurandEnum = "Display.RemainingTimeToMaximumSOC"
	TransactionEventRequestMeasurandEnumDisplayChargingComplete                   TransactionEventRequestMeasurandEnum = "Display.ChargingComplete"
	TransactionEventRequestMeasurandEnumDisplayBatteryEnergyCapacity              TransactionEventRequestMeasurandEnum = "Display.BatteryEnergyCapacity"
	TransactionEventRequestMeasurandEnumDisplayInletHot                           TransactionEventRequestMeasurandEnum = "Display.InletHot"
	TransactionEventRequestMeasurandEnumEnergyActiveExportInterval                TransactionEventRequestMeasurandEnum = "Energy.Active.Export.Interval"
	TransactionEventRequestMeasurandEnumEnergyActiveExportRegister                TransactionEventRequestMeasurandEnum = "Energy.Active.Export.Register"
	TransactionEventRequestMeasurandEnumEnergyActiveImportInterval                TransactionEventRequestMeasurandEnum = "Energy.Active.Import.Interval"
	TransactionEventRequestMeasurandEnumEnergyActiveImportRegister                TransactionEventRequestMeasurandEnum = "Energy.Active.Import.Register"
	TransactionEventRequestMeasurandEnumEnergyActiveImportCableLoss               TransactionEventRequestMeasurandEnum = "Energy.Active.Import.CableLoss"
	TransactionEventRequestMeasurandEnumEnergyActiveImportLocalGenerationRegister TransactionEventRequestMeasurandEnum = "Energy.Active.Import.LocalGeneration.Register"
	TransactionEventRequestMeasurandEnumEnergyActiveNet                           TransactionEventRequestMeasurandEnum = "Energy.Active.Net"
	TransactionEventRequestMeasurandEnumEnergyActiveSetpointInterval              TransactionEventRequestMeasurandEnum = "Energy.Active.Setpoint.Interval"
	TransactionEventRequestMeasurandEnumEnergyApparentExport                      TransactionEventRequestMeasurandEnum = "Energy.Apparent.Export"
	TransactionEventRequestMeasurandEnumEnergyApparentImport                      TransactionEventRequestMeasurandEnum = "Energy.Apparent.Import"
	TransactionEventRequestMeasurandEnumEnergyApparentNet                         TransactionEventRequestMeasurandEnum = "Energy.Apparent.Net"
	TransactionEventRequestMeasurandEnumEnergyReactiveExportInterval              TransactionEventRequestMeasurandEnum = "Energy.Reactive.Export.Interval"
	TransactionEventRequestMeasurandEnumEnergyReactiveExportRegister              TransactionEventRequestMeasurandEnum = "Energy.Reactive.Export.Register"
	TransactionEventRequestMeasurandEnumEnergyReactiveImportInterval              TransactionEventRequestMeasurandEnum = "Energy.Reactive.Import.Interval"
	TransactionEventRequestMeasurandEnumEnergyReactiveImportRegister              TransactionEventRequestMeasurandEnum = "Energy.Reactive.Import.Register"
	TransactionEventRequestMeasurandEnumEnergyReactiveNet                         TransactionEventRequestMeasurandEnum = "Energy.Reactive.Net"
	TransactionEventRequestMeasurandEnumEnergyRequestTarget                       TransactionEventRequestMeasurandEnum = "EnergyRequest.Target"
	TransactionEventRequestMeasurandEnumEnergyRequestMinimum                      TransactionEventRequestMeasurandEnum = "EnergyRequest.Minimum"
	TransactionEventRequestMeasurandEnumEnergyRequestMaximum                      TransactionEventRequestMeasurandEnum = "EnergyRequest.Maximum"
	TransactionEventRequestMeasurandEnumEnergyRequestMinimumV2X                   TransactionEventRequestMeasurandEnum = "EnergyRequest.Minimum.V2X"
	TransactionEventRequestMeasurandEnumEnergyRequestMaximumV2X                   TransactionEventRequestMeasurandEnum = "EnergyRequest.Maximum.V2X"
	TransactionEventRequestMeasurandEnumEnergyRequestBulk                         TransactionEventRequestMeasurandEnum = "EnergyRequest.Bulk"
	TransactionEventRequestMeasurandEnumFrequency                                 TransactionEventRequestMeasurandEnum = "Frequency"
	TransactionEventRequestMeasurandEnumPowerActiveExport                         TransactionEventRequestMeasurandEnum = "Power.Active.Export"
	TransactionEventRequestMeasurandEnumPowerActiveImport                         TransactionEventRequestMeasurandEnum = "Power.Active.Import"
	TransactionEventRequestMeasurandEnumPowerActiveSetpoint                       TransactionEventRequestMeasurandEnum = "Power.Active.Setpoint"
	TransactionEventRequestMeasurandEnumPowerActiveResidual                       TransactionEventRequestMeasurandEnum = "Power.Active.Residual"
	TransactionEventRequestMeasurandEnumPowerExportMinimum                        TransactionEventRequestMeasurandEnum = "Power.Export.Minimum"
	TransactionEventRequestMeasurandEnumPowerExportOffered                        TransactionEventRequestMeasurandEnum = "Power.Export.Offered"
	TransactionEventRequestMeasurandEnumPowerFactor                               TransactionEventRequestMeasurandEnum = "Power.Factor"
	TransactionEventRequestMeasurandEnumPowerImportOffered                        TransactionEventRequestMeasurandEnum = "Power.Import.Offered"
	TransactionEventRequestMeasurandEnumPowerImportMinimum                        TransactionEventRequestMeasurandEnum = "Power.Import.Minimum"
	TransactionEventRequestMeasurandEnumPowerOffered                              TransactionEventRequestMeasurandEnum = "Power.Offered"
	TransactionEventRequestMeasurandEnumPowerReactiveExport                       TransactionEventRequestMeasurandEnum = "Power.Reactive.Export"
	TransactionEventRequestMeasurandEnumPowerReactiveImport                       TransactionEventRequestMeasurandEnum = "Power.Reactive.Import"
	TransactionEventRequestMeasurandEnumSoC                                       TransactionEventRequestMeasurandEnum = "SoC"
	TransactionEventRequestMeasurandEnumVoltage                                   TransactionEventRequestMeasurandEnum = "Voltage"
	TransactionEventRequestMeasurandEnumVoltageMinimum                            TransactionEventRequestMeasurandEnum = "Voltage.Minimum"
	TransactionEventRequestMeasurandEnumVoltageMaximum                            TransactionEventRequestMeasurandEnum = "Voltage.Maximum"
)

type TransactionEventRequestTransactionEventEnum string

const (
	TransactionEventRequestTransactionEventEnumEnded   TransactionEventRequestTransactionEventEnum = "Ended"
	TransactionEventRequestTransactionEventEnumStarted TransactionEventRequestTransactionEventEnum = "Started"
	TransactionEventRequestTransactionEventEnumUpdated TransactionEventRequestTransactionEventEnum = "Updated"
)

type TransactionEventRequestCostDetails struct {
	ChargingPeriods    []TransactionEventRequestChargingPeriod `json:"chargingPeriods,omitempty"`
	TotalCost          TransactionEventRequestTotalCost        `json:"totalCost"`
	TotalUsage         TransactionEventRequestTotalUsage       `json:"totalUsage"`
	FailureToCalculate *bool                                   `json:"failureToCalculate,omitempty"`
	FailureReason      *string                                 `json:"failureReason,omitempty"`
	CustomData         *TransactionEventRequestCustomData      `json:"customData,omitempty"`
}

type TransactionEventRequestTotalUsage struct {
	Energy          float64                            `json:"energy"`
	ChargingTime    int                                `json:"chargingTime"`
	IdleTime        int                                `json:"idleTime"`
	ReservationTime *int                               `json:"reservationTime,omitempty"`
	CustomData      *TransactionEventRequestCustomData `json:"customData,omitempty"`
}

type TransactionEventRequestTotalCost struct {
	Currency         string                                `json:"currency"`
	TypeOfCost       TransactionEventRequestTariffCostEnum `json:"typeOfCost"`
	Fixed            *TransactionEventRequestPrice         `json:"fixed,omitempty"`
	Energy           *TransactionEventRequestPrice         `json:"energy,omitempty"`
	ChargingTime     *TransactionEventRequestPrice         `json:"chargingTime,omitempty"`
	IdleTime         *TransactionEventRequestPrice         `json:"idleTime,omitempty"`
	ReservationTime  *TransactionEventRequestPrice         `json:"reservationTime,omitempty"`
	ReservationFixed *TransactionEventRequestPrice         `json:"reservationFixed,omitempty"`
	Total            TransactionEventRequestTotalPrice     `json:"total"`
	CustomData       *TransactionEventRequestCustomData    `json:"customData,omitempty"`
}

type TransactionEventRequestTotalPrice struct {
	ExclTax    *float64                           `json:"exclTax,omitempty"`
	InclTax    *float64                           `json:"inclTax,omitempty"`
	CustomData *TransactionEventRequestCustomData `json:"customData,omitempty"`
}

type TransactionEventRequestPrice struct {
	ExclTax    *float64                           `json:"exclTax,omitempty"`
	InclTax    *float64                           `json:"inclTax,omitempty"`
	TaxRates   []TransactionEventRequestTaxRate   `json:"taxRates,omitempty"`
	CustomData *TransactionEventRequestCustomData `json:"customData,omitempty"`
}

type TransactionEventRequestTaxRate struct {
	Type       string                             `json:"type"`
	Tax        float64                            `json:"tax"`
	Stack      *int                               `json:"stack,omitempty"`
	CustomData *TransactionEventRequestCustomData `json:"customData,omitempty"`
}

type TransactionEventRequestTariffCostEnum string

const (
	TransactionEventRequestTariffCostEnumNormalCost TransactionEventRequestTariffCostEnum = "NormalCost"
	TransactionEventRequestTariffCostEnumMinCost    TransactionEventRequestTariffCostEnum = "MinCost"
	TransactionEventRequestTariffCostEnumMaxCost    TransactionEventRequestTariffCostEnum = "MaxCost"
)

type TransactionEventRequestChargingPeriod struct {
	Dimensions  []TransactionEventRequestCostDimension `json:"dimensions,omitempty"`
	TariffID    *string                                `json:"tariffId,omitempty"`
	StartPeriod string                                 `json:"startPeriod"`
	CustomData  *TransactionEventRequestCustomData     `json:"customData,omitempty"`
}

type TransactionEventRequestCostDimension struct {
	Type       TransactionEventRequestCostDimensionEnum `json:"type"`
	Volume     float64                                  `json:"volume"`
	CustomData *TransactionEventRequestCustomData       `json:"customData,omitempty"`
}

type TransactionEventRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

type TransactionEventRequestCostDimensionEnum string

const (
	TransactionEventRequestCostDimensionEnumEnergy       TransactionEventRequestCostDimensionEnum = "Energy"
	TransactionEventRequestCostDimensionEnumMaxCurrent   TransactionEventRequestCostDimensionEnum = "MaxCurrent"
	TransactionEventRequestCostDimensionEnumMinCurrent   TransactionEventRequestCostDimensionEnum = "MinCurrent"
	TransactionEventRequestCostDimensionEnumMaxPower     TransactionEventRequestCostDimensionEnum = "MaxPower"
	TransactionEventRequestCostDimensionEnumMinPower     TransactionEventRequestCostDimensionEnum = "MinPower"
	TransactionEventRequestCostDimensionEnumIdleTIme     TransactionEventRequestCostDimensionEnum = "IdleTIme"
	TransactionEventRequestCostDimensionEnumChargingTime TransactionEventRequestCostDimensionEnum = "ChargingTime"
)

func (TransactionEventRequest) ActionName() string { return "TransactionEvent" }

func (TransactionEventRequest) Version() protocol.Version { return protocol.OCPP21 }

func (TransactionEventRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (TransactionEventRequest) SchemaName() string { return "TransactionEventRequest.json" }

func (message TransactionEventRequest) Validate() error {
	return validation.Validate("TransactionEventRequest.json", schemaTransactionEventRequest, message)
}

func (TransactionEventRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("TransactionEventRequest.json", schemaTransactionEventRequest, data)
}
