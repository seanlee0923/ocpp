// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = ClearChargingProfileConfirmation{}

var schemaClearChargingProfileConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Unknown"}}}, Required: []string{"status"}}

type ClearChargingProfileConfirmation struct {
	Status ClearChargingProfileConfirmationStatus `json:"status"`
}

type ClearChargingProfileConfirmationStatus string

const (
	ClearChargingProfileConfirmationStatusAccepted ClearChargingProfileConfirmationStatus = "Accepted"
	ClearChargingProfileConfirmationStatusUnknown  ClearChargingProfileConfirmationStatus = "Unknown"
)

func (ClearChargingProfileConfirmation) ActionName() string { return "ClearChargingProfile" }

func (ClearChargingProfileConfirmation) Version() protocol.Version { return protocol.OCPP16 }

func (ClearChargingProfileConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (ClearChargingProfileConfirmation) SchemaName() string {
	return "ClearChargingProfileResponse.json"
}

func (message ClearChargingProfileConfirmation) Validate() error {
	return validation.Validate("ClearChargingProfileResponse.json", schemaClearChargingProfileConfirmation, message)
}

func (ClearChargingProfileConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ClearChargingProfileResponse.json", schemaClearChargingProfileConfirmation, data)
}
