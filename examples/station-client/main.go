// Command station-client runs a virtual OCPP 1.6 Charging Station against a
// CSMS. It reconnects automatically, sends BootNotification followed by a
// periodic Heartbeat, and answers a CSMS-initiated Reset.
//
// Run examples/ocpp16-csms first (it listens on :8080), then run this in a
// second terminal:
//
//	go run ./examples/ocpp16-csms
//	go run ./examples/station-client
package main

import (
	"context"
	"log"
	"time"

	"github.com/seanlee0923/ocpp/protocol"
	"github.com/seanlee0923/ocpp/station"
	"github.com/seanlee0923/ocpp/v16"
)

// defaultHeartbeatInterval is used if a CSMS accepts the boot but returns a
// non-positive Interval; the schema allows this (Interval has no minimum),
// but time.NewTicker panics on a non-positive duration.
const defaultHeartbeatInterval = 300 * time.Second

func main() {
	st, err := station.New(station.Config{
		URL:      "ws://localhost:8080",
		Identity: "STATION-01",
		Version:  protocol.OCPP16,
		ReconnectPolicy: &station.ReconnectPolicy{
			InitialDelay: time.Second, MaxDelay: 30 * time.Second, Multiplier: 2,
		},
		// bootAndHeartbeat runs once per successful connection: Run's
		// reconnect loop calls OnConnect again after every reconnect, so a
		// dropped connection gets a fresh BootNotification/Heartbeat cycle
		// instead of the station going silent after the first drop.
		OnConnect:    func(st *station.Station) { log.Print("connected"); go bootAndHeartbeat(context.Background(), st) },
		OnDisconnect: func(_ *station.Station, cause error) { log.Printf("disconnected: %v", cause) },
	})
	if err != nil {
		log.Fatal(err)
	}
	if err := station.Handle(st, handleReset); err != nil {
		log.Fatal(err)
	}

	log.Fatal(st.Run(context.Background()))
}

// bootAndHeartbeat sends BootNotification and then periodic Heartbeats
// until a Call fails (most commonly because the connection dropped), at
// which point it returns — the next OnConnect starts a fresh instance, so
// exactly one instance of this loop runs per live connection.
func bootAndHeartbeat(ctx context.Context, st *station.Station) {
	boot, err := station.Call[v16.BootNotificationRequest, v16.BootNotificationConfirmation](
		ctx, st, v16.BootNotificationRequest{ChargePointVendor: "Acme", ChargePointModel: "Simulator"},
	)
	if err != nil {
		log.Printf("BootNotification failed: %v", err)
		return
	}
	log.Printf("BootNotification status: %s, interval: %ds", boot.Status, boot.Interval)
	if boot.Status != v16.BootNotificationConfirmationStatusAccepted {
		return
	}

	interval := time.Duration(boot.Interval) * time.Second
	if interval <= 0 {
		interval = defaultHeartbeatInterval
	}
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		confirmation, err := station.Call[v16.HeartbeatRequest, v16.HeartbeatConfirmation](ctx, st, v16.HeartbeatRequest{})
		if err != nil {
			log.Printf("Heartbeat failed: %v", err)
			return
		}
		log.Printf("Heartbeat: %s", confirmation.CurrentTime)
	}
}

func handleReset(_ context.Context, request v16.ResetRequest) (v16.ResetConfirmation, error) {
	log.Printf("Reset requested: %s", request.Type)
	return v16.ResetConfirmation{Status: v16.ResetConfirmationStatusAccepted}, nil
}
