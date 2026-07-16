// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = SendLocalListRequest{}

var schemaSendLocalListRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"localAuthorizationList": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"idToken": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"additionalInfo": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"additionalIdToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}, "type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"additionalIdToken", "type"}}, MinItems: 1, HasMinItems: true}, "idToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}, "type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"idToken", "type"}}, "idTokenInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Blocked", "ConcurrentTx", "Expired", "Invalid", "NoCredit", "NotAllowedTypeEVSE", "NotAtThisLocation", "NotAtThisTime", "Unknown"}}, "cacheExpiryDateTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "chargingPriority": &validation.Schema{Type: "integer", AllowAdditional: true}, "groupIdToken": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"additionalInfo": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"additionalIdToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}, "type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"additionalIdToken", "type"}}, MinItems: 1, HasMinItems: true}, "idToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}, "type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"idToken", "type"}}, "language1": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 8, HasMaxLength: true}, "language2": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 8, HasMaxLength: true}, "evseId": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, MinItems: 1, HasMinItems: true}, "personalMessage": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"format": &validation.Schema{Type: "string", Enum: []string{"ASCII", "HTML", "URI", "UTF8", "QRCODE"}}, "language": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 8, HasMaxLength: true}, "content": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"format", "content"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"idToken"}}, MinItems: 1, HasMinItems: true}, "versionNumber": &validation.Schema{Type: "integer", AllowAdditional: true}, "updateType": &validation.Schema{Type: "string", Enum: []string{"Differential", "Full"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"versionNumber", "updateType"}}

type SendLocalListRequest struct {
	LocalAuthorizationList []SendLocalListRequestAuthorizationData `json:"localAuthorizationList,omitempty"`
	VersionNumber          int                                     `json:"versionNumber"`
	UpdateType             SendLocalListRequestUpdateEnum          `json:"updateType"`
	CustomData             *SendLocalListRequestCustomData         `json:"customData,omitempty"`
}

type SendLocalListRequestUpdateEnum string

const (
	SendLocalListRequestUpdateEnumDifferential SendLocalListRequestUpdateEnum = "Differential"
	SendLocalListRequestUpdateEnumFull         SendLocalListRequestUpdateEnum = "Full"
)

type SendLocalListRequestAuthorizationData struct {
	IDToken     SendLocalListRequestIDToken      `json:"idToken"`
	IDTokenInfo *SendLocalListRequestIDTokenInfo `json:"idTokenInfo,omitempty"`
	CustomData  *SendLocalListRequestCustomData  `json:"customData,omitempty"`
}

type SendLocalListRequestIDTokenInfo struct {
	Status              SendLocalListRequestAuthorizationStatusEnum `json:"status"`
	CacheExpiryDateTime *string                                     `json:"cacheExpiryDateTime,omitempty"`
	ChargingPriority    *int                                        `json:"chargingPriority,omitempty"`
	GroupIDToken        *SendLocalListRequestIDToken                `json:"groupIdToken,omitempty"`
	Language1           *string                                     `json:"language1,omitempty"`
	Language2           *string                                     `json:"language2,omitempty"`
	EVSEID              []int                                       `json:"evseId,omitempty"`
	PersonalMessage     *SendLocalListRequestMessageContent         `json:"personalMessage,omitempty"`
	CustomData          *SendLocalListRequestCustomData             `json:"customData,omitempty"`
}

type SendLocalListRequestMessageContent struct {
	Format     SendLocalListRequestMessageFormatEnum `json:"format"`
	Language   *string                               `json:"language,omitempty"`
	Content    string                                `json:"content"`
	CustomData *SendLocalListRequestCustomData       `json:"customData,omitempty"`
}

type SendLocalListRequestMessageFormatEnum string

const (
	SendLocalListRequestMessageFormatEnumASCII  SendLocalListRequestMessageFormatEnum = "ASCII"
	SendLocalListRequestMessageFormatEnumHTML   SendLocalListRequestMessageFormatEnum = "HTML"
	SendLocalListRequestMessageFormatEnumURI    SendLocalListRequestMessageFormatEnum = "URI"
	SendLocalListRequestMessageFormatEnumUTF8   SendLocalListRequestMessageFormatEnum = "UTF8"
	SendLocalListRequestMessageFormatEnumQRCODE SendLocalListRequestMessageFormatEnum = "QRCODE"
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
	AdditionalInfo []SendLocalListRequestAdditionalInfo `json:"additionalInfo,omitempty"`
	IDToken        string                               `json:"idToken"`
	Type           string                               `json:"type"`
	CustomData     *SendLocalListRequestCustomData      `json:"customData,omitempty"`
}

type SendLocalListRequestAdditionalInfo struct {
	AdditionalIDToken string                          `json:"additionalIdToken"`
	Type              string                          `json:"type"`
	CustomData        *SendLocalListRequestCustomData `json:"customData,omitempty"`
}

type SendLocalListRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (SendLocalListRequest) ActionName() string { return "SendLocalList" }

func (SendLocalListRequest) Version() protocol.Version { return protocol.OCPP21 }

func (SendLocalListRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (SendLocalListRequest) SchemaName() string { return "SendLocalListRequest.json" }

func (message SendLocalListRequest) Validate() error {
	return validation.Validate("SendLocalListRequest.json", schemaSendLocalListRequest, message)
}

func (SendLocalListRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("SendLocalListRequest.json", schemaSendLocalListRequest, data)
}
