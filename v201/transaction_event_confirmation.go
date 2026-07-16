// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = TransactionEventConfirmation{}

var schemaTransactionEventConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "totalCost": &validation.Schema{Type: "number", AllowAdditional: true}, "chargingPriority": &validation.Schema{Type: "integer", AllowAdditional: true}, "idTokenInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Blocked", "ConcurrentTx", "Expired", "Invalid", "NoCredit", "NotAllowedTypeEVSE", "NotAtThisLocation", "NotAtThisTime", "Unknown"}}, "cacheExpiryDateTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "chargingPriority": &validation.Schema{Type: "integer", AllowAdditional: true}, "language1": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 8, HasMaxLength: true}, "evseId": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "integer", AllowAdditional: true}, MinItems: 1, HasMinItems: true}, "groupIdToken": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "additionalInfo": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "additionalIdToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}}, Required: []string{"additionalIdToken", "type"}}, MinItems: 1, HasMinItems: true}, "idToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "type": &validation.Schema{Type: "string", Enum: []string{"Central", "eMAID", "ISO14443", "ISO15693", "KeyCode", "Local", "MacAddress", "NoAuthorization"}}}, Required: []string{"idToken", "type"}}, "language2": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 8, HasMaxLength: true}, "personalMessage": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "format": &validation.Schema{Type: "string", Enum: []string{"ASCII", "HTML", "URI", "UTF8"}}, "language": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 8, HasMaxLength: true}, "content": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"format", "content"}}}, Required: []string{"status"}}, "updatedPersonalMessage": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "format": &validation.Schema{Type: "string", Enum: []string{"ASCII", "HTML", "URI", "UTF8"}}, "language": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 8, HasMaxLength: true}, "content": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"format", "content"}}}}

type TransactionEventConfirmation struct {
	CustomData             *TransactionEventConfirmationCustomData     `json:"customData,omitempty"`
	TotalCost              *float64                                    `json:"totalCost,omitempty"`
	ChargingPriority       *int                                        `json:"chargingPriority,omitempty"`
	IDTokenInfo            *TransactionEventConfirmationIDTokenInfo    `json:"idTokenInfo,omitempty"`
	UpdatedPersonalMessage *TransactionEventConfirmationMessageContent `json:"updatedPersonalMessage,omitempty"`
}

type TransactionEventConfirmationIDTokenInfo struct {
	CustomData          *TransactionEventConfirmationCustomData             `json:"customData,omitempty"`
	Status              TransactionEventConfirmationAuthorizationStatusEnum `json:"status"`
	CacheExpiryDateTime *string                                             `json:"cacheExpiryDateTime,omitempty"`
	ChargingPriority    *int                                                `json:"chargingPriority,omitempty"`
	Language1           *string                                             `json:"language1,omitempty"`
	EVSEID              []int                                               `json:"evseId,omitempty"`
	GroupIDToken        *TransactionEventConfirmationIDToken                `json:"groupIdToken,omitempty"`
	Language2           *string                                             `json:"language2,omitempty"`
	PersonalMessage     *TransactionEventConfirmationMessageContent         `json:"personalMessage,omitempty"`
}

type TransactionEventConfirmationMessageContent struct {
	CustomData *TransactionEventConfirmationCustomData       `json:"customData,omitempty"`
	Format     TransactionEventConfirmationMessageFormatEnum `json:"format"`
	Language   *string                                       `json:"language,omitempty"`
	Content    string                                        `json:"content"`
}

type TransactionEventConfirmationMessageFormatEnum string

const (
	TransactionEventConfirmationMessageFormatEnumASCII TransactionEventConfirmationMessageFormatEnum = "ASCII"
	TransactionEventConfirmationMessageFormatEnumHTML  TransactionEventConfirmationMessageFormatEnum = "HTML"
	TransactionEventConfirmationMessageFormatEnumURI   TransactionEventConfirmationMessageFormatEnum = "URI"
	TransactionEventConfirmationMessageFormatEnumUTF8  TransactionEventConfirmationMessageFormatEnum = "UTF8"
)

type TransactionEventConfirmationIDToken struct {
	CustomData     *TransactionEventConfirmationCustomData      `json:"customData,omitempty"`
	AdditionalInfo []TransactionEventConfirmationAdditionalInfo `json:"additionalInfo,omitempty"`
	IDToken        string                                       `json:"idToken"`
	Type           TransactionEventConfirmationIDTokenEnum      `json:"type"`
}

type TransactionEventConfirmationIDTokenEnum string

const (
	TransactionEventConfirmationIDTokenEnumCentral         TransactionEventConfirmationIDTokenEnum = "Central"
	TransactionEventConfirmationIDTokenEnumEMAID           TransactionEventConfirmationIDTokenEnum = "eMAID"
	TransactionEventConfirmationIDTokenEnumISO14443        TransactionEventConfirmationIDTokenEnum = "ISO14443"
	TransactionEventConfirmationIDTokenEnumISO15693        TransactionEventConfirmationIDTokenEnum = "ISO15693"
	TransactionEventConfirmationIDTokenEnumKeyCode         TransactionEventConfirmationIDTokenEnum = "KeyCode"
	TransactionEventConfirmationIDTokenEnumLocal           TransactionEventConfirmationIDTokenEnum = "Local"
	TransactionEventConfirmationIDTokenEnumMacAddress      TransactionEventConfirmationIDTokenEnum = "MacAddress"
	TransactionEventConfirmationIDTokenEnumNoAuthorization TransactionEventConfirmationIDTokenEnum = "NoAuthorization"
)

type TransactionEventConfirmationAdditionalInfo struct {
	CustomData        *TransactionEventConfirmationCustomData `json:"customData,omitempty"`
	AdditionalIDToken string                                  `json:"additionalIdToken"`
	Type              string                                  `json:"type"`
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

func (TransactionEventConfirmation) ActionName() string { return "TransactionEvent" }

func (TransactionEventConfirmation) Version() protocol.Version { return protocol.OCPP201 }

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
