// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = GetInstalledCertificateIdsConfirmation{}

var schemaGetInstalledCertificateIdsConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "NotFound"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}, "certificateHashDataChain": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "certificateHashData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "hashAlgorithm": &validation.Schema{Type: "string", Enum: []string{"SHA256", "SHA384", "SHA512"}}, "issuerNameHash": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 128, HasMaxLength: true}, "issuerKeyHash": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 128, HasMaxLength: true}, "serialNumber": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 40, HasMaxLength: true}}, Required: []string{"hashAlgorithm", "issuerNameHash", "issuerKeyHash", "serialNumber"}}, "certificateType": &validation.Schema{Type: "string", Enum: []string{"V2GRootCertificate", "MORootCertificate", "CSMSRootCertificate", "V2GCertificateChain", "ManufacturerRootCertificate"}}, "childCertificateHashData": &validation.Schema{Type: "array", AllowAdditional: true, Items: &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "hashAlgorithm": &validation.Schema{Type: "string", Enum: []string{"SHA256", "SHA384", "SHA512"}}, "issuerNameHash": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 128, HasMaxLength: true}, "issuerKeyHash": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 128, HasMaxLength: true}, "serialNumber": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 40, HasMaxLength: true}}, Required: []string{"hashAlgorithm", "issuerNameHash", "issuerKeyHash", "serialNumber"}}, MinItems: 1, HasMinItems: true, MaxItems: 4, HasMaxItems: true}}, Required: []string{"certificateType", "certificateHashData"}}, MinItems: 1, HasMinItems: true}}, Required: []string{"status"}}

type GetInstalledCertificateIdsConfirmation struct {
	CustomData               *GetInstalledCertificateIdsConfirmationCustomData                       `json:"customData,omitempty"`
	Status                   GetInstalledCertificateIdsConfirmationGetInstalledCertificateStatusEnum `json:"status"`
	StatusInfo               *GetInstalledCertificateIdsConfirmationStatusInfo                       `json:"statusInfo,omitempty"`
	CertificateHashDataChain []GetInstalledCertificateIdsConfirmationCertificateHashDataChain        `json:"certificateHashDataChain,omitempty"`
}

type GetInstalledCertificateIdsConfirmationCertificateHashDataChain struct {
	CustomData               *GetInstalledCertificateIdsConfirmationCustomData             `json:"customData,omitempty"`
	CertificateHashData      GetInstalledCertificateIdsConfirmationCertificateHashData     `json:"certificateHashData"`
	CertificateType          GetInstalledCertificateIdsConfirmationGetCertificateIDUseEnum `json:"certificateType"`
	ChildCertificateHashData []GetInstalledCertificateIdsConfirmationCertificateHashData   `json:"childCertificateHashData,omitempty"`
}

type GetInstalledCertificateIdsConfirmationGetCertificateIDUseEnum string

const (
	GetInstalledCertificateIdsConfirmationGetCertificateIDUseEnumV2GRootCertificate          GetInstalledCertificateIdsConfirmationGetCertificateIDUseEnum = "V2GRootCertificate"
	GetInstalledCertificateIdsConfirmationGetCertificateIDUseEnumMORootCertificate           GetInstalledCertificateIdsConfirmationGetCertificateIDUseEnum = "MORootCertificate"
	GetInstalledCertificateIdsConfirmationGetCertificateIDUseEnumCSMSRootCertificate         GetInstalledCertificateIdsConfirmationGetCertificateIDUseEnum = "CSMSRootCertificate"
	GetInstalledCertificateIdsConfirmationGetCertificateIDUseEnumV2GCertificateChain         GetInstalledCertificateIdsConfirmationGetCertificateIDUseEnum = "V2GCertificateChain"
	GetInstalledCertificateIdsConfirmationGetCertificateIDUseEnumManufacturerRootCertificate GetInstalledCertificateIdsConfirmationGetCertificateIDUseEnum = "ManufacturerRootCertificate"
)

type GetInstalledCertificateIdsConfirmationCertificateHashData struct {
	CustomData     *GetInstalledCertificateIdsConfirmationCustomData       `json:"customData,omitempty"`
	HashAlgorithm  GetInstalledCertificateIdsConfirmationHashAlgorithmEnum `json:"hashAlgorithm"`
	IssuerNameHash string                                                  `json:"issuerNameHash"`
	IssuerKeyHash  string                                                  `json:"issuerKeyHash"`
	SerialNumber   string                                                  `json:"serialNumber"`
}

type GetInstalledCertificateIdsConfirmationHashAlgorithmEnum string

const (
	GetInstalledCertificateIdsConfirmationHashAlgorithmEnumSHA256 GetInstalledCertificateIdsConfirmationHashAlgorithmEnum = "SHA256"
	GetInstalledCertificateIdsConfirmationHashAlgorithmEnumSHA384 GetInstalledCertificateIdsConfirmationHashAlgorithmEnum = "SHA384"
	GetInstalledCertificateIdsConfirmationHashAlgorithmEnumSHA512 GetInstalledCertificateIdsConfirmationHashAlgorithmEnum = "SHA512"
)

type GetInstalledCertificateIdsConfirmationStatusInfo struct {
	CustomData     *GetInstalledCertificateIdsConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                                            `json:"reasonCode"`
	AdditionalInfo *string                                           `json:"additionalInfo,omitempty"`
}

type GetInstalledCertificateIdsConfirmationGetInstalledCertificateStatusEnum string

const (
	GetInstalledCertificateIdsConfirmationGetInstalledCertificateStatusEnumAccepted GetInstalledCertificateIdsConfirmationGetInstalledCertificateStatusEnum = "Accepted"
	GetInstalledCertificateIdsConfirmationGetInstalledCertificateStatusEnumNotFound GetInstalledCertificateIdsConfirmationGetInstalledCertificateStatusEnum = "NotFound"
)

type GetInstalledCertificateIdsConfirmationCustomData struct {
	VendorID string `json:"vendorId"`
}

func (GetInstalledCertificateIdsConfirmation) ActionName() string {
	return "GetInstalledCertificateIds"
}

func (GetInstalledCertificateIdsConfirmation) Version() protocol.Version { return protocol.OCPP201 }

func (GetInstalledCertificateIdsConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (GetInstalledCertificateIdsConfirmation) SchemaName() string {
	return "GetInstalledCertificateIdsResponse.json"
}

func (message GetInstalledCertificateIdsConfirmation) Validate() error {
	return validation.Validate("GetInstalledCertificateIdsResponse.json", schemaGetInstalledCertificateIdsConfirmation, message)
}

func (GetInstalledCertificateIdsConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("GetInstalledCertificateIdsResponse.json", schemaGetInstalledCertificateIdsConfirmation, data)
}
