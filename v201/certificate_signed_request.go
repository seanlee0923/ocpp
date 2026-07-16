// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = CertificateSignedRequest{}

var schemaCertificateSignedRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "certificateChain": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 10000, HasMaxLength: true}, "certificateType": &validation.Schema{Type: "string", Enum: []string{"ChargingStationCertificate", "V2GCertificate"}}}, Required: []string{"certificateChain"}}

type CertificateSignedRequest struct {
	CustomData       *CertificateSignedRequestCustomData                `json:"customData,omitempty"`
	CertificateChain string                                             `json:"certificateChain"`
	CertificateType  *CertificateSignedRequestCertificateSigningUseEnum `json:"certificateType,omitempty"`
}

type CertificateSignedRequestCertificateSigningUseEnum string

const (
	CertificateSignedRequestCertificateSigningUseEnumChargingStationCertificate CertificateSignedRequestCertificateSigningUseEnum = "ChargingStationCertificate"
	CertificateSignedRequestCertificateSigningUseEnumV2GCertificate             CertificateSignedRequestCertificateSigningUseEnum = "V2GCertificate"
)

type CertificateSignedRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (CertificateSignedRequest) ActionName() string { return "CertificateSigned" }

func (CertificateSignedRequest) Version() protocol.Version { return protocol.OCPP201 }

func (CertificateSignedRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (CertificateSignedRequest) SchemaName() string { return "CertificateSignedRequest.json" }

func (message CertificateSignedRequest) Validate() error {
	return validation.Validate("CertificateSignedRequest.json", schemaCertificateSignedRequest, message)
}

func (CertificateSignedRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("CertificateSignedRequest.json", schemaCertificateSignedRequest, data)
}
