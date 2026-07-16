// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = SetMonitoringLevelConfirmation{}

var schemaSetMonitoringLevelConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}}, Required: []string{"status"}}

type SetMonitoringLevelConfirmation struct {
	CustomData *SetMonitoringLevelConfirmationCustomData       `json:"customData,omitempty"`
	Status     SetMonitoringLevelConfirmationGenericStatusEnum `json:"status"`
	StatusInfo *SetMonitoringLevelConfirmationStatusInfo       `json:"statusInfo,omitempty"`
}

type SetMonitoringLevelConfirmationStatusInfo struct {
	CustomData     *SetMonitoringLevelConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                                    `json:"reasonCode"`
	AdditionalInfo *string                                   `json:"additionalInfo,omitempty"`
}

type SetMonitoringLevelConfirmationGenericStatusEnum string

const (
	SetMonitoringLevelConfirmationGenericStatusEnumAccepted SetMonitoringLevelConfirmationGenericStatusEnum = "Accepted"
	SetMonitoringLevelConfirmationGenericStatusEnumRejected SetMonitoringLevelConfirmationGenericStatusEnum = "Rejected"
)

type SetMonitoringLevelConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (SetMonitoringLevelConfirmation) ActionName() string { return "SetMonitoringLevel" }

func (SetMonitoringLevelConfirmation) Version() protocol.Version { return protocol.OCPP201 }

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
