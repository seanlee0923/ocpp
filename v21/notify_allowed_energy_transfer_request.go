// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = NotifyAllowedEnergyTransferRequest{}

var schemaNotifyAllowedEnergyTransferRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"transactionId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "allowedEnergyTransfer": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", Enum: []string{"AC_single_phase", "AC_two_phase", "AC_three_phase", "DC", "AC_BPT", "AC_BPT_DER", "AC_DER", "DC_BPT", "DC_ACDP", "DC_ACDP_BPT", "WPT"}}, MinItems: 1, HasMinItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"transactionId", "allowedEnergyTransfer"}}

type NotifyAllowedEnergyTransferRequest struct {
	TransactionID         string                                                     `json:"transactionId"`
	AllowedEnergyTransfer []NotifyAllowedEnergyTransferRequestEnergyTransferModeEnum `json:"allowedEnergyTransfer"`
	CustomData            *NotifyAllowedEnergyTransferRequestCustomData              `json:"customData,omitempty"`
}

type NotifyAllowedEnergyTransferRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value NotifyAllowedEnergyTransferRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *NotifyAllowedEnergyTransferRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

type NotifyAllowedEnergyTransferRequestEnergyTransferModeEnum string

const (
	NotifyAllowedEnergyTransferRequestEnergyTransferModeEnumACSinglePhase NotifyAllowedEnergyTransferRequestEnergyTransferModeEnum = "AC_single_phase"
	NotifyAllowedEnergyTransferRequestEnergyTransferModeEnumACTwoPhase    NotifyAllowedEnergyTransferRequestEnergyTransferModeEnum = "AC_two_phase"
	NotifyAllowedEnergyTransferRequestEnergyTransferModeEnumACThreePhase  NotifyAllowedEnergyTransferRequestEnergyTransferModeEnum = "AC_three_phase"
	NotifyAllowedEnergyTransferRequestEnergyTransferModeEnumDC            NotifyAllowedEnergyTransferRequestEnergyTransferModeEnum = "DC"
	NotifyAllowedEnergyTransferRequestEnergyTransferModeEnumACBPT         NotifyAllowedEnergyTransferRequestEnergyTransferModeEnum = "AC_BPT"
	NotifyAllowedEnergyTransferRequestEnergyTransferModeEnumACBPTDER      NotifyAllowedEnergyTransferRequestEnergyTransferModeEnum = "AC_BPT_DER"
	NotifyAllowedEnergyTransferRequestEnergyTransferModeEnumACDER         NotifyAllowedEnergyTransferRequestEnergyTransferModeEnum = "AC_DER"
	NotifyAllowedEnergyTransferRequestEnergyTransferModeEnumDCBPT         NotifyAllowedEnergyTransferRequestEnergyTransferModeEnum = "DC_BPT"
	NotifyAllowedEnergyTransferRequestEnergyTransferModeEnumDCACDP        NotifyAllowedEnergyTransferRequestEnergyTransferModeEnum = "DC_ACDP"
	NotifyAllowedEnergyTransferRequestEnergyTransferModeEnumDCACDPBPT     NotifyAllowedEnergyTransferRequestEnergyTransferModeEnum = "DC_ACDP_BPT"
	NotifyAllowedEnergyTransferRequestEnergyTransferModeEnumWPT           NotifyAllowedEnergyTransferRequestEnergyTransferModeEnum = "WPT"
)

func (NotifyAllowedEnergyTransferRequest) ActionName() string { return "NotifyAllowedEnergyTransfer" }

func (NotifyAllowedEnergyTransferRequest) Version() protocol.Version { return protocol.OCPP21 }

func (NotifyAllowedEnergyTransferRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (NotifyAllowedEnergyTransferRequest) SchemaName() string {
	return "NotifyAllowedEnergyTransferRequest.json"
}

func (message NotifyAllowedEnergyTransferRequest) Validate() error {
	return validation.Validate("NotifyAllowedEnergyTransferRequest.json", schemaNotifyAllowedEnergyTransferRequest, message)
}

func (NotifyAllowedEnergyTransferRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("NotifyAllowedEnergyTransferRequest.json", schemaNotifyAllowedEnergyTransferRequest, data)
}
