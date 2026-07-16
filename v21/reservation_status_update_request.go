// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ReservationStatusUpdateRequest{}

var schemaReservationStatusUpdateRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reservationId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "reservationUpdateStatus": &validation.Schema{Type: "string", Enum: []string{"Expired", "Removed", "NoTransaction"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reservationId", "reservationUpdateStatus"}}

type ReservationStatusUpdateRequest struct {
	ReservationID           int                                                       `json:"reservationId"`
	ReservationUpdateStatus ReservationStatusUpdateRequestReservationUpdateStatusEnum `json:"reservationUpdateStatus"`
	CustomData              *ReservationStatusUpdateRequestCustomData                 `json:"customData,omitempty"`
}

type ReservationStatusUpdateRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value ReservationStatusUpdateRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *ReservationStatusUpdateRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

type ReservationStatusUpdateRequestReservationUpdateStatusEnum string

const (
	ReservationStatusUpdateRequestReservationUpdateStatusEnumExpired       ReservationStatusUpdateRequestReservationUpdateStatusEnum = "Expired"
	ReservationStatusUpdateRequestReservationUpdateStatusEnumRemoved       ReservationStatusUpdateRequestReservationUpdateStatusEnum = "Removed"
	ReservationStatusUpdateRequestReservationUpdateStatusEnumNoTransaction ReservationStatusUpdateRequestReservationUpdateStatusEnum = "NoTransaction"
)

func (ReservationStatusUpdateRequest) ActionName() string { return "ReservationStatusUpdate" }

func (ReservationStatusUpdateRequest) Version() protocol.Version { return protocol.OCPP21 }

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
