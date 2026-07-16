// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = ReserveNowConfirmation{}

var schemaReserveNowConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Faulted", "Occupied", "Rejected", "Unavailable"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type ReserveNowConfirmation struct {
	Status     ReserveNowConfirmationReserveNowStatusEnum `json:"status"`
	StatusInfo *ReserveNowConfirmationStatusInfo          `json:"statusInfo,omitempty"`
	CustomData *ReserveNowConfirmationCustomData          `json:"customData,omitempty"`
}

type ReserveNowConfirmationStatusInfo struct {
	ReasonCode     string                            `json:"reasonCode"`
	AdditionalInfo *string                           `json:"additionalInfo,omitempty"`
	CustomData     *ReserveNowConfirmationCustomData `json:"customData,omitempty"`
}

type ReserveNowConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

type ReserveNowConfirmationReserveNowStatusEnum string

const (
	ReserveNowConfirmationReserveNowStatusEnumAccepted    ReserveNowConfirmationReserveNowStatusEnum = "Accepted"
	ReserveNowConfirmationReserveNowStatusEnumFaulted     ReserveNowConfirmationReserveNowStatusEnum = "Faulted"
	ReserveNowConfirmationReserveNowStatusEnumOccupied    ReserveNowConfirmationReserveNowStatusEnum = "Occupied"
	ReserveNowConfirmationReserveNowStatusEnumRejected    ReserveNowConfirmationReserveNowStatusEnum = "Rejected"
	ReserveNowConfirmationReserveNowStatusEnumUnavailable ReserveNowConfirmationReserveNowStatusEnum = "Unavailable"
)

func (ReserveNowConfirmation) ActionName() string { return "ReserveNow" }

func (ReserveNowConfirmation) Version() protocol.Version { return protocol.OCPP21 }

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
