// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = PublishFirmwareStatusNotificationConfirmation{}

var schemaPublishFirmwareStatusNotificationConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type PublishFirmwareStatusNotificationConfirmation struct {
	CustomData *PublishFirmwareStatusNotificationConfirmationCustomData `json:"customData,omitempty"`
}

type PublishFirmwareStatusNotificationConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value PublishFirmwareStatusNotificationConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *PublishFirmwareStatusNotificationConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (PublishFirmwareStatusNotificationConfirmation) ActionName() string {
	return "PublishFirmwareStatusNotification"
}

func (PublishFirmwareStatusNotificationConfirmation) Version() protocol.Version {
	return protocol.OCPP201
}

func (PublishFirmwareStatusNotificationConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (PublishFirmwareStatusNotificationConfirmation) SchemaName() string {
	return "PublishFirmwareStatusNotificationResponse.json"
}

func (message PublishFirmwareStatusNotificationConfirmation) Validate() error {
	return validation.Validate("PublishFirmwareStatusNotificationResponse.json", schemaPublishFirmwareStatusNotificationConfirmation, message)
}

func (PublishFirmwareStatusNotificationConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("PublishFirmwareStatusNotificationResponse.json", schemaPublishFirmwareStatusNotificationConfirmation, data)
}
