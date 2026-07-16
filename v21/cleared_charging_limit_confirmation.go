// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ClearedChargingLimitConfirmation{}

var schemaClearedChargingLimitConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type ClearedChargingLimitConfirmation struct {
	CustomData *ClearedChargingLimitConfirmationCustomData `json:"customData,omitempty"`
}

type ClearedChargingLimitConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (ClearedChargingLimitConfirmation) ActionName() string { return "ClearedChargingLimit" }

func (ClearedChargingLimitConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (ClearedChargingLimitConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (ClearedChargingLimitConfirmation) SchemaName() string {
	return "ClearedChargingLimitResponse.json"
}

func (message ClearedChargingLimitConfirmation) Validate() error {
	return validation.Validate("ClearedChargingLimitResponse.json", schemaClearedChargingLimitConfirmation, message)
}

func (ClearedChargingLimitConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ClearedChargingLimitResponse.json", schemaClearedChargingLimitConfirmation, data)
}
