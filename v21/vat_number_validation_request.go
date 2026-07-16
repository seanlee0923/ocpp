// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = VatNumberValidationRequest{}

var schemaVatNumberValidationRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vatNumber": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "evseId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"vatNumber"}}

type VatNumberValidationRequest struct {
	VatNumber  string                                `json:"vatNumber"`
	EVSEID     *int                                  `json:"evseId,omitempty"`
	CustomData *VatNumberValidationRequestCustomData `json:"customData,omitempty"`
}

type VatNumberValidationRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value VatNumberValidationRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *VatNumberValidationRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (VatNumberValidationRequest) ActionName() string { return "VatNumberValidation" }

func (VatNumberValidationRequest) Version() protocol.Version { return protocol.OCPP21 }

func (VatNumberValidationRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (VatNumberValidationRequest) SchemaName() string { return "VatNumberValidationRequest.json" }

func (message VatNumberValidationRequest) Validate() error {
	return validation.Validate("VatNumberValidationRequest.json", schemaVatNumberValidationRequest, message)
}

func (VatNumberValidationRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("VatNumberValidationRequest.json", schemaVatNumberValidationRequest, data)
}
