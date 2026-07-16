// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = BootNotificationRequest{}

var schemaBootNotificationRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"chargePointVendor": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "chargePointModel": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "chargePointSerialNumber": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 25, HasMaxLength: true}, "chargeBoxSerialNumber": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 25, HasMaxLength: true}, "firmwareVersion": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "iccid": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "imsi": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "meterType": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 25, HasMaxLength: true}, "meterSerialNumber": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 25, HasMaxLength: true}}, Required: []string{"chargePointVendor", "chargePointModel"}}

type BootNotificationRequest struct {
	ChargePointVendor       string  `json:"chargePointVendor"`
	ChargePointModel        string  `json:"chargePointModel"`
	ChargePointSerialNumber *string `json:"chargePointSerialNumber,omitempty"`
	ChargeBoxSerialNumber   *string `json:"chargeBoxSerialNumber,omitempty"`
	FirmwareVersion         *string `json:"firmwareVersion,omitempty"`
	ICCID                   *string `json:"iccid,omitempty"`
	IMSI                    *string `json:"imsi,omitempty"`
	MeterType               *string `json:"meterType,omitempty"`
	MeterSerialNumber       *string `json:"meterSerialNumber,omitempty"`
}

func (BootNotificationRequest) ActionName() string { return "BootNotification" }

func (BootNotificationRequest) Version() protocol.Version { return protocol.OCPP16 }

func (BootNotificationRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (BootNotificationRequest) SchemaName() string { return "BootNotification.json" }

func (message BootNotificationRequest) Validate() error {
	return validation.Validate("BootNotification.json", schemaBootNotificationRequest, message)
}

func (BootNotificationRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("BootNotification.json", schemaBootNotificationRequest, data)
}
