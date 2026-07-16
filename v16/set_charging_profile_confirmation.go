// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = SetChargingProfileConfirmation{}

var schemaSetChargingProfileConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "NotSupported"}}}, Required: []string{"status"}}

type SetChargingProfileConfirmation struct {
	Status SetChargingProfileConfirmationStatus `json:"status"`
}

type SetChargingProfileConfirmationStatus string

const (
	SetChargingProfileConfirmationStatusAccepted     SetChargingProfileConfirmationStatus = "Accepted"
	SetChargingProfileConfirmationStatusRejected     SetChargingProfileConfirmationStatus = "Rejected"
	SetChargingProfileConfirmationStatusNotSupported SetChargingProfileConfirmationStatus = "NotSupported"
)

func (SetChargingProfileConfirmation) ActionName() string { return "SetChargingProfile" }

func (SetChargingProfileConfirmation) Version() protocol.Version { return protocol.OCPP16 }

func (SetChargingProfileConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (SetChargingProfileConfirmation) SchemaName() string { return "SetChargingProfileResponse.json" }

func (message SetChargingProfileConfirmation) Validate() error {
	return validation.Validate("SetChargingProfileResponse.json", schemaSetChargingProfileConfirmation, message)
}

func (SetChargingProfileConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("SetChargingProfileResponse.json", schemaSetChargingProfileConfirmation, data)
}
