// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = HeartbeatConfirmation{}

var schemaHeartbeatConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "currentTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}}, Required: []string{"currentTime"}}

type HeartbeatConfirmation struct {
	CustomData  *HeartbeatConfirmationCustomData `json:"customData,omitempty"`
	CurrentTime string                           `json:"currentTime"`
}

type HeartbeatConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value HeartbeatConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *HeartbeatConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (HeartbeatConfirmation) ActionName() string { return "Heartbeat" }

func (HeartbeatConfirmation) Version() protocol.Version { return protocol.OCPP201 }

func (HeartbeatConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (HeartbeatConfirmation) SchemaName() string { return "HeartbeatResponse.json" }

func (message HeartbeatConfirmation) Validate() error {
	return validation.Validate("HeartbeatResponse.json", schemaHeartbeatConfirmation, message)
}

func (HeartbeatConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("HeartbeatResponse.json", schemaHeartbeatConfirmation, data)
}
