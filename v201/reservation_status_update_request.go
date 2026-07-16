// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ReservationStatusUpdateRequest{}

var schemaReservationStatusUpdateRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reservationId": &validation.Schema{Type: "integer", AllowAdditional: true}, "reservationUpdateStatus": &validation.Schema{Type: "string", Enum: []string{"Expired", "Removed"}}}, Required: []string{"reservationId", "reservationUpdateStatus"}}

type ReservationStatusUpdateRequest struct {
	CustomData              *ReservationStatusUpdateRequestCustomData                 `json:"customData,omitempty"`
	ReservationID           int                                                       `json:"reservationId"`
	ReservationUpdateStatus ReservationStatusUpdateRequestReservationUpdateStatusEnum `json:"reservationUpdateStatus"`
}

type ReservationStatusUpdateRequestReservationUpdateStatusEnum string

const (
	ReservationStatusUpdateRequestReservationUpdateStatusEnumExpired ReservationStatusUpdateRequestReservationUpdateStatusEnum = "Expired"
	ReservationStatusUpdateRequestReservationUpdateStatusEnumRemoved ReservationStatusUpdateRequestReservationUpdateStatusEnum = "Removed"
)

type ReservationStatusUpdateRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (ReservationStatusUpdateRequest) ActionName() string { return "ReservationStatusUpdate" }

func (ReservationStatusUpdateRequest) Version() protocol.Version { return protocol.OCPP201 }

func (ReservationStatusUpdateRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (ReservationStatusUpdateRequest) SchemaName() string {
	return "ReservationStatusUpdateRequest.json"
}

func (message ReservationStatusUpdateRequest) Validate() error {
	return validation.Validate("ReservationStatusUpdateRequest.json", schemaReservationStatusUpdateRequest, message)
}

func (ReservationStatusUpdateRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ReservationStatusUpdateRequest.json", schemaReservationStatusUpdateRequest, data)
}
