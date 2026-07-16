// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = UpdateFirmwareRequest{}

var schemaUpdateFirmwareRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"retries": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "retryInterval": &validation.Schema{Type: "integer", AllowAdditional: true}, "requestId": &validation.Schema{Type: "integer", AllowAdditional: true}, "firmware": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"location": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 2000, HasMaxLength: true}, "retrieveDateTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "installDateTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "signingCertificate": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 5500, HasMaxLength: true}, "signature": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 800, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"location", "retrieveDateTime"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"requestId", "firmware"}}

type UpdateFirmwareRequest struct {
	Retries       *int                             `json:"retries,omitempty"`
	RetryInterval *int                             `json:"retryInterval,omitempty"`
	RequestID     int                              `json:"requestId"`
	Firmware      UpdateFirmwareRequestFirmware    `json:"firmware"`
	CustomData    *UpdateFirmwareRequestCustomData `json:"customData,omitempty"`
}

type UpdateFirmwareRequestFirmware struct {
	Location           string                           `json:"location"`
	RetrieveDateTime   string                           `json:"retrieveDateTime"`
	InstallDateTime    *string                          `json:"installDateTime,omitempty"`
	SigningCertificate *string                          `json:"signingCertificate,omitempty"`
	Signature          *string                          `json:"signature,omitempty"`
	CustomData         *UpdateFirmwareRequestCustomData `json:"customData,omitempty"`
}

type UpdateFirmwareRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (UpdateFirmwareRequest) ActionName() string { return "UpdateFirmware" }

func (UpdateFirmwareRequest) Version() protocol.Version { return protocol.OCPP21 }

func (UpdateFirmwareRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (UpdateFirmwareRequest) SchemaName() string { return "UpdateFirmwareRequest.json" }

func (message UpdateFirmwareRequest) Validate() error {
	return validation.Validate("UpdateFirmwareRequest.json", schemaUpdateFirmwareRequest, message)
}

func (UpdateFirmwareRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("UpdateFirmwareRequest.json", schemaUpdateFirmwareRequest, data)
}
