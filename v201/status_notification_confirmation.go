// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = StatusNotificationConfirmation{}

var schemaStatusNotificationConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type StatusNotificationConfirmation struct {
	CustomData *StatusNotificationConfirmationCustomData `json:"customData,omitempty"`
}

type StatusNotificationConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value StatusNotificationConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *StatusNotificationConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (StatusNotificationConfirmation) ActionName() string { return "StatusNotification" }

func (StatusNotificationConfirmation) Version() protocol.Version { return protocol.OCPP201 }

func (StatusNotificationConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (StatusNotificationConfirmation) SchemaName() string { return "StatusNotificationResponse.json" }

func (message StatusNotificationConfirmation) Validate() error {
	return validation.Validate("StatusNotificationResponse.json", schemaStatusNotificationConfirmation, message)
}

func (StatusNotificationConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("StatusNotificationResponse.json", schemaStatusNotificationConfirmation, data)
}
