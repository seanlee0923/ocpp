// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = SetVariableMonitoringRequest{}

var schemaSetVariableMonitoringRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "setMonitoringData": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "id": &validation.Schema{Type: "integer", AllowAdditional: true}, "transaction": &validation.Schema{Type: "boolean", AllowAdditional: true}, "value": &validation.Schema{Type: "number", AllowAdditional: true}, "type": &validation.Schema{Type: "string", Enum: []string{"UpperThreshold", "LowerThreshold", "Delta", "Periodic", "PeriodicClockAligned"}}, "severity": &validation.Schema{Type: "integer", AllowAdditional: true}, "component": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "evse": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "id": &validation.Schema{Type: "integer", AllowAdditional: true}, "connectorId": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"id"}}, "name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}}, Required: []string{"name"}}, "variable": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}}, Required: []string{"name"}}}, Required: []string{"value", "type", "severity", "component", "variable"}}, MinItems: 1, HasMinItems: true}}, Required: []string{"setMonitoringData"}}

type SetVariableMonitoringRequest struct {
	CustomData        *SetVariableMonitoringRequestCustomData         `json:"customData,omitempty"`
	SetMonitoringData []SetVariableMonitoringRequestSetMonitoringData `json:"setMonitoringData"`
}

type SetVariableMonitoringRequestSetMonitoringData struct {
	CustomData  *SetVariableMonitoringRequestCustomData `json:"customData,omitempty"`
	ID          *int                                    `json:"id,omitempty"`
	Transaction *bool                                   `json:"transaction,omitempty"`
	Value       float64                                 `json:"value"`
	Type        SetVariableMonitoringRequestMonitorEnum `json:"type"`
	Severity    int                                     `json:"severity"`
	Component   SetVariableMonitoringRequestComponent   `json:"component"`
	Variable    SetVariableMonitoringRequestVariable    `json:"variable"`
}

type SetVariableMonitoringRequestVariable struct {
	CustomData *SetVariableMonitoringRequestCustomData `json:"customData,omitempty"`
	Name       string                                  `json:"name"`
	Instance   *string                                 `json:"instance,omitempty"`
}

type SetVariableMonitoringRequestComponent struct {
	CustomData *SetVariableMonitoringRequestCustomData `json:"customData,omitempty"`
	EVSE       *SetVariableMonitoringRequestEVSE       `json:"evse,omitempty"`
	Name       string                                  `json:"name"`
	Instance   *string                                 `json:"instance,omitempty"`
}

type SetVariableMonitoringRequestEVSE struct {
	CustomData  *SetVariableMonitoringRequestCustomData `json:"customData,omitempty"`
	ID          int                                     `json:"id"`
	ConnectorID *int                                    `json:"connectorId,omitempty"`
}

type SetVariableMonitoringRequestMonitorEnum string

const (
	SetVariableMonitoringRequestMonitorEnumUpperThreshold       SetVariableMonitoringRequestMonitorEnum = "UpperThreshold"
	SetVariableMonitoringRequestMonitorEnumLowerThreshold       SetVariableMonitoringRequestMonitorEnum = "LowerThreshold"
	SetVariableMonitoringRequestMonitorEnumDelta                SetVariableMonitoringRequestMonitorEnum = "Delta"
	SetVariableMonitoringRequestMonitorEnumPeriodic             SetVariableMonitoringRequestMonitorEnum = "Periodic"
	SetVariableMonitoringRequestMonitorEnumPeriodicClockAligned SetVariableMonitoringRequestMonitorEnum = "PeriodicClockAligned"
)

type SetVariableMonitoringRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (SetVariableMonitoringRequest) ActionName() string { return "SetVariableMonitoring" }

func (SetVariableMonitoringRequest) Version() protocol.Version { return protocol.OCPP201 }

func (SetVariableMonitoringRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (SetVariableMonitoringRequest) SchemaName() string { return "SetVariableMonitoringRequest.json" }

func (message SetVariableMonitoringRequest) Validate() error {
	return validation.Validate("SetVariableMonitoringRequest.json", schemaSetVariableMonitoringRequest, message)
}

func (SetVariableMonitoringRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("SetVariableMonitoringRequest.json", schemaSetVariableMonitoringRequest, data)
}
