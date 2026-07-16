// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = SetMonitoringBaseConfirmation{}

var schemaSetMonitoringBaseConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "NotSupported", "EmptyResultSet"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}}, Required: []string{"status"}}

type SetMonitoringBaseConfirmation struct {
	CustomData *SetMonitoringBaseConfirmationCustomData                  `json:"customData,omitempty"`
	Status     SetMonitoringBaseConfirmationGenericDeviceModelStatusEnum `json:"status"`
	StatusInfo *SetMonitoringBaseConfirmationStatusInfo                  `json:"statusInfo,omitempty"`
}

type SetMonitoringBaseConfirmationStatusInfo struct {
	CustomData     *SetMonitoringBaseConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                                   `json:"reasonCode"`
	AdditionalInfo *string                                  `json:"additionalInfo,omitempty"`
}

type SetMonitoringBaseConfirmationGenericDeviceModelStatusEnum string

const (
	SetMonitoringBaseConfirmationGenericDeviceModelStatusEnumAccepted       SetMonitoringBaseConfirmationGenericDeviceModelStatusEnum = "Accepted"
	SetMonitoringBaseConfirmationGenericDeviceModelStatusEnumRejected       SetMonitoringBaseConfirmationGenericDeviceModelStatusEnum = "Rejected"
	SetMonitoringBaseConfirmationGenericDeviceModelStatusEnumNotSupported   SetMonitoringBaseConfirmationGenericDeviceModelStatusEnum = "NotSupported"
	SetMonitoringBaseConfirmationGenericDeviceModelStatusEnumEmptyResultSet SetMonitoringBaseConfirmationGenericDeviceModelStatusEnum = "EmptyResultSet"
)

type SetMonitoringBaseConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value SetMonitoringBaseConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *SetMonitoringBaseConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (SetMonitoringBaseConfirmation) ActionName() string { return "SetMonitoringBase" }

func (SetMonitoringBaseConfirmation) Version() protocol.Version { return protocol.OCPP201 }

func (SetMonitoringBaseConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (SetMonitoringBaseConfirmation) SchemaName() string { return "SetMonitoringBaseResponse.json" }

func (message SetMonitoringBaseConfirmation) Validate() error {
	return validation.Validate("SetMonitoringBaseResponse.json", schemaSetMonitoringBaseConfirmation, message)
}

func (SetMonitoringBaseConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("SetMonitoringBaseResponse.json", schemaSetMonitoringBaseConfirmation, data)
}
