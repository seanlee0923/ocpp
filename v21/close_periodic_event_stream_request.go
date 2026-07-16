// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ClosePeriodicEventStreamRequest{}

var schemaClosePeriodicEventStreamRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"id": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id"}}

type ClosePeriodicEventStreamRequest struct {
	ID         int                                        `json:"id"`
	CustomData *ClosePeriodicEventStreamRequestCustomData `json:"customData,omitempty"`
}

type ClosePeriodicEventStreamRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value ClosePeriodicEventStreamRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *ClosePeriodicEventStreamRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (ClosePeriodicEventStreamRequest) ActionName() string { return "ClosePeriodicEventStream" }

func (ClosePeriodicEventStreamRequest) Version() protocol.Version { return protocol.OCPP21 }

func (ClosePeriodicEventStreamRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (ClosePeriodicEventStreamRequest) SchemaName() string {
	return "ClosePeriodicEventStreamRequest.json"
}

func (message ClosePeriodicEventStreamRequest) Validate() error {
	return validation.Validate("ClosePeriodicEventStreamRequest.json", schemaClosePeriodicEventStreamRequest, message)
}

func (ClosePeriodicEventStreamRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ClosePeriodicEventStreamRequest.json", schemaClosePeriodicEventStreamRequest, data)
}
