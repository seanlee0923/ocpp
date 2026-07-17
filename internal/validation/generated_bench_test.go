package validation_test

import (
	"encoding/json"
	"testing"

	"github.com/seanlee0923/ocpp/v201"
	"github.com/seanlee0923/ocpp/v21"
)

func benchSmallPayload() ([]byte, v201.BootNotificationRequest) {
	payload := v201.BootNotificationRequest{
		Reason: v201.BootNotificationRequestBootReasonEnumPowerUp,
		ChargingStation: v201.BootNotificationRequestChargingStation{
			Model:      "AC-22K",
			VendorName: "Example",
		},
	}
	raw, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	return raw, payload
}

// benchLargePayload builds a NotifyPeriodicEventStreamRequest with 100 data
// points, representative of a batched OCPP 2.1 stream payload rather than a
// small control message.
func benchLargePayload() ([]byte, v21.NotifyPeriodicEventStreamRequest) {
	data := make([]v21.NotifyPeriodicEventStreamRequestStreamDataElement, 100)
	for i := range data {
		data[i] = v21.NotifyPeriodicEventStreamRequestStreamDataElement{
			T: float64(i),
			V: "230.4",
		}
	}
	payload := v21.NotifyPeriodicEventStreamRequest{
		ID:       1,
		Pending:  0,
		Basetime: "2026-07-16T00:00:00Z",
		Data:     data,
	}
	raw, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	return raw, payload
}

func BenchmarkValidateJSONSmall(b *testing.B) {
	raw, payload := benchSmallPayload()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if err := payload.ValidateJSON(raw); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkValidateJSONLarge(b *testing.B) {
	raw, payload := benchLargePayload()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if err := payload.ValidateJSON(raw); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnmarshalOnlySmall(b *testing.B) {
	raw, _ := benchSmallPayload()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var decoded v201.BootNotificationRequest
		if err := json.Unmarshal(raw, &decoded); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnmarshalOnlyLarge(b *testing.B) {
	raw, _ := benchLargePayload()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var decoded v21.NotifyPeriodicEventStreamRequest
		if err := json.Unmarshal(raw, &decoded); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkValidateThenUnmarshal* mirrors the real inbound handler path
// (typed_handler.go): ValidateJSON followed by json.Unmarshal on the same
// bytes, each a full pass over the payload.
func BenchmarkValidateThenUnmarshalSmall(b *testing.B) {
	raw, payload := benchSmallPayload()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if err := payload.ValidateJSON(raw); err != nil {
			b.Fatal(err)
		}
		var decoded v201.BootNotificationRequest
		if err := json.Unmarshal(raw, &decoded); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkValidateThenUnmarshalLarge(b *testing.B) {
	raw, payload := benchLargePayload()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if err := payload.ValidateJSON(raw); err != nil {
			b.Fatal(err)
		}
		var decoded v21.NotifyPeriodicEventStreamRequest
		if err := json.Unmarshal(raw, &decoded); err != nil {
			b.Fatal(err)
		}
	}
}
