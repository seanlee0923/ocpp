// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = NotifyPeriodicEventStreamRequest{}

var schemaNotifyPeriodicEventStreamRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"data": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"t": &validation.Schema{Type: "number", AllowAdditional: true}, "v": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 2500, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"t", "v"}}, MinItems: 1, HasMinItems: true}, "id": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "pending": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "basetime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id", "pending", "basetime", "data"}}

type NotifyPeriodicEventStreamRequest struct {
	Data       []NotifyPeriodicEventStreamRequestStreamDataElement `json:"data"`
	ID         int                                                 `json:"id"`
	Pending    int                                                 `json:"pending"`
	Basetime   string                                              `json:"basetime"`
	CustomData *NotifyPeriodicEventStreamRequestCustomData         `json:"customData,omitempty"`
}

type NotifyPeriodicEventStreamRequestStreamDataElement struct {
	T          float64                                     `json:"t"`
	V          string                                      `json:"v"`
	CustomData *NotifyPeriodicEventStreamRequestCustomData `json:"customData,omitempty"`
}

type NotifyPeriodicEventStreamRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (NotifyPeriodicEventStreamRequest) ActionName() string { return "NotifyPeriodicEventStream" }

func (NotifyPeriodicEventStreamRequest) Version() protocol.Version { return protocol.OCPP21 }

func (NotifyPeriodicEventStreamRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (NotifyPeriodicEventStreamRequest) SchemaName() string { return "NotifyPeriodicEventStream.json" }

func (message NotifyPeriodicEventStreamRequest) Validate() error {
	return validation.Validate("NotifyPeriodicEventStream.json", schemaNotifyPeriodicEventStreamRequest, message)
}

func (NotifyPeriodicEventStreamRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("NotifyPeriodicEventStream.json", schemaNotifyPeriodicEventStreamRequest, data)
}
