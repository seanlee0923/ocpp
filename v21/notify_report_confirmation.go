// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = NotifyReportConfirmation{}

var schemaNotifyReportConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type NotifyReportConfirmation struct {
	CustomData *NotifyReportConfirmationCustomData `json:"customData,omitempty"`
}

type NotifyReportConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (NotifyReportConfirmation) ActionName() string { return "NotifyReport" }

func (NotifyReportConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (NotifyReportConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (NotifyReportConfirmation) SchemaName() string { return "NotifyReportResponse.json" }

func (message NotifyReportConfirmation) Validate() error {
	return validation.Validate("NotifyReportResponse.json", schemaNotifyReportConfirmation, message)
}

func (NotifyReportConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("NotifyReportResponse.json", schemaNotifyReportConfirmation, data)
}
