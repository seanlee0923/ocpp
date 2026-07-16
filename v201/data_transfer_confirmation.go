// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = DataTransferConfirmation{}

var schemaDataTransferConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "UnknownMessageId", "UnknownVendorId"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}, "data": &validation.Schema{AllowAdditional: true}}, Required: []string{"status"}}

type DataTransferConfirmation struct {
	CustomData *DataTransferConfirmationCustomData            `json:"customData,omitempty"`
	Status     DataTransferConfirmationDataTransferStatusEnum `json:"status"`
	StatusInfo *DataTransferConfirmationStatusInfo            `json:"statusInfo,omitempty"`
	Data       any                                            `json:"data,omitempty"`
}

type DataTransferConfirmationStatusInfo struct {
	CustomData     *DataTransferConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                              `json:"reasonCode"`
	AdditionalInfo *string                             `json:"additionalInfo,omitempty"`
}

type DataTransferConfirmationDataTransferStatusEnum string

const (
	DataTransferConfirmationDataTransferStatusEnumAccepted         DataTransferConfirmationDataTransferStatusEnum = "Accepted"
	DataTransferConfirmationDataTransferStatusEnumRejected         DataTransferConfirmationDataTransferStatusEnum = "Rejected"
	DataTransferConfirmationDataTransferStatusEnumUnknownMessageID DataTransferConfirmationDataTransferStatusEnum = "UnknownMessageId"
	DataTransferConfirmationDataTransferStatusEnumUnknownVendorID  DataTransferConfirmationDataTransferStatusEnum = "UnknownVendorId"
)

type DataTransferConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (DataTransferConfirmation) ActionName() string { return "DataTransfer" }

func (DataTransferConfirmation) Version() protocol.Version { return protocol.OCPP201 }

func (DataTransferConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (DataTransferConfirmation) SchemaName() string { return "DataTransferResponse.json" }

func (message DataTransferConfirmation) Validate() error {
	return validation.Validate("DataTransferResponse.json", schemaDataTransferConfirmation, message)
}

func (DataTransferConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("DataTransferResponse.json", schemaDataTransferConfirmation, data)
}
