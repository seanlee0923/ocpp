// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = GetLogRequest{}

var schemaGetLogRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "log": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "remoteLocation": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}, "oldestTimestamp": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "latestTimestamp": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}}, Required: []string{"remoteLocation"}}, "logType": &validation.Schema{Type: "string", Enum: []string{"DiagnosticsLog", "SecurityLog"}}, "requestId": &validation.Schema{Type: "integer", AllowAdditional: true}, "retries": &validation.Schema{Type: "integer", AllowAdditional: true}, "retryInterval": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"logType", "requestId", "log"}}

type GetLogRequest struct {
	CustomData    *GetLogRequestCustomData   `json:"customData,omitempty"`
	Log           GetLogRequestLogParameters `json:"log"`
	LogType       GetLogRequestLogEnum       `json:"logType"`
	RequestID     int                        `json:"requestId"`
	Retries       *int                       `json:"retries,omitempty"`
	RetryInterval *int                       `json:"retryInterval,omitempty"`
}

type GetLogRequestLogEnum string

const (
	GetLogRequestLogEnumDiagnosticsLog GetLogRequestLogEnum = "DiagnosticsLog"
	GetLogRequestLogEnumSecurityLog    GetLogRequestLogEnum = "SecurityLog"
)

type GetLogRequestLogParameters struct {
	CustomData      *GetLogRequestCustomData `json:"customData,omitempty"`
	RemoteLocation  string                   `json:"remoteLocation"`
	OldestTimestamp *string                  `json:"oldestTimestamp,omitempty"`
	LatestTimestamp *string                  `json:"latestTimestamp,omitempty"`
}

type GetLogRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value GetLogRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *GetLogRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (GetLogRequest) ActionName() string { return "GetLog" }

func (GetLogRequest) Version() protocol.Version { return protocol.OCPP201 }

func (GetLogRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (GetLogRequest) SchemaName() string { return "GetLogRequest.json" }

func (message GetLogRequest) Validate() error {
	return validation.Validate("GetLogRequest.json", schemaGetLogRequest, message)
}

func (GetLogRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetLogRequest.json", schemaGetLogRequest, data)
}
