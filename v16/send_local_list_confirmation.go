// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = SendLocalListConfirmation{}

var schemaSendLocalListConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Failed", "NotSupported", "VersionMismatch"}}}, Required: []string{"status"}}

type SendLocalListConfirmation struct {
	Status SendLocalListConfirmationStatus `json:"status"`
}

type SendLocalListConfirmationStatus string

const (
	SendLocalListConfirmationStatusAccepted        SendLocalListConfirmationStatus = "Accepted"
	SendLocalListConfirmationStatusFailed          SendLocalListConfirmationStatus = "Failed"
	SendLocalListConfirmationStatusNotSupported    SendLocalListConfirmationStatus = "NotSupported"
	SendLocalListConfirmationStatusVersionMismatch SendLocalListConfirmationStatus = "VersionMismatch"
)

func (SendLocalListConfirmation) ActionName() string { return "SendLocalList" }

func (SendLocalListConfirmation) Version() protocol.Version { return protocol.OCPP16 }

func (SendLocalListConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (SendLocalListConfirmation) SchemaName() string { return "SendLocalListResponse.json" }

func (message SendLocalListConfirmation) Validate() error {
	return validation.Validate("SendLocalListResponse.json", schemaSendLocalListConfirmation, message)
}

func (SendLocalListConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("SendLocalListResponse.json", schemaSendLocalListConfirmation, data)
}
