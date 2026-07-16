// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = DataTransferRequest{}

var schemaDataTransferRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}, "messageId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "data": &validation.Schema{Type: "string", AllowAdditional: true}}, Required: []string{"vendorId"}}

type DataTransferRequest struct {
	VendorID  string  `json:"vendorId"`
	MessageID *string `json:"messageId,omitempty"`
	Data      *string `json:"data,omitempty"`
}

func (DataTransferRequest) ActionName() string { return "DataTransfer" }

func (DataTransferRequest) Version() protocol.Version { return protocol.OCPP16 }

func (DataTransferRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (DataTransferRequest) SchemaName() string { return "DataTransfer.json" }

func (message DataTransferRequest) Validate() error {
	return validation.Validate("DataTransfer.json", schemaDataTransferRequest, message)
}

func (DataTransferRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("DataTransfer.json", schemaDataTransferRequest, data)
}
