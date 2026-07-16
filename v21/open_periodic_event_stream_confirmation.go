// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = OpenPeriodicEventStreamConfirmation{}

var schemaOpenPeriodicEventStreamConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type OpenPeriodicEventStreamConfirmation struct {
	Status     OpenPeriodicEventStreamConfirmationGenericStatusEnum `json:"status"`
	StatusInfo *OpenPeriodicEventStreamConfirmationStatusInfo       `json:"statusInfo,omitempty"`
	CustomData *OpenPeriodicEventStreamConfirmationCustomData       `json:"customData,omitempty"`
}

type OpenPeriodicEventStreamConfirmationStatusInfo struct {
	ReasonCode     string                                         `json:"reasonCode"`
	AdditionalInfo *string                                        `json:"additionalInfo,omitempty"`
	CustomData     *OpenPeriodicEventStreamConfirmationCustomData `json:"customData,omitempty"`
}

type OpenPeriodicEventStreamConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

type OpenPeriodicEventStreamConfirmationGenericStatusEnum string

const (
	OpenPeriodicEventStreamConfirmationGenericStatusEnumAccepted OpenPeriodicEventStreamConfirmationGenericStatusEnum = "Accepted"
	OpenPeriodicEventStreamConfirmationGenericStatusEnumRejected OpenPeriodicEventStreamConfirmationGenericStatusEnum = "Rejected"
)

func (OpenPeriodicEventStreamConfirmation) ActionName() string { return "OpenPeriodicEventStream" }

func (OpenPeriodicEventStreamConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (OpenPeriodicEventStreamConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (OpenPeriodicEventStreamConfirmation) SchemaName() string {
	return "OpenPeriodicEventStreamResponse.json"
}

func (message OpenPeriodicEventStreamConfirmation) Validate() error {
	return validation.Validate("OpenPeriodicEventStreamResponse.json", schemaOpenPeriodicEventStreamConfirmation, message)
}

func (OpenPeriodicEventStreamConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("OpenPeriodicEventStreamResponse.json", schemaOpenPeriodicEventStreamConfirmation, data)
}
