// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ClearDisplayMessageRequest{}

var schemaClearDisplayMessageRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "id": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"id"}}

type ClearDisplayMessageRequest struct {
	CustomData *ClearDisplayMessageRequestCustomData `json:"customData,omitempty"`
	ID         int                                   `json:"id"`
}

type ClearDisplayMessageRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value ClearDisplayMessageRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *ClearDisplayMessageRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (ClearDisplayMessageRequest) ActionName() string { return "ClearDisplayMessage" }

func (ClearDisplayMessageRequest) Version() protocol.Version { return protocol.OCPP201 }

func (ClearDisplayMessageRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (ClearDisplayMessageRequest) SchemaName() string { return "ClearDisplayMessageRequest.json" }

func (message ClearDisplayMessageRequest) Validate() error {
	return validation.Validate("ClearDisplayMessageRequest.json", schemaClearDisplayMessageRequest, message)
}

func (ClearDisplayMessageRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ClearDisplayMessageRequest.json", schemaClearDisplayMessageRequest, data)
}
