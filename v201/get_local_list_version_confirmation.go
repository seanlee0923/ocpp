// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = GetLocalListVersionConfirmation{}

var schemaGetLocalListVersionConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "versionNumber": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"versionNumber"}}

type GetLocalListVersionConfirmation struct {
	CustomData    *GetLocalListVersionConfirmationCustomData `json:"customData,omitempty"`
	VersionNumber int                                        `json:"versionNumber"`
}

type GetLocalListVersionConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value GetLocalListVersionConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *GetLocalListVersionConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (GetLocalListVersionConfirmation) ActionName() string { return "GetLocalListVersion" }

func (GetLocalListVersionConfirmation) Version() protocol.Version { return protocol.OCPP201 }

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
