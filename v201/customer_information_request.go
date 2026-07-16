// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = CustomerInformationRequest{}

var schemaCustomerInformationRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "customerCertificate": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "hashAlgorithm": &validation.Schema{Type: "string", Enum: []string{"SHA256", "SHA384", "SHA512"}}, "issuerNameHash": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 128, HasMaxLength: true}, "issuerKeyHash": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 128, HasMaxLength: true}, "serialNumber": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 40, HasMaxLength: true}}, Required: []string{"hashAlgorithm", "issuerNameHash", "issuerKeyHash", "serialNumber"}}, "idToken": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "additionalInfo": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "additionalIdToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}}, Required: []string{"additionalIdToken", "type"}}, MinItems: 1, HasMinItems: true}, "idToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 36, HasMaxLength: true}, "type": &validation.Schema{Type: "string", Enum: []string{"Central", "eMAID", "ISO14443", "ISO15693", "KeyCode", "Local", "MacAddress", "NoAuthorization"}}}, Required: []string{"idToken", "type"}}, "requestId": &validation.Schema{Type: "integer", AllowAdditional: true}, "report": &validation.Schema{Type: "boolean", AllowAdditional: true}, "clear": &validation.Schema{Type: "boolean", AllowAdditional: true}, "customerIdentifier": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 64, HasMaxLength: true}}, Required: []string{"requestId", "report", "clear"}}

type CustomerInformationRequest struct {
	CustomData          *CustomerInformationRequestCustomData          `json:"customData,omitempty"`
	CustomerCertificate *CustomerInformationRequestCertificateHashData `json:"customerCertificate,omitempty"`
	IDToken             *CustomerInformationRequestIDToken             `json:"idToken,omitempty"`
	RequestID           int                                            `json:"requestId"`
	Report              bool                                           `json:"report"`
	Clear               bool                                           `json:"clear"`
	CustomerIdentifier  *string                                        `json:"customerIdentifier,omitempty"`
}

type CustomerInformationRequestIDToken struct {
	CustomData     *CustomerInformationRequestCustomData      `json:"customData,omitempty"`
	AdditionalInfo []CustomerInformationRequestAdditionalInfo `json:"additionalInfo,omitempty"`
	IDToken        string                                     `json:"idToken"`
	Type           CustomerInformationRequestIDTokenEnum      `json:"type"`
}

type CustomerInformationRequestIDTokenEnum string

const (
	CustomerInformationRequestIDTokenEnumCentral         CustomerInformationRequestIDTokenEnum = "Central"
	CustomerInformationRequestIDTokenEnumEMAID           CustomerInformationRequestIDTokenEnum = "eMAID"
	CustomerInformationRequestIDTokenEnumISO14443        CustomerInformationRequestIDTokenEnum = "ISO14443"
	CustomerInformationRequestIDTokenEnumISO15693        CustomerInformationRequestIDTokenEnum = "ISO15693"
	CustomerInformationRequestIDTokenEnumKeyCode         CustomerInformationRequestIDTokenEnum = "KeyCode"
	CustomerInformationRequestIDTokenEnumLocal           CustomerInformationRequestIDTokenEnum = "Local"
	CustomerInformationRequestIDTokenEnumMacAddress      CustomerInformationRequestIDTokenEnum = "MacAddress"
	CustomerInformationRequestIDTokenEnumNoAuthorization CustomerInformationRequestIDTokenEnum = "NoAuthorization"
)

type CustomerInformationRequestAdditionalInfo struct {
	CustomData        *CustomerInformationRequestCustomData `json:"customData,omitempty"`
	AdditionalIDToken string                                `json:"additionalIdToken"`
	Type              string                                `json:"type"`
}

type CustomerInformationRequestCertificateHashData struct {
	CustomData     *CustomerInformationRequestCustomData       `json:"customData,omitempty"`
	HashAlgorithm  CustomerInformationRequestHashAlgorithmEnum `json:"hashAlgorithm"`
	IssuerNameHash string                                      `json:"issuerNameHash"`
	IssuerKeyHash  string                                      `json:"issuerKeyHash"`
	SerialNumber   string                                      `json:"serialNumber"`
}

type CustomerInformationRequestHashAlgorithmEnum string

const (
	CustomerInformationRequestHashAlgorithmEnumSHA256 CustomerInformationRequestHashAlgorithmEnum = "SHA256"
	CustomerInformationRequestHashAlgorithmEnumSHA384 CustomerInformationRequestHashAlgorithmEnum = "SHA384"
	CustomerInformationRequestHashAlgorithmEnumSHA512 CustomerInformationRequestHashAlgorithmEnum = "SHA512"
)

type CustomerInformationRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (CustomerInformationRequest) ActionName() string { return "CustomerInformation" }

func (CustomerInformationRequest) Version() protocol.Version { return protocol.OCPP201 }

func (CustomerInformationRequest) Direction() protocol.PayloadDirection {
	return protocol.RequestPayload
}

func (CustomerInformationRequest) SchemaName() string { return "CustomerInformationRequest.json" }

func (message CustomerInformationRequest) Validate() error {
	return validation.Validate("CustomerInformationRequest.json", schemaCustomerInformationRequest, message)
}

func (CustomerInformationRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("CustomerInformationRequest.json", schemaCustomerInformationRequest, data)
}
