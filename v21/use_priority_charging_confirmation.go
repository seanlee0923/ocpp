// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = UsePriorityChargingConfirmation{}

var schemaUsePriorityChargingConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "NoProfile"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type UsePriorityChargingConfirmation struct {
	Status     UsePriorityChargingConfirmationPriorityChargingStatusEnum `json:"status"`
	StatusInfo *UsePriorityChargingConfirmationStatusInfo                `json:"statusInfo,omitempty"`
	CustomData *UsePriorityChargingConfirmationCustomData                `json:"customData,omitempty"`
}

type UsePriorityChargingConfirmationStatusInfo struct {
	ReasonCode     string                                     `json:"reasonCode"`
	AdditionalInfo *string                                    `json:"additionalInfo,omitempty"`
	CustomData     *UsePriorityChargingConfirmationCustomData `json:"customData,omitempty"`
}

type UsePriorityChargingConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value UsePriorityChargingConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *UsePriorityChargingConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

type UsePriorityChargingConfirmationPriorityChargingStatusEnum string

const (
	UsePriorityChargingConfirmationPriorityChargingStatusEnumAccepted  UsePriorityChargingConfirmationPriorityChargingStatusEnum = "Accepted"
	UsePriorityChargingConfirmationPriorityChargingStatusEnumRejected  UsePriorityChargingConfirmationPriorityChargingStatusEnum = "Rejected"
	UsePriorityChargingConfirmationPriorityChargingStatusEnumNoProfile UsePriorityChargingConfirmationPriorityChargingStatusEnum = "NoProfile"
)

func (UsePriorityChargingConfirmation) ActionName() string { return "UsePriorityCharging" }

func (UsePriorityChargingConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (UsePriorityChargingConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (UsePriorityChargingConfirmation) SchemaName() string { return "UsePriorityChargingResponse.json" }

func (message UsePriorityChargingConfirmation) Validate() error {
	return validation.Validate("UsePriorityChargingResponse.json", schemaUsePriorityChargingConfirmation, message)
}

func (UsePriorityChargingConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("UsePriorityChargingResponse.json", schemaUsePriorityChargingConfirmation, data)
}
