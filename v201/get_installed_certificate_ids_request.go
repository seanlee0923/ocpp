// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = GetInstalledCertificateIdsRequest{}

var schemaGetInstalledCertificateIdsRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "certificateType": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", Enum: []string{"V2GRootCertificate", "MORootCertificate", "CSMSRootCertificate", "V2GCertificateChain", "ManufacturerRootCertificate"}}, MinItems: 1, HasMinItems: true}}}

type GetInstalledCertificateIdsRequest struct {
	CustomData      *GetInstalledCertificateIdsRequestCustomData               `json:"customData,omitempty"`
	CertificateType []GetInstalledCertificateIdsRequestGetCertificateIDUseEnum `json:"certificateType,omitempty"`
}

type GetInstalledCertificateIdsRequestGetCertificateIDUseEnum string

const (
	GetInstalledCertificateIdsRequestGetCertificateIDUseEnumV2GRootCertificate          GetInstalledCertificateIdsRequestGetCertificateIDUseEnum = "V2GRootCertificate"
	GetInstalledCertificateIdsRequestGetCertificateIDUseEnumMORootCertificate           GetInstalledCertificateIdsRequestGetCertificateIDUseEnum = "MORootCertificate"
	GetInstalledCertificateIdsRequestGetCertificateIDUseEnumCSMSRootCertificate         GetInstalledCertificateIdsRequestGetCertificateIDUseEnum = "CSMSRootCertificate"
	GetInstalledCertificateIdsRequestGetCertificateIDUseEnumV2GCertificateChain         GetInstalledCertificateIdsRequestGetCertificateIDUseEnum = "V2GCertificateChain"
	GetInstalledCertificateIdsRequestGetCertificateIDUseEnumManufacturerRootCertificate GetInstalledCertificateIdsRequestGetCertificateIDUseEnum = "ManufacturerRootCertificate"
)

type GetInstalledCertificateIdsRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (GetInstalledCertificateIdsRequest) ActionName() string { return "GetInstalledCertificateIds" }

func (GetInstalledCertificateIdsRequest) Version() protocol.Version { return protocol.OCPP201 }

func (GetInstalledCertificateIdsRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (GetInstalledCertificateIdsRequest) SchemaName() string {
	return "GetInstalledCertificateIdsRequest.json"
}

func (message GetInstalledCertificateIdsRequest) Validate() error {
	return validation.Validate("GetInstalledCertificateIdsRequest.json", schemaGetInstalledCertificateIdsRequest, message)
}

func (GetInstalledCertificateIdsRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetInstalledCertificateIdsRequest.json", schemaGetInstalledCertificateIdsRequest, data)
}
