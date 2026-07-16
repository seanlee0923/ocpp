// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = FirmwareStatusNotificationRequest{}

var schemaFirmwareStatusNotificationRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Downloaded", "DownloadFailed", "Downloading", "DownloadScheduled", "DownloadPaused", "Idle", "InstallationFailed", "Installing", "Installed", "InstallRebooting", "InstallScheduled", "InstallVerificationFailed", "InvalidSignature", "SignatureVerified"}}, "requestId": &validation.Schema{Type: "integer", AllowAdditional: true}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type FirmwareStatusNotificationRequest struct {
	Status     FirmwareStatusNotificationRequestFirmwareStatusEnum `json:"status"`
	RequestID  *int                                                `json:"requestId,omitempty"`
	StatusInfo *FirmwareStatusNotificationRequestStatusInfo        `json:"statusInfo,omitempty"`
	CustomData *FirmwareStatusNotificationRequestCustomData        `json:"customData,omitempty"`
}

type FirmwareStatusNotificationRequestStatusInfo struct {
	ReasonCode     string                                       `json:"reasonCode"`
	AdditionalInfo *string                                      `json:"additionalInfo,omitempty"`
	CustomData     *FirmwareStatusNotificationRequestCustomData `json:"customData,omitempty"`
}

type FirmwareStatusNotificationRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value FirmwareStatusNotificationRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *FirmwareStatusNotificationRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
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

func (FirmwareStatusNotificationRequest) ActionName() string { return "FirmwareStatusNotification" }

func (FirmwareStatusNotificationRequest) Version() protocol.Version { return protocol.OCPP21 }

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
