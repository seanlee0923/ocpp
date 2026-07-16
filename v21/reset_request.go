// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ResetRequest{}

var schemaResetRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"type": &validation.Schema{Type: "string", Enum: []string{"Immediate", "OnIdle", "ImmediateAndResume"}}, "evseId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"type"}}

type ResetRequest struct {
	Type       ResetRequestResetEnum   `json:"type"`
	EVSEID     *int                    `json:"evseId,omitempty"`
	CustomData *ResetRequestCustomData `json:"customData,omitempty"`
}

type ResetRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value ResetRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *ResetRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

type ResetRequestResetEnum string

const (
	ResetRequestResetEnumImmediate          ResetRequestResetEnum = "Immediate"
	ResetRequestResetEnumOnIdle             ResetRequestResetEnum = "OnIdle"
	ResetRequestResetEnumImmediateAndResume ResetRequestResetEnum = "ImmediateAndResume"
)

func (ResetRequest) ActionName() string { return "Reset" }

func (ResetRequest) Version() protocol.Version { return protocol.OCPP21 }

func (ResetRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (ResetRequest) SchemaName() string { return "ResetRequest.json" }

func (message ResetRequest) Validate() error {
	return validation.Validate("ResetRequest.json", schemaResetRequest, message)
}

func (ResetRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ResetRequest.json", schemaResetRequest, data)
}
