// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ReportChargingProfilesConfirmation{}

var schemaReportChargingProfilesConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type ReportChargingProfilesConfirmation struct {
	CustomData *ReportChargingProfilesConfirmationCustomData `json:"customData,omitempty"`
}

type ReportChargingProfilesConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (ReportChargingProfilesConfirmation) ActionName() string { return "ReportChargingProfiles" }

func (ReportChargingProfilesConfirmation) Version() protocol.Version { return protocol.OCPP201 }

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
