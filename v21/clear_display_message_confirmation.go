// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ClearDisplayMessageConfirmation{}

var schemaClearDisplayMessageConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Unknown", "Rejected"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type ClearDisplayMessageConfirmation struct {
	Status     ClearDisplayMessageConfirmationClearMessageStatusEnum `json:"status"`
	StatusInfo *ClearDisplayMessageConfirmationStatusInfo            `json:"statusInfo,omitempty"`
	CustomData *ClearDisplayMessageConfirmationCustomData            `json:"customData,omitempty"`
}

type ClearDisplayMessageConfirmationStatusInfo struct {
	ReasonCode     string                                     `json:"reasonCode"`
	AdditionalInfo *string                                    `json:"additionalInfo,omitempty"`
	CustomData     *ClearDisplayMessageConfirmationCustomData `json:"customData,omitempty"`
}

type ClearDisplayMessageConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

type ClearDisplayMessageConfirmationClearMessageStatusEnum string

const (
	ClearDisplayMessageConfirmationClearMessageStatusEnumAccepted ClearDisplayMessageConfirmationClearMessageStatusEnum = "Accepted"
	ClearDisplayMessageConfirmationClearMessageStatusEnumUnknown  ClearDisplayMessageConfirmationClearMessageStatusEnum = "Unknown"
	ClearDisplayMessageConfirmationClearMessageStatusEnumRejected ClearDisplayMessageConfirmationClearMessageStatusEnum = "Rejected"
)

func (ClearDisplayMessageConfirmation) ActionName() string { return "ClearDisplayMessage" }

func (ClearDisplayMessageConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (ClearDisplayMessageConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (ClearDisplayMessageConfirmation) SchemaName() string { return "ClearDisplayMessageResponse.json" }

func (message ClearDisplayMessageConfirmation) Validate() error {
	return validation.Validate("ClearDisplayMessageResponse.json", schemaClearDisplayMessageConfirmation, message)
}

func (ClearDisplayMessageConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ClearDisplayMessageResponse.json", schemaClearDisplayMessageConfirmation, data)
}
