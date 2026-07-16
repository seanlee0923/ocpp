// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = UpdateDynamicScheduleConfirmation{}

var schemaUpdateDynamicScheduleConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type UpdateDynamicScheduleConfirmation struct {
	Status     UpdateDynamicScheduleConfirmationChargingProfileStatusEnum `json:"status"`
	StatusInfo *UpdateDynamicScheduleConfirmationStatusInfo               `json:"statusInfo,omitempty"`
	CustomData *UpdateDynamicScheduleConfirmationCustomData               `json:"customData,omitempty"`
}

type UpdateDynamicScheduleConfirmationStatusInfo struct {
	ReasonCode     string                                       `json:"reasonCode"`
	AdditionalInfo *string                                      `json:"additionalInfo,omitempty"`
	CustomData     *UpdateDynamicScheduleConfirmationCustomData `json:"customData,omitempty"`
}

type UpdateDynamicScheduleConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

type UpdateDynamicScheduleConfirmationChargingProfileStatusEnum string

const (
	UpdateDynamicScheduleConfirmationChargingProfileStatusEnumAccepted UpdateDynamicScheduleConfirmationChargingProfileStatusEnum = "Accepted"
	UpdateDynamicScheduleConfirmationChargingProfileStatusEnumRejected UpdateDynamicScheduleConfirmationChargingProfileStatusEnum = "Rejected"
)

func (UpdateDynamicScheduleConfirmation) ActionName() string { return "UpdateDynamicSchedule" }

func (UpdateDynamicScheduleConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (UpdateDynamicScheduleConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (UpdateDynamicScheduleConfirmation) SchemaName() string {
	return "UpdateDynamicScheduleResponse.json"
}

func (message UpdateDynamicScheduleConfirmation) Validate() error {
	return validation.Validate("UpdateDynamicScheduleResponse.json", schemaUpdateDynamicScheduleConfirmation, message)
}

func (UpdateDynamicScheduleConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("UpdateDynamicScheduleResponse.json", schemaUpdateDynamicScheduleConfirmation, data)
}
