package csms

import (
	"context"
	"fmt"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/seanlee0923/ocpp/internal/stationtest"
	"github.com/seanlee0923/ocpp/protocol"
	"github.com/seanlee0923/ocpp/v16"
)

// TestMultiStationTypedBootAndHeartbeat drives a full bidirectional typed
// round trip across several simulated Charging Stations concurrently:
// Station-initiated BootNotification/Heartbeat via stationtest.Call, and a
// CSMS-initiated Reset via csms.Call answered through stationtest.Handle.
// Every payload in both directions goes through the same generated-type
// schema validation a real Charging Station or CSMS would use, unlike this
// package's other tests, which construct raw wire JSON by hand.
func TestMultiStationTypedBootAndHeartbeat(t *testing.T) {
	const stationCount = 8

	router := NewRouter()
	if err := Handle(router, func(_ context.Context, _ *Session, req v16.BootNotificationRequest) (v16.BootNotificationConfirmation, error) {
		if req.ChargePointVendor == "" {
			t.Errorf("BootNotification vendor is empty")
		}
		return v16.BootNotificationConfirmation{
			Status: v16.BootNotificationConfirmationStatusAccepted, CurrentTime: "2026-07-19T00:00:00Z", Interval: 300,
		}, nil
	}); err != nil {
		t.Fatal(err)
	}
	if err := Handle(router, func(context.Context, *Session, v16.HeartbeatRequest) (v16.HeartbeatConfirmation, error) {
		return v16.HeartbeatConfirmation{CurrentTime: "2026-07-19T00:00:01Z"}, nil
	}); err != nil {
		t.Fatal(err)
	}

	connected := make(chan *Session, stationCount)
	server, err := New(Config{
		Router: router, Versions: []protocol.Version{protocol.OCPP16},
		OnConnect: func(session *Session) { connected <- session },
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()

	var wg sync.WaitGroup
	for i := range stationCount {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			identity := fmt.Sprintf("TYPED-%02d", i)
			station := stationtest.Dial(t, httpServer.URL, identity, protocol.OCPP16)
			if station == nil {
				return
			}
			if err := stationtest.Handle(station, func(_ context.Context, req v16.ResetRequest) (v16.ResetConfirmation, error) {
				if req.Type != v16.ResetRequestTypeSoft {
					t.Errorf("%s: Reset type = %q, want Soft", identity, req.Type)
				}
				return v16.ResetConfirmation{Status: v16.ResetConfirmationStatusAccepted}, nil
			}); err != nil {
				t.Errorf("%s: register Reset handler: %v", identity, err)
				return
			}

			boot, err := stationtest.Call[v16.BootNotificationRequest, v16.BootNotificationConfirmation](
				context.Background(), station, v16.BootNotificationRequest{ChargePointVendor: "Acme", ChargePointModel: "X1"},
			)
			if err != nil {
				t.Errorf("%s: BootNotification: %v", identity, err)
				return
			}
			if boot.Status != v16.BootNotificationConfirmationStatusAccepted {
				t.Errorf("%s: BootNotification status = %q", identity, boot.Status)
			}

			if _, err := stationtest.Call[v16.HeartbeatRequest, v16.HeartbeatConfirmation](
				context.Background(), station, v16.HeartbeatRequest{},
			); err != nil {
				t.Errorf("%s: Heartbeat: %v", identity, err)
			}
		}(i)
	}
	wg.Wait()

	// Every station has booted, sent a heartbeat, and registered its Reset
	// handler before this point (wg.Wait above), so it is now safe to drive
	// a CSMS-initiated typed Call against each connected session.
	for range stationCount {
		session := <-connected
		confirmation, err := Call[v16.ResetRequest, v16.ResetConfirmation](
			context.Background(), session, v16.ResetRequest{Type: v16.ResetRequestTypeSoft},
		)
		if err != nil {
			t.Errorf("outbound Reset to %s: %v", session.Identity(), err)
			continue
		}
		if confirmation.Status != v16.ResetConfirmationStatusAccepted {
			t.Errorf("outbound Reset to %s status = %q", session.Identity(), confirmation.Status)
		}
	}
}
