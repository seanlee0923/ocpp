// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = CustomerInformationRequest{}

var schemaCustomerInformationRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customerCertificate": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"hashAlgorithm": &validation.Schema{Type: "string", Enum: []string{"SHA256", "SHA384", "SHA512"}}, "issuerNameHash": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 128, HasMaxLength: true}, "issuerKeyHash": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 128, HasMaxLength: true}, "serialNumber": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 40, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"hashAlgorithm", "issuerNameHash", "issuerKeyHash", "serialNumber"}}, "idToken": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"additionalInfo": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"additionalIdToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}, "type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"additionalIdToken", "type"}}, MinItems: 1, HasMinItems: true}, "idToken": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}, "type": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"idToken", "type"}}, "requestId": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "report": &validation.Schema{Type: "boolean", AllowAdditional: true}, "clear": &validation.Schema{Type: "boolean", AllowAdditional: true}, "customerIdentifier": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 64, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"requestId", "report", "clear"}}

type CustomerInformationRequest struct {
	CustomerCertificate *CustomerInformationRequestCertificateHashData `json:"customerCertificate,omitempty"`
	IDToken             *CustomerInformationRequestIDToken             `json:"idToken,omitempty"`
	RequestID           int                                            `json:"requestId"`
	Report              bool                                           `json:"report"`
	Clear               bool                                           `json:"clear"`
	CustomerIdentifier  *string                                        `json:"customerIdentifier,omitempty"`
	CustomData          *CustomerInformationRequestCustomData          `json:"customData,omitempty"`
}

type CustomerInformationRequestIDToken struct {
	AdditionalInfo []CustomerInformationRequestAdditionalInfo `json:"additionalInfo,omitempty"`
	IDToken        string                                     `json:"idToken"`
	Type           string                                     `json:"type"`
	CustomData     *CustomerInformationRequestCustomData      `json:"customData,omitempty"`
}

type CustomerInformationRequestAdditionalInfo struct {
	AdditionalIDToken string                                `json:"additionalIdToken"`
	Type              string                                `json:"type"`
	CustomData        *CustomerInformationRequestCustomData `json:"customData,omitempty"`
}

type CustomerInformationRequestCertificateHashData struct {
	HashAlgorithm  CustomerInformationRequestHashAlgorithmEnum `json:"hashAlgorithm"`
	IssuerNameHash string                                      `json:"issuerNameHash"`
	IssuerKeyHash  string                                      `json:"issuerKeyHash"`
	SerialNumber   string                                      `json:"serialNumber"`
	CustomData     *CustomerInformationRequestCustomData       `json:"customData,omitempty"`
}

type CustomerInformationRequestCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value CustomerInformationRequestCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *CustomerInformationRequestCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

type CustomerInformationRequestHashAlgorithmEnum string

const (
	CustomerInformationRequestHashAlgorithmEnumSHA256 CustomerInformationRequestHashAlgorithmEnum = "SHA256"
	CustomerInformationRequestHashAlgorithmEnumSHA384 CustomerInformationRequestHashAlgorithmEnum = "SHA384"
	CustomerInformationRequestHashAlgorithmEnumSHA512 CustomerInformationRequestHashAlgorithmEnum = "SHA512"
)

func (CustomerInformationRequest) ActionName() string { return "CustomerInformation" }

func (CustomerInformationRequest) Version() protocol.Version { return protocol.OCPP21 }

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
