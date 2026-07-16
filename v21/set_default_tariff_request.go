// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = SetDefaultTariffRequest{}

var schemaSetDefaultTariffRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"evseId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "tariff": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"tariffId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 60, HasMaxLength: true}, "description": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"format": &validation.Schema{Type: "string", Enum: []string{"ASCII", "HTML", "URI", "UTF8", "QRCODE"}}, "language": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 8, HasMaxLength: true}, "content": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"format", "content"}}, MinItems: 1, HasMinItems: true, MaxItems: 10, HasMaxItems: true}, "currency": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 3, HasMaxLength: true}, "energy": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"prices": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"priceKwh": &validation.Schema{Type: "number", AllowAdditional: true}, "conditions": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"startTimeOfDay": &validation.Schema{Type: "string", AllowAdditional: true}, "endTimeOfDay": &validation.Schema{Type: "string", AllowAdditional: true}, "dayOfWeek": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", Enum: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}}, MinItems: 1, HasMinItems: true, MaxItems: 7, HasMaxItems: true}, "validFromDate": &validation.Schema{Type: "string", AllowAdditional: true}, "validToDate": &validation.Schema{Type: "string", AllowAdditional: true}, "evseKind": &validation.Schema{Type: "string", Enum: []string{"AC", "DC"}}, "minEnergy": &validation.Schema{Type: "number", AllowAdditional: true}, "maxEnergy": &validation.Schema{Type: "number", AllowAdditional: true}, "minCurrent": &validation.Schema{Type: "number", AllowAdditional: true}, "maxCurrent": &validation.Schema{Type: "number", AllowAdditional: true}, "minPower": &validation.Schema{Type: "number", AllowAdditional: true}, "maxPower": &validation.Schema{Type: "number", AllowAdditional: true}, "minTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "maxTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "minChargingTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "maxChargingTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "minIdleTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "maxIdleTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"priceKwh"}}, MinItems: 1, HasMinItems: true}, "taxRates": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "tax": &validation.Schema{Type: "number", AllowAdditional: true}, "stack": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"type", "tax"}}, MinItems: 1, HasMinItems: true, MaxItems: 5, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"prices"}}, "validFrom": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "chargingTime": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"prices": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"priceMinute": &validation.Schema{Type: "number", AllowAdditional: true}, "conditions": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"startTimeOfDay": &validation.Schema{Type: "string", AllowAdditional: true}, "endTimeOfDay": &validation.Schema{Type: "string", AllowAdditional: true}, "dayOfWeek": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", Enum: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}}, MinItems: 1, HasMinItems: true, MaxItems: 7, HasMaxItems: true}, "validFromDate": &validation.Schema{Type: "string", AllowAdditional: true}, "validToDate": &validation.Schema{Type: "string", AllowAdditional: true}, "evseKind": &validation.Schema{Type: "string", Enum: []string{"AC", "DC"}}, "minEnergy": &validation.Schema{Type: "number", AllowAdditional: true}, "maxEnergy": &validation.Schema{Type: "number", AllowAdditional: true}, "minCurrent": &validation.Schema{Type: "number", AllowAdditional: true}, "maxCurrent": &validation.Schema{Type: "number", AllowAdditional: true}, "minPower": &validation.Schema{Type: "number", AllowAdditional: true}, "maxPower": &validation.Schema{Type: "number", AllowAdditional: true}, "minTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "maxTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "minChargingTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "maxChargingTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "minIdleTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "maxIdleTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"priceMinute"}}, MinItems: 1, HasMinItems: true}, "taxRates": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "tax": &validation.Schema{Type: "number", AllowAdditional: true}, "stack": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"type", "tax"}}, MinItems: 1, HasMinItems: true, MaxItems: 5, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"prices"}}, "idleTime": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"prices": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"priceMinute": &validation.Schema{Type: "number", AllowAdditional: true}, "conditions": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"startTimeOfDay": &validation.Schema{Type: "string", AllowAdditional: true}, "endTimeOfDay": &validation.Schema{Type: "string", AllowAdditional: true}, "dayOfWeek": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", Enum: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}}, MinItems: 1, HasMinItems: true, MaxItems: 7, HasMaxItems: true}, "validFromDate": &validation.Schema{Type: "string", AllowAdditional: true}, "validToDate": &validation.Schema{Type: "string", AllowAdditional: true}, "evseKind": &validation.Schema{Type: "string", Enum: []string{"AC", "DC"}}, "minEnergy": &validation.Schema{Type: "number", AllowAdditional: true}, "maxEnergy": &validation.Schema{Type: "number", AllowAdditional: true}, "minCurrent": &validation.Schema{Type: "number", AllowAdditional: true}, "maxCurrent": &validation.Schema{Type: "number", AllowAdditional: true}, "minPower": &validation.Schema{Type: "number", AllowAdditional: true}, "maxPower": &validation.Schema{Type: "number", AllowAdditional: true}, "minTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "maxTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "minChargingTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "maxChargingTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "minIdleTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "maxIdleTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"priceMinute"}}, MinItems: 1, HasMinItems: true}, "taxRates": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "tax": &validation.Schema{Type: "number", AllowAdditional: true}, "stack": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"type", "tax"}}, MinItems: 1, HasMinItems: true, MaxItems: 5, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"prices"}}, "fixedFee": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"prices": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"conditions": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"startTimeOfDay": &validation.Schema{Type: "string", AllowAdditional: true}, "endTimeOfDay": &validation.Schema{Type: "string", AllowAdditional: true}, "dayOfWeek": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", Enum: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}}, MinItems: 1, HasMinItems: true, MaxItems: 7, HasMaxItems: true}, "validFromDate": &validation.Schema{Type: "string", AllowAdditional: true}, "validToDate": &validation.Schema{Type: "string", AllowAdditional: true}, "evseKind": &validation.Schema{Type: "string", Enum: []string{"AC", "DC"}}, "paymentBrand": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "paymentRecognition": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "priceFixed": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"priceFixed"}}, MinItems: 1, HasMinItems: true}, "taxRates": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "tax": &validation.Schema{Type: "number", AllowAdditional: true}, "stack": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"type", "tax"}}, MinItems: 1, HasMinItems: true, MaxItems: 5, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"prices"}}, "reservationTime": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"prices": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"priceMinute": &validation.Schema{Type: "number", AllowAdditional: true}, "conditions": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"startTimeOfDay": &validation.Schema{Type: "string", AllowAdditional: true}, "endTimeOfDay": &validation.Schema{Type: "string", AllowAdditional: true}, "dayOfWeek": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", Enum: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}}, MinItems: 1, HasMinItems: true, MaxItems: 7, HasMaxItems: true}, "validFromDate": &validation.Schema{Type: "string", AllowAdditional: true}, "validToDate": &validation.Schema{Type: "string", AllowAdditional: true}, "evseKind": &validation.Schema{Type: "string", Enum: []string{"AC", "DC"}}, "minEnergy": &validation.Schema{Type: "number", AllowAdditional: true}, "maxEnergy": &validation.Schema{Type: "number", AllowAdditional: true}, "minCurrent": &validation.Schema{Type: "number", AllowAdditional: true}, "maxCurrent": &validation.Schema{Type: "number", AllowAdditional: true}, "minPower": &validation.Schema{Type: "number", AllowAdditional: true}, "maxPower": &validation.Schema{Type: "number", AllowAdditional: true}, "minTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "maxTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "minChargingTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "maxChargingTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "minIdleTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "maxIdleTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"priceMinute"}}, MinItems: 1, HasMinItems: true}, "taxRates": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "tax": &validation.Schema{Type: "number", AllowAdditional: true}, "stack": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"type", "tax"}}, MinItems: 1, HasMinItems: true, MaxItems: 5, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"prices"}}, "reservationFixed": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"prices": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"conditions": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"startTimeOfDay": &validation.Schema{Type: "string", AllowAdditional: true}, "endTimeOfDay": &validation.Schema{Type: "string", AllowAdditional: true}, "dayOfWeek": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", Enum: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}}, MinItems: 1, HasMinItems: true, MaxItems: 7, HasMaxItems: true}, "validFromDate": &validation.Schema{Type: "string", AllowAdditional: true}, "validToDate": &validation.Schema{Type: "string", AllowAdditional: true}, "evseKind": &validation.Schema{Type: "string", Enum: []string{"AC", "DC"}}, "paymentBrand": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "paymentRecognition": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "priceFixed": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"priceFixed"}}, MinItems: 1, HasMinItems: true}, "taxRates": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "tax": &validation.Schema{Type: "number", AllowAdditional: true}, "stack": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"type", "tax"}}, MinItems: 1, HasMinItems: true, MaxItems: 5, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"prices"}}, "minCost": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exclTax": &validation.Schema{Type: "number", AllowAdditional: true}, "inclTax": &validation.Schema{Type: "number", AllowAdditional: true}, "taxRates": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "tax": &validation.Schema{Type: "number", AllowAdditional: true}, "stack": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"type", "tax"}}, MinItems: 1, HasMinItems: true, MaxItems: 5, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "maxCost": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exclTax": &validation.Schema{Type: "number", AllowAdditional: true}, "inclTax": &validation.Schema{Type: "number", AllowAdditional: true}, "taxRates": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "tax": &validation.Schema{Type: "number", AllowAdditional: true}, "stack": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"type", "tax"}}, MinItems: 1, HasMinItems: true, MaxItems: 5, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"tariffId", "currency"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"evseId", "tariff"}}

type SetDefaultTariffRequest struct {
	EVSEID     int                                `json:"evseId"`
	Tariff     SetDefaultTariffRequestTariff      `json:"tariff"`
	CustomData *SetDefaultTariffRequestCustomData `json:"customData,omitempty"`
}

type SetDefaultTariffRequestTariff struct {
	TariffID         string                                  `json:"tariffId"`
	Description      []SetDefaultTariffRequestMessageContent `json:"description,omitempty"`
	Currency         string                                  `json:"currency"`
	Energy           *SetDefaultTariffRequestTariffEnergy    `json:"energy,omitempty"`
	ValidFrom        *string                                 `json:"validFrom,omitempty"`
	ChargingTime     *SetDefaultTariffRequestTariffTime      `json:"chargingTime,omitempty"`
	IdleTime         *SetDefaultTariffRequestTariffTime      `json:"idleTime,omitempty"`
	FixedFee         *SetDefaultTariffRequestTariffFixed     `json:"fixedFee,omitempty"`
	ReservationTime  *SetDefaultTariffRequestTariffTime      `json:"reservationTime,omitempty"`
	ReservationFixed *SetDefaultTariffRequestTariffFixed     `json:"reservationFixed,omitempty"`
	MinCost          *SetDefaultTariffRequestPrice           `json:"minCost,omitempty"`
	MaxCost          *SetDefaultTariffRequestPrice           `json:"maxCost,omitempty"`
	CustomData       *SetDefaultTariffRequestCustomData      `json:"customData,omitempty"`
}

type SetDefaultTariffRequestPrice struct {
	ExclTax    *float64                           `json:"exclTax,omitempty"`
	InclTax    *float64                           `json:"inclTax,omitempty"`
	TaxRates   []SetDefaultTariffRequestTaxRate   `json:"taxRates,omitempty"`
	CustomData *SetDefaultTariffRequestCustomData `json:"customData,omitempty"`
}

type SetDefaultTariffRequestTariffFixed struct {
	Prices     []SetDefaultTariffRequestTariffFixedPrice `json:"prices"`
	TaxRates   []SetDefaultTariffRequestTaxRate          `json:"taxRates,omitempty"`
	CustomData *SetDefaultTariffRequestCustomData        `json:"customData,omitempty"`
}

type SetDefaultTariffRequestTariffFixedPrice struct {
	Conditions *SetDefaultTariffRequestTariffConditionsFixed `json:"conditions,omitempty"`
	PriceFixed float64                                       `json:"priceFixed"`
	CustomData *SetDefaultTariffRequestCustomData            `json:"customData,omitempty"`
}

type SetDefaultTariffRequestTariffConditionsFixed struct {
	StartTimeOfDay     *string                                `json:"startTimeOfDay,omitempty"`
	EndTimeOfDay       *string                                `json:"endTimeOfDay,omitempty"`
	DayOfWeek          []SetDefaultTariffRequestDayOfWeekEnum `json:"dayOfWeek,omitempty"`
	ValidFromDate      *string                                `json:"validFromDate,omitempty"`
	ValidToDate        *string                                `json:"validToDate,omitempty"`
	EVSEKind           *SetDefaultTariffRequestEVSEKindEnum   `json:"evseKind,omitempty"`
	PaymentBrand       *string                                `json:"paymentBrand,omitempty"`
	PaymentRecognition *string                                `json:"paymentRecognition,omitempty"`
	CustomData         *SetDefaultTariffRequestCustomData     `json:"customData,omitempty"`
}

type SetDefaultTariffRequestTariffTime struct {
	Prices     []SetDefaultTariffRequestTariffTimePrice `json:"prices"`
	TaxRates   []SetDefaultTariffRequestTaxRate         `json:"taxRates,omitempty"`
	CustomData *SetDefaultTariffRequestCustomData       `json:"customData,omitempty"`
}

type SetDefaultTariffRequestTariffTimePrice struct {
	PriceMinute float64                                  `json:"priceMinute"`
	Conditions  *SetDefaultTariffRequestTariffConditions `json:"conditions,omitempty"`
	CustomData  *SetDefaultTariffRequestCustomData       `json:"customData,omitempty"`
}

type SetDefaultTariffRequestTariffEnergy struct {
	Prices     []SetDefaultTariffRequestTariffEnergyPrice `json:"prices"`
	TaxRates   []SetDefaultTariffRequestTaxRate           `json:"taxRates,omitempty"`
	CustomData *SetDefaultTariffRequestCustomData         `json:"customData,omitempty"`
}

type SetDefaultTariffRequestTaxRate struct {
	Type       string                             `json:"type"`
	Tax        float64                            `json:"tax"`
	Stack      *int                               `json:"stack,omitempty"`
	CustomData *SetDefaultTariffRequestCustomData `json:"customData,omitempty"`
}

type SetDefaultTariffRequestTariffEnergyPrice struct {
	PriceKwh   float64                                  `json:"priceKwh"`
	Conditions *SetDefaultTariffRequestTariffConditions `json:"conditions,omitempty"`
	CustomData *SetDefaultTariffRequestCustomData       `json:"customData,omitempty"`
}

type SetDefaultTariffRequestTariffConditions struct {
	StartTimeOfDay  *string                                `json:"startTimeOfDay,omitempty"`
	EndTimeOfDay    *string                                `json:"endTimeOfDay,omitempty"`
	DayOfWeek       []SetDefaultTariffRequestDayOfWeekEnum `json:"dayOfWeek,omitempty"`
	ValidFromDate   *string                                `json:"validFromDate,omitempty"`
	ValidToDate     *string                                `json:"validToDate,omitempty"`
	EVSEKind        *SetDefaultTariffRequestEVSEKindEnum   `json:"evseKind,omitempty"`
	MinEnergy       *float64                               `json:"minEnergy,omitempty"`
	MaxEnergy       *float64                               `json:"maxEnergy,omitempty"`
	MinCurrent      *float64                               `json:"minCurrent,omitempty"`
	MaxCurrent      *float64                               `json:"maxCurrent,omitempty"`
	MinPower        *float64                               `json:"minPower,omitempty"`
	MaxPower        *float64                               `json:"maxPower,omitempty"`
	MinTime         *int                                   `json:"minTime,omitempty"`
	MaxTime         *int                                   `json:"maxTime,omitempty"`
	MinChargingTime *int                                   `json:"minChargingTime,omitempty"`
	MaxChargingTime *int                                   `json:"maxChargingTime,omitempty"`
	MinIdleTime     *int                                   `json:"minIdleTime,omitempty"`
	MaxIdleTime     *int                                   `json:"maxIdleTime,omitempty"`
	CustomData      *SetDefaultTariffRequestCustomData     `json:"customData,omitempty"`
}

type SetDefaultTariffRequestEVSEKindEnum string

const (
	SetDefaultTariffRequestEVSEKindEnumAC SetDefaultTariffRequestEVSEKindEnum = "AC"
	SetDefaultTariffRequestEVSEKindEnumDC SetDefaultTariffRequestEVSEKindEnum = "DC"
)

type SetDefaultTariffRequestDayOfWeekEnum string

const (
	SetDefaultTariffRequestDayOfWeekEnumMonday    SetDefaultTariffRequestDayOfWeekEnum = "Monday"
	SetDefaultTariffRequestDayOfWeekEnumTuesday   SetDefaultTariffRequestDayOfWeekEnum = "Tuesday"
	SetDefaultTariffRequestDayOfWeekEnumWednesday SetDefaultTariffRequestDayOfWeekEnum = "Wednesday"
	SetDefaultTariffRequestDayOfWeekEnumThursday  SetDefaultTariffRequestDayOfWeekEnum = "Thursday"
	SetDefaultTariffRequestDayOfWeekEnumFriday    SetDefaultTariffRequestDayOfWeekEnum = "Friday"
	SetDefaultTariffRequestDayOfWeekEnumSaturday  SetDefaultTariffRequestDayOfWeekEnum = "Saturday"
	SetDefaultTariffRequestDayOfWeekEnumSunday    SetDefaultTariffRequestDayOfWeekEnum = "Sunday"
)

type SetDefaultTariffRequestMessageContent struct {
	Format     SetDefaultTariffRequestMessageFormatEnum `json:"format"`
	Language   *string                                  `json:"language,omitempty"`
	Content    string                                   `json:"content"`
	CustomData *SetDefaultTariffRequestCustomData       `json:"customData,omitempty"`
}

type SetDefaultTariffRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

type SetDefaultTariffRequestMessageFormatEnum string

const (
	SetDefaultTariffRequestMessageFormatEnumASCII  SetDefaultTariffRequestMessageFormatEnum = "ASCII"
	SetDefaultTariffRequestMessageFormatEnumHTML   SetDefaultTariffRequestMessageFormatEnum = "HTML"
	SetDefaultTariffRequestMessageFormatEnumURI    SetDefaultTariffRequestMessageFormatEnum = "URI"
	SetDefaultTariffRequestMessageFormatEnumUTF8   SetDefaultTariffRequestMessageFormatEnum = "UTF8"
	SetDefaultTariffRequestMessageFormatEnumQRCODE SetDefaultTariffRequestMessageFormatEnum = "QRCODE"
)

func (SetDefaultTariffRequest) ActionName() string { return "SetDefaultTariff" }

func (SetDefaultTariffRequest) Version() protocol.Version { return protocol.OCPP21 }

func (SetDefaultTariffRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (SetDefaultTariffRequest) SchemaName() string { return "SetDefaultTariffRequest.json" }

func (message SetDefaultTariffRequest) Validate() error {
	return validation.Validate("SetDefaultTariffRequest.json", schemaSetDefaultTariffRequest, message)
}

func (SetDefaultTariffRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("SetDefaultTariffRequest.json", schemaSetDefaultTariffRequest, data)
}
