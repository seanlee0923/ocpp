// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = TriggerMessageConfirmation{}

var schemaTriggerMessageConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "NotImplemented"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}}, Required: []string{"status"}}

type TriggerMessageConfirmation struct {
	CustomData *TriggerMessageConfirmationCustomData              `json:"customData,omitempty"`
	Status     TriggerMessageConfirmationTriggerMessageStatusEnum `json:"status"`
	StatusInfo *TriggerMessageConfirmationStatusInfo              `json:"statusInfo,omitempty"`
}

type TriggerMessageConfirmationStatusInfo struct {
	CustomData     *TriggerMessageConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                                `json:"reasonCode"`
	AdditionalInfo *string                               `json:"additionalInfo,omitempty"`
}

type TriggerMessageConfirmationTriggerMessageStatusEnum string

const (
	TriggerMessageConfirmationTriggerMessageStatusEnumAccepted       TriggerMessageConfirmationTriggerMessageStatusEnum = "Accepted"
	TriggerMessageConfirmationTriggerMessageStatusEnumRejected       TriggerMessageConfirmationTriggerMessageStatusEnum = "Rejected"
	TriggerMessageConfirmationTriggerMessageStatusEnumNotImplemented TriggerMessageConfirmationTriggerMessageStatusEnum = "NotImplemented"
)

type TriggerMessageConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (TriggerMessageConfirmation) ActionName() string { return "TriggerMessage" }

func (TriggerMessageConfirmation) Version() protocol.Version { return protocol.OCPP201 }

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
