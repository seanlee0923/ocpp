// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = GetChargingProfilesRequest{}

var schemaGetChargingProfilesRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "requestId": &validation.Schema{Type: "integer", AllowAdditional: true}, "evseId": &validation.Schema{Type: "integer", AllowAdditional: true}, "chargingProfile": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "chargingProfilePurpose": &validation.Schema{Type: "string", Enum: []string{"ChargingStationExternalConstraints", "ChargingStationMaxProfile", "TxDefaultProfile", "TxProfile"}}, "stackLevel": &validation.Schema{Type: "integer", AllowAdditional: true}, "chargingProfileId": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "integer", AllowAdditional: true}, MinItems: 1, HasMinItems: true}, "chargingLimitSource": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", Enum: []string{"EMS", "Other", "SO", "CSO"}}, MinItems: 1, HasMinItems: true, MaxItems: 4, HasMaxItems: true}}}}, Required: []string{"requestId", "chargingProfile"}}

type GetChargingProfilesRequest struct {
	CustomData      *GetChargingProfilesRequestCustomData              `json:"customData,omitempty"`
	RequestID       int                                                `json:"requestId"`
	EVSEID          *int                                               `json:"evseId,omitempty"`
	ChargingProfile GetChargingProfilesRequestChargingProfileCriterion `json:"chargingProfile"`
}

type GetChargingProfilesRequestChargingProfileCriterion struct {
	CustomData             *GetChargingProfilesRequestCustomData                 `json:"customData,omitempty"`
	ChargingProfilePurpose *GetChargingProfilesRequestChargingProfilePurposeEnum `json:"chargingProfilePurpose,omitempty"`
	StackLevel             *int                                                  `json:"stackLevel,omitempty"`
	ChargingProfileID      []int                                                 `json:"chargingProfileId,omitempty"`
	ChargingLimitSource    []GetChargingProfilesRequestChargingLimitSourceEnum   `json:"chargingLimitSource,omitempty"`
}

type GetChargingProfilesRequestChargingLimitSourceEnum string

const (
	GetChargingProfilesRequestChargingLimitSourceEnumEMS   GetChargingProfilesRequestChargingLimitSourceEnum = "EMS"
	GetChargingProfilesRequestChargingLimitSourceEnumOther GetChargingProfilesRequestChargingLimitSourceEnum = "Other"
	GetChargingProfilesRequestChargingLimitSourceEnumSO    GetChargingProfilesRequestChargingLimitSourceEnum = "SO"
	GetChargingProfilesRequestChargingLimitSourceEnumCSO   GetChargingProfilesRequestChargingLimitSourceEnum = "CSO"
)

type GetChargingProfilesRequestChargingProfilePurposeEnum string

const (
	GetChargingProfilesRequestChargingProfilePurposeEnumChargingStationExternalConstraints GetChargingProfilesRequestChargingProfilePurposeEnum = "ChargingStationExternalConstraints"
	GetChargingProfilesRequestChargingProfilePurposeEnumChargingStationMaxProfile          GetChargingProfilesRequestChargingProfilePurposeEnum = "ChargingStationMaxProfile"
	GetChargingProfilesRequestChargingProfilePurposeEnumTxDefaultProfile                   GetChargingProfilesRequestChargingProfilePurposeEnum = "TxDefaultProfile"
	GetChargingProfilesRequestChargingProfilePurposeEnumTxProfile                          GetChargingProfilesRequestChargingProfilePurposeEnum = "TxProfile"
)

type GetChargingProfilesRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value GetChargingProfilesRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *GetChargingProfilesRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (GetChargingProfilesRequest) ActionName() string { return "GetChargingProfiles" }

func (GetChargingProfilesRequest) Version() protocol.Version { return protocol.OCPP201 }

func (GetChargingProfilesRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (GetChargingProfilesRequest) SchemaName() string { return "GetChargingProfilesRequest.json" }

func (message GetChargingProfilesRequest) Validate() error {
	return validation.Validate("GetChargingProfilesRequest.json", schemaGetChargingProfilesRequest, message)
}

func (GetChargingProfilesRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetChargingProfilesRequest.json", schemaGetChargingProfilesRequest, data)
}
