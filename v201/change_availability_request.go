// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = ChangeAvailabilityRequest{}

var schemaChangeAvailabilityRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "evse": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "id": &validation.Schema{Type: "integer", AllowAdditional: true}, "connectorId": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"id"}}, "operationalStatus": &validation.Schema{Type: "string", Enum: []string{"Inoperative", "Operative"}}}, Required: []string{"operationalStatus"}}

type ChangeAvailabilityRequest struct {
	CustomData        *ChangeAvailabilityRequestCustomData           `json:"customData,omitempty"`
	EVSE              *ChangeAvailabilityRequestEVSE                 `json:"evse,omitempty"`
	OperationalStatus ChangeAvailabilityRequestOperationalStatusEnum `json:"operationalStatus"`
}

type ChangeAvailabilityRequestOperationalStatusEnum string

const (
	ChangeAvailabilityRequestOperationalStatusEnumInoperative ChangeAvailabilityRequestOperationalStatusEnum = "Inoperative"
	ChangeAvailabilityRequestOperationalStatusEnumOperative   ChangeAvailabilityRequestOperationalStatusEnum = "Operative"
)

type ChangeAvailabilityRequestEVSE struct {
	CustomData  *ChangeAvailabilityRequestCustomData `json:"customData,omitempty"`
	ID          int                                  `json:"id"`
	ConnectorID *int                                 `json:"connectorId,omitempty"`
}

type ChangeAvailabilityRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (ChangeAvailabilityRequest) ActionName() string { return "ChangeAvailability" }

func (ChangeAvailabilityRequest) Version() protocol.Version { return protocol.OCPP201 }

func (ChangeAvailabilityRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (ChangeAvailabilityRequest) SchemaName() string { return "ChangeAvailabilityRequest.json" }

func (message ChangeAvailabilityRequest) Validate() error {
	return validation.Validate("ChangeAvailabilityRequest.json", schemaChangeAvailabilityRequest, message)
}

func (ChangeAvailabilityRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ChangeAvailabilityRequest.json", schemaChangeAvailabilityRequest, data)
}
