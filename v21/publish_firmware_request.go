// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = PublishFirmwareRequest{}

var schemaPublishFirmwareRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"location": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 2000, HasMaxLength: true}, "retries": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "checksum": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 32, HasMaxLength: true}, "requestId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "retryInterval": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"location", "checksum", "requestId"}}

type PublishFirmwareRequest struct {
	Location      string                            `json:"location"`
	Retries       *int                              `json:"retries,omitempty"`
	Checksum      string                            `json:"checksum"`
	RequestID     int                               `json:"requestId"`
	RetryInterval *int                              `json:"retryInterval,omitempty"`
	CustomData    *PublishFirmwareRequestCustomData `json:"customData,omitempty"`
}

type PublishFirmwareRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (PublishFirmwareRequest) ActionName() string { return "PublishFirmware" }

func (PublishFirmwareRequest) Version() protocol.Version { return protocol.OCPP21 }

func (PublishFirmwareRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (PublishFirmwareRequest) SchemaName() string { return "PublishFirmwareRequest.json" }

func (message PublishFirmwareRequest) Validate() error {
	return validation.Validate("PublishFirmwareRequest.json", schemaPublishFirmwareRequest, message)
}

func (PublishFirmwareRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("PublishFirmwareRequest.json", schemaPublishFirmwareRequest, data)
}
