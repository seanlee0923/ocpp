// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = TransactionEventConfirmation{}

var schemaTransactionEventConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"totalCost": &validation.Schema{Type: "number", AllowAdditional: true}, "chargingPriority": &validation.Schema{Type: "integer", AllowAdditional: true}, "idTokenInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Blocked", "ConcurrentTx", "Expired", "Invalid", "NoCredit", "NotAllowedTypeEVSE", "NotAtThisLocation", "NotAtThisTime", "Unknown"}}, "cacheExpiryDateTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "chargingPriority": &validation.Schema{Type: "integer", AllowAdditional: true}, "groupIdToken": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"additionalInfo": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"additionalIdToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}, "type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"additionalIdToken", "type"}}, MinItems: 1, HasMinItems: true}, "idToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}, "type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"idToken", "type"}}, "language1": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 8, HasMaxLength: true}, "language2": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 8, HasMaxLength: true}, "evseId": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, MinItems: 1, HasMinItems: true}, "personalMessage": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"format": &validation.Schema{Type: "string", Enum: []string{"ASCII", "HTML", "URI", "UTF8", "QRCODE"}}, "language": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 8, HasMaxLength: true}, "content": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"format", "content"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}, "transactionLimit": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"maxCost": &validation.Schema{Type: "number", AllowAdditional: true}, "maxEnergy": &validation.Schema{Type: "number", AllowAdditional: true}, "maxTime": &validation.Schema{Type: "integer", AllowAdditional: true}, "maxSoC": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true, Maximum: 100, HasMaximum: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}, "updatedPersonalMessage": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"format": &validation.Schema{Type: "string", Enum: []string{"ASCII", "HTML", "URI", "UTF8", "QRCODE"}}, "language": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 8, HasMaxLength: true}, "content": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"format", "content"}}, "updatedPersonalMessageExtra": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"format": &validation.Schema{Type: "string", Enum: []string{"ASCII", "HTML", "URI", "UTF8", "QRCODE"}}, "language": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 8, HasMaxLength: true}, "content": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"format", "content"}}, MinItems: 1, HasMinItems: true, MaxItems: 4, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}}

type TransactionEventConfirmation struct {
	TotalCost                   *float64                                      `json:"totalCost,omitempty"`
	ChargingPriority            *int                                          `json:"chargingPriority,omitempty"`
	IDTokenInfo                 *TransactionEventConfirmationIDTokenInfo      `json:"idTokenInfo,omitempty"`
	TransactionLimit            *TransactionEventConfirmationTransactionLimit `json:"transactionLimit,omitempty"`
	UpdatedPersonalMessage      *TransactionEventConfirmationMessageContent   `json:"updatedPersonalMessage,omitempty"`
	UpdatedPersonalMessageExtra []TransactionEventConfirmationMessageContent  `json:"updatedPersonalMessageExtra,omitempty"`
	CustomData                  *TransactionEventConfirmationCustomData       `json:"customData,omitempty"`
}

type TransactionEventConfirmationTransactionLimit struct {
	MaxCost    *float64                                `json:"maxCost,omitempty"`
	MaxEnergy  *float64                                `json:"maxEnergy,omitempty"`
	MaxTime    *int                                    `json:"maxTime,omitempty"`
	MaxSoC     *int                                    `json:"maxSoC,omitempty"`
	CustomData *TransactionEventConfirmationCustomData `json:"customData,omitempty"`
}

type TransactionEventConfirmationIDTokenInfo struct {
	Status              TransactionEventConfirmationAuthorizationStatusEnum `json:"status"`
	CacheExpiryDateTime *string                                             `json:"cacheExpiryDateTime,omitempty"`
	ChargingPriority    *int                                                `json:"chargingPriority,omitempty"`
	GroupIDToken        *TransactionEventConfirmationIDToken                `json:"groupIdToken,omitempty"`
	Language1           *string                                             `json:"language1,omitempty"`
	Language2           *string                                             `json:"language2,omitempty"`
	EVSEID              []int                                               `json:"evseId,omitempty"`
	PersonalMessage     *TransactionEventConfirmationMessageContent         `json:"personalMessage,omitempty"`
	CustomData          *TransactionEventConfirmationCustomData             `json:"customData,omitempty"`
}

type TransactionEventConfirmationMessageContent struct {
	Format     TransactionEventConfirmationMessageFormatEnum `json:"format"`
	Language   *string                                       `json:"language,omitempty"`
	Content    string                                        `json:"content"`
	CustomData *TransactionEventConfirmationCustomData       `json:"customData,omitempty"`
}

type TransactionEventConfirmationMessageFormatEnum string

const (
	TransactionEventConfirmationMessageFormatEnumASCII  TransactionEventConfirmationMessageFormatEnum = "ASCII"
	TransactionEventConfirmationMessageFormatEnumHTML   TransactionEventConfirmationMessageFormatEnum = "HTML"
	TransactionEventConfirmationMessageFormatEnumURI    TransactionEventConfirmationMessageFormatEnum = "URI"
	TransactionEventConfirmationMessageFormatEnumUTF8   TransactionEventConfirmationMessageFormatEnum = "UTF8"
	TransactionEventConfirmationMessageFormatEnumQRCODE TransactionEventConfirmationMessageFormatEnum = "QRCODE"
)

type TransactionEventConfirmationIDToken struct {
	AdditionalInfo []TransactionEventConfirmationAdditionalInfo `json:"additionalInfo,omitempty"`
	IDToken        string                                       `json:"idToken"`
	Type           string                                       `json:"type"`
	CustomData     *TransactionEventConfirmationCustomData      `json:"customData,omitempty"`
}

type TransactionEventConfirmationAdditionalInfo struct {
	AdditionalIDToken string                                  `json:"additionalIdToken"`
	Type              string                                  `json:"type"`
	CustomData        *TransactionEventConfirmationCustomData `json:"customData,omitempty"`
}

type TransactionEventConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value TransactionEventConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *TransactionEventConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

type TransactionEventConfirmationAuthorizationStatusEnum string

const (
	TransactionEventConfirmationAuthorizationStatusEnumAccepted           TransactionEventConfirmationAuthorizationStatusEnum = "Accepted"
	TransactionEventConfirmationAuthorizationStatusEnumBlocked            TransactionEventConfirmationAuthorizationStatusEnum = "Blocked"
	TransactionEventConfirmationAuthorizationStatusEnumConcurrentTx       TransactionEventConfirmationAuthorizationStatusEnum = "ConcurrentTx"
	TransactionEventConfirmationAuthorizationStatusEnumExpired            TransactionEventConfirmationAuthorizationStatusEnum = "Expired"
	TransactionEventConfirmationAuthorizationStatusEnumInvalid            TransactionEventConfirmationAuthorizationStatusEnum = "Invalid"
	TransactionEventConfirmationAuthorizationStatusEnumNoCredit           TransactionEventConfirmationAuthorizationStatusEnum = "NoCredit"
	TransactionEventConfirmationAuthorizationStatusEnumNotAllowedTypeEVSE TransactionEventConfirmationAuthorizationStatusEnum = "NotAllowedTypeEVSE"
	TransactionEventConfirmationAuthorizationStatusEnumNotAtThisLocation  TransactionEventConfirmationAuthorizationStatusEnum = "NotAtThisLocation"
	TransactionEventConfirmationAuthorizationStatusEnumNotAtThisTime      TransactionEventConfirmationAuthorizationStatusEnum = "NotAtThisTime"
	TransactionEventConfirmationAuthorizationStatusEnumUnknown            TransactionEventConfirmationAuthorizationStatusEnum = "Unknown"
)

func (TransactionEventConfirmation) ActionName() string { return "TransactionEvent" }

func (TransactionEventConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (TransactionEventConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (TransactionEventConfirmation) SchemaName() string { return "TransactionEventResponse.json" }

func (message TransactionEventConfirmation) Validate() error {
	return validation.Validate("TransactionEventResponse.json", schemaTransactionEventConfirmation, message)
}

func (TransactionEventConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("TransactionEventResponse.json", schemaTransactionEventConfirmation, data)
}
