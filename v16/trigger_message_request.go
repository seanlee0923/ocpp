// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = TriggerMessageRequest{}

var schemaTriggerMessageRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"requestedMessage": &validation.Schema{Type: "string", Enum: []string{"BootNotification", "DiagnosticsStatusNotification", "FirmwareStatusNotification", "Heartbeat", "MeterValues", "StatusNotification"}}, "connectorId": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"requestedMessage"}}

type TriggerMessageRequest struct {
	RequestedMessage TriggerMessageRequestRequestedMessage `json:"requestedMessage"`
	ConnectorID      *int                                  `json:"connectorId,omitempty"`
}

type TriggerMessageRequestRequestedMessage string

const (
	TriggerMessageRequestRequestedMessageBootNotification              TriggerMessageRequestRequestedMessage = "BootNotification"
	TriggerMessageRequestRequestedMessageDiagnosticsStatusNotification TriggerMessageRequestRequestedMessage = "DiagnosticsStatusNotification"
	TriggerMessageRequestRequestedMessageFirmwareStatusNotification    TriggerMessageRequestRequestedMessage = "FirmwareStatusNotification"
	TriggerMessageRequestRequestedMessageHeartbeat                     TriggerMessageRequestRequestedMessage = "Heartbeat"
	TriggerMessageRequestRequestedMessageMeterValues                   TriggerMessageRequestRequestedMessage = "MeterValues"
	TriggerMessageRequestRequestedMessageStatusNotification            TriggerMessageRequestRequestedMessage = "StatusNotification"
)

func (TriggerMessageRequest) ActionName() string { return "TriggerMessage" }

func (TriggerMessageRequest) Version() protocol.Version { return protocol.OCPP16 }

func (TriggerMessageRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (TriggerMessageRequest) SchemaName() string { return "TriggerMessage.json" }

func (message TriggerMessageRequest) Validate() error {
	return validation.Validate("TriggerMessage.json", schemaTriggerMessageRequest, message)
}

func (TriggerMessageRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("TriggerMessage.json", schemaTriggerMessageRequest, data)
}
