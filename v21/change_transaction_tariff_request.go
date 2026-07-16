// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ChangeTransactionTariffRequest{}

var schemaChangeTransactionTariffRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"tariff": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"tariffId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 60, HasMaxLength: true}, "description": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"format": &validation.Schema{Type: "string", Enum: []string{"ASCII", "HTML", "URI", "UTF8", "QRCODE"}}, "language": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 8, HasMaxLength: true}, "content": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"format", "content"}}, MinItems: 1, HasMinItems: true, MaxItems: 10, HasMaxItems: true}, "currency": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 3, HasMaxLength: true}, "energy": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"prices": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"priceKwh": &validation.Schema{Type: "number", AllowAdditional: true}, "conditions": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"startTimeOfDay": &validation.Schema{Type: "string", AllowAdditional: true}, "endTimeOfDay": &validation.Schema{Type: "string", AllowAdditional: true}, "dayOfWeek": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", Enum: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}}, MinItems: 1, HasMinItems: true, MaxItems: 7, HasMaxItems: true}, "validFromDate": &validation.Schema{Type: "string", AllowAdditional: true}, "validToDate": &validation.Schema{Type: "string", AllowAdditional: true}, "evseKind": &validation.Schema{Type: "string", Enum: []string{"AC", "DC"}}, "minEnergy": &validation.Schema{Type: "number", AllowAdditional: true}, "maxEnergy": &validation.Schema{Type: "number", AllowAdditional: true}, "minCurrent": &validation.Schema{Type: "number", AllowAdditional: true}, "maxCurrent": &validation.Schema{Type: "number", AllowAdditional: true}, "minPower": &validation.Schema{Type: "number", AllowAdditional: true}, "maxPower": &validation.Schema{Type: "number", AllowAdditional: true}, "minTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "maxTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "minChargingTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "maxChargingTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "minIdleTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "maxIdleTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"priceKwh"}}, MinItems: 1, HasMinItems: true}, "taxRates": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "tax": &validation.Schema{Type: "number", AllowAdditional: true}, "stack": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"type", "tax"}}, MinItems: 1, HasMinItems: true, MaxItems: 5, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"prices"}}, "validFrom": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "chargingTime": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"prices": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"priceMinute": &validation.Schema{Type: "number", AllowAdditional: true}, "conditions": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"startTimeOfDay": &validation.Schema{Type: "string", AllowAdditional: true}, "endTimeOfDay": &validation.Schema{Type: "string", AllowAdditional: true}, "dayOfWeek": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", Enum: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}}, MinItems: 1, HasMinItems: true, MaxItems: 7, HasMaxItems: true}, "validFromDate": &validation.Schema{Type: "string", AllowAdditional: true}, "validToDate": &validation.Schema{Type: "string", AllowAdditional: true}, "evseKind": &validation.Schema{Type: "string", Enum: []string{"AC", "DC"}}, "minEnergy": &validation.Schema{Type: "number", AllowAdditional: true}, "maxEnergy": &validation.Schema{Type: "number", AllowAdditional: true}, "minCurrent": &validation.Schema{Type: "number", AllowAdditional: true}, "maxCurrent": &validation.Schema{Type: "number", AllowAdditional: true}, "minPower": &validation.Schema{Type: "number", AllowAdditional: true}, "maxPower": &validation.Schema{Type: "number", AllowAdditional: true}, "minTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "maxTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "minChargingTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "maxChargingTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "minIdleTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "maxIdleTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"priceMinute"}}, MinItems: 1, HasMinItems: true}, "taxRates": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "tax": &validation.Schema{Type: "number", AllowAdditional: true}, "stack": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"type", "tax"}}, MinItems: 1, HasMinItems: true, MaxItems: 5, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"prices"}}, "idleTime": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"prices": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"priceMinute": &validation.Schema{Type: "number", AllowAdditional: true}, "conditions": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"startTimeOfDay": &validation.Schema{Type: "string", AllowAdditional: true}, "endTimeOfDay": &validation.Schema{Type: "string", AllowAdditional: true}, "dayOfWeek": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", Enum: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}}, MinItems: 1, HasMinItems: true, MaxItems: 7, HasMaxItems: true}, "validFromDate": &validation.Schema{Type: "string", AllowAdditional: true}, "validToDate": &validation.Schema{Type: "string", AllowAdditional: true}, "evseKind": &validation.Schema{Type: "string", Enum: []string{"AC", "DC"}}, "minEnergy": &validation.Schema{Type: "number", AllowAdditional: true}, "maxEnergy": &validation.Schema{Type: "number", AllowAdditional: true}, "minCurrent": &validation.Schema{Type: "number", AllowAdditional: true}, "maxCurrent": &validation.Schema{Type: "number", AllowAdditional: true}, "minPower": &validation.Schema{Type: "number", AllowAdditional: true}, "maxPower": &validation.Schema{Type: "number", AllowAdditional: true}, "minTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "maxTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "minChargingTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "maxChargingTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "minIdleTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "maxIdleTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"priceMinute"}}, MinItems: 1, HasMinItems: true}, "taxRates": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "tax": &validation.Schema{Type: "number", AllowAdditional: true}, "stack": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"type", "tax"}}, MinItems: 1, HasMinItems: true, MaxItems: 5, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"prices"}}, "fixedFee": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"prices": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"conditions": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"startTimeOfDay": &validation.Schema{Type: "string", AllowAdditional: true}, "endTimeOfDay": &validation.Schema{Type: "string", AllowAdditional: true}, "dayOfWeek": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", Enum: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}}, MinItems: 1, HasMinItems: true, MaxItems: 7, HasMaxItems: true}, "validFromDate": &validation.Schema{Type: "string", AllowAdditional: true}, "validToDate": &validation.Schema{Type: "string", AllowAdditional: true}, "evseKind": &validation.Schema{Type: "string", Enum: []string{"AC", "DC"}}, "paymentBrand": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "paymentRecognition": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "priceFixed": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"priceFixed"}}, MinItems: 1, HasMinItems: true}, "taxRates": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "tax": &validation.Schema{Type: "number", AllowAdditional: true}, "stack": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"type", "tax"}}, MinItems: 1, HasMinItems: true, MaxItems: 5, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"prices"}}, "reservationTime": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"prices": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"priceMinute": &validation.Schema{Type: "number", AllowAdditional: true}, "conditions": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"startTimeOfDay": &validation.Schema{Type: "string", AllowAdditional: true}, "endTimeOfDay": &validation.Schema{Type: "string", AllowAdditional: true}, "dayOfWeek": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", Enum: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}}, MinItems: 1, HasMinItems: true, MaxItems: 7, HasMaxItems: true}, "validFromDate": &validation.Schema{Type: "string", AllowAdditional: true}, "validToDate": &validation.Schema{Type: "string", AllowAdditional: true}, "evseKind": &validation.Schema{Type: "string", Enum: []string{"AC", "DC"}}, "minEnergy": &validation.Schema{Type: "number", AllowAdditional: true}, "maxEnergy": &validation.Schema{Type: "number", AllowAdditional: true}, "minCurrent": &validation.Schema{Type: "number", AllowAdditional: true}, "maxCurrent": &validation.Schema{Type: "number", AllowAdditional: true}, "minPower": &validation.Schema{Type: "number", AllowAdditional: true}, "maxPower": &validation.Schema{Type: "number", AllowAdditional: true}, "minTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "maxTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "minChargingTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "maxChargingTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "minIdleTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "maxIdleTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"priceMinute"}}, MinItems: 1, HasMinItems: true}, "taxRates": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "tax": &validation.Schema{Type: "number", AllowAdditional: true}, "stack": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"type", "tax"}}, MinItems: 1, HasMinItems: true, MaxItems: 5, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"prices"}}, "reservationFixed": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"prices": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"conditions": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"startTimeOfDay": &validation.Schema{Type: "string", AllowAdditional: true}, "endTimeOfDay": &validation.Schema{Type: "string", AllowAdditional: true}, "dayOfWeek": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", Enum: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}}, MinItems: 1, HasMinItems: true, MaxItems: 7, HasMaxItems: true}, "validFromDate": &validation.Schema{Type: "string", AllowAdditional: true}, "validToDate": &validation.Schema{Type: "string", AllowAdditional: true}, "evseKind": &validation.Schema{Type: "string", Enum: []string{"AC", "DC"}}, "paymentBrand": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "paymentRecognition": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "priceFixed": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"priceFixed"}}, MinItems: 1, HasMinItems: true}, "taxRates": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "tax": &validation.Schema{Type: "number", AllowAdditional: true}, "stack": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"type", "tax"}}, MinItems: 1, HasMinItems: true, MaxItems: 5, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"prices"}}, "minCost": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exclTax": &validation.Schema{Type: "number", AllowAdditional: true}, "inclTax": &validation.Schema{Type: "number", AllowAdditional: true}, "taxRates": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "tax": &validation.Schema{Type: "number", AllowAdditional: true}, "stack": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"type", "tax"}}, MinItems: 1, HasMinItems: true, MaxItems: 5, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "maxCost": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exclTax": &validation.Schema{Type: "number", AllowAdditional: true}, "inclTax": &validation.Schema{Type: "number", AllowAdditional: true}, "taxRates": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "tax": &validation.Schema{Type: "number", AllowAdditional: true}, "stack": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"type", "tax"}}, MinItems: 1, HasMinItems: true, MaxItems: 5, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"tariffId", "currency"}}, "transactionId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"transactionId", "tariff"}}

type ChangeTransactionTariffRequest struct {
	Tariff        ChangeTransactionTariffRequestTariff      `json:"tariff"`
	TransactionID string                                    `json:"transactionId"`
	CustomData    *ChangeTransactionTariffRequestCustomData `json:"customData,omitempty"`
}

type ChangeTransactionTariffRequestTariff struct {
	TariffID         string                                         `json:"tariffId"`
	Description      []ChangeTransactionTariffRequestMessageContent `json:"description,omitempty"`
	Currency         string                                         `json:"currency"`
	Energy           *ChangeTransactionTariffRequestTariffEnergy    `json:"energy,omitempty"`
	ValidFrom        *string                                        `json:"validFrom,omitempty"`
	ChargingTime     *ChangeTransactionTariffRequestTariffTime      `json:"chargingTime,omitempty"`
	IdleTime         *ChangeTransactionTariffRequestTariffTime      `json:"idleTime,omitempty"`
	FixedFee         *ChangeTransactionTariffRequestTariffFixed     `json:"fixedFee,omitempty"`
	ReservationTime  *ChangeTransactionTariffRequestTariffTime      `json:"reservationTime,omitempty"`
	ReservationFixed *ChangeTransactionTariffRequestTariffFixed     `json:"reservationFixed,omitempty"`
	MinCost          *ChangeTransactionTariffRequestPrice           `json:"minCost,omitempty"`
	MaxCost          *ChangeTransactionTariffRequestPrice           `json:"maxCost,omitempty"`
	CustomData       *ChangeTransactionTariffRequestCustomData      `json:"customData,omitempty"`
}

type ChangeTransactionTariffRequestPrice struct {
	ExclTax    *float64                                  `json:"exclTax,omitempty"`
	InclTax    *float64                                  `json:"inclTax,omitempty"`
	TaxRates   []ChangeTransactionTariffRequestTaxRate   `json:"taxRates,omitempty"`
	CustomData *ChangeTransactionTariffRequestCustomData `json:"customData,omitempty"`
}

type ChangeTransactionTariffRequestTariffFixed struct {
	Prices     []ChangeTransactionTariffRequestTariffFixedPrice `json:"prices"`
	TaxRates   []ChangeTransactionTariffRequestTaxRate          `json:"taxRates,omitempty"`
	CustomData *ChangeTransactionTariffRequestCustomData        `json:"customData,omitempty"`
}

type ChangeTransactionTariffRequestTariffFixedPrice struct {
	Conditions *ChangeTransactionTariffRequestTariffConditionsFixed `json:"conditions,omitempty"`
	PriceFixed float64                                              `json:"priceFixed"`
	CustomData *ChangeTransactionTariffRequestCustomData            `json:"customData,omitempty"`
}

type ChangeTransactionTariffRequestTariffConditionsFixed struct {
	StartTimeOfDay     *string                                       `json:"startTimeOfDay,omitempty"`
	EndTimeOfDay       *string                                       `json:"endTimeOfDay,omitempty"`
	DayOfWeek          []ChangeTransactionTariffRequestDayOfWeekEnum `json:"dayOfWeek,omitempty"`
	ValidFromDate      *string                                       `json:"validFromDate,omitempty"`
	ValidToDate        *string                                       `json:"validToDate,omitempty"`
	EVSEKind           *ChangeTransactionTariffRequestEVSEKindEnum   `json:"evseKind,omitempty"`
	PaymentBrand       *string                                       `json:"paymentBrand,omitempty"`
	PaymentRecognition *string                                       `json:"paymentRecognition,omitempty"`
	CustomData         *ChangeTransactionTariffRequestCustomData     `json:"customData,omitempty"`
}

type ChangeTransactionTariffRequestTariffTime struct {
	Prices     []ChangeTransactionTariffRequestTariffTimePrice `json:"prices"`
	TaxRates   []ChangeTransactionTariffRequestTaxRate         `json:"taxRates,omitempty"`
	CustomData *ChangeTransactionTariffRequestCustomData       `json:"customData,omitempty"`
}

type ChangeTransactionTariffRequestTariffTimePrice struct {
	PriceMinute float64                                         `json:"priceMinute"`
	Conditions  *ChangeTransactionTariffRequestTariffConditions `json:"conditions,omitempty"`
	CustomData  *ChangeTransactionTariffRequestCustomData       `json:"customData,omitempty"`
}

type ChangeTransactionTariffRequestTariffEnergy struct {
	Prices     []ChangeTransactionTariffRequestTariffEnergyPrice `json:"prices"`
	TaxRates   []ChangeTransactionTariffRequestTaxRate           `json:"taxRates,omitempty"`
	CustomData *ChangeTransactionTariffRequestCustomData         `json:"customData,omitempty"`
}

type ChangeTransactionTariffRequestTaxRate struct {
	Type       string                                    `json:"type"`
	Tax        float64                                   `json:"tax"`
	Stack      *int                                      `json:"stack,omitempty"`
	CustomData *ChangeTransactionTariffRequestCustomData `json:"customData,omitempty"`
}

type ChangeTransactionTariffRequestTariffEnergyPrice struct {
	PriceKwh   float64                                         `json:"priceKwh"`
	Conditions *ChangeTransactionTariffRequestTariffConditions `json:"conditions,omitempty"`
	CustomData *ChangeTransactionTariffRequestCustomData       `json:"customData,omitempty"`
}

type ChangeTransactionTariffRequestTariffConditions struct {
	StartTimeOfDay  *string                                       `json:"startTimeOfDay,omitempty"`
	EndTimeOfDay    *string                                       `json:"endTimeOfDay,omitempty"`
	DayOfWeek       []ChangeTransactionTariffRequestDayOfWeekEnum `json:"dayOfWeek,omitempty"`
	ValidFromDate   *string                                       `json:"validFromDate,omitempty"`
	ValidToDate     *string                                       `json:"validToDate,omitempty"`
	EVSEKind        *ChangeTransactionTariffRequestEVSEKindEnum   `json:"evseKind,omitempty"`
	MinEnergy       *float64                                      `json:"minEnergy,omitempty"`
	MaxEnergy       *float64                                      `json:"maxEnergy,omitempty"`
	MinCurrent      *float64                                      `json:"minCurrent,omitempty"`
	MaxCurrent      *float64                                      `json:"maxCurrent,omitempty"`
	MinPower        *float64                                      `json:"minPower,omitempty"`
	MaxPower        *float64                                      `json:"maxPower,omitempty"`
	MinTime         *int                                          `json:"minTime,omitempty"`
	MaxTime         *int                                          `json:"maxTime,omitempty"`
	MinChargingTime *int                                          `json:"minChargingTime,omitempty"`
	MaxChargingTime *int                                          `json:"maxChargingTime,omitempty"`
	MinIdleTime     *int                                          `json:"minIdleTime,omitempty"`
	MaxIdleTime     *int                                          `json:"maxIdleTime,omitempty"`
	CustomData      *ChangeTransactionTariffRequestCustomData     `json:"customData,omitempty"`
}

type ChangeTransactionTariffRequestEVSEKindEnum string

const (
	ChangeTransactionTariffRequestEVSEKindEnumAC ChangeTransactionTariffRequestEVSEKindEnum = "AC"
	ChangeTransactionTariffRequestEVSEKindEnumDC ChangeTransactionTariffRequestEVSEKindEnum = "DC"
)

type ChangeTransactionTariffRequestDayOfWeekEnum string

const (
	ChangeTransactionTariffRequestDayOfWeekEnumMonday    ChangeTransactionTariffRequestDayOfWeekEnum = "Monday"
	ChangeTransactionTariffRequestDayOfWeekEnumTuesday   ChangeTransactionTariffRequestDayOfWeekEnum = "Tuesday"
	ChangeTransactionTariffRequestDayOfWeekEnumWednesday ChangeTransactionTariffRequestDayOfWeekEnum = "Wednesday"
	ChangeTransactionTariffRequestDayOfWeekEnumThursday  ChangeTransactionTariffRequestDayOfWeekEnum = "Thursday"
	ChangeTransactionTariffRequestDayOfWeekEnumFriday    ChangeTransactionTariffRequestDayOfWeekEnum = "Friday"
	ChangeTransactionTariffRequestDayOfWeekEnumSaturday  ChangeTransactionTariffRequestDayOfWeekEnum = "Saturday"
	ChangeTransactionTariffRequestDayOfWeekEnumSunday    ChangeTransactionTariffRequestDayOfWeekEnum = "Sunday"
)

type ChangeTransactionTariffRequestMessageContent struct {
	Format     ChangeTransactionTariffRequestMessageFormatEnum `json:"format"`
	Language   *string                                         `json:"language,omitempty"`
	Content    string                                          `json:"content"`
	CustomData *ChangeTransactionTariffRequestCustomData       `json:"customData,omitempty"`
}

type ChangeTransactionTariffRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

type ChangeTransactionTariffRequestMessageFormatEnum string

const (
	ChangeTransactionTariffRequestMessageFormatEnumASCII  ChangeTransactionTariffRequestMessageFormatEnum = "ASCII"
	ChangeTransactionTariffRequestMessageFormatEnumHTML   ChangeTransactionTariffRequestMessageFormatEnum = "HTML"
	ChangeTransactionTariffRequestMessageFormatEnumURI    ChangeTransactionTariffRequestMessageFormatEnum = "URI"
	ChangeTransactionTariffRequestMessageFormatEnumUTF8   ChangeTransactionTariffRequestMessageFormatEnum = "UTF8"
	ChangeTransactionTariffRequestMessageFormatEnumQRCODE ChangeTransactionTariffRequestMessageFormatEnum = "QRCODE"
)

func (ChangeTransactionTariffRequest) ActionName() string { return "ChangeTransactionTariff" }

func (ChangeTransactionTariffRequest) Version() protocol.Version { return protocol.OCPP21 }

func (ChangeTransactionTariffRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (ChangeTransactionTariffRequest) SchemaName() string {
	return "ChangeTransactionTariffRequest.json"
}

func (message ChangeTransactionTariffRequest) Validate() error {
	return validation.Validate("ChangeTransactionTariffRequest.json", schemaChangeTransactionTariffRequest, message)
}

func (ChangeTransactionTariffRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ChangeTransactionTariffRequest.json", schemaChangeTransactionTariffRequest, data)
}
