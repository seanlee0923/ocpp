// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = UpdateFirmwareConfirmation{}

var schemaUpdateFirmwareConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "AcceptedCanceled", "InvalidCertificate", "RevokedCertificate"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}}, Required: []string{"status"}}

type UpdateFirmwareConfirmation struct {
	CustomData *UpdateFirmwareConfirmationCustomData              `json:"customData,omitempty"`
	Status     UpdateFirmwareConfirmationUpdateFirmwareStatusEnum `json:"status"`
	StatusInfo *UpdateFirmwareConfirmationStatusInfo              `json:"statusInfo,omitempty"`
}

type UpdateFirmwareConfirmationStatusInfo struct {
	CustomData     *UpdateFirmwareConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                                `json:"reasonCode"`
	AdditionalInfo *string                               `json:"additionalInfo,omitempty"`
}

type UpdateFirmwareConfirmationUpdateFirmwareStatusEnum string

const (
	UpdateFirmwareConfirmationUpdateFirmwareStatusEnumAccepted           UpdateFirmwareConfirmationUpdateFirmwareStatusEnum = "Accepted"
	UpdateFirmwareConfirmationUpdateFirmwareStatusEnumRejected           UpdateFirmwareConfirmationUpdateFirmwareStatusEnum = "Rejected"
	UpdateFirmwareConfirmationUpdateFirmwareStatusEnumAcceptedCanceled   UpdateFirmwareConfirmationUpdateFirmwareStatusEnum = "AcceptedCanceled"
	UpdateFirmwareConfirmationUpdateFirmwareStatusEnumInvalidCertificate UpdateFirmwareConfirmationUpdateFirmwareStatusEnum = "InvalidCertificate"
	UpdateFirmwareConfirmationUpdateFirmwareStatusEnumRevokedCertificate UpdateFirmwareConfirmationUpdateFirmwareStatusEnum = "RevokedCertificate"
)

type UpdateFirmwareConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (UpdateFirmwareConfirmation) ActionName() string { return "UpdateFirmware" }

func (UpdateFirmwareConfirmation) Version() protocol.Version { return protocol.OCPP201 }

func (UpdateFirmwareConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (UpdateFirmwareConfirmation) SchemaName() string { return "UpdateFirmwareResponse.json" }

func (message UpdateFirmwareConfirmation) Validate() error {
	return validation.Validate("UpdateFirmwareResponse.json", schemaUpdateFirmwareConfirmation, message)
}

func (UpdateFirmwareConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("UpdateFirmwareResponse.json", schemaUpdateFirmwareConfirmation, data)
}
