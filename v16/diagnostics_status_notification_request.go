// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = DiagnosticsStatusNotificationRequest{}

var schemaDiagnosticsStatusNotificationRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Idle", "Uploaded", "UploadFailed", "Uploading"}}}, Required: []string{"status"}}

type DiagnosticsStatusNotificationRequest struct {
	Status DiagnosticsStatusNotificationRequestStatus `json:"status"`
}

type DiagnosticsStatusNotificationRequestStatus string

const (
	DiagnosticsStatusNotificationRequestStatusIdle         DiagnosticsStatusNotificationRequestStatus = "Idle"
	DiagnosticsStatusNotificationRequestStatusUploaded     DiagnosticsStatusNotificationRequestStatus = "Uploaded"
	DiagnosticsStatusNotificationRequestStatusUploadFailed DiagnosticsStatusNotificationRequestStatus = "UploadFailed"
	DiagnosticsStatusNotificationRequestStatusUploading    DiagnosticsStatusNotificationRequestStatus = "Uploading"
)

func (DiagnosticsStatusNotificationRequest) ActionName() string {
	return "DiagnosticsStatusNotification"
}

func (DiagnosticsStatusNotificationRequest) Version() protocol.Version { return protocol.OCPP16 }

func (DiagnosticsStatusNotificationRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (DiagnosticsStatusNotificationRequest) SchemaName() string {
	return "DiagnosticsStatusNotification.json"
}

func (message DiagnosticsStatusNotificationRequest) Validate() error {
	return validation.Validate("DiagnosticsStatusNotification.json", schemaDiagnosticsStatusNotificationRequest, message)
}

func (DiagnosticsStatusNotificationRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("DiagnosticsStatusNotification.json", schemaDiagnosticsStatusNotificationRequest, data)
}
