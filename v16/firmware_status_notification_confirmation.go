// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = FirmwareStatusNotificationConfirmation{}

var schemaFirmwareStatusNotificationConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{}}

type FirmwareStatusNotificationConfirmation struct {
}

func (FirmwareStatusNotificationConfirmation) ActionName() string {
	return "FirmwareStatusNotification"
}

func (FirmwareStatusNotificationConfirmation) Version() protocol.Version { return protocol.OCPP16 }

func (FirmwareStatusNotificationConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (FirmwareStatusNotificationConfirmation) SchemaName() string {
	return "FirmwareStatusNotificationResponse.json"
}

func (message FirmwareStatusNotificationConfirmation) Validate() error {
	return validation.Validate("FirmwareStatusNotificationResponse.json", schemaFirmwareStatusNotificationConfirmation, message)
}

func (FirmwareStatusNotificationConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("FirmwareStatusNotificationResponse.json", schemaFirmwareStatusNotificationConfirmation, data)
}
