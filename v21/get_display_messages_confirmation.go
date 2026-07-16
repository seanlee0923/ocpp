// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = GetDisplayMessagesConfirmation{}

var schemaGetDisplayMessagesConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Unknown"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type GetDisplayMessagesConfirmation struct {
	Status     GetDisplayMessagesConfirmationGetDisplayMessagesStatusEnum `json:"status"`
	StatusInfo *GetDisplayMessagesConfirmationStatusInfo                  `json:"statusInfo,omitempty"`
	CustomData *GetDisplayMessagesConfirmationCustomData                  `json:"customData,omitempty"`
}

type GetDisplayMessagesConfirmationStatusInfo struct {
	ReasonCode     string                                    `json:"reasonCode"`
	AdditionalInfo *string                                   `json:"additionalInfo,omitempty"`
	CustomData     *GetDisplayMessagesConfirmationCustomData `json:"customData,omitempty"`
}

type GetDisplayMessagesConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value GetDisplayMessagesConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *GetDisplayMessagesConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

type GetDisplayMessagesConfirmationGetDisplayMessagesStatusEnum string

const (
	GetDisplayMessagesConfirmationGetDisplayMessagesStatusEnumAccepted GetDisplayMessagesConfirmationGetDisplayMessagesStatusEnum = "Accepted"
	GetDisplayMessagesConfirmationGetDisplayMessagesStatusEnumUnknown  GetDisplayMessagesConfirmationGetDisplayMessagesStatusEnum = "Unknown"
)

func (GetDisplayMessagesConfirmation) ActionName() string { return "GetDisplayMessages" }

func (GetDisplayMessagesConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (GetDisplayMessagesConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (GetDisplayMessagesConfirmation) SchemaName() string { return "GetDisplayMessagesResponse.json" }

func (message GetDisplayMessagesConfirmation) Validate() error {
	return validation.Validate("GetDisplayMessagesResponse.json", schemaGetDisplayMessagesConfirmation, message)
}

func (GetDisplayMessagesConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetDisplayMessagesResponse.json", schemaGetDisplayMessagesConfirmation, data)
}
