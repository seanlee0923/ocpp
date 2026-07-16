// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = StatusNotificationRequest{}

var schemaStatusNotificationRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "timestamp": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "connectorStatus": &validation.Schema{Type: "string", Enum: []string{"Available", "Occupied", "Reserved", "Unavailable", "Faulted"}}, "evseId": &validation.Schema{Type: "integer", AllowAdditional: true}, "connectorId": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"timestamp", "connectorStatus", "evseId", "connectorId"}}

type StatusNotificationRequest struct {
	CustomData      *StatusNotificationRequestCustomData         `json:"customData,omitempty"`
	Timestamp       string                                       `json:"timestamp"`
	ConnectorStatus StatusNotificationRequestConnectorStatusEnum `json:"connectorStatus"`
	EVSEID          int                                          `json:"evseId"`
	ConnectorID     int                                          `json:"connectorId"`
}

type StatusNotificationRequestConnectorStatusEnum string

const (
	StatusNotificationRequestConnectorStatusEnumAvailable   StatusNotificationRequestConnectorStatusEnum = "Available"
	StatusNotificationRequestConnectorStatusEnumOccupied    StatusNotificationRequestConnectorStatusEnum = "Occupied"
	StatusNotificationRequestConnectorStatusEnumReserved    StatusNotificationRequestConnectorStatusEnum = "Reserved"
	StatusNotificationRequestConnectorStatusEnumUnavailable StatusNotificationRequestConnectorStatusEnum = "Unavailable"
	StatusNotificationRequestConnectorStatusEnumFaulted     StatusNotificationRequestConnectorStatusEnum = "Faulted"
)

type StatusNotificationRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value StatusNotificationRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *StatusNotificationRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (StatusNotificationRequest) ActionName() string { return "StatusNotification" }

func (StatusNotificationRequest) Version() protocol.Version { return protocol.OCPP201 }

func (StatusNotificationRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (StatusNotificationRequest) SchemaName() string { return "StatusNotificationRequest.json" }

func (message StatusNotificationRequest) Validate() error {
	return validation.Validate("StatusNotificationRequest.json", schemaStatusNotificationRequest, message)
}

func (StatusNotificationRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("StatusNotificationRequest.json", schemaStatusNotificationRequest, data)
}
