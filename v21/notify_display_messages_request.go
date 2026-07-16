// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = NotifyDisplayMessagesRequest{}

var schemaNotifyDisplayMessagesRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"messageInfo": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"display": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"evse": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"id": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "connectorId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id"}}, "name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"name"}}, "id": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "priority": &validation.Schema{Type: "string", Enum: []string{"AlwaysFront", "InFront", "NormalCycle"}}, "state": &validation.Schema{Type: "string", Enum: []string{"Charging", "Faulted", "Idle", "Unavailable", "Suspended", "Discharging"}}, "startDateTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "endDateTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "transactionId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "message": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"format": &validation.Schema{Type: "string", Enum: []string{"ASCII", "HTML", "URI", "UTF8", "QRCODE"}}, "language": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 8, HasMaxLength: true}, "content": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"format", "content"}}, "messageExtra": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"format": &validation.Schema{Type: "string", Enum: []string{"ASCII", "HTML", "URI", "UTF8", "QRCODE"}}, "language": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 8, HasMaxLength: true}, "content": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"format", "content"}}, MinItems: 1, HasMinItems: true, MaxItems: 4, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id", "priority", "message"}}, MinItems: 1, HasMinItems: true}, "requestId": &validation.Schema{Type: "integer", AllowAdditional: true}, "tbc": &validation.Schema{Type: "boolean", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"requestId"}}

type NotifyDisplayMessagesRequest struct {
	MessageInfo []NotifyDisplayMessagesRequestMessageInfo `json:"messageInfo,omitempty"`
	RequestID   int                                       `json:"requestId"`
	Tbc         *bool                                     `json:"tbc,omitempty"`
	CustomData  *NotifyDisplayMessagesRequestCustomData   `json:"customData,omitempty"`
}

type NotifyDisplayMessagesRequestMessageInfo struct {
	Display       *NotifyDisplayMessagesRequestComponent          `json:"display,omitempty"`
	ID            int                                             `json:"id"`
	Priority      NotifyDisplayMessagesRequestMessagePriorityEnum `json:"priority"`
	State         *NotifyDisplayMessagesRequestMessageStateEnum   `json:"state,omitempty"`
	StartDateTime *string                                         `json:"startDateTime,omitempty"`
	EndDateTime   *string                                         `json:"endDateTime,omitempty"`
	TransactionID *string                                         `json:"transactionId,omitempty"`
	Message       NotifyDisplayMessagesRequestMessageContent      `json:"message"`
	MessageExtra  []NotifyDisplayMessagesRequestMessageContent    `json:"messageExtra,omitempty"`
	CustomData    *NotifyDisplayMessagesRequestCustomData         `json:"customData,omitempty"`
}

type NotifyDisplayMessagesRequestMessageContent struct {
	Format     NotifyDisplayMessagesRequestMessageFormatEnum `json:"format"`
	Language   *string                                       `json:"language,omitempty"`
	Content    string                                        `json:"content"`
	CustomData *NotifyDisplayMessagesRequestCustomData       `json:"customData,omitempty"`
}

type NotifyDisplayMessagesRequestMessageFormatEnum string

const (
	NotifyDisplayMessagesRequestMessageFormatEnumASCII  NotifyDisplayMessagesRequestMessageFormatEnum = "ASCII"
	NotifyDisplayMessagesRequestMessageFormatEnumHTML   NotifyDisplayMessagesRequestMessageFormatEnum = "HTML"
	NotifyDisplayMessagesRequestMessageFormatEnumURI    NotifyDisplayMessagesRequestMessageFormatEnum = "URI"
	NotifyDisplayMessagesRequestMessageFormatEnumUTF8   NotifyDisplayMessagesRequestMessageFormatEnum = "UTF8"
	NotifyDisplayMessagesRequestMessageFormatEnumQRCODE NotifyDisplayMessagesRequestMessageFormatEnum = "QRCODE"
)

type NotifyDisplayMessagesRequestMessageStateEnum string

const (
	NotifyDisplayMessagesRequestMessageStateEnumCharging    NotifyDisplayMessagesRequestMessageStateEnum = "Charging"
	NotifyDisplayMessagesRequestMessageStateEnumFaulted     NotifyDisplayMessagesRequestMessageStateEnum = "Faulted"
	NotifyDisplayMessagesRequestMessageStateEnumIdle        NotifyDisplayMessagesRequestMessageStateEnum = "Idle"
	NotifyDisplayMessagesRequestMessageStateEnumUnavailable NotifyDisplayMessagesRequestMessageStateEnum = "Unavailable"
	NotifyDisplayMessagesRequestMessageStateEnumSuspended   NotifyDisplayMessagesRequestMessageStateEnum = "Suspended"
	NotifyDisplayMessagesRequestMessageStateEnumDischarging NotifyDisplayMessagesRequestMessageStateEnum = "Discharging"
)

type NotifyDisplayMessagesRequestMessagePriorityEnum string

const (
	NotifyDisplayMessagesRequestMessagePriorityEnumAlwaysFront NotifyDisplayMessagesRequestMessagePriorityEnum = "AlwaysFront"
	NotifyDisplayMessagesRequestMessagePriorityEnumInFront     NotifyDisplayMessagesRequestMessagePriorityEnum = "InFront"
	NotifyDisplayMessagesRequestMessagePriorityEnumNormalCycle NotifyDisplayMessagesRequestMessagePriorityEnum = "NormalCycle"
)

type NotifyDisplayMessagesRequestComponent struct {
	EVSE       *NotifyDisplayMessagesRequestEVSE       `json:"evse,omitempty"`
	Name       string                                  `json:"name"`
	Instance   *string                                 `json:"instance,omitempty"`
	CustomData *NotifyDisplayMessagesRequestCustomData `json:"customData,omitempty"`
}

type NotifyDisplayMessagesRequestEVSE struct {
	ID          int                                     `json:"id"`
	ConnectorID *int                                    `json:"connectorId,omitempty"`
	CustomData  *NotifyDisplayMessagesRequestCustomData `json:"customData,omitempty"`
}

type NotifyDisplayMessagesRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value NotifyDisplayMessagesRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *NotifyDisplayMessagesRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (NotifyDisplayMessagesRequest) ActionName() string { return "NotifyDisplayMessages" }

func (NotifyDisplayMessagesRequest) Version() protocol.Version { return protocol.OCPP21 }

func (NotifyDisplayMessagesRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (NotifyDisplayMessagesRequest) SchemaName() string { return "NotifyDisplayMessagesRequest.json" }

func (message NotifyDisplayMessagesRequest) Validate() error {
	return validation.Validate("NotifyDisplayMessagesRequest.json", schemaNotifyDisplayMessagesRequest, message)
}

func (NotifyDisplayMessagesRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("NotifyDisplayMessagesRequest.json", schemaNotifyDisplayMessagesRequest, data)
}
