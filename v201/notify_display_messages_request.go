// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = NotifyDisplayMessagesRequest{}

var schemaNotifyDisplayMessagesRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "messageInfo": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "display": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "evse": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "id": &validation.Schema{Type: "integer", AllowAdditional: true}, "connectorId": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"id"}}, "name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}}, Required: []string{"name"}}, "id": &validation.Schema{Type: "integer", AllowAdditional: true}, "priority": &validation.Schema{Type: "string", Enum: []string{"AlwaysFront", "InFront", "NormalCycle"}}, "state": &validation.Schema{Type: "string", Enum: []string{"Charging", "Faulted", "Idle", "Unavailable"}}, "startDateTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "endDateTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "transactionId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "message": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "format": &validation.Schema{Type: "string", Enum: []string{"ASCII", "HTML", "URI", "UTF8"}}, "language": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 8, HasMaxLength: true}, "content": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"format", "content"}}}, Required: []string{"id", "priority", "message"}}, MinItems: 1, HasMinItems: true}, "requestId": &validation.Schema{Type: "integer", AllowAdditional: true}, "tbc": &validation.Schema{Type: "boolean", AllowAdditional: true}}, Required: []string{"requestId"}}

type NotifyDisplayMessagesRequest struct {
	CustomData  *NotifyDisplayMessagesRequestCustomData   `json:"customData,omitempty"`
	MessageInfo []NotifyDisplayMessagesRequestMessageInfo `json:"messageInfo,omitempty"`
	RequestID   int                                       `json:"requestId"`
	Tbc         *bool                                     `json:"tbc,omitempty"`
}

type NotifyDisplayMessagesRequestMessageInfo struct {
	CustomData    *NotifyDisplayMessagesRequestCustomData         `json:"customData,omitempty"`
	Display       *NotifyDisplayMessagesRequestComponent          `json:"display,omitempty"`
	ID            int                                             `json:"id"`
	Priority      NotifyDisplayMessagesRequestMessagePriorityEnum `json:"priority"`
	State         *NotifyDisplayMessagesRequestMessageStateEnum   `json:"state,omitempty"`
	StartDateTime *string                                         `json:"startDateTime,omitempty"`
	EndDateTime   *string                                         `json:"endDateTime,omitempty"`
	TransactionID *string                                         `json:"transactionId,omitempty"`
	Message       NotifyDisplayMessagesRequestMessageContent      `json:"message"`
}

type NotifyDisplayMessagesRequestMessageContent struct {
	CustomData *NotifyDisplayMessagesRequestCustomData       `json:"customData,omitempty"`
	Format     NotifyDisplayMessagesRequestMessageFormatEnum `json:"format"`
	Language   *string                                       `json:"language,omitempty"`
	Content    string                                        `json:"content"`
}

type NotifyDisplayMessagesRequestMessageFormatEnum string

const (
	NotifyDisplayMessagesRequestMessageFormatEnumASCII NotifyDisplayMessagesRequestMessageFormatEnum = "ASCII"
	NotifyDisplayMessagesRequestMessageFormatEnumHTML  NotifyDisplayMessagesRequestMessageFormatEnum = "HTML"
	NotifyDisplayMessagesRequestMessageFormatEnumURI   NotifyDisplayMessagesRequestMessageFormatEnum = "URI"
	NotifyDisplayMessagesRequestMessageFormatEnumUTF8  NotifyDisplayMessagesRequestMessageFormatEnum = "UTF8"
)

type NotifyDisplayMessagesRequestMessageStateEnum string

const (
	NotifyDisplayMessagesRequestMessageStateEnumCharging    NotifyDisplayMessagesRequestMessageStateEnum = "Charging"
	NotifyDisplayMessagesRequestMessageStateEnumFaulted     NotifyDisplayMessagesRequestMessageStateEnum = "Faulted"
	NotifyDisplayMessagesRequestMessageStateEnumIdle        NotifyDisplayMessagesRequestMessageStateEnum = "Idle"
	NotifyDisplayMessagesRequestMessageStateEnumUnavailable NotifyDisplayMessagesRequestMessageStateEnum = "Unavailable"
)

type NotifyDisplayMessagesRequestMessagePriorityEnum string

const (
	NotifyDisplayMessagesRequestMessagePriorityEnumAlwaysFront NotifyDisplayMessagesRequestMessagePriorityEnum = "AlwaysFront"
	NotifyDisplayMessagesRequestMessagePriorityEnumInFront     NotifyDisplayMessagesRequestMessagePriorityEnum = "InFront"
	NotifyDisplayMessagesRequestMessagePriorityEnumNormalCycle NotifyDisplayMessagesRequestMessagePriorityEnum = "NormalCycle"
)

type NotifyDisplayMessagesRequestComponent struct {
	CustomData *NotifyDisplayMessagesRequestCustomData `json:"customData,omitempty"`
	EVSE       *NotifyDisplayMessagesRequestEVSE       `json:"evse,omitempty"`
	Name       string                                  `json:"name"`
	Instance   *string                                 `json:"instance,omitempty"`
}

type NotifyDisplayMessagesRequestEVSE struct {
	CustomData  *NotifyDisplayMessagesRequestCustomData `json:"customData,omitempty"`
	ID          int                                     `json:"id"`
	ConnectorID *int                                    `json:"connectorId,omitempty"`
}

type NotifyDisplayMessagesRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (NotifyDisplayMessagesRequest) ActionName() string { return "NotifyDisplayMessages" }

func (NotifyDisplayMessagesRequest) Version() protocol.Version { return protocol.OCPP201 }

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
