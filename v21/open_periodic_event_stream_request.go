// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = OpenPeriodicEventStreamRequest{}

var schemaOpenPeriodicEventStreamRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"constantStreamData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"id": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "params": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"interval": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "values": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "variableMonitoringId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id", "variableMonitoringId", "params"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"constantStreamData"}}

type OpenPeriodicEventStreamRequest struct {
	ConstantStreamData OpenPeriodicEventStreamRequestConstantStreamData `json:"constantStreamData"`
	CustomData         *OpenPeriodicEventStreamRequestCustomData        `json:"customData,omitempty"`
}

type OpenPeriodicEventStreamRequestConstantStreamData struct {
	ID                   int                                                     `json:"id"`
	Params               OpenPeriodicEventStreamRequestPeriodicEventStreamParams `json:"params"`
	VariableMonitoringID int                                                     `json:"variableMonitoringId"`
	CustomData           *OpenPeriodicEventStreamRequestCustomData               `json:"customData,omitempty"`
}

type OpenPeriodicEventStreamRequestPeriodicEventStreamParams struct {
	Interval   *int                                      `json:"interval,omitempty"`
	Values     *int                                      `json:"values,omitempty"`
	CustomData *OpenPeriodicEventStreamRequestCustomData `json:"customData,omitempty"`
}

type OpenPeriodicEventStreamRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value OpenPeriodicEventStreamRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *OpenPeriodicEventStreamRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (OpenPeriodicEventStreamRequest) ActionName() string { return "OpenPeriodicEventStream" }

func (OpenPeriodicEventStreamRequest) Version() protocol.Version { return protocol.OCPP21 }

func (OpenPeriodicEventStreamRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (OpenPeriodicEventStreamRequest) SchemaName() string {
	return "OpenPeriodicEventStreamRequest.json"
}

func (message OpenPeriodicEventStreamRequest) Validate() error {
	return validation.Validate("OpenPeriodicEventStreamRequest.json", schemaOpenPeriodicEventStreamRequest, message)
}

func (OpenPeriodicEventStreamRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("OpenPeriodicEventStreamRequest.json", schemaOpenPeriodicEventStreamRequest, data)
}
