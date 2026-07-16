// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = SetDisplayMessageRequest{}

var schemaSetDisplayMessageRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "message": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "display": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "evse": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "id": &validation.Schema{Type: "integer", AllowAdditional: true}, "connectorId": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"id"}}, "name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}}, Required: []string{"name"}}, "id": &validation.Schema{Type: "integer", AllowAdditional: true}, "priority": &validation.Schema{Type: "string", Enum: []string{"AlwaysFront", "InFront", "NormalCycle"}}, "state": &validation.Schema{Type: "string", Enum: []string{"Charging", "Faulted", "Idle", "Unavailable"}}, "startDateTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "endDateTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "transactionId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "message": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "format": &validation.Schema{Type: "string", Enum: []string{"ASCII", "HTML", "URI", "UTF8"}}, "language": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 8, HasMaxLength: true}, "content": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"format", "content"}}}, Required: []string{"id", "priority", "message"}}}, Required: []string{"message"}}

type SetDisplayMessageRequest struct {
	CustomData *SetDisplayMessageRequestCustomData `json:"customData,omitempty"`
	Message    SetDisplayMessageRequestMessageInfo `json:"message"`
}

type SetDisplayMessageRequestMessageInfo struct {
	CustomData    *SetDisplayMessageRequestCustomData         `json:"customData,omitempty"`
	Display       *SetDisplayMessageRequestComponent          `json:"display,omitempty"`
	ID            int                                         `json:"id"`
	Priority      SetDisplayMessageRequestMessagePriorityEnum `json:"priority"`
	State         *SetDisplayMessageRequestMessageStateEnum   `json:"state,omitempty"`
	StartDateTime *string                                     `json:"startDateTime,omitempty"`
	EndDateTime   *string                                     `json:"endDateTime,omitempty"`
	TransactionID *string                                     `json:"transactionId,omitempty"`
	Message       SetDisplayMessageRequestMessageContent      `json:"message"`
}

type SetDisplayMessageRequestMessageContent struct {
	CustomData *SetDisplayMessageRequestCustomData       `json:"customData,omitempty"`
	Format     SetDisplayMessageRequestMessageFormatEnum `json:"format"`
	Language   *string                                   `json:"language,omitempty"`
	Content    string                                    `json:"content"`
}

type SetDisplayMessageRequestMessageFormatEnum string

const (
	SetDisplayMessageRequestMessageFormatEnumASCII SetDisplayMessageRequestMessageFormatEnum = "ASCII"
	SetDisplayMessageRequestMessageFormatEnumHTML  SetDisplayMessageRequestMessageFormatEnum = "HTML"
	SetDisplayMessageRequestMessageFormatEnumURI   SetDisplayMessageRequestMessageFormatEnum = "URI"
	SetDisplayMessageRequestMessageFormatEnumUTF8  SetDisplayMessageRequestMessageFormatEnum = "UTF8"
)

type SetDisplayMessageRequestMessageStateEnum string

const (
	SetDisplayMessageRequestMessageStateEnumCharging    SetDisplayMessageRequestMessageStateEnum = "Charging"
	SetDisplayMessageRequestMessageStateEnumFaulted     SetDisplayMessageRequestMessageStateEnum = "Faulted"
	SetDisplayMessageRequestMessageStateEnumIdle        SetDisplayMessageRequestMessageStateEnum = "Idle"
	SetDisplayMessageRequestMessageStateEnumUnavailable SetDisplayMessageRequestMessageStateEnum = "Unavailable"
)

type SetDisplayMessageRequestMessagePriorityEnum string

const (
	SetDisplayMessageRequestMessagePriorityEnumAlwaysFront SetDisplayMessageRequestMessagePriorityEnum = "AlwaysFront"
	SetDisplayMessageRequestMessagePriorityEnumInFront     SetDisplayMessageRequestMessagePriorityEnum = "InFront"
	SetDisplayMessageRequestMessagePriorityEnumNormalCycle SetDisplayMessageRequestMessagePriorityEnum = "NormalCycle"
)

type SetDisplayMessageRequestComponent struct {
	CustomData *SetDisplayMessageRequestCustomData `json:"customData,omitempty"`
	EVSE       *SetDisplayMessageRequestEVSE       `json:"evse,omitempty"`
	Name       string                              `json:"name"`
	Instance   *string                             `json:"instance,omitempty"`
}

type SetDisplayMessageRequestEVSE struct {
	CustomData  *SetDisplayMessageRequestCustomData `json:"customData,omitempty"`
	ID          int                                 `json:"id"`
	ConnectorID *int                                `json:"connectorId,omitempty"`
}

type SetDisplayMessageRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (SetDisplayMessageRequest) ActionName() string { return "SetDisplayMessage" }

func (SetDisplayMessageRequest) Version() protocol.Version { return protocol.OCPP201 }

func (SetDisplayMessageRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (SetDisplayMessageRequest) SchemaName() string { return "SetDisplayMessageRequest.json" }

func (message SetDisplayMessageRequest) Validate() error {
	return validation.Validate("SetDisplayMessageRequest.json", schemaSetDisplayMessageRequest, message)
}

func (SetDisplayMessageRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("SetDisplayMessageRequest.json", schemaSetDisplayMessageRequest, data)
}
