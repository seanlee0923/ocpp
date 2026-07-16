// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = NotifyCustomerInformationRequest{}

var schemaNotifyCustomerInformationRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"data": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}, "tbc": &validation.Schema{Type: "boolean", AllowAdditional: true}, "seqNo": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "generatedAt": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "requestId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"data", "seqNo", "generatedAt", "requestId"}}

type NotifyCustomerInformationRequest struct {
	Data        string                                      `json:"data"`
	Tbc         *bool                                       `json:"tbc,omitempty"`
	SeqNo       int                                         `json:"seqNo"`
	GeneratedAt string                                      `json:"generatedAt"`
	RequestID   int                                         `json:"requestId"`
	CustomData  *NotifyCustomerInformationRequestCustomData `json:"customData,omitempty"`
}

type NotifyCustomerInformationRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (NotifyCustomerInformationRequest) ActionName() string { return "NotifyCustomerInformation" }

func (NotifyCustomerInformationRequest) Version() protocol.Version { return protocol.OCPP21 }

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
