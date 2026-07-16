// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = BatterySwapRequest{}

var schemaBatterySwapRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"batteryData": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"evseId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "serialNumber": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "soC": &validation.Schema{Type: "number", AllowAdditional: true, Minimum: 0, HasMinimum: true, Maximum: 100, HasMaximum: true}, "soH": &validation.Schema{Type: "number", AllowAdditional: true, Minimum: 0, HasMinimum: true, Maximum: 100, HasMaximum: true}, "productionDate": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "vendorInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 500, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"evseId", "serialNumber", "soC", "soH"}}, MinItems: 1, HasMinItems: true}, "eventType": &validation.Schema{Type: "string", Enum: []string{"BatteryIn", "BatteryOut", "BatteryOutTimeout"}}, "idToken": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"additionalInfo": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"additionalIdToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}, "type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"additionalIdToken", "type"}}, MinItems: 1, HasMinItems: true}, "idToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}, "type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"idToken", "type"}}, "requestId": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"eventType", "requestId", "idToken", "batteryData"}}

type BatterySwapRequest struct {
	BatteryData []BatterySwapRequestBatteryData        `json:"batteryData"`
	EventType   BatterySwapRequestBatterySwapEventEnum `json:"eventType"`
	IDToken     BatterySwapRequestIDToken              `json:"idToken"`
	RequestID   int                                    `json:"requestId"`
	CustomData  *BatterySwapRequestCustomData          `json:"customData,omitempty"`
}

type BatterySwapRequestIDToken struct {
	AdditionalInfo []BatterySwapRequestAdditionalInfo `json:"additionalInfo,omitempty"`
	IDToken        string                             `json:"idToken"`
	Type           string                             `json:"type"`
	CustomData     *BatterySwapRequestCustomData      `json:"customData,omitempty"`
}

type BatterySwapRequestAdditionalInfo struct {
	AdditionalIDToken string                        `json:"additionalIdToken"`
	Type              string                        `json:"type"`
	CustomData        *BatterySwapRequestCustomData `json:"customData,omitempty"`
}

type BatterySwapRequestBatterySwapEventEnum string

const (
	BatterySwapRequestBatterySwapEventEnumBatteryIn         BatterySwapRequestBatterySwapEventEnum = "BatteryIn"
	BatterySwapRequestBatterySwapEventEnumBatteryOut        BatterySwapRequestBatterySwapEventEnum = "BatteryOut"
	BatterySwapRequestBatterySwapEventEnumBatteryOutTimeout BatterySwapRequestBatterySwapEventEnum = "BatteryOutTimeout"
)

type BatterySwapRequestBatteryData struct {
	EVSEID         int                           `json:"evseId"`
	SerialNumber   string                        `json:"serialNumber"`
	SoC            float64                       `json:"soC"`
	SoH            float64                       `json:"soH"`
	ProductionDate *string                       `json:"productionDate,omitempty"`
	VendorInfo     *string                       `json:"vendorInfo,omitempty"`
	CustomData     *BatterySwapRequestCustomData `json:"customData,omitempty"`
}

type BatterySwapRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (BatterySwapRequest) ActionName() string { return "BatterySwap" }

func (BatterySwapRequest) Version() protocol.Version { return protocol.OCPP21 }

func (BatterySwapRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (BatterySwapRequest) SchemaName() string { return "BatterySwapRequest.json" }

func (message BatterySwapRequest) Validate() error {
	return validation.Validate("BatterySwapRequest.json", schemaBatterySwapRequest, message)
}

func (BatterySwapRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("BatterySwapRequest.json", schemaBatterySwapRequest, data)
}
