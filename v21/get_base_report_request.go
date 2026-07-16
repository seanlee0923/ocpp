// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = GetBaseReportRequest{}

var schemaGetBaseReportRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"requestId": &validation.Schema{Type: "integer", AllowAdditional: true}, "reportBase": &validation.Schema{Type: "string", Enum: []string{"ConfigurationInventory", "FullInventory", "SummaryInventory"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"requestId", "reportBase"}}

type GetBaseReportRequest struct {
	RequestID  int                                `json:"requestId"`
	ReportBase GetBaseReportRequestReportBaseEnum `json:"reportBase"`
	CustomData *GetBaseReportRequestCustomData    `json:"customData,omitempty"`
}

type GetBaseReportRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value GetBaseReportRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *GetBaseReportRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

type GetBaseReportRequestReportBaseEnum string

const (
	GetBaseReportRequestReportBaseEnumConfigurationInventory GetBaseReportRequestReportBaseEnum = "ConfigurationInventory"
	GetBaseReportRequestReportBaseEnumFullInventory          GetBaseReportRequestReportBaseEnum = "FullInventory"
	GetBaseReportRequestReportBaseEnumSummaryInventory       GetBaseReportRequestReportBaseEnum = "SummaryInventory"
)

func (GetBaseReportRequest) ActionName() string { return "GetBaseReport" }

func (GetBaseReportRequest) Version() protocol.Version { return protocol.OCPP21 }

func (GetBaseReportRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (GetBaseReportRequest) SchemaName() string { return "GetBaseReportRequest.json" }

func (message GetBaseReportRequest) Validate() error {
	return validation.Validate("GetBaseReportRequest.json", schemaGetBaseReportRequest, message)
}

func (GetBaseReportRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetBaseReportRequest.json", schemaGetBaseReportRequest, data)
}
