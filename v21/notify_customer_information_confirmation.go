// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = NotifyCustomerInformationConfirmation{}

var schemaNotifyCustomerInformationConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type NotifyCustomerInformationConfirmation struct {
	CustomData *NotifyCustomerInformationConfirmationCustomData `json:"customData,omitempty"`
}

type NotifyCustomerInformationConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value NotifyCustomerInformationConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *NotifyCustomerInformationConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (NotifyCustomerInformationConfirmation) ActionName() string { return "NotifyCustomerInformation" }

func (NotifyCustomerInformationConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (NotifyCustomerInformationConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (NotifyCustomerInformationConfirmation) SchemaName() string {
	return "NotifyCustomerInformationResponse.json"
}

func (message NotifyCustomerInformationConfirmation) Validate() error {
	return validation.Validate("NotifyCustomerInformationResponse.json", schemaNotifyCustomerInformationConfirmation, message)
}

func (NotifyCustomerInformationConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("NotifyCustomerInformationResponse.json", schemaNotifyCustomerInformationConfirmation, data)
}
