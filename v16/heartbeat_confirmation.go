// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = HeartbeatConfirmation{}

var schemaHeartbeatConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"currentTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}}, Required: []string{"currentTime"}}

type HeartbeatConfirmation struct {
	CurrentTime string `json:"currentTime"`
}

func (HeartbeatConfirmation) ActionName() string { return "Heartbeat" }

func (HeartbeatConfirmation) Version() protocol.Version { return protocol.OCPP16 }

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
