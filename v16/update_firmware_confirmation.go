// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = UpdateFirmwareConfirmation{}

var schemaUpdateFirmwareConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{}}

type UpdateFirmwareConfirmation struct {
}

func (UpdateFirmwareConfirmation) ActionName() string { return "UpdateFirmware" }

func (UpdateFirmwareConfirmation) Version() protocol.Version { return protocol.OCPP16 }

func (UpdateFirmwareConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (UpdateFirmwareConfirmation) SchemaName() string { return "UpdateFirmwareResponse.json" }

func (message UpdateFirmwareConfirmation) Validate() error {
	return validation.Validate("UpdateFirmwareResponse.json", schemaUpdateFirmwareConfirmation, message)
}

func (UpdateFirmwareConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("UpdateFirmwareResponse.json", schemaUpdateFirmwareConfirmation, data)
}
