// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = RemoteStopTransactionRequest{}

var schemaRemoteStopTransactionRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"transactionId": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"transactionId"}}

type RemoteStopTransactionRequest struct {
	TransactionID int `json:"transactionId"`
}

func (RemoteStopTransactionRequest) ActionName() string { return "RemoteStopTransaction" }

func (RemoteStopTransactionRequest) Version() protocol.Version { return protocol.OCPP16 }

func (RemoteStopTransactionRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (RemoteStopTransactionRequest) SchemaName() string { return "RemoteStopTransaction.json" }

func (message RemoteStopTransactionRequest) Validate() error {
	return validation.Validate("RemoteStopTransaction.json", schemaRemoteStopTransactionRequest, message)
}

func (RemoteStopTransactionRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("RemoteStopTransaction.json", schemaRemoteStopTransactionRequest, data)
}
