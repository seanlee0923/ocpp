// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = CancelReservationConfirmation{}

var schemaCancelReservationConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}}, Required: []string{"status"}}

type CancelReservationConfirmation struct {
	CustomData *CancelReservationConfirmationCustomData                 `json:"customData,omitempty"`
	Status     CancelReservationConfirmationCancelReservationStatusEnum `json:"status"`
	StatusInfo *CancelReservationConfirmationStatusInfo                 `json:"statusInfo,omitempty"`
}

type CancelReservationConfirmationStatusInfo struct {
	CustomData     *CancelReservationConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                                   `json:"reasonCode"`
	AdditionalInfo *string                                  `json:"additionalInfo,omitempty"`
}

type CancelReservationConfirmationCancelReservationStatusEnum string

const (
	CancelReservationConfirmationCancelReservationStatusEnumAccepted CancelReservationConfirmationCancelReservationStatusEnum = "Accepted"
	CancelReservationConfirmationCancelReservationStatusEnumRejected CancelReservationConfirmationCancelReservationStatusEnum = "Rejected"
)

type CancelReservationConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value CancelReservationConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *CancelReservationConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (CancelReservationConfirmation) ActionName() string { return "CancelReservation" }

func (CancelReservationConfirmation) Version() protocol.Version { return protocol.OCPP201 }

func (CancelReservationConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (CancelReservationConfirmation) SchemaName() string { return "CancelReservationResponse.json" }

func (message CancelReservationConfirmation) Validate() error {
	return validation.Validate("CancelReservationResponse.json", schemaCancelReservationConfirmation, message)
}

func (CancelReservationConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("CancelReservationResponse.json", schemaCancelReservationConfirmation, data)
}
