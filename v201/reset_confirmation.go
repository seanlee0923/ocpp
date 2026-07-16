// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ResetConfirmation{}

var schemaResetConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "Scheduled"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}}, Required: []string{"status"}}

type ResetConfirmation struct {
	CustomData *ResetConfirmationCustomData     `json:"customData,omitempty"`
	Status     ResetConfirmationResetStatusEnum `json:"status"`
	StatusInfo *ResetConfirmationStatusInfo     `json:"statusInfo,omitempty"`
}

type ResetConfirmationStatusInfo struct {
	CustomData     *ResetConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                       `json:"reasonCode"`
	AdditionalInfo *string                      `json:"additionalInfo,omitempty"`
}

type ResetConfirmationResetStatusEnum string

const (
	ResetConfirmationResetStatusEnumAccepted  ResetConfirmationResetStatusEnum = "Accepted"
	ResetConfirmationResetStatusEnumRejected  ResetConfirmationResetStatusEnum = "Rejected"
	ResetConfirmationResetStatusEnumScheduled ResetConfirmationResetStatusEnum = "Scheduled"
)

type ResetConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (ResetConfirmation) ActionName() string { return "Reset" }

func (ResetConfirmation) Version() protocol.Version { return protocol.OCPP201 }

func (ResetConfirmation) Direction() protocol.PayloadDirection { return protocol.ConfirmationPayload }

func (ResetConfirmation) SchemaName() string { return "ResetResponse.json" }

func (message ResetConfirmation) Validate() error {
	return validation.Validate("ResetResponse.json", schemaResetConfirmation, message)
}

func (ResetConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ResetResponse.json", schemaResetConfirmation, data)
}
