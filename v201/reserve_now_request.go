// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = ReserveNowRequest{}

var schemaReserveNowRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "id": &validation.Schema{Type: "integer", AllowAdditional: true}, "expiryDateTime": &validation.Schema{Type: "string", AllowAdditional: true, Format: "date-time"}, "connectorType": &validation.Schema{Type: "string", Enum: []string{"cCCS1", "cCCS2", "cG105", "cTesla", "cType1", "cType2", "s309-1P-16A", "s309-1P-32A", "s309-3P-16A", "s309-3P-32A", "sBS1361", "sCEE-7-7", "sType2", "sType3", "Other1PhMax16A", "Other1PhOver16A", "Other3Ph", "Pan", "wInductive", "wResonant", "Undetermined", "Unknown"}}, "idToken": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "additionalInfo": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "additionalIdToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}}, Required: []string{"additionalIdToken", "type"}}, MinItems: 1, HasMinItems: true}, "idToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "type": &validation.Schema{Type: "string", Enum: []string{"Central", "eMAID", "ISO14443", "ISO15693", "KeyCode", "Local", "MacAddress", "NoAuthorization"}}}, Required: []string{"idToken", "type"}}, "evseId": &validation.Schema{Type: "integer", AllowAdditional: true}, "groupIdToken": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "additionalInfo": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "additionalIdToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}}, Required: []string{"additionalIdToken", "type"}}, MinItems: 1, HasMinItems: true}, "idToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "type": &validation.Schema{Type: "string", Enum: []string{"Central", "eMAID", "ISO14443", "ISO15693", "KeyCode", "Local", "MacAddress", "NoAuthorization"}}}, Required: []string{"idToken", "type"}}}, Required: []string{"id", "expiryDateTime", "idToken"}}

type ReserveNowRequest struct {
	CustomData     *ReserveNowRequestCustomData    `json:"customData,omitempty"`
	ID             int                             `json:"id"`
	ExpiryDateTime string                          `json:"expiryDateTime"`
	ConnectorType  *ReserveNowRequestConnectorEnum `json:"connectorType,omitempty"`
	IDToken        ReserveNowRequestIDToken        `json:"idToken"`
	EVSEID         *int                            `json:"evseId,omitempty"`
	GroupIDToken   *ReserveNowRequestIDToken       `json:"groupIdToken,omitempty"`
}

type ReserveNowRequestIDToken struct {
	CustomData     *ReserveNowRequestCustomData      `json:"customData,omitempty"`
	AdditionalInfo []ReserveNowRequestAdditionalInfo `json:"additionalInfo,omitempty"`
	IDToken        string                            `json:"idToken"`
	Type           ReserveNowRequestIDTokenEnum      `json:"type"`
}

type ReserveNowRequestIDTokenEnum string

const (
	ReserveNowRequestIDTokenEnumCentral         ReserveNowRequestIDTokenEnum = "Central"
	ReserveNowRequestIDTokenEnumEMAID           ReserveNowRequestIDTokenEnum = "eMAID"
	ReserveNowRequestIDTokenEnumISO14443        ReserveNowRequestIDTokenEnum = "ISO14443"
	ReserveNowRequestIDTokenEnumISO15693        ReserveNowRequestIDTokenEnum = "ISO15693"
	ReserveNowRequestIDTokenEnumKeyCode         ReserveNowRequestIDTokenEnum = "KeyCode"
	ReserveNowRequestIDTokenEnumLocal           ReserveNowRequestIDTokenEnum = "Local"
	ReserveNowRequestIDTokenEnumMacAddress      ReserveNowRequestIDTokenEnum = "MacAddress"
	ReserveNowRequestIDTokenEnumNoAuthorization ReserveNowRequestIDTokenEnum = "NoAuthorization"
)

type ReserveNowRequestAdditionalInfo struct {
	CustomData        *ReserveNowRequestCustomData `json:"customData,omitempty"`
	AdditionalIDToken string                       `json:"additionalIdToken"`
	Type              string                       `json:"type"`
}

type ReserveNowRequestConnectorEnum string

const (
	ReserveNowRequestConnectorEnumCCCS1           ReserveNowRequestConnectorEnum = "cCCS1"
	ReserveNowRequestConnectorEnumCCCS2           ReserveNowRequestConnectorEnum = "cCCS2"
	ReserveNowRequestConnectorEnumCG105           ReserveNowRequestConnectorEnum = "cG105"
	ReserveNowRequestConnectorEnumCTesla          ReserveNowRequestConnectorEnum = "cTesla"
	ReserveNowRequestConnectorEnumCType1          ReserveNowRequestConnectorEnum = "cType1"
	ReserveNowRequestConnectorEnumCType2          ReserveNowRequestConnectorEnum = "cType2"
	ReserveNowRequestConnectorEnumS3091P16A       ReserveNowRequestConnectorEnum = "s309-1P-16A"
	ReserveNowRequestConnectorEnumS3091P32A       ReserveNowRequestConnectorEnum = "s309-1P-32A"
	ReserveNowRequestConnectorEnumS3093P16A       ReserveNowRequestConnectorEnum = "s309-3P-16A"
	ReserveNowRequestConnectorEnumS3093P32A       ReserveNowRequestConnectorEnum = "s309-3P-32A"
	ReserveNowRequestConnectorEnumSBS1361         ReserveNowRequestConnectorEnum = "sBS1361"
	ReserveNowRequestConnectorEnumSCEE77          ReserveNowRequestConnectorEnum = "sCEE-7-7"
	ReserveNowRequestConnectorEnumSType2          ReserveNowRequestConnectorEnum = "sType2"
	ReserveNowRequestConnectorEnumSType3          ReserveNowRequestConnectorEnum = "sType3"
	ReserveNowRequestConnectorEnumOther1PhMax16A  ReserveNowRequestConnectorEnum = "Other1PhMax16A"
	ReserveNowRequestConnectorEnumOther1PhOver16A ReserveNowRequestConnectorEnum = "Other1PhOver16A"
	ReserveNowRequestConnectorEnumOther3Ph        ReserveNowRequestConnectorEnum = "Other3Ph"
	ReserveNowRequestConnectorEnumPan             ReserveNowRequestConnectorEnum = "Pan"
	ReserveNowRequestConnectorEnumWInductive      ReserveNowRequestConnectorEnum = "wInductive"
	ReserveNowRequestConnectorEnumWResonant       ReserveNowRequestConnectorEnum = "wResonant"
	ReserveNowRequestConnectorEnumUndetermined    ReserveNowRequestConnectorEnum = "Undetermined"
	ReserveNowRequestConnectorEnumUnknown         ReserveNowRequestConnectorEnum = "Unknown"
)

type ReserveNowRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value ReserveNowRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *ReserveNowRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (ReserveNowRequest) ActionName() string { return "ReserveNow" }

func (ReserveNowRequest) Version() protocol.Version { return protocol.OCPP201 }

func (ReserveNowRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (ReserveNowRequest) SchemaName() string { return "ReserveNowRequest.json" }

func (message ReserveNowRequest) Validate() error {
	return validation.Validate("ReserveNowRequest.json", schemaReserveNowRequest, message)
}

func (ReserveNowRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ReserveNowRequest.json", schemaReserveNowRequest, data)
}
