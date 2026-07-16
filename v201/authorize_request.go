// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = AuthorizeRequest{}

var schemaAuthorizeRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "idToken": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "additionalInfo": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "additionalIdToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}}, Required: []string{"additionalIdToken", "type"}}, MinItems: 1, HasMinItems: true}, "idToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "type": &validation.Schema{Type: "string", Enum: []string{"Central", "eMAID", "ISO14443", "ISO15693", "KeyCode", "Local", "MacAddress", "NoAuthorization"}}}, Required: []string{"idToken", "type"}}, "certificate": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 5500, HasMaxLength: true}, "iso15118CertificateHashData": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "hashAlgorithm": &validation.Schema{Type: "string", Enum: []string{"SHA256", "SHA384", "SHA512"}}, "issuerNameHash": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 128, HasMaxLength: true}, "issuerKeyHash": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 128, HasMaxLength: true}, "serialNumber": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 40, HasMaxLength: true}, "responderURL": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"hashAlgorithm", "issuerNameHash", "issuerKeyHash", "serialNumber", "responderURL"}}, MinItems: 1, HasMinItems: true, MaxItems: 4, HasMaxItems: true}}, Required: []string{"idToken"}}

type AuthorizeRequest struct {
	CustomData                  *AuthorizeRequestCustomData       `json:"customData,omitempty"`
	IDToken                     AuthorizeRequestIDToken           `json:"idToken"`
	Certificate                 *string                           `json:"certificate,omitempty"`
	Iso15118CertificateHashData []AuthorizeRequestOCSPRequestData `json:"iso15118CertificateHashData,omitempty"`
}

type AuthorizeRequestOCSPRequestData struct {
	CustomData     *AuthorizeRequestCustomData       `json:"customData,omitempty"`
	HashAlgorithm  AuthorizeRequestHashAlgorithmEnum `json:"hashAlgorithm"`
	IssuerNameHash string                            `json:"issuerNameHash"`
	IssuerKeyHash  string                            `json:"issuerKeyHash"`
	SerialNumber   string                            `json:"serialNumber"`
	ResponderURL   string                            `json:"responderURL"`
}

type AuthorizeRequestHashAlgorithmEnum string

const (
	AuthorizeRequestHashAlgorithmEnumSHA256 AuthorizeRequestHashAlgorithmEnum = "SHA256"
	AuthorizeRequestHashAlgorithmEnumSHA384 AuthorizeRequestHashAlgorithmEnum = "SHA384"
	AuthorizeRequestHashAlgorithmEnumSHA512 AuthorizeRequestHashAlgorithmEnum = "SHA512"
)

type AuthorizeRequestIDToken struct {
	CustomData     *AuthorizeRequestCustomData      `json:"customData,omitempty"`
	AdditionalInfo []AuthorizeRequestAdditionalInfo `json:"additionalInfo,omitempty"`
	IDToken        string                           `json:"idToken"`
	Type           AuthorizeRequestIDTokenEnum      `json:"type"`
}

type AuthorizeRequestIDTokenEnum string

const (
	AuthorizeRequestIDTokenEnumCentral         AuthorizeRequestIDTokenEnum = "Central"
	AuthorizeRequestIDTokenEnumEMAID           AuthorizeRequestIDTokenEnum = "eMAID"
	AuthorizeRequestIDTokenEnumISO14443        AuthorizeRequestIDTokenEnum = "ISO14443"
	AuthorizeRequestIDTokenEnumISO15693        AuthorizeRequestIDTokenEnum = "ISO15693"
	AuthorizeRequestIDTokenEnumKeyCode         AuthorizeRequestIDTokenEnum = "KeyCode"
	AuthorizeRequestIDTokenEnumLocal           AuthorizeRequestIDTokenEnum = "Local"
	AuthorizeRequestIDTokenEnumMacAddress      AuthorizeRequestIDTokenEnum = "MacAddress"
	AuthorizeRequestIDTokenEnumNoAuthorization AuthorizeRequestIDTokenEnum = "NoAuthorization"
)

type AuthorizeRequestAdditionalInfo struct {
	CustomData        *AuthorizeRequestCustomData `json:"customData,omitempty"`
	AdditionalIDToken string                      `json:"additionalIdToken"`
	Type              string                      `json:"type"`
}

type AuthorizeRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (AuthorizeRequest) ActionName() string { return "Authorize" }

func (AuthorizeRequest) Version() protocol.Version { return protocol.OCPP201 }

func (AuthorizeRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (AuthorizeRequest) SchemaName() string { return "AuthorizeRequest.json" }

func (message AuthorizeRequest) Validate() error {
	return validation.Validate("AuthorizeRequest.json", schemaAuthorizeRequest, message)
}

func (AuthorizeRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("AuthorizeRequest.json", schemaAuthorizeRequest, data)
}
