// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = AuthorizeConfirmation{}

var schemaAuthorizeConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"idTagInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"expiryDate": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "parentIdTag": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Blocked", "Expired", "Invalid", "ConcurrentTx"}}}, Required: []string{"status"}}}, Required: []string{"idTagInfo"}}

type AuthorizeConfirmation struct {
	IDTagInfo AuthorizeConfirmationIDTagInfo `json:"idTagInfo"`
}

type AuthorizeConfirmationIDTagInfo struct {
	ExpiryDate  *string                              `json:"expiryDate,omitempty"`
	ParentIDTag *string                              `json:"parentIdTag,omitempty"`
	Status      AuthorizeConfirmationIDTagInfoStatus `json:"status"`
}

type AuthorizeConfirmationIDTagInfoStatus string

const (
	AuthorizeConfirmationIDTagInfoStatusAccepted     AuthorizeConfirmationIDTagInfoStatus = "Accepted"
	AuthorizeConfirmationIDTagInfoStatusBlocked      AuthorizeConfirmationIDTagInfoStatus = "Blocked"
	AuthorizeConfirmationIDTagInfoStatusExpired      AuthorizeConfirmationIDTagInfoStatus = "Expired"
	AuthorizeConfirmationIDTagInfoStatusInvalid      AuthorizeConfirmationIDTagInfoStatus = "Invalid"
	AuthorizeConfirmationIDTagInfoStatusConcurrentTx AuthorizeConfirmationIDTagInfoStatus = "ConcurrentTx"
)

func (AuthorizeConfirmation) ActionName() string { return "Authorize" }

func (AuthorizeConfirmation) Version() protocol.Version { return protocol.OCPP16 }

func (AuthorizeConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (AuthorizeConfirmation) SchemaName() string { return "AuthorizeResponse.json" }

func (message AuthorizeConfirmation) Validate() error {
	return validation.Validate("AuthorizeResponse.json", schemaAuthorizeConfirmation, message)
}

func (AuthorizeConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("AuthorizeResponse.json", schemaAuthorizeConfirmation, data)
}
