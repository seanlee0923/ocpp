// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = AdjustPeriodicEventStreamRequest{}

var schemaAdjustPeriodicEventStreamRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"id": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "params": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"interval": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "values": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id", "params"}}

type AdjustPeriodicEventStreamRequest struct {
	ID         int                                                       `json:"id"`
	Params     AdjustPeriodicEventStreamRequestPeriodicEventStreamParams `json:"params"`
	CustomData *AdjustPeriodicEventStreamRequestCustomData               `json:"customData,omitempty"`
}

type AdjustPeriodicEventStreamRequestPeriodicEventStreamParams struct {
	Interval   *int                                        `json:"interval,omitempty"`
	Values     *int                                        `json:"values,omitempty"`
	CustomData *AdjustPeriodicEventStreamRequestCustomData `json:"customData,omitempty"`
}

type AdjustPeriodicEventStreamRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (AdjustPeriodicEventStreamRequest) ActionName() string { return "AdjustPeriodicEventStream" }

func (AdjustPeriodicEventStreamRequest) Version() protocol.Version { return protocol.OCPP21 }

func (AdjustPeriodicEventStreamRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (AdjustPeriodicEventStreamRequest) SchemaName() string {
	return "AdjustPeriodicEventStreamRequest.json"
}

func (message AdjustPeriodicEventStreamRequest) Validate() error {
	return validation.Validate("AdjustPeriodicEventStreamRequest.json", schemaAdjustPeriodicEventStreamRequest, message)
}

func (AdjustPeriodicEventStreamRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("AdjustPeriodicEventStreamRequest.json", schemaAdjustPeriodicEventStreamRequest, data)
}
