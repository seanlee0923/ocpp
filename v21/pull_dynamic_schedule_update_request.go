// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = PullDynamicScheduleUpdateRequest{}

var schemaPullDynamicScheduleUpdateRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"chargingProfileId": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"chargingProfileId"}}

type PullDynamicScheduleUpdateRequest struct {
	ChargingProfileID int                                         `json:"chargingProfileId"`
	CustomData        *PullDynamicScheduleUpdateRequestCustomData `json:"customData,omitempty"`
}

type PullDynamicScheduleUpdateRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value PullDynamicScheduleUpdateRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *PullDynamicScheduleUpdateRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (PullDynamicScheduleUpdateRequest) ActionName() string { return "PullDynamicScheduleUpdate" }

func (PullDynamicScheduleUpdateRequest) Version() protocol.Version { return protocol.OCPP21 }

func (PullDynamicScheduleUpdateRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (PullDynamicScheduleUpdateRequest) SchemaName() string {
	return "PullDynamicScheduleUpdateRequest.json"
}

func (message PullDynamicScheduleUpdateRequest) Validate() error {
	return validation.Validate("PullDynamicScheduleUpdateRequest.json", schemaPullDynamicScheduleUpdateRequest, message)
}

func (PullDynamicScheduleUpdateRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("PullDynamicScheduleUpdateRequest.json", schemaPullDynamicScheduleUpdateRequest, data)
}
