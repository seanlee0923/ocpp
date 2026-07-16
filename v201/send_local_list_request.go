// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = SendLocalListRequest{}

var schemaSendLocalListRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "localAuthorizationList": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "idToken": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "additionalInfo": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "additionalIdToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}}, Required: []string{"additionalIdToken", "type"}}, MinItems: 1, HasMinItems: true}, "idToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "type": &validation.Schema{Type: "string", Enum: []string{"Central", "eMAID", "ISO14443", "ISO15693", "KeyCode", "Local", "MacAddress", "NoAuthorization"}}}, Required: []string{"idToken", "type"}}, "idTokenInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Blocked", "ConcurrentTx", "Expired", "Invalid", "NoCredit", "NotAllowedTypeEVSE", "NotAtThisLocation", "NotAtThisTime", "Unknown"}}, "cacheExpiryDateTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "chargingPriority": &validation.Schema{Type: "integer", AllowAdditional: true}, "language1": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 8, HasMaxLength: true}, "evseId": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "integer", AllowAdditional: true}, MinItems: 1, HasMinItems: true}, "groupIdToken": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "additionalInfo": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "additionalIdToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}}, Required: []string{"additionalIdToken", "type"}}, MinItems: 1, HasMinItems: true}, "idToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "type": &validation.Schema{Type: "string", Enum: []string{"Central", "eMAID", "ISO14443", "ISO15693", "KeyCode", "Local", "MacAddress", "NoAuthorization"}}}, Required: []string{"idToken", "type"}}, "language2": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 8, HasMaxLength: true}, "personalMessage": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "format": &validation.Schema{Type: "string", Enum: []string{"ASCII", "HTML", "URI", "UTF8"}}, "language": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 8, HasMaxLength: true}, "content": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"format", "content"}}}, Required: []string{"status"}}}, Required: []string{"idToken"}}, MinItems: 1, HasMinItems: true}, "versionNumber": &validation.Schema{Type: "integer", AllowAdditional: true}, "updateType": &validation.Schema{Type: "string", Enum: []string{"Differential", "Full"}}}, Required: []string{"versionNumber", "updateType"}}

type SendLocalListRequest struct {
	CustomData             *SendLocalListRequestCustomData         `json:"customData,omitempty"`
	LocalAuthorizationList []SendLocalListRequestAuthorizationData `json:"localAuthorizationList,omitempty"`
	VersionNumber          int                                     `json:"versionNumber"`
	UpdateType             SendLocalListRequestUpdateEnum          `json:"updateType"`
}

type SendLocalListRequestUpdateEnum string

const (
	SendLocalListRequestUpdateEnumDifferential SendLocalListRequestUpdateEnum = "Differential"
	SendLocalListRequestUpdateEnumFull         SendLocalListRequestUpdateEnum = "Full"
)

type SendLocalListRequestAuthorizationData struct {
	CustomData  *SendLocalListRequestCustomData  `json:"customData,omitempty"`
	IDToken     SendLocalListRequestIDToken      `json:"idToken"`
	IDTokenInfo *SendLocalListRequestIDTokenInfo `json:"idTokenInfo,omitempty"`
}

type SendLocalListRequestIDTokenInfo struct {
	CustomData          *SendLocalListRequestCustomData             `json:"customData,omitempty"`
	Status              SendLocalListRequestAuthorizationStatusEnum `json:"status"`
	CacheExpiryDateTime *string                                     `json:"cacheExpiryDateTime,omitempty"`
	ChargingPriority    *int                                        `json:"chargingPriority,omitempty"`
	Language1           *string                                     `json:"language1,omitempty"`
	EVSEID              []int                                       `json:"evseId,omitempty"`
	GroupIDToken        *SendLocalListRequestIDToken                `json:"groupIdToken,omitempty"`
	Language2           *string                                     `json:"language2,omitempty"`
	PersonalMessage     *SendLocalListRequestMessageContent         `json:"personalMessage,omitempty"`
}

type SendLocalListRequestMessageContent struct {
	CustomData *SendLocalListRequestCustomData       `json:"customData,omitempty"`
	Format     SendLocalListRequestMessageFormatEnum `json:"format"`
	Language   *string                               `json:"language,omitempty"`
	Content    string                                `json:"content"`
}

type SendLocalListRequestMessageFormatEnum string

const (
	SendLocalListRequestMessageFormatEnumASCII SendLocalListRequestMessageFormatEnum = "ASCII"
	SendLocalListRequestMessageFormatEnumHTML  SendLocalListRequestMessageFormatEnum = "HTML"
	SendLocalListRequestMessageFormatEnumURI   SendLocalListRequestMessageFormatEnum = "URI"
	SendLocalListRequestMessageFormatEnumUTF8  SendLocalListRequestMessageFormatEnum = "UTF8"
)

type SendLocalListRequestAuthorizationStatusEnum string

const (
	SendLocalListRequestAuthorizationStatusEnumAccepted           SendLocalListRequestAuthorizationStatusEnum = "Accepted"
	SendLocalListRequestAuthorizationStatusEnumBlocked            SendLocalListRequestAuthorizationStatusEnum = "Blocked"
	SendLocalListRequestAuthorizationStatusEnumConcurrentTx       SendLocalListRequestAuthorizationStatusEnum = "ConcurrentTx"
	SendLocalListRequestAuthorizationStatusEnumExpired            SendLocalListRequestAuthorizationStatusEnum = "Expired"
	SendLocalListRequestAuthorizationStatusEnumInvalid            SendLocalListRequestAuthorizationStatusEnum = "Invalid"
	SendLocalListRequestAuthorizationStatusEnumNoCredit           SendLocalListRequestAuthorizationStatusEnum = "NoCredit"
	SendLocalListRequestAuthorizationStatusEnumNotAllowedTypeEVSE SendLocalListRequestAuthorizationStatusEnum = "NotAllowedTypeEVSE"
	SendLocalListRequestAuthorizationStatusEnumNotAtThisLocation  SendLocalListRequestAuthorizationStatusEnum = "NotAtThisLocation"
	SendLocalListRequestAuthorizationStatusEnumNotAtThisTime      SendLocalListRequestAuthorizationStatusEnum = "NotAtThisTime"
	SendLocalListRequestAuthorizationStatusEnumUnknown            SendLocalListRequestAuthorizationStatusEnum = "Unknown"
)

type SendLocalListRequestIDToken struct {
	CustomData     *SendLocalListRequestCustomData      `json:"customData,omitempty"`
	AdditionalInfo []SendLocalListRequestAdditionalInfo `json:"additionalInfo,omitempty"`
	IDToken        string                               `json:"idToken"`
	Type           SendLocalListRequestIDTokenEnum      `json:"type"`
}

type SendLocalListRequestIDTokenEnum string

const (
	SendLocalListRequestIDTokenEnumCentral         SendLocalListRequestIDTokenEnum = "Central"
	SendLocalListRequestIDTokenEnumEMAID           SendLocalListRequestIDTokenEnum = "eMAID"
	SendLocalListRequestIDTokenEnumISO14443        SendLocalListRequestIDTokenEnum = "ISO14443"
	SendLocalListRequestIDTokenEnumISO15693        SendLocalListRequestIDTokenEnum = "ISO15693"
	SendLocalListRequestIDTokenEnumKeyCode         SendLocalListRequestIDTokenEnum = "KeyCode"
	SendLocalListRequestIDTokenEnumLocal           SendLocalListRequestIDTokenEnum = "Local"
	SendLocalListRequestIDTokenEnumMacAddress      SendLocalListRequestIDTokenEnum = "MacAddress"
	SendLocalListRequestIDTokenEnumNoAuthorization SendLocalListRequestIDTokenEnum = "NoAuthorization"
)

type SendLocalListRequestAdditionalInfo struct {
	CustomData        *SendLocalListRequestCustomData `json:"customData,omitempty"`
	AdditionalIDToken string                          `json:"additionalIdToken"`
	Type              string                          `json:"type"`
}

type SendLocalListRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (SendLocalListRequest) ActionName() string { return "SendLocalList" }

func (SendLocalListRequest) Version() protocol.Version { return protocol.OCPP201 }

func (SendLocalListRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (SendLocalListRequest) SchemaName() string { return "SendLocalListRequest.json" }

func (message SendLocalListRequest) Validate() error {
	return validation.Validate("SendLocalListRequest.json", schemaSendLocalListRequest, message)
}

func (SendLocalListRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("SendLocalListRequest.json", schemaSendLocalListRequest, data)
}
