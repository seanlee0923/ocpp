// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = SignCertificateConfirmation{}

var schemaSignCertificateConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}}, Required: []string{"status"}}

type SignCertificateConfirmation struct {
	CustomData *SignCertificateConfirmationCustomData       `json:"customData,omitempty"`
	Status     SignCertificateConfirmationGenericStatusEnum `json:"status"`
	StatusInfo *SignCertificateConfirmationStatusInfo       `json:"statusInfo,omitempty"`
}

type SignCertificateConfirmationStatusInfo struct {
	CustomData     *SignCertificateConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                                 `json:"reasonCode"`
	AdditionalInfo *string                                `json:"additionalInfo,omitempty"`
}

type SignCertificateConfirmationGenericStatusEnum string

const (
	SignCertificateConfirmationGenericStatusEnumAccepted SignCertificateConfirmationGenericStatusEnum = "Accepted"
	SignCertificateConfirmationGenericStatusEnumRejected SignCertificateConfirmationGenericStatusEnum = "Rejected"
)

type SignCertificateConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (SignCertificateConfirmation) ActionName() string { return "SignCertificate" }

func (SignCertificateConfirmation) Version() protocol.Version { return protocol.OCPP201 }

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
