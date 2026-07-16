// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = GetChargingProfilesRequest{}

var schemaGetChargingProfilesRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"requestId": &validation.Schema{Type: "integer", AllowAdditional: true}, "evseId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "chargingProfile": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"chargingProfilePurpose": &validation.Schema{Type: "string", Enum: []string{"ChargingStationExternalConstraints", "ChargingStationMaxProfile", "TxDefaultProfile", "TxProfile", "PriorityCharging", "LocalGeneration"}}, "stackLevel": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "chargingProfileId": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "integer", AllowAdditional: true}, MinItems: 1, HasMinItems: true}, "chargingLimitSource": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, MinItems: 1, HasMinItems: true, MaxItems: 4, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"requestId", "chargingProfile"}}

type GetChargingProfilesRequest struct {
	RequestID       int                                                `json:"requestId"`
	EVSEID          *int                                               `json:"evseId,omitempty"`
	ChargingProfile GetChargingProfilesRequestChargingProfileCriterion `json:"chargingProfile"`
	CustomData      *GetChargingProfilesRequestCustomData              `json:"customData,omitempty"`
}

type GetChargingProfilesRequestChargingProfileCriterion struct {
	ChargingProfilePurpose *GetChargingProfilesRequestChargingProfilePurposeEnum `json:"chargingProfilePurpose,omitempty"`
	StackLevel             *int                                                  `json:"stackLevel,omitempty"`
	ChargingProfileID      []int                                                 `json:"chargingProfileId,omitempty"`
	ChargingLimitSource    []string                                              `json:"chargingLimitSource,omitempty"`
	CustomData             *GetChargingProfilesRequestCustomData                 `json:"customData,omitempty"`
}

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

type GetChargingProfilesRequestChargingProfilePurposeEnum string

const (
	GetChargingProfilesRequestChargingProfilePurposeEnumChargingStationExternalConstraints GetChargingProfilesRequestChargingProfilePurposeEnum = "ChargingStationExternalConstraints"
	GetChargingProfilesRequestChargingProfilePurposeEnumChargingStationMaxProfile          GetChargingProfilesRequestChargingProfilePurposeEnum = "ChargingStationMaxProfile"
	GetChargingProfilesRequestChargingProfilePurposeEnumTxDefaultProfile                   GetChargingProfilesRequestChargingProfilePurposeEnum = "TxDefaultProfile"
	GetChargingProfilesRequestChargingProfilePurposeEnumTxProfile                          GetChargingProfilesRequestChargingProfilePurposeEnum = "TxProfile"
	GetChargingProfilesRequestChargingProfilePurposeEnumPriorityCharging                   GetChargingProfilesRequestChargingProfilePurposeEnum = "PriorityCharging"
	GetChargingProfilesRequestChargingProfilePurposeEnumLocalGeneration                    GetChargingProfilesRequestChargingProfilePurposeEnum = "LocalGeneration"
)

func (GetChargingProfilesRequest) ActionName() string { return "GetChargingProfiles" }

func (GetChargingProfilesRequest) Version() protocol.Version { return protocol.OCPP21 }

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
