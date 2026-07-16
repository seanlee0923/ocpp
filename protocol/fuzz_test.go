package protocol_test

import (
	"testing"

	"ocpp-go/protocol"
	"ocpp-go/v16"
	"ocpp-go/v201"
	"ocpp-go/v21"
)

func FuzzDecode(f *testing.F) {
	for _, seed := range [][]byte{
		[]byte(`[2,"id","Heartbeat",{}]`),
		[]byte(`[3,"id",{}]`),
		[]byte(`[4,"id","InternalError","failure",{}]`),
		[]byte(`[6,"id","NotifyPeriodicEventStream",{}]`),
		[]byte(`null`),
		[]byte(`{}`),
	} {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, data []byte) {
		_, _ = protocol.Decode(data)
	})
}

func FuzzGeneratedPayloadValidation(f *testing.F) {
	for _, seed := range [][]byte{
		[]byte(`{}`),
		[]byte(`{"chargePointVendor":"Example","chargePointModel":"AC-22K"}`),
		[]byte(`{"reason":"PowerUp","chargingStation":{"model":"AC-22K","vendorName":"Example"}}`),
		[]byte(`{"id":1,"pending":0,"basetime":"2026-07-16T00:00:00Z","data":[{"t":0,"v":"230.4"}]}`),
	} {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, data []byte) {
		_ = (v16.BootNotificationRequest{}).ValidateJSON(data)
		_ = (v201.BootNotificationRequest{}).ValidateJSON(data)
		_ = (v21.BootNotificationRequest{}).ValidateJSON(data)
		_ = (v21.NotifyPeriodicEventStreamRequest{}).ValidateJSON(data)
	})
}
