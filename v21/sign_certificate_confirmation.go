// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = SignCertificateConfirmation{}

var schemaSignCertificateConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type SignCertificateConfirmation struct {
	Status     SignCertificateConfirmationGenericStatusEnum `json:"status"`
	StatusInfo *SignCertificateConfirmationStatusInfo       `json:"statusInfo,omitempty"`
	CustomData *SignCertificateConfirmationCustomData       `json:"customData,omitempty"`
}

type SignCertificateConfirmationStatusInfo struct {
	ReasonCode     string                                 `json:"reasonCode"`
	AdditionalInfo *string                                `json:"additionalInfo,omitempty"`
	CustomData     *SignCertificateConfirmationCustomData `json:"customData,omitempty"`
}

type SignCertificateConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

type SignCertificateConfirmationGenericStatusEnum string

const (
	SignCertificateConfirmationGenericStatusEnumAccepted SignCertificateConfirmationGenericStatusEnum = "Accepted"
	SignCertificateConfirmationGenericStatusEnumRejected SignCertificateConfirmationGenericStatusEnum = "Rejected"
)

func (SignCertificateConfirmation) ActionName() string { return "SignCertificate" }

func (SignCertificateConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (SignCertificateConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (SignCertificateConfirmation) SchemaName() string { return "SignCertificateResponse.json" }

func (message SignCertificateConfirmation) Validate() error {
	return validation.Validate("SignCertificateResponse.json", schemaSignCertificateConfirmation, message)
}

func (SignCertificateConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("SignCertificateResponse.json", schemaSignCertificateConfirmation, data)
}
