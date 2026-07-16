// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = NotifyCustomerInformationRequest{}

var schemaNotifyCustomerInformationRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "data": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}, "tbc": &validation.Schema{Type: "boolean", AllowAdditional: true}, "seqNo": &validation.Schema{Type: "integer", AllowAdditional: true}, "generatedAt": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "requestId": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"data", "seqNo", "generatedAt", "requestId"}}

type NotifyCustomerInformationRequest struct {
	CustomData  *NotifyCustomerInformationRequestCustomData `json:"customData,omitempty"`
	Data        string                                      `json:"data"`
	Tbc         *bool                                       `json:"tbc,omitempty"`
	SeqNo       int                                         `json:"seqNo"`
	GeneratedAt string                                      `json:"generatedAt"`
	RequestID   int                                         `json:"requestId"`
}

type NotifyCustomerInformationRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (NotifyCustomerInformationRequest) ActionName() string { return "NotifyCustomerInformation" }

func (NotifyCustomerInformationRequest) Version() protocol.Version { return protocol.OCPP201 }

func (NotifyCustomerInformationRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (NotifyCustomerInformationRequest) SchemaName() string {
	return "NotifyCustomerInformationRequest.json"
}

func (message NotifyCustomerInformationRequest) Validate() error {
	return validation.Validate("NotifyCustomerInformationRequest.json", schemaNotifyCustomerInformationRequest, message)
}

func (NotifyCustomerInformationRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("NotifyCustomerInformationRequest.json", schemaNotifyCustomerInformationRequest, data)
}
