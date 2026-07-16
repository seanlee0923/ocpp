// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = GetDisplayMessagesConfirmation{}

var schemaGetDisplayMessagesConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Unknown"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}}, Required: []string{"status"}}

type GetDisplayMessagesConfirmation struct {
	CustomData *GetDisplayMessagesConfirmationCustomData                  `json:"customData,omitempty"`
	Status     GetDisplayMessagesConfirmationGetDisplayMessagesStatusEnum `json:"status"`
	StatusInfo *GetDisplayMessagesConfirmationStatusInfo                  `json:"statusInfo,omitempty"`
}

type GetDisplayMessagesConfirmationStatusInfo struct {
	CustomData     *GetDisplayMessagesConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                                    `json:"reasonCode"`
	AdditionalInfo *string                                   `json:"additionalInfo,omitempty"`
}

type GetDisplayMessagesConfirmationGetDisplayMessagesStatusEnum string

const (
	GetDisplayMessagesConfirmationGetDisplayMessagesStatusEnumAccepted GetDisplayMessagesConfirmationGetDisplayMessagesStatusEnum = "Accepted"
	GetDisplayMessagesConfirmationGetDisplayMessagesStatusEnumUnknown  GetDisplayMessagesConfirmationGetDisplayMessagesStatusEnum = "Unknown"
)

type GetDisplayMessagesConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (GetDisplayMessagesConfirmation) ActionName() string { return "GetDisplayMessages" }

func (GetDisplayMessagesConfirmation) Version() protocol.Version { return protocol.OCPP201 }

func (GetDisplayMessagesConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (GetDisplayMessagesConfirmation) SchemaName() string { return "GetDisplayMessagesResponse.json" }

func (message GetDisplayMessagesConfirmation) Validate() error {
	return validation.Validate("GetDisplayMessagesResponse.json", schemaGetDisplayMessagesConfirmation, message)
}

func (GetDisplayMessagesConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetDisplayMessagesResponse.json", schemaGetDisplayMessagesConfirmation, data)
}
