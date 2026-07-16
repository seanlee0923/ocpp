// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = ClearCacheConfirmation{}

var schemaClearCacheConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"status": &validation.Schema{Type: "string", Enum: []string{"Accepted", "Rejected"}}}, Required: []string{"status"}}

type ClearCacheConfirmation struct {
	Status ClearCacheConfirmationStatus `json:"status"`
}

type ClearCacheConfirmationStatus string

const (
	ClearCacheConfirmationStatusAccepted ClearCacheConfirmationStatus = "Accepted"
	ClearCacheConfirmationStatusRejected ClearCacheConfirmationStatus = "Rejected"
)

func (ClearCacheConfirmation) ActionName() string { return "ClearCache" }

func (ClearCacheConfirmation) Version() protocol.Version { return protocol.OCPP16 }

func (ClearCacheConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (ClearCacheConfirmation) SchemaName() string { return "ClearCacheResponse.json" }

func (message ClearCacheConfirmation) Validate() error {
	return validation.Validate("ClearCacheResponse.json", schemaClearCacheConfirmation, message)
}

func (ClearCacheConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("ClearCacheResponse.json", schemaClearCacheConfirmation, data)
}
