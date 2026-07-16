// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = ReserveNowConfirmation{}

var schemaReserveNowConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Faulted", "Occupied", "Rejected", "Unavailable"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}}, Required: []string{"status"}}

type ReserveNowConfirmation struct {
	CustomData *ReserveNowConfirmationCustomData          `json:"customData,omitempty"`
	Status     ReserveNowConfirmationReserveNowStatusEnum `json:"status"`
	StatusInfo *ReserveNowConfirmationStatusInfo          `json:"statusInfo,omitempty"`
}

type ReserveNowConfirmationStatusInfo struct {
	CustomData     *ReserveNowConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                            `json:"reasonCode"`
	AdditionalInfo *string                           `json:"additionalInfo,omitempty"`
}

type ReserveNowConfirmationReserveNowStatusEnum string

const (
	ReserveNowConfirmationReserveNowStatusEnumAccepted    ReserveNowConfirmationReserveNowStatusEnum = "Accepted"
	ReserveNowConfirmationReserveNowStatusEnumFaulted     ReserveNowConfirmationReserveNowStatusEnum = "Faulted"
	ReserveNowConfirmationReserveNowStatusEnumOccupied    ReserveNowConfirmationReserveNowStatusEnum = "Occupied"
	ReserveNowConfirmationReserveNowStatusEnumRejected    ReserveNowConfirmationReserveNowStatusEnum = "Rejected"
	ReserveNowConfirmationReserveNowStatusEnumUnavailable ReserveNowConfirmationReserveNowStatusEnum = "Unavailable"
)

type ReserveNowConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (ReserveNowConfirmation) ActionName() string { return "ReserveNow" }

func (ReserveNowConfirmation) Version() protocol.Version { return protocol.OCPP201 }

func (ReserveNowConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (ReserveNowConfirmation) SchemaName() string { return "ReserveNowResponse.json" }

func (message ReserveNowConfirmation) Validate() error {
	return validation.Validate("ReserveNowResponse.json", schemaReserveNowConfirmation, message)
}

func (ReserveNowConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ReserveNowResponse.json", schemaReserveNowConfirmation, data)
}
