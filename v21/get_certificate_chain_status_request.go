// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = GetCertificateChainStatusRequest{}

var schemaGetCertificateChainStatusRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"certificateStatusRequests": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"certificateHashData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"hashAlgorithm": &validation.Schema{Type: "string", Enum: []string{"SHA256", "SHA384", "SHA512"}}, "issuerNameHash": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 128, HasMaxLength: true}, "issuerKeyHash": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 128, HasMaxLength: true}, "serialNumber": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 40, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"hashAlgorithm", "issuerNameHash", "issuerKeyHash", "serialNumber"}}, "source": &validation.Schema{Type: "string", Enum: []string{"CRL", "OCSP"}}, "urls": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 2000, HasMaxLength: true}, MinItems: 1, HasMinItems: true, MaxItems: 5, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"source", "urls", "certificateHashData"}}, MinItems: 1, HasMinItems: true, MaxItems: 4, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"certificateStatusRequests"}}

type GetCertificateChainStatusRequest struct {
	CertificateStatusRequests []GetCertificateChainStatusRequestCertificateStatusRequestInfo `json:"certificateStatusRequests"`
	CustomData                *GetCertificateChainStatusRequestCustomData                    `json:"customData,omitempty"`
}

type GetCertificateChainStatusRequestCertificateStatusRequestInfo struct {
	CertificateHashData GetCertificateChainStatusRequestCertificateHashData         `json:"certificateHashData"`
	Source              GetCertificateChainStatusRequestCertificateStatusSourceEnum `json:"source"`
	Urls                []string                                                    `json:"urls"`
	CustomData          *GetCertificateChainStatusRequestCustomData                 `json:"customData,omitempty"`
}

type GetCertificateChainStatusRequestCertificateStatusSourceEnum string

const (
	GetCertificateChainStatusRequestCertificateStatusSourceEnumCRL  GetCertificateChainStatusRequestCertificateStatusSourceEnum = "CRL"
	GetCertificateChainStatusRequestCertificateStatusSourceEnumOCSP GetCertificateChainStatusRequestCertificateStatusSourceEnum = "OCSP"
)

type GetCertificateChainStatusRequestCertificateHashData struct {
	HashAlgorithm  GetCertificateChainStatusRequestHashAlgorithmEnum `json:"hashAlgorithm"`
	IssuerNameHash string                                            `json:"issuerNameHash"`
	IssuerKeyHash  string                                            `json:"issuerKeyHash"`
	SerialNumber   string                                            `json:"serialNumber"`
	CustomData     *GetCertificateChainStatusRequestCustomData       `json:"customData,omitempty"`
}

type GetCertificateChainStatusRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value GetCertificateChainStatusRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *GetCertificateChainStatusRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

type GetCertificateChainStatusRequestHashAlgorithmEnum string

const (
	GetCertificateChainStatusRequestHashAlgorithmEnumSHA256 GetCertificateChainStatusRequestHashAlgorithmEnum = "SHA256"
	GetCertificateChainStatusRequestHashAlgorithmEnumSHA384 GetCertificateChainStatusRequestHashAlgorithmEnum = "SHA384"
	GetCertificateChainStatusRequestHashAlgorithmEnumSHA512 GetCertificateChainStatusRequestHashAlgorithmEnum = "SHA512"
)

func (GetCertificateChainStatusRequest) ActionName() string { return "GetCertificateChainStatus" }

func (GetCertificateChainStatusRequest) Version() protocol.Version { return protocol.OCPP21 }

func (GetCertificateChainStatusRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (GetCertificateChainStatusRequest) SchemaName() string {
	return "GetCertificateChainStatusRequest.json"
}

func (message GetCertificateChainStatusRequest) Validate() error {
	return validation.Validate("GetCertificateChainStatusRequest.json", schemaGetCertificateChainStatusRequest, message)
}

func (GetCertificateChainStatusRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetCertificateChainStatusRequest.json", schemaGetCertificateChainStatusRequest, data)
}
