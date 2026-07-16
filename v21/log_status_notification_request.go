// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = LogStatusNotificationRequest{}

var schemaLogStatusNotificationRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"BadMessage", "Idle", "NotSupportedOperation", "PermissionDenied", "Uploaded", "UploadFailure", "Uploading", "AcceptedCanceled"}}, "requestId": &validation.Schema{Type: "integer", AllowAdditional: true}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type LogStatusNotificationRequest struct {
	Status     LogStatusNotificationRequestUploadLogStatusEnum `json:"status"`
	RequestID  *int                                            `json:"requestId,omitempty"`
	StatusInfo *LogStatusNotificationRequestStatusInfo         `json:"statusInfo,omitempty"`
	CustomData *LogStatusNotificationRequestCustomData         `json:"customData,omitempty"`
}

type LogStatusNotificationRequestStatusInfo struct {
	ReasonCode     string                                  `json:"reasonCode"`
	AdditionalInfo *string                                 `json:"additionalInfo,omitempty"`
	CustomData     *LogStatusNotificationRequestCustomData `json:"customData,omitempty"`
}

type LogStatusNotificationRequestCustomData struct {
	VendorID string `json:"vendorId"`
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

func (LogStatusNotificationRequest) ActionName() string { return "LogStatusNotification" }

func (LogStatusNotificationRequest) Version() protocol.Version { return protocol.OCPP21 }

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
