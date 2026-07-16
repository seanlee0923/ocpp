// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = StartTransactionConfirmation{}

var schemaStartTransactionConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"idTagInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"expiryDate": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "parentIdTag": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Blocked", "Expired", "Invalid", "ConcurrentTx"}}}, Required: []string{"status"}}, "transactionId": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"idTagInfo", "transactionId"}}

type StartTransactionConfirmation struct {
	IDTagInfo     StartTransactionConfirmationIDTagInfo `json:"idTagInfo"`
	TransactionID int                                   `json:"transactionId"`
}

type StartTransactionConfirmationIDTagInfo struct {
	ExpiryDate  *string                                     `json:"expiryDate,omitempty"`
	ParentIDTag *string                                     `json:"parentIdTag,omitempty"`
	Status      StartTransactionConfirmationIDTagInfoStatus `json:"status"`
}

type StartTransactionConfirmationIDTagInfoStatus string

const (
	StartTransactionConfirmationIDTagInfoStatusAccepted     StartTransactionConfirmationIDTagInfoStatus = "Accepted"
	StartTransactionConfirmationIDTagInfoStatusBlocked      StartTransactionConfirmationIDTagInfoStatus = "Blocked"
	StartTransactionConfirmationIDTagInfoStatusExpired      StartTransactionConfirmationIDTagInfoStatus = "Expired"
	StartTransactionConfirmationIDTagInfoStatusInvalid      StartTransactionConfirmationIDTagInfoStatus = "Invalid"
	StartTransactionConfirmationIDTagInfoStatusConcurrentTx StartTransactionConfirmationIDTagInfoStatus = "ConcurrentTx"
)

func (StartTransactionConfirmation) ActionName() string { return "StartTransaction" }

func (StartTransactionConfirmation) Version() protocol.Version { return protocol.OCPP16 }

func (StartTransactionConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (StartTransactionConfirmation) SchemaName() string { return "StartTransactionResponse.json" }

func (message StartTransactionConfirmation) Validate() error {
	return validation.Validate("StartTransactionResponse.json", schemaStartTransactionConfirmation, message)
}

func (StartTransactionConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("StartTransactionResponse.json", schemaStartTransactionConfirmation, data)
}
