// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = RequestStopTransactionConfirmation{}

var schemaRequestStopTransactionConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}}, Required: []string{"status"}}

type RequestStopTransactionConfirmation struct {
	CustomData *RequestStopTransactionConfirmationCustomData                `json:"customData,omitempty"`
	Status     RequestStopTransactionConfirmationRequestStartStopStatusEnum `json:"status"`
	StatusInfo *RequestStopTransactionConfirmationStatusInfo                `json:"statusInfo,omitempty"`
}

type RequestStopTransactionConfirmationStatusInfo struct {
	CustomData     *RequestStopTransactionConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                                        `json:"reasonCode"`
	AdditionalInfo *string                                       `json:"additionalInfo,omitempty"`
}

type RequestStopTransactionConfirmationRequestStartStopStatusEnum string

const (
	RequestStopTransactionConfirmationRequestStartStopStatusEnumAccepted RequestStopTransactionConfirmationRequestStartStopStatusEnum = "Accepted"
	RequestStopTransactionConfirmationRequestStartStopStatusEnumRejected RequestStopTransactionConfirmationRequestStartStopStatusEnum = "Rejected"
)

type RequestStopTransactionConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (RequestStopTransactionConfirmation) ActionName() string { return "RequestStopTransaction" }

func (RequestStopTransactionConfirmation) Version() protocol.Version { return protocol.OCPP201 }

func (RequestStopTransactionConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (RequestStopTransactionConfirmation) SchemaName() string {
	return "RequestStopTransactionResponse.json"
}

func (message RequestStopTransactionConfirmation) Validate() error {
	return validation.Validate("RequestStopTransactionResponse.json", schemaRequestStopTransactionConfirmation, message)
}

func (RequestStopTransactionConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("RequestStopTransactionResponse.json", schemaRequestStopTransactionConfirmation, data)
}
