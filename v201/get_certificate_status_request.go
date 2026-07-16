// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = GetCertificateStatusRequest{}

var schemaGetCertificateStatusRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "ocspRequestData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "hashAlgorithm": &validation.Schema{Type: "string", Enum: []string{"SHA256", "SHA384", "SHA512"}}, "issuerNameHash": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 128, HasMaxLength: true}, "issuerKeyHash": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 128, HasMaxLength: true}, "serialNumber": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 40, HasMaxLength: true}, "responderURL": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"hashAlgorithm", "issuerNameHash", "issuerKeyHash", "serialNumber", "responderURL"}}}, Required: []string{"ocspRequestData"}}

type GetCertificateStatusRequest struct {
	CustomData      *GetCertificateStatusRequestCustomData     `json:"customData,omitempty"`
	OcspRequestData GetCertificateStatusRequestOCSPRequestData `json:"ocspRequestData"`
}

type GetCertificateStatusRequestOCSPRequestData struct {
	CustomData     *GetCertificateStatusRequestCustomData       `json:"customData,omitempty"`
	HashAlgorithm  GetCertificateStatusRequestHashAlgorithmEnum `json:"hashAlgorithm"`
	IssuerNameHash string                                       `json:"issuerNameHash"`
	IssuerKeyHash  string                                       `json:"issuerKeyHash"`
	SerialNumber   string                                       `json:"serialNumber"`
	ResponderURL   string                                       `json:"responderURL"`
}

type GetCertificateStatusRequestHashAlgorithmEnum string

const (
	GetCertificateStatusRequestHashAlgorithmEnumSHA256 GetCertificateStatusRequestHashAlgorithmEnum = "SHA256"
	GetCertificateStatusRequestHashAlgorithmEnumSHA384 GetCertificateStatusRequestHashAlgorithmEnum = "SHA384"
	GetCertificateStatusRequestHashAlgorithmEnumSHA512 GetCertificateStatusRequestHashAlgorithmEnum = "SHA512"
)

type GetCertificateStatusRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value GetCertificateStatusRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *GetCertificateStatusRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (GetCertificateStatusRequest) ActionName() string { return "GetCertificateStatus" }

func (GetCertificateStatusRequest) Version() protocol.Version { return protocol.OCPP201 }

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
