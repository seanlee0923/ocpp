// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = NotifyAllowedEnergyTransferConfirmation{}

var schemaNotifyAllowedEnergyTransferConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type NotifyAllowedEnergyTransferConfirmation struct {
	Status     NotifyAllowedEnergyTransferConfirmationNotifyAllowedEnergyTransferStatusEnum `json:"status"`
	StatusInfo *NotifyAllowedEnergyTransferConfirmationStatusInfo                           `json:"statusInfo,omitempty"`
	CustomData *NotifyAllowedEnergyTransferConfirmationCustomData                           `json:"customData,omitempty"`
}

type NotifyAllowedEnergyTransferConfirmationStatusInfo struct {
	ReasonCode     string                                             `json:"reasonCode"`
	AdditionalInfo *string                                            `json:"additionalInfo,omitempty"`
	CustomData     *NotifyAllowedEnergyTransferConfirmationCustomData `json:"customData,omitempty"`
}

type NotifyAllowedEnergyTransferConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

type NotifyAllowedEnergyTransferConfirmationNotifyAllowedEnergyTransferStatusEnum string

const (
	NotifyAllowedEnergyTransferConfirmationNotifyAllowedEnergyTransferStatusEnumAccepted NotifyAllowedEnergyTransferConfirmationNotifyAllowedEnergyTransferStatusEnum = "Accepted"
	NotifyAllowedEnergyTransferConfirmationNotifyAllowedEnergyTransferStatusEnumRejected NotifyAllowedEnergyTransferConfirmationNotifyAllowedEnergyTransferStatusEnum = "Rejected"
)

func (NotifyAllowedEnergyTransferConfirmation) ActionName() string {
	return "NotifyAllowedEnergyTransfer"
}

func (NotifyAllowedEnergyTransferConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (NotifyAllowedEnergyTransferConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (NotifyAllowedEnergyTransferConfirmation) SchemaName() string {
	return "NotifyAllowedEnergyTransferResponse.json"
}

func (message NotifyAllowedEnergyTransferConfirmation) Validate() error {
	return validation.Validate("NotifyAllowedEnergyTransferResponse.json", schemaNotifyAllowedEnergyTransferConfirmation, message)
}

func (NotifyAllowedEnergyTransferConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("NotifyAllowedEnergyTransferResponse.json", schemaNotifyAllowedEnergyTransferConfirmation, data)
}
