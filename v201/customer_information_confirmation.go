// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = CustomerInformationConfirmation{}

var schemaCustomerInformationConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "Invalid"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}}, Required: []string{"status"}}

type CustomerInformationConfirmation struct {
	CustomData *CustomerInformationConfirmationCustomData                   `json:"customData,omitempty"`
	Status     CustomerInformationConfirmationCustomerInformationStatusEnum `json:"status"`
	StatusInfo *CustomerInformationConfirmationStatusInfo                   `json:"statusInfo,omitempty"`
}

type CustomerInformationConfirmationStatusInfo struct {
	CustomData     *CustomerInformationConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                                     `json:"reasonCode"`
	AdditionalInfo *string                                    `json:"additionalInfo,omitempty"`
}

type CustomerInformationConfirmationCustomerInformationStatusEnum string

const (
	CustomerInformationConfirmationCustomerInformationStatusEnumAccepted CustomerInformationConfirmationCustomerInformationStatusEnum = "Accepted"
	CustomerInformationConfirmationCustomerInformationStatusEnumRejected CustomerInformationConfirmationCustomerInformationStatusEnum = "Rejected"
	CustomerInformationConfirmationCustomerInformationStatusEnumInvalid  CustomerInformationConfirmationCustomerInformationStatusEnum = "Invalid"
)

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

func (CustomerInformationConfirmation) ActionName() string { return "CustomerInformation" }

func (CustomerInformationConfirmation) Version() protocol.Version { return protocol.OCPP201 }

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
