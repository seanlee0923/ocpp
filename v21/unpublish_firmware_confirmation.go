// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = UnpublishFirmwareConfirmation{}

var schemaUnpublishFirmwareConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"DownloadOngoing", "NoFirmware", "Unpublished"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type UnpublishFirmwareConfirmation struct {
	Status     UnpublishFirmwareConfirmationUnpublishFirmwareStatusEnum `json:"status"`
	CustomData *UnpublishFirmwareConfirmationCustomData                 `json:"customData,omitempty"`
}

type UnpublishFirmwareConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

type UnpublishFirmwareConfirmationUnpublishFirmwareStatusEnum string

const (
	UnpublishFirmwareConfirmationUnpublishFirmwareStatusEnumDownloadOngoing UnpublishFirmwareConfirmationUnpublishFirmwareStatusEnum = "DownloadOngoing"
	UnpublishFirmwareConfirmationUnpublishFirmwareStatusEnumNoFirmware      UnpublishFirmwareConfirmationUnpublishFirmwareStatusEnum = "NoFirmware"
	UnpublishFirmwareConfirmationUnpublishFirmwareStatusEnumUnpublished     UnpublishFirmwareConfirmationUnpublishFirmwareStatusEnum = "Unpublished"
)

func (UnpublishFirmwareConfirmation) ActionName() string { return "UnpublishFirmware" }

func (UnpublishFirmwareConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (UnpublishFirmwareConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (UnpublishFirmwareConfirmation) SchemaName() string { return "UnpublishFirmwareResponse.json" }

func (message UnpublishFirmwareConfirmation) Validate() error {
	return validation.Validate("UnpublishFirmwareResponse.json", schemaUnpublishFirmwareConfirmation, message)
}

func (UnpublishFirmwareConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("UnpublishFirmwareResponse.json", schemaUnpublishFirmwareConfirmation, data)
}
