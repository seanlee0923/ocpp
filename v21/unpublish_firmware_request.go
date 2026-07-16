// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = UnpublishFirmwareRequest{}

var schemaUnpublishFirmwareRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"checksum": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 32, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"checksum"}}

type UnpublishFirmwareRequest struct {
	Checksum   string                              `json:"checksum"`
	CustomData *UnpublishFirmwareRequestCustomData `json:"customData,omitempty"`
}

type UnpublishFirmwareRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (UnpublishFirmwareRequest) ActionName() string { return "UnpublishFirmware" }

func (UnpublishFirmwareRequest) Version() protocol.Version { return protocol.OCPP21 }

func (UnpublishFirmwareRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (UnpublishFirmwareRequest) SchemaName() string { return "UnpublishFirmwareRequest.json" }

func (message UnpublishFirmwareRequest) Validate() error {
	return validation.Validate("UnpublishFirmwareRequest.json", schemaUnpublishFirmwareRequest, message)
}

func (UnpublishFirmwareRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("UnpublishFirmwareRequest.json", schemaUnpublishFirmwareRequest, data)
}
