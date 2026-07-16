// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = GetBaseReportConfirmation{}

var schemaGetBaseReportConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "NotSupported", "EmptyResultSet"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type GetBaseReportConfirmation struct {
	Status     GetBaseReportConfirmationGenericDeviceModelStatusEnum `json:"status"`
	StatusInfo *GetBaseReportConfirmationStatusInfo                  `json:"statusInfo,omitempty"`
	CustomData *GetBaseReportConfirmationCustomData                  `json:"customData,omitempty"`
}

type GetBaseReportConfirmationStatusInfo struct {
	ReasonCode     string                               `json:"reasonCode"`
	AdditionalInfo *string                              `json:"additionalInfo,omitempty"`
	CustomData     *GetBaseReportConfirmationCustomData `json:"customData,omitempty"`
}

type GetBaseReportConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

type GetBaseReportConfirmationGenericDeviceModelStatusEnum string

const (
	GetBaseReportConfirmationGenericDeviceModelStatusEnumAccepted       GetBaseReportConfirmationGenericDeviceModelStatusEnum = "Accepted"
	GetBaseReportConfirmationGenericDeviceModelStatusEnumRejected       GetBaseReportConfirmationGenericDeviceModelStatusEnum = "Rejected"
	GetBaseReportConfirmationGenericDeviceModelStatusEnumNotSupported   GetBaseReportConfirmationGenericDeviceModelStatusEnum = "NotSupported"
	GetBaseReportConfirmationGenericDeviceModelStatusEnumEmptyResultSet GetBaseReportConfirmationGenericDeviceModelStatusEnum = "EmptyResultSet"
)

func (GetBaseReportConfirmation) ActionName() string { return "GetBaseReport" }

func (GetBaseReportConfirmation) Version() protocol.Version { return protocol.OCPP21 }

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
