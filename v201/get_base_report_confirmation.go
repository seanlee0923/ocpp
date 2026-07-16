// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = GetBaseReportConfirmation{}

var schemaGetBaseReportConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "NotSupported", "EmptyResultSet"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}}, Required: []string{"status"}}

type GetBaseReportConfirmation struct {
	CustomData *GetBaseReportConfirmationCustomData                  `json:"customData,omitempty"`
	Status     GetBaseReportConfirmationGenericDeviceModelStatusEnum `json:"status"`
	StatusInfo *GetBaseReportConfirmationStatusInfo                  `json:"statusInfo,omitempty"`
}

type GetBaseReportConfirmationStatusInfo struct {
	CustomData     *GetBaseReportConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                               `json:"reasonCode"`
	AdditionalInfo *string                              `json:"additionalInfo,omitempty"`
}

type GetBaseReportConfirmationGenericDeviceModelStatusEnum string

const (
	GetBaseReportConfirmationGenericDeviceModelStatusEnumAccepted       GetBaseReportConfirmationGenericDeviceModelStatusEnum = "Accepted"
	GetBaseReportConfirmationGenericDeviceModelStatusEnumRejected       GetBaseReportConfirmationGenericDeviceModelStatusEnum = "Rejected"
	GetBaseReportConfirmationGenericDeviceModelStatusEnumNotSupported   GetBaseReportConfirmationGenericDeviceModelStatusEnum = "NotSupported"
	GetBaseReportConfirmationGenericDeviceModelStatusEnumEmptyResultSet GetBaseReportConfirmationGenericDeviceModelStatusEnum = "EmptyResultSet"
)

type GetBaseReportConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value GetBaseReportConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *GetBaseReportConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (GetBaseReportConfirmation) ActionName() string { return "GetBaseReport" }

func (GetBaseReportConfirmation) Version() protocol.Version { return protocol.OCPP201 }

func (GetBaseReportConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (GetBaseReportConfirmation) SchemaName() string { return "GetBaseReportResponse.json" }

func (message GetBaseReportConfirmation) Validate() error {
	return validation.Validate("GetBaseReportResponse.json", schemaGetBaseReportConfirmation, message)
}

func (GetBaseReportConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetBaseReportResponse.json", schemaGetBaseReportConfirmation, data)
}
