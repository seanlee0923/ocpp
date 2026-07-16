// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = FirmwareStatusNotificationRequest{}

var schemaFirmwareStatusNotificationRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Downloaded", "DownloadFailed", "Downloading", "Idle", "InstallationFailed", "Installing", "Installed"}}}, Required: []string{"status"}}

type FirmwareStatusNotificationRequest struct {
	Status FirmwareStatusNotificationRequestStatus `json:"status"`
}

type FirmwareStatusNotificationRequestStatus string

const (
	FirmwareStatusNotificationRequestStatusDownloaded         FirmwareStatusNotificationRequestStatus = "Downloaded"
	FirmwareStatusNotificationRequestStatusDownloadFailed     FirmwareStatusNotificationRequestStatus = "DownloadFailed"
	FirmwareStatusNotificationRequestStatusDownloading        FirmwareStatusNotificationRequestStatus = "Downloading"
	FirmwareStatusNotificationRequestStatusIdle               FirmwareStatusNotificationRequestStatus = "Idle"
	FirmwareStatusNotificationRequestStatusInstallationFailed FirmwareStatusNotificationRequestStatus = "InstallationFailed"
	FirmwareStatusNotificationRequestStatusInstalling         FirmwareStatusNotificationRequestStatus = "Installing"
	FirmwareStatusNotificationRequestStatusInstalled          FirmwareStatusNotificationRequestStatus = "Installed"
)

func (FirmwareStatusNotificationRequest) ActionName() string { return "FirmwareStatusNotification" }

func (FirmwareStatusNotificationRequest) Version() protocol.Version { return protocol.OCPP16 }

func (FirmwareStatusNotificationRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (FirmwareStatusNotificationRequest) SchemaName() string {
	return "FirmwareStatusNotification.json"
}

func (message FirmwareStatusNotificationRequest) Validate() error {
	return validation.Validate("FirmwareStatusNotification.json", schemaFirmwareStatusNotificationRequest, message)
}

func (FirmwareStatusNotificationRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("FirmwareStatusNotification.json", schemaFirmwareStatusNotificationRequest, data)
}
