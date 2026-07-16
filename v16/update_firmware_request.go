// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = UpdateFirmwareRequest{}

var schemaUpdateFirmwareRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"location": &validation.Schema{Type: "string", AllowAdditional: true, Format: "uri"}, "retries": &validation.Schema{Type: "integer", AllowAdditional: true}, "retrieveDate": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "retryInterval": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"location", "retrieveDate"}}

type UpdateFirmwareRequest struct {
	Location      string `json:"location"`
	Retries       *int   `json:"retries,omitempty"`
	RetrieveDate  string `json:"retrieveDate"`
	RetryInterval *int   `json:"retryInterval,omitempty"`
}

func (UpdateFirmwareRequest) ActionName() string { return "UpdateFirmware" }

func (UpdateFirmwareRequest) Version() protocol.Version { return protocol.OCPP16 }

func (UpdateFirmwareRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (UpdateFirmwareRequest) SchemaName() string { return "UpdateFirmware.json" }

func (message UpdateFirmwareRequest) Validate() error {
	return validation.Validate("UpdateFirmware.json", schemaUpdateFirmwareRequest, message)
}

func (UpdateFirmwareRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("UpdateFirmware.json", schemaUpdateFirmwareRequest, data)
}
