// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = GetPeriodicEventStreamRequest{}

var schemaGetPeriodicEventStreamRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type GetPeriodicEventStreamRequest struct {
	CustomData *GetPeriodicEventStreamRequestCustomData `json:"customData,omitempty"`
}

type GetPeriodicEventStreamRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value GetPeriodicEventStreamRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *GetPeriodicEventStreamRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (GetPeriodicEventStreamRequest) ActionName() string { return "GetPeriodicEventStream" }

func (GetPeriodicEventStreamRequest) Version() protocol.Version { return protocol.OCPP21 }

func (GetPeriodicEventStreamRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (GetPeriodicEventStreamRequest) SchemaName() string { return "GetPeriodicEventStreamRequest.json" }

func (message GetPeriodicEventStreamRequest) Validate() error {
	return validation.Validate("GetPeriodicEventStreamRequest.json", schemaGetPeriodicEventStreamRequest, message)
}

func (GetPeriodicEventStreamRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetPeriodicEventStreamRequest.json", schemaGetPeriodicEventStreamRequest, data)
}
