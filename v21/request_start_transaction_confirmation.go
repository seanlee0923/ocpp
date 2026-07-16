// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = RequestStartTransactionConfirmation{}

var schemaRequestStartTransactionConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "transactionId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type RequestStartTransactionConfirmation struct {
	Status        RequestStartTransactionConfirmationRequestStartStopStatusEnum `json:"status"`
	StatusInfo    *RequestStartTransactionConfirmationStatusInfo                `json:"statusInfo,omitempty"`
	TransactionID *string                                                       `json:"transactionId,omitempty"`
	CustomData    *RequestStartTransactionConfirmationCustomData                `json:"customData,omitempty"`
}

type RequestStartTransactionConfirmationStatusInfo struct {
	ReasonCode     string                                         `json:"reasonCode"`
	AdditionalInfo *string                                        `json:"additionalInfo,omitempty"`
	CustomData     *RequestStartTransactionConfirmationCustomData `json:"customData,omitempty"`
}

type RequestStartTransactionConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

type RequestStartTransactionConfirmationRequestStartStopStatusEnum string

const (
	RequestStartTransactionConfirmationRequestStartStopStatusEnumAccepted RequestStartTransactionConfirmationRequestStartStopStatusEnum = "Accepted"
	RequestStartTransactionConfirmationRequestStartStopStatusEnumRejected RequestStartTransactionConfirmationRequestStartStopStatusEnum = "Rejected"
)

func (RequestStartTransactionConfirmation) ActionName() string { return "RequestStartTransaction" }

func (RequestStartTransactionConfirmation) Version() protocol.Version { return protocol.OCPP21 }

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
