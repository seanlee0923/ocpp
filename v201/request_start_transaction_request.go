// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = RequestStartTransactionRequest{}

var schemaRequestStartTransactionRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "evseId": &validation.Schema{Type: "integer", AllowAdditional: true}, "groupIdToken": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "additionalInfo": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "additionalIdToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}}, Required: []string{"additionalIdToken", "type"}}, MinItems: 1, HasMinItems: true}, "idToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "type": &validation.Schema{Type: "string", Enum: []string{"Central", "eMAID", "ISO14443", "ISO15693", "KeyCode", "Local", "MacAddress", "NoAuthorization"}}}, Required: []string{"idToken", "type"}}, "idToken": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "additionalInfo": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "additionalIdToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}}, Required: []string{"additionalIdToken", "type"}}, MinItems: 1, HasMinItems: true}, "idToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "type": &validation.Schema{Type: "string", Enum: []string{"Central", "eMAID", "ISO14443", "ISO15693", "KeyCode", "Local", "MacAddress", "NoAuthorization"}}}, Required: []string{"idToken", "type"}}, "remoteStartId": &validation.Schema{Type: "integer", AllowAdditional: true}, "chargingProfile": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "id": &validation.Schema{Type: "integer", AllowAdditional: true}, "stackLevel": &validation.Schema{Type: "integer", AllowAdditional: true}, "chargingProfilePurpose": &validation.Schema{Type: "string", Enum: []string{"ChargingStationExternalConstraints", "ChargingStationMaxProfile", "TxDefaultProfile", "TxProfile"}}, "chargingProfileKind": &validation.Schema{Type: "string", Enum: []string{"Absolute", "Recurring", "Relative"}}, "recurrencyKind": &validation.Schema{Type: "string", Enum: []string{"Daily", "Weekly"}}, "validFrom": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "validTo": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "chargingSchedule": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "id": &validation.Schema{Type: "integer", AllowAdditional: true}, "startSchedule": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "duration": &validation.Schema{Type: "integer", AllowAdditional: true}, "chargingRateUnit": &validation.Schema{Type: "string", Enum: []string{"W", "A"}}, "chargingSchedulePeriod": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "startPeriod": &validation.Schema{Type: "integer", AllowAdditional: true}, "limit": &validation.Schema{Type: "number", AllowAdditional: true}, "numberPhases": &validation.Schema{Type: "integer", AllowAdditional: true}, "phaseToUse": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"startPeriod", "limit"}}, MinItems: 1, HasMinItems: true, MaxItems: 1024, HasMaxItems: true}, "minChargingRate": &validation.Schema{Type: "number", AllowAdditional: true}, "salesTariff": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "id": &validation.Schema{Type: "integer", AllowAdditional: true}, "salesTariffDescription": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 32, HasMaxLength: true}, "numEPriceLevels": &validation.Schema{Type: "integer", AllowAdditional: true}, "salesTariffEntry": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "relativeTimeInterval": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "start": &validation.Schema{Type: "integer", AllowAdditional: true}, "duration": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"start"}}, "ePriceLevel": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "consumptionCost": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "startValue": &validation.Schema{Type: "number", AllowAdditional: true}, "cost": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "costKind": &validation.Schema{Type: "string", Enum: []string{"CarbonDioxideEmission", "RelativePricePercentage", "RenewableGenerationPercentage"}}, "amount": &validation.Schema{Type: "integer", AllowAdditional: true}, "amountMultiplier": &validation.Schema{Type: "integer", AllowAdditional: true}}, Required: []string{"costKind", "amount"}}, MinItems: 1, HasMinItems: true, MaxItems: 3, HasMaxItems: true}}, Required: []string{"startValue", "cost"}}, MinItems: 1, HasMinItems: true, MaxItems: 3, HasMaxItems: true}}, Required: []string{"relativeTimeInterval"}}, MinItems: 1, HasMinItems: true, MaxItems: 1024, HasMaxItems: true}}, Required: []string{"id", "salesTariffEntry"}}}, Required: []string{"id", "chargingRateUnit", "chargingSchedulePeriod"}}, MinItems: 1, HasMinItems: true, MaxItems: 3, HasMaxItems: true}, "transactionId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}}, Required: []string{"id", "stackLevel", "chargingProfilePurpose", "chargingProfileKind", "chargingSchedule"}}}, Required: []string{"remoteStartId", "idToken"}}

type RequestStartTransactionRequest struct {
	CustomData      *RequestStartTransactionRequestCustomData      `json:"customData,omitempty"`
	EVSEID          *int                                           `json:"evseId,omitempty"`
	GroupIDToken    *RequestStartTransactionRequestIDToken         `json:"groupIdToken,omitempty"`
	IDToken         RequestStartTransactionRequestIDToken          `json:"idToken"`
	RemoteStartID   int                                            `json:"remoteStartId"`
	ChargingProfile *RequestStartTransactionRequestChargingProfile `json:"chargingProfile,omitempty"`
}

type RequestStartTransactionRequestChargingProfile struct {
	CustomData             *RequestStartTransactionRequestCustomData                `json:"customData,omitempty"`
	ID                     int                                                      `json:"id"`
	StackLevel             int                                                      `json:"stackLevel"`
	ChargingProfilePurpose RequestStartTransactionRequestChargingProfilePurposeEnum `json:"chargingProfilePurpose"`
	ChargingProfileKind    RequestStartTransactionRequestChargingProfileKindEnum    `json:"chargingProfileKind"`
	RecurrencyKind         *RequestStartTransactionRequestRecurrencyKindEnum        `json:"recurrencyKind,omitempty"`
	ValidFrom              *string                                                  `json:"validFrom,omitempty"`
	ValidTo                *string                                                  `json:"validTo,omitempty"`
	ChargingSchedule       []RequestStartTransactionRequestChargingSchedule         `json:"chargingSchedule"`
	TransactionID          *string                                                  `json:"transactionId,omitempty"`
}

type RequestStartTransactionRequestChargingSchedule struct {
	CustomData             *RequestStartTransactionRequestCustomData              `json:"customData,omitempty"`
	ID                     int                                                    `json:"id"`
	StartSchedule          *string                                                `json:"startSchedule,omitempty"`
	Duration               *int                                                   `json:"duration,omitempty"`
	ChargingRateUnit       RequestStartTransactionRequestChargingRateUnitEnum     `json:"chargingRateUnit"`
	ChargingSchedulePeriod []RequestStartTransactionRequestChargingSchedulePeriod `json:"chargingSchedulePeriod"`
	MinChargingRate        *float64                                               `json:"minChargingRate,omitempty"`
	SalesTariff            *RequestStartTransactionRequestSalesTariff             `json:"salesTariff,omitempty"`
}

type RequestStartTransactionRequestSalesTariff struct {
	CustomData             *RequestStartTransactionRequestCustomData        `json:"customData,omitempty"`
	ID                     int                                              `json:"id"`
	SalesTariffDescription *string                                          `json:"salesTariffDescription,omitempty"`
	NumEPriceLevels        *int                                             `json:"numEPriceLevels,omitempty"`
	SalesTariffEntry       []RequestStartTransactionRequestSalesTariffEntry `json:"salesTariffEntry"`
}

type RequestStartTransactionRequestSalesTariffEntry struct {
	CustomData           *RequestStartTransactionRequestCustomData          `json:"customData,omitempty"`
	RelativeTimeInterval RequestStartTransactionRequestRelativeTimeInterval `json:"relativeTimeInterval"`
	EPriceLevel          *int                                               `json:"ePriceLevel,omitempty"`
	ConsumptionCost      []RequestStartTransactionRequestConsumptionCost    `json:"consumptionCost,omitempty"`
}

type RequestStartTransactionRequestConsumptionCost struct {
	CustomData *RequestStartTransactionRequestCustomData `json:"customData,omitempty"`
	StartValue float64                                   `json:"startValue"`
	Cost       []RequestStartTransactionRequestCost      `json:"cost"`
}

type RequestStartTransactionRequestCost struct {
	CustomData       *RequestStartTransactionRequestCustomData  `json:"customData,omitempty"`
	CostKind         RequestStartTransactionRequestCostKindEnum `json:"costKind"`
	Amount           int                                        `json:"amount"`
	AmountMultiplier *int                                       `json:"amountMultiplier,omitempty"`
}

type RequestStartTransactionRequestCostKindEnum string

const (
	RequestStartTransactionRequestCostKindEnumCarbonDioxideEmission         RequestStartTransactionRequestCostKindEnum = "CarbonDioxideEmission"
	RequestStartTransactionRequestCostKindEnumRelativePricePercentage       RequestStartTransactionRequestCostKindEnum = "RelativePricePercentage"
	RequestStartTransactionRequestCostKindEnumRenewableGenerationPercentage RequestStartTransactionRequestCostKindEnum = "RenewableGenerationPercentage"
)

type RequestStartTransactionRequestRelativeTimeInterval struct {
	CustomData *RequestStartTransactionRequestCustomData `json:"customData,omitempty"`
	Start      int                                       `json:"start"`
	Duration   *int                                      `json:"duration,omitempty"`
}

type RequestStartTransactionRequestChargingSchedulePeriod struct {
	CustomData   *RequestStartTransactionRequestCustomData `json:"customData,omitempty"`
	StartPeriod  int                                       `json:"startPeriod"`
	Limit        float64                                   `json:"limit"`
	NumberPhases *int                                      `json:"numberPhases,omitempty"`
	PhaseToUse   *int                                      `json:"phaseToUse,omitempty"`
}

type RequestStartTransactionRequestChargingRateUnitEnum string

const (
	RequestStartTransactionRequestChargingRateUnitEnumW RequestStartTransactionRequestChargingRateUnitEnum = "W"
	RequestStartTransactionRequestChargingRateUnitEnumA RequestStartTransactionRequestChargingRateUnitEnum = "A"
)

type RequestStartTransactionRequestRecurrencyKindEnum string

const (
	RequestStartTransactionRequestRecurrencyKindEnumDaily  RequestStartTransactionRequestRecurrencyKindEnum = "Daily"
	RequestStartTransactionRequestRecurrencyKindEnumWeekly RequestStartTransactionRequestRecurrencyKindEnum = "Weekly"
)

type RequestStartTransactionRequestChargingProfileKindEnum string

const (
	RequestStartTransactionRequestChargingProfileKindEnumAbsolute  RequestStartTransactionRequestChargingProfileKindEnum = "Absolute"
	RequestStartTransactionRequestChargingProfileKindEnumRecurring RequestStartTransactionRequestChargingProfileKindEnum = "Recurring"
	RequestStartTransactionRequestChargingProfileKindEnumRelative  RequestStartTransactionRequestChargingProfileKindEnum = "Relative"
)

type RequestStartTransactionRequestChargingProfilePurposeEnum string

const (
	RequestStartTransactionRequestChargingProfilePurposeEnumChargingStationExternalConstraints RequestStartTransactionRequestChargingProfilePurposeEnum = "ChargingStationExternalConstraints"
	RequestStartTransactionRequestChargingProfilePurposeEnumChargingStationMaxProfile          RequestStartTransactionRequestChargingProfilePurposeEnum = "ChargingStationMaxProfile"
	RequestStartTransactionRequestChargingProfilePurposeEnumTxDefaultProfile                   RequestStartTransactionRequestChargingProfilePurposeEnum = "TxDefaultProfile"
	RequestStartTransactionRequestChargingProfilePurposeEnumTxProfile                          RequestStartTransactionRequestChargingProfilePurposeEnum = "TxProfile"
)

type RequestStartTransactionRequestIDToken struct {
	CustomData     *RequestStartTransactionRequestCustomData      `json:"customData,omitempty"`
	AdditionalInfo []RequestStartTransactionRequestAdditionalInfo `json:"additionalInfo,omitempty"`
	IDToken        string                                         `json:"idToken"`
	Type           RequestStartTransactionRequestIDTokenEnum      `json:"type"`
}

type RequestStartTransactionRequestIDTokenEnum string

const (
	RequestStartTransactionRequestIDTokenEnumCentral         RequestStartTransactionRequestIDTokenEnum = "Central"
	RequestStartTransactionRequestIDTokenEnumEMAID           RequestStartTransactionRequestIDTokenEnum = "eMAID"
	RequestStartTransactionRequestIDTokenEnumISO14443        RequestStartTransactionRequestIDTokenEnum = "ISO14443"
	RequestStartTransactionRequestIDTokenEnumISO15693        RequestStartTransactionRequestIDTokenEnum = "ISO15693"
	RequestStartTransactionRequestIDTokenEnumKeyCode         RequestStartTransactionRequestIDTokenEnum = "KeyCode"
	RequestStartTransactionRequestIDTokenEnumLocal           RequestStartTransactionRequestIDTokenEnum = "Local"
	RequestStartTransactionRequestIDTokenEnumMacAddress      RequestStartTransactionRequestIDTokenEnum = "MacAddress"
	RequestStartTransactionRequestIDTokenEnumNoAuthorization RequestStartTransactionRequestIDTokenEnum = "NoAuthorization"
)

type RequestStartTransactionRequestAdditionalInfo struct {
	CustomData        *RequestStartTransactionRequestCustomData `json:"customData,omitempty"`
	AdditionalIDToken string                                    `json:"additionalIdToken"`
	Type              string                                    `json:"type"`
}

type RequestStartTransactionRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value RequestStartTransactionRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *RequestStartTransactionRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (RequestStartTransactionRequest) ActionName() string { return "RequestStartTransaction" }

func (RequestStartTransactionRequest) Version() protocol.Version { return protocol.OCPP201 }

func (RequestStartTransactionRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (RequestStartTransactionRequest) SchemaName() string {
	return "RequestStartTransactionRequest.json"
}

func (message RequestStartTransactionRequest) Validate() error {
	return validation.Validate("RequestStartTransactionRequest.json", schemaRequestStartTransactionRequest, message)
}

func (RequestStartTransactionRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("RequestStartTransactionRequest.json", schemaRequestStartTransactionRequest, data)
}
