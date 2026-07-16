// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = AdjustPeriodicEventStreamConfirmation{}

var schemaAdjustPeriodicEventStreamConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type AdjustPeriodicEventStreamConfirmation struct {
	Status     AdjustPeriodicEventStreamConfirmationGenericStatusEnum `json:"status"`
	StatusInfo *AdjustPeriodicEventStreamConfirmationStatusInfo       `json:"statusInfo,omitempty"`
	CustomData *AdjustPeriodicEventStreamConfirmationCustomData       `json:"customData,omitempty"`
}

type AdjustPeriodicEventStreamConfirmationStatusInfo struct {
	ReasonCode     string                                           `json:"reasonCode"`
	AdditionalInfo *string                                          `json:"additionalInfo,omitempty"`
	CustomData     *AdjustPeriodicEventStreamConfirmationCustomData `json:"customData,omitempty"`
}

type AdjustPeriodicEventStreamConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value AdjustPeriodicEventStreamConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *AdjustPeriodicEventStreamConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

type AdjustPeriodicEventStreamConfirmationGenericStatusEnum string

const (
	AdjustPeriodicEventStreamConfirmationGenericStatusEnumAccepted AdjustPeriodicEventStreamConfirmationGenericStatusEnum = "Accepted"
	AdjustPeriodicEventStreamConfirmationGenericStatusEnumRejected AdjustPeriodicEventStreamConfirmationGenericStatusEnum = "Rejected"
)

func (AdjustPeriodicEventStreamConfirmation) ActionName() string { return "AdjustPeriodicEventStream" }

func (AdjustPeriodicEventStreamConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (AdjustPeriodicEventStreamConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (AdjustPeriodicEventStreamConfirmation) SchemaName() string {
	return "AdjustPeriodicEventStreamResponse.json"
}

func (message AdjustPeriodicEventStreamConfirmation) Validate() error {
	return validation.Validate("AdjustPeriodicEventStreamResponse.json", schemaAdjustPeriodicEventStreamConfirmation, message)
}

func (AdjustPeriodicEventStreamConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("AdjustPeriodicEventStreamResponse.json", schemaAdjustPeriodicEventStreamConfirmation, data)
}
