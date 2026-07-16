package validation_test

import (
	"encoding/json"
	"strings"
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

func TestGeneratedCustomDataPreservesVendorExtensions(t *testing.T) {
	raw := []byte(`{"reason":"PowerUp","chargingStation":{"model":"AC-22K","vendorName":"Example"},"customData":{"vendorId":"Example","site":"NL-01","nested":{"a":1}}}`)

	var payload v201.BootNotificationRequest
	if err := payload.ValidateJSON(raw); err != nil {
		t.Fatalf("ValidateJSON: %v", err)
	}
	if err := json.Unmarshal(raw, &payload); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if payload.CustomData == nil {
		t.Fatal("CustomData is nil")
	}
	if payload.CustomData.VendorID != "Example" {
		t.Fatalf("VendorID = %q", payload.CustomData.VendorID)
	}
	if string(payload.CustomData.Extra["site"]) != `"NL-01"` {
		t.Fatalf("Extra[site] = %s", payload.CustomData.Extra["site"])
	}
	if _, ok := payload.CustomData.Extra["vendorId"]; ok {
		t.Fatal("Extra must not duplicate vendorId")
	}

	encoded, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	if err := payload.ValidateJSON(encoded); err != nil {
		t.Fatalf("round-tripped payload failed validation: %v", err)
	}
	var decoded struct {
		CustomData map[string]json.RawMessage `json:"customData"`
	}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatalf("decode round trip: %v", err)
	}
	if string(decoded.CustomData["site"]) != `"NL-01"` {
		t.Fatalf("round-tripped customData[site] = %s", decoded.CustomData["site"])
	}
	if string(decoded.CustomData["nested"]) != `{"a":1}` {
		t.Fatalf("round-tripped customData[nested] = %s", decoded.CustomData["nested"])
	}
}

func TestGeneratedCustomDataPreservesNestedVendorExtensions(t *testing.T) {
	raw := []byte(`{"id":1,"pending":0,"basetime":"2026-07-16T00:00:00Z","data":[{"t":0,"v":"1.0","customData":{"vendorId":"Example","raw":"x"}}]}`)

	var payload v21.NotifyPeriodicEventStreamRequest
	if err := payload.ValidateJSON(raw); err != nil {
		t.Fatalf("ValidateJSON: %v", err)
	}
	if err := json.Unmarshal(raw, &payload); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if len(payload.Data) != 1 || payload.Data[0].CustomData == nil {
		t.Fatalf("Data[0].CustomData missing: %+v", payload.Data)
	}
	if string(payload.Data[0].CustomData.Extra["raw"]) != `"x"` {
		t.Fatalf("Extra[raw] = %s", payload.Data[0].CustomData.Extra["raw"])
	}

	encoded, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	if err := payload.ValidateJSON(encoded); err != nil {
		t.Fatalf("round-tripped payload failed validation: %v", err)
	}
}

func TestGeneratedCustomDataOmitsExtraWhenAbsent(t *testing.T) {
	payload := v201.BootNotificationRequest{
		Reason: v201.BootNotificationRequestBootReasonEnumPowerUp,
		ChargingStation: v201.BootNotificationRequestChargingStation{
			Model:      "AC-22K",
			VendorName: "Example",
		},
	}
	encoded, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	if strings.Contains(string(encoded), "customData") {
		t.Fatalf("expected no customData field, got %s", encoded)
	}
}
