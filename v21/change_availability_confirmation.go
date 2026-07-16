// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ChangeAvailabilityConfirmation{}

var schemaChangeAvailabilityConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "Scheduled"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type ChangeAvailabilityConfirmation struct {
	Status     ChangeAvailabilityConfirmationChangeAvailabilityStatusEnum `json:"status"`
	StatusInfo *ChangeAvailabilityConfirmationStatusInfo                  `json:"statusInfo,omitempty"`
	CustomData *ChangeAvailabilityConfirmationCustomData                  `json:"customData,omitempty"`
}

type ChangeAvailabilityConfirmationStatusInfo struct {
	ReasonCode     string                                    `json:"reasonCode"`
	AdditionalInfo *string                                   `json:"additionalInfo,omitempty"`
	CustomData     *ChangeAvailabilityConfirmationCustomData `json:"customData,omitempty"`
}

type ChangeAvailabilityConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

type ChangeAvailabilityConfirmationChangeAvailabilityStatusEnum string

const (
	ChangeAvailabilityConfirmationChangeAvailabilityStatusEnumAccepted  ChangeAvailabilityConfirmationChangeAvailabilityStatusEnum = "Accepted"
	ChangeAvailabilityConfirmationChangeAvailabilityStatusEnumRejected  ChangeAvailabilityConfirmationChangeAvailabilityStatusEnum = "Rejected"
	ChangeAvailabilityConfirmationChangeAvailabilityStatusEnumScheduled ChangeAvailabilityConfirmationChangeAvailabilityStatusEnum = "Scheduled"
)

func (ChangeAvailabilityConfirmation) ActionName() string { return "ChangeAvailability" }

func (ChangeAvailabilityConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (ChangeAvailabilityConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (ChangeAvailabilityConfirmation) SchemaName() string { return "ChangeAvailabilityResponse.json" }

func (message ChangeAvailabilityConfirmation) Validate() error {
	return validation.Validate("ChangeAvailabilityResponse.json", schemaChangeAvailabilityConfirmation, message)
}

func (ChangeAvailabilityConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ChangeAvailabilityResponse.json", schemaChangeAvailabilityConfirmation, data)
}
