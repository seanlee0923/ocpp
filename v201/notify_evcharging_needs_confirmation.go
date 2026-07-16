// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = NotifyEVChargingNeedsConfirmation{}

var schemaNotifyEVChargingNeedsConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "Processing"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}}, Required: []string{"status"}}

type NotifyEVChargingNeedsConfirmation struct {
	CustomData *NotifyEVChargingNeedsConfirmationCustomData                     `json:"customData,omitempty"`
	Status     NotifyEVChargingNeedsConfirmationNotifyEVChargingNeedsStatusEnum `json:"status"`
	StatusInfo *NotifyEVChargingNeedsConfirmationStatusInfo                     `json:"statusInfo,omitempty"`
}

type NotifyEVChargingNeedsConfirmationStatusInfo struct {
	CustomData     *NotifyEVChargingNeedsConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                                       `json:"reasonCode"`
	AdditionalInfo *string                                      `json:"additionalInfo,omitempty"`
}

type NotifyEVChargingNeedsConfirmationNotifyEVChargingNeedsStatusEnum string

const (
	NotifyEVChargingNeedsConfirmationNotifyEVChargingNeedsStatusEnumAccepted   NotifyEVChargingNeedsConfirmationNotifyEVChargingNeedsStatusEnum = "Accepted"
	NotifyEVChargingNeedsConfirmationNotifyEVChargingNeedsStatusEnumRejected   NotifyEVChargingNeedsConfirmationNotifyEVChargingNeedsStatusEnum = "Rejected"
	NotifyEVChargingNeedsConfirmationNotifyEVChargingNeedsStatusEnumProcessing NotifyEVChargingNeedsConfirmationNotifyEVChargingNeedsStatusEnum = "Processing"
)

type NotifyEVChargingNeedsConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (NotifyEVChargingNeedsConfirmation) ActionName() string { return "NotifyEVChargingNeeds" }

func (NotifyEVChargingNeedsConfirmation) Version() protocol.Version { return protocol.OCPP201 }

func (NotifyEVChargingNeedsConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (NotifyEVChargingNeedsConfirmation) SchemaName() string {
	return "NotifyEVChargingNeedsResponse.json"
}

func (message NotifyEVChargingNeedsConfirmation) Validate() error {
	return validation.Validate("NotifyEVChargingNeedsResponse.json", schemaNotifyEVChargingNeedsConfirmation, message)
}

func (NotifyEVChargingNeedsConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("NotifyEVChargingNeedsResponse.json", schemaNotifyEVChargingNeedsConfirmation, data)
}
