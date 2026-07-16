// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = CustomerInformationConfirmation{}

var schemaCustomerInformationConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "Invalid"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type CustomerInformationConfirmation struct {
	Status     CustomerInformationConfirmationCustomerInformationStatusEnum `json:"status"`
	StatusInfo *CustomerInformationConfirmationStatusInfo                   `json:"statusInfo,omitempty"`
	CustomData *CustomerInformationConfirmationCustomData                   `json:"customData,omitempty"`
}

type CustomerInformationConfirmationStatusInfo struct {
	ReasonCode     string                                     `json:"reasonCode"`
	AdditionalInfo *string                                    `json:"additionalInfo,omitempty"`
	CustomData     *CustomerInformationConfirmationCustomData `json:"customData,omitempty"`
}

type CustomerInformationConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value CustomerInformationConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *CustomerInformationConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

type CustomerInformationConfirmationCustomerInformationStatusEnum string

const (
	CustomerInformationConfirmationCustomerInformationStatusEnumAccepted CustomerInformationConfirmationCustomerInformationStatusEnum = "Accepted"
	CustomerInformationConfirmationCustomerInformationStatusEnumRejected CustomerInformationConfirmationCustomerInformationStatusEnum = "Rejected"
	CustomerInformationConfirmationCustomerInformationStatusEnumInvalid  CustomerInformationConfirmationCustomerInformationStatusEnum = "Invalid"
)

func (CustomerInformationConfirmation) ActionName() string { return "CustomerInformation" }

func (CustomerInformationConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (CustomerInformationConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (CustomerInformationConfirmation) SchemaName() string { return "CustomerInformationResponse.json" }

func (message CustomerInformationConfirmation) Validate() error {
	return validation.Validate("CustomerInformationResponse.json", schemaCustomerInformationConfirmation, message)
}

func (CustomerInformationConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("CustomerInformationResponse.json", schemaCustomerInformationConfirmation, data)
}
