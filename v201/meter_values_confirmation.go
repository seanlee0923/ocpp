// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = MeterValuesConfirmation{}

var schemaMeterValuesConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type MeterValuesConfirmation struct {
	CustomData *MeterValuesConfirmationCustomData `json:"customData,omitempty"`
}

type MeterValuesConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value MeterValuesConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *MeterValuesConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (MeterValuesConfirmation) ActionName() string { return "MeterValues" }

func (MeterValuesConfirmation) Version() protocol.Version { return protocol.OCPP201 }

func (MeterValuesConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (MeterValuesConfirmation) SchemaName() string { return "MeterValuesResponse.json" }

func (message MeterValuesConfirmation) Validate() error {
	return validation.Validate("MeterValuesResponse.json", schemaMeterValuesConfirmation, message)
}

func (MeterValuesConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("MeterValuesResponse.json", schemaMeterValuesConfirmation, data)
}
