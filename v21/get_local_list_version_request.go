// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = GetLocalListVersionRequest{}

var schemaGetLocalListVersionRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type GetLocalListVersionRequest struct {
	CustomData *GetLocalListVersionRequestCustomData `json:"customData,omitempty"`
}

type GetLocalListVersionRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (GetLocalListVersionRequest) ActionName() string { return "GetLocalListVersion" }

func (GetLocalListVersionRequest) Version() protocol.Version { return protocol.OCPP21 }

func (GetLocalListVersionRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (GetLocalListVersionRequest) SchemaName() string { return "GetLocalListVersionRequest.json" }

func (message GetLocalListVersionRequest) Validate() error {
	return validation.Validate("GetLocalListVersionRequest.json", schemaGetLocalListVersionRequest, message)
}

func (GetLocalListVersionRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetLocalListVersionRequest.json", schemaGetLocalListVersionRequest, data)
}
