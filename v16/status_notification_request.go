// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = StatusNotificationRequest{}

var schemaStatusNotificationRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"connectorId": &validation.Schema{Type: "integer", AllowAdditional: true}, "errorCode": &validation.Schema{Type: "string", Enum: []string{"ConnectorLockFailure", "EVCommunicationError", "GroundFailure", "HighTemperature", "InternalError", "LocalListConflict", "NoError", "OtherError", "OverCurrentFailure", "PowerMeterFailure", "PowerSwitchFailure", "ReaderFailure", "ResetFailure", "UnderVoltage", "OverVoltage", "WeakSignal"}}, "info": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Available", "Preparing", "Charging", "SuspendedEVSE", "SuspendedEV", "Finishing", "Reserved", "Unavailable", "Faulted"}}, "timestamp": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}, "vendorErrorCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}}, Required: []string{"connectorId", "errorCode", "status"}}

type StatusNotificationRequest struct {
	ConnectorID     int                                `json:"connectorId"`
	ErrorCode       StatusNotificationRequestErrorCode `json:"errorCode"`
	Info            *string                            `json:"info,omitempty"`
	Status          StatusNotificationRequestStatus    `json:"status"`
	Timestamp       *string                            `json:"timestamp,omitempty"`
	VendorID        *string                            `json:"vendorId,omitempty"`
	VendorErrorCode *string                            `json:"vendorErrorCode,omitempty"`
}

type StatusNotificationRequestStatus string

const (
	StatusNotificationRequestStatusAvailable     StatusNotificationRequestStatus = "Available"
	StatusNotificationRequestStatusPreparing     StatusNotificationRequestStatus = "Preparing"
	StatusNotificationRequestStatusCharging      StatusNotificationRequestStatus = "Charging"
	StatusNotificationRequestStatusSuspendedEVSE StatusNotificationRequestStatus = "SuspendedEVSE"
	StatusNotificationRequestStatusSuspendedEV   StatusNotificationRequestStatus = "SuspendedEV"
	StatusNotificationRequestStatusFinishing     StatusNotificationRequestStatus = "Finishing"
	StatusNotificationRequestStatusReserved      StatusNotificationRequestStatus = "Reserved"
	StatusNotificationRequestStatusUnavailable   StatusNotificationRequestStatus = "Unavailable"
	StatusNotificationRequestStatusFaulted       StatusNotificationRequestStatus = "Faulted"
)

type StatusNotificationRequestErrorCode string

const (
	StatusNotificationRequestErrorCodeConnectorLockFailure StatusNotificationRequestErrorCode = "ConnectorLockFailure"
	StatusNotificationRequestErrorCodeEVCommunicationError StatusNotificationRequestErrorCode = "EVCommunicationError"
	StatusNotificationRequestErrorCodeGroundFailure        StatusNotificationRequestErrorCode = "GroundFailure"
	StatusNotificationRequestErrorCodeHighTemperature      StatusNotificationRequestErrorCode = "HighTemperature"
	StatusNotificationRequestErrorCodeInternalError        StatusNotificationRequestErrorCode = "InternalError"
	StatusNotificationRequestErrorCodeLocalListConflict    StatusNotificationRequestErrorCode = "LocalListConflict"
	StatusNotificationRequestErrorCodeNoError              StatusNotificationRequestErrorCode = "NoError"
	StatusNotificationRequestErrorCodeOtherError           StatusNotificationRequestErrorCode = "OtherError"
	StatusNotificationRequestErrorCodeOverCurrentFailure   StatusNotificationRequestErrorCode = "OverCurrentFailure"
	StatusNotificationRequestErrorCodePowerMeterFailure    StatusNotificationRequestErrorCode = "PowerMeterFailure"
	StatusNotificationRequestErrorCodePowerSwitchFailure   StatusNotificationRequestErrorCode = "PowerSwitchFailure"
	StatusNotificationRequestErrorCodeReaderFailure        StatusNotificationRequestErrorCode = "ReaderFailure"
	StatusNotificationRequestErrorCodeResetFailure         StatusNotificationRequestErrorCode = "ResetFailure"
	StatusNotificationRequestErrorCodeUnderVoltage         StatusNotificationRequestErrorCode = "UnderVoltage"
	StatusNotificationRequestErrorCodeOverVoltage          StatusNotificationRequestErrorCode = "OverVoltage"
	StatusNotificationRequestErrorCodeWeakSignal           StatusNotificationRequestErrorCode = "WeakSignal"
)

func (StatusNotificationRequest) ActionName() string { return "StatusNotification" }

func (StatusNotificationRequest) Version() protocol.Version { return protocol.OCPP16 }

func (StatusNotificationRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (StatusNotificationRequest) SchemaName() string { return "StatusNotification.json" }

func (message StatusNotificationRequest) Validate() error {
	return validation.Validate("StatusNotification.json", schemaStatusNotificationRequest, message)
}

func (StatusNotificationRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("StatusNotification.json", schemaStatusNotificationRequest, data)
}
