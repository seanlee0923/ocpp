// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = VatNumberValidationConfirmation{}

var schemaVatNumberValidationConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"company": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "address1": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 100, HasMaxLength: true}, "address2": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 100, HasMaxLength: true}, "city": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 100, HasMaxLength: true}, "postalCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "country": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"name", "address1", "city", "country"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "vatNumber": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "evseId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"vatNumber", "status"}}

type VatNumberValidationConfirmation struct {
	Company    *VatNumberValidationConfirmationAddress          `json:"company,omitempty"`
	StatusInfo *VatNumberValidationConfirmationStatusInfo       `json:"statusInfo,omitempty"`
	VatNumber  string                                           `json:"vatNumber"`
	EVSEID     *int                                             `json:"evseId,omitempty"`
	Status     VatNumberValidationConfirmationGenericStatusEnum `json:"status"`
	CustomData *VatNumberValidationConfirmationCustomData       `json:"customData,omitempty"`
}

type VatNumberValidationConfirmationGenericStatusEnum string

const (
	VatNumberValidationConfirmationGenericStatusEnumAccepted VatNumberValidationConfirmationGenericStatusEnum = "Accepted"
	VatNumberValidationConfirmationGenericStatusEnumRejected VatNumberValidationConfirmationGenericStatusEnum = "Rejected"
)

type VatNumberValidationConfirmationStatusInfo struct {
	ReasonCode     string                                     `json:"reasonCode"`
	AdditionalInfo *string                                    `json:"additionalInfo,omitempty"`
	CustomData     *VatNumberValidationConfirmationCustomData `json:"customData,omitempty"`
}

type VatNumberValidationConfirmationAddress struct {
	Name       string                                     `json:"name"`
	Address1   string                                     `json:"address1"`
	Address2   *string                                    `json:"address2,omitempty"`
	City       string                                     `json:"city"`
	PostalCode *string                                    `json:"postalCode,omitempty"`
	Country    string                                     `json:"country"`
	CustomData *VatNumberValidationConfirmationCustomData `json:"customData,omitempty"`
}

type VatNumberValidationConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value VatNumberValidationConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *VatNumberValidationConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (VatNumberValidationConfirmation) ActionName() string { return "VatNumberValidation" }

func (VatNumberValidationConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (VatNumberValidationConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (VatNumberValidationConfirmation) SchemaName() string { return "VatNumberValidationResponse.json" }

func (message VatNumberValidationConfirmation) Validate() error {
	return validation.Validate("VatNumberValidationResponse.json", schemaVatNumberValidationConfirmation, message)
}

func (VatNumberValidationConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("VatNumberValidationResponse.json", schemaVatNumberValidationConfirmation, data)
}
