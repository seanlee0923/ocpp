// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = TriggerMessageRequest{}

var schemaTriggerMessageRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"evse": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"id": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "connectorId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"id"}}, "requestedMessage": &validation.Schema{Type: "string", Enum: []string{"BootNotification", "LogStatusNotification", "FirmwareStatusNotification", "Heartbeat", "MeterValues", "SignChargingStationCertificate", "SignV2GCertificate", "SignV2G20Certificate", "StatusNotification", "TransactionEvent", "SignCombinedCertificate", "PublishFirmwareStatusNotification", "CustomTrigger"}}, "customTrigger": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"requestedMessage"}}

type TriggerMessageRequest struct {
	EVSE             *TriggerMessageRequestEVSE              `json:"evse,omitempty"`
	RequestedMessage TriggerMessageRequestMessageTriggerEnum `json:"requestedMessage"`
	CustomTrigger    *string                                 `json:"customTrigger,omitempty"`
	CustomData       *TriggerMessageRequestCustomData        `json:"customData,omitempty"`
}

type TriggerMessageRequestMessageTriggerEnum string

const (
	TriggerMessageRequestMessageTriggerEnumBootNotification                  TriggerMessageRequestMessageTriggerEnum = "BootNotification"
	TriggerMessageRequestMessageTriggerEnumLogStatusNotification             TriggerMessageRequestMessageTriggerEnum = "LogStatusNotification"
	TriggerMessageRequestMessageTriggerEnumFirmwareStatusNotification        TriggerMessageRequestMessageTriggerEnum = "FirmwareStatusNotification"
	TriggerMessageRequestMessageTriggerEnumHeartbeat                         TriggerMessageRequestMessageTriggerEnum = "Heartbeat"
	TriggerMessageRequestMessageTriggerEnumMeterValues                       TriggerMessageRequestMessageTriggerEnum = "MeterValues"
	TriggerMessageRequestMessageTriggerEnumSignChargingStationCertificate    TriggerMessageRequestMessageTriggerEnum = "SignChargingStationCertificate"
	TriggerMessageRequestMessageTriggerEnumSignV2GCertificate                TriggerMessageRequestMessageTriggerEnum = "SignV2GCertificate"
	TriggerMessageRequestMessageTriggerEnumSignV2G20Certificate              TriggerMessageRequestMessageTriggerEnum = "SignV2G20Certificate"
	TriggerMessageRequestMessageTriggerEnumStatusNotification                TriggerMessageRequestMessageTriggerEnum = "StatusNotification"
	TriggerMessageRequestMessageTriggerEnumTransactionEvent                  TriggerMessageRequestMessageTriggerEnum = "TransactionEvent"
	TriggerMessageRequestMessageTriggerEnumSignCombinedCertificate           TriggerMessageRequestMessageTriggerEnum = "SignCombinedCertificate"
	TriggerMessageRequestMessageTriggerEnumPublishFirmwareStatusNotification TriggerMessageRequestMessageTriggerEnum = "PublishFirmwareStatusNotification"
	TriggerMessageRequestMessageTriggerEnumCustomTrigger                     TriggerMessageRequestMessageTriggerEnum = "CustomTrigger"
)

type TriggerMessageRequestEVSE struct {
	ID          int                              `json:"id"`
	ConnectorID *int                             `json:"connectorId,omitempty"`
	CustomData  *TriggerMessageRequestCustomData `json:"customData,omitempty"`
}

type TriggerMessageRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (TriggerMessageRequest) ActionName() string { return "TriggerMessage" }

func (TriggerMessageRequest) Version() protocol.Version { return protocol.OCPP21 }

func (TriggerMessageRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (TriggerMessageRequest) SchemaName() string { return "TriggerMessageRequest.json" }

func (message TriggerMessageRequest) Validate() error {
	return validation.Validate("TriggerMessageRequest.json", schemaTriggerMessageRequest, message)
}

func (TriggerMessageRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("TriggerMessageRequest.json", schemaTriggerMessageRequest, data)
}
