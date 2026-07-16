// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = GetLocalListVersionConfirmation{}

var schemaGetLocalListVersionConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"versionNumber": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"versionNumber"}}

type GetLocalListVersionConfirmation struct {
	VersionNumber int                                        `json:"versionNumber"`
	CustomData    *GetLocalListVersionConfirmationCustomData `json:"customData,omitempty"`
}

type GetLocalListVersionConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (GetLocalListVersionConfirmation) ActionName() string { return "GetLocalListVersion" }

func (GetLocalListVersionConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (GetLocalListVersionConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (GetLocalListVersionConfirmation) SchemaName() string { return "GetLocalListVersionResponse.json" }

func (message GetLocalListVersionConfirmation) Validate() error {
	return validation.Validate("GetLocalListVersionResponse.json", schemaGetLocalListVersionConfirmation, message)
}

func (GetLocalListVersionConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetLocalListVersionResponse.json", schemaGetLocalListVersionConfirmation, data)
}
