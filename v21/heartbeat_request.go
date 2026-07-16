// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = HeartbeatRequest{}

var schemaHeartbeatRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type HeartbeatRequest struct {
	CustomData *HeartbeatRequestCustomData `json:"customData,omitempty"`
}

type HeartbeatRequestCustomData struct {
	VendorID string `json:"vendorId"`
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
