// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = StartTransactionRequest{}

var schemaStartTransactionRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"connectorId": &validation.Schema{Type: "integer", AllowAdditional: true}, "idTag": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "meterStart": &validation.Schema{Type: "integer", AllowAdditional: true}, "reservationId": &validation.Schema{Type: "integer", AllowAdditional: true}, "timestamp": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}}, Required: []string{"connectorId", "idTag", "meterStart", "timestamp"}}

type StartTransactionRequest struct {
	ConnectorID   int    `json:"connectorId"`
	IDTag         string `json:"idTag"`
	MeterStart    int    `json:"meterStart"`
	ReservationID *int   `json:"reservationId,omitempty"`
	Timestamp     string `json:"timestamp"`
}

func (StartTransactionRequest) ActionName() string { return "StartTransaction" }

func (StartTransactionRequest) Version() protocol.Version { return protocol.OCPP16 }

func (StartTransactionRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (StartTransactionRequest) SchemaName() string { return "StartTransaction.json" }

func (message StartTransactionRequest) Validate() error {
	return validation.Validate("StartTransaction.json", schemaStartTransactionRequest, message)
}

func (StartTransactionRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("StartTransaction.json", schemaStartTransactionRequest, data)
}
