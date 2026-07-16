// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = GetLocalListVersionConfirmation{}

var schemaGetLocalListVersionConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"listVersion": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"listVersion"}}

type GetLocalListVersionConfirmation struct {
	ListVersion int `json:"listVersion"`
}

func (GetLocalListVersionConfirmation) ActionName() string { return "GetLocalListVersion" }

func (GetLocalListVersionConfirmation) Version() protocol.Version { return protocol.OCPP16 }

func (GetLocalListVersionConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (GetLocalListVersionConfirmation) SchemaName() string { return "GetLocalListVersionResponse.json" }

func (message GetLocalListVersionConfirmation) Validate() error {
	return validation.Validate("GetLocalListVersionResponse.json", schemaGetLocalListVersionConfirmation, message)
}

func (GetLocalListVersionConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetLocalListVersionResponse.json", schemaGetLocalListVersionConfirmation, data)
}
