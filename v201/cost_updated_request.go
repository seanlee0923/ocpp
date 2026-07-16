// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = CostUpdatedRequest{}

var schemaCostUpdatedRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "totalCost": &validation.Schema{Type: "number", AllowAdditional: true}, "transactionId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}}, Required: []string{"totalCost", "transactionId"}}

type CostUpdatedRequest struct {
	CustomData    *CostUpdatedRequestCustomData `json:"customData,omitempty"`
	TotalCost     float64                       `json:"totalCost"`
	TransactionID string                        `json:"transactionId"`
}

type CostUpdatedRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (CostUpdatedRequest) ActionName() string { return "CostUpdated" }

func (CostUpdatedRequest) Version() protocol.Version { return protocol.OCPP201 }

func (CostUpdatedRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (CostUpdatedRequest) SchemaName() string { return "CostUpdatedRequest.json" }

func (message CostUpdatedRequest) Validate() error {
	return validation.Validate("CostUpdatedRequest.json", schemaCostUpdatedRequest, message)
}

func (CostUpdatedRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("CostUpdatedRequest.json", schemaCostUpdatedRequest, data)
}
