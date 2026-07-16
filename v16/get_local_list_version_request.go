// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = GetLocalListVersionRequest{}

var schemaGetLocalListVersionRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{}}

type GetLocalListVersionRequest struct {
}

func (GetLocalListVersionRequest) ActionName() string { return "GetLocalListVersion" }

func (GetLocalListVersionRequest) Version() protocol.Version { return protocol.OCPP16 }

func (GetLocalListVersionRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (GetLocalListVersionRequest) SchemaName() string { return "GetLocalListVersion.json" }

func (message GetLocalListVersionRequest) Validate() error {
	return validation.Validate("GetLocalListVersion.json", schemaGetLocalListVersionRequest, message)
}

func (GetLocalListVersionRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetLocalListVersion.json", schemaGetLocalListVersionRequest, data)
}
