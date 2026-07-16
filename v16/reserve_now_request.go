// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ReserveNowRequest{}

var schemaReserveNowRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"connectorId": &validation.Schema{Type: "integer", AllowAdditional: true}, "expiryDate": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "idTag": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "parentIdTag": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "reservationId": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"connectorId", "expiryDate", "idTag", "reservationId"}}

type ReserveNowRequest struct {
	ConnectorID   int     `json:"connectorId"`
	ExpiryDate    string  `json:"expiryDate"`
	IDTag         string  `json:"idTag"`
	ParentIDTag   *string `json:"parentIdTag,omitempty"`
	ReservationID int     `json:"reservationId"`
}

func (ReserveNowRequest) ActionName() string { return "ReserveNow" }

func (ReserveNowRequest) Version() protocol.Version { return protocol.OCPP16 }

func (ReserveNowRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (ReserveNowRequest) SchemaName() string { return "ReserveNow.json" }

func (message ReserveNowRequest) Validate() error {
	return validation.Validate("ReserveNow.json", schemaReserveNowRequest, message)
}

func (ReserveNowRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ReserveNow.json", schemaReserveNowRequest, data)
}
