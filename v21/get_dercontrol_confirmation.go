// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = GetDERControlConfirmation{}

var schemaGetDERControlConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "NotSupported", "NotFound"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type GetDERControlConfirmation struct {
	Status     GetDERControlConfirmationDERControlStatusEnum `json:"status"`
	StatusInfo *GetDERControlConfirmationStatusInfo          `json:"statusInfo,omitempty"`
	CustomData *GetDERControlConfirmationCustomData          `json:"customData,omitempty"`
}

type GetDERControlConfirmationStatusInfo struct {
	ReasonCode     string                               `json:"reasonCode"`
	AdditionalInfo *string                              `json:"additionalInfo,omitempty"`
	CustomData     *GetDERControlConfirmationCustomData `json:"customData,omitempty"`
}

type GetDERControlConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

type GetDERControlConfirmationDERControlStatusEnum string

const (
	GetDERControlConfirmationDERControlStatusEnumAccepted     GetDERControlConfirmationDERControlStatusEnum = "Accepted"
	GetDERControlConfirmationDERControlStatusEnumRejected     GetDERControlConfirmationDERControlStatusEnum = "Rejected"
	GetDERControlConfirmationDERControlStatusEnumNotSupported GetDERControlConfirmationDERControlStatusEnum = "NotSupported"
	GetDERControlConfirmationDERControlStatusEnumNotFound     GetDERControlConfirmationDERControlStatusEnum = "NotFound"
)

func (GetDERControlConfirmation) ActionName() string { return "GetDERControl" }

func (GetDERControlConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (GetDERControlConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (GetDERControlConfirmation) SchemaName() string { return "GetDERControlResponse.json" }

func (message GetDERControlConfirmation) Validate() error {
	return validation.Validate("GetDERControlResponse.json", schemaGetDERControlConfirmation, message)
}

func (GetDERControlConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetDERControlResponse.json", schemaGetDERControlConfirmation, data)
}
