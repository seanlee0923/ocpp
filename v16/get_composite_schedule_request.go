// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = GetCompositeScheduleRequest{}

var schemaGetCompositeScheduleRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"connectorId": &validation.Schema{Type: "integer", AllowAdditional: true}, "duration": &validation.Schema{Type: "integer", AllowAdditional: true}, "chargingRateUnit": &validation.Schema{Type: "string", Enum: []string{"A", "W"}}}, Required: []string{"connectorId", "duration"}}

type GetCompositeScheduleRequest struct {
	ConnectorID      int                                          `json:"connectorId"`
	Duration         int                                          `json:"duration"`
	ChargingRateUnit *GetCompositeScheduleRequestChargingRateUnit `json:"chargingRateUnit,omitempty"`
}

type GetCompositeScheduleRequestChargingRateUnit string

const (
	GetCompositeScheduleRequestChargingRateUnitA GetCompositeScheduleRequestChargingRateUnit = "A"
	GetCompositeScheduleRequestChargingRateUnitW GetCompositeScheduleRequestChargingRateUnit = "W"
)

func (GetCompositeScheduleRequest) ActionName() string { return "GetCompositeSchedule" }

func (GetCompositeScheduleRequest) Version() protocol.Version { return protocol.OCPP16 }

func (GetCompositeScheduleRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (GetCompositeScheduleRequest) SchemaName() string { return "GetCompositeSchedule.json" }

func (message GetCompositeScheduleRequest) Validate() error {
	return validation.Validate("GetCompositeSchedule.json", schemaGetCompositeScheduleRequest, message)
}

func (GetCompositeScheduleRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetCompositeSchedule.json", schemaGetCompositeScheduleRequest, data)
}
