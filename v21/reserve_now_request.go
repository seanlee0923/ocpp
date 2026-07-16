// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ReserveNowRequest{}

var schemaReserveNowRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"id": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "expiryDateTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "connectorType": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "idToken": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"additionalInfo": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"additionalIdToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}, "type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"additionalIdToken", "type"}}, MinItems: 1, HasMinItems: true}, "idToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}, "type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"idToken", "type"}}, "evseId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "groupIdToken": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"additionalInfo": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"additionalIdToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}, "type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"additionalIdToken", "type"}}, MinItems: 1, HasMinItems: true}, "idToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}, "type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"idToken", "type"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id", "expiryDateTime", "idToken"}}

type ReserveNowRequest struct {
	ID             int                          `json:"id"`
	ExpiryDateTime string                       `json:"expiryDateTime"`
	ConnectorType  *string                      `json:"connectorType,omitempty"`
	IDToken        ReserveNowRequestIDToken     `json:"idToken"`
	EVSEID         *int                         `json:"evseId,omitempty"`
	GroupIDToken   *ReserveNowRequestIDToken    `json:"groupIdToken,omitempty"`
	CustomData     *ReserveNowRequestCustomData `json:"customData,omitempty"`
}

type ReserveNowRequestIDToken struct {
	AdditionalInfo []ReserveNowRequestAdditionalInfo `json:"additionalInfo,omitempty"`
	IDToken        string                            `json:"idToken"`
	Type           string                            `json:"type"`
	CustomData     *ReserveNowRequestCustomData      `json:"customData,omitempty"`
}

type ReserveNowRequestAdditionalInfo struct {
	AdditionalIDToken string                       `json:"additionalIdToken"`
	Type              string                       `json:"type"`
	CustomData        *ReserveNowRequestCustomData `json:"customData,omitempty"`
}

type ReserveNowRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (ReserveNowRequest) ActionName() string { return "ReserveNow" }

func (ReserveNowRequest) Version() protocol.Version { return protocol.OCPP21 }

func (ReserveNowRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (ReserveNowRequest) SchemaName() string { return "ReserveNowRequest.json" }

func (message ReserveNowRequest) Validate() error {
	return validation.Validate("ReserveNowRequest.json", schemaReserveNowRequest, message)
}

func (ReserveNowRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ReserveNowRequest.json", schemaReserveNowRequest, data)
}
