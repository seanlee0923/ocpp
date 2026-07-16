// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = CertificateSignedConfirmation{}

var schemaCertificateSignedConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}}, Required: []string{"status"}}

type CertificateSignedConfirmation struct {
	CustomData *CertificateSignedConfirmationCustomData                 `json:"customData,omitempty"`
	Status     CertificateSignedConfirmationCertificateSignedStatusEnum `json:"status"`
	StatusInfo *CertificateSignedConfirmationStatusInfo                 `json:"statusInfo,omitempty"`
}

type CertificateSignedConfirmationStatusInfo struct {
	CustomData     *CertificateSignedConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                                   `json:"reasonCode"`
	AdditionalInfo *string                                  `json:"additionalInfo,omitempty"`
}

type CertificateSignedConfirmationCertificateSignedStatusEnum string

const (
	CertificateSignedConfirmationCertificateSignedStatusEnumAccepted CertificateSignedConfirmationCertificateSignedStatusEnum = "Accepted"
	CertificateSignedConfirmationCertificateSignedStatusEnumRejected CertificateSignedConfirmationCertificateSignedStatusEnum = "Rejected"
)

type CertificateSignedConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value CertificateSignedConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *CertificateSignedConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (CertificateSignedConfirmation) ActionName() string { return "CertificateSigned" }

func (CertificateSignedConfirmation) Version() protocol.Version { return protocol.OCPP201 }

func (CertificateSignedConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (CertificateSignedConfirmation) SchemaName() string { return "CertificateSignedResponse.json" }

func (message CertificateSignedConfirmation) Validate() error {
	return validation.Validate("CertificateSignedResponse.json", schemaCertificateSignedConfirmation, message)
}

func (CertificateSignedConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("CertificateSignedResponse.json", schemaCertificateSignedConfirmation, data)
}
