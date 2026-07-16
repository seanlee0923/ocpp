// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ReservationStatusUpdateConfirmation{}

var schemaReservationStatusUpdateConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type ReservationStatusUpdateConfirmation struct {
	CustomData *ReservationStatusUpdateConfirmationCustomData `json:"customData,omitempty"`
}

type ReservationStatusUpdateConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value ReservationStatusUpdateConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *ReservationStatusUpdateConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (ReservationStatusUpdateConfirmation) ActionName() string { return "ReservationStatusUpdate" }

func (ReservationStatusUpdateConfirmation) Version() protocol.Version { return protocol.OCPP201 }

func (ReservationStatusUpdateConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (ReservationStatusUpdateConfirmation) SchemaName() string {
	return "ReservationStatusUpdateResponse.json"
}

func (message ReservationStatusUpdateConfirmation) Validate() error {
	return validation.Validate("ReservationStatusUpdateResponse.json", schemaReservationStatusUpdateConfirmation, message)
}

func (ReservationStatusUpdateConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ReservationStatusUpdateResponse.json", schemaReservationStatusUpdateConfirmation, data)
}
