// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = BootNotificationRequest{}

var schemaBootNotificationRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"chargingStation": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"serialNumber": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 25, HasMaxLength: true}, "model": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "modem": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"iccid": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "imsi": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "vendorName": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "firmwareVersion": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"model", "vendorName"}}, "reason": &validation.Schema{Type: "string", Enum: []string{"ApplicationReset", "FirmwareUpdate", "LocalReset", "PowerUp", "RemoteReset", "ScheduledReset", "Triggered", "Unknown", "Watchdog"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reason", "chargingStation"}}

type BootNotificationRequest struct {
	ChargingStation BootNotificationRequestChargingStation `json:"chargingStation"`
	Reason          BootNotificationRequestBootReasonEnum  `json:"reason"`
	CustomData      *BootNotificationRequestCustomData     `json:"customData,omitempty"`
}

type BootNotificationRequestBootReasonEnum string

const (
	BootNotificationRequestBootReasonEnumApplicationReset BootNotificationRequestBootReasonEnum = "ApplicationReset"
	BootNotificationRequestBootReasonEnumFirmwareUpdate   BootNotificationRequestBootReasonEnum = "FirmwareUpdate"
	BootNotificationRequestBootReasonEnumLocalReset       BootNotificationRequestBootReasonEnum = "LocalReset"
	BootNotificationRequestBootReasonEnumPowerUp          BootNotificationRequestBootReasonEnum = "PowerUp"
	BootNotificationRequestBootReasonEnumRemoteReset      BootNotificationRequestBootReasonEnum = "RemoteReset"
	BootNotificationRequestBootReasonEnumScheduledReset   BootNotificationRequestBootReasonEnum = "ScheduledReset"
	BootNotificationRequestBootReasonEnumTriggered        BootNotificationRequestBootReasonEnum = "Triggered"
	BootNotificationRequestBootReasonEnumUnknown          BootNotificationRequestBootReasonEnum = "Unknown"
	BootNotificationRequestBootReasonEnumWatchdog         BootNotificationRequestBootReasonEnum = "Watchdog"
)

type BootNotificationRequestChargingStation struct {
	SerialNumber    *string                            `json:"serialNumber,omitempty"`
	Model           string                             `json:"model"`
	Modem           *BootNotificationRequestModem      `json:"modem,omitempty"`
	VendorName      string                             `json:"vendorName"`
	FirmwareVersion *string                            `json:"firmwareVersion,omitempty"`
	CustomData      *BootNotificationRequestCustomData `json:"customData,omitempty"`
}

type BootNotificationRequestModem struct {
	ICCID      *string                            `json:"iccid,omitempty"`
	IMSI       *string                            `json:"imsi,omitempty"`
	CustomData *BootNotificationRequestCustomData `json:"customData,omitempty"`
}

type BootNotificationRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value BootNotificationRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *BootNotificationRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (BootNotificationRequest) ActionName() string { return "BootNotification" }

func (BootNotificationRequest) Version() protocol.Version { return protocol.OCPP21 }

func (BootNotificationRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (BootNotificationRequest) SchemaName() string { return "BootNotificationRequest.json" }

func (message BootNotificationRequest) Validate() error {
	return validation.Validate("BootNotificationRequest.json", schemaBootNotificationRequest, message)
}

func (BootNotificationRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("BootNotificationRequest.json", schemaBootNotificationRequest, data)
}
