// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = Get15118EVCertificateRequest{}

var schemaGet15118EVCertificateRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"iso15118SchemaVersion": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "action": &validation.Schema{Type: "string", Enum: []string{"Install", "Update"}}, "exiRequest": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 11000, HasMaxLength: true}, "maximumContractCertificateChains": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "prioritizedEMAIDs": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}, MinItems: 1, HasMinItems: true, MaxItems: 8, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"iso15118SchemaVersion", "action", "exiRequest"}}

type Get15118EVCertificateRequest struct {
	Iso15118SchemaVersion            string                                            `json:"iso15118SchemaVersion"`
	Action                           Get15118EVCertificateRequestCertificateActionEnum `json:"action"`
	ExiRequest                       string                                            `json:"exiRequest"`
	MaximumContractCertificateChains *int                                              `json:"maximumContractCertificateChains,omitempty"`
	PrioritizedEMAIDs                []string                                          `json:"prioritizedEMAIDs,omitempty"`
	CustomData                       *Get15118EVCertificateRequestCustomData           `json:"customData,omitempty"`
}

type Get15118EVCertificateRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

type Get15118EVCertificateRequestCertificateActionEnum string

const (
	Get15118EVCertificateRequestCertificateActionEnumInstall Get15118EVCertificateRequestCertificateActionEnum = "Install"
	Get15118EVCertificateRequestCertificateActionEnumUpdate  Get15118EVCertificateRequestCertificateActionEnum = "Update"
)

func (Get15118EVCertificateRequest) ActionName() string { return "Get15118EVCertificate" }

func (Get15118EVCertificateRequest) Version() protocol.Version { return protocol.OCPP21 }

func (Get15118EVCertificateRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (Get15118EVCertificateRequest) SchemaName() string { return "Get15118EVCertificateRequest.json" }

func (message Get15118EVCertificateRequest) Validate() error {
	return validation.Validate("Get15118EVCertificateRequest.json", schemaGet15118EVCertificateRequest, message)
}

func (Get15118EVCertificateRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("Get15118EVCertificateRequest.json", schemaGet15118EVCertificateRequest, data)
}
