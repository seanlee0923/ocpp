// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ResetRequest{}

var schemaResetRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"type": &validation.Schema{Type: "string", Enum: []string{"Hard", "Soft"}}}, Required: []string{"type"}}

type ResetRequest struct {
	Type ResetRequestType `json:"type"`
}

type ResetRequestType string

const (
	ResetRequestTypeHard ResetRequestType = "Hard"
	ResetRequestTypeSoft ResetRequestType = "Soft"
)

func (ResetRequest) ActionName() string { return "Reset" }

func (ResetRequest) Version() protocol.Version { return protocol.OCPP16 }

func (ResetRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (ResetRequest) SchemaName() string { return "Reset.json" }

func (message ResetRequest) Validate() error {
	return validation.Validate("Reset.json", schemaResetRequest, message)
}

func (ResetRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("Reset.json", schemaResetRequest, data)
}
