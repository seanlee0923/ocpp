// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = SetDisplayMessageConfirmation{}

var schemaSetDisplayMessageConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "NotSupportedMessageFormat", "Rejected", "NotSupportedPriority", "NotSupportedState", "UnknownTransaction"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}}, Required: []string{"status"}}

type SetDisplayMessageConfirmation struct {
	CustomData *SetDisplayMessageConfirmationCustomData              `json:"customData,omitempty"`
	Status     SetDisplayMessageConfirmationDisplayMessageStatusEnum `json:"status"`
	StatusInfo *SetDisplayMessageConfirmationStatusInfo              `json:"statusInfo,omitempty"`
}

type SetDisplayMessageConfirmationStatusInfo struct {
	CustomData     *SetDisplayMessageConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                                   `json:"reasonCode"`
	AdditionalInfo *string                                  `json:"additionalInfo,omitempty"`
}

type SetDisplayMessageConfirmationDisplayMessageStatusEnum string

const (
	SetDisplayMessageConfirmationDisplayMessageStatusEnumAccepted                  SetDisplayMessageConfirmationDisplayMessageStatusEnum = "Accepted"
	SetDisplayMessageConfirmationDisplayMessageStatusEnumNotSupportedMessageFormat SetDisplayMessageConfirmationDisplayMessageStatusEnum = "NotSupportedMessageFormat"
	SetDisplayMessageConfirmationDisplayMessageStatusEnumRejected                  SetDisplayMessageConfirmationDisplayMessageStatusEnum = "Rejected"
	SetDisplayMessageConfirmationDisplayMessageStatusEnumNotSupportedPriority      SetDisplayMessageConfirmationDisplayMessageStatusEnum = "NotSupportedPriority"
	SetDisplayMessageConfirmationDisplayMessageStatusEnumNotSupportedState         SetDisplayMessageConfirmationDisplayMessageStatusEnum = "NotSupportedState"
	SetDisplayMessageConfirmationDisplayMessageStatusEnumUnknownTransaction        SetDisplayMessageConfirmationDisplayMessageStatusEnum = "UnknownTransaction"
)

type SetDisplayMessageConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (SetDisplayMessageConfirmation) ActionName() string { return "SetDisplayMessage" }

func (SetDisplayMessageConfirmation) Version() protocol.Version { return protocol.OCPP201 }

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
