// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = FirmwareStatusNotificationConfirmation{}

var schemaFirmwareStatusNotificationConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type FirmwareStatusNotificationConfirmation struct {
	CustomData *FirmwareStatusNotificationConfirmationCustomData `json:"customData,omitempty"`
}

type FirmwareStatusNotificationConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value FirmwareStatusNotificationConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *FirmwareStatusNotificationConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (FirmwareStatusNotificationConfirmation) ActionName() string {
	return "FirmwareStatusNotification"
}

func (FirmwareStatusNotificationConfirmation) Version() protocol.Version { return protocol.OCPP201 }

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
