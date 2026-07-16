// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ClearDisplayMessageConfirmation{}

var schemaClearDisplayMessageConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Unknown"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}}, Required: []string{"status"}}

type ClearDisplayMessageConfirmation struct {
	CustomData *ClearDisplayMessageConfirmationCustomData            `json:"customData,omitempty"`
	Status     ClearDisplayMessageConfirmationClearMessageStatusEnum `json:"status"`
	StatusInfo *ClearDisplayMessageConfirmationStatusInfo            `json:"statusInfo,omitempty"`
}

type ClearDisplayMessageConfirmationStatusInfo struct {
	CustomData     *ClearDisplayMessageConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                                     `json:"reasonCode"`
	AdditionalInfo *string                                    `json:"additionalInfo,omitempty"`
}

type ClearDisplayMessageConfirmationClearMessageStatusEnum string

const (
	ClearDisplayMessageConfirmationClearMessageStatusEnumAccepted ClearDisplayMessageConfirmationClearMessageStatusEnum = "Accepted"
	ClearDisplayMessageConfirmationClearMessageStatusEnumUnknown  ClearDisplayMessageConfirmationClearMessageStatusEnum = "Unknown"
)

type ClearDisplayMessageConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value ClearDisplayMessageConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *ClearDisplayMessageConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (ClearDisplayMessageConfirmation) ActionName() string { return "ClearDisplayMessage" }

func (ClearDisplayMessageConfirmation) Version() protocol.Version { return protocol.OCPP201 }

func (ClearDisplayMessageConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (ClearDisplayMessageConfirmation) SchemaName() string { return "ClearDisplayMessageResponse.json" }

func (message ClearDisplayMessageConfirmation) Validate() error {
	return validation.Validate("ClearDisplayMessageResponse.json", schemaClearDisplayMessageConfirmation, message)
}

func (ClearDisplayMessageConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ClearDisplayMessageResponse.json", schemaClearDisplayMessageConfirmation, data)
}
