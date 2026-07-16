// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = SecurityEventNotificationConfirmation{}

var schemaSecurityEventNotificationConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type SecurityEventNotificationConfirmation struct {
	CustomData *SecurityEventNotificationConfirmationCustomData `json:"customData,omitempty"`
}

type SecurityEventNotificationConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value SecurityEventNotificationConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *SecurityEventNotificationConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (SecurityEventNotificationConfirmation) ActionName() string { return "SecurityEventNotification" }

func (SecurityEventNotificationConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (SecurityEventNotificationConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (SecurityEventNotificationConfirmation) SchemaName() string {
	return "SecurityEventNotificationResponse.json"
}

func (message SecurityEventNotificationConfirmation) Validate() error {
	return validation.Validate("SecurityEventNotificationResponse.json", schemaSecurityEventNotificationConfirmation, message)
}

func (SecurityEventNotificationConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("SecurityEventNotificationResponse.json", schemaSecurityEventNotificationConfirmation, data)
}
