// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = PublishFirmwareRequest{}

var schemaPublishFirmwareRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "location": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}, "retries": &validation.Schema{Type: "integer", AllowAdditional: true}, "checksum": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 32, HasMaxLength: true}, "requestId": &validation.Schema{Type: "integer", AllowAdditional: true}, "retryInterval": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"location", "checksum", "requestId"}}

type PublishFirmwareRequest struct {
	CustomData    *PublishFirmwareRequestCustomData `json:"customData,omitempty"`
	Location      string                            `json:"location"`
	Retries       *int                              `json:"retries,omitempty"`
	Checksum      string                            `json:"checksum"`
	RequestID     int                               `json:"requestId"`
	RetryInterval *int                              `json:"retryInterval,omitempty"`
}

type PublishFirmwareRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (PublishFirmwareRequest) ActionName() string { return "PublishFirmware" }

func (PublishFirmwareRequest) Version() protocol.Version { return protocol.OCPP201 }

func (PublishFirmwareRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (PublishFirmwareRequest) SchemaName() string { return "PublishFirmwareRequest.json" }

func (message PublishFirmwareRequest) Validate() error {
	return validation.Validate("PublishFirmwareRequest.json", schemaPublishFirmwareRequest, message)
}

func (PublishFirmwareRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("PublishFirmwareRequest.json", schemaPublishFirmwareRequest, data)
}
