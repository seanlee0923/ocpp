// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = NotifySettlementConfirmation{}

var schemaNotifySettlementConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"receiptUrl": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 2000, HasMaxLength: true}, "receiptId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type NotifySettlementConfirmation struct {
	ReceiptUrl *string                                 `json:"receiptUrl,omitempty"`
	ReceiptID  *string                                 `json:"receiptId,omitempty"`
	CustomData *NotifySettlementConfirmationCustomData `json:"customData,omitempty"`
}

type NotifySettlementConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (NotifySettlementConfirmation) ActionName() string { return "NotifySettlement" }

func (NotifySettlementConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (NotifySettlementConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (NotifySettlementConfirmation) SchemaName() string { return "NotifySettlementResponse.json" }

func (message NotifySettlementConfirmation) Validate() error {
	return validation.Validate("NotifySettlementResponse.json", schemaNotifySettlementConfirmation, message)
}

func (NotifySettlementConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("NotifySettlementResponse.json", schemaNotifySettlementConfirmation, data)
}
