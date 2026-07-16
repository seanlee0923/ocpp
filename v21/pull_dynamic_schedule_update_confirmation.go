// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = PullDynamicScheduleUpdateConfirmation{}

var schemaPullDynamicScheduleUpdateConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"scheduleUpdate": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"limit": &validation.Schema{Type: "number", AllowAdditional: true}, "limit_L2": &validation.Schema{Type: "number", AllowAdditional: true}, "limit_L3": &validation.Schema{Type: "number", AllowAdditional: true}, "dischargeLimit": &validation.Schema{Type: "number", AllowAdditional: true, Maximum: 0, HasMaximum: true}, "dischargeLimit_L2": &validation.Schema{Type: "number", AllowAdditional: true, Maximum: 0, HasMaximum: true}, "dischargeLimit_L3": &validation.Schema{Type: "number", AllowAdditional: true, Maximum: 0, HasMaximum: true}, "setpoint": &validation.Schema{Type: "number", AllowAdditional: true}, "setpoint_L2": &validation.Schema{Type: "number", AllowAdditional: true}, "setpoint_L3": &validation.Schema{Type: "number", AllowAdditional: true}, "setpointReactive": &validation.Schema{Type: "number", AllowAdditional: true}, "setpointReactive_L2": &validation.Schema{Type: "number", AllowAdditional: true}, "setpointReactive_L3": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type PullDynamicScheduleUpdateConfirmation struct {
	ScheduleUpdate *PullDynamicScheduleUpdateConfirmationChargingScheduleUpdate   `json:"scheduleUpdate,omitempty"`
	Status         PullDynamicScheduleUpdateConfirmationChargingProfileStatusEnum `json:"status"`
	StatusInfo     *PullDynamicScheduleUpdateConfirmationStatusInfo               `json:"statusInfo,omitempty"`
	CustomData     *PullDynamicScheduleUpdateConfirmationCustomData               `json:"customData,omitempty"`
}

type PullDynamicScheduleUpdateConfirmationStatusInfo struct {
	ReasonCode     string                                           `json:"reasonCode"`
	AdditionalInfo *string                                          `json:"additionalInfo,omitempty"`
	CustomData     *PullDynamicScheduleUpdateConfirmationCustomData `json:"customData,omitempty"`
}

type PullDynamicScheduleUpdateConfirmationChargingProfileStatusEnum string

const (
	PullDynamicScheduleUpdateConfirmationChargingProfileStatusEnumAccepted PullDynamicScheduleUpdateConfirmationChargingProfileStatusEnum = "Accepted"
	PullDynamicScheduleUpdateConfirmationChargingProfileStatusEnumRejected PullDynamicScheduleUpdateConfirmationChargingProfileStatusEnum = "Rejected"
)

type PullDynamicScheduleUpdateConfirmationChargingScheduleUpdate struct {
	Limit              *float64                                         `json:"limit,omitempty"`
	LimitL2            *float64                                         `json:"limit_L2,omitempty"`
	LimitL3            *float64                                         `json:"limit_L3,omitempty"`
	DischargeLimit     *float64                                         `json:"dischargeLimit,omitempty"`
	DischargeLimitL2   *float64                                         `json:"dischargeLimit_L2,omitempty"`
	DischargeLimitL3   *float64                                         `json:"dischargeLimit_L3,omitempty"`
	Setpoint           *float64                                         `json:"setpoint,omitempty"`
	SetpointL2         *float64                                         `json:"setpoint_L2,omitempty"`
	SetpointL3         *float64                                         `json:"setpoint_L3,omitempty"`
	SetpointReactive   *float64                                         `json:"setpointReactive,omitempty"`
	SetpointReactiveL2 *float64                                         `json:"setpointReactive_L2,omitempty"`
	SetpointReactiveL3 *float64                                         `json:"setpointReactive_L3,omitempty"`
	CustomData         *PullDynamicScheduleUpdateConfirmationCustomData `json:"customData,omitempty"`
}

type PullDynamicScheduleUpdateConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value PullDynamicScheduleUpdateConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *PullDynamicScheduleUpdateConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (PullDynamicScheduleUpdateConfirmation) ActionName() string { return "PullDynamicScheduleUpdate" }

func (PullDynamicScheduleUpdateConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (PullDynamicScheduleUpdateConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (PullDynamicScheduleUpdateConfirmation) SchemaName() string {
	return "PullDynamicScheduleUpdateResponse.json"
}

func (message PullDynamicScheduleUpdateConfirmation) Validate() error {
	return validation.Validate("PullDynamicScheduleUpdateResponse.json", schemaPullDynamicScheduleUpdateConfirmation, message)
}

func (PullDynamicScheduleUpdateConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("PullDynamicScheduleUpdateResponse.json", schemaPullDynamicScheduleUpdateConfirmation, data)
}
