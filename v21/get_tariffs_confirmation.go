// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = GetTariffsConfirmation{}

var schemaGetTariffsConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected", "NoTariff"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "tariffAssignments": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"tariffId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 60, HasMaxLength: true}, "tariffKind": &validation.Schema{Type: "string", Enum: []string{"DefaultTariff", "DriverTariff"}}, "validFrom": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "evseIds": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, MinItems: 1, HasMinItems: true}, "idTokens": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}, MinItems: 1, HasMinItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"tariffId", "tariffKind"}}, MinItems: 1, HasMinItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type GetTariffsConfirmation struct {
	Status            GetTariffsConfirmationTariffGetStatusEnum `json:"status"`
	StatusInfo        *GetTariffsConfirmationStatusInfo         `json:"statusInfo,omitempty"`
	TariffAssignments []GetTariffsConfirmationTariffAssignment  `json:"tariffAssignments,omitempty"`
	CustomData        *GetTariffsConfirmationCustomData         `json:"customData,omitempty"`
}

type GetTariffsConfirmationTariffAssignment struct {
	TariffID   string                               `json:"tariffId"`
	TariffKind GetTariffsConfirmationTariffKindEnum `json:"tariffKind"`
	ValidFrom  *string                              `json:"validFrom,omitempty"`
	EVSEIds    []int                                `json:"evseIds,omitempty"`
	IDTokens   []string                             `json:"idTokens,omitempty"`
	CustomData *GetTariffsConfirmationCustomData    `json:"customData,omitempty"`
}

type GetTariffsConfirmationTariffKindEnum string

const (
	GetTariffsConfirmationTariffKindEnumDefaultTariff GetTariffsConfirmationTariffKindEnum = "DefaultTariff"
	GetTariffsConfirmationTariffKindEnumDriverTariff  GetTariffsConfirmationTariffKindEnum = "DriverTariff"
)

type GetTariffsConfirmationStatusInfo struct {
	ReasonCode     string                            `json:"reasonCode"`
	AdditionalInfo *string                           `json:"additionalInfo,omitempty"`
	CustomData     *GetTariffsConfirmationCustomData `json:"customData,omitempty"`
}

type GetTariffsConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value GetTariffsConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *GetTariffsConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

type GetTariffsConfirmationTariffGetStatusEnum string

const (
	GetTariffsConfirmationTariffGetStatusEnumAccepted GetTariffsConfirmationTariffGetStatusEnum = "Accepted"
	GetTariffsConfirmationTariffGetStatusEnumRejected GetTariffsConfirmationTariffGetStatusEnum = "Rejected"
	GetTariffsConfirmationTariffGetStatusEnumNoTariff GetTariffsConfirmationTariffGetStatusEnum = "NoTariff"
)

func (GetTariffsConfirmation) ActionName() string { return "GetTariffs" }

func (GetTariffsConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (GetTariffsConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (GetTariffsConfirmation) SchemaName() string { return "GetTariffsResponse.json" }

func (message GetTariffsConfirmation) Validate() error {
	return validation.Validate("GetTariffsResponse.json", schemaGetTariffsConfirmation, message)
}

func (GetTariffsConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetTariffsResponse.json", schemaGetTariffsConfirmation, data)
}
