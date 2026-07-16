// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ClearChargingProfileRequest{}

var schemaClearChargingProfileRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"id": &validation.Schema{Type: "integer", AllowAdditional: true}, "connectorId": &validation.Schema{Type: "integer", AllowAdditional: true}, "chargingProfilePurpose": &validation.Schema{Type: "string", Enum: []string{"ChargePointMaxProfile", "TxDefaultProfile", "TxProfile"}}, "stackLevel": &validation.Schema{Type: "integer", AllowAdditional: true}}}

type ClearChargingProfileRequest struct {
	ID                     *int                                               `json:"id,omitempty"`
	ConnectorID            *int                                               `json:"connectorId,omitempty"`
	ChargingProfilePurpose *ClearChargingProfileRequestChargingProfilePurpose `json:"chargingProfilePurpose,omitempty"`
	StackLevel             *int                                               `json:"stackLevel,omitempty"`
}

type ClearChargingProfileRequestChargingProfilePurpose string

const (
	ClearChargingProfileRequestChargingProfilePurposeChargePointMaxProfile ClearChargingProfileRequestChargingProfilePurpose = "ChargePointMaxProfile"
	ClearChargingProfileRequestChargingProfilePurposeTxDefaultProfile      ClearChargingProfileRequestChargingProfilePurpose = "TxDefaultProfile"
	ClearChargingProfileRequestChargingProfilePurposeTxProfile             ClearChargingProfileRequestChargingProfilePurpose = "TxProfile"
)

func (ClearChargingProfileRequest) ActionName() string { return "ClearChargingProfile" }

func (ClearChargingProfileRequest) Version() protocol.Version { return protocol.OCPP16 }

func (ClearChargingProfileRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (ClearChargingProfileRequest) SchemaName() string { return "ClearChargingProfile.json" }

func (message ClearChargingProfileRequest) Validate() error {
	return validation.Validate("ClearChargingProfile.json", schemaClearChargingProfileRequest, message)
}

func (ClearChargingProfileRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ClearChargingProfile.json", schemaClearChargingProfileRequest, data)
}
