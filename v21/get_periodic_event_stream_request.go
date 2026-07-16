// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = GetPeriodicEventStreamRequest{}

var schemaGetPeriodicEventStreamRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type GetPeriodicEventStreamRequest struct {
	CustomData *GetPeriodicEventStreamRequestCustomData `json:"customData,omitempty"`
}

type GetPeriodicEventStreamRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (GetPeriodicEventStreamRequest) ActionName() string { return "GetPeriodicEventStream" }

func (GetPeriodicEventStreamRequest) Version() protocol.Version { return protocol.OCPP21 }

func (GetPeriodicEventStreamRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (GetPeriodicEventStreamRequest) SchemaName() string { return "GetPeriodicEventStreamRequest.json" }

func (message GetPeriodicEventStreamRequest) Validate() error {
	return validation.Validate("GetPeriodicEventStreamRequest.json", schemaGetPeriodicEventStreamRequest, message)
}

func (GetPeriodicEventStreamRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetPeriodicEventStreamRequest.json", schemaGetPeriodicEventStreamRequest, data)
}
