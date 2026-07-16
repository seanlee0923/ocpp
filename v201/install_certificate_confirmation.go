// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = InstallCertificateConfirmation{}

var schemaInstallCertificateConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "Failed"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}}, Required: []string{"status"}}

type InstallCertificateConfirmation struct {
	CustomData *InstallCertificateConfirmationCustomData                  `json:"customData,omitempty"`
	Status     InstallCertificateConfirmationInstallCertificateStatusEnum `json:"status"`
	StatusInfo *InstallCertificateConfirmationStatusInfo                  `json:"statusInfo,omitempty"`
}

type InstallCertificateConfirmationStatusInfo struct {
	CustomData     *InstallCertificateConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                                    `json:"reasonCode"`
	AdditionalInfo *string                                   `json:"additionalInfo,omitempty"`
}

type InstallCertificateConfirmationInstallCertificateStatusEnum string

const (
	InstallCertificateConfirmationInstallCertificateStatusEnumAccepted InstallCertificateConfirmationInstallCertificateStatusEnum = "Accepted"
	InstallCertificateConfirmationInstallCertificateStatusEnumRejected InstallCertificateConfirmationInstallCertificateStatusEnum = "Rejected"
	InstallCertificateConfirmationInstallCertificateStatusEnumFailed   InstallCertificateConfirmationInstallCertificateStatusEnum = "Failed"
)

type InstallCertificateConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value InstallCertificateConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *InstallCertificateConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (InstallCertificateConfirmation) ActionName() string { return "InstallCertificate" }

func (InstallCertificateConfirmation) Version() protocol.Version { return protocol.OCPP201 }

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
