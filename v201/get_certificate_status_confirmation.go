// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = GetCertificateStatusConfirmation{}

var schemaGetCertificateStatusConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Failed"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}, "ocspResult": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 5500, HasMaxLength: true}}, Required: []string{"status"}}

type GetCertificateStatusConfirmation struct {
	CustomData *GetCertificateStatusConfirmationCustomData              `json:"customData,omitempty"`
	Status     GetCertificateStatusConfirmationGetCertificateStatusEnum `json:"status"`
	StatusInfo *GetCertificateStatusConfirmationStatusInfo              `json:"statusInfo,omitempty"`
	OcspResult *string                                                  `json:"ocspResult,omitempty"`
}

type GetCertificateStatusConfirmationStatusInfo struct {
	CustomData     *GetCertificateStatusConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                                      `json:"reasonCode"`
	AdditionalInfo *string                                     `json:"additionalInfo,omitempty"`
}

type GetCertificateStatusConfirmationGetCertificateStatusEnum string

const (
	GetCertificateStatusConfirmationGetCertificateStatusEnumAccepted GetCertificateStatusConfirmationGetCertificateStatusEnum = "Accepted"
	GetCertificateStatusConfirmationGetCertificateStatusEnumFailed   GetCertificateStatusConfirmationGetCertificateStatusEnum = "Failed"
)

type GetCertificateStatusConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (GetCertificateStatusConfirmation) ActionName() string { return "GetCertificateStatus" }

func (GetCertificateStatusConfirmation) Version() protocol.Version { return protocol.OCPP201 }

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
