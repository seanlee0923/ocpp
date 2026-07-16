// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v16

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = MeterValuesConfirmation{}

var schemaMeterValuesConfirmation = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{}}

type MeterValuesConfirmation struct {
}

func (MeterValuesConfirmation) ActionName() string { return "MeterValues" }

func (MeterValuesConfirmation) Version() protocol.Version { return protocol.OCPP16 }

func (MeterValuesConfirmation) Direction() protocol.PayloadDirection {
	return protocol.ConfirmationPayload
}

func (MeterValuesConfirmation) SchemaName() string { return "MeterValuesResponse.json" }

func (message MeterValuesConfirmation) Validate() error {
	return validation.Validate("MeterValuesResponse.json", schemaMeterValuesConfirmation, message)
}

func (MeterValuesConfirmation) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("MeterValuesResponse.json", schemaMeterValuesConfirmation, data)
}
