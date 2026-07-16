// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = NotifyEVChargingNeedsRequest{}

var schemaNotifyEVChargingNeedsRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "maxScheduleTuples": &validation.Schema{Type: "integer", AllowAdditional: true}, "chargingNeeds": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "acChargingParameters": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "energyAmount": &validation.Schema{Type: "integer", AllowAdditional: true}, "evMinCurrent": &validation.Schema{Type: "integer", AllowAdditional: true}, "evMaxCurrent": &validation.Schema{Type: "integer", AllowAdditional: true}, "evMaxVoltage": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"energyAmount", "evMinCurrent", "evMaxCurrent", "evMaxVoltage"}}, "dcChargingParameters": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "evMaxCurrent": &validation.Schema{Type: "integer", AllowAdditional: true}, "evMaxVoltage": &validation.Schema{Type: "integer", AllowAdditional: true}, "energyAmount": &validation.Schema{Type: "integer", AllowAdditional: true}, "evMaxPower": &validation.Schema{Type: "integer", AllowAdditional: true}, "stateOfCharge": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true, Maximum: 100, HasMaximum: true}, "evEnergyCapacity": &validation.Schema{Type: "integer", AllowAdditional: true}, "fullSoC": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true, Maximum: 100, HasMaximum: true}, "bulkSoC": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true, Maximum: 100, HasMaximum: true}}, Required: []string{"evMaxCurrent", "evMaxVoltage"}}, "requestedEnergyTransfer": &validation.Schema{Type: "string", Enum: []string{"DC", "AC_single_phase", "AC_two_phase", "AC_three_phase"}}, "departureTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}}, Required: []string{"requestedEnergyTransfer"}}, "evseId": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"evseId", "chargingNeeds"}}

type NotifyEVChargingNeedsRequest struct {
	CustomData        *NotifyEVChargingNeedsRequestCustomData   `json:"customData,omitempty"`
	MaxScheduleTuples *int                                      `json:"maxScheduleTuples,omitempty"`
	ChargingNeeds     NotifyEVChargingNeedsRequestChargingNeeds `json:"chargingNeeds"`
	EVSEID            int                                       `json:"evseId"`
}

type NotifyEVChargingNeedsRequestChargingNeeds struct {
	CustomData              *NotifyEVChargingNeedsRequestCustomData            `json:"customData,omitempty"`
	AcChargingParameters    *NotifyEVChargingNeedsRequestACChargingParameters  `json:"acChargingParameters,omitempty"`
	DcChargingParameters    *NotifyEVChargingNeedsRequestDCChargingParameters  `json:"dcChargingParameters,omitempty"`
	RequestedEnergyTransfer NotifyEVChargingNeedsRequestEnergyTransferModeEnum `json:"requestedEnergyTransfer"`
	DepartureTime           *string                                            `json:"departureTime,omitempty"`
}

type NotifyEVChargingNeedsRequestEnergyTransferModeEnum string

const (
	NotifyEVChargingNeedsRequestEnergyTransferModeEnumDC            NotifyEVChargingNeedsRequestEnergyTransferModeEnum = "DC"
	NotifyEVChargingNeedsRequestEnergyTransferModeEnumACSinglePhase NotifyEVChargingNeedsRequestEnergyTransferModeEnum = "AC_single_phase"
	NotifyEVChargingNeedsRequestEnergyTransferModeEnumACTwoPhase    NotifyEVChargingNeedsRequestEnergyTransferModeEnum = "AC_two_phase"
	NotifyEVChargingNeedsRequestEnergyTransferModeEnumACThreePhase  NotifyEVChargingNeedsRequestEnergyTransferModeEnum = "AC_three_phase"
)

type NotifyEVChargingNeedsRequestDCChargingParameters struct {
	CustomData       *NotifyEVChargingNeedsRequestCustomData `json:"customData,omitempty"`
	EVMaxCurrent     int                                     `json:"evMaxCurrent"`
	EVMaxVoltage     int                                     `json:"evMaxVoltage"`
	EnergyAmount     *int                                    `json:"energyAmount,omitempty"`
	EVMaxPower       *int                                    `json:"evMaxPower,omitempty"`
	StateOfCharge    *int                                    `json:"stateOfCharge,omitempty"`
	EVEnergyCapacity *int                                    `json:"evEnergyCapacity,omitempty"`
	FullSoC          *int                                    `json:"fullSoC,omitempty"`
	BulkSoC          *int                                    `json:"bulkSoC,omitempty"`
}

type NotifyEVChargingNeedsRequestACChargingParameters struct {
	CustomData   *NotifyEVChargingNeedsRequestCustomData `json:"customData,omitempty"`
	EnergyAmount int                                     `json:"energyAmount"`
	EVMinCurrent int                                     `json:"evMinCurrent"`
	EVMaxCurrent int                                     `json:"evMaxCurrent"`
	EVMaxVoltage int                                     `json:"evMaxVoltage"`
}

type NotifyEVChargingNeedsRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value NotifyEVChargingNeedsRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *NotifyEVChargingNeedsRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (NotifyEVChargingNeedsRequest) ActionName() string { return "NotifyEVChargingNeeds" }

func (NotifyEVChargingNeedsRequest) Version() protocol.Version { return protocol.OCPP201 }

func (NotifyEVChargingNeedsRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (NotifyEVChargingNeedsRequest) SchemaName() string { return "NotifyEVChargingNeedsRequest.json" }

func (message NotifyEVChargingNeedsRequest) Validate() error {
	return validation.Validate("NotifyEVChargingNeedsRequest.json", schemaNotifyEVChargingNeedsRequest, message)
}

func (NotifyEVChargingNeedsRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("NotifyEVChargingNeedsRequest.json", schemaNotifyEVChargingNeedsRequest, data)
}
