// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ClearVariableMonitoringRequest{}

var schemaClearVariableMonitoringRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"id": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, MinItems: 1, HasMinItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id"}}

type ClearVariableMonitoringRequest struct {
	ID         []int                                     `json:"id"`
	CustomData *ClearVariableMonitoringRequestCustomData `json:"customData,omitempty"`
}

type ClearVariableMonitoringRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (ClearVariableMonitoringRequest) ActionName() string { return "ClearVariableMonitoring" }

func (ClearVariableMonitoringRequest) Version() protocol.Version { return protocol.OCPP21 }

func (ClearVariableMonitoringRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (ClearVariableMonitoringRequest) SchemaName() string {
	return "ClearVariableMonitoringRequest.json"
}

func (message ClearVariableMonitoringRequest) Validate() error {
	return validation.Validate("ClearVariableMonitoringRequest.json", schemaClearVariableMonitoringRequest, message)
}

func (ClearVariableMonitoringRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ClearVariableMonitoringRequest.json", schemaClearVariableMonitoringRequest, data)
}
