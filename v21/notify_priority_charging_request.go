// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = NotifyPriorityChargingRequest{}

var schemaNotifyPriorityChargingRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"transactionId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "activated": &validation.Schema{Type: "boolean", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"transactionId", "activated"}}

type NotifyPriorityChargingRequest struct {
	TransactionID string                                   `json:"transactionId"`
	Activated     bool                                     `json:"activated"`
	CustomData    *NotifyPriorityChargingRequestCustomData `json:"customData,omitempty"`
}

type NotifyPriorityChargingRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (NotifyPriorityChargingRequest) ActionName() string { return "NotifyPriorityCharging" }

func (NotifyPriorityChargingRequest) Version() protocol.Version { return protocol.OCPP21 }

func (NotifyPriorityChargingRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (NotifyPriorityChargingRequest) SchemaName() string { return "NotifyPriorityChargingRequest.json" }

func (message NotifyPriorityChargingRequest) Validate() error {
	return validation.Validate("NotifyPriorityChargingRequest.json", schemaNotifyPriorityChargingRequest, message)
}

func (NotifyPriorityChargingRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("NotifyPriorityChargingRequest.json", schemaNotifyPriorityChargingRequest, data)
}
