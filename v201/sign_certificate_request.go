// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = SignCertificateRequest{}

var schemaSignCertificateRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "csr": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 5500, HasMaxLength: true}, "certificateType": &validation.Schema{Type: "string", Enum: []string{"ChargingStationCertificate", "V2GCertificate"}}}, Required: []string{"csr"}}

type SignCertificateRequest struct {
	CustomData      *SignCertificateRequestCustomData                `json:"customData,omitempty"`
	Csr             string                                           `json:"csr"`
	CertificateType *SignCertificateRequestCertificateSigningUseEnum `json:"certificateType,omitempty"`
}

type SignCertificateRequestCertificateSigningUseEnum string

const (
	SignCertificateRequestCertificateSigningUseEnumChargingStationCertificate SignCertificateRequestCertificateSigningUseEnum = "ChargingStationCertificate"
	SignCertificateRequestCertificateSigningUseEnumV2GCertificate             SignCertificateRequestCertificateSigningUseEnum = "V2GCertificate"
)

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

func (SignCertificateRequest) ActionName() string { return "SignCertificate" }

func (SignCertificateRequest) Version() protocol.Version { return protocol.OCPP201 }

func (SignCertificateRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (SignCertificateRequest) SchemaName() string { return "SignCertificateRequest.json" }

func (message SignCertificateRequest) Validate() error {
	return validation.Validate("SignCertificateRequest.json", schemaSignCertificateRequest, message)
}

func (SignCertificateRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("SignCertificateRequest.json", schemaSignCertificateRequest, data)
}
