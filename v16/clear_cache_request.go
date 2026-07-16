// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = ClearCacheRequest{}

var schemaClearCacheRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{}}

type ClearCacheRequest struct {
}

func (ClearCacheRequest) ActionName() string { return "ClearCache" }

func (ClearCacheRequest) Version() protocol.Version { return protocol.OCPP16 }

func (ClearCacheRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (ClearCacheRequest) SchemaName() string { return "ClearCache.json" }

func (message ClearCacheRequest) Validate() error {
	return validation.Validate("ClearCache.json", schemaClearCacheRequest, message)
}

func (ClearCacheRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ClearCache.json", schemaClearCacheRequest, data)
}
