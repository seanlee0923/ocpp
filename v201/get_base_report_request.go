// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = GetBaseReportRequest{}

var schemaGetBaseReportRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "requestId": &validation.Schema{Type: "integer", AllowAdditional: true}, "reportBase": &validation.Schema{Type: "string", Enum: []string{"ConfigurationInventory", "FullInventory", "SummaryInventory"}}}, Required: []string{"requestId", "reportBase"}}

type GetBaseReportRequest struct {
	CustomData *GetBaseReportRequestCustomData    `json:"customData,omitempty"`
	RequestID  int                                `json:"requestId"`
	ReportBase GetBaseReportRequestReportBaseEnum `json:"reportBase"`
}

type GetBaseReportRequestReportBaseEnum string

const (
	GetBaseReportRequestReportBaseEnumConfigurationInventory GetBaseReportRequestReportBaseEnum = "ConfigurationInventory"
	GetBaseReportRequestReportBaseEnumFullInventory          GetBaseReportRequestReportBaseEnum = "FullInventory"
	GetBaseReportRequestReportBaseEnumSummaryInventory       GetBaseReportRequestReportBaseEnum = "SummaryInventory"
)

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

func (GetBaseReportRequest) ActionName() string { return "GetBaseReport" }

func (GetBaseReportRequest) Version() protocol.Version { return protocol.OCPP201 }

func (GetBaseReportRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (GetBaseReportRequest) SchemaName() string { return "GetBaseReportRequest.json" }

func (message GetBaseReportRequest) Validate() error {
	return validation.Validate("GetBaseReportRequest.json", schemaGetBaseReportRequest, message)
}

func (GetBaseReportRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetBaseReportRequest.json", schemaGetBaseReportRequest, data)
}
