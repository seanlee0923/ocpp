// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = DataTransferConfirmation{}

var schemaDataTransferConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "UnknownMessageId", "UnknownVendorId"}}, "data": &validation.Schema{Type: "string", AllowAdditional: true}}, Required: []string{"status"}}

type DataTransferConfirmation struct {
	Status DataTransferConfirmationStatus `json:"status"`
	Data   *string                        `json:"data,omitempty"`
}

type DataTransferConfirmationStatus string

const (
	DataTransferConfirmationStatusAccepted         DataTransferConfirmationStatus = "Accepted"
	DataTransferConfirmationStatusRejected         DataTransferConfirmationStatus = "Rejected"
	DataTransferConfirmationStatusUnknownMessageID DataTransferConfirmationStatus = "UnknownMessageId"
	DataTransferConfirmationStatusUnknownVendorID  DataTransferConfirmationStatus = "UnknownVendorId"
)

func (DataTransferConfirmation) ActionName() string { return "DataTransfer" }

func (DataTransferConfirmation) Version() protocol.Version { return protocol.OCPP16 }

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
