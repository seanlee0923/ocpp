// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ClearChargingProfileConfirmation{}

var schemaClearChargingProfileConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Unknown"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}}, Required: []string{"status"}}

type ClearChargingProfileConfirmation struct {
	CustomData *ClearChargingProfileConfirmationCustomData                    `json:"customData,omitempty"`
	Status     ClearChargingProfileConfirmationClearChargingProfileStatusEnum `json:"status"`
	StatusInfo *ClearChargingProfileConfirmationStatusInfo                    `json:"statusInfo,omitempty"`
}

type ClearChargingProfileConfirmationStatusInfo struct {
	CustomData     *ClearChargingProfileConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                                      `json:"reasonCode"`
	AdditionalInfo *string                                     `json:"additionalInfo,omitempty"`
}

type ClearChargingProfileConfirmationClearChargingProfileStatusEnum string

const (
	ClearChargingProfileConfirmationClearChargingProfileStatusEnumAccepted ClearChargingProfileConfirmationClearChargingProfileStatusEnum = "Accepted"
	ClearChargingProfileConfirmationClearChargingProfileStatusEnumUnknown  ClearChargingProfileConfirmationClearChargingProfileStatusEnum = "Unknown"
)

type ClearChargingProfileConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (ClearChargingProfileConfirmation) ActionName() string { return "ClearChargingProfile" }

func (ClearChargingProfileConfirmation) Version() protocol.Version { return protocol.OCPP201 }

func (ClearChargingProfileConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (ClearChargingProfileConfirmation) SchemaName() string {
	return "ClearChargingProfileResponse.json"
}

func (message ClearChargingProfileConfirmation) Validate() error {
	return validation.Validate("ClearChargingProfileResponse.json", schemaClearChargingProfileConfirmation, message)
}

func (ClearChargingProfileConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ClearChargingProfileResponse.json", schemaClearChargingProfileConfirmation, data)
}
