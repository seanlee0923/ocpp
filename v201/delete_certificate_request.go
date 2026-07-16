// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = DeleteCertificateRequest{}

var schemaDeleteCertificateRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "certificateHashData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "hashAlgorithm": &validation.Schema{Type: "string", Enum: []string{"SHA256", "SHA384", "SHA512"}}, "issuerNameHash": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 128, HasMaxLength: true}, "issuerKeyHash": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 128, HasMaxLength: true}, "serialNumber": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 40, HasMaxLength: true}}, Required: []string{"hashAlgorithm", "issuerNameHash", "issuerKeyHash", "serialNumber"}}}, Required: []string{"certificateHashData"}}

type DeleteCertificateRequest struct {
	CustomData          *DeleteCertificateRequestCustomData         `json:"customData,omitempty"`
	CertificateHashData DeleteCertificateRequestCertificateHashData `json:"certificateHashData"`
}

type DeleteCertificateRequestCertificateHashData struct {
	CustomData     *DeleteCertificateRequestCustomData       `json:"customData,omitempty"`
	HashAlgorithm  DeleteCertificateRequestHashAlgorithmEnum `json:"hashAlgorithm"`
	IssuerNameHash string                                    `json:"issuerNameHash"`
	IssuerKeyHash  string                                    `json:"issuerKeyHash"`
	SerialNumber   string                                    `json:"serialNumber"`
}

type DeleteCertificateRequestHashAlgorithmEnum string

const (
	DeleteCertificateRequestHashAlgorithmEnumSHA256 DeleteCertificateRequestHashAlgorithmEnum = "SHA256"
	DeleteCertificateRequestHashAlgorithmEnumSHA384 DeleteCertificateRequestHashAlgorithmEnum = "SHA384"
	DeleteCertificateRequestHashAlgorithmEnumSHA512 DeleteCertificateRequestHashAlgorithmEnum = "SHA512"
)

type DeleteCertificateRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (DeleteCertificateRequest) ActionName() string { return "DeleteCertificate" }

func (DeleteCertificateRequest) Version() protocol.Version { return protocol.OCPP201 }

func (DeleteCertificateRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (DeleteCertificateRequest) SchemaName() string { return "DeleteCertificateRequest.json" }

func (message DeleteCertificateRequest) Validate() error {
	return validation.Validate("DeleteCertificateRequest.json", schemaDeleteCertificateRequest, message)
}

func (DeleteCertificateRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("DeleteCertificateRequest.json", schemaDeleteCertificateRequest, data)
}
