// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ChangeAvailabilityRequest{}

var schemaChangeAvailabilityRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"connectorId": &validation.Schema{Type: "integer", AllowAdditional: true}, "type": &validation.Schema{Type: "string", Enum: []string{"Inoperative", "Operative"}}}, Required: []string{"connectorId", "type"}}

type ChangeAvailabilityRequest struct {
	ConnectorID int                           `json:"connectorId"`
	Type        ChangeAvailabilityRequestType `json:"type"`
}

type ChangeAvailabilityRequestType string

const (
	ChangeAvailabilityRequestTypeInoperative ChangeAvailabilityRequestType = "Inoperative"
	ChangeAvailabilityRequestTypeOperative   ChangeAvailabilityRequestType = "Operative"
)

func (ChangeAvailabilityRequest) ActionName() string { return "ChangeAvailability" }

func (ChangeAvailabilityRequest) Version() protocol.Version { return protocol.OCPP16 }

func (ChangeAvailabilityRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (ChangeAvailabilityRequest) SchemaName() string { return "ChangeAvailability.json" }

func (message ChangeAvailabilityRequest) Validate() error {
	return validation.Validate("ChangeAvailability.json", schemaChangeAvailabilityRequest, message)
}

func (ChangeAvailabilityRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ChangeAvailability.json", schemaChangeAvailabilityRequest, data)
}
