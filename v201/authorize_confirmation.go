// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = AuthorizeConfirmation{}

var schemaAuthorizeConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "idTokenInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Blocked", "ConcurrentTx", "Expired", "Invalid", "NoCredit", "NotAllowedTypeEVSE", "NotAtThisLocation", "NotAtThisTime", "Unknown"}}, "cacheExpiryDateTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "chargingPriority": &validation.Schema{Type: "integer", AllowAdditional: true}, "language1": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 8, HasMaxLength: true}, "evseId": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "integer", AllowAdditional: true}, MinItems: 1, HasMinItems: true}, "groupIdToken": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "additionalInfo": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "additionalIdToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}}, Required: []string{"additionalIdToken", "type"}}, MinItems: 1, HasMinItems: true}, "idToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "type": &validation.Schema{Type: "string", Enum: []string{"Central", "eMAID", "ISO14443", "ISO15693", "KeyCode", "Local", "MacAddress", "NoAuthorization"}}}, Required: []string{"idToken", "type"}}, "language2": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 8, HasMaxLength: true}, "personalMessage": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "format": &validation.Schema{Type: "string", Enum: []string{"ASCII", "HTML", "URI", "UTF8"}}, "language": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 8, HasMaxLength: true}, "content": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"format", "content"}}}, Required: []string{"status"}}, "certificateStatus": &validation.Schema{Type: "string", Enum: []string{"Accepted", "SignatureError", "CertificateExpired", "CertificateRevoked", "NoCertificateAvailable", "CertChainError", "ContractCancelled"}}}, Required: []string{"idTokenInfo"}}

type AuthorizeConfirmation struct {
	CustomData        *AuthorizeConfirmationCustomData                     `json:"customData,omitempty"`
	IDTokenInfo       AuthorizeConfirmationIDTokenInfo                     `json:"idTokenInfo"`
	CertificateStatus *AuthorizeConfirmationAuthorizeCertificateStatusEnum `json:"certificateStatus,omitempty"`
}

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
	CustomData          *AuthorizeConfirmationCustomData             `json:"customData,omitempty"`
	Status              AuthorizeConfirmationAuthorizationStatusEnum `json:"status"`
	CacheExpiryDateTime *string                                      `json:"cacheExpiryDateTime,omitempty"`
	ChargingPriority    *int                                         `json:"chargingPriority,omitempty"`
	Language1           *string                                      `json:"language1,omitempty"`
	EVSEID              []int                                        `json:"evseId,omitempty"`
	GroupIDToken        *AuthorizeConfirmationIDToken                `json:"groupIdToken,omitempty"`
	Language2           *string                                      `json:"language2,omitempty"`
	PersonalMessage     *AuthorizeConfirmationMessageContent         `json:"personalMessage,omitempty"`
}

type AuthorizeConfirmationMessageContent struct {
	CustomData *AuthorizeConfirmationCustomData       `json:"customData,omitempty"`
	Format     AuthorizeConfirmationMessageFormatEnum `json:"format"`
	Language   *string                                `json:"language,omitempty"`
	Content    string                                 `json:"content"`
}

type AuthorizeConfirmationMessageFormatEnum string

const (
	AuthorizeConfirmationMessageFormatEnumASCII AuthorizeConfirmationMessageFormatEnum = "ASCII"
	AuthorizeConfirmationMessageFormatEnumHTML  AuthorizeConfirmationMessageFormatEnum = "HTML"
	AuthorizeConfirmationMessageFormatEnumURI   AuthorizeConfirmationMessageFormatEnum = "URI"
	AuthorizeConfirmationMessageFormatEnumUTF8  AuthorizeConfirmationMessageFormatEnum = "UTF8"
)

type AuthorizeConfirmationIDToken struct {
	CustomData     *AuthorizeConfirmationCustomData      `json:"customData,omitempty"`
	AdditionalInfo []AuthorizeConfirmationAdditionalInfo `json:"additionalInfo,omitempty"`
	IDToken        string                                `json:"idToken"`
	Type           AuthorizeConfirmationIDTokenEnum      `json:"type"`
}

type AuthorizeConfirmationIDTokenEnum string

const (
	AuthorizeConfirmationIDTokenEnumCentral         AuthorizeConfirmationIDTokenEnum = "Central"
	AuthorizeConfirmationIDTokenEnumEMAID           AuthorizeConfirmationIDTokenEnum = "eMAID"
	AuthorizeConfirmationIDTokenEnumISO14443        AuthorizeConfirmationIDTokenEnum = "ISO14443"
	AuthorizeConfirmationIDTokenEnumISO15693        AuthorizeConfirmationIDTokenEnum = "ISO15693"
	AuthorizeConfirmationIDTokenEnumKeyCode         AuthorizeConfirmationIDTokenEnum = "KeyCode"
	AuthorizeConfirmationIDTokenEnumLocal           AuthorizeConfirmationIDTokenEnum = "Local"
	AuthorizeConfirmationIDTokenEnumMacAddress      AuthorizeConfirmationIDTokenEnum = "MacAddress"
	AuthorizeConfirmationIDTokenEnumNoAuthorization AuthorizeConfirmationIDTokenEnum = "NoAuthorization"
)

type AuthorizeConfirmationAdditionalInfo struct {
	CustomData        *AuthorizeConfirmationCustomData `json:"customData,omitempty"`
	AdditionalIDToken string                           `json:"additionalIdToken"`
	Type              string                           `json:"type"`
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

func (AuthorizeConfirmation) ActionName() string { return "Authorize" }

func (AuthorizeConfirmation) Version() protocol.Version { return protocol.OCPP201 }

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
