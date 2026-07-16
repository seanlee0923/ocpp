// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = GetTariffsRequest{}

var schemaGetTariffsRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"evseId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"evseId"}}

type GetTariffsRequest struct {
	EVSEID     int                          `json:"evseId"`
	CustomData *GetTariffsRequestCustomData `json:"customData,omitempty"`
}

type GetTariffsRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value GetTariffsRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *GetTariffsRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (GetTariffsRequest) ActionName() string { return "GetTariffs" }

func (GetTariffsRequest) Version() protocol.Version { return protocol.OCPP21 }

func (GetTariffsRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (GetTariffsRequest) SchemaName() string { return "GetTariffsRequest.json" }

func (message GetTariffsRequest) Validate() error {
	return validation.Validate("GetTariffsRequest.json", schemaGetTariffsRequest, message)
}

func (GetTariffsRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetTariffsRequest.json", schemaGetTariffsRequest, data)
}
