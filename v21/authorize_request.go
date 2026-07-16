// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = AuthorizeRequest{}

var schemaAuthorizeRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"idToken": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"additionalInfo": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"additionalIdToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}, "type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"additionalIdToken", "type"}}, MinItems: 1, HasMinItems: true}, "idToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}, "type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"idToken", "type"}}, "certificate": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 10000, HasMaxLength: true}, "iso15118CertificateHashData": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"hashAlgorithm": &validation.Schema{Type: "string", Enum: []string{"SHA256", "SHA384", "SHA512"}}, "issuerNameHash": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 128, HasMaxLength: true}, "issuerKeyHash": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 128, HasMaxLength: true}, "serialNumber": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 40, HasMaxLength: true}, "responderURL": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 2000, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"hashAlgorithm", "issuerNameHash", "issuerKeyHash", "serialNumber", "responderURL"}}, MinItems: 1, HasMinItems: true, MaxItems: 4, HasMaxItems: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"idToken"}}

type AuthorizeRequest struct {
	IDToken                     AuthorizeRequestIDToken           `json:"idToken"`
	Certificate                 *string                           `json:"certificate,omitempty"`
	Iso15118CertificateHashData []AuthorizeRequestOCSPRequestData `json:"iso15118CertificateHashData,omitempty"`
	CustomData                  *AuthorizeRequestCustomData       `json:"customData,omitempty"`
}

type AuthorizeRequestOCSPRequestData struct {
	HashAlgorithm  AuthorizeRequestHashAlgorithmEnum `json:"hashAlgorithm"`
	IssuerNameHash string                            `json:"issuerNameHash"`
	IssuerKeyHash  string                            `json:"issuerKeyHash"`
	SerialNumber   string                            `json:"serialNumber"`
	ResponderURL   string                            `json:"responderURL"`
	CustomData     *AuthorizeRequestCustomData       `json:"customData,omitempty"`
}

type AuthorizeRequestHashAlgorithmEnum string

const (
	AuthorizeRequestHashAlgorithmEnumSHA256 AuthorizeRequestHashAlgorithmEnum = "SHA256"
	AuthorizeRequestHashAlgorithmEnumSHA384 AuthorizeRequestHashAlgorithmEnum = "SHA384"
	AuthorizeRequestHashAlgorithmEnumSHA512 AuthorizeRequestHashAlgorithmEnum = "SHA512"
)

type AuthorizeRequestIDToken struct {
	AdditionalInfo []AuthorizeRequestAdditionalInfo `json:"additionalInfo,omitempty"`
	IDToken        string                           `json:"idToken"`
	Type           string                           `json:"type"`
	CustomData     *AuthorizeRequestCustomData      `json:"customData,omitempty"`
}

type AuthorizeRequestAdditionalInfo struct {
	AdditionalIDToken string                      `json:"additionalIdToken"`
	Type              string                      `json:"type"`
	CustomData        *AuthorizeRequestCustomData `json:"customData,omitempty"`
}

type AuthorizeRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (AuthorizeRequest) ActionName() string { return "Authorize" }

func (AuthorizeRequest) Version() protocol.Version { return protocol.OCPP21 }

func (AuthorizeRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (AuthorizeRequest) SchemaName() string { return "AuthorizeRequest.json" }

func (message AuthorizeRequest) Validate() error {
	return validation.Validate("AuthorizeRequest.json", schemaAuthorizeRequest, message)
}

func (AuthorizeRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("AuthorizeRequest.json", schemaAuthorizeRequest, data)
}
