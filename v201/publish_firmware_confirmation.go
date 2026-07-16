// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = PublishFirmwareConfirmation{}

var schemaPublishFirmwareConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}}, Required: []string{"status"}}

type PublishFirmwareConfirmation struct {
	CustomData *PublishFirmwareConfirmationCustomData       `json:"customData,omitempty"`
	Status     PublishFirmwareConfirmationGenericStatusEnum `json:"status"`
	StatusInfo *PublishFirmwareConfirmationStatusInfo       `json:"statusInfo,omitempty"`
}

type PublishFirmwareConfirmationStatusInfo struct {
	CustomData     *PublishFirmwareConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                                 `json:"reasonCode"`
	AdditionalInfo *string                                `json:"additionalInfo,omitempty"`
}

type PublishFirmwareConfirmationGenericStatusEnum string

const (
	PublishFirmwareConfirmationGenericStatusEnumAccepted PublishFirmwareConfirmationGenericStatusEnum = "Accepted"
	PublishFirmwareConfirmationGenericStatusEnumRejected PublishFirmwareConfirmationGenericStatusEnum = "Rejected"
)

type PublishFirmwareConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value PublishFirmwareConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *PublishFirmwareConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (PublishFirmwareConfirmation) ActionName() string { return "PublishFirmware" }

func (PublishFirmwareConfirmation) Version() protocol.Version { return protocol.OCPP201 }

func (PublishFirmwareConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (PublishFirmwareConfirmation) SchemaName() string { return "PublishFirmwareResponse.json" }

func (message PublishFirmwareConfirmation) Validate() error {
	return validation.Validate("PublishFirmwareResponse.json", schemaPublishFirmwareConfirmation, message)
}

func (PublishFirmwareConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("PublishFirmwareResponse.json", schemaPublishFirmwareConfirmation, data)
}
