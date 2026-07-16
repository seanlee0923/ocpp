// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ClearCacheConfirmation{}

var schemaClearCacheConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type ClearCacheConfirmation struct {
	Status     ClearCacheConfirmationClearCacheStatusEnum `json:"status"`
	StatusInfo *ClearCacheConfirmationStatusInfo          `json:"statusInfo,omitempty"`
	CustomData *ClearCacheConfirmationCustomData          `json:"customData,omitempty"`
}

type ClearCacheConfirmationStatusInfo struct {
	ReasonCode     string                            `json:"reasonCode"`
	AdditionalInfo *string                           `json:"additionalInfo,omitempty"`
	CustomData     *ClearCacheConfirmationCustomData `json:"customData,omitempty"`
}

type ClearCacheConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

type ClearCacheConfirmationClearCacheStatusEnum string

const (
	ClearCacheConfirmationClearCacheStatusEnumAccepted ClearCacheConfirmationClearCacheStatusEnum = "Accepted"
	ClearCacheConfirmationClearCacheStatusEnumRejected ClearCacheConfirmationClearCacheStatusEnum = "Rejected"
)

func (ClearCacheConfirmation) ActionName() string { return "ClearCache" }

func (ClearCacheConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (ClearCacheConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (ClearCacheConfirmation) SchemaName() string { return "ClearCacheResponse.json" }

func (message ClearCacheConfirmation) Validate() error {
	return validation.Validate("ClearCacheResponse.json", schemaClearCacheConfirmation, message)
}

func (ClearCacheConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ClearCacheResponse.json", schemaClearCacheConfirmation, data)
}
