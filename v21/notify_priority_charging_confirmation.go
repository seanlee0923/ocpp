// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = NotifyPriorityChargingConfirmation{}

var schemaNotifyPriorityChargingConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type NotifyPriorityChargingConfirmation struct {
	CustomData *NotifyPriorityChargingConfirmationCustomData `json:"customData,omitempty"`
}

type NotifyPriorityChargingConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value NotifyPriorityChargingConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *NotifyPriorityChargingConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (NotifyPriorityChargingConfirmation) ActionName() string { return "NotifyPriorityCharging" }

func (NotifyPriorityChargingConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (NotifyPriorityChargingConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (NotifyPriorityChargingConfirmation) SchemaName() string {
	return "NotifyPriorityChargingResponse.json"
}

func (message NotifyPriorityChargingConfirmation) Validate() error {
	return validation.Validate("NotifyPriorityChargingResponse.json", schemaNotifyPriorityChargingConfirmation, message)
}

func (NotifyPriorityChargingConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("NotifyPriorityChargingResponse.json", schemaNotifyPriorityChargingConfirmation, data)
}
