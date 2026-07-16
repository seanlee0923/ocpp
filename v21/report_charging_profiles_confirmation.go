// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ReportChargingProfilesConfirmation{}

var schemaReportChargingProfilesConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type ReportChargingProfilesConfirmation struct {
	CustomData *ReportChargingProfilesConfirmationCustomData `json:"customData,omitempty"`
}

type ReportChargingProfilesConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value ReportChargingProfilesConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *ReportChargingProfilesConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (ReportChargingProfilesConfirmation) ActionName() string { return "ReportChargingProfiles" }

func (ReportChargingProfilesConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (ReportChargingProfilesConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (ReportChargingProfilesConfirmation) SchemaName() string {
	return "ReportChargingProfilesResponse.json"
}

func (message ReportChargingProfilesConfirmation) Validate() error {
	return validation.Validate("ReportChargingProfilesResponse.json", schemaReportChargingProfilesConfirmation, message)
}

func (ReportChargingProfilesConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ReportChargingProfilesResponse.json", schemaReportChargingProfilesConfirmation, data)
}
