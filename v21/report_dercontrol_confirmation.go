// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ReportDERControlConfirmation{}

var schemaReportDERControlConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type ReportDERControlConfirmation struct {
	CustomData *ReportDERControlConfirmationCustomData `json:"customData,omitempty"`
}

type ReportDERControlConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value ReportDERControlConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *ReportDERControlConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
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
