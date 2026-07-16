// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = SendLocalListRequest{}

var schemaSendLocalListRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"listVersion": &validation.Schema{Type: "integer", AllowAdditional: true}, "localAuthorizationList": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"idTag": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "idTagInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"expiryDate": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "parentIdTag": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Blocked", "Expired", "Invalid", "ConcurrentTx"}}}, Required: []string{"status"}}}, Required: []string{"idTag"}}}, "updateType": &validation.Schema{Type: "string", Enum: []string{"Differential", "Full"}}}, Required: []string{"listVersion", "updateType"}}

type SendLocalListRequest struct {
	ListVersion            int                                              `json:"listVersion"`
	LocalAuthorizationList []SendLocalListRequestLocalAuthorizationListItem `json:"localAuthorizationList,omitempty"`
	UpdateType             SendLocalListRequestUpdateType                   `json:"updateType"`
}

type SendLocalListRequestUpdateType string

const (
	SendLocalListRequestUpdateTypeDifferential SendLocalListRequestUpdateType = "Differential"
	SendLocalListRequestUpdateTypeFull         SendLocalListRequestUpdateType = "Full"
)

type SendLocalListRequestLocalAuthorizationListItem struct {
	IDTag     string                                                   `json:"idTag"`
	IDTagInfo *SendLocalListRequestLocalAuthorizationListItemIDTagInfo `json:"idTagInfo,omitempty"`
}

type SendLocalListRequestLocalAuthorizationListItemIDTagInfo struct {
	ExpiryDate  *string                                                       `json:"expiryDate,omitempty"`
	ParentIDTag *string                                                       `json:"parentIdTag,omitempty"`
	Status      SendLocalListRequestLocalAuthorizationListItemIDTagInfoStatus `json:"status"`
}

type SendLocalListRequestLocalAuthorizationListItemIDTagInfoStatus string

const (
	SendLocalListRequestLocalAuthorizationListItemIDTagInfoStatusAccepted     SendLocalListRequestLocalAuthorizationListItemIDTagInfoStatus = "Accepted"
	SendLocalListRequestLocalAuthorizationListItemIDTagInfoStatusBlocked      SendLocalListRequestLocalAuthorizationListItemIDTagInfoStatus = "Blocked"
	SendLocalListRequestLocalAuthorizationListItemIDTagInfoStatusExpired      SendLocalListRequestLocalAuthorizationListItemIDTagInfoStatus = "Expired"
	SendLocalListRequestLocalAuthorizationListItemIDTagInfoStatusInvalid      SendLocalListRequestLocalAuthorizationListItemIDTagInfoStatus = "Invalid"
	SendLocalListRequestLocalAuthorizationListItemIDTagInfoStatusConcurrentTx SendLocalListRequestLocalAuthorizationListItemIDTagInfoStatus = "ConcurrentTx"
)

func (SendLocalListRequest) ActionName() string { return "SendLocalList" }

func (SendLocalListRequest) Version() protocol.Version { return protocol.OCPP16 }

func (SendLocalListRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (SendLocalListRequest) SchemaName() string { return "SendLocalList.json" }

func (message SendLocalListRequest) Validate() error {
	return validation.Validate("SendLocalList.json", schemaSendLocalListRequest, message)
}

func (SendLocalListRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("SendLocalList.json", schemaSendLocalListRequest, data)
}
