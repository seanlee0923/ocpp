// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = UnlockConnectorRequest{}

var schemaUnlockConnectorRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"evseId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "connectorId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"evseId", "connectorId"}}

type UnlockConnectorRequest struct {
	EVSEID      int                               `json:"evseId"`
	ConnectorID int                               `json:"connectorId"`
	CustomData  *UnlockConnectorRequestCustomData `json:"customData,omitempty"`
}

type UnlockConnectorRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value UnlockConnectorRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *UnlockConnectorRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (UnlockConnectorRequest) ActionName() string { return "UnlockConnector" }

func (UnlockConnectorRequest) Version() protocol.Version { return protocol.OCPP21 }

func (UnlockConnectorRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (UnlockConnectorRequest) SchemaName() string { return "UnlockConnectorRequest.json" }

func (message UnlockConnectorRequest) Validate() error {
	return validation.Validate("UnlockConnectorRequest.json", schemaUnlockConnectorRequest, message)
}

func (UnlockConnectorRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("UnlockConnectorRequest.json", schemaUnlockConnectorRequest, data)
}
