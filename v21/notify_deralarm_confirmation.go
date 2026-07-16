// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = NotifyDERAlarmConfirmation{}

var schemaNotifyDERAlarmConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type NotifyDERAlarmConfirmation struct {
	CustomData *NotifyDERAlarmConfirmationCustomData `json:"customData,omitempty"`
}

type NotifyDERAlarmConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value NotifyDERAlarmConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *NotifyDERAlarmConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (NotifyDERAlarmConfirmation) ActionName() string { return "NotifyDERAlarm" }

func (NotifyDERAlarmConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (NotifyDERAlarmConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (NotifyDERAlarmConfirmation) SchemaName() string { return "NotifyDERAlarmResponse.json" }

func (message NotifyDERAlarmConfirmation) Validate() error {
	return validation.Validate("NotifyDERAlarmResponse.json", schemaNotifyDERAlarmConfirmation, message)
}

func (NotifyDERAlarmConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("NotifyDERAlarmResponse.json", schemaNotifyDERAlarmConfirmation, data)
}
