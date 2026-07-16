// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = CertificateSignedConfirmation{}

var schemaCertificateSignedConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type CertificateSignedConfirmation struct {
	Status     CertificateSignedConfirmationCertificateSignedStatusEnum `json:"status"`
	StatusInfo *CertificateSignedConfirmationStatusInfo                 `json:"statusInfo,omitempty"`
	CustomData *CertificateSignedConfirmationCustomData                 `json:"customData,omitempty"`
}

type CertificateSignedConfirmationStatusInfo struct {
	ReasonCode     string                                   `json:"reasonCode"`
	AdditionalInfo *string                                  `json:"additionalInfo,omitempty"`
	CustomData     *CertificateSignedConfirmationCustomData `json:"customData,omitempty"`
}

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

type CertificateSignedConfirmationCertificateSignedStatusEnum string

const (
	CertificateSignedConfirmationCertificateSignedStatusEnumAccepted CertificateSignedConfirmationCertificateSignedStatusEnum = "Accepted"
	CertificateSignedConfirmationCertificateSignedStatusEnumRejected CertificateSignedConfirmationCertificateSignedStatusEnum = "Rejected"
)

func (CertificateSignedConfirmation) ActionName() string { return "CertificateSigned" }

func (CertificateSignedConfirmation) Version() protocol.Version { return protocol.OCPP21 }

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
