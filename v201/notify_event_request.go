// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = NotifyEventRequest{}

var schemaNotifyEventRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "generatedAt": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "tbc": &validation.Schema{Type: "boolean", AllowAdditional: true}, "seqNo": &validation.Schema{Type: "integer", AllowAdditional: true}, "eventData": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "eventId": &validation.Schema{Type: "integer", AllowAdditional: true}, "timestamp": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "trigger": &validation.Schema{Type: "string", Enum: []string{"Alerting", "Delta", "Periodic"}}, "cause": &validation.Schema{Type: "integer", AllowAdditional: true}, "actualValue": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 2500, HasMaxLength: true}, "techCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "techInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 500, HasMaxLength: true}, "cleared": &validation.Schema{Type: "boolean", AllowAdditional: true}, "transactionId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "component": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "evse": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "id": &validation.Schema{Type: "integer", AllowAdditional: true}, "connectorId": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"id"}}, "name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}}, Required: []string{"name"}}, "variableMonitoringId": &validation.Schema{Type: "integer", AllowAdditional: true}, "eventNotificationType": &validation.Schema{Type: "string", Enum: []string{"HardWiredNotification", "HardWiredMonitor", "PreconfiguredMonitor", "CustomMonitor"}}, "variable": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "instance": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}}, Required: []string{"name"}}}, Required: []string{"eventId", "timestamp", "trigger", "actualValue", "eventNotificationType", "component", "variable"}}, MinItems: 1, HasMinItems: true}}, Required: []string{"generatedAt", "seqNo", "eventData"}}

type NotifyEventRequest struct {
	CustomData  *NotifyEventRequestCustomData `json:"customData,omitempty"`
	GeneratedAt string                        `json:"generatedAt"`
	Tbc         *bool                         `json:"tbc,omitempty"`
	SeqNo       int                           `json:"seqNo"`
	EventData   []NotifyEventRequestEventData `json:"eventData"`
}

type NotifyEventRequestEventData struct {
	CustomData            *NotifyEventRequestCustomData           `json:"customData,omitempty"`
	EventID               int                                     `json:"eventId"`
	Timestamp             string                                  `json:"timestamp"`
	Trigger               NotifyEventRequestEventTriggerEnum      `json:"trigger"`
	Cause                 *int                                    `json:"cause,omitempty"`
	ActualValue           string                                  `json:"actualValue"`
	TechCode              *string                                 `json:"techCode,omitempty"`
	TechInfo              *string                                 `json:"techInfo,omitempty"`
	Cleared               *bool                                   `json:"cleared,omitempty"`
	TransactionID         *string                                 `json:"transactionId,omitempty"`
	Component             NotifyEventRequestComponent             `json:"component"`
	VariableMonitoringID  *int                                    `json:"variableMonitoringId,omitempty"`
	EventNotificationType NotifyEventRequestEventNotificationEnum `json:"eventNotificationType"`
	Variable              NotifyEventRequestVariable              `json:"variable"`
}

type NotifyEventRequestVariable struct {
	CustomData *NotifyEventRequestCustomData `json:"customData,omitempty"`
	Name       string                        `json:"name"`
	Instance   *string                       `json:"instance,omitempty"`
}

type NotifyEventRequestEventNotificationEnum string

const (
	NotifyEventRequestEventNotificationEnumHardWiredNotification NotifyEventRequestEventNotificationEnum = "HardWiredNotification"
	NotifyEventRequestEventNotificationEnumHardWiredMonitor      NotifyEventRequestEventNotificationEnum = "HardWiredMonitor"
	NotifyEventRequestEventNotificationEnumPreconfiguredMonitor  NotifyEventRequestEventNotificationEnum = "PreconfiguredMonitor"
	NotifyEventRequestEventNotificationEnumCustomMonitor         NotifyEventRequestEventNotificationEnum = "CustomMonitor"
)

type NotifyEventRequestComponent struct {
	CustomData *NotifyEventRequestCustomData `json:"customData,omitempty"`
	EVSE       *NotifyEventRequestEVSE       `json:"evse,omitempty"`
	Name       string                        `json:"name"`
	Instance   *string                       `json:"instance,omitempty"`
}

type NotifyEventRequestEVSE struct {
	CustomData  *NotifyEventRequestCustomData `json:"customData,omitempty"`
	ID          int                           `json:"id"`
	ConnectorID *int                          `json:"connectorId,omitempty"`
}

type NotifyEventRequestEventTriggerEnum string

const (
	NotifyEventRequestEventTriggerEnumAlerting NotifyEventRequestEventTriggerEnum = "Alerting"
	NotifyEventRequestEventTriggerEnumDelta    NotifyEventRequestEventTriggerEnum = "Delta"
	NotifyEventRequestEventTriggerEnumPeriodic NotifyEventRequestEventTriggerEnum = "Periodic"
)

type NotifyEventRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value NotifyEventRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *NotifyEventRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (NotifyEventRequest) ActionName() string { return "NotifyEvent" }

func (NotifyEventRequest) Version() protocol.Version { return protocol.OCPP201 }

func (NotifyEventRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (NotifyEventRequest) SchemaName() string { return "NotifyEventRequest.json" }

func (message NotifyEventRequest) Validate() error {
	return validation.Validate("NotifyEventRequest.json", schemaNotifyEventRequest, message)
}

func (NotifyEventRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("NotifyEventRequest.json", schemaNotifyEventRequest, data)
}
