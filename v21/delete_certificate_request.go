// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = DeleteCertificateRequest{}

var schemaDeleteCertificateRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"certificateHashData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"hashAlgorithm": &validation.Schema{Type: "string", Enum: []string{"SHA256", "SHA384", "SHA512"}}, "issuerNameHash": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 128, HasMaxLength: true}, "issuerKeyHash": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 128, HasMaxLength: true}, "serialNumber": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 40, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"hashAlgorithm", "issuerNameHash", "issuerKeyHash", "serialNumber"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"certificateHashData"}}

type DeleteCertificateRequest struct {
	CertificateHashData DeleteCertificateRequestCertificateHashData `json:"certificateHashData"`
	CustomData          *DeleteCertificateRequestCustomData         `json:"customData,omitempty"`
}

type DeleteCertificateRequestCertificateHashData struct {
	HashAlgorithm  DeleteCertificateRequestHashAlgorithmEnum `json:"hashAlgorithm"`
	IssuerNameHash string                                    `json:"issuerNameHash"`
	IssuerKeyHash  string                                    `json:"issuerKeyHash"`
	SerialNumber   string                                    `json:"serialNumber"`
	CustomData     *DeleteCertificateRequestCustomData       `json:"customData,omitempty"`
}

type DeleteCertificateRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value DeleteCertificateRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *DeleteCertificateRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

type DeleteCertificateRequestHashAlgorithmEnum string

const (
	DeleteCertificateRequestHashAlgorithmEnumSHA256 DeleteCertificateRequestHashAlgorithmEnum = "SHA256"
	DeleteCertificateRequestHashAlgorithmEnumSHA384 DeleteCertificateRequestHashAlgorithmEnum = "SHA384"
	DeleteCertificateRequestHashAlgorithmEnumSHA512 DeleteCertificateRequestHashAlgorithmEnum = "SHA512"
)

func (DeleteCertificateRequest) ActionName() string { return "DeleteCertificate" }

func (DeleteCertificateRequest) Version() protocol.Version { return protocol.OCPP21 }

func (DeleteCertificateRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (DeleteCertificateRequest) SchemaName() string { return "DeleteCertificateRequest.json" }

func (message DeleteCertificateRequest) Validate() error {
	return validation.Validate("DeleteCertificateRequest.json", schemaDeleteCertificateRequest, message)
}

func (DeleteCertificateRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("DeleteCertificateRequest.json", schemaDeleteCertificateRequest, data)
}
