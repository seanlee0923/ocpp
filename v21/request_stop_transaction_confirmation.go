// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = RequestStopTransactionConfirmation{}

var schemaRequestStopTransactionConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type RequestStopTransactionConfirmation struct {
	Status     RequestStopTransactionConfirmationRequestStartStopStatusEnum `json:"status"`
	StatusInfo *RequestStopTransactionConfirmationStatusInfo                `json:"statusInfo,omitempty"`
	CustomData *RequestStopTransactionConfirmationCustomData                `json:"customData,omitempty"`
}

type RequestStopTransactionConfirmationStatusInfo struct {
	ReasonCode     string                                        `json:"reasonCode"`
	AdditionalInfo *string                                       `json:"additionalInfo,omitempty"`
	CustomData     *RequestStopTransactionConfirmationCustomData `json:"customData,omitempty"`
}

type RequestStopTransactionConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

type RequestStopTransactionConfirmationRequestStartStopStatusEnum string

const (
	RequestStopTransactionConfirmationRequestStartStopStatusEnumAccepted RequestStopTransactionConfirmationRequestStartStopStatusEnum = "Accepted"
	RequestStopTransactionConfirmationRequestStartStopStatusEnumRejected RequestStopTransactionConfirmationRequestStartStopStatusEnum = "Rejected"
)

func (RequestStopTransactionConfirmation) ActionName() string { return "RequestStopTransaction" }

func (RequestStopTransactionConfirmation) Version() protocol.Version { return protocol.OCPP21 }

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
