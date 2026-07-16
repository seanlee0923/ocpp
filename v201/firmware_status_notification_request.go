// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = FirmwareStatusNotificationRequest{}

var schemaFirmwareStatusNotificationRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Downloaded", "DownloadFailed", "Downloading", "DownloadScheduled", "DownloadPaused", "Idle", "InstallationFailed", "Installing", "Installed", "InstallRebooting", "InstallScheduled", "InstallVerificationFailed", "InvalidSignature", "SignatureVerified"}}, "requestId": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"status"}}

type FirmwareStatusNotificationRequest struct {
	CustomData *FirmwareStatusNotificationRequestCustomData        `json:"customData,omitempty"`
	Status     FirmwareStatusNotificationRequestFirmwareStatusEnum `json:"status"`
	RequestID  *int                                                `json:"requestId,omitempty"`
}

type FirmwareStatusNotificationRequestFirmwareStatusEnum string

const (
	FirmwareStatusNotificationRequestFirmwareStatusEnumDownloaded                FirmwareStatusNotificationRequestFirmwareStatusEnum = "Downloaded"
	FirmwareStatusNotificationRequestFirmwareStatusEnumDownloadFailed            FirmwareStatusNotificationRequestFirmwareStatusEnum = "DownloadFailed"
	FirmwareStatusNotificationRequestFirmwareStatusEnumDownloading               FirmwareStatusNotificationRequestFirmwareStatusEnum = "Downloading"
	FirmwareStatusNotificationRequestFirmwareStatusEnumDownloadScheduled         FirmwareStatusNotificationRequestFirmwareStatusEnum = "DownloadScheduled"
	FirmwareStatusNotificationRequestFirmwareStatusEnumDownloadPaused            FirmwareStatusNotificationRequestFirmwareStatusEnum = "DownloadPaused"
	FirmwareStatusNotificationRequestFirmwareStatusEnumIdle                      FirmwareStatusNotificationRequestFirmwareStatusEnum = "Idle"
	FirmwareStatusNotificationRequestFirmwareStatusEnumInstallationFailed        FirmwareStatusNotificationRequestFirmwareStatusEnum = "InstallationFailed"
	FirmwareStatusNotificationRequestFirmwareStatusEnumInstalling                FirmwareStatusNotificationRequestFirmwareStatusEnum = "Installing"
	FirmwareStatusNotificationRequestFirmwareStatusEnumInstalled                 FirmwareStatusNotificationRequestFirmwareStatusEnum = "Installed"
	FirmwareStatusNotificationRequestFirmwareStatusEnumInstallRebooting          FirmwareStatusNotificationRequestFirmwareStatusEnum = "InstallRebooting"
	FirmwareStatusNotificationRequestFirmwareStatusEnumInstallScheduled          FirmwareStatusNotificationRequestFirmwareStatusEnum = "InstallScheduled"
	FirmwareStatusNotificationRequestFirmwareStatusEnumInstallVerificationFailed FirmwareStatusNotificationRequestFirmwareStatusEnum = "InstallVerificationFailed"
	FirmwareStatusNotificationRequestFirmwareStatusEnumInvalidSignature          FirmwareStatusNotificationRequestFirmwareStatusEnum = "InvalidSignature"
	FirmwareStatusNotificationRequestFirmwareStatusEnumSignatureVerified         FirmwareStatusNotificationRequestFirmwareStatusEnum = "SignatureVerified"
)

type FirmwareStatusNotificationRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (FirmwareStatusNotificationRequest) ActionName() string { return "FirmwareStatusNotification" }

func (FirmwareStatusNotificationRequest) Version() protocol.Version { return protocol.OCPP201 }

func (FirmwareStatusNotificationRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (FirmwareStatusNotificationRequest) SchemaName() string {
	return "FirmwareStatusNotificationRequest.json"
}

func (message FirmwareStatusNotificationRequest) Validate() error {
	return validation.Validate("FirmwareStatusNotificationRequest.json", schemaFirmwareStatusNotificationRequest, message)
}

func (FirmwareStatusNotificationRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("FirmwareStatusNotificationRequest.json", schemaFirmwareStatusNotificationRequest, data)
}
