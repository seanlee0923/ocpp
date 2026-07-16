// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = RequestStopTransactionRequest{}

var schemaRequestStopTransactionRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "transactionId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}}, Required: []string{"transactionId"}}

type RequestStopTransactionRequest struct {
	CustomData    *RequestStopTransactionRequestCustomData `json:"customData,omitempty"`
	TransactionID string                                   `json:"transactionId"`
}

type RequestStopTransactionRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (RequestStopTransactionRequest) ActionName() string { return "RequestStopTransaction" }

func (RequestStopTransactionRequest) Version() protocol.Version { return protocol.OCPP201 }

func (RequestStopTransactionRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (RequestStopTransactionRequest) SchemaName() string { return "RequestStopTransactionRequest.json" }

func (message RequestStopTransactionRequest) Validate() error {
	return validation.Validate("RequestStopTransactionRequest.json", schemaRequestStopTransactionRequest, message)
}

func (RequestStopTransactionRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("RequestStopTransactionRequest.json", schemaRequestStopTransactionRequest, data)
}
