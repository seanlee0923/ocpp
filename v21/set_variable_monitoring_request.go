// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = SetVariableMonitoringRequest{}

var schemaSetVariableMonitoringRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"setMonitoringData": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"id": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "periodicEventStream": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"interval": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "values": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "transaction": &validation.Schema{Type: "boolean", AllowAdditional: true}, "value": &validation.Schema{Type: "number", AllowAdditional: true}, "type": &validation.Schema{Type: "string", Enum: []string{"UpperThreshold", "LowerThreshold", "Delta", "Periodic", "PeriodicClockAligned", "TargetDelta", "TargetDeltaRelative"}}, "severity": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "component": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"evse": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"id": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "connectorId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id"}}, "name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"name"}}, "variable": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"name"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"value", "type", "severity", "component", "variable"}}, MinItems: 1, HasMinItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"setMonitoringData"}}

type SetVariableMonitoringRequest struct {
	SetMonitoringData []SetVariableMonitoringRequestSetMonitoringData `json:"setMonitoringData"`
	CustomData        *SetVariableMonitoringRequestCustomData         `json:"customData,omitempty"`
}

type SetVariableMonitoringRequestSetMonitoringData struct {
	ID                  *int                                                   `json:"id,omitempty"`
	PeriodicEventStream *SetVariableMonitoringRequestPeriodicEventStreamParams `json:"periodicEventStream,omitempty"`
	Transaction         *bool                                                  `json:"transaction,omitempty"`
	Value               float64                                                `json:"value"`
	Type                SetVariableMonitoringRequestMonitorEnum                `json:"type"`
	Severity            int                                                    `json:"severity"`
	Component           SetVariableMonitoringRequestComponent                  `json:"component"`
	Variable            SetVariableMonitoringRequestVariable                   `json:"variable"`
	CustomData          *SetVariableMonitoringRequestCustomData                `json:"customData,omitempty"`
}

type SetVariableMonitoringRequestVariable struct {
	Name       string                                  `json:"name"`
	Instance   *string                                 `json:"instance,omitempty"`
	CustomData *SetVariableMonitoringRequestCustomData `json:"customData,omitempty"`
}

type SetVariableMonitoringRequestComponent struct {
	EVSE       *SetVariableMonitoringRequestEVSE       `json:"evse,omitempty"`
	Name       string                                  `json:"name"`
	Instance   *string                                 `json:"instance,omitempty"`
	CustomData *SetVariableMonitoringRequestCustomData `json:"customData,omitempty"`
}

type SetVariableMonitoringRequestEVSE struct {
	ID          int                                     `json:"id"`
	ConnectorID *int                                    `json:"connectorId,omitempty"`
	CustomData  *SetVariableMonitoringRequestCustomData `json:"customData,omitempty"`
}

type SetVariableMonitoringRequestMonitorEnum string

const (
	SetVariableMonitoringRequestMonitorEnumUpperThreshold       SetVariableMonitoringRequestMonitorEnum = "UpperThreshold"
	SetVariableMonitoringRequestMonitorEnumLowerThreshold       SetVariableMonitoringRequestMonitorEnum = "LowerThreshold"
	SetVariableMonitoringRequestMonitorEnumDelta                SetVariableMonitoringRequestMonitorEnum = "Delta"
	SetVariableMonitoringRequestMonitorEnumPeriodic             SetVariableMonitoringRequestMonitorEnum = "Periodic"
	SetVariableMonitoringRequestMonitorEnumPeriodicClockAligned SetVariableMonitoringRequestMonitorEnum = "PeriodicClockAligned"
	SetVariableMonitoringRequestMonitorEnumTargetDelta          SetVariableMonitoringRequestMonitorEnum = "TargetDelta"
	SetVariableMonitoringRequestMonitorEnumTargetDeltaRelative  SetVariableMonitoringRequestMonitorEnum = "TargetDeltaRelative"
)

type SetVariableMonitoringRequestPeriodicEventStreamParams struct {
	Interval   *int                                    `json:"interval,omitempty"`
	Values     *int                                    `json:"values,omitempty"`
	CustomData *SetVariableMonitoringRequestCustomData `json:"customData,omitempty"`
}

type SetVariableMonitoringRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value SetVariableMonitoringRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *SetVariableMonitoringRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (SetVariableMonitoringRequest) ActionName() string { return "SetVariableMonitoring" }

func (SetVariableMonitoringRequest) Version() protocol.Version { return protocol.OCPP21 }

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
