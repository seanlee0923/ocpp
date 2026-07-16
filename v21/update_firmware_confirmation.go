// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = UpdateFirmwareConfirmation{}

var schemaUpdateFirmwareConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "AcceptedCanceled", "InvalidCertificate", "RevokedCertificate"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type UpdateFirmwareConfirmation struct {
	Status     UpdateFirmwareConfirmationUpdateFirmwareStatusEnum `json:"status"`
	StatusInfo *UpdateFirmwareConfirmationStatusInfo              `json:"statusInfo,omitempty"`
	CustomData *UpdateFirmwareConfirmationCustomData              `json:"customData,omitempty"`
}

type UpdateFirmwareConfirmationStatusInfo struct {
	ReasonCode     string                                `json:"reasonCode"`
	AdditionalInfo *string                               `json:"additionalInfo,omitempty"`
	CustomData     *UpdateFirmwareConfirmationCustomData `json:"customData,omitempty"`
}

type UpdateFirmwareConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

type UpdateFirmwareConfirmationUpdateFirmwareStatusEnum string

const (
	UpdateFirmwareConfirmationUpdateFirmwareStatusEnumAccepted           UpdateFirmwareConfirmationUpdateFirmwareStatusEnum = "Accepted"
	UpdateFirmwareConfirmationUpdateFirmwareStatusEnumRejected           UpdateFirmwareConfirmationUpdateFirmwareStatusEnum = "Rejected"
	UpdateFirmwareConfirmationUpdateFirmwareStatusEnumAcceptedCanceled   UpdateFirmwareConfirmationUpdateFirmwareStatusEnum = "AcceptedCanceled"
	UpdateFirmwareConfirmationUpdateFirmwareStatusEnumInvalidCertificate UpdateFirmwareConfirmationUpdateFirmwareStatusEnum = "InvalidCertificate"
	UpdateFirmwareConfirmationUpdateFirmwareStatusEnumRevokedCertificate UpdateFirmwareConfirmationUpdateFirmwareStatusEnum = "RevokedCertificate"
)

func (UpdateFirmwareConfirmation) ActionName() string { return "UpdateFirmware" }

func (UpdateFirmwareConfirmation) Version() protocol.Version { return protocol.OCPP21 }

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
