// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = GetTransactionStatusConfirmation{}

var schemaGetTransactionStatusConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"ongoingIndicator": &validation.Schema{Type: "boolean", AllowAdditional: true}, "messagesInQueue": &validation.Schema{Type: "boolean", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"messagesInQueue"}}

type GetTransactionStatusConfirmation struct {
	OngoingIndicator *bool                                       `json:"ongoingIndicator,omitempty"`
	MessagesInQueue  bool                                        `json:"messagesInQueue"`
	CustomData       *GetTransactionStatusConfirmationCustomData `json:"customData,omitempty"`
}

type GetTransactionStatusConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (GetTransactionStatusConfirmation) ActionName() string { return "GetTransactionStatus" }

func (GetTransactionStatusConfirmation) Version() protocol.Version { return protocol.OCPP21 }

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
