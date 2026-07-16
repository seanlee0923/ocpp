// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = GetDisplayMessagesRequest{}

var schemaGetDisplayMessagesRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"id": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, MinItems: 1, HasMinItems: true}, "requestId": &validation.Schema{Type: "integer", AllowAdditional: true}, "priority": &validation.Schema{Type: "string", Enum: []string{"AlwaysFront", "InFront", "NormalCycle"}}, "state": &validation.Schema{Type: "string", Enum: []string{"Charging", "Faulted", "Idle", "Unavailable", "Suspended", "Discharging"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"requestId"}}

type GetDisplayMessagesRequest struct {
	ID         []int                                         `json:"id,omitempty"`
	RequestID  int                                           `json:"requestId"`
	Priority   *GetDisplayMessagesRequestMessagePriorityEnum `json:"priority,omitempty"`
	State      *GetDisplayMessagesRequestMessageStateEnum    `json:"state,omitempty"`
	CustomData *GetDisplayMessagesRequestCustomData          `json:"customData,omitempty"`
}

type GetDisplayMessagesRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

type GetDisplayMessagesRequestMessageStateEnum string

const (
	GetDisplayMessagesRequestMessageStateEnumCharging    GetDisplayMessagesRequestMessageStateEnum = "Charging"
	GetDisplayMessagesRequestMessageStateEnumFaulted     GetDisplayMessagesRequestMessageStateEnum = "Faulted"
	GetDisplayMessagesRequestMessageStateEnumIdle        GetDisplayMessagesRequestMessageStateEnum = "Idle"
	GetDisplayMessagesRequestMessageStateEnumUnavailable GetDisplayMessagesRequestMessageStateEnum = "Unavailable"
	GetDisplayMessagesRequestMessageStateEnumSuspended   GetDisplayMessagesRequestMessageStateEnum = "Suspended"
	GetDisplayMessagesRequestMessageStateEnumDischarging GetDisplayMessagesRequestMessageStateEnum = "Discharging"
)

type GetDisplayMessagesRequestMessagePriorityEnum string

const (
	GetDisplayMessagesRequestMessagePriorityEnumAlwaysFront GetDisplayMessagesRequestMessagePriorityEnum = "AlwaysFront"
	GetDisplayMessagesRequestMessagePriorityEnumInFront     GetDisplayMessagesRequestMessagePriorityEnum = "InFront"
	GetDisplayMessagesRequestMessagePriorityEnumNormalCycle GetDisplayMessagesRequestMessagePriorityEnum = "NormalCycle"
)

func (GetDisplayMessagesRequest) ActionName() string { return "GetDisplayMessages" }

func (GetDisplayMessagesRequest) Version() protocol.Version { return protocol.OCPP21 }

func (GetDisplayMessagesRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (GetDisplayMessagesRequest) SchemaName() string { return "GetDisplayMessagesRequest.json" }

func (message GetDisplayMessagesRequest) Validate() error {
	return validation.Validate("GetDisplayMessagesRequest.json", schemaGetDisplayMessagesRequest, message)
}

func (GetDisplayMessagesRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetDisplayMessagesRequest.json", schemaGetDisplayMessagesRequest, data)
}
