// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = SetMonitoringBaseRequest{}

var schemaSetMonitoringBaseRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"monitoringBase": &validation.Schema{Type: "string", Enum: []string{"All", "FactoryDefault", "HardWiredOnly"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"monitoringBase"}}

type SetMonitoringBaseRequest struct {
	MonitoringBase SetMonitoringBaseRequestMonitoringBaseEnum `json:"monitoringBase"`
	CustomData     *SetMonitoringBaseRequestCustomData        `json:"customData,omitempty"`
}

type SetMonitoringBaseRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

type SetMonitoringBaseRequestMonitoringBaseEnum string

const (
	SetMonitoringBaseRequestMonitoringBaseEnumAll            SetMonitoringBaseRequestMonitoringBaseEnum = "All"
	SetMonitoringBaseRequestMonitoringBaseEnumFactoryDefault SetMonitoringBaseRequestMonitoringBaseEnum = "FactoryDefault"
	SetMonitoringBaseRequestMonitoringBaseEnumHardWiredOnly  SetMonitoringBaseRequestMonitoringBaseEnum = "HardWiredOnly"
)

func (SetMonitoringBaseRequest) ActionName() string { return "SetMonitoringBase" }

func (SetMonitoringBaseRequest) Version() protocol.Version { return protocol.OCPP21 }

func (SetMonitoringBaseRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (SetMonitoringBaseRequest) SchemaName() string { return "SetMonitoringBaseRequest.json" }

func (message SetMonitoringBaseRequest) Validate() error {
	return validation.Validate("SetMonitoringBaseRequest.json", schemaSetMonitoringBaseRequest, message)
}

func (SetMonitoringBaseRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("SetMonitoringBaseRequest.json", schemaSetMonitoringBaseRequest, data)
}
