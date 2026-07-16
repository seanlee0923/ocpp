// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = UnlockConnectorRequest{}

var schemaUnlockConnectorRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"connectorId": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"connectorId"}}

type UnlockConnectorRequest struct {
	ConnectorID int `json:"connectorId"`
}

func (UnlockConnectorRequest) ActionName() string { return "UnlockConnector" }

func (UnlockConnectorRequest) Version() protocol.Version { return protocol.OCPP16 }

func (UnlockConnectorRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (UnlockConnectorRequest) SchemaName() string { return "UnlockConnector.json" }

func (message UnlockConnectorRequest) Validate() error {
	return validation.Validate("UnlockConnector.json", schemaUnlockConnectorRequest, message)
}

func (UnlockConnectorRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("UnlockConnector.json", schemaUnlockConnectorRequest, data)
}
