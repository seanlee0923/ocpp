// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ClearChargingProfileConfirmation{}

var schemaClearChargingProfileConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Unknown"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type ClearChargingProfileConfirmation struct {
	Status     ClearChargingProfileConfirmationClearChargingProfileStatusEnum `json:"status"`
	StatusInfo *ClearChargingProfileConfirmationStatusInfo                    `json:"statusInfo,omitempty"`
	CustomData *ClearChargingProfileConfirmationCustomData                    `json:"customData,omitempty"`
}

type ClearChargingProfileConfirmationStatusInfo struct {
	ReasonCode     string                                      `json:"reasonCode"`
	AdditionalInfo *string                                     `json:"additionalInfo,omitempty"`
	CustomData     *ClearChargingProfileConfirmationCustomData `json:"customData,omitempty"`
}

type ClearChargingProfileConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value ClearChargingProfileConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *ClearChargingProfileConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

type ClearChargingProfileConfirmationClearChargingProfileStatusEnum string

const (
	ClearChargingProfileConfirmationClearChargingProfileStatusEnumAccepted ClearChargingProfileConfirmationClearChargingProfileStatusEnum = "Accepted"
	ClearChargingProfileConfirmationClearChargingProfileStatusEnumUnknown  ClearChargingProfileConfirmationClearChargingProfileStatusEnum = "Unknown"
)

func (ClearChargingProfileConfirmation) ActionName() string { return "ClearChargingProfile" }

func (ClearChargingProfileConfirmation) Version() protocol.Version { return protocol.OCPP21 }

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
