// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = SetNetworkProfileConfirmation{}

var schemaSetNetworkProfileConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "Failed"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}}, Required: []string{"status"}}

type SetNetworkProfileConfirmation struct {
	CustomData *SetNetworkProfileConfirmationCustomData                 `json:"customData,omitempty"`
	Status     SetNetworkProfileConfirmationSetNetworkProfileStatusEnum `json:"status"`
	StatusInfo *SetNetworkProfileConfirmationStatusInfo                 `json:"statusInfo,omitempty"`
}

type SetNetworkProfileConfirmationStatusInfo struct {
	CustomData     *SetNetworkProfileConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                                   `json:"reasonCode"`
	AdditionalInfo *string                                  `json:"additionalInfo,omitempty"`
}

type SetNetworkProfileConfirmationSetNetworkProfileStatusEnum string

const (
	SetNetworkProfileConfirmationSetNetworkProfileStatusEnumAccepted SetNetworkProfileConfirmationSetNetworkProfileStatusEnum = "Accepted"
	SetNetworkProfileConfirmationSetNetworkProfileStatusEnumRejected SetNetworkProfileConfirmationSetNetworkProfileStatusEnum = "Rejected"
	SetNetworkProfileConfirmationSetNetworkProfileStatusEnumFailed   SetNetworkProfileConfirmationSetNetworkProfileStatusEnum = "Failed"
)

type SetNetworkProfileConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value SetNetworkProfileConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *SetNetworkProfileConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (SetNetworkProfileConfirmation) ActionName() string { return "SetNetworkProfile" }

func (SetNetworkProfileConfirmation) Version() protocol.Version { return protocol.OCPP201 }

func (SetNetworkProfileConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (SetNetworkProfileConfirmation) SchemaName() string { return "SetNetworkProfileResponse.json" }

func (message SetNetworkProfileConfirmation) Validate() error {
	return validation.Validate("SetNetworkProfileResponse.json", schemaSetNetworkProfileConfirmation, message)
}

func (SetNetworkProfileConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("SetNetworkProfileResponse.json", schemaSetNetworkProfileConfirmation, data)
}
