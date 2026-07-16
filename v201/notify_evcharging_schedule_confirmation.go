// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = NotifyEVChargingScheduleConfirmation{}

var schemaNotifyEVChargingScheduleConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}}, Required: []string{"status"}}

type NotifyEVChargingScheduleConfirmation struct {
	CustomData *NotifyEVChargingScheduleConfirmationCustomData       `json:"customData,omitempty"`
	Status     NotifyEVChargingScheduleConfirmationGenericStatusEnum `json:"status"`
	StatusInfo *NotifyEVChargingScheduleConfirmationStatusInfo       `json:"statusInfo,omitempty"`
}

type NotifyEVChargingScheduleConfirmationStatusInfo struct {
	CustomData     *NotifyEVChargingScheduleConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                                          `json:"reasonCode"`
	AdditionalInfo *string                                         `json:"additionalInfo,omitempty"`
}

type NotifyEVChargingScheduleConfirmationGenericStatusEnum string

const (
	NotifyEVChargingScheduleConfirmationGenericStatusEnumAccepted NotifyEVChargingScheduleConfirmationGenericStatusEnum = "Accepted"
	NotifyEVChargingScheduleConfirmationGenericStatusEnumRejected NotifyEVChargingScheduleConfirmationGenericStatusEnum = "Rejected"
)

type NotifyEVChargingScheduleConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (NotifyEVChargingScheduleConfirmation) ActionName() string { return "NotifyEVChargingSchedule" }

func (NotifyEVChargingScheduleConfirmation) Version() protocol.Version { return protocol.OCPP201 }

func (NotifyEVChargingScheduleConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (NotifyEVChargingScheduleConfirmation) SchemaName() string {
	return "NotifyEVChargingScheduleResponse.json"
}

func (message NotifyEVChargingScheduleConfirmation) Validate() error {
	return validation.Validate("NotifyEVChargingScheduleResponse.json", schemaNotifyEVChargingScheduleConfirmation, message)
}

func (NotifyEVChargingScheduleConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("NotifyEVChargingScheduleResponse.json", schemaNotifyEVChargingScheduleConfirmation, data)
}
