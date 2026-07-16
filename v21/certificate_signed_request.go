// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = CertificateSignedRequest{}

var schemaCertificateSignedRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"certificateChain": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 10000, HasMaxLength: true}, "certificateType": &validation.Schema{Type: "string", Enum: []string{"ChargingStationCertificate", "V2GCertificate", "V2G20Certificate"}}, "requestId": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"certificateChain"}}

type CertificateSignedRequest struct {
	CertificateChain string                                             `json:"certificateChain"`
	CertificateType  *CertificateSignedRequestCertificateSigningUseEnum `json:"certificateType,omitempty"`
	RequestID        *int                                               `json:"requestId,omitempty"`
	CustomData       *CertificateSignedRequestCustomData                `json:"customData,omitempty"`
}

type CertificateSignedRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value CertificateSignedRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *CertificateSignedRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

type CertificateSignedRequestCertificateSigningUseEnum string

const (
	CertificateSignedRequestCertificateSigningUseEnumChargingStationCertificate CertificateSignedRequestCertificateSigningUseEnum = "ChargingStationCertificate"
	CertificateSignedRequestCertificateSigningUseEnumV2GCertificate             CertificateSignedRequestCertificateSigningUseEnum = "V2GCertificate"
	CertificateSignedRequestCertificateSigningUseEnumV2G20Certificate           CertificateSignedRequestCertificateSigningUseEnum = "V2G20Certificate"
)

func (CertificateSignedRequest) ActionName() string { return "CertificateSigned" }

func (CertificateSignedRequest) Version() protocol.Version { return protocol.OCPP21 }

func (CertificateSignedRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (CertificateSignedRequest) SchemaName() string { return "CertificateSignedRequest.json" }

func (message CertificateSignedRequest) Validate() error {
	return validation.Validate("CertificateSignedRequest.json", schemaCertificateSignedRequest, message)
}

func (CertificateSignedRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("CertificateSignedRequest.json", schemaCertificateSignedRequest, data)
}
