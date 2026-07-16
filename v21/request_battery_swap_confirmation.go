// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = RequestBatterySwapConfirmation{}

var schemaRequestBatterySwapConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type RequestBatterySwapConfirmation struct {
	Status     RequestBatterySwapConfirmationGenericStatusEnum `json:"status"`
	StatusInfo *RequestBatterySwapConfirmationStatusInfo       `json:"statusInfo,omitempty"`
	CustomData *RequestBatterySwapConfirmationCustomData       `json:"customData,omitempty"`
}

type RequestBatterySwapConfirmationStatusInfo struct {
	ReasonCode     string                                    `json:"reasonCode"`
	AdditionalInfo *string                                   `json:"additionalInfo,omitempty"`
	CustomData     *RequestBatterySwapConfirmationCustomData `json:"customData,omitempty"`
}

type RequestBatterySwapConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

type RequestBatterySwapConfirmationGenericStatusEnum string

const (
	RequestBatterySwapConfirmationGenericStatusEnumAccepted RequestBatterySwapConfirmationGenericStatusEnum = "Accepted"
	RequestBatterySwapConfirmationGenericStatusEnumRejected RequestBatterySwapConfirmationGenericStatusEnum = "Rejected"
)

func (RequestBatterySwapConfirmation) ActionName() string { return "RequestBatterySwap" }

func (RequestBatterySwapConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (RequestBatterySwapConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (RequestBatterySwapConfirmation) SchemaName() string { return "RequestBatterySwapResponse.json" }

func (message RequestBatterySwapConfirmation) Validate() error {
	return validation.Validate("RequestBatterySwapResponse.json", schemaRequestBatterySwapConfirmation, message)
}

func (RequestBatterySwapConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("RequestBatterySwapResponse.json", schemaRequestBatterySwapConfirmation, data)
}
