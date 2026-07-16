// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = NotifyEVChargingNeedsConfirmation{}

var schemaNotifyEVChargingNeedsConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "Processing", "NoChargingProfile"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type NotifyEVChargingNeedsConfirmation struct {
	Status     NotifyEVChargingNeedsConfirmationNotifyEVChargingNeedsStatusEnum `json:"status"`
	StatusInfo *NotifyEVChargingNeedsConfirmationStatusInfo                     `json:"statusInfo,omitempty"`
	CustomData *NotifyEVChargingNeedsConfirmationCustomData                     `json:"customData,omitempty"`
}

type NotifyEVChargingNeedsConfirmationStatusInfo struct {
	ReasonCode     string                                       `json:"reasonCode"`
	AdditionalInfo *string                                      `json:"additionalInfo,omitempty"`
	CustomData     *NotifyEVChargingNeedsConfirmationCustomData `json:"customData,omitempty"`
}

type NotifyEVChargingNeedsConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

type NotifyEVChargingNeedsConfirmationNotifyEVChargingNeedsStatusEnum string

const (
	NotifyEVChargingNeedsConfirmationNotifyEVChargingNeedsStatusEnumAccepted          NotifyEVChargingNeedsConfirmationNotifyEVChargingNeedsStatusEnum = "Accepted"
	NotifyEVChargingNeedsConfirmationNotifyEVChargingNeedsStatusEnumRejected          NotifyEVChargingNeedsConfirmationNotifyEVChargingNeedsStatusEnum = "Rejected"
	NotifyEVChargingNeedsConfirmationNotifyEVChargingNeedsStatusEnumProcessing        NotifyEVChargingNeedsConfirmationNotifyEVChargingNeedsStatusEnum = "Processing"
	NotifyEVChargingNeedsConfirmationNotifyEVChargingNeedsStatusEnumNoChargingProfile NotifyEVChargingNeedsConfirmationNotifyEVChargingNeedsStatusEnum = "NoChargingProfile"
)

func (NotifyEVChargingNeedsConfirmation) ActionName() string { return "NotifyEVChargingNeeds" }

func (NotifyEVChargingNeedsConfirmation) Version() protocol.Version { return protocol.OCPP21 }

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
