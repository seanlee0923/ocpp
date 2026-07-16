// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
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
	VendorID string `json:"vendorId"`
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
