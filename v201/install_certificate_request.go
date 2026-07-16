// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = InstallCertificateRequest{}

var schemaInstallCertificateRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "certificateType": &validation.Schema{Type: "string", Enum: []string{"V2GRootCertificate", "MORootCertificate", "CSMSRootCertificate", "ManufacturerRootCertificate"}}, "certificate": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 5500, HasMaxLength: true}}, Required: []string{"certificateType", "certificate"}}

type InstallCertificateRequest struct {
	CustomData      *InstallCertificateRequestCustomData               `json:"customData,omitempty"`
	CertificateType InstallCertificateRequestInstallCertificateUseEnum `json:"certificateType"`
	Certificate     string                                             `json:"certificate"`
}

type InstallCertificateRequestInstallCertificateUseEnum string

const (
	InstallCertificateRequestInstallCertificateUseEnumV2GRootCertificate          InstallCertificateRequestInstallCertificateUseEnum = "V2GRootCertificate"
	InstallCertificateRequestInstallCertificateUseEnumMORootCertificate           InstallCertificateRequestInstallCertificateUseEnum = "MORootCertificate"
	InstallCertificateRequestInstallCertificateUseEnumCSMSRootCertificate         InstallCertificateRequestInstallCertificateUseEnum = "CSMSRootCertificate"
	InstallCertificateRequestInstallCertificateUseEnumManufacturerRootCertificate InstallCertificateRequestInstallCertificateUseEnum = "ManufacturerRootCertificate"
)

type InstallCertificateRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (InstallCertificateRequest) ActionName() string { return "InstallCertificate" }

func (InstallCertificateRequest) Version() protocol.Version { return protocol.OCPP201 }

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
