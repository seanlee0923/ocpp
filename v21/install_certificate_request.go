// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = InstallCertificateRequest{}

var schemaInstallCertificateRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"certificateType": &validation.Schema{Type: "string", Enum: []string{"V2GRootCertificate", "MORootCertificate", "ManufacturerRootCertificate", "CSMSRootCertificate", "OEMRootCertificate"}}, "certificate": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 10000, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"certificateType", "certificate"}}

type InstallCertificateRequest struct {
	CertificateType InstallCertificateRequestInstallCertificateUseEnum `json:"certificateType"`
	Certificate     string                                             `json:"certificate"`
	CustomData      *InstallCertificateRequestCustomData               `json:"customData,omitempty"`
}

type InstallCertificateRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

type InstallCertificateRequestInstallCertificateUseEnum string

const (
	InstallCertificateRequestInstallCertificateUseEnumV2GRootCertificate          InstallCertificateRequestInstallCertificateUseEnum = "V2GRootCertificate"
	InstallCertificateRequestInstallCertificateUseEnumMORootCertificate           InstallCertificateRequestInstallCertificateUseEnum = "MORootCertificate"
	InstallCertificateRequestInstallCertificateUseEnumManufacturerRootCertificate InstallCertificateRequestInstallCertificateUseEnum = "ManufacturerRootCertificate"
	InstallCertificateRequestInstallCertificateUseEnumCSMSRootCertificate         InstallCertificateRequestInstallCertificateUseEnum = "CSMSRootCertificate"
	InstallCertificateRequestInstallCertificateUseEnumOEMRootCertificate          InstallCertificateRequestInstallCertificateUseEnum = "OEMRootCertificate"
)

func (InstallCertificateRequest) ActionName() string { return "InstallCertificate" }

func (InstallCertificateRequest) Version() protocol.Version { return protocol.OCPP21 }

func (InstallCertificateRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (InstallCertificateRequest) SchemaName() string { return "InstallCertificateRequest.json" }

func (message InstallCertificateRequest) Validate() error {
	return validation.Validate("InstallCertificateRequest.json", schemaInstallCertificateRequest, message)
}

func (InstallCertificateRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("InstallCertificateRequest.json", schemaInstallCertificateRequest, data)
}
