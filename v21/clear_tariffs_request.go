// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ClearTariffsRequest{}

var schemaClearTariffsRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"tariffIds": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 60, HasMaxLength: true}, MinItems: 1, HasMinItems: true}, "evseId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type ClearTariffsRequest struct {
	TariffIds  []string                       `json:"tariffIds,omitempty"`
	EVSEID     *int                           `json:"evseId,omitempty"`
	CustomData *ClearTariffsRequestCustomData `json:"customData,omitempty"`
}

type ClearTariffsRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (ClearTariffsRequest) ActionName() string { return "ClearTariffs" }

func (ClearTariffsRequest) Version() protocol.Version { return protocol.OCPP21 }

func (ClearTariffsRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (ClearTariffsRequest) SchemaName() string { return "ClearTariffsRequest.json" }

func (message ClearTariffsRequest) Validate() error {
	return validation.Validate("ClearTariffsRequest.json", schemaClearTariffsRequest, message)
}

func (ClearTariffsRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ClearTariffsRequest.json", schemaClearTariffsRequest, data)
}
