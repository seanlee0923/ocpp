// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = AFRRSignalRequest{}

var schemaAFRRSignalRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"timestamp": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "signal": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"timestamp", "signal"}}

type AFRRSignalRequest struct {
	Timestamp  string                       `json:"timestamp"`
	Signal     int                          `json:"signal"`
	CustomData *AFRRSignalRequestCustomData `json:"customData,omitempty"`
}

type AFRRSignalRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (AFRRSignalRequest) ActionName() string { return "AFRRSignal" }

func (AFRRSignalRequest) Version() protocol.Version { return protocol.OCPP21 }

func (AFRRSignalRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (AFRRSignalRequest) SchemaName() string { return "AFRRSignalRequest.json" }

func (message AFRRSignalRequest) Validate() error {
	return validation.Validate("AFRRSignalRequest.json", schemaAFRRSignalRequest, message)
}

func (AFRRSignalRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("AFRRSignalRequest.json", schemaAFRRSignalRequest, data)
}
