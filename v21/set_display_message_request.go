// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = SetDisplayMessageRequest{}

var schemaSetDisplayMessageRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"message": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"display": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"evse": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"id": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "connectorId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id"}}, "name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"name"}}, "id": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "priority": &validation.Schema{Type: "string", Enum: []string{"AlwaysFront", "InFront", "NormalCycle"}}, "state": &validation.Schema{Type: "string", Enum: []string{"Charging", "Faulted", "Idle", "Unavailable", "Suspended", "Discharging"}}, "startDateTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "endDateTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "transactionId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "message": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"format": &validation.Schema{Type: "string", Enum: []string{"ASCII", "HTML", "URI", "UTF8", "QRCODE"}}, "language": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 8, HasMaxLength: true}, "content": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"format", "content"}}, "messageExtra": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"format": &validation.Schema{Type: "string", Enum: []string{"ASCII", "HTML", "URI", "UTF8", "QRCODE"}}, "language": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 8, HasMaxLength: true}, "content": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"format", "content"}}, MinItems: 1, HasMinItems: true, MaxItems: 4, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id", "priority", "message"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"message"}}

type SetDisplayMessageRequest struct {
	Message    SetDisplayMessageRequestMessageInfo `json:"message"`
	CustomData *SetDisplayMessageRequestCustomData `json:"customData,omitempty"`
}

type SetDisplayMessageRequestMessageInfo struct {
	Display       *SetDisplayMessageRequestComponent          `json:"display,omitempty"`
	ID            int                                         `json:"id"`
	Priority      SetDisplayMessageRequestMessagePriorityEnum `json:"priority"`
	State         *SetDisplayMessageRequestMessageStateEnum   `json:"state,omitempty"`
	StartDateTime *string                                     `json:"startDateTime,omitempty"`
	EndDateTime   *string                                     `json:"endDateTime,omitempty"`
	TransactionID *string                                     `json:"transactionId,omitempty"`
	Message       SetDisplayMessageRequestMessageContent      `json:"message"`
	MessageExtra  []SetDisplayMessageRequestMessageContent    `json:"messageExtra,omitempty"`
	CustomData    *SetDisplayMessageRequestCustomData         `json:"customData,omitempty"`
}

type SetDisplayMessageRequestMessageContent struct {
	Format     SetDisplayMessageRequestMessageFormatEnum `json:"format"`
	Language   *string                                   `json:"language,omitempty"`
	Content    string                                    `json:"content"`
	CustomData *SetDisplayMessageRequestCustomData       `json:"customData,omitempty"`
}

type SetDisplayMessageRequestMessageFormatEnum string

const (
	SetDisplayMessageRequestMessageFormatEnumASCII  SetDisplayMessageRequestMessageFormatEnum = "ASCII"
	SetDisplayMessageRequestMessageFormatEnumHTML   SetDisplayMessageRequestMessageFormatEnum = "HTML"
	SetDisplayMessageRequestMessageFormatEnumURI    SetDisplayMessageRequestMessageFormatEnum = "URI"
	SetDisplayMessageRequestMessageFormatEnumUTF8   SetDisplayMessageRequestMessageFormatEnum = "UTF8"
	SetDisplayMessageRequestMessageFormatEnumQRCODE SetDisplayMessageRequestMessageFormatEnum = "QRCODE"
)

type SetDisplayMessageRequestMessageStateEnum string

const (
	SetDisplayMessageRequestMessageStateEnumCharging    SetDisplayMessageRequestMessageStateEnum = "Charging"
	SetDisplayMessageRequestMessageStateEnumFaulted     SetDisplayMessageRequestMessageStateEnum = "Faulted"
	SetDisplayMessageRequestMessageStateEnumIdle        SetDisplayMessageRequestMessageStateEnum = "Idle"
	SetDisplayMessageRequestMessageStateEnumUnavailable SetDisplayMessageRequestMessageStateEnum = "Unavailable"
	SetDisplayMessageRequestMessageStateEnumSuspended   SetDisplayMessageRequestMessageStateEnum = "Suspended"
	SetDisplayMessageRequestMessageStateEnumDischarging SetDisplayMessageRequestMessageStateEnum = "Discharging"
)

type SetDisplayMessageRequestMessagePriorityEnum string

const (
	SetDisplayMessageRequestMessagePriorityEnumAlwaysFront SetDisplayMessageRequestMessagePriorityEnum = "AlwaysFront"
	SetDisplayMessageRequestMessagePriorityEnumInFront     SetDisplayMessageRequestMessagePriorityEnum = "InFront"
	SetDisplayMessageRequestMessagePriorityEnumNormalCycle SetDisplayMessageRequestMessagePriorityEnum = "NormalCycle"
)

type SetDisplayMessageRequestComponent struct {
	EVSE       *SetDisplayMessageRequestEVSE       `json:"evse,omitempty"`
	Name       string                              `json:"name"`
	Instance   *string                             `json:"instance,omitempty"`
	CustomData *SetDisplayMessageRequestCustomData `json:"customData,omitempty"`
}

type SetDisplayMessageRequestEVSE struct {
	ID          int                                 `json:"id"`
	ConnectorID *int                                `json:"connectorId,omitempty"`
	CustomData  *SetDisplayMessageRequestCustomData `json:"customData,omitempty"`
}

type SetDisplayMessageRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value SetDisplayMessageRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *SetDisplayMessageRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (SetDisplayMessageRequest) ActionName() string { return "SetDisplayMessage" }

func (SetDisplayMessageRequest) Version() protocol.Version { return protocol.OCPP21 }

func (SetDisplayMessageRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (SetDisplayMessageRequest) SchemaName() string { return "SetDisplayMessageRequest.json" }

func (message SetDisplayMessageRequest) Validate() error {
	return validation.Validate("SetDisplayMessageRequest.json", schemaSetDisplayMessageRequest, message)
}

func (SetDisplayMessageRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("SetDisplayMessageRequest.json", schemaSetDisplayMessageRequest, data)
}
