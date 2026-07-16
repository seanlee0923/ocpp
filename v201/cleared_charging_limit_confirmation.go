// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ClearedChargingLimitConfirmation{}

var schemaClearedChargingLimitConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type ClearedChargingLimitConfirmation struct {
	CustomData *ClearedChargingLimitConfirmationCustomData `json:"customData,omitempty"`
}

type ClearedChargingLimitConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value ClearedChargingLimitConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *ClearedChargingLimitConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (ClearedChargingLimitConfirmation) ActionName() string { return "ClearedChargingLimit" }

func (ClearedChargingLimitConfirmation) Version() protocol.Version { return protocol.OCPP201 }

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
