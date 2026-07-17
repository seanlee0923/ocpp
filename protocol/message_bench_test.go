package protocol

import "testing"

var benchCallFrame = []byte(`[2,"bench-id","BootNotification",{"reason":"PowerUp","chargingStation":{"model":"AC-22K","vendorName":"Example"}}]`)

func BenchmarkDecodeCall(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := Decode(benchCallFrame); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkEncodeCallResult(b *testing.B) {
	message := CallResult{ID: "bench-id", Payload: []byte(`{"currentTime":"2026-07-16T00:00:00Z","interval":300,"status":"Accepted"}`)}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := Encode(message); err != nil {
			b.Fatal(err)
		}
	}
}
