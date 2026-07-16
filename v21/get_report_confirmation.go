// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = GetReportConfirmation{}

var schemaGetReportConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "NotSupported", "EmptyResultSet"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type GetReportConfirmation struct {
	Status     GetReportConfirmationGenericDeviceModelStatusEnum `json:"status"`
	StatusInfo *GetReportConfirmationStatusInfo                  `json:"statusInfo,omitempty"`
	CustomData *GetReportConfirmationCustomData                  `json:"customData,omitempty"`
}

type GetReportConfirmationStatusInfo struct {
	ReasonCode     string                           `json:"reasonCode"`
	AdditionalInfo *string                          `json:"additionalInfo,omitempty"`
	CustomData     *GetReportConfirmationCustomData `json:"customData,omitempty"`
}

type GetReportConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

type GetReportConfirmationGenericDeviceModelStatusEnum string

const (
	GetReportConfirmationGenericDeviceModelStatusEnumAccepted       GetReportConfirmationGenericDeviceModelStatusEnum = "Accepted"
	GetReportConfirmationGenericDeviceModelStatusEnumRejected       GetReportConfirmationGenericDeviceModelStatusEnum = "Rejected"
	GetReportConfirmationGenericDeviceModelStatusEnumNotSupported   GetReportConfirmationGenericDeviceModelStatusEnum = "NotSupported"
	GetReportConfirmationGenericDeviceModelStatusEnumEmptyResultSet GetReportConfirmationGenericDeviceModelStatusEnum = "EmptyResultSet"
)

func (GetReportConfirmation) ActionName() string { return "GetReport" }

func (GetReportConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (GetReportConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (GetReportConfirmation) SchemaName() string { return "GetReportResponse.json" }

func (message GetReportConfirmation) Validate() error {
	return validation.Validate("GetReportResponse.json", schemaGetReportConfirmation, message)
}

func (GetReportConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetReportResponse.json", schemaGetReportConfirmation, data)
}
