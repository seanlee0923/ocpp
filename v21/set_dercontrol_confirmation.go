// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = SetDERControlConfirmation{}

var schemaSetDERControlConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "NotSupported", "NotFound"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "supersededIds": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, MinItems: 1, HasMinItems: true, MaxItems: 24, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type SetDERControlConfirmation struct {
	Status        SetDERControlConfirmationDERControlStatusEnum `json:"status"`
	StatusInfo    *SetDERControlConfirmationStatusInfo          `json:"statusInfo,omitempty"`
	SupersededIds []string                                      `json:"supersededIds,omitempty"`
	CustomData    *SetDERControlConfirmationCustomData          `json:"customData,omitempty"`
}

type SetDERControlConfirmationStatusInfo struct {
	ReasonCode     string                               `json:"reasonCode"`
	AdditionalInfo *string                              `json:"additionalInfo,omitempty"`
	CustomData     *SetDERControlConfirmationCustomData `json:"customData,omitempty"`
}

type SetDERControlConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value SetDERControlConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *SetDERControlConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

type SetDERControlConfirmationDERControlStatusEnum string

const (
	SetDERControlConfirmationDERControlStatusEnumAccepted     SetDERControlConfirmationDERControlStatusEnum = "Accepted"
	SetDERControlConfirmationDERControlStatusEnumRejected     SetDERControlConfirmationDERControlStatusEnum = "Rejected"
	SetDERControlConfirmationDERControlStatusEnumNotSupported SetDERControlConfirmationDERControlStatusEnum = "NotSupported"
	SetDERControlConfirmationDERControlStatusEnumNotFound     SetDERControlConfirmationDERControlStatusEnum = "NotFound"
)

func (SetDERControlConfirmation) ActionName() string { return "SetDERControl" }

func (SetDERControlConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (SetDERControlConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (SetDERControlConfirmation) SchemaName() string { return "SetDERControlResponse.json" }

func (message SetDERControlConfirmation) Validate() error {
	return validation.Validate("SetDERControlResponse.json", schemaSetDERControlConfirmation, message)
}

func (SetDERControlConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("SetDERControlResponse.json", schemaSetDERControlConfirmation, data)
}
