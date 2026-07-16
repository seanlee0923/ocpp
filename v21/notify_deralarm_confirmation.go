// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = NotifyDERAlarmConfirmation{}

var schemaNotifyDERAlarmConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type NotifyDERAlarmConfirmation struct {
	CustomData *NotifyDERAlarmConfirmationCustomData `json:"customData,omitempty"`
}

type NotifyDERAlarmConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
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
