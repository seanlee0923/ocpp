// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = GetChargingProfilesConfirmation{}

var schemaGetChargingProfilesConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "NoProfiles"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}}, Required: []string{"status"}}

type GetChargingProfilesConfirmation struct {
	CustomData *GetChargingProfilesConfirmationCustomData                  `json:"customData,omitempty"`
	Status     GetChargingProfilesConfirmationGetChargingProfileStatusEnum `json:"status"`
	StatusInfo *GetChargingProfilesConfirmationStatusInfo                  `json:"statusInfo,omitempty"`
}

type GetChargingProfilesConfirmationStatusInfo struct {
	CustomData     *GetChargingProfilesConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                                     `json:"reasonCode"`
	AdditionalInfo *string                                    `json:"additionalInfo,omitempty"`
}

type GetChargingProfilesConfirmationGetChargingProfileStatusEnum string

const (
	GetChargingProfilesConfirmationGetChargingProfileStatusEnumAccepted   GetChargingProfilesConfirmationGetChargingProfileStatusEnum = "Accepted"
	GetChargingProfilesConfirmationGetChargingProfileStatusEnumNoProfiles GetChargingProfilesConfirmationGetChargingProfileStatusEnum = "NoProfiles"
)

type GetChargingProfilesConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (GetChargingProfilesConfirmation) ActionName() string { return "GetChargingProfiles" }

func (GetChargingProfilesConfirmation) Version() protocol.Version { return protocol.OCPP201 }

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
