// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = PublishFirmwareStatusNotificationRequest{}

var schemaPublishFirmwareStatusNotificationRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Idle", "DownloadScheduled", "Downloading", "Downloaded", "Published", "DownloadFailed", "DownloadPaused", "InvalidChecksum", "ChecksumVerified", "PublishFailed"}}, "location": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 2000, HasMaxLength: true}, MinItems: 1, HasMinItems: true}, "requestId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type PublishFirmwareStatusNotificationRequest struct {
	Status     PublishFirmwareStatusNotificationRequestPublishFirmwareStatusEnum `json:"status"`
	Location   []string                                                          `json:"location,omitempty"`
	RequestID  *int                                                              `json:"requestId,omitempty"`
	StatusInfo *PublishFirmwareStatusNotificationRequestStatusInfo               `json:"statusInfo,omitempty"`
	CustomData *PublishFirmwareStatusNotificationRequestCustomData               `json:"customData,omitempty"`
}

type PublishFirmwareStatusNotificationRequestStatusInfo struct {
	ReasonCode     string                                              `json:"reasonCode"`
	AdditionalInfo *string                                             `json:"additionalInfo,omitempty"`
	CustomData     *PublishFirmwareStatusNotificationRequestCustomData `json:"customData,omitempty"`
}

type PublishFirmwareStatusNotificationRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

type PublishFirmwareStatusNotificationRequestPublishFirmwareStatusEnum string

const (
	PublishFirmwareStatusNotificationRequestPublishFirmwareStatusEnumIdle              PublishFirmwareStatusNotificationRequestPublishFirmwareStatusEnum = "Idle"
	PublishFirmwareStatusNotificationRequestPublishFirmwareStatusEnumDownloadScheduled PublishFirmwareStatusNotificationRequestPublishFirmwareStatusEnum = "DownloadScheduled"
	PublishFirmwareStatusNotificationRequestPublishFirmwareStatusEnumDownloading       PublishFirmwareStatusNotificationRequestPublishFirmwareStatusEnum = "Downloading"
	PublishFirmwareStatusNotificationRequestPublishFirmwareStatusEnumDownloaded        PublishFirmwareStatusNotificationRequestPublishFirmwareStatusEnum = "Downloaded"
	PublishFirmwareStatusNotificationRequestPublishFirmwareStatusEnumPublished         PublishFirmwareStatusNotificationRequestPublishFirmwareStatusEnum = "Published"
	PublishFirmwareStatusNotificationRequestPublishFirmwareStatusEnumDownloadFailed    PublishFirmwareStatusNotificationRequestPublishFirmwareStatusEnum = "DownloadFailed"
	PublishFirmwareStatusNotificationRequestPublishFirmwareStatusEnumDownloadPaused    PublishFirmwareStatusNotificationRequestPublishFirmwareStatusEnum = "DownloadPaused"
	PublishFirmwareStatusNotificationRequestPublishFirmwareStatusEnumInvalidChecksum   PublishFirmwareStatusNotificationRequestPublishFirmwareStatusEnum = "InvalidChecksum"
	PublishFirmwareStatusNotificationRequestPublishFirmwareStatusEnumChecksumVerified  PublishFirmwareStatusNotificationRequestPublishFirmwareStatusEnum = "ChecksumVerified"
	PublishFirmwareStatusNotificationRequestPublishFirmwareStatusEnumPublishFailed     PublishFirmwareStatusNotificationRequestPublishFirmwareStatusEnum = "PublishFailed"
)

func (PublishFirmwareStatusNotificationRequest) ActionName() string {
	return "PublishFirmwareStatusNotification"
}

func (PublishFirmwareStatusNotificationRequest) Version() protocol.Version { return protocol.OCPP21 }

func (PublishFirmwareStatusNotificationRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (PublishFirmwareStatusNotificationRequest) SchemaName() string {
	return "PublishFirmwareStatusNotificationRequest.json"
}

func (message PublishFirmwareStatusNotificationRequest) Validate() error {
	return validation.Validate("PublishFirmwareStatusNotificationRequest.json", schemaPublishFirmwareStatusNotificationRequest, message)
}

func (PublishFirmwareStatusNotificationRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("PublishFirmwareStatusNotificationRequest.json", schemaPublishFirmwareStatusNotificationRequest, data)
}
