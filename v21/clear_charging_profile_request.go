// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ClearChargingProfileRequest{}

var schemaClearChargingProfileRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"chargingProfileId": &validation.Schema{Type: "integer", AllowAdditional: true}, "chargingProfileCriteria": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"evseId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "chargingProfilePurpose": &validation.Schema{Type: "string", Enum: []string{"ChargingStationExternalConstraints", "ChargingStationMaxProfile", "TxDefaultProfile", "TxProfile", "PriorityCharging", "LocalGeneration"}}, "stackLevel": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type ClearChargingProfileRequest struct {
	ChargingProfileID       *int                                             `json:"chargingProfileId,omitempty"`
	ChargingProfileCriteria *ClearChargingProfileRequestClearChargingProfile `json:"chargingProfileCriteria,omitempty"`
	CustomData              *ClearChargingProfileRequestCustomData           `json:"customData,omitempty"`
}

type ClearChargingProfileRequestClearChargingProfile struct {
	EVSEID                 *int                                                   `json:"evseId,omitempty"`
	ChargingProfilePurpose *ClearChargingProfileRequestChargingProfilePurposeEnum `json:"chargingProfilePurpose,omitempty"`
	StackLevel             *int                                                   `json:"stackLevel,omitempty"`
	CustomData             *ClearChargingProfileRequestCustomData                 `json:"customData,omitempty"`
}

type ClearChargingProfileRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

type ClearChargingProfileRequestChargingProfilePurposeEnum string

const (
	ClearChargingProfileRequestChargingProfilePurposeEnumChargingStationExternalConstraints ClearChargingProfileRequestChargingProfilePurposeEnum = "ChargingStationExternalConstraints"
	ClearChargingProfileRequestChargingProfilePurposeEnumChargingStationMaxProfile          ClearChargingProfileRequestChargingProfilePurposeEnum = "ChargingStationMaxProfile"
	ClearChargingProfileRequestChargingProfilePurposeEnumTxDefaultProfile                   ClearChargingProfileRequestChargingProfilePurposeEnum = "TxDefaultProfile"
	ClearChargingProfileRequestChargingProfilePurposeEnumTxProfile                          ClearChargingProfileRequestChargingProfilePurposeEnum = "TxProfile"
	ClearChargingProfileRequestChargingProfilePurposeEnumPriorityCharging                   ClearChargingProfileRequestChargingProfilePurposeEnum = "PriorityCharging"
	ClearChargingProfileRequestChargingProfilePurposeEnumLocalGeneration                    ClearChargingProfileRequestChargingProfilePurposeEnum = "LocalGeneration"
)

func (ClearChargingProfileRequest) ActionName() string { return "ClearChargingProfile" }

func (ClearChargingProfileRequest) Version() protocol.Version { return protocol.OCPP21 }

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
