// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = DeleteCertificateConfirmation{}

var schemaDeleteCertificateConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Failed", "NotFound"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}}, Required: []string{"status"}}

type DeleteCertificateConfirmation struct {
	CustomData *DeleteCertificateConfirmationCustomData                 `json:"customData,omitempty"`
	Status     DeleteCertificateConfirmationDeleteCertificateStatusEnum `json:"status"`
	StatusInfo *DeleteCertificateConfirmationStatusInfo                 `json:"statusInfo,omitempty"`
}

type DeleteCertificateConfirmationStatusInfo struct {
	CustomData     *DeleteCertificateConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                                   `json:"reasonCode"`
	AdditionalInfo *string                                  `json:"additionalInfo,omitempty"`
}

type DeleteCertificateConfirmationDeleteCertificateStatusEnum string

const (
	DeleteCertificateConfirmationDeleteCertificateStatusEnumAccepted DeleteCertificateConfirmationDeleteCertificateStatusEnum = "Accepted"
	DeleteCertificateConfirmationDeleteCertificateStatusEnumFailed   DeleteCertificateConfirmationDeleteCertificateStatusEnum = "Failed"
	DeleteCertificateConfirmationDeleteCertificateStatusEnumNotFound DeleteCertificateConfirmationDeleteCertificateStatusEnum = "NotFound"
)

type DeleteCertificateConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (DeleteCertificateConfirmation) ActionName() string { return "DeleteCertificate" }

func (DeleteCertificateConfirmation) Version() protocol.Version { return protocol.OCPP201 }

func (DeleteCertificateConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (DeleteCertificateConfirmation) SchemaName() string { return "DeleteCertificateResponse.json" }

func (message DeleteCertificateConfirmation) Validate() error {
	return validation.Validate("DeleteCertificateResponse.json", schemaDeleteCertificateConfirmation, message)
}

func (DeleteCertificateConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("DeleteCertificateResponse.json", schemaDeleteCertificateConfirmation, data)
}
