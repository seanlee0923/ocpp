// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ClearedChargingLimitRequest{}

var schemaClearedChargingLimitRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "chargingLimitSource": &validation.Schema{Type: "string", Enum: []string{"EMS", "Other", "SO", "CSO"}}, "evseId": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"chargingLimitSource"}}

type ClearedChargingLimitRequest struct {
	CustomData          *ClearedChargingLimitRequestCustomData             `json:"customData,omitempty"`
	ChargingLimitSource ClearedChargingLimitRequestChargingLimitSourceEnum `json:"chargingLimitSource"`
	EVSEID              *int                                               `json:"evseId,omitempty"`
}

type ClearedChargingLimitRequestChargingLimitSourceEnum string

const (
	ClearedChargingLimitRequestChargingLimitSourceEnumEMS   ClearedChargingLimitRequestChargingLimitSourceEnum = "EMS"
	ClearedChargingLimitRequestChargingLimitSourceEnumOther ClearedChargingLimitRequestChargingLimitSourceEnum = "Other"
	ClearedChargingLimitRequestChargingLimitSourceEnumSO    ClearedChargingLimitRequestChargingLimitSourceEnum = "SO"
	ClearedChargingLimitRequestChargingLimitSourceEnumCSO   ClearedChargingLimitRequestChargingLimitSourceEnum = "CSO"
)

type ClearedChargingLimitRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value ClearedChargingLimitRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *ClearedChargingLimitRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (ClearedChargingLimitRequest) ActionName() string { return "ClearedChargingLimit" }

func (ClearedChargingLimitRequest) Version() protocol.Version { return protocol.OCPP201 }

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
