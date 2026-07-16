// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = ClearVariableMonitoringConfirmation{}

var schemaClearVariableMonitoringConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "clearMonitoringResult": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "NotFound"}}, "id": &validation.Schema{Type: "integer", AllowAdditional: true}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}}, Required: []string{"status", "id"}}, MinItems: 1, HasMinItems: true}}, Required: []string{"clearMonitoringResult"}}

type ClearVariableMonitoringConfirmation struct {
	CustomData            *ClearVariableMonitoringConfirmationCustomData             `json:"customData,omitempty"`
	ClearMonitoringResult []ClearVariableMonitoringConfirmationClearMonitoringResult `json:"clearMonitoringResult"`
}

type ClearVariableMonitoringConfirmationClearMonitoringResult struct {
	CustomData *ClearVariableMonitoringConfirmationCustomData               `json:"customData,omitempty"`
	Status     ClearVariableMonitoringConfirmationClearMonitoringStatusEnum `json:"status"`
	ID         int                                                          `json:"id"`
	StatusInfo *ClearVariableMonitoringConfirmationStatusInfo               `json:"statusInfo,omitempty"`
}

type ClearVariableMonitoringConfirmationStatusInfo struct {
	CustomData     *ClearVariableMonitoringConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                                         `json:"reasonCode"`
	AdditionalInfo *string                                        `json:"additionalInfo,omitempty"`
}

type ClearVariableMonitoringConfirmationClearMonitoringStatusEnum string

const (
	ClearVariableMonitoringConfirmationClearMonitoringStatusEnumAccepted ClearVariableMonitoringConfirmationClearMonitoringStatusEnum = "Accepted"
	ClearVariableMonitoringConfirmationClearMonitoringStatusEnumRejected ClearVariableMonitoringConfirmationClearMonitoringStatusEnum = "Rejected"
	ClearVariableMonitoringConfirmationClearMonitoringStatusEnumNotFound ClearVariableMonitoringConfirmationClearMonitoringStatusEnum = "NotFound"
)

type ClearVariableMonitoringConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (ClearVariableMonitoringConfirmation) ActionName() string { return "ClearVariableMonitoring" }

func (ClearVariableMonitoringConfirmation) Version() protocol.Version { return protocol.OCPP201 }

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
