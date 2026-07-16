// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = GetDisplayMessagesRequest{}

var schemaGetDisplayMessagesRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "id": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "integer", AllowAdditional: true}, MinItems: 1, HasMinItems: true}, "requestId": &validation.Schema{Type: "integer", AllowAdditional: true}, "priority": &validation.Schema{Type: "string", Enum: []string{"AlwaysFront", "InFront", "NormalCycle"}}, "state": &validation.Schema{Type: "string", Enum: []string{"Charging", "Faulted", "Idle", "Unavailable"}}}, Required: []string{"requestId"}}

type GetDisplayMessagesRequest struct {
	CustomData *GetDisplayMessagesRequestCustomData          `json:"customData,omitempty"`
	ID         []int                                         `json:"id,omitempty"`
	RequestID  int                                           `json:"requestId"`
	Priority   *GetDisplayMessagesRequestMessagePriorityEnum `json:"priority,omitempty"`
	State      *GetDisplayMessagesRequestMessageStateEnum    `json:"state,omitempty"`
}

type GetDisplayMessagesRequestMessageStateEnum string

const (
	GetDisplayMessagesRequestMessageStateEnumCharging    GetDisplayMessagesRequestMessageStateEnum = "Charging"
	GetDisplayMessagesRequestMessageStateEnumFaulted     GetDisplayMessagesRequestMessageStateEnum = "Faulted"
	GetDisplayMessagesRequestMessageStateEnumIdle        GetDisplayMessagesRequestMessageStateEnum = "Idle"
	GetDisplayMessagesRequestMessageStateEnumUnavailable GetDisplayMessagesRequestMessageStateEnum = "Unavailable"
)

type GetDisplayMessagesRequestMessagePriorityEnum string

const (
	GetDisplayMessagesRequestMessagePriorityEnumAlwaysFront GetDisplayMessagesRequestMessagePriorityEnum = "AlwaysFront"
	GetDisplayMessagesRequestMessagePriorityEnumInFront     GetDisplayMessagesRequestMessagePriorityEnum = "InFront"
	GetDisplayMessagesRequestMessagePriorityEnumNormalCycle GetDisplayMessagesRequestMessagePriorityEnum = "NormalCycle"
)

type GetDisplayMessagesRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value GetDisplayMessagesRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *GetDisplayMessagesRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (GetDisplayMessagesRequest) ActionName() string { return "GetDisplayMessages" }

func (GetDisplayMessagesRequest) Version() protocol.Version { return protocol.OCPP201 }

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
