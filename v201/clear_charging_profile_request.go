// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ClearChargingProfileRequest{}

var schemaClearChargingProfileRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "chargingProfileId": &validation.Schema{Type: "integer", AllowAdditional: true}, "chargingProfileCriteria": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "evseId": &validation.Schema{Type: "integer", AllowAdditional: true}, "chargingProfilePurpose": &validation.Schema{Type: "string", Enum: []string{"ChargingStationExternalConstraints", "ChargingStationMaxProfile", "TxDefaultProfile", "TxProfile"}}, "stackLevel": &validation.Schema{Type: "integer", AllowAdditional: true}}}}}

type ClearChargingProfileRequest struct {
	CustomData              *ClearChargingProfileRequestCustomData           `json:"customData,omitempty"`
	ChargingProfileID       *int                                             `json:"chargingProfileId,omitempty"`
	ChargingProfileCriteria *ClearChargingProfileRequestClearChargingProfile `json:"chargingProfileCriteria,omitempty"`
}

type ClearChargingProfileRequestClearChargingProfile struct {
	CustomData             *ClearChargingProfileRequestCustomData                 `json:"customData,omitempty"`
	EVSEID                 *int                                                   `json:"evseId,omitempty"`
	ChargingProfilePurpose *ClearChargingProfileRequestChargingProfilePurposeEnum `json:"chargingProfilePurpose,omitempty"`
	StackLevel             *int                                                   `json:"stackLevel,omitempty"`
}

type ClearChargingProfileRequestChargingProfilePurposeEnum string

const (
	ClearChargingProfileRequestChargingProfilePurposeEnumChargingStationExternalConstraints ClearChargingProfileRequestChargingProfilePurposeEnum = "ChargingStationExternalConstraints"
	ClearChargingProfileRequestChargingProfilePurposeEnumChargingStationMaxProfile          ClearChargingProfileRequestChargingProfilePurposeEnum = "ChargingStationMaxProfile"
	ClearChargingProfileRequestChargingProfilePurposeEnumTxDefaultProfile                   ClearChargingProfileRequestChargingProfilePurposeEnum = "TxDefaultProfile"
	ClearChargingProfileRequestChargingProfilePurposeEnumTxProfile                          ClearChargingProfileRequestChargingProfilePurposeEnum = "TxProfile"
)

type ClearChargingProfileRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value ClearChargingProfileRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *ClearChargingProfileRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (ClearChargingProfileRequest) ActionName() string { return "ClearChargingProfile" }

func (ClearChargingProfileRequest) Version() protocol.Version { return protocol.OCPP201 }

func (ClearChargingProfileRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (ClearChargingProfileRequest) SchemaName() string { return "ClearChargingProfileRequest.json" }

func (message ClearChargingProfileRequest) Validate() error {
	return validation.Validate("ClearChargingProfileRequest.json", schemaClearChargingProfileRequest, message)
}

func (ClearChargingProfileRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ClearChargingProfileRequest.json", schemaClearChargingProfileRequest, data)
}
