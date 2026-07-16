// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"encoding/json"

	"github.com/seanlee0923/ocpp/internal/customdata"
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = DeleteCertificateConfirmation{}

var schemaDeleteCertificateConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Failed", "NotFound"}}, "statusInfo": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"reasonCode": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "additionalInfo": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 1024, HasMaxLength: true}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"reasonCode"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"status"}}

type DeleteCertificateConfirmation struct {
	Status     DeleteCertificateConfirmationDeleteCertificateStatusEnum `json:"status"`
	StatusInfo *DeleteCertificateConfirmationStatusInfo                 `json:"statusInfo,omitempty"`
	CustomData *DeleteCertificateConfirmationCustomData                 `json:"customData,omitempty"`
}

type DeleteCertificateConfirmationStatusInfo struct {
	ReasonCode     string                                   `json:"reasonCode"`
	AdditionalInfo *string                                  `json:"additionalInfo,omitempty"`
	CustomData     *DeleteCertificateConfirmationCustomData `json:"customData,omitempty"`
}

type DeleteCertificateConfirmationCustomData struct {
	VendorID string                     `json:"vendorId"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (value DeleteCertificateConfirmationCustomData) MarshalJSON() ([]byte, error) {
	return customdata.Marshal(value.VendorID, value.Extra)
}

func (value *DeleteCertificateConfirmationCustomData) UnmarshalJSON(data []byte) error {
	return customdata.Unmarshal(data, &value.VendorID, &value.Extra)
}

type DeleteCertificateConfirmationDeleteCertificateStatusEnum string

const (
	DeleteCertificateConfirmationDeleteCertificateStatusEnumAccepted DeleteCertificateConfirmationDeleteCertificateStatusEnum = "Accepted"
	DeleteCertificateConfirmationDeleteCertificateStatusEnumFailed   DeleteCertificateConfirmationDeleteCertificateStatusEnum = "Failed"
	DeleteCertificateConfirmationDeleteCertificateStatusEnumNotFound DeleteCertificateConfirmationDeleteCertificateStatusEnum = "NotFound"
)

func (DeleteCertificateConfirmation) ActionName() string { return "DeleteCertificate" }

func (DeleteCertificateConfirmation) Version() protocol.Version { return protocol.OCPP21 }

func (DeleteCertificateConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (DeleteCertificateConfirmation) SchemaName() string { return "DeleteCertificateResponse.json" }

func (message DeleteCertificateConfirmation) Validate() error {
	return validation.Validate("DeleteCertificateResponse.json", schemaDeleteCertificateConfirmation, message)
}

func (DeleteCertificateConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("DeleteCertificateResponse.json", schemaDeleteCertificateConfirmation, data)
}
