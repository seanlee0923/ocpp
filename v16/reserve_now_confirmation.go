// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ReserveNowConfirmation{}

var schemaReserveNowConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Faulted", "Occupied", "Rejected", "Unavailable"}}}, Required: []string{"status"}}

type ReserveNowConfirmation struct {
	Status ReserveNowConfirmationStatus `json:"status"`
}

type ReserveNowConfirmationStatus string

const (
	ReserveNowConfirmationStatusAccepted    ReserveNowConfirmationStatus = "Accepted"
	ReserveNowConfirmationStatusFaulted     ReserveNowConfirmationStatus = "Faulted"
	ReserveNowConfirmationStatusOccupied    ReserveNowConfirmationStatus = "Occupied"
	ReserveNowConfirmationStatusRejected    ReserveNowConfirmationStatus = "Rejected"
	ReserveNowConfirmationStatusUnavailable ReserveNowConfirmationStatus = "Unavailable"
)

func (ReserveNowConfirmation) ActionName() string { return "ReserveNow" }

func (ReserveNowConfirmation) Version() protocol.Version { return protocol.OCPP16 }

func (ReserveNowConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (ReserveNowConfirmation) SchemaName() string { return "ReserveNowResponse.json" }

func (message ReserveNowConfirmation) Validate() error {
	return validation.Validate("ReserveNowResponse.json", schemaReserveNowConfirmation, message)
}

func (ReserveNowConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ReserveNowResponse.json", schemaReserveNowConfirmation, data)
}
