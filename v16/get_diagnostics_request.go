// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = GetDiagnosticsRequest{}

var schemaGetDiagnosticsRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"location": &validation.Schema{Type: "string", AllowAdditional: true, Format: "uri"}, "retries": &validation.Schema{Type: "integer", AllowAdditional: true}, "retryInterval": &validation.Schema{Type: "integer", AllowAdditional: true}, "startTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "stopTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}}, Required: []string{"location"}}

type GetDiagnosticsRequest struct {
	Location      string  `json:"location"`
	Retries       *int    `json:"retries,omitempty"`
	RetryInterval *int    `json:"retryInterval,omitempty"`
	StartTime     *string `json:"startTime,omitempty"`
	StopTime      *string `json:"stopTime,omitempty"`
}

func (GetDiagnosticsRequest) ActionName() string { return "GetDiagnostics" }

func (GetDiagnosticsRequest) Version() protocol.Version { return protocol.OCPP16 }

func (GetDiagnosticsRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (GetDiagnosticsRequest) SchemaName() string { return "GetDiagnostics.json" }

func (message GetDiagnosticsRequest) Validate() error {
	return validation.Validate("GetDiagnostics.json", schemaGetDiagnosticsRequest, message)
}

func (GetDiagnosticsRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetDiagnostics.json", schemaGetDiagnosticsRequest, data)
}
