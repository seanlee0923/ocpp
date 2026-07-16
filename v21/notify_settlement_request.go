// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = NotifySettlementRequest{}

var schemaNotifySettlementRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"transactionId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "pspRef": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Settled", "Canceled", "Rejected", "Failed"}}, "statusInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 500, HasMaxLength: true}, "settlementAmount": &validation.Schema{Type: "number", AllowAdditional: true}, "settlementTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "receiptId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "receiptUrl": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 2000, HasMaxLength: true}, "vatCompany": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"name": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "address1": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 100, HasMaxLength: true}, "address2": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 100, HasMaxLength: true}, "city": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 100, HasMaxLength: true}, "postalCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "country": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"name", "address1", "city", "country"}}, "vatNumber": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"pspRef", "status", "settlementAmount", "settlementTime"}}

type NotifySettlementRequest struct {
	TransactionID    *string                                  `json:"transactionId,omitempty"`
	PspRef           string                                   `json:"pspRef"`
	Status           NotifySettlementRequestPaymentStatusEnum `json:"status"`
	StatusInfo       *string                                  `json:"statusInfo,omitempty"`
	SettlementAmount float64                                  `json:"settlementAmount"`
	SettlementTime   string                                   `json:"settlementTime"`
	ReceiptID        *string                                  `json:"receiptId,omitempty"`
	ReceiptUrl       *string                                  `json:"receiptUrl,omitempty"`
	VatCompany       *NotifySettlementRequestAddress          `json:"vatCompany,omitempty"`
	VatNumber        *string                                  `json:"vatNumber,omitempty"`
	CustomData       *NotifySettlementRequestCustomData       `json:"customData,omitempty"`
}

type NotifySettlementRequestAddress struct {
	Name       string                             `json:"name"`
	Address1   string                             `json:"address1"`
	Address2   *string                            `json:"address2,omitempty"`
	City       string                             `json:"city"`
	PostalCode *string                            `json:"postalCode,omitempty"`
	Country    string                             `json:"country"`
	CustomData *NotifySettlementRequestCustomData `json:"customData,omitempty"`
}

type NotifySettlementRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value NotifySettlementRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *NotifySettlementRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

type NotifySettlementRequestPaymentStatusEnum string

const (
	NotifySettlementRequestPaymentStatusEnumSettled  NotifySettlementRequestPaymentStatusEnum = "Settled"
	NotifySettlementRequestPaymentStatusEnumCanceled NotifySettlementRequestPaymentStatusEnum = "Canceled"
	NotifySettlementRequestPaymentStatusEnumRejected NotifySettlementRequestPaymentStatusEnum = "Rejected"
	NotifySettlementRequestPaymentStatusEnumFailed   NotifySettlementRequestPaymentStatusEnum = "Failed"
)

func (NotifySettlementRequest) ActionName() string { return "NotifySettlement" }

func (NotifySettlementRequest) Version() protocol.Version { return protocol.OCPP21 }

func (NotifySettlementRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (NotifySettlementRequest) SchemaName() string { return "NotifySettlementRequest.json" }

func (message NotifySettlementRequest) Validate() error {
	return validation.Validate("NotifySettlementRequest.json", schemaNotifySettlementRequest, message)
}

func (NotifySettlementRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("NotifySettlementRequest.json", schemaNotifySettlementRequest, data)
}
