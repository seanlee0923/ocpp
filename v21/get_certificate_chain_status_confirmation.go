// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = GetCertificateChainStatusConfirmation{}

var schemaGetCertificateChainStatusConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"certificateStatus": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"certificateHashData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"hashAlgorithm": &validation.Schema{Type: "string", Enum: []string{"SHA256", "SHA384", "SHA512"}}, "issuerNameHash": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 128, HasMaxLength: true}, "issuerKeyHash": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 128, HasMaxLength: true}, "serialNumber": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 40, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"hashAlgorithm", "issuerNameHash", "issuerKeyHash", "serialNumber"}}, "source": &validation.Schema{Type: "string", Enum: []string{"CRL", "OCSP"}}, "status": &validation.Schema{Type: "string", Enum: []string{"Good", "Revoked", "Unknown", "Failed"}}, "nextUpdate": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"source", "status", "nextUpdate", "certificateHashData"}}, MinItems: 1, HasMinItems: true, MaxItems: 4, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"certificateStatus"}}

type GetCertificateChainStatusConfirmation struct {
	CertificateStatus []GetCertificateChainStatusConfirmationCertificateStatus `json:"certificateStatus"`
	CustomData        *GetCertificateChainStatusConfirmationCustomData         `json:"customData,omitempty"`
}

type GetCertificateChainStatusConfirmationCertificateStatus struct {
	CertificateHashData GetCertificateChainStatusConfirmationCertificateHashData         `json:"certificateHashData"`
	Source              GetCertificateChainStatusConfirmationCertificateStatusSourceEnum `json:"source"`
	Status              GetCertificateChainStatusConfirmationCertificateStatusEnum       `json:"status"`
	NextUpdate          string                                                           `json:"nextUpdate"`
	CustomData          *GetCertificateChainStatusConfirmationCustomData                 `json:"customData,omitempty"`
}

type GetCertificateChainStatusConfirmationCertificateStatusEnum string

const (
	GetCertificateChainStatusConfirmationCertificateStatusEnumGood    GetCertificateChainStatusConfirmationCertificateStatusEnum = "Good"
	GetCertificateChainStatusConfirmationCertificateStatusEnumRevoked GetCertificateChainStatusConfirmationCertificateStatusEnum = "Revoked"
	GetCertificateChainStatusConfirmationCertificateStatusEnumUnknown GetCertificateChainStatusConfirmationCertificateStatusEnum = "Unknown"
	GetCertificateChainStatusConfirmationCertificateStatusEnumFailed  GetCertificateChainStatusConfirmationCertificateStatusEnum = "Failed"
)

type GetCertificateChainStatusConfirmationCertificateStatusSourceEnum string

const (
	GetCertificateChainStatusConfirmationCertificateStatusSourceEnumCRL  GetCertificateChainStatusConfirmationCertificateStatusSourceEnum = "CRL"
	GetCertificateChainStatusConfirmationCertificateStatusSourceEnumOCSP GetCertificateChainStatusConfirmationCertificateStatusSourceEnum = "OCSP"
)

type GetCertificateChainStatusConfirmationCertificateHashData struct {
	HashAlgorithm  GetCertificateChainStatusConfirmationHashAlgorithmEnum `json:"hashAlgorithm"`
	IssuerNameHash string                                                 `json:"issuerNameHash"`
	IssuerKeyHash  string                                                 `json:"issuerKeyHash"`
	SerialNumber   string                                                 `json:"serialNumber"`
	CustomData     *GetCertificateChainStatusConfirmationCustomData       `json:"customData,omitempty"`
}

type GetCertificateChainStatusConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

type GetCertificateChainStatusConfirmationHashAlgorithmEnum string

const (
	GetCertificateChainStatusConfirmationHashAlgorithmEnumSHA256 GetCertificateChainStatusConfirmationHashAlgorithmEnum = "SHA256"
	GetCertificateChainStatusConfirmationHashAlgorithmEnumSHA384 GetCertificateChainStatusConfirmationHashAlgorithmEnum = "SHA384"
	GetCertificateChainStatusConfirmationHashAlgorithmEnumSHA512 GetCertificateChainStatusConfirmationHashAlgorithmEnum = "SHA512"
)

func (GetCertificateChainStatusConfirmation) ActionName() string { return "GetCertificateChainStatus" }

func (GetCertificateChainStatusConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (GetCertificateChainStatusConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (GetCertificateChainStatusConfirmation) SchemaName() string {
	return "GetCertificateChainStatusResponse.json"
}

func (message GetCertificateChainStatusConfirmation) Validate() error {
	return validation.Validate("GetCertificateChainStatusResponse.json", schemaGetCertificateChainStatusConfirmation, message)
}

func (GetCertificateChainStatusConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetCertificateChainStatusResponse.json", schemaGetCertificateChainStatusConfirmation, data)
}
