// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = PublishFirmwareStatusNotificationRequest{}

var schemaPublishFirmwareStatusNotificationRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Idle", "DownloadScheduled", "Downloading", "Downloaded", "Published", "DownloadFailed", "DownloadPaused", "InvalidChecksum", "ChecksumVerified", "PublishFailed"}}, "location": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}, MinItems: 1, HasMinItems: true}, "requestId": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"status"}}

type PublishFirmwareStatusNotificationRequest struct {
	CustomData *PublishFirmwareStatusNotificationRequestCustomData               `json:"customData,omitempty"`
	Status     PublishFirmwareStatusNotificationRequestPublishFirmwareStatusEnum `json:"status"`
	Location   []string                                                          `json:"location,omitempty"`
	RequestID  *int                                                              `json:"requestId,omitempty"`
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

type PublishFirmwareStatusNotificationRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (PublishFirmwareStatusNotificationRequest) ActionName() string {
	return "PublishFirmwareStatusNotification"
}

func (PublishFirmwareStatusNotificationRequest) Version() protocol.Version { return protocol.OCPP201 }

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
