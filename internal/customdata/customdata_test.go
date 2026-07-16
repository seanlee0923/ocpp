package customdata_test

import (
	"encoding/json"
	"testing"

	"github.com/seanlee0923/ocpp/internal/customdata"
)

func TestMarshalUnmarshalRoundTrip(t *testing.T) {
	input := []byte(`{"vendorId":"Acme","site":"NL-01","flags":{"a":1,"b":true}}`)

	var vendorID string
	var extra map[string]json.RawMessage
	if err := customdata.Unmarshal(input, &vendorID, &extra); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if vendorID != "Acme" {
		t.Fatalf("vendorID = %q, want %q", vendorID, "Acme")
	}
	if len(extra) != 2 {
		t.Fatalf("extra = %v, want 2 entries", extra)
	}
	if string(extra["site"]) != `"NL-01"` {
		t.Fatalf("extra[site] = %s", extra["site"])
	}

	encoded, err := customdata.Marshal(vendorID, extra)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	var roundTripped map[string]json.RawMessage
	if err := json.Unmarshal(encoded, &roundTripped); err != nil {
		t.Fatalf("decode round-tripped output: %v", err)
	}
	if len(roundTripped) != 3 {
		t.Fatalf("round-tripped object has %d keys, want 3: %s", len(roundTripped), encoded)
	}
	if string(roundTripped["vendorId"]) != `"Acme"` {
		t.Fatalf("round-tripped vendorId = %s", roundTripped["vendorId"])
	}
	if string(roundTripped["flags"]) != `{"a":1,"b":true}` {
		t.Fatalf("round-tripped flags = %s", roundTripped["flags"])
	}
}

func TestUnmarshalNoExtraFields(t *testing.T) {
	var vendorID string
	var extra map[string]json.RawMessage
	if err := customdata.Unmarshal([]byte(`{"vendorId":"Acme"}`), &vendorID, &extra); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if vendorID != "Acme" {
		t.Fatalf("vendorID = %q", vendorID)
	}
	if extra != nil {
		t.Fatalf("extra = %v, want nil", extra)
	}
}

func TestMarshalVendorIDAlwaysWins(t *testing.T) {
	// A caller-populated extra map that also carries a "vendorId" key must
	// never shadow the dedicated field.
	extra := map[string]json.RawMessage{"vendorId": json.RawMessage(`"Impostor"`)}
	encoded, err := customdata.Marshal("Real", extra)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	var decoded map[string]json.RawMessage
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if string(decoded["vendorId"]) != `"Real"` {
		t.Fatalf("vendorId = %s, want %q", decoded["vendorId"], `"Real"`)
	}
}

func TestUnmarshalInvalidJSON(t *testing.T) {
	var vendorID string
	var extra map[string]json.RawMessage
	if err := customdata.Unmarshal([]byte(`not json`), &vendorID, &extra); err == nil {
		t.Fatal("expected an error for invalid JSON")
	}
}

func TestUnmarshalWrongVendorIDType(t *testing.T) {
	var vendorID string
	var extra map[string]json.RawMessage
	if err := customdata.Unmarshal([]byte(`{"vendorId":42}`), &vendorID, &extra); err == nil {
		t.Fatal("expected an error when vendorId is not a string")
	}
}
