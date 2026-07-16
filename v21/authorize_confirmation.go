// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = AuthorizeConfirmation{}

var schemaAuthorizeConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"idTokenInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Blocked", "ConcurrentTx", "Expired", "Invalid", "NoCredit", "NotAllowedTypeEVSE", "NotAtThisLocation", "NotAtThisTime", "Unknown"}}, "cacheExpiryDateTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "chargingPriority": &validation.Schema{Type: "integer", AllowAdditional: true}, "groupIdToken": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"additionalInfo": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"additionalIdToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}, "type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"additionalIdToken", "type"}}, MinItems: 1, HasMinItems: true}, "idToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}, "type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"idToken", "type"}}, "language1": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 8, HasMaxLength: true}, "language2": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 8, HasMaxLength: true}, "evseId": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, MinItems: 1, HasMinItems: true}, "personalMessage": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"format": &validation.Schema{Type: "string", Enum: []string{"ASCII", "HTML", "URI", "UTF8", "QRCODE"}}, "language": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 8, HasMaxLength: true}, "content": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"format", "content"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}, "certificateStatus": &validation.Schema{Type: "string", Enum: []string{"Accepted", "SignatureError", "CertificateExpired", "CertificateRevoked", "NoCertificateAvailable", "CertChainError", "ContractCancelled"}}, "allowedEnergyTransfer": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", Enum: []string{"AC_single_phase", "AC_two_phase", "AC_three_phase", "DC", "AC_BPT", "AC_BPT_DER", "AC_DER", "DC_BPT", "DC_ACDP", "DC_ACDP_BPT", "WPT"}}, MinItems: 1, HasMinItems: true}, "tariff": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"tariffId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 60, HasMaxLength: true}, "description": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"format": &validation.Schema{Type: "string", Enum: []string{"ASCII", "HTML", "URI", "UTF8", "QRCODE"}}, "language": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 8, HasMaxLength: true}, "content": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"format", "content"}}, MinItems: 1, HasMinItems: true, MaxItems: 10, HasMaxItems: true}, "currency": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 3, HasMaxLength: true}, "energy": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"prices": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"priceKwh": &validation.Schema{Type: "number", AllowAdditional: true}, "conditions": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"startTimeOfDay": &validation.Schema{Type: "string", AllowAdditional: true}, "endTimeOfDay": &validation.Schema{Type: "string", AllowAdditional: true}, "dayOfWeek": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", Enum: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}}, MinItems: 1, HasMinItems: true, MaxItems: 7, HasMaxItems: true}, "validFromDate": &validation.Schema{Type: "string", AllowAdditional: true}, "validToDate": &validation.Schema{Type: "string", AllowAdditional: true}, "evseKind": &validation.Schema{Type: "string", Enum: []string{"AC", "DC"}}, "minEnergy": &validation.Schema{Type: "number", AllowAdditional: true}, "maxEnergy": &validation.Schema{Type: "number", AllowAdditional: true}, "minCurrent": &validation.Schema{Type: "number", AllowAdditional: true}, "maxCurrent": &validation.Schema{Type: "number", AllowAdditional: true}, "minPower": &validation.Schema{Type: "number", AllowAdditional: true}, "maxPower": &validation.Schema{Type: "number", AllowAdditional: true}, "minTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "maxTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "minChargingTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "maxChargingTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "minIdleTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "maxIdleTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"priceKwh"}}, MinItems: 1, HasMinItems: true}, "taxRates": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "tax": &validation.Schema{Type: "number", AllowAdditional: true}, "stack": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"type", "tax"}}, MinItems: 1, HasMinItems: true, MaxItems: 5, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"prices"}}, "validFrom": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "chargingTime": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"prices": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"priceMinute": &validation.Schema{Type: "number", AllowAdditional: true}, "conditions": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"startTimeOfDay": &validation.Schema{Type: "string", AllowAdditional: true}, "endTimeOfDay": &validation.Schema{Type: "string", AllowAdditional: true}, "dayOfWeek": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", Enum: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}}, MinItems: 1, HasMinItems: true, MaxItems: 7, HasMaxItems: true}, "validFromDate": &validation.Schema{Type: "string", AllowAdditional: true}, "validToDate": &validation.Schema{Type: "string", AllowAdditional: true}, "evseKind": &validation.Schema{Type: "string", Enum: []string{"AC", "DC"}}, "minEnergy": &validation.Schema{Type: "number", AllowAdditional: true}, "maxEnergy": &validation.Schema{Type: "number", AllowAdditional: true}, "minCurrent": &validation.Schema{Type: "number", AllowAdditional: true}, "maxCurrent": &validation.Schema{Type: "number", AllowAdditional: true}, "minPower": &validation.Schema{Type: "number", AllowAdditional: true}, "maxPower": &validation.Schema{Type: "number", AllowAdditional: true}, "minTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "maxTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "minChargingTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "maxChargingTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "minIdleTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "maxIdleTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"priceMinute"}}, MinItems: 1, HasMinItems: true}, "taxRates": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "tax": &validation.Schema{Type: "number", AllowAdditional: true}, "stack": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"type", "tax"}}, MinItems: 1, HasMinItems: true, MaxItems: 5, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"prices"}}, "idleTime": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"prices": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"priceMinute": &validation.Schema{Type: "number", AllowAdditional: true}, "conditions": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"startTimeOfDay": &validation.Schema{Type: "string", AllowAdditional: true}, "endTimeOfDay": &validation.Schema{Type: "string", AllowAdditional: true}, "dayOfWeek": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", Enum: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}}, MinItems: 1, HasMinItems: true, MaxItems: 7, HasMaxItems: true}, "validFromDate": &validation.Schema{Type: "string", AllowAdditional: true}, "validToDate": &validation.Schema{Type: "string", AllowAdditional: true}, "evseKind": &validation.Schema{Type: "string", Enum: []string{"AC", "DC"}}, "minEnergy": &validation.Schema{Type: "number", AllowAdditional: true}, "maxEnergy": &validation.Schema{Type: "number", AllowAdditional: true}, "minCurrent": &validation.Schema{Type: "number", AllowAdditional: true}, "maxCurrent": &validation.Schema{Type: "number", AllowAdditional: true}, "minPower": &validation.Schema{Type: "number", AllowAdditional: true}, "maxPower": &validation.Schema{Type: "number", AllowAdditional: true}, "minTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "maxTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "minChargingTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "maxChargingTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "minIdleTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "maxIdleTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"priceMinute"}}, MinItems: 1, HasMinItems: true}, "taxRates": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "tax": &validation.Schema{Type: "number", AllowAdditional: true}, "stack": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"type", "tax"}}, MinItems: 1, HasMinItems: true, MaxItems: 5, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"prices"}}, "fixedFee": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"prices": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"conditions": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"startTimeOfDay": &validation.Schema{Type: "string", AllowAdditional: true}, "endTimeOfDay": &validation.Schema{Type: "string", AllowAdditional: true}, "dayOfWeek": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", Enum: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}}, MinItems: 1, HasMinItems: true, MaxItems: 7, HasMaxItems: true}, "validFromDate": &validation.Schema{Type: "string", AllowAdditional: true}, "validToDate": &validation.Schema{Type: "string", AllowAdditional: true}, "evseKind": &validation.Schema{Type: "string", Enum: []string{"AC", "DC"}}, "paymentBrand": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "paymentRecognition": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "priceFixed": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"priceFixed"}}, MinItems: 1, HasMinItems: true}, "taxRates": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "tax": &validation.Schema{Type: "number", AllowAdditional: true}, "stack": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"type", "tax"}}, MinItems: 1, HasMinItems: true, MaxItems: 5, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"prices"}}, "reservationTime": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"prices": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"priceMinute": &validation.Schema{Type: "number", AllowAdditional: true}, "conditions": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"startTimeOfDay": &validation.Schema{Type: "string", AllowAdditional: true}, "endTimeOfDay": &validation.Schema{Type: "string", AllowAdditional: true}, "dayOfWeek": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", Enum: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}}, MinItems: 1, HasMinItems: true, MaxItems: 7, HasMaxItems: true}, "validFromDate": &validation.Schema{Type: "string", AllowAdditional: true}, "validToDate": &validation.Schema{Type: "string", AllowAdditional: true}, "evseKind": &validation.Schema{Type: "string", Enum: []string{"AC", "DC"}}, "minEnergy": &validation.Schema{Type: "number", AllowAdditional: true}, "maxEnergy": &validation.Schema{Type: "number", AllowAdditional: true}, "minCurrent": &validation.Schema{Type: "number", AllowAdditional: true}, "maxCurrent": &validation.Schema{Type: "number", AllowAdditional: true}, "minPower": &validation.Schema{Type: "number", AllowAdditional: true}, "maxPower": &validation.Schema{Type: "number", AllowAdditional: true}, "minTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "maxTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "minChargingTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "maxChargingTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "minIdleTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "maxIdleTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"priceMinute"}}, MinItems: 1, HasMinItems: true}, "taxRates": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "tax": &validation.Schema{Type: "number", AllowAdditional: true}, "stack": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"type", "tax"}}, MinItems: 1, HasMinItems: true, MaxItems: 5, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"prices"}}, "reservationFixed": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"prices": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"conditions": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"startTimeOfDay": &validation.Schema{Type: "string", AllowAdditional: true}, "endTimeOfDay": &validation.Schema{Type: "string", AllowAdditional: true}, "dayOfWeek": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "string", Enum: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}}, MinItems: 1, HasMinItems: true, MaxItems: 7, HasMaxItems: true}, "validFromDate": &validation.Schema{Type: "string", AllowAdditional: true}, "validToDate": &validation.Schema{Type: "string", AllowAdditional: true}, "evseKind": &validation.Schema{Type: "string", Enum: []string{"AC", "DC"}}, "paymentBrand": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "paymentRecognition": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "priceFixed": &validation.Schema{Type: "number", AllowAdditional: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"priceFixed"}}, MinItems: 1, HasMinItems: true}, "taxRates": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "tax": &validation.Schema{Type: "number", AllowAdditional: true}, "stack": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"type", "tax"}}, MinItems: 1, HasMinItems: true, MaxItems: 5, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"prices"}}, "minCost": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exclTax": &validation.Schema{Type: "number", AllowAdditional: true}, "inclTax": &validation.Schema{Type: "number", AllowAdditional: true}, "taxRates": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "tax": &validation.Schema{Type: "number", AllowAdditional: true}, "stack": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"type", "tax"}}, MinItems: 1, HasMinItems: true, MaxItems: 5, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "maxCost": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"exclTax": &validation.Schema{Type: "number", AllowAdditional: true}, "inclTax": &validation.Schema{Type: "number", AllowAdditional: true}, "taxRates": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "tax": &validation.Schema{Type: "number", AllowAdditional: true}, "stack": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"type", "tax"}}, MinItems: 1, HasMinItems: true, MaxItems: 5, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"tariffId", "currency"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"idTokenInfo"}}

type AuthorizeConfirmation struct {
	IDTokenInfo           AuthorizeConfirmationIDTokenInfo                     `json:"idTokenInfo"`
	CertificateStatus     *AuthorizeConfirmationAuthorizeCertificateStatusEnum `json:"certificateStatus,omitempty"`
	AllowedEnergyTransfer []AuthorizeConfirmationEnergyTransferModeEnum        `json:"allowedEnergyTransfer,omitempty"`
	Tariff                *AuthorizeConfirmationTariff                         `json:"tariff,omitempty"`
	CustomData            *AuthorizeConfirmationCustomData                     `json:"customData,omitempty"`
}

type AuthorizeConfirmationTariff struct {
	TariffID         string                                `json:"tariffId"`
	Description      []AuthorizeConfirmationMessageContent `json:"description,omitempty"`
	Currency         string                                `json:"currency"`
	Energy           *AuthorizeConfirmationTariffEnergy    `json:"energy,omitempty"`
	ValidFrom        *string                               `json:"validFrom,omitempty"`
	ChargingTime     *AuthorizeConfirmationTariffTime      `json:"chargingTime,omitempty"`
	IdleTime         *AuthorizeConfirmationTariffTime      `json:"idleTime,omitempty"`
	FixedFee         *AuthorizeConfirmationTariffFixed     `json:"fixedFee,omitempty"`
	ReservationTime  *AuthorizeConfirmationTariffTime      `json:"reservationTime,omitempty"`
	ReservationFixed *AuthorizeConfirmationTariffFixed     `json:"reservationFixed,omitempty"`
	MinCost          *AuthorizeConfirmationPrice           `json:"minCost,omitempty"`
	MaxCost          *AuthorizeConfirmationPrice           `json:"maxCost,omitempty"`
	CustomData       *AuthorizeConfirmationCustomData      `json:"customData,omitempty"`
}

type AuthorizeConfirmationPrice struct {
	ExclTax    *float64                         `json:"exclTax,omitempty"`
	InclTax    *float64                         `json:"inclTax,omitempty"`
	TaxRates   []AuthorizeConfirmationTaxRate   `json:"taxRates,omitempty"`
	CustomData *AuthorizeConfirmationCustomData `json:"customData,omitempty"`
}

type AuthorizeConfirmationTariffFixed struct {
	Prices     []AuthorizeConfirmationTariffFixedPrice `json:"prices"`
	TaxRates   []AuthorizeConfirmationTaxRate          `json:"taxRates,omitempty"`
	CustomData *AuthorizeConfirmationCustomData        `json:"customData,omitempty"`
}

type AuthorizeConfirmationTariffFixedPrice struct {
	Conditions *AuthorizeConfirmationTariffConditionsFixed `json:"conditions,omitempty"`
	PriceFixed float64                                     `json:"priceFixed"`
	CustomData *AuthorizeConfirmationCustomData            `json:"customData,omitempty"`
}

type AuthorizeConfirmationTariffConditionsFixed struct {
	StartTimeOfDay     *string                              `json:"startTimeOfDay,omitempty"`
	EndTimeOfDay       *string                              `json:"endTimeOfDay,omitempty"`
	DayOfWeek          []AuthorizeConfirmationDayOfWeekEnum `json:"dayOfWeek,omitempty"`
	ValidFromDate      *string                              `json:"validFromDate,omitempty"`
	ValidToDate        *string                              `json:"validToDate,omitempty"`
	EVSEKind           *AuthorizeConfirmationEVSEKindEnum   `json:"evseKind,omitempty"`
	PaymentBrand       *string                              `json:"paymentBrand,omitempty"`
	PaymentRecognition *string                              `json:"paymentRecognition,omitempty"`
	CustomData         *AuthorizeConfirmationCustomData     `json:"customData,omitempty"`
}

type AuthorizeConfirmationTariffTime struct {
	Prices     []AuthorizeConfirmationTariffTimePrice `json:"prices"`
	TaxRates   []AuthorizeConfirmationTaxRate         `json:"taxRates,omitempty"`
	CustomData *AuthorizeConfirmationCustomData       `json:"customData,omitempty"`
}

type AuthorizeConfirmationTariffTimePrice struct {
	PriceMinute float64                                `json:"priceMinute"`
	Conditions  *AuthorizeConfirmationTariffConditions `json:"conditions,omitempty"`
	CustomData  *AuthorizeConfirmationCustomData       `json:"customData,omitempty"`
}

type AuthorizeConfirmationTariffEnergy struct {
	Prices     []AuthorizeConfirmationTariffEnergyPrice `json:"prices"`
	TaxRates   []AuthorizeConfirmationTaxRate           `json:"taxRates,omitempty"`
	CustomData *AuthorizeConfirmationCustomData         `json:"customData,omitempty"`
}

type AuthorizeConfirmationTaxRate struct {
	Type       string                           `json:"type"`
	Tax        float64                          `json:"tax"`
	Stack      *int                             `json:"stack,omitempty"`
	CustomData *AuthorizeConfirmationCustomData `json:"customData,omitempty"`
}

type AuthorizeConfirmationTariffEnergyPrice struct {
	PriceKwh   float64                                `json:"priceKwh"`
	Conditions *AuthorizeConfirmationTariffConditions `json:"conditions,omitempty"`
	CustomData *AuthorizeConfirmationCustomData       `json:"customData,omitempty"`
}

type AuthorizeConfirmationTariffConditions struct {
	StartTimeOfDay  *string                              `json:"startTimeOfDay,omitempty"`
	EndTimeOfDay    *string                              `json:"endTimeOfDay,omitempty"`
	DayOfWeek       []AuthorizeConfirmationDayOfWeekEnum `json:"dayOfWeek,omitempty"`
	ValidFromDate   *string                              `json:"validFromDate,omitempty"`
	ValidToDate     *string                              `json:"validToDate,omitempty"`
	EVSEKind        *AuthorizeConfirmationEVSEKindEnum   `json:"evseKind,omitempty"`
	MinEnergy       *float64                             `json:"minEnergy,omitempty"`
	MaxEnergy       *float64                             `json:"maxEnergy,omitempty"`
	MinCurrent      *float64                             `json:"minCurrent,omitempty"`
	MaxCurrent      *float64                             `json:"maxCurrent,omitempty"`
	MinPower        *float64                             `json:"minPower,omitempty"`
	MaxPower        *float64                             `json:"maxPower,omitempty"`
	MinTime         *int                                 `json:"minTime,omitempty"`
	MaxTime         *int                                 `json:"maxTime,omitempty"`
	MinChargingTime *int                                 `json:"minChargingTime,omitempty"`
	MaxChargingTime *int                                 `json:"maxChargingTime,omitempty"`
	MinIdleTime     *int                                 `json:"minIdleTime,omitempty"`
	MaxIdleTime     *int                                 `json:"maxIdleTime,omitempty"`
	CustomData      *AuthorizeConfirmationCustomData     `json:"customData,omitempty"`
}

type AuthorizeConfirmationEVSEKindEnum string

const (
	AuthorizeConfirmationEVSEKindEnumAC AuthorizeConfirmationEVSEKindEnum = "AC"
	AuthorizeConfirmationEVSEKindEnumDC AuthorizeConfirmationEVSEKindEnum = "DC"
)

type AuthorizeConfirmationDayOfWeekEnum string

const (
	AuthorizeConfirmationDayOfWeekEnumMonday    AuthorizeConfirmationDayOfWeekEnum = "Monday"
	AuthorizeConfirmationDayOfWeekEnumTuesday   AuthorizeConfirmationDayOfWeekEnum = "Tuesday"
	AuthorizeConfirmationDayOfWeekEnumWednesday AuthorizeConfirmationDayOfWeekEnum = "Wednesday"
	AuthorizeConfirmationDayOfWeekEnumThursday  AuthorizeConfirmationDayOfWeekEnum = "Thursday"
	AuthorizeConfirmationDayOfWeekEnumFriday    AuthorizeConfirmationDayOfWeekEnum = "Friday"
	AuthorizeConfirmationDayOfWeekEnumSaturday  AuthorizeConfirmationDayOfWeekEnum = "Saturday"
	AuthorizeConfirmationDayOfWeekEnumSunday    AuthorizeConfirmationDayOfWeekEnum = "Sunday"
)

type AuthorizeConfirmationEnergyTransferModeEnum string

const (
	AuthorizeConfirmationEnergyTransferModeEnumACSinglePhase AuthorizeConfirmationEnergyTransferModeEnum = "AC_single_phase"
	AuthorizeConfirmationEnergyTransferModeEnumACTwoPhase    AuthorizeConfirmationEnergyTransferModeEnum = "AC_two_phase"
	AuthorizeConfirmationEnergyTransferModeEnumACThreePhase  AuthorizeConfirmationEnergyTransferModeEnum = "AC_three_phase"
	AuthorizeConfirmationEnergyTransferModeEnumDC            AuthorizeConfirmationEnergyTransferModeEnum = "DC"
	AuthorizeConfirmationEnergyTransferModeEnumACBPT         AuthorizeConfirmationEnergyTransferModeEnum = "AC_BPT"
	AuthorizeConfirmationEnergyTransferModeEnumACBPTDER      AuthorizeConfirmationEnergyTransferModeEnum = "AC_BPT_DER"
	AuthorizeConfirmationEnergyTransferModeEnumACDER         AuthorizeConfirmationEnergyTransferModeEnum = "AC_DER"
	AuthorizeConfirmationEnergyTransferModeEnumDCBPT         AuthorizeConfirmationEnergyTransferModeEnum = "DC_BPT"
	AuthorizeConfirmationEnergyTransferModeEnumDCACDP        AuthorizeConfirmationEnergyTransferModeEnum = "DC_ACDP"
	AuthorizeConfirmationEnergyTransferModeEnumDCACDPBPT     AuthorizeConfirmationEnergyTransferModeEnum = "DC_ACDP_BPT"
	AuthorizeConfirmationEnergyTransferModeEnumWPT           AuthorizeConfirmationEnergyTransferModeEnum = "WPT"
)

type AuthorizeConfirmationAuthorizeCertificateStatusEnum string

const (
	AuthorizeConfirmationAuthorizeCertificateStatusEnumAccepted               AuthorizeConfirmationAuthorizeCertificateStatusEnum = "Accepted"
	AuthorizeConfirmationAuthorizeCertificateStatusEnumSignatureError         AuthorizeConfirmationAuthorizeCertificateStatusEnum = "SignatureError"
	AuthorizeConfirmationAuthorizeCertificateStatusEnumCertificateExpired     AuthorizeConfirmationAuthorizeCertificateStatusEnum = "CertificateExpired"
	AuthorizeConfirmationAuthorizeCertificateStatusEnumCertificateRevoked     AuthorizeConfirmationAuthorizeCertificateStatusEnum = "CertificateRevoked"
	AuthorizeConfirmationAuthorizeCertificateStatusEnumNoCertificateAvailable AuthorizeConfirmationAuthorizeCertificateStatusEnum = "NoCertificateAvailable"
	AuthorizeConfirmationAuthorizeCertificateStatusEnumCertChainError         AuthorizeConfirmationAuthorizeCertificateStatusEnum = "CertChainError"
	AuthorizeConfirmationAuthorizeCertificateStatusEnumContractCancelled      AuthorizeConfirmationAuthorizeCertificateStatusEnum = "ContractCancelled"
)

type AuthorizeConfirmationIDTokenInfo struct {
	Status              AuthorizeConfirmationAuthorizationStatusEnum `json:"status"`
	CacheExpiryDateTime *string                                      `json:"cacheExpiryDateTime,omitempty"`
	ChargingPriority    *int                                         `json:"chargingPriority,omitempty"`
	GroupIDToken        *AuthorizeConfirmationIDToken                `json:"groupIdToken,omitempty"`
	Language1           *string                                      `json:"language1,omitempty"`
	Language2           *string                                      `json:"language2,omitempty"`
	EVSEID              []int                                        `json:"evseId,omitempty"`
	PersonalMessage     *AuthorizeConfirmationMessageContent         `json:"personalMessage,omitempty"`
	CustomData          *AuthorizeConfirmationCustomData             `json:"customData,omitempty"`
}

type AuthorizeConfirmationMessageContent struct {
	Format     AuthorizeConfirmationMessageFormatEnum `json:"format"`
	Language   *string                                `json:"language,omitempty"`
	Content    string                                 `json:"content"`
	CustomData *AuthorizeConfirmationCustomData       `json:"customData,omitempty"`
}

type AuthorizeConfirmationMessageFormatEnum string

const (
	AuthorizeConfirmationMessageFormatEnumASCII  AuthorizeConfirmationMessageFormatEnum = "ASCII"
	AuthorizeConfirmationMessageFormatEnumHTML   AuthorizeConfirmationMessageFormatEnum = "HTML"
	AuthorizeConfirmationMessageFormatEnumURI    AuthorizeConfirmationMessageFormatEnum = "URI"
	AuthorizeConfirmationMessageFormatEnumUTF8   AuthorizeConfirmationMessageFormatEnum = "UTF8"
	AuthorizeConfirmationMessageFormatEnumQRCODE AuthorizeConfirmationMessageFormatEnum = "QRCODE"
)

type AuthorizeConfirmationIDToken struct {
	AdditionalInfo []AuthorizeConfirmationAdditionalInfo `json:"additionalInfo,omitempty"`
	IDToken        string                                `json:"idToken"`
	Type           string                                `json:"type"`
	CustomData     *AuthorizeConfirmationCustomData      `json:"customData,omitempty"`
}

type AuthorizeConfirmationAdditionalInfo struct {
	AdditionalIDToken string                           `json:"additionalIdToken"`
	Type              string                           `json:"type"`
	CustomData        *AuthorizeConfirmationCustomData `json:"customData,omitempty"`
}

type AuthorizeConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value AuthorizeConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *AuthorizeConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

type AuthorizeConfirmationAuthorizationStatusEnum string

const (
	AuthorizeConfirmationAuthorizationStatusEnumAccepted           AuthorizeConfirmationAuthorizationStatusEnum = "Accepted"
	AuthorizeConfirmationAuthorizationStatusEnumBlocked            AuthorizeConfirmationAuthorizationStatusEnum = "Blocked"
	AuthorizeConfirmationAuthorizationStatusEnumConcurrentTx       AuthorizeConfirmationAuthorizationStatusEnum = "ConcurrentTx"
	AuthorizeConfirmationAuthorizationStatusEnumExpired            AuthorizeConfirmationAuthorizationStatusEnum = "Expired"
	AuthorizeConfirmationAuthorizationStatusEnumInvalid            AuthorizeConfirmationAuthorizationStatusEnum = "Invalid"
	AuthorizeConfirmationAuthorizationStatusEnumNoCredit           AuthorizeConfirmationAuthorizationStatusEnum = "NoCredit"
	AuthorizeConfirmationAuthorizationStatusEnumNotAllowedTypeEVSE AuthorizeConfirmationAuthorizationStatusEnum = "NotAllowedTypeEVSE"
	AuthorizeConfirmationAuthorizationStatusEnumNotAtThisLocation  AuthorizeConfirmationAuthorizationStatusEnum = "NotAtThisLocation"
	AuthorizeConfirmationAuthorizationStatusEnumNotAtThisTime      AuthorizeConfirmationAuthorizationStatusEnum = "NotAtThisTime"
	AuthorizeConfirmationAuthorizationStatusEnumUnknown            AuthorizeConfirmationAuthorizationStatusEnum = "Unknown"
)

func (AuthorizeConfirmation) ActionName() string { return "Authorize" }

func (AuthorizeConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (AuthorizeConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (AuthorizeConfirmation) SchemaName() string { return "AuthorizeResponse.json" }

func (message AuthorizeConfirmation) Validate() error {
	return validation.Validate("AuthorizeResponse.json", schemaAuthorizeConfirmation, message)
}

func (AuthorizeConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("AuthorizeResponse.json", schemaAuthorizeConfirmation, data)
}
