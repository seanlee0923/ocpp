// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = CostUpdatedConfirmation{}

var schemaCostUpdatedConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type CostUpdatedConfirmation struct {
	CustomData *CostUpdatedConfirmationCustomData `json:"customData,omitempty"`
}

type CostUpdatedConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value CostUpdatedConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *CostUpdatedConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (CostUpdatedConfirmation) ActionName() string { return "CostUpdated" }

func (CostUpdatedConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (CostUpdatedConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (CostUpdatedConfirmation) SchemaName() string { return "CostUpdatedResponse.json" }

func (message CostUpdatedConfirmation) Validate() error {
	return validation.Validate("CostUpdatedResponse.json", schemaCostUpdatedConfirmation, message)
}

func (CostUpdatedConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("CostUpdatedResponse.json", schemaCostUpdatedConfirmation, data)
}
