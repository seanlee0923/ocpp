// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = GetLogConfirmation{}

var schemaGetLogConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "AcceptedCanceled"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "filename": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type GetLogConfirmation struct {
	Status     GetLogConfirmationLogStatusEnum `json:"status"`
	StatusInfo *GetLogConfirmationStatusInfo   `json:"statusInfo,omitempty"`
	Filename   *string                         `json:"filename,omitempty"`
	CustomData *GetLogConfirmationCustomData   `json:"customData,omitempty"`
}

type GetLogConfirmationStatusInfo struct {
	ReasonCode     string                        `json:"reasonCode"`
	AdditionalInfo *string                       `json:"additionalInfo,omitempty"`
	CustomData     *GetLogConfirmationCustomData `json:"customData,omitempty"`
}

type GetLogConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

type GetLogConfirmationLogStatusEnum string

const (
	GetLogConfirmationLogStatusEnumAccepted         GetLogConfirmationLogStatusEnum = "Accepted"
	GetLogConfirmationLogStatusEnumRejected         GetLogConfirmationLogStatusEnum = "Rejected"
	GetLogConfirmationLogStatusEnumAcceptedCanceled GetLogConfirmationLogStatusEnum = "AcceptedCanceled"
)

func (GetLogConfirmation) ActionName() string { return "GetLog" }

func (GetLogConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (GetLogConfirmation) Direction() protocol.PayloadDirection { return protocol.ConfirmationPayload }

func (GetLogConfirmation) SchemaName() string { return "GetLogResponse.json" }

func (message GetLogConfirmation) Validate() error {
	return validation.Validate("GetLogResponse.json", schemaGetLogConfirmation, message)
}

func (GetLogConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetLogResponse.json", schemaGetLogConfirmation, data)
}
