// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = SetDefaultTariffConfirmation{}

var schemaSetDefaultTariffConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "TooManyElements", "ConditionNotSupported", "DuplicateTariffId"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type SetDefaultTariffConfirmation struct {
	Status     SetDefaultTariffConfirmationTariffSetStatusEnum `json:"status"`
	StatusInfo *SetDefaultTariffConfirmationStatusInfo         `json:"statusInfo,omitempty"`
	CustomData *SetDefaultTariffConfirmationCustomData         `json:"customData,omitempty"`
}

type SetDefaultTariffConfirmationStatusInfo struct {
	ReasonCode     string                                  `json:"reasonCode"`
	AdditionalInfo *string                                 `json:"additionalInfo,omitempty"`
	CustomData     *SetDefaultTariffConfirmationCustomData `json:"customData,omitempty"`
}

type SetDefaultTariffConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

type SetDefaultTariffConfirmationTariffSetStatusEnum string

const (
	SetDefaultTariffConfirmationTariffSetStatusEnumAccepted              SetDefaultTariffConfirmationTariffSetStatusEnum = "Accepted"
	SetDefaultTariffConfirmationTariffSetStatusEnumRejected              SetDefaultTariffConfirmationTariffSetStatusEnum = "Rejected"
	SetDefaultTariffConfirmationTariffSetStatusEnumTooManyElements       SetDefaultTariffConfirmationTariffSetStatusEnum = "TooManyElements"
	SetDefaultTariffConfirmationTariffSetStatusEnumConditionNotSupported SetDefaultTariffConfirmationTariffSetStatusEnum = "ConditionNotSupported"
	SetDefaultTariffConfirmationTariffSetStatusEnumDuplicateTariffID     SetDefaultTariffConfirmationTariffSetStatusEnum = "DuplicateTariffId"
)

func (SetDefaultTariffConfirmation) ActionName() string { return "SetDefaultTariff" }

func (SetDefaultTariffConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (SetDefaultTariffConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (SetDefaultTariffConfirmation) SchemaName() string { return "SetDefaultTariffResponse.json" }

func (message SetDefaultTariffConfirmation) Validate() error {
	return validation.Validate("SetDefaultTariffResponse.json", schemaSetDefaultTariffConfirmation, message)
}

func (SetDefaultTariffConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("SetDefaultTariffResponse.json", schemaSetDefaultTariffConfirmation, data)
}
