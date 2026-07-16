// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = ClearedChargingLimitRequest{}

var schemaClearedChargingLimitRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"chargingLimitSource": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "evseId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"chargingLimitSource"}}

type ClearedChargingLimitRequest struct {
	ChargingLimitSource string                                 `json:"chargingLimitSource"`
	EVSEID              *int                                   `json:"evseId,omitempty"`
	CustomData          *ClearedChargingLimitRequestCustomData `json:"customData,omitempty"`
}

type ClearedChargingLimitRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (ClearedChargingLimitRequest) ActionName() string { return "ClearedChargingLimit" }

func (ClearedChargingLimitRequest) Version() protocol.Version { return protocol.OCPP21 }

func (ClearedChargingLimitRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (ClearedChargingLimitRequest) SchemaName() string { return "ClearedChargingLimitRequest.json" }

func (message ClearedChargingLimitRequest) Validate() error {
	return validation.Validate("ClearedChargingLimitRequest.json", schemaClearedChargingLimitRequest, message)
}

func (ClearedChargingLimitRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ClearedChargingLimitRequest.json", schemaClearedChargingLimitRequest, data)
}
