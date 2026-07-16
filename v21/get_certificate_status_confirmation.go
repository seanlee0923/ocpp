// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = GetCertificateStatusConfirmation{}

var schemaGetCertificateStatusConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Failed"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "ocspResult": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 18000, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type GetCertificateStatusConfirmation struct {
	Status     GetCertificateStatusConfirmationGetCertificateStatusEnum `json:"status"`
	StatusInfo *GetCertificateStatusConfirmationStatusInfo              `json:"statusInfo,omitempty"`
	OcspResult *string                                                  `json:"ocspResult,omitempty"`
	CustomData *GetCertificateStatusConfirmationCustomData              `json:"customData,omitempty"`
}

type GetCertificateStatusConfirmationStatusInfo struct {
	ReasonCode     string                                      `json:"reasonCode"`
	AdditionalInfo *string                                     `json:"additionalInfo,omitempty"`
	CustomData     *GetCertificateStatusConfirmationCustomData `json:"customData,omitempty"`
}

type GetCertificateStatusConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

type GetCertificateStatusConfirmationGetCertificateStatusEnum string

const (
	GetCertificateStatusConfirmationGetCertificateStatusEnumAccepted GetCertificateStatusConfirmationGetCertificateStatusEnum = "Accepted"
	GetCertificateStatusConfirmationGetCertificateStatusEnumFailed   GetCertificateStatusConfirmationGetCertificateStatusEnum = "Failed"
)

func (GetCertificateStatusConfirmation) ActionName() string { return "GetCertificateStatus" }

func (GetCertificateStatusConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (GetCertificateStatusConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (GetCertificateStatusConfirmation) SchemaName() string {
	return "GetCertificateStatusResponse.json"
}

func (message GetCertificateStatusConfirmation) Validate() error {
	return validation.Validate("GetCertificateStatusResponse.json", schemaGetCertificateStatusConfirmation, message)
}

func (GetCertificateStatusConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetCertificateStatusResponse.json", schemaGetCertificateStatusConfirmation, data)
}
