// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = SetNetworkProfileConfirmation{}

var schemaSetNetworkProfileConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "Failed"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type SetNetworkProfileConfirmation struct {
	Status     SetNetworkProfileConfirmationSetNetworkProfileStatusEnum `json:"status"`
	StatusInfo *SetNetworkProfileConfirmationStatusInfo                 `json:"statusInfo,omitempty"`
	CustomData *SetNetworkProfileConfirmationCustomData                 `json:"customData,omitempty"`
}

type SetNetworkProfileConfirmationStatusInfo struct {
	ReasonCode     string                                   `json:"reasonCode"`
	AdditionalInfo *string                                  `json:"additionalInfo,omitempty"`
	CustomData     *SetNetworkProfileConfirmationCustomData `json:"customData,omitempty"`
}

type SetNetworkProfileConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

type SetNetworkProfileConfirmationSetNetworkProfileStatusEnum string

const (
	SetNetworkProfileConfirmationSetNetworkProfileStatusEnumAccepted SetNetworkProfileConfirmationSetNetworkProfileStatusEnum = "Accepted"
	SetNetworkProfileConfirmationSetNetworkProfileStatusEnumRejected SetNetworkProfileConfirmationSetNetworkProfileStatusEnum = "Rejected"
	SetNetworkProfileConfirmationSetNetworkProfileStatusEnumFailed   SetNetworkProfileConfirmationSetNetworkProfileStatusEnum = "Failed"
)

func (SetNetworkProfileConfirmation) ActionName() string { return "SetNetworkProfile" }

func (SetNetworkProfileConfirmation) Version() protocol.Version { return protocol.OCPP21 }

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
