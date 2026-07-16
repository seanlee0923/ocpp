// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = ClearDisplayMessageRequest{}

var schemaClearDisplayMessageRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "id": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"id"}}

type ClearDisplayMessageRequest struct {
	CustomData *ClearDisplayMessageRequestCustomData `json:"customData,omitempty"`
	ID         int                                   `json:"id"`
}

type ClearDisplayMessageRequestCustomData struct {
	VendorID string `json:"vendorId"`
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
