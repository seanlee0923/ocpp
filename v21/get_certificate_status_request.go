// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = GetCertificateStatusRequest{}

var schemaGetCertificateStatusRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"ocspRequestData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"hashAlgorithm": &validation.Schema{Type: "string", Enum: []string{"SHA256", "SHA384", "SHA512"}}, "issuerNameHash": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 128, HasMaxLength: true}, "issuerKeyHash": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 128, HasMaxLength: true}, "serialNumber": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 40, HasMaxLength: true}, "responderURL": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 2000, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"hashAlgorithm", "issuerNameHash", "issuerKeyHash", "serialNumber", "responderURL"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"ocspRequestData"}}

type GetCertificateStatusRequest struct {
	OcspRequestData GetCertificateStatusRequestOCSPRequestData `json:"ocspRequestData"`
	CustomData      *GetCertificateStatusRequestCustomData     `json:"customData,omitempty"`
}

type GetCertificateStatusRequestOCSPRequestData struct {
	HashAlgorithm  GetCertificateStatusRequestHashAlgorithmEnum `json:"hashAlgorithm"`
	IssuerNameHash string                                       `json:"issuerNameHash"`
	IssuerKeyHash  string                                       `json:"issuerKeyHash"`
	SerialNumber   string                                       `json:"serialNumber"`
	ResponderURL   string                                       `json:"responderURL"`
	CustomData     *GetCertificateStatusRequestCustomData       `json:"customData,omitempty"`
}

type GetCertificateStatusRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

type GetCertificateStatusRequestHashAlgorithmEnum string

const (
	GetCertificateStatusRequestHashAlgorithmEnumSHA256 GetCertificateStatusRequestHashAlgorithmEnum = "SHA256"
	GetCertificateStatusRequestHashAlgorithmEnumSHA384 GetCertificateStatusRequestHashAlgorithmEnum = "SHA384"
	GetCertificateStatusRequestHashAlgorithmEnumSHA512 GetCertificateStatusRequestHashAlgorithmEnum = "SHA512"
)

func (GetCertificateStatusRequest) ActionName() string { return "GetCertificateStatus" }

func (GetCertificateStatusRequest) Version() protocol.Version { return protocol.OCPP21 }

func (GetCertificateStatusRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (GetCertificateStatusRequest) SchemaName() string { return "GetCertificateStatusRequest.json" }

func (message GetCertificateStatusRequest) Validate() error {
	return validation.Validate("GetCertificateStatusRequest.json", schemaGetCertificateStatusRequest, message)
}

func (GetCertificateStatusRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetCertificateStatusRequest.json", schemaGetCertificateStatusRequest, data)
}
