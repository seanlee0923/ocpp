// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = GetTransactionStatusRequest{}

var schemaGetTransactionStatusRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"transactionId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type GetTransactionStatusRequest struct {
	TransactionID *string                                `json:"transactionId,omitempty"`
	CustomData    *GetTransactionStatusRequestCustomData `json:"customData,omitempty"`
}

type GetTransactionStatusRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (GetTransactionStatusRequest) ActionName() string { return "GetTransactionStatus" }

func (GetTransactionStatusRequest) Version() protocol.Version { return protocol.OCPP21 }

func (GetTransactionStatusRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (GetTransactionStatusRequest) SchemaName() string { return "GetTransactionStatusRequest.json" }

func (message GetTransactionStatusRequest) Validate() error {
	return validation.Validate("GetTransactionStatusRequest.json", schemaGetTransactionStatusRequest, message)
}

func (GetTransactionStatusRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetTransactionStatusRequest.json", schemaGetTransactionStatusRequest, data)
}
