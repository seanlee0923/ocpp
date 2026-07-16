// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = RequestStartTransactionConfirmation{}

var schemaRequestStartTransactionConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}, "transactionId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}}, Required: []string{"status"}}

type RequestStartTransactionConfirmation struct {
	CustomData    *RequestStartTransactionConfirmationCustomData                `json:"customData,omitempty"`
	Status        RequestStartTransactionConfirmationRequestStartStopStatusEnum `json:"status"`
	StatusInfo    *RequestStartTransactionConfirmationStatusInfo                `json:"statusInfo,omitempty"`
	TransactionID *string                                                       `json:"transactionId,omitempty"`
}

type RequestStartTransactionConfirmationStatusInfo struct {
	CustomData     *RequestStartTransactionConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                                         `json:"reasonCode"`
	AdditionalInfo *string                                        `json:"additionalInfo,omitempty"`
}

type RequestStartTransactionConfirmationRequestStartStopStatusEnum string

const (
	RequestStartTransactionConfirmationRequestStartStopStatusEnumAccepted RequestStartTransactionConfirmationRequestStartStopStatusEnum = "Accepted"
	RequestStartTransactionConfirmationRequestStartStopStatusEnumRejected RequestStartTransactionConfirmationRequestStartStopStatusEnum = "Rejected"
)

type RequestStartTransactionConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (RequestStartTransactionConfirmation) ActionName() string { return "RequestStartTransaction" }

func (RequestStartTransactionConfirmation) Version() protocol.Version { return protocol.OCPP201 }

func (RequestStartTransactionConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (RequestStartTransactionConfirmation) SchemaName() string {
	return "RequestStartTransactionResponse.json"
}

func (message RequestStartTransactionConfirmation) Validate() error {
	return validation.Validate("RequestStartTransactionResponse.json", schemaRequestStartTransactionConfirmation, message)
}

func (RequestStartTransactionConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("RequestStartTransactionResponse.json", schemaRequestStartTransactionConfirmation, data)
}
