// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = LogStatusNotificationRequest{}

var schemaLogStatusNotificationRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"BadMessage", "Idle", "NotSupportedOperation", "PermissionDenied", "Uploaded", "UploadFailure", "Uploading", "AcceptedCanceled"}}, "requestId": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"status"}}

type LogStatusNotificationRequest struct {
	CustomData *LogStatusNotificationRequestCustomData         `json:"customData,omitempty"`
	Status     LogStatusNotificationRequestUploadLogStatusEnum `json:"status"`
	RequestID  *int                                            `json:"requestId,omitempty"`
}

type LogStatusNotificationRequestUploadLogStatusEnum string

const (
	LogStatusNotificationRequestUploadLogStatusEnumBadMessage            LogStatusNotificationRequestUploadLogStatusEnum = "BadMessage"
	LogStatusNotificationRequestUploadLogStatusEnumIdle                  LogStatusNotificationRequestUploadLogStatusEnum = "Idle"
	LogStatusNotificationRequestUploadLogStatusEnumNotSupportedOperation LogStatusNotificationRequestUploadLogStatusEnum = "NotSupportedOperation"
	LogStatusNotificationRequestUploadLogStatusEnumPermissionDenied      LogStatusNotificationRequestUploadLogStatusEnum = "PermissionDenied"
	LogStatusNotificationRequestUploadLogStatusEnumUploaded              LogStatusNotificationRequestUploadLogStatusEnum = "Uploaded"
	LogStatusNotificationRequestUploadLogStatusEnumUploadFailure         LogStatusNotificationRequestUploadLogStatusEnum = "UploadFailure"
	LogStatusNotificationRequestUploadLogStatusEnumUploading             LogStatusNotificationRequestUploadLogStatusEnum = "Uploading"
	LogStatusNotificationRequestUploadLogStatusEnumAcceptedCanceled      LogStatusNotificationRequestUploadLogStatusEnum = "AcceptedCanceled"
)

type LogStatusNotificationRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (LogStatusNotificationRequest) ActionName() string { return "LogStatusNotification" }

func (LogStatusNotificationRequest) Version() protocol.Version { return protocol.OCPP201 }

func (LogStatusNotificationRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (LogStatusNotificationRequest) SchemaName() string { return "LogStatusNotificationRequest.json" }

func (message LogStatusNotificationRequest) Validate() error {
	return validation.Validate("LogStatusNotificationRequest.json", schemaLogStatusNotificationRequest, message)
}

func (LogStatusNotificationRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("LogStatusNotificationRequest.json", schemaLogStatusNotificationRequest, data)
}
