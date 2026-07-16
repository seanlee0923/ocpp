// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = DataTransferRequest{}

var schemaDataTransferRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"messageId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "data": &validation.Schema{AllowAdditional: true}, "vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"vendorId"}}

type DataTransferRequest struct {
	MessageID  *string                        `json:"messageId,omitempty"`
	Data       any                            `json:"data,omitempty"`
	VendorID   string                         `json:"vendorId"`
	CustomData *DataTransferRequestCustomData `json:"customData,omitempty"`
}

type DataTransferRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (DataTransferRequest) ActionName() string { return "DataTransfer" }

func (DataTransferRequest) Version() protocol.Version { return protocol.OCPP21 }

func (DataTransferRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (DataTransferRequest) SchemaName() string { return "DataTransferRequest.json" }

func (message DataTransferRequest) Validate() error {
	return validation.Validate("DataTransferRequest.json", schemaDataTransferRequest, message)
}

func (DataTransferRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("DataTransferRequest.json", schemaDataTransferRequest, data)
}
