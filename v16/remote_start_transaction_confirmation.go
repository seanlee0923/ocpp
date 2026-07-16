// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = RemoteStartTransactionConfirmation{}

var schemaRemoteStartTransactionConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected"}}}, Required: []string{"status"}}

type RemoteStartTransactionConfirmation struct {
	Status RemoteStartTransactionConfirmationStatus `json:"status"`
}

type RemoteStartTransactionConfirmationStatus string

const (
	RemoteStartTransactionConfirmationStatusAccepted RemoteStartTransactionConfirmationStatus = "Accepted"
	RemoteStartTransactionConfirmationStatusRejected RemoteStartTransactionConfirmationStatus = "Rejected"
)

func (RemoteStartTransactionConfirmation) ActionName() string { return "RemoteStartTransaction" }

func (RemoteStartTransactionConfirmation) Version() protocol.Version { return protocol.OCPP16 }

func (RemoteStartTransactionConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (RemoteStartTransactionConfirmation) SchemaName() string {
	return "RemoteStartTransactionResponse.json"
}

func (message RemoteStartTransactionConfirmation) Validate() error {
	return validation.Validate("RemoteStartTransactionResponse.json", schemaRemoteStartTransactionConfirmation, message)
}

func (RemoteStartTransactionConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("RemoteStartTransactionResponse.json", schemaRemoteStartTransactionConfirmation, data)
}
