// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = GetLogRequest{}

var schemaGetLogRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"log": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"remoteLocation": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 2000, HasMaxLength: true}, "oldestTimestamp": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "latestTimestamp": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"remoteLocation"}}, "logType": &validation.Schema{Type: "string", Enum: []string{"DiagnosticsLog", "SecurityLog", "DataCollectorLog"}}, "requestId": &validation.Schema{Type: "integer", AllowAdditional: true}, "retries": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "retryInterval": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"logType", "requestId", "log"}}

type GetLogRequest struct {
	Log           GetLogRequestLogParameters `json:"log"`
	LogType       GetLogRequestLogEnum       `json:"logType"`
	RequestID     int                        `json:"requestId"`
	Retries       *int                       `json:"retries,omitempty"`
	RetryInterval *int                       `json:"retryInterval,omitempty"`
	CustomData    *GetLogRequestCustomData   `json:"customData,omitempty"`
}

type GetLogRequestLogEnum string

const (
	GetLogRequestLogEnumDiagnosticsLog   GetLogRequestLogEnum = "DiagnosticsLog"
	GetLogRequestLogEnumSecurityLog      GetLogRequestLogEnum = "SecurityLog"
	GetLogRequestLogEnumDataCollectorLog GetLogRequestLogEnum = "DataCollectorLog"
)

type GetLogRequestLogParameters struct {
	RemoteLocation  string                   `json:"remoteLocation"`
	OldestTimestamp *string                  `json:"oldestTimestamp,omitempty"`
	LatestTimestamp *string                  `json:"latestTimestamp,omitempty"`
	CustomData      *GetLogRequestCustomData `json:"customData,omitempty"`
}

type GetLogRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (GetLogRequest) ActionName() string { return "GetLog" }

func (GetLogRequest) Version() protocol.Version { return protocol.OCPP21 }

func (GetLogRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (GetLogRequest) SchemaName() string { return "GetLogRequest.json" }

func (message GetLogRequest) Validate() error {
	return validation.Validate("GetLogRequest.json", schemaGetLogRequest, message)
}

func (GetLogRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetLogRequest.json", schemaGetLogRequest, data)
}
