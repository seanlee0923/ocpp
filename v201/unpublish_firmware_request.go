// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = UnpublishFirmwareRequest{}

var schemaUnpublishFirmwareRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "checksum": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 32, HasMaxLength: true}}, Required: []string{"checksum"}}

type UnpublishFirmwareRequest struct {
	CustomData *UnpublishFirmwareRequestCustomData `json:"customData,omitempty"`
	Checksum   string                              `json:"checksum"`
}

type UnpublishFirmwareRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value UnpublishFirmwareRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *UnpublishFirmwareRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (UnpublishFirmwareRequest) ActionName() string { return "UnpublishFirmware" }

func (UnpublishFirmwareRequest) Version() protocol.Version { return protocol.OCPP201 }

func (UnpublishFirmwareRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (UnpublishFirmwareRequest) SchemaName() string { return "UnpublishFirmwareRequest.json" }

func (message UnpublishFirmwareRequest) Validate() error {
	return validation.Validate("UnpublishFirmwareRequest.json", schemaUnpublishFirmwareRequest, message)
}

func (UnpublishFirmwareRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("UnpublishFirmwareRequest.json", schemaUnpublishFirmwareRequest, data)
}
