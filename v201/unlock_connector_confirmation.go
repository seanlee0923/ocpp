// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = UnlockConnectorConfirmation{}

var schemaUnlockConnectorConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Unlocked", "UnlockFailed", "OngoingAuthorizedTransaction", "UnknownConnector"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}}, Required: []string{"status"}}

type UnlockConnectorConfirmation struct {
	CustomData *UnlockConnectorConfirmationCustomData      `json:"customData,omitempty"`
	Status     UnlockConnectorConfirmationUnlockStatusEnum `json:"status"`
	StatusInfo *UnlockConnectorConfirmationStatusInfo      `json:"statusInfo,omitempty"`
}

type UnlockConnectorConfirmationStatusInfo struct {
	CustomData     *UnlockConnectorConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                                 `json:"reasonCode"`
	AdditionalInfo *string                                `json:"additionalInfo,omitempty"`
}

type UnlockConnectorConfirmationUnlockStatusEnum string

const (
	UnlockConnectorConfirmationUnlockStatusEnumUnlocked                     UnlockConnectorConfirmationUnlockStatusEnum = "Unlocked"
	UnlockConnectorConfirmationUnlockStatusEnumUnlockFailed                 UnlockConnectorConfirmationUnlockStatusEnum = "UnlockFailed"
	UnlockConnectorConfirmationUnlockStatusEnumOngoingAuthorizedTransaction UnlockConnectorConfirmationUnlockStatusEnum = "OngoingAuthorizedTransaction"
	UnlockConnectorConfirmationUnlockStatusEnumUnknownConnector             UnlockConnectorConfirmationUnlockStatusEnum = "UnknownConnector"
)

type UnlockConnectorConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value UnlockConnectorConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *UnlockConnectorConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (UnlockConnectorConfirmation) ActionName() string { return "UnlockConnector" }

func (UnlockConnectorConfirmation) Version() protocol.Version { return protocol.OCPP201 }

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
