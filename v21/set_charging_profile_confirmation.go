// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = SetChargingProfileConfirmation{}

var schemaSetChargingProfileConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type SetChargingProfileConfirmation struct {
	Status     SetChargingProfileConfirmationChargingProfileStatusEnum `json:"status"`
	StatusInfo *SetChargingProfileConfirmationStatusInfo               `json:"statusInfo,omitempty"`
	CustomData *SetChargingProfileConfirmationCustomData               `json:"customData,omitempty"`
}

type SetChargingProfileConfirmationStatusInfo struct {
	ReasonCode     string                                    `json:"reasonCode"`
	AdditionalInfo *string                                   `json:"additionalInfo,omitempty"`
	CustomData     *SetChargingProfileConfirmationCustomData `json:"customData,omitempty"`
}

type SetChargingProfileConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

type SetChargingProfileConfirmationChargingProfileStatusEnum string

const (
	SetChargingProfileConfirmationChargingProfileStatusEnumAccepted SetChargingProfileConfirmationChargingProfileStatusEnum = "Accepted"
	SetChargingProfileConfirmationChargingProfileStatusEnumRejected SetChargingProfileConfirmationChargingProfileStatusEnum = "Rejected"
)

func (SetChargingProfileConfirmation) ActionName() string { return "SetChargingProfile" }

func (SetChargingProfileConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (SetChargingProfileConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (SetChargingProfileConfirmation) SchemaName() string { return "SetChargingProfileResponse.json" }

func (message SetChargingProfileConfirmation) Validate() error {
	return validation.Validate("SetChargingProfileResponse.json", schemaSetChargingProfileConfirmation, message)
}

func (SetChargingProfileConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("SetChargingProfileResponse.json", schemaSetChargingProfileConfirmation, data)
}
