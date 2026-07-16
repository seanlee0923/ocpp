// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = ReportDERControlConfirmation{}

var schemaReportDERControlConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type ReportDERControlConfirmation struct {
	CustomData *ReportDERControlConfirmationCustomData `json:"customData,omitempty"`
}

type ReportDERControlConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (ReportDERControlConfirmation) ActionName() string { return "ReportDERControl" }

func (ReportDERControlConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (ReportDERControlConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (ReportDERControlConfirmation) SchemaName() string { return "ReportDERControlResponse.json" }

func (message ReportDERControlConfirmation) Validate() error {
	return validation.Validate("ReportDERControlResponse.json", schemaReportDERControlConfirmation, message)
}

func (ReportDERControlConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ReportDERControlResponse.json", schemaReportDERControlConfirmation, data)
}
