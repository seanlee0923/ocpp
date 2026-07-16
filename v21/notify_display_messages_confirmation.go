// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = NotifyDisplayMessagesConfirmation{}

var schemaNotifyDisplayMessagesConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type NotifyDisplayMessagesConfirmation struct {
	CustomData *NotifyDisplayMessagesConfirmationCustomData `json:"customData,omitempty"`
}

type NotifyDisplayMessagesConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value NotifyDisplayMessagesConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *NotifyDisplayMessagesConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (NotifyDisplayMessagesConfirmation) ActionName() string { return "NotifyDisplayMessages" }

func (NotifyDisplayMessagesConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (NotifyDisplayMessagesConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (NotifyDisplayMessagesConfirmation) SchemaName() string {
	return "NotifyDisplayMessagesResponse.json"
}

func (message NotifyDisplayMessagesConfirmation) Validate() error {
	return validation.Validate("NotifyDisplayMessagesResponse.json", schemaNotifyDisplayMessagesConfirmation, message)
}

func (NotifyDisplayMessagesConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("NotifyDisplayMessagesResponse.json", schemaNotifyDisplayMessagesConfirmation, data)
}
