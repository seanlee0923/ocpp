// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = UsePriorityChargingRequest{}

var schemaUsePriorityChargingRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"transactionId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "activate": &validation.Schema{Type: "boolean", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"transactionId", "activate"}}

type UsePriorityChargingRequest struct {
	TransactionID string                                `json:"transactionId"`
	Activate      bool                                  `json:"activate"`
	CustomData    *UsePriorityChargingRequestCustomData `json:"customData,omitempty"`
}

type UsePriorityChargingRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (UsePriorityChargingRequest) ActionName() string { return "UsePriorityCharging" }

func (UsePriorityChargingRequest) Version() protocol.Version { return protocol.OCPP21 }

func (UsePriorityChargingRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (UsePriorityChargingRequest) SchemaName() string { return "UsePriorityChargingRequest.json" }

func (message UsePriorityChargingRequest) Validate() error {
	return validation.Validate("UsePriorityChargingRequest.json", schemaUsePriorityChargingRequest, message)
}

func (UsePriorityChargingRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("UsePriorityChargingRequest.json", schemaUsePriorityChargingRequest, data)
}
