// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = RemoteStopTransactionConfirmation{}

var schemaRemoteStopTransactionConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected"}}}, Required: []string{"status"}}

type RemoteStopTransactionConfirmation struct {
	Status RemoteStopTransactionConfirmationStatus `json:"status"`
}

type RemoteStopTransactionConfirmationStatus string

const (
	RemoteStopTransactionConfirmationStatusAccepted RemoteStopTransactionConfirmationStatus = "Accepted"
	RemoteStopTransactionConfirmationStatusRejected RemoteStopTransactionConfirmationStatus = "Rejected"
)

func (RemoteStopTransactionConfirmation) ActionName() string { return "RemoteStopTransaction" }

func (RemoteStopTransactionConfirmation) Version() protocol.Version { return protocol.OCPP16 }

func (RemoteStopTransactionConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (RemoteStopTransactionConfirmation) SchemaName() string {
	return "RemoteStopTransactionResponse.json"
}

func (message RemoteStopTransactionConfirmation) Validate() error {
	return validation.Validate("RemoteStopTransactionResponse.json", schemaRemoteStopTransactionConfirmation, message)
}

func (RemoteStopTransactionConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("RemoteStopTransactionResponse.json", schemaRemoteStopTransactionConfirmation, data)
}
