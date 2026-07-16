// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = LogStatusNotificationConfirmation{}

var schemaLogStatusNotificationConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type LogStatusNotificationConfirmation struct {
	CustomData *LogStatusNotificationConfirmationCustomData `json:"customData,omitempty"`
}

type LogStatusNotificationConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value LogStatusNotificationConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *LogStatusNotificationConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (LogStatusNotificationConfirmation) ActionName() string { return "LogStatusNotification" }

func (LogStatusNotificationConfirmation) Version() protocol.Version { return protocol.OCPP201 }

func (LogStatusNotificationConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (LogStatusNotificationConfirmation) SchemaName() string {
	return "LogStatusNotificationResponse.json"
}

func (message LogStatusNotificationConfirmation) Validate() error {
	return validation.Validate("LogStatusNotificationResponse.json", schemaLogStatusNotificationConfirmation, message)
}

func (LogStatusNotificationConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("LogStatusNotificationResponse.json", schemaLogStatusNotificationConfirmation, data)
}
