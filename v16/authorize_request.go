// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = AuthorizeRequest{}

var schemaAuthorizeRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"idTag": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}}, Required: []string{"idTag"}}

type AuthorizeRequest struct {
	IDTag string `json:"idTag"`
}

func (AuthorizeRequest) ActionName() string { return "Authorize" }

func (AuthorizeRequest) Version() protocol.Version { return protocol.OCPP16 }

func (AuthorizeRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (AuthorizeRequest) SchemaName() string { return "Authorize.json" }

func (message AuthorizeRequest) Validate() error {
	return validation.Validate("Authorize.json", schemaAuthorizeRequest, message)
}

func (AuthorizeRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("Authorize.json", schemaAuthorizeRequest, data)
}
