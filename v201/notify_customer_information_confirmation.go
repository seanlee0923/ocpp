// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = NotifyCustomerInformationConfirmation{}

var schemaNotifyCustomerInformationConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type NotifyCustomerInformationConfirmation struct {
	CustomData *NotifyCustomerInformationConfirmationCustomData `json:"customData,omitempty"`
}

type NotifyCustomerInformationConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (NotifyCustomerInformationConfirmation) ActionName() string { return "NotifyCustomerInformation" }

func (NotifyCustomerInformationConfirmation) Version() protocol.Version { return protocol.OCPP201 }

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
