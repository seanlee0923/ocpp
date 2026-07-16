// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = SetDisplayMessageConfirmation{}

var schemaSetDisplayMessageConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "NotSupportedMessageFormat", "Rejected", "NotSupportedPriority", "NotSupportedState", "UnknownTransaction", "LanguageNotSupported"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type SetDisplayMessageConfirmation struct {
	Status     SetDisplayMessageConfirmationDisplayMessageStatusEnum `json:"status"`
	StatusInfo *SetDisplayMessageConfirmationStatusInfo              `json:"statusInfo,omitempty"`
	CustomData *SetDisplayMessageConfirmationCustomData              `json:"customData,omitempty"`
}

type SetDisplayMessageConfirmationStatusInfo struct {
	ReasonCode     string                                   `json:"reasonCode"`
	AdditionalInfo *string                                  `json:"additionalInfo,omitempty"`
	CustomData     *SetDisplayMessageConfirmationCustomData `json:"customData,omitempty"`
}

type SetDisplayMessageConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

type SetDisplayMessageConfirmationDisplayMessageStatusEnum string

const (
	SetDisplayMessageConfirmationDisplayMessageStatusEnumAccepted                  SetDisplayMessageConfirmationDisplayMessageStatusEnum = "Accepted"
	SetDisplayMessageConfirmationDisplayMessageStatusEnumNotSupportedMessageFormat SetDisplayMessageConfirmationDisplayMessageStatusEnum = "NotSupportedMessageFormat"
	SetDisplayMessageConfirmationDisplayMessageStatusEnumRejected                  SetDisplayMessageConfirmationDisplayMessageStatusEnum = "Rejected"
	SetDisplayMessageConfirmationDisplayMessageStatusEnumNotSupportedPriority      SetDisplayMessageConfirmationDisplayMessageStatusEnum = "NotSupportedPriority"
	SetDisplayMessageConfirmationDisplayMessageStatusEnumNotSupportedState         SetDisplayMessageConfirmationDisplayMessageStatusEnum = "NotSupportedState"
	SetDisplayMessageConfirmationDisplayMessageStatusEnumUnknownTransaction        SetDisplayMessageConfirmationDisplayMessageStatusEnum = "UnknownTransaction"
	SetDisplayMessageConfirmationDisplayMessageStatusEnumLanguageNotSupported      SetDisplayMessageConfirmationDisplayMessageStatusEnum = "LanguageNotSupported"
)

func (SetDisplayMessageConfirmation) ActionName() string { return "SetDisplayMessage" }

func (SetDisplayMessageConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (SetDisplayMessageConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (SetDisplayMessageConfirmation) SchemaName() string { return "SetDisplayMessageResponse.json" }

func (message SetDisplayMessageConfirmation) Validate() error {
	return validation.Validate("SetDisplayMessageResponse.json", schemaSetDisplayMessageConfirmation, message)
}

func (SetDisplayMessageConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("SetDisplayMessageResponse.json", schemaSetDisplayMessageConfirmation, data)
}
