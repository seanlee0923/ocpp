// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = UpdateDynamicScheduleRequest{}

var schemaUpdateDynamicScheduleRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"chargingProfileId": &validation.Schema{Type: "integer", AllowAdditional: true}, "scheduleUpdate": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"limit": &validation.Schema{Type: "number", AllowAdditional: true}, "limit_L2": &validation.Schema{Type: "number", AllowAdditional: true}, "limit_L3": &validation.Schema{Type: "number", AllowAdditional: true}, "dischargeLimit": &validation.Schema{Type: "number", AllowAdditional: true, Maximum: 0, HasMaximum: true}, "dischargeLimit_L2": &validation.Schema{Type: "number", AllowAdditional: true, Maximum: 0, HasMaximum: true}, "dischargeLimit_L3": &validation.Schema{Type: "number", AllowAdditional: true, Maximum: 0, HasMaximum: true}, "setpoint": &validation.Schema{Type: "number", AllowAdditional: true}, "setpoint_L2": &validation.Schema{Type: "number", AllowAdditional: true}, "setpoint_L3": &validation.Schema{Type: "number", AllowAdditional: true}, "setpointReactive": &validation.Schema{Type: "number", AllowAdditional: true}, "setpointReactive_L2": &validation.Schema{Type: "number", AllowAdditional: true}, "setpointReactive_L3": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"chargingProfileId", "scheduleUpdate"}}

type UpdateDynamicScheduleRequest struct {
	ChargingProfileID int                                                `json:"chargingProfileId"`
	ScheduleUpdate    UpdateDynamicScheduleRequestChargingScheduleUpdate `json:"scheduleUpdate"`
	CustomData        *UpdateDynamicScheduleRequestCustomData            `json:"customData,omitempty"`
}

type UpdateDynamicScheduleRequestChargingScheduleUpdate struct {
	Limit              *float64                                `json:"limit,omitempty"`
	LimitL2            *float64                                `json:"limit_L2,omitempty"`
	LimitL3            *float64                                `json:"limit_L3,omitempty"`
	DischargeLimit     *float64                                `json:"dischargeLimit,omitempty"`
	DischargeLimitL2   *float64                                `json:"dischargeLimit_L2,omitempty"`
	DischargeLimitL3   *float64                                `json:"dischargeLimit_L3,omitempty"`
	Setpoint           *float64                                `json:"setpoint,omitempty"`
	SetpointL2         *float64                                `json:"setpoint_L2,omitempty"`
	SetpointL3         *float64                                `json:"setpoint_L3,omitempty"`
	SetpointReactive   *float64                                `json:"setpointReactive,omitempty"`
	SetpointReactiveL2 *float64                                `json:"setpointReactive_L2,omitempty"`
	SetpointReactiveL3 *float64                                `json:"setpointReactive_L3,omitempty"`
	CustomData         *UpdateDynamicScheduleRequestCustomData `json:"customData,omitempty"`
}

type UpdateDynamicScheduleRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (UpdateDynamicScheduleRequest) ActionName() string { return "UpdateDynamicSchedule" }

func (UpdateDynamicScheduleRequest) Version() protocol.Version { return protocol.OCPP21 }

func (UpdateDynamicScheduleRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (UpdateDynamicScheduleRequest) SchemaName() string { return "UpdateDynamicScheduleRequest.json" }

func (message UpdateDynamicScheduleRequest) Validate() error {
	return validation.Validate("UpdateDynamicScheduleRequest.json", schemaUpdateDynamicScheduleRequest, message)
}

func (UpdateDynamicScheduleRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("UpdateDynamicScheduleRequest.json", schemaUpdateDynamicScheduleRequest, data)
}
