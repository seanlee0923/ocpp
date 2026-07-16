// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = InstallCertificateConfirmation{}

var schemaInstallCertificateConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "Failed"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type InstallCertificateConfirmation struct {
	Status     InstallCertificateConfirmationInstallCertificateStatusEnum `json:"status"`
	StatusInfo *InstallCertificateConfirmationStatusInfo                  `json:"statusInfo,omitempty"`
	CustomData *InstallCertificateConfirmationCustomData                  `json:"customData,omitempty"`
}

type InstallCertificateConfirmationStatusInfo struct {
	ReasonCode     string                                    `json:"reasonCode"`
	AdditionalInfo *string                                   `json:"additionalInfo,omitempty"`
	CustomData     *InstallCertificateConfirmationCustomData `json:"customData,omitempty"`
}

type InstallCertificateConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

type InstallCertificateConfirmationInstallCertificateStatusEnum string

const (
	InstallCertificateConfirmationInstallCertificateStatusEnumAccepted InstallCertificateConfirmationInstallCertificateStatusEnum = "Accepted"
	InstallCertificateConfirmationInstallCertificateStatusEnumRejected InstallCertificateConfirmationInstallCertificateStatusEnum = "Rejected"
	InstallCertificateConfirmationInstallCertificateStatusEnumFailed   InstallCertificateConfirmationInstallCertificateStatusEnum = "Failed"
)

func (InstallCertificateConfirmation) ActionName() string { return "InstallCertificate" }

func (InstallCertificateConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (InstallCertificateConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (InstallCertificateConfirmation) SchemaName() string { return "InstallCertificateResponse.json" }

func (message InstallCertificateConfirmation) Validate() error {
	return validation.Validate("InstallCertificateResponse.json", schemaInstallCertificateConfirmation, message)
}

func (InstallCertificateConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("InstallCertificateResponse.json", schemaInstallCertificateConfirmation, data)
}
