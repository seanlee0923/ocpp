// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = SetMonitoringLevelRequest{}

var schemaSetMonitoringLevelRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"severity": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"severity"}}

type SetMonitoringLevelRequest struct {
	Severity   int                                  `json:"severity"`
	CustomData *SetMonitoringLevelRequestCustomData `json:"customData,omitempty"`
}

type SetMonitoringLevelRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (SetMonitoringLevelRequest) ActionName() string { return "SetMonitoringLevel" }

func (SetMonitoringLevelRequest) Version() protocol.Version { return protocol.OCPP21 }

func (SetMonitoringLevelRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (SetMonitoringLevelRequest) SchemaName() string { return "SetMonitoringLevelRequest.json" }

func (message SetMonitoringLevelRequest) Validate() error {
	return validation.Validate("SetMonitoringLevelRequest.json", schemaSetMonitoringLevelRequest, message)
}

func (SetMonitoringLevelRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("SetMonitoringLevelRequest.json", schemaSetMonitoringLevelRequest, data)
}
