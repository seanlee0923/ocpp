// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ChangeAvailabilityConfirmation{}

var schemaChangeAvailabilityConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "Scheduled"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}}, Required: []string{"status"}}

type ChangeAvailabilityConfirmation struct {
	CustomData *ChangeAvailabilityConfirmationCustomData                  `json:"customData,omitempty"`
	Status     ChangeAvailabilityConfirmationChangeAvailabilityStatusEnum `json:"status"`
	StatusInfo *ChangeAvailabilityConfirmationStatusInfo                  `json:"statusInfo,omitempty"`
}

type ChangeAvailabilityConfirmationStatusInfo struct {
	CustomData     *ChangeAvailabilityConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                                    `json:"reasonCode"`
	AdditionalInfo *string                                   `json:"additionalInfo,omitempty"`
}

type ChangeAvailabilityConfirmationChangeAvailabilityStatusEnum string

const (
	ChangeAvailabilityConfirmationChangeAvailabilityStatusEnumAccepted  ChangeAvailabilityConfirmationChangeAvailabilityStatusEnum = "Accepted"
	ChangeAvailabilityConfirmationChangeAvailabilityStatusEnumRejected  ChangeAvailabilityConfirmationChangeAvailabilityStatusEnum = "Rejected"
	ChangeAvailabilityConfirmationChangeAvailabilityStatusEnumScheduled ChangeAvailabilityConfirmationChangeAvailabilityStatusEnum = "Scheduled"
)

type ChangeAvailabilityConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value ChangeAvailabilityConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *ChangeAvailabilityConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (ChangeAvailabilityConfirmation) ActionName() string { return "ChangeAvailability" }

func (ChangeAvailabilityConfirmation) Version() protocol.Version { return protocol.OCPP201 }

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
