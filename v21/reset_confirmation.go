// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ResetConfirmation{}

var schemaResetConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "Scheduled"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type ResetConfirmation struct {
	Status     ResetConfirmationResetStatusEnum `json:"status"`
	StatusInfo *ResetConfirmationStatusInfo     `json:"statusInfo,omitempty"`
	CustomData *ResetConfirmationCustomData     `json:"customData,omitempty"`
}

type ResetConfirmationStatusInfo struct {
	ReasonCode     string                       `json:"reasonCode"`
	AdditionalInfo *string                      `json:"additionalInfo,omitempty"`
	CustomData     *ResetConfirmationCustomData `json:"customData,omitempty"`
}

type ResetConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

type ResetConfirmationResetStatusEnum string

const (
	ResetConfirmationResetStatusEnumAccepted  ResetConfirmationResetStatusEnum = "Accepted"
	ResetConfirmationResetStatusEnumRejected  ResetConfirmationResetStatusEnum = "Rejected"
	ResetConfirmationResetStatusEnumScheduled ResetConfirmationResetStatusEnum = "Scheduled"
)

func (ResetConfirmation) ActionName() string { return "Reset" }

func (ResetConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (ResetConfirmation) Direction() protocol.PayloadDirection { return protocol.ConfirmationPayload }

func (ResetConfirmation) SchemaName() string { return "ResetResponse.json" }

func (message ResetConfirmation) Validate() error {
	return validation.Validate("ResetResponse.json", schemaResetConfirmation, message)
}

func (ResetConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ResetResponse.json", schemaResetConfirmation, data)
}
