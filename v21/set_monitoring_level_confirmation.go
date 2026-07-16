// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = SetMonitoringLevelConfirmation{}

var schemaSetMonitoringLevelConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type SetMonitoringLevelConfirmation struct {
	Status     SetMonitoringLevelConfirmationGenericStatusEnum `json:"status"`
	StatusInfo *SetMonitoringLevelConfirmationStatusInfo       `json:"statusInfo,omitempty"`
	CustomData *SetMonitoringLevelConfirmationCustomData       `json:"customData,omitempty"`
}

type SetMonitoringLevelConfirmationStatusInfo struct {
	ReasonCode     string                                    `json:"reasonCode"`
	AdditionalInfo *string                                   `json:"additionalInfo,omitempty"`
	CustomData     *SetMonitoringLevelConfirmationCustomData `json:"customData,omitempty"`
}

type SetMonitoringLevelConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

type SetMonitoringLevelConfirmationGenericStatusEnum string

const (
	SetMonitoringLevelConfirmationGenericStatusEnumAccepted SetMonitoringLevelConfirmationGenericStatusEnum = "Accepted"
	SetMonitoringLevelConfirmationGenericStatusEnumRejected SetMonitoringLevelConfirmationGenericStatusEnum = "Rejected"
)

func (SetMonitoringLevelConfirmation) ActionName() string { return "SetMonitoringLevel" }

func (SetMonitoringLevelConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (SetMonitoringLevelConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (SetMonitoringLevelConfirmation) SchemaName() string { return "SetMonitoringLevelResponse.json" }

func (message SetMonitoringLevelConfirmation) Validate() error {
	return validation.Validate("SetMonitoringLevelResponse.json", schemaSetMonitoringLevelConfirmation, message)
}

func (SetMonitoringLevelConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("SetMonitoringLevelResponse.json", schemaSetMonitoringLevelConfirmation, data)
}
