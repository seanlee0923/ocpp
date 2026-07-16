// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = GetInstalledCertificateIdsRequest{}

var schemaGetInstalledCertificateIdsRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"certificateType": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", Enum: []string{"V2GRootCertificate", "MORootCertificate", "CSMSRootCertificate", "V2GCertificateChain", "ManufacturerRootCertificate", "OEMRootCertificate"}}, MinItems: 1, HasMinItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type GetInstalledCertificateIdsRequest struct {
	CertificateType []GetInstalledCertificateIdsRequestGetCertificateIDUseEnum `json:"certificateType,omitempty"`
	CustomData      *GetInstalledCertificateIdsRequestCustomData               `json:"customData,omitempty"`
}

type GetInstalledCertificateIdsRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

type GetInstalledCertificateIdsRequestGetCertificateIDUseEnum string

const (
	GetInstalledCertificateIdsRequestGetCertificateIDUseEnumV2GRootCertificate          GetInstalledCertificateIdsRequestGetCertificateIDUseEnum = "V2GRootCertificate"
	GetInstalledCertificateIdsRequestGetCertificateIDUseEnumMORootCertificate           GetInstalledCertificateIdsRequestGetCertificateIDUseEnum = "MORootCertificate"
	GetInstalledCertificateIdsRequestGetCertificateIDUseEnumCSMSRootCertificate         GetInstalledCertificateIdsRequestGetCertificateIDUseEnum = "CSMSRootCertificate"
	GetInstalledCertificateIdsRequestGetCertificateIDUseEnumV2GCertificateChain         GetInstalledCertificateIdsRequestGetCertificateIDUseEnum = "V2GCertificateChain"
	GetInstalledCertificateIdsRequestGetCertificateIDUseEnumManufacturerRootCertificate GetInstalledCertificateIdsRequestGetCertificateIDUseEnum = "ManufacturerRootCertificate"
	GetInstalledCertificateIdsRequestGetCertificateIDUseEnumOEMRootCertificate          GetInstalledCertificateIdsRequestGetCertificateIDUseEnum = "OEMRootCertificate"
)

func (GetInstalledCertificateIdsRequest) ActionName() string { return "GetInstalledCertificateIds" }

func (GetInstalledCertificateIdsRequest) Version() protocol.Version { return protocol.OCPP21 }

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
