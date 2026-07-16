// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = SignCertificateRequest{}

var schemaSignCertificateRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"csr": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 5500, HasMaxLength: true}, "certificateType": &validation.Schema{Type: "string", Enum: []string{"ChargingStationCertificate", "V2GCertificate", "V2G20Certificate"}}, "hashRootCertificate": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"hashAlgorithm": &validation.Schema{Type: "string", Enum: []string{"SHA256", "SHA384", "SHA512"}}, "issuerNameHash": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 128, HasMaxLength: true}, "issuerKeyHash": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 128, HasMaxLength: true}, "serialNumber": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 40, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"hashAlgorithm", "issuerNameHash", "issuerKeyHash", "serialNumber"}}, "requestId": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"csr"}}

type SignCertificateRequest struct {
	Csr                 string                                           `json:"csr"`
	CertificateType     *SignCertificateRequestCertificateSigningUseEnum `json:"certificateType,omitempty"`
	HashRootCertificate *SignCertificateRequestCertificateHashData       `json:"hashRootCertificate,omitempty"`
	RequestID           *int                                             `json:"requestId,omitempty"`
	CustomData          *SignCertificateRequestCustomData                `json:"customData,omitempty"`
}

type SignCertificateRequestCertificateHashData struct {
	HashAlgorithm  SignCertificateRequestHashAlgorithmEnum `json:"hashAlgorithm"`
	IssuerNameHash string                                  `json:"issuerNameHash"`
	IssuerKeyHash  string                                  `json:"issuerKeyHash"`
	SerialNumber   string                                  `json:"serialNumber"`
	CustomData     *SignCertificateRequestCustomData       `json:"customData,omitempty"`
}

type SignCertificateRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value SignCertificateRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *SignCertificateRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

type SignCertificateRequestHashAlgorithmEnum string

const (
	SignCertificateRequestHashAlgorithmEnumSHA256 SignCertificateRequestHashAlgorithmEnum = "SHA256"
	SignCertificateRequestHashAlgorithmEnumSHA384 SignCertificateRequestHashAlgorithmEnum = "SHA384"
	SignCertificateRequestHashAlgorithmEnumSHA512 SignCertificateRequestHashAlgorithmEnum = "SHA512"
)

type SignCertificateRequestCertificateSigningUseEnum string

const (
	SignCertificateRequestCertificateSigningUseEnumChargingStationCertificate SignCertificateRequestCertificateSigningUseEnum = "ChargingStationCertificate"
	SignCertificateRequestCertificateSigningUseEnumV2GCertificate             SignCertificateRequestCertificateSigningUseEnum = "V2GCertificate"
	SignCertificateRequestCertificateSigningUseEnumV2G20Certificate           SignCertificateRequestCertificateSigningUseEnum = "V2G20Certificate"
)

func (SignCertificateRequest) ActionName() string { return "SignCertificate" }

func (SignCertificateRequest) Version() protocol.Version { return protocol.OCPP21 }

func (SignCertificateRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (SignCertificateRequest) SchemaName() string { return "SignCertificateRequest.json" }

func (message SignCertificateRequest) Validate() error {
	return validation.Validate("SignCertificateRequest.json", schemaSignCertificateRequest, message)
}

func (SignCertificateRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("SignCertificateRequest.json", schemaSignCertificateRequest, data)
}
