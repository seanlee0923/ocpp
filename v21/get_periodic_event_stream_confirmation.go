// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = GetPeriodicEventStreamConfirmation{}

var schemaGetPeriodicEventStreamConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"constantStreamData": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"id": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "params": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"interval": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "values": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "variableMonitoringId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id", "variableMonitoringId", "params"}}, MinItems: 1, HasMinItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type GetPeriodicEventStreamConfirmation struct {
	ConstantStreamData []GetPeriodicEventStreamConfirmationConstantStreamData `json:"constantStreamData,omitempty"`
	CustomData         *GetPeriodicEventStreamConfirmationCustomData          `json:"customData,omitempty"`
}

type GetPeriodicEventStreamConfirmationConstantStreamData struct {
	ID                   int                                                         `json:"id"`
	Params               GetPeriodicEventStreamConfirmationPeriodicEventStreamParams `json:"params"`
	VariableMonitoringID int                                                         `json:"variableMonitoringId"`
	CustomData           *GetPeriodicEventStreamConfirmationCustomData               `json:"customData,omitempty"`
}

type GetPeriodicEventStreamConfirmationPeriodicEventStreamParams struct {
	Interval   *int                                          `json:"interval,omitempty"`
	Values     *int                                          `json:"values,omitempty"`
	CustomData *GetPeriodicEventStreamConfirmationCustomData `json:"customData,omitempty"`
}

type GetPeriodicEventStreamConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value GetPeriodicEventStreamConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *GetPeriodicEventStreamConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (GetPeriodicEventStreamConfirmation) ActionName() string { return "GetPeriodicEventStream" }

func (GetPeriodicEventStreamConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (GetPeriodicEventStreamConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (GetPeriodicEventStreamConfirmation) SchemaName() string {
	return "GetPeriodicEventStreamResponse.json"
}

func (message GetPeriodicEventStreamConfirmation) Validate() error {
	return validation.Validate("GetPeriodicEventStreamResponse.json", schemaGetPeriodicEventStreamConfirmation, message)
}

func (GetPeriodicEventStreamConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetPeriodicEventStreamResponse.json", schemaGetPeriodicEventStreamConfirmation, data)
}
