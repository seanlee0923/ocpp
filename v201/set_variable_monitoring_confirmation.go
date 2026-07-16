// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = SetVariableMonitoringConfirmation{}

var schemaSetVariableMonitoringConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "setMonitoringResult": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "id": &validation.Schema{Type: "integer", AllowAdditional: true}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "UnknownComponent", "UnknownVariable", "UnsupportedMonitorType", "Rejected", "Duplicate"}}, "type": &validation.Schema{Type: "string", Enum: []string{"UpperThreshold", "LowerThreshold", "Delta", "Periodic", "PeriodicClockAligned"}}, "component": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "evse": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "id": &validation.Schema{Type: "integer", AllowAdditional: true}, "connectorId": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"id"}}, "name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}}, Required: []string{"name"}}, "variable": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}}, Required: []string{"name"}}, "severity": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"status", "type", "severity", "component", "variable"}}, MinItems: 1, HasMinItems: true}}, Required: []string{"setMonitoringResult"}}

type SetVariableMonitoringConfirmation struct {
	CustomData          *SetVariableMonitoringConfirmationCustomData           `json:"customData,omitempty"`
	SetMonitoringResult []SetVariableMonitoringConfirmationSetMonitoringResult `json:"setMonitoringResult"`
}

type SetVariableMonitoringConfirmationSetMonitoringResult struct {
	CustomData *SetVariableMonitoringConfirmationCustomData             `json:"customData,omitempty"`
	ID         *int                                                     `json:"id,omitempty"`
	StatusInfo *SetVariableMonitoringConfirmationStatusInfo             `json:"statusInfo,omitempty"`
	Status     SetVariableMonitoringConfirmationSetMonitoringStatusEnum `json:"status"`
	Type       SetVariableMonitoringConfirmationMonitorEnum             `json:"type"`
	Component  SetVariableMonitoringConfirmationComponent               `json:"component"`
	Variable   SetVariableMonitoringConfirmationVariable                `json:"variable"`
	Severity   int                                                      `json:"severity"`
}

type SetVariableMonitoringConfirmationVariable struct {
	CustomData *SetVariableMonitoringConfirmationCustomData `json:"customData,omitempty"`
	Name       string                                       `json:"name"`
	Instance   *string                                      `json:"instance,omitempty"`
}

type SetVariableMonitoringConfirmationComponent struct {
	CustomData *SetVariableMonitoringConfirmationCustomData `json:"customData,omitempty"`
	EVSE       *SetVariableMonitoringConfirmationEVSE       `json:"evse,omitempty"`
	Name       string                                       `json:"name"`
	Instance   *string                                      `json:"instance,omitempty"`
}

type SetVariableMonitoringConfirmationEVSE struct {
	CustomData  *SetVariableMonitoringConfirmationCustomData `json:"customData,omitempty"`
	ID          int                                          `json:"id"`
	ConnectorID *int                                         `json:"connectorId,omitempty"`
}

type SetVariableMonitoringConfirmationMonitorEnum string

const (
	SetVariableMonitoringConfirmationMonitorEnumUpperThreshold       SetVariableMonitoringConfirmationMonitorEnum = "UpperThreshold"
	SetVariableMonitoringConfirmationMonitorEnumLowerThreshold       SetVariableMonitoringConfirmationMonitorEnum = "LowerThreshold"
	SetVariableMonitoringConfirmationMonitorEnumDelta                SetVariableMonitoringConfirmationMonitorEnum = "Delta"
	SetVariableMonitoringConfirmationMonitorEnumPeriodic             SetVariableMonitoringConfirmationMonitorEnum = "Periodic"
	SetVariableMonitoringConfirmationMonitorEnumPeriodicClockAligned SetVariableMonitoringConfirmationMonitorEnum = "PeriodicClockAligned"
)

type SetVariableMonitoringConfirmationSetMonitoringStatusEnum string

const (
	SetVariableMonitoringConfirmationSetMonitoringStatusEnumAccepted               SetVariableMonitoringConfirmationSetMonitoringStatusEnum = "Accepted"
	SetVariableMonitoringConfirmationSetMonitoringStatusEnumUnknownComponent       SetVariableMonitoringConfirmationSetMonitoringStatusEnum = "UnknownComponent"
	SetVariableMonitoringConfirmationSetMonitoringStatusEnumUnknownVariable        SetVariableMonitoringConfirmationSetMonitoringStatusEnum = "UnknownVariable"
	SetVariableMonitoringConfirmationSetMonitoringStatusEnumUnsupportedMonitorType SetVariableMonitoringConfirmationSetMonitoringStatusEnum = "UnsupportedMonitorType"
	SetVariableMonitoringConfirmationSetMonitoringStatusEnumRejected               SetVariableMonitoringConfirmationSetMonitoringStatusEnum = "Rejected"
	SetVariableMonitoringConfirmationSetMonitoringStatusEnumDuplicate              SetVariableMonitoringConfirmationSetMonitoringStatusEnum = "Duplicate"
)

type SetVariableMonitoringConfirmationStatusInfo struct {
	CustomData     *SetVariableMonitoringConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                                       `json:"reasonCode"`
	AdditionalInfo *string                                      `json:"additionalInfo,omitempty"`
}

type SetVariableMonitoringConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (SetVariableMonitoringConfirmation) ActionName() string { return "SetVariableMonitoring" }

func (SetVariableMonitoringConfirmation) Version() protocol.Version { return protocol.OCPP201 }

func (SetVariableMonitoringConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (SetVariableMonitoringConfirmation) SchemaName() string {
	return "SetVariableMonitoringResponse.json"
}

func (message SetVariableMonitoringConfirmation) Validate() error {
	return validation.Validate("SetVariableMonitoringResponse.json", schemaSetVariableMonitoringConfirmation, message)
}

func (SetVariableMonitoringConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("SetVariableMonitoringResponse.json", schemaSetVariableMonitoringConfirmation, data)
}
