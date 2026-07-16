// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = GetCompositeScheduleRequest{}

var schemaGetCompositeScheduleRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"duration": &validation.Schema{Type: "integer", AllowAdditional: true}, "chargingRateUnit": &validation.Schema{Type: "string", Enum: []string{"W", "A"}}, "evseId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"duration", "evseId"}}

type GetCompositeScheduleRequest struct {
	Duration         int                                              `json:"duration"`
	ChargingRateUnit *GetCompositeScheduleRequestChargingRateUnitEnum `json:"chargingRateUnit,omitempty"`
	EVSEID           int                                              `json:"evseId"`
	CustomData       *GetCompositeScheduleRequestCustomData           `json:"customData,omitempty"`
}

type GetCompositeScheduleRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value GetCompositeScheduleRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *GetCompositeScheduleRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

type GetCompositeScheduleRequestChargingRateUnitEnum string

const (
	GetCompositeScheduleRequestChargingRateUnitEnumW GetCompositeScheduleRequestChargingRateUnitEnum = "W"
	GetCompositeScheduleRequestChargingRateUnitEnumA GetCompositeScheduleRequestChargingRateUnitEnum = "A"
)

func (GetCompositeScheduleRequest) ActionName() string { return "GetCompositeSchedule" }

func (GetCompositeScheduleRequest) Version() protocol.Version { return protocol.OCPP21 }

func (GetCompositeScheduleRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (GetCompositeScheduleRequest) SchemaName() string { return "GetCompositeScheduleRequest.json" }

func (message GetCompositeScheduleRequest) Validate() error {
	return validation.Validate("GetCompositeScheduleRequest.json", schemaGetCompositeScheduleRequest, message)
}

func (GetCompositeScheduleRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetCompositeScheduleRequest.json", schemaGetCompositeScheduleRequest, data)
}
