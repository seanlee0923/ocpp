// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = SetChargingProfileConfirmation{}

var schemaSetChargingProfileConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}}, Required: []string{"status"}}

type SetChargingProfileConfirmation struct {
	CustomData *SetChargingProfileConfirmationCustomData               `json:"customData,omitempty"`
	Status     SetChargingProfileConfirmationChargingProfileStatusEnum `json:"status"`
	StatusInfo *SetChargingProfileConfirmationStatusInfo               `json:"statusInfo,omitempty"`
}

type SetChargingProfileConfirmationStatusInfo struct {
	CustomData     *SetChargingProfileConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                                    `json:"reasonCode"`
	AdditionalInfo *string                                   `json:"additionalInfo,omitempty"`
}

type SetChargingProfileConfirmationChargingProfileStatusEnum string

const (
	SetChargingProfileConfirmationChargingProfileStatusEnumAccepted SetChargingProfileConfirmationChargingProfileStatusEnum = "Accepted"
	SetChargingProfileConfirmationChargingProfileStatusEnumRejected SetChargingProfileConfirmationChargingProfileStatusEnum = "Rejected"
)

type SetChargingProfileConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value SetChargingProfileConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *SetChargingProfileConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (SetChargingProfileConfirmation) ActionName() string { return "SetChargingProfile" }

func (SetChargingProfileConfirmation) Version() protocol.Version { return protocol.OCPP201 }

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
