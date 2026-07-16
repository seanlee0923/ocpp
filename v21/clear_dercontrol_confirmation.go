// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ClearDERControlConfirmation{}

var schemaClearDERControlConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "NotSupported", "NotFound"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type ClearDERControlConfirmation struct {
	Status     ClearDERControlConfirmationDERControlStatusEnum `json:"status"`
	StatusInfo *ClearDERControlConfirmationStatusInfo          `json:"statusInfo,omitempty"`
	CustomData *ClearDERControlConfirmationCustomData          `json:"customData,omitempty"`
}

type ClearDERControlConfirmationStatusInfo struct {
	ReasonCode     string                                 `json:"reasonCode"`
	AdditionalInfo *string                                `json:"additionalInfo,omitempty"`
	CustomData     *ClearDERControlConfirmationCustomData `json:"customData,omitempty"`
}

type ClearDERControlConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value ClearDERControlConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *ClearDERControlConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

type ClearDERControlConfirmationDERControlStatusEnum string

const (
	ClearDERControlConfirmationDERControlStatusEnumAccepted     ClearDERControlConfirmationDERControlStatusEnum = "Accepted"
	ClearDERControlConfirmationDERControlStatusEnumRejected     ClearDERControlConfirmationDERControlStatusEnum = "Rejected"
	ClearDERControlConfirmationDERControlStatusEnumNotSupported ClearDERControlConfirmationDERControlStatusEnum = "NotSupported"
	ClearDERControlConfirmationDERControlStatusEnumNotFound     ClearDERControlConfirmationDERControlStatusEnum = "NotFound"
)

func (ClearDERControlConfirmation) ActionName() string { return "ClearDERControl" }

func (ClearDERControlConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (ClearDERControlConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (ClearDERControlConfirmation) SchemaName() string { return "ClearDERControlResponse.json" }

func (message ClearDERControlConfirmation) Validate() error {
	return validation.Validate("ClearDERControlResponse.json", schemaClearDERControlConfirmation, message)
}

func (ClearDERControlConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ClearDERControlResponse.json", schemaClearDERControlConfirmation, data)
}
