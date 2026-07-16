// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = SetMonitoringBaseConfirmation{}

var schemaSetMonitoringBaseConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "NotSupported", "EmptyResultSet"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type SetMonitoringBaseConfirmation struct {
	Status     SetMonitoringBaseConfirmationGenericDeviceModelStatusEnum `json:"status"`
	StatusInfo *SetMonitoringBaseConfirmationStatusInfo                  `json:"statusInfo,omitempty"`
	CustomData *SetMonitoringBaseConfirmationCustomData                  `json:"customData,omitempty"`
}

type SetMonitoringBaseConfirmationStatusInfo struct {
	ReasonCode     string                                   `json:"reasonCode"`
	AdditionalInfo *string                                  `json:"additionalInfo,omitempty"`
	CustomData     *SetMonitoringBaseConfirmationCustomData `json:"customData,omitempty"`
}

type SetMonitoringBaseConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

type SetMonitoringBaseConfirmationGenericDeviceModelStatusEnum string

const (
	SetMonitoringBaseConfirmationGenericDeviceModelStatusEnumAccepted       SetMonitoringBaseConfirmationGenericDeviceModelStatusEnum = "Accepted"
	SetMonitoringBaseConfirmationGenericDeviceModelStatusEnumRejected       SetMonitoringBaseConfirmationGenericDeviceModelStatusEnum = "Rejected"
	SetMonitoringBaseConfirmationGenericDeviceModelStatusEnumNotSupported   SetMonitoringBaseConfirmationGenericDeviceModelStatusEnum = "NotSupported"
	SetMonitoringBaseConfirmationGenericDeviceModelStatusEnumEmptyResultSet SetMonitoringBaseConfirmationGenericDeviceModelStatusEnum = "EmptyResultSet"
)

func (SetMonitoringBaseConfirmation) ActionName() string { return "SetMonitoringBase" }

func (SetMonitoringBaseConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (SetMonitoringBaseConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (SetMonitoringBaseConfirmation) SchemaName() string { return "SetMonitoringBaseResponse.json" }

func (message SetMonitoringBaseConfirmation) Validate() error {
	return validation.Validate("SetMonitoringBaseResponse.json", schemaSetMonitoringBaseConfirmation, message)
}

func (SetMonitoringBaseConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("SetMonitoringBaseResponse.json", schemaSetMonitoringBaseConfirmation, data)
}
