// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = GetReportConfirmation{}

var schemaGetReportConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "NotSupported", "EmptyResultSet"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}}, Required: []string{"status"}}

type GetReportConfirmation struct {
	CustomData *GetReportConfirmationCustomData                  `json:"customData,omitempty"`
	Status     GetReportConfirmationGenericDeviceModelStatusEnum `json:"status"`
	StatusInfo *GetReportConfirmationStatusInfo                  `json:"statusInfo,omitempty"`
}

type GetReportConfirmationStatusInfo struct {
	CustomData     *GetReportConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                           `json:"reasonCode"`
	AdditionalInfo *string                          `json:"additionalInfo,omitempty"`
}

type GetReportConfirmationGenericDeviceModelStatusEnum string

const (
	GetReportConfirmationGenericDeviceModelStatusEnumAccepted       GetReportConfirmationGenericDeviceModelStatusEnum = "Accepted"
	GetReportConfirmationGenericDeviceModelStatusEnumRejected       GetReportConfirmationGenericDeviceModelStatusEnum = "Rejected"
	GetReportConfirmationGenericDeviceModelStatusEnumNotSupported   GetReportConfirmationGenericDeviceModelStatusEnum = "NotSupported"
	GetReportConfirmationGenericDeviceModelStatusEnumEmptyResultSet GetReportConfirmationGenericDeviceModelStatusEnum = "EmptyResultSet"
)

type GetReportConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (GetReportConfirmation) ActionName() string { return "GetReport" }

func (GetReportConfirmation) Version() protocol.Version { return protocol.OCPP201 }

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
