// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = NotifyReportConfirmation{}

var schemaNotifyReportConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type NotifyReportConfirmation struct {
	CustomData *NotifyReportConfirmationCustomData `json:"customData,omitempty"`
}

type NotifyReportConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value NotifyReportConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *NotifyReportConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
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
