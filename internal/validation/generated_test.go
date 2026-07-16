package validation_test

import (
	"testing"

	"github.com/seanlee0923/ocpp/protocol"
	"github.com/seanlee0923/ocpp/v16"
	"github.com/seanlee0923/ocpp/v201"
	"github.com/seanlee0923/ocpp/v21"
)

func TestGeneratedValidationAcceptsValidPayload(t *testing.T) {
	payload := v16.BootNotificationRequest{
		ChargePointVendor: "Example",
		ChargePointModel:  "AC-22K",
	}
	if err := payload.Validate(); err != nil {
		t.Fatal(err)
	}
}

func TestGeneratedValidationRejectsConstraintsAndUnknownProperties(t *testing.T) {
	tests := []struct {
		name    string
		payload protocol.Payload
		data    []byte
	}{
		{"max length", v16.BootNotificationRequest{}, []byte(`{"chargePointVendor":"Example","chargePointModel":"model-name-that-is-longer-than-the-schema-allows"}`)},
		{"required", v16.BootNotificationRequest{}, []byte(`{"chargePointVendor":"Example"}`)},
		{"root unknown property", v16.BootNotificationRequest{}, []byte(`{"chargePointVendor":"Example","chargePointModel":"AC-22K","unknown":true}`)},
		{"enum", v16.StatusNotificationRequest{}, []byte(`{"connectorId":0,"errorCode":"NoError","status":"Invalid"}`)},
		{"date-time", v16.StatusNotificationRequest{}, []byte(`{"connectorId":0,"errorCode":"NoError","status":"Available","timestamp":"not-a-date"}`)},
		{"nested unknown property", v201.BootNotificationRequest{}, []byte(`{"reason":"PowerUp","chargingStation":{"model":"AC-22K","vendorName":"Example","unknown":true}}`)},
		{"minimum", v21.UnlockConnectorRequest{}, []byte(`{"evseId":-1,"connectorId":0}`)},
		{"min items", v21.NotifyPeriodicEventStreamRequest{}, []byte(`{"id":0,"pending":0,"basetime":"2026-07-16T00:00:00Z","data":[]}`)},
		{"trailing JSON", v16.BootNotificationRequest{}, []byte(`{"chargePointVendor":"Example","chargePointModel":"AC-22K"} {}`)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.payload.ValidateJSON(test.data)
			if err == nil {
				t.Fatalf("ValidateJSON(%s) succeeded", test.data)
			}
			if !protocol.IsValidationError(err) {
				t.Fatalf("error = %T %v, want generated validation error", err, err)
			}
		})
	}
}

func TestGeneratedValidationAcceptsRequiredZeroValuesAndCustomData(t *testing.T) {
	tests := []struct {
		payload protocol.Payload
		data    []byte
	}{
		{v16.StatusNotificationRequest{}, []byte(`{"connectorId":0,"errorCode":"NoError","status":"Available"}`)},
		{v201.BootNotificationRequest{}, []byte(`{"reason":"PowerUp","chargingStation":{"model":"AC-22K","vendorName":"Example"},"customData":{"vendorId":"example","extension":true}}`)},
	}
	for _, test := range tests {
		if err := test.payload.ValidateJSON(test.data); err != nil {
			t.Fatalf("ValidateJSON(%s): %v", test.data, err)
		}
	}
}
