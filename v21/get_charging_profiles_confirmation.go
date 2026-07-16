// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = GetChargingProfilesConfirmation{}

var schemaGetChargingProfilesConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "NoProfiles"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type GetChargingProfilesConfirmation struct {
	Status     GetChargingProfilesConfirmationGetChargingProfileStatusEnum `json:"status"`
	StatusInfo *GetChargingProfilesConfirmationStatusInfo                  `json:"statusInfo,omitempty"`
	CustomData *GetChargingProfilesConfirmationCustomData                  `json:"customData,omitempty"`
}

type GetChargingProfilesConfirmationStatusInfo struct {
	ReasonCode     string                                     `json:"reasonCode"`
	AdditionalInfo *string                                    `json:"additionalInfo,omitempty"`
	CustomData     *GetChargingProfilesConfirmationCustomData `json:"customData,omitempty"`
}

type GetChargingProfilesConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

type GetChargingProfilesConfirmationGetChargingProfileStatusEnum string

const (
	GetChargingProfilesConfirmationGetChargingProfileStatusEnumAccepted   GetChargingProfilesConfirmationGetChargingProfileStatusEnum = "Accepted"
	GetChargingProfilesConfirmationGetChargingProfileStatusEnumNoProfiles GetChargingProfilesConfirmationGetChargingProfileStatusEnum = "NoProfiles"
)

func (GetChargingProfilesConfirmation) ActionName() string { return "GetChargingProfiles" }

func (GetChargingProfilesConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (GetChargingProfilesConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (GetChargingProfilesConfirmation) SchemaName() string { return "GetChargingProfilesResponse.json" }

func (message GetChargingProfilesConfirmation) Validate() error {
	return validation.Validate("GetChargingProfilesResponse.json", schemaGetChargingProfilesConfirmation, message)
}

func (GetChargingProfilesConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetChargingProfilesResponse.json", schemaGetChargingProfilesConfirmation, data)
}
