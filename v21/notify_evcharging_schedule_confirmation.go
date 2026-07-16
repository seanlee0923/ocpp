// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = NotifyEVChargingScheduleConfirmation{}

var schemaNotifyEVChargingScheduleConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type NotifyEVChargingScheduleConfirmation struct {
	Status     NotifyEVChargingScheduleConfirmationGenericStatusEnum `json:"status"`
	StatusInfo *NotifyEVChargingScheduleConfirmationStatusInfo       `json:"statusInfo,omitempty"`
	CustomData *NotifyEVChargingScheduleConfirmationCustomData       `json:"customData,omitempty"`
}

type NotifyEVChargingScheduleConfirmationStatusInfo struct {
	ReasonCode     string                                          `json:"reasonCode"`
	AdditionalInfo *string                                         `json:"additionalInfo,omitempty"`
	CustomData     *NotifyEVChargingScheduleConfirmationCustomData `json:"customData,omitempty"`
}

type NotifyEVChargingScheduleConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

type NotifyEVChargingScheduleConfirmationGenericStatusEnum string

const (
	NotifyEVChargingScheduleConfirmationGenericStatusEnumAccepted NotifyEVChargingScheduleConfirmationGenericStatusEnum = "Accepted"
	NotifyEVChargingScheduleConfirmationGenericStatusEnumRejected NotifyEVChargingScheduleConfirmationGenericStatusEnum = "Rejected"
)

func (NotifyEVChargingScheduleConfirmation) ActionName() string { return "NotifyEVChargingSchedule" }

func (NotifyEVChargingScheduleConfirmation) Version() protocol.Version { return protocol.OCPP21 }

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
