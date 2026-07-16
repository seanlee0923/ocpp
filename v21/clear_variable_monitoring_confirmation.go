// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = ClearVariableMonitoringConfirmation{}

var schemaClearVariableMonitoringConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"clearMonitoringResult": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "NotFound"}}, "id": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status", "id"}}, MinItems: 1, HasMinItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"clearMonitoringResult"}}

type ClearVariableMonitoringConfirmation struct {
	ClearMonitoringResult []ClearVariableMonitoringConfirmationClearMonitoringResult `json:"clearMonitoringResult"`
	CustomData            *ClearVariableMonitoringConfirmationCustomData             `json:"customData,omitempty"`
}

type ClearVariableMonitoringConfirmationClearMonitoringResult struct {
	Status     ClearVariableMonitoringConfirmationClearMonitoringStatusEnum `json:"status"`
	ID         int                                                          `json:"id"`
	StatusInfo *ClearVariableMonitoringConfirmationStatusInfo               `json:"statusInfo,omitempty"`
	CustomData *ClearVariableMonitoringConfirmationCustomData               `json:"customData,omitempty"`
}

type ClearVariableMonitoringConfirmationStatusInfo struct {
	ReasonCode     string                                         `json:"reasonCode"`
	AdditionalInfo *string                                        `json:"additionalInfo,omitempty"`
	CustomData     *ClearVariableMonitoringConfirmationCustomData `json:"customData,omitempty"`
}

type ClearVariableMonitoringConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

type ClearVariableMonitoringConfirmationClearMonitoringStatusEnum string

const (
	ClearVariableMonitoringConfirmationClearMonitoringStatusEnumAccepted ClearVariableMonitoringConfirmationClearMonitoringStatusEnum = "Accepted"
	ClearVariableMonitoringConfirmationClearMonitoringStatusEnumRejected ClearVariableMonitoringConfirmationClearMonitoringStatusEnum = "Rejected"
	ClearVariableMonitoringConfirmationClearMonitoringStatusEnumNotFound ClearVariableMonitoringConfirmationClearMonitoringStatusEnum = "NotFound"
)

func (ClearVariableMonitoringConfirmation) ActionName() string { return "ClearVariableMonitoring" }

func (ClearVariableMonitoringConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (ClearVariableMonitoringConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (ClearVariableMonitoringConfirmation) SchemaName() string {
	return "ClearVariableMonitoringResponse.json"
}

func (message ClearVariableMonitoringConfirmation) Validate() error {
	return validation.Validate("ClearVariableMonitoringResponse.json", schemaClearVariableMonitoringConfirmation, message)
}

func (ClearVariableMonitoringConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ClearVariableMonitoringResponse.json", schemaClearVariableMonitoringConfirmation, data)
}
