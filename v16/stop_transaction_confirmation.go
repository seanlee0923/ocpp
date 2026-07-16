// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = StopTransactionConfirmation{}

var schemaStopTransactionConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"idTagInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"expiryDate": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "parentIdTag": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Blocked", "Expired", "Invalid", "ConcurrentTx"}}}, Required: []string{"status"}}}}

type StopTransactionConfirmation struct {
	IDTagInfo *StopTransactionConfirmationIDTagInfo `json:"idTagInfo,omitempty"`
}

type StopTransactionConfirmationIDTagInfo struct {
	ExpiryDate  *string                                    `json:"expiryDate,omitempty"`
	ParentIDTag *string                                    `json:"parentIdTag,omitempty"`
	Status      StopTransactionConfirmationIDTagInfoStatus `json:"status"`
}

type StopTransactionConfirmationIDTagInfoStatus string

const (
	StopTransactionConfirmationIDTagInfoStatusAccepted     StopTransactionConfirmationIDTagInfoStatus = "Accepted"
	StopTransactionConfirmationIDTagInfoStatusBlocked      StopTransactionConfirmationIDTagInfoStatus = "Blocked"
	StopTransactionConfirmationIDTagInfoStatusExpired      StopTransactionConfirmationIDTagInfoStatus = "Expired"
	StopTransactionConfirmationIDTagInfoStatusInvalid      StopTransactionConfirmationIDTagInfoStatus = "Invalid"
	StopTransactionConfirmationIDTagInfoStatusConcurrentTx StopTransactionConfirmationIDTagInfoStatus = "ConcurrentTx"
)

func (StopTransactionConfirmation) ActionName() string { return "StopTransaction" }

func (StopTransactionConfirmation) Version() protocol.Version { return protocol.OCPP16 }

func (StopTransactionConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (StopTransactionConfirmation) SchemaName() string { return "StopTransactionResponse.json" }

func (message StopTransactionConfirmation) Validate() error {
	return validation.Validate("StopTransactionResponse.json", schemaStopTransactionConfirmation, message)
}

func (StopTransactionConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("StopTransactionResponse.json", schemaStopTransactionConfirmation, data)
}
