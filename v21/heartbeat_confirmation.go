// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = HeartbeatConfirmation{}

var schemaHeartbeatConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"currentTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"currentTime"}}

type HeartbeatConfirmation struct {
	CurrentTime string                           `json:"currentTime"`
	CustomData  *HeartbeatConfirmationCustomData `json:"customData,omitempty"`
}

type HeartbeatConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (HeartbeatConfirmation) ActionName() string { return "Heartbeat" }

func (HeartbeatConfirmation) Version() protocol.Version { return protocol.OCPP21 }

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
