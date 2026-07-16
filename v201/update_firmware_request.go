// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = UpdateFirmwareRequest{}

var schemaUpdateFirmwareRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "retries": &validation.Schema{Type: "integer", AllowAdditional: true}, "retryInterval": &validation.Schema{Type: "integer", AllowAdditional: true}, "requestId": &validation.Schema{Type: "integer", AllowAdditional: true}, "firmware": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "location": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}, "retrieveDateTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "installDateTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "signingCertificate": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 5500, HasMaxLength: true}, "signature": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 800, HasMaxLength: true}}, Required: []string{"location", "retrieveDateTime"}}}, Required: []string{"requestId", "firmware"}}

type UpdateFirmwareRequest struct {
	CustomData    *UpdateFirmwareRequestCustomData `json:"customData,omitempty"`
	Retries       *int                             `json:"retries,omitempty"`
	RetryInterval *int                             `json:"retryInterval,omitempty"`
	RequestID     int                              `json:"requestId"`
	Firmware      UpdateFirmwareRequestFirmware    `json:"firmware"`
}

type UpdateFirmwareRequestFirmware struct {
	CustomData         *UpdateFirmwareRequestCustomData `json:"customData,omitempty"`
	Location           string                           `json:"location"`
	RetrieveDateTime   string                           `json:"retrieveDateTime"`
	InstallDateTime    *string                          `json:"installDateTime,omitempty"`
	SigningCertificate *string                          `json:"signingCertificate,omitempty"`
	Signature          *string                          `json:"signature,omitempty"`
}

type UpdateFirmwareRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value UpdateFirmwareRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *UpdateFirmwareRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (UpdateFirmwareRequest) ActionName() string { return "UpdateFirmware" }

func (UpdateFirmwareRequest) Version() protocol.Version { return protocol.OCPP201 }

func (UpdateFirmwareRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (UpdateFirmwareRequest) SchemaName() string { return "UpdateFirmwareRequest.json" }

func (message UpdateFirmwareRequest) Validate() error {
	return validation.Validate("UpdateFirmwareRequest.json", schemaUpdateFirmwareRequest, message)
}

func (UpdateFirmwareRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("UpdateFirmwareRequest.json", schemaUpdateFirmwareRequest, data)
}
