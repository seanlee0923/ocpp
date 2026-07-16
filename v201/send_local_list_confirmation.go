// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = SendLocalListConfirmation{}

var schemaSendLocalListConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Failed", "VersionMismatch"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}}, Required: []string{"status"}}

type SendLocalListConfirmation struct {
	CustomData *SendLocalListConfirmationCustomData             `json:"customData,omitempty"`
	Status     SendLocalListConfirmationSendLocalListStatusEnum `json:"status"`
	StatusInfo *SendLocalListConfirmationStatusInfo             `json:"statusInfo,omitempty"`
}

type SendLocalListConfirmationStatusInfo struct {
	CustomData     *SendLocalListConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                               `json:"reasonCode"`
	AdditionalInfo *string                              `json:"additionalInfo,omitempty"`
}

type SendLocalListConfirmationSendLocalListStatusEnum string

const (
	SendLocalListConfirmationSendLocalListStatusEnumAccepted        SendLocalListConfirmationSendLocalListStatusEnum = "Accepted"
	SendLocalListConfirmationSendLocalListStatusEnumFailed          SendLocalListConfirmationSendLocalListStatusEnum = "Failed"
	SendLocalListConfirmationSendLocalListStatusEnumVersionMismatch SendLocalListConfirmationSendLocalListStatusEnum = "VersionMismatch"
)

type SendLocalListConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value SendLocalListConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *SendLocalListConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (SendLocalListConfirmation) ActionName() string { return "SendLocalList" }

func (SendLocalListConfirmation) Version() protocol.Version { return protocol.OCPP201 }

func (SendLocalListConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (SendLocalListConfirmation) SchemaName() string { return "SendLocalListResponse.json" }

func (message SendLocalListConfirmation) Validate() error {
	return validation.Validate("SendLocalListResponse.json", schemaSendLocalListConfirmation, message)
}

func (SendLocalListConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("SendLocalListResponse.json", schemaSendLocalListConfirmation, data)
}
