// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = CancelReservationRequest{}

var schemaCancelReservationRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reservationId": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"reservationId"}}

type CancelReservationRequest struct {
	ReservationID int `json:"reservationId"`
}

func (CancelReservationRequest) ActionName() string { return "CancelReservation" }

func (CancelReservationRequest) Version() protocol.Version { return protocol.OCPP16 }

func (CancelReservationRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (CancelReservationRequest) SchemaName() string { return "CancelReservation.json" }

func (message CancelReservationRequest) Validate() error {
	return validation.Validate("CancelReservation.json", schemaCancelReservationRequest, message)
}

func (CancelReservationRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("CancelReservation.json", schemaCancelReservationRequest, data)
}
