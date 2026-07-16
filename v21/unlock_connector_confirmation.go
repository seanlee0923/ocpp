// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = UnlockConnectorConfirmation{}

var schemaUnlockConnectorConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Unlocked", "UnlockFailed", "OngoingAuthorizedTransaction", "UnknownConnector"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type UnlockConnectorConfirmation struct {
	Status     UnlockConnectorConfirmationUnlockStatusEnum `json:"status"`
	StatusInfo *UnlockConnectorConfirmationStatusInfo      `json:"statusInfo,omitempty"`
	CustomData *UnlockConnectorConfirmationCustomData      `json:"customData,omitempty"`
}

type UnlockConnectorConfirmationStatusInfo struct {
	ReasonCode     string                                 `json:"reasonCode"`
	AdditionalInfo *string                                `json:"additionalInfo,omitempty"`
	CustomData     *UnlockConnectorConfirmationCustomData `json:"customData,omitempty"`
}

type UnlockConnectorConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

type UnlockConnectorConfirmationUnlockStatusEnum string

const (
	UnlockConnectorConfirmationUnlockStatusEnumUnlocked                     UnlockConnectorConfirmationUnlockStatusEnum = "Unlocked"
	UnlockConnectorConfirmationUnlockStatusEnumUnlockFailed                 UnlockConnectorConfirmationUnlockStatusEnum = "UnlockFailed"
	UnlockConnectorConfirmationUnlockStatusEnumOngoingAuthorizedTransaction UnlockConnectorConfirmationUnlockStatusEnum = "OngoingAuthorizedTransaction"
	UnlockConnectorConfirmationUnlockStatusEnumUnknownConnector             UnlockConnectorConfirmationUnlockStatusEnum = "UnknownConnector"
)

func (UnlockConnectorConfirmation) ActionName() string { return "UnlockConnector" }

func (UnlockConnectorConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (UnlockConnectorConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (UnlockConnectorConfirmation) SchemaName() string { return "UnlockConnectorResponse.json" }

func (message UnlockConnectorConfirmation) Validate() error {
	return validation.Validate("UnlockConnectorResponse.json", schemaUnlockConnectorConfirmation, message)
}

func (UnlockConnectorConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("UnlockConnectorResponse.json", schemaUnlockConnectorConfirmation, data)
}
