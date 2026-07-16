// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = RequestStopTransactionRequest{}

var schemaRequestStopTransactionRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"transactionId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"transactionId"}}

type RequestStopTransactionRequest struct {
	TransactionID string                                   `json:"transactionId"`
	CustomData    *RequestStopTransactionRequestCustomData `json:"customData,omitempty"`
}

type RequestStopTransactionRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (RequestStopTransactionRequest) ActionName() string { return "RequestStopTransaction" }

func (RequestStopTransactionRequest) Version() protocol.Version { return protocol.OCPP21 }

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
