// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = TriggerMessageRequest{}

var schemaTriggerMessageRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "evse": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "id": &validation.Schema{Type: "integer", AllowAdditional: true}, "connectorId": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"id"}}, "requestedMessage": &validation.Schema{Type: "string", Enum: []string{"BootNotification", "LogStatusNotification", "FirmwareStatusNotification", "Heartbeat", "MeterValues", "SignChargingStationCertificate", "SignV2GCertificate", "StatusNotification", "TransactionEvent", "SignCombinedCertificate", "PublishFirmwareStatusNotification"}}}, Required: []string{"requestedMessage"}}

type TriggerMessageRequest struct {
	CustomData       *TriggerMessageRequestCustomData        `json:"customData,omitempty"`
	EVSE             *TriggerMessageRequestEVSE              `json:"evse,omitempty"`
	RequestedMessage TriggerMessageRequestMessageTriggerEnum `json:"requestedMessage"`
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
	TriggerMessageRequestMessageTriggerEnumStatusNotification                TriggerMessageRequestMessageTriggerEnum = "StatusNotification"
	TriggerMessageRequestMessageTriggerEnumTransactionEvent                  TriggerMessageRequestMessageTriggerEnum = "TransactionEvent"
	TriggerMessageRequestMessageTriggerEnumSignCombinedCertificate           TriggerMessageRequestMessageTriggerEnum = "SignCombinedCertificate"
	TriggerMessageRequestMessageTriggerEnumPublishFirmwareStatusNotification TriggerMessageRequestMessageTriggerEnum = "PublishFirmwareStatusNotification"
)

type TriggerMessageRequestEVSE struct {
	CustomData  *TriggerMessageRequestCustomData `json:"customData,omitempty"`
	ID          int                              `json:"id"`
	ConnectorID *int                             `json:"connectorId,omitempty"`
}

type TriggerMessageRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value TriggerMessageRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *TriggerMessageRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (TriggerMessageRequest) ActionName() string { return "TriggerMessage" }

func (TriggerMessageRequest) Version() protocol.Version { return protocol.OCPP201 }

func (TriggerMessageRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (TriggerMessageRequest) SchemaName() string { return "TriggerMessageRequest.json" }

func (message TriggerMessageRequest) Validate() error {
	return validation.Validate("TriggerMessageRequest.json", schemaTriggerMessageRequest, message)
}

func (TriggerMessageRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("TriggerMessageRequest.json", schemaTriggerMessageRequest, data)
}
