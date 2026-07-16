// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = NotifyWebPaymentStartedConfirmation{}

var schemaNotifyWebPaymentStartedConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type NotifyWebPaymentStartedConfirmation struct {
	CustomData *NotifyWebPaymentStartedConfirmationCustomData `json:"customData,omitempty"`
}

type NotifyWebPaymentStartedConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value NotifyWebPaymentStartedConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *NotifyWebPaymentStartedConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (NotifyWebPaymentStartedConfirmation) ActionName() string { return "NotifyWebPaymentStarted" }

func (NotifyWebPaymentStartedConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (NotifyWebPaymentStartedConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (NotifyWebPaymentStartedConfirmation) SchemaName() string {
	return "NotifyWebPaymentStartedResponse.json"
}

func (message NotifyWebPaymentStartedConfirmation) Validate() error {
	return validation.Validate("NotifyWebPaymentStartedResponse.json", schemaNotifyWebPaymentStartedConfirmation, message)
}

func (NotifyWebPaymentStartedConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("NotifyWebPaymentStartedResponse.json", schemaNotifyWebPaymentStartedConfirmation, data)
}
