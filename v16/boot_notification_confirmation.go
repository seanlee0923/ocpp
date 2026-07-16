// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = BootNotificationConfirmation{}

var schemaBootNotificationConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Pending", "Rejected"}}, "currentTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "interval": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"status", "currentTime", "interval"}}

type BootNotificationConfirmation struct {
	Status      BootNotificationConfirmationStatus `json:"status"`
	CurrentTime string                             `json:"currentTime"`
	Interval    int                                `json:"interval"`
}

type BootNotificationConfirmationStatus string

const (
	BootNotificationConfirmationStatusAccepted BootNotificationConfirmationStatus = "Accepted"
	BootNotificationConfirmationStatusPending  BootNotificationConfirmationStatus = "Pending"
	BootNotificationConfirmationStatusRejected BootNotificationConfirmationStatus = "Rejected"
)

func (BootNotificationConfirmation) ActionName() string { return "BootNotification" }

func (BootNotificationConfirmation) Version() protocol.Version { return protocol.OCPP16 }

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
