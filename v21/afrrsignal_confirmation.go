// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = AFRRSignalConfirmation{}

var schemaAFRRSignalConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type AFRRSignalConfirmation struct {
	Status     AFRRSignalConfirmationGenericStatusEnum `json:"status"`
	StatusInfo *AFRRSignalConfirmationStatusInfo       `json:"statusInfo,omitempty"`
	CustomData *AFRRSignalConfirmationCustomData       `json:"customData,omitempty"`
}

type AFRRSignalConfirmationStatusInfo struct {
	ReasonCode     string                            `json:"reasonCode"`
	AdditionalInfo *string                           `json:"additionalInfo,omitempty"`
	CustomData     *AFRRSignalConfirmationCustomData `json:"customData,omitempty"`
}

type AFRRSignalConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value AFRRSignalConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *AFRRSignalConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

type AFRRSignalConfirmationGenericStatusEnum string

const (
	AFRRSignalConfirmationGenericStatusEnumAccepted AFRRSignalConfirmationGenericStatusEnum = "Accepted"
	AFRRSignalConfirmationGenericStatusEnumRejected AFRRSignalConfirmationGenericStatusEnum = "Rejected"
)

func (AFRRSignalConfirmation) ActionName() string { return "AFRRSignal" }

func (AFRRSignalConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (AFRRSignalConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (AFRRSignalConfirmation) SchemaName() string { return "AFRRSignalResponse.json" }

func (message AFRRSignalConfirmation) Validate() error {
	return validation.Validate("AFRRSignalResponse.json", schemaAFRRSignalConfirmation, message)
}

func (AFRRSignalConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("AFRRSignalResponse.json", schemaAFRRSignalConfirmation, data)
}
