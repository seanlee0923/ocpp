// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = GetLogConfirmation{}

var schemaGetLogConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "AcceptedCanceled"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}, "filename": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"status"}}

type GetLogConfirmation struct {
	CustomData *GetLogConfirmationCustomData   `json:"customData,omitempty"`
	Status     GetLogConfirmationLogStatusEnum `json:"status"`
	StatusInfo *GetLogConfirmationStatusInfo   `json:"statusInfo,omitempty"`
	Filename   *string                         `json:"filename,omitempty"`
}

type GetLogConfirmationStatusInfo struct {
	CustomData     *GetLogConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                        `json:"reasonCode"`
	AdditionalInfo *string                       `json:"additionalInfo,omitempty"`
}

type GetLogConfirmationLogStatusEnum string

const (
	GetLogConfirmationLogStatusEnumAccepted         GetLogConfirmationLogStatusEnum = "Accepted"
	GetLogConfirmationLogStatusEnumRejected         GetLogConfirmationLogStatusEnum = "Rejected"
	GetLogConfirmationLogStatusEnumAcceptedCanceled GetLogConfirmationLogStatusEnum = "AcceptedCanceled"
)

type GetLogConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value GetLogConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *GetLogConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (GetLogConfirmation) ActionName() string { return "GetLog" }

func (GetLogConfirmation) Version() protocol.Version { return protocol.OCPP201 }

func (GetLogConfirmation) Direction() protocol.PayloadDirection { return protocol.ConfirmationPayload }

func (GetLogConfirmation) SchemaName() string { return "GetLogResponse.json" }

func (message GetLogConfirmation) Validate() error {
	return validation.Validate("GetLogResponse.json", schemaGetLogConfirmation, message)
}

func (GetLogConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetLogResponse.json", schemaGetLogConfirmation, data)
}
