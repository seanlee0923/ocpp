// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = ClearCacheRequest{}

var schemaClearCacheRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type ClearCacheRequest struct {
	CustomData *ClearCacheRequestCustomData `json:"customData,omitempty"`
}

type ClearCacheRequestCustomData struct {
	VendorID string `json:"vendorId"`
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
