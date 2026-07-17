package validation_test

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/seanlee0923/ocpp/protocol"
)

// TestAllGeneratedTypesValidateAndRoundTrip exercises every generated OCPP
// 1.6/2.0.1/2.1 request and confirmation type at least once. Most of the 365
// generated types are otherwise only ever compiled, never run: integration
// tests cover a handful of representative actions, leaving the rest
// unverified against their own schema. This caught the customData
// vendor-extension data loss bug and guards against similar regressions in
// future schema updates or generator changes.
func TestAllGeneratedTypesValidateAndRoundTrip(t *testing.T) {
	for _, payload := range allGeneratedPayloads {
		name := string(payload.Version()) + "/" + payload.ActionName() + "/" + directionName(payload.Direction())
		t.Run(name, func(t *testing.T) {
			zeroErr := payload.Validate()
			if zeroErr != nil && !protocol.IsValidationError(zeroErr) {
				t.Fatalf("Validate() returned a non-validation error: %v", zeroErr)
			}

			raw, err := json.Marshal(payload)
			if err != nil {
				t.Fatalf("Marshal: %v", err)
			}

			jsonErr := payload.ValidateJSON(raw)
			if jsonErr != nil && !protocol.IsValidationError(jsonErr) {
				t.Fatalf("ValidateJSON() returned a non-validation error: %v", jsonErr)
			}
			if (zeroErr == nil) != (jsonErr == nil) {
				t.Fatalf("Validate() and ValidateJSON(own marshal) disagree: Validate=%v ValidateJSON=%v raw=%s", zeroErr, jsonErr, raw)
			}

			decoded := reflect.New(reflect.TypeOf(payload))
			if err := json.Unmarshal(raw, decoded.Interface()); err != nil {
				t.Fatalf("Unmarshal: %v", err)
			}
			reEncoded, err := json.Marshal(decoded.Elem().Interface())
			if err != nil {
				t.Fatalf("re-Marshal: %v", err)
			}
			if !bytes.Equal(raw, reEncoded) {
				t.Fatalf("marshal round-trip is not stable, a field was lost or reordered:\n first:  %s\n second: %s", raw, reEncoded)
			}

			injected, err := injectUnknownTopLevelProperty(raw)
			if err != nil {
				t.Fatalf("inject unknown property: %v", err)
			}
			if err := payload.ValidateJSON(injected); err == nil {
				t.Fatalf("ValidateJSON accepted an unrecognized top-level property: %s", injected)
			} else if !protocol.IsValidationError(err) {
				t.Fatalf("ValidateJSON(injected) returned a non-validation error: %v", err)
			}
		})
	}
}

func directionName(direction protocol.PayloadDirection) string {
	if direction == protocol.RequestPayload {
		return "request"
	}
	return "confirmation"
}

func injectUnknownTopLevelProperty(raw []byte) ([]byte, error) {
	var object map[string]json.RawMessage
	if err := json.Unmarshal(raw, &object); err != nil {
		return nil, err
	}
	object["zzUnrecognizedVendorProbe"] = json.RawMessage("true")
	return json.Marshal(object)
}
