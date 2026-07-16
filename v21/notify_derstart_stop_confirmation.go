// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = NotifyDERStartStopConfirmation{}

var schemaNotifyDERStartStopConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type NotifyDERStartStopConfirmation struct {
	CustomData *NotifyDERStartStopConfirmationCustomData `json:"customData,omitempty"`
}

type NotifyDERStartStopConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (NotifyDERStartStopConfirmation) ActionName() string { return "NotifyDERStartStop" }

func (NotifyDERStartStopConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (NotifyDERStartStopConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (NotifyDERStartStopConfirmation) SchemaName() string { return "NotifyDERStartStopResponse.json" }

func (message NotifyDERStartStopConfirmation) Validate() error {
	return validation.Validate("NotifyDERStartStopResponse.json", schemaNotifyDERStartStopConfirmation, message)
}

func (NotifyDERStartStopConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("NotifyDERStartStopResponse.json", schemaNotifyDERStartStopConfirmation, data)
}
