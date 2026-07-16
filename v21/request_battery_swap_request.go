// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = RequestBatterySwapRequest{}

var schemaRequestBatterySwapRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"idToken": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"additionalInfo": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"additionalIdToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}, "type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"additionalIdToken", "type"}}, MinItems: 1, HasMinItems: true}, "idToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}, "type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"idToken", "type"}}, "requestId": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"requestId", "idToken"}}

type RequestBatterySwapRequest struct {
	IDToken    RequestBatterySwapRequestIDToken     `json:"idToken"`
	RequestID  int                                  `json:"requestId"`
	CustomData *RequestBatterySwapRequestCustomData `json:"customData,omitempty"`
}

type RequestBatterySwapRequestIDToken struct {
	AdditionalInfo []RequestBatterySwapRequestAdditionalInfo `json:"additionalInfo,omitempty"`
	IDToken        string                                    `json:"idToken"`
	Type           string                                    `json:"type"`
	CustomData     *RequestBatterySwapRequestCustomData      `json:"customData,omitempty"`
}

type RequestBatterySwapRequestAdditionalInfo struct {
	AdditionalIDToken string                               `json:"additionalIdToken"`
	Type              string                               `json:"type"`
	CustomData        *RequestBatterySwapRequestCustomData `json:"customData,omitempty"`
}

type RequestBatterySwapRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (RequestBatterySwapRequest) ActionName() string { return "RequestBatterySwap" }

func (RequestBatterySwapRequest) Version() protocol.Version { return protocol.OCPP21 }

func (RequestBatterySwapRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (RequestBatterySwapRequest) SchemaName() string { return "RequestBatterySwapRequest.json" }

func (message RequestBatterySwapRequest) Validate() error {
	return validation.Validate("RequestBatterySwapRequest.json", schemaRequestBatterySwapRequest, message)
}

func (RequestBatterySwapRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("RequestBatterySwapRequest.json", schemaRequestBatterySwapRequest, data)
}
