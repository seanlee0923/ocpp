// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = GetTransactionStatusRequest{}

var schemaGetTransactionStatusRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"transactionId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type GetTransactionStatusRequest struct {
	TransactionID *string                                `json:"transactionId,omitempty"`
	CustomData    *GetTransactionStatusRequestCustomData `json:"customData,omitempty"`
}

type GetTransactionStatusRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value GetTransactionStatusRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *GetTransactionStatusRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (GetTransactionStatusRequest) ActionName() string { return "GetTransactionStatus" }

func (GetTransactionStatusRequest) Version() protocol.Version { return protocol.OCPP21 }

func (GetTransactionStatusRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (GetTransactionStatusRequest) SchemaName() string { return "GetTransactionStatusRequest.json" }

func (message GetTransactionStatusRequest) Validate() error {
	return validation.Validate("GetTransactionStatusRequest.json", schemaGetTransactionStatusRequest, message)
}

func (GetTransactionStatusRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetTransactionStatusRequest.json", schemaGetTransactionStatusRequest, data)
}
