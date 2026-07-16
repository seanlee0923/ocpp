// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = HeartbeatRequest{}

var schemaHeartbeatRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type HeartbeatRequest struct {
	CustomData *HeartbeatRequestCustomData `json:"customData,omitempty"`
}

type HeartbeatRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value HeartbeatRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *HeartbeatRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (HeartbeatRequest) ActionName() string { return "Heartbeat" }

func (HeartbeatRequest) Version() protocol.Version { return protocol.OCPP21 }

func (HeartbeatRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (HeartbeatRequest) SchemaName() string { return "HeartbeatRequest.json" }

func (message HeartbeatRequest) Validate() error {
	return validation.Validate("HeartbeatRequest.json", schemaHeartbeatRequest, message)
}

func (HeartbeatRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("HeartbeatRequest.json", schemaHeartbeatRequest, data)
}
