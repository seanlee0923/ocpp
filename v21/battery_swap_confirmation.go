// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = BatterySwapConfirmation{}

var schemaBatterySwapConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type BatterySwapConfirmation struct {
	CustomData *BatterySwapConfirmationCustomData `json:"customData,omitempty"`
}

type BatterySwapConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (BatterySwapConfirmation) ActionName() string { return "BatterySwap" }

func (BatterySwapConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (BatterySwapConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (BatterySwapConfirmation) SchemaName() string { return "BatterySwapResponse.json" }

func (message BatterySwapConfirmation) Validate() error {
	return validation.Validate("BatterySwapResponse.json", schemaBatterySwapConfirmation, message)
}

func (BatterySwapConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("BatterySwapResponse.json", schemaBatterySwapConfirmation, data)
}
