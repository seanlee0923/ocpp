// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = VatNumberValidationRequest{}

var schemaVatNumberValidationRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vatNumber": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "evseId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"vatNumber"}}

type VatNumberValidationRequest struct {
	VatNumber  string                                `json:"vatNumber"`
	EVSEID     *int                                  `json:"evseId,omitempty"`
	CustomData *VatNumberValidationRequestCustomData `json:"customData,omitempty"`
}

type VatNumberValidationRequestCustomData struct {
	VendorID string `json:"vendorId"`
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
