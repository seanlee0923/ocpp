// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = Get15118EVCertificateConfirmation{}

var schemaGet15118EVCertificateConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Failed"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}}, Required: []string{"reasonCode"}}, "exiResponse": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 5600, HasMaxLength: true}}, Required: []string{"status", "exiResponse"}}

type Get15118EVCertificateConfirmation struct {
	CustomData  *Get15118EVCertificateConfirmationCustomData                     `json:"customData,omitempty"`
	Status      Get15118EVCertificateConfirmationIso15118EVCertificateStatusEnum `json:"status"`
	StatusInfo  *Get15118EVCertificateConfirmationStatusInfo                     `json:"statusInfo,omitempty"`
	ExiResponse string                                                           `json:"exiResponse"`
}

type Get15118EVCertificateConfirmationStatusInfo struct {
	CustomData     *Get15118EVCertificateConfirmationCustomData `json:"customData,omitempty"`
	ReasonCode     string                                       `json:"reasonCode"`
	AdditionalInfo *string                                      `json:"additionalInfo,omitempty"`
}

type Get15118EVCertificateConfirmationIso15118EVCertificateStatusEnum string

const (
	Get15118EVCertificateConfirmationIso15118EVCertificateStatusEnumAccepted Get15118EVCertificateConfirmationIso15118EVCertificateStatusEnum = "Accepted"
	Get15118EVCertificateConfirmationIso15118EVCertificateStatusEnumFailed   Get15118EVCertificateConfirmationIso15118EVCertificateStatusEnum = "Failed"
)

type Get15118EVCertificateConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value Get15118EVCertificateConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *Get15118EVCertificateConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

func (Get15118EVCertificateConfirmation) ActionName() string { return "Get15118EVCertificate" }

func (Get15118EVCertificateConfirmation) Version() protocol.Version { return protocol.OCPP201 }

func (Get15118EVCertificateConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (Get15118EVCertificateConfirmation) SchemaName() string {
	return "Get15118EVCertificateResponse.json"
}

func (message Get15118EVCertificateConfirmation) Validate() error {
	return validation.Validate("Get15118EVCertificateResponse.json", schemaGet15118EVCertificateConfirmation, message)
}

func (Get15118EVCertificateConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("Get15118EVCertificateResponse.json", schemaGet15118EVCertificateConfirmation, data)
}
