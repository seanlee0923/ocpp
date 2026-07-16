// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ChangeAvailabilityRequest{}

var schemaChangeAvailabilityRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"evse": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"id": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "connectorId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id"}}, "operationalStatus": &validation.Schema{Type: "string", Enum: []string{"Inoperative", "Operative"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"operationalStatus"}}

type ChangeAvailabilityRequest struct {
	EVSE              *ChangeAvailabilityRequestEVSE                 `json:"evse,omitempty"`
	OperationalStatus ChangeAvailabilityRequestOperationalStatusEnum `json:"operationalStatus"`
	CustomData        *ChangeAvailabilityRequestCustomData           `json:"customData,omitempty"`
}

type ChangeAvailabilityRequestOperationalStatusEnum string

const (
	ChangeAvailabilityRequestOperationalStatusEnumInoperative ChangeAvailabilityRequestOperationalStatusEnum = "Inoperative"
	ChangeAvailabilityRequestOperationalStatusEnumOperative   ChangeAvailabilityRequestOperationalStatusEnum = "Operative"
)

type ChangeAvailabilityRequestEVSE struct {
	ID          int                                  `json:"id"`
	ConnectorID *int                                 `json:"connectorId,omitempty"`
	CustomData  *ChangeAvailabilityRequestCustomData `json:"customData,omitempty"`
}

type ChangeAvailabilityRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value ChangeAvailabilityRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *ChangeAvailabilityRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (ChangeAvailabilityRequest) ActionName() string { return "ChangeAvailability" }

func (ChangeAvailabilityRequest) Version() protocol.Version { return protocol.OCPP21 }

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
