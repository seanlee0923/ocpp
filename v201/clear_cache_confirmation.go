// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ClearCacheConfirmation{}

var schemaClearCacheConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}}, Required: []string{"status"}}

type ClearCacheConfirmation struct {
	CustomData *ClearCacheConfirmationCustomData          `json:"customData,omitempty"`
	Status     ClearCacheConfirmationClearCacheStatusEnum `json:"status"`
	StatusInfo *ClearCacheConfirmationStatusInfo          `json:"statusInfo,omitempty"`
}

type ClearCacheConfirmationStatusInfo struct {
	CustomData     *ClearCacheConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                            `json:"reasonCode"`
	AdditionalInfo *string                           `json:"additionalInfo,omitempty"`
}

type ClearCacheConfirmationClearCacheStatusEnum string

const (
	ClearCacheConfirmationClearCacheStatusEnumAccepted ClearCacheConfirmationClearCacheStatusEnum = "Accepted"
	ClearCacheConfirmationClearCacheStatusEnumRejected ClearCacheConfirmationClearCacheStatusEnum = "Rejected"
)

type ClearCacheConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value ClearCacheConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *ClearCacheConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (ClearCacheConfirmation) ActionName() string { return "ClearCache" }

func (ClearCacheConfirmation) Version() protocol.Version { return protocol.OCPP201 }

func (ClearCacheConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (ClearCacheConfirmation) SchemaName() string { return "ClearCacheResponse.json" }

func (message ClearCacheConfirmation) Validate() error {
	return validation.Validate("ClearCacheResponse.json", schemaClearCacheConfirmation, message)
}

func (ClearCacheConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ClearCacheResponse.json", schemaClearCacheConfirmation, data)
}
