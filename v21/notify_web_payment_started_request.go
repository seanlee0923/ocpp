// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = NotifyWebPaymentStartedRequest{}

var schemaNotifyWebPaymentStartedRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"evseId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "timeout": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"evseId", "timeout"}}

type NotifyWebPaymentStartedRequest struct {
	EVSEID     int                                       `json:"evseId"`
	Timeout    int                                       `json:"timeout"`
	CustomData *NotifyWebPaymentStartedRequestCustomData `json:"customData,omitempty"`
}

type NotifyWebPaymentStartedRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (NotifyWebPaymentStartedRequest) ActionName() string { return "NotifyWebPaymentStarted" }

func (NotifyWebPaymentStartedRequest) Version() protocol.Version { return protocol.OCPP21 }

func (NotifyWebPaymentStartedRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (NotifyWebPaymentStartedRequest) SchemaName() string {
	return "NotifyWebPaymentStartedRequest.json"
}

func (message NotifyWebPaymentStartedRequest) Validate() error {
	return validation.Validate("NotifyWebPaymentStartedRequest.json", schemaNotifyWebPaymentStartedRequest, message)
}

func (NotifyWebPaymentStartedRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("NotifyWebPaymentStartedRequest.json", schemaNotifyWebPaymentStartedRequest, data)
}
