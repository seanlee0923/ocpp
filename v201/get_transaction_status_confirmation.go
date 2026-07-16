// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = GetTransactionStatusConfirmation{}

var schemaGetTransactionStatusConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "ongoingIndicator": &validation.Schema{Type: "boolean", AllowAdditional: true}, "messagesInQueue": &validation.Schema{Type: "boolean", AllowAdditional: true}}, Required: []string{"messagesInQueue"}}

type GetTransactionStatusConfirmation struct {
	CustomData       *GetTransactionStatusConfirmationCustomData `json:"customData,omitempty"`
	OngoingIndicator *bool                                       `json:"ongoingIndicator,omitempty"`
	MessagesInQueue  bool                                        `json:"messagesInQueue"`
}

type GetTransactionStatusConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (GetTransactionStatusConfirmation) ActionName() string { return "GetTransactionStatus" }

func (GetTransactionStatusConfirmation) Version() protocol.Version { return protocol.OCPP201 }

func (GetTransactionStatusConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (GetTransactionStatusConfirmation) SchemaName() string {
	return "GetTransactionStatusResponse.json"
}

func (message GetTransactionStatusConfirmation) Validate() error {
	return validation.Validate("GetTransactionStatusResponse.json", schemaGetTransactionStatusConfirmation, message)
}

func (GetTransactionStatusConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetTransactionStatusResponse.json", schemaGetTransactionStatusConfirmation, data)
}
