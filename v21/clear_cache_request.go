// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ClearCacheRequest{}

var schemaClearCacheRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type ClearCacheRequest struct {
	CustomData *ClearCacheRequestCustomData `json:"customData,omitempty"`
}

type ClearCacheRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value ClearCacheRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *ClearCacheRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (ClearCacheRequest) ActionName() string { return "ClearCache" }

func (ClearCacheRequest) Version() protocol.Version { return protocol.OCPP21 }

func (ClearCacheRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (ClearCacheRequest) SchemaName() string { return "ClearCacheRequest.json" }

func (message ClearCacheRequest) Validate() error {
	return validation.Validate("ClearCacheRequest.json", schemaClearCacheRequest, message)
}

func (ClearCacheRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ClearCacheRequest.json", schemaClearCacheRequest, data)
}
