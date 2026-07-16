// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = NotifyDERStartStopRequest{}

var schemaNotifyDERStartStopRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"controlId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "started": &validation.Schema{Type: "boolean", AllowAdditional: true}, "timestamp": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "supersededIds": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, MinItems: 1, HasMinItems: true, MaxItems: 24, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"controlId", "started", "timestamp"}}

type NotifyDERStartStopRequest struct {
	ControlID     string                               `json:"controlId"`
	Started       bool                                 `json:"started"`
	Timestamp     string                               `json:"timestamp"`
	SupersededIds []string                             `json:"supersededIds,omitempty"`
	CustomData    *NotifyDERStartStopRequestCustomData `json:"customData,omitempty"`
}

type NotifyDERStartStopRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value NotifyDERStartStopRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *NotifyDERStartStopRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (NotifyDERStartStopRequest) ActionName() string { return "NotifyDERStartStop" }

func (NotifyDERStartStopRequest) Version() protocol.Version { return protocol.OCPP21 }

func (NotifyDERStartStopRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (NotifyDERStartStopRequest) SchemaName() string { return "NotifyDERStartStopRequest.json" }

func (message NotifyDERStartStopRequest) Validate() error {
	return validation.Validate("NotifyDERStartStopRequest.json", schemaNotifyDERStartStopRequest, message)
}

func (NotifyDERStartStopRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("NotifyDERStartStopRequest.json", schemaNotifyDERStartStopRequest, data)
}
