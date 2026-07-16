// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = NotifyMonitoringReportConfirmation{}

var schemaNotifyMonitoringReportConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type NotifyMonitoringReportConfirmation struct {
	CustomData *NotifyMonitoringReportConfirmationCustomData `json:"customData,omitempty"`
}

type NotifyMonitoringReportConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (NotifyMonitoringReportConfirmation) ActionName() string { return "NotifyMonitoringReport" }

func (NotifyMonitoringReportConfirmation) Version() protocol.Version { return protocol.OCPP201 }

func (NotifyMonitoringReportConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (NotifyMonitoringReportConfirmation) SchemaName() string {
	return "NotifyMonitoringReportResponse.json"
}

func (message NotifyMonitoringReportConfirmation) Validate() error {
	return validation.Validate("NotifyMonitoringReportResponse.json", schemaNotifyMonitoringReportConfirmation, message)
}

func (NotifyMonitoringReportConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("NotifyMonitoringReportResponse.json", schemaNotifyMonitoringReportConfirmation, data)
}
