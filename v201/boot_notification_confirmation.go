// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = BootNotificationConfirmation{}

var schemaBootNotificationConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "currentTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "interval": &validation.Schema{Type: "integer", AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Pending", "Rejected"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}}, Required: []string{"currentTime", "interval", "status"}}

type BootNotificationConfirmation struct {
	CustomData  *BootNotificationConfirmationCustomData            `json:"customData,omitempty"`
	CurrentTime string                                             `json:"currentTime"`
	Interval    int                                                `json:"interval"`
	Status      BootNotificationConfirmationRegistrationStatusEnum `json:"status"`
	StatusInfo  *BootNotificationConfirmationStatusInfo            `json:"statusInfo,omitempty"`
}

type BootNotificationConfirmationStatusInfo struct {
	CustomData     *BootNotificationConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                                  `json:"reasonCode"`
	AdditionalInfo *string                                 `json:"additionalInfo,omitempty"`
}

type BootNotificationConfirmationRegistrationStatusEnum string

const (
	BootNotificationConfirmationRegistrationStatusEnumAccepted BootNotificationConfirmationRegistrationStatusEnum = "Accepted"
	BootNotificationConfirmationRegistrationStatusEnumPending  BootNotificationConfirmationRegistrationStatusEnum = "Pending"
	BootNotificationConfirmationRegistrationStatusEnumRejected BootNotificationConfirmationRegistrationStatusEnum = "Rejected"
)

type BootNotificationConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (BootNotificationConfirmation) ActionName() string { return "BootNotification" }

func (BootNotificationConfirmation) Version() protocol.Version { return protocol.OCPP201 }

func (BootNotificationConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (BootNotificationConfirmation) SchemaName() string { return "BootNotificationResponse.json" }

func (message BootNotificationConfirmation) Validate() error {
	return validation.Validate("BootNotificationResponse.json", schemaBootNotificationConfirmation, message)
}

func (BootNotificationConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("BootNotificationResponse.json", schemaBootNotificationConfirmation, data)
}
