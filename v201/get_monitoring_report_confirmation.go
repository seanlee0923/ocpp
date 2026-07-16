// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = GetMonitoringReportConfirmation{}

var schemaGetMonitoringReportConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "NotSupported", "EmptyResultSet"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}}, Required: []string{"status"}}

type GetMonitoringReportConfirmation struct {
	CustomData *GetMonitoringReportConfirmationCustomData                  `json:"customData,omitempty"`
	Status     GetMonitoringReportConfirmationGenericDeviceModelStatusEnum `json:"status"`
	StatusInfo *GetMonitoringReportConfirmationStatusInfo                  `json:"statusInfo,omitempty"`
}

type GetMonitoringReportConfirmationStatusInfo struct {
	CustomData     *GetMonitoringReportConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                                     `json:"reasonCode"`
	AdditionalInfo *string                                    `json:"additionalInfo,omitempty"`
}

type GetMonitoringReportConfirmationGenericDeviceModelStatusEnum string

const (
	GetMonitoringReportConfirmationGenericDeviceModelStatusEnumAccepted       GetMonitoringReportConfirmationGenericDeviceModelStatusEnum = "Accepted"
	GetMonitoringReportConfirmationGenericDeviceModelStatusEnumRejected       GetMonitoringReportConfirmationGenericDeviceModelStatusEnum = "Rejected"
	GetMonitoringReportConfirmationGenericDeviceModelStatusEnumNotSupported   GetMonitoringReportConfirmationGenericDeviceModelStatusEnum = "NotSupported"
	GetMonitoringReportConfirmationGenericDeviceModelStatusEnumEmptyResultSet GetMonitoringReportConfirmationGenericDeviceModelStatusEnum = "EmptyResultSet"
)

type GetMonitoringReportConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value GetMonitoringReportConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *GetMonitoringReportConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (GetMonitoringReportConfirmation) ActionName() string { return "GetMonitoringReport" }

func (GetMonitoringReportConfirmation) Version() protocol.Version { return protocol.OCPP201 }

func (GetMonitoringReportConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (GetMonitoringReportConfirmation) SchemaName() string { return "GetMonitoringReportResponse.json" }

func (message GetMonitoringReportConfirmation) Validate() error {
	return validation.Validate("GetMonitoringReportResponse.json", schemaGetMonitoringReportConfirmation, message)
}

func (GetMonitoringReportConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetMonitoringReportResponse.json", schemaGetMonitoringReportConfirmation, data)
}
