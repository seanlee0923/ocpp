// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = TriggerMessageConfirmation{}

var schemaTriggerMessageConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "NotImplemented"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type TriggerMessageConfirmation struct {
	Status     TriggerMessageConfirmationTriggerMessageStatusEnum `json:"status"`
	StatusInfo *TriggerMessageConfirmationStatusInfo              `json:"statusInfo,omitempty"`
	CustomData *TriggerMessageConfirmationCustomData              `json:"customData,omitempty"`
}

type TriggerMessageConfirmationStatusInfo struct {
	ReasonCode     string                                `json:"reasonCode"`
	AdditionalInfo *string                               `json:"additionalInfo,omitempty"`
	CustomData     *TriggerMessageConfirmationCustomData `json:"customData,omitempty"`
}

type TriggerMessageConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

type TriggerMessageConfirmationTriggerMessageStatusEnum string

const (
	TriggerMessageConfirmationTriggerMessageStatusEnumAccepted       TriggerMessageConfirmationTriggerMessageStatusEnum = "Accepted"
	TriggerMessageConfirmationTriggerMessageStatusEnumRejected       TriggerMessageConfirmationTriggerMessageStatusEnum = "Rejected"
	TriggerMessageConfirmationTriggerMessageStatusEnumNotImplemented TriggerMessageConfirmationTriggerMessageStatusEnum = "NotImplemented"
)

func (TriggerMessageConfirmation) ActionName() string { return "TriggerMessage" }

func (TriggerMessageConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (TriggerMessageConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (TriggerMessageConfirmation) SchemaName() string { return "TriggerMessageResponse.json" }

func (message TriggerMessageConfirmation) Validate() error {
	return validation.Validate("TriggerMessageResponse.json", schemaTriggerMessageConfirmation, message)
}

func (TriggerMessageConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("TriggerMessageResponse.json", schemaTriggerMessageConfirmation, data)
}
