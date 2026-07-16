// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = ClearTariffsConfirmation{}

var schemaClearTariffsConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"clearTariffsResult": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "tariffId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 60, HasMaxLength: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "NoTariff"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}, MinItems: 1, HasMinItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"clearTariffsResult"}}

type ClearTariffsConfirmation struct {
	ClearTariffsResult []ClearTariffsConfirmationClearTariffsResult `json:"clearTariffsResult"`
	CustomData         *ClearTariffsConfirmationCustomData          `json:"customData,omitempty"`
}

type ClearTariffsConfirmationClearTariffsResult struct {
	StatusInfo *ClearTariffsConfirmationStatusInfo           `json:"statusInfo,omitempty"`
	TariffID   *string                                       `json:"tariffId,omitempty"`
	Status     ClearTariffsConfirmationTariffClearStatusEnum `json:"status"`
	CustomData *ClearTariffsConfirmationCustomData           `json:"customData,omitempty"`
}

type ClearTariffsConfirmationTariffClearStatusEnum string

const (
	ClearTariffsConfirmationTariffClearStatusEnumAccepted ClearTariffsConfirmationTariffClearStatusEnum = "Accepted"
	ClearTariffsConfirmationTariffClearStatusEnumRejected ClearTariffsConfirmationTariffClearStatusEnum = "Rejected"
	ClearTariffsConfirmationTariffClearStatusEnumNoTariff ClearTariffsConfirmationTariffClearStatusEnum = "NoTariff"
)

type ClearTariffsConfirmationStatusInfo struct {
	ReasonCode     string                              `json:"reasonCode"`
	AdditionalInfo *string                             `json:"additionalInfo,omitempty"`
	CustomData     *ClearTariffsConfirmationCustomData `json:"customData,omitempty"`
}

type ClearTariffsConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (ClearTariffsConfirmation) ActionName() string { return "ClearTariffs" }

func (ClearTariffsConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (ClearTariffsConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (ClearTariffsConfirmation) SchemaName() string { return "ClearTariffsResponse.json" }

func (message ClearTariffsConfirmation) Validate() error {
	return validation.Validate("ClearTariffsResponse.json", schemaClearTariffsConfirmation, message)
}

func (ClearTariffsConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ClearTariffsResponse.json", schemaClearTariffsConfirmation, data)
}
