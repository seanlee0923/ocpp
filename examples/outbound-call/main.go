// Command outbound-call demonstrates a CSMS-initiated OCPP 1.6 Reset CALL.
package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/seanlee0923/ocpp/csms"
	"github.com/seanlee0923/ocpp/profiles/ocpp16"
	"github.com/seanlee0923/ocpp/protocol"
	"github.com/seanlee0923/ocpp/v16"
)

func main() {
	router := csms.NewRouter()
	profile, err := ocpp16.NewProfile(router)
	if err != nil {
		log.Fatal(err)
	}
	if err := profile.HandleBootNotification(acceptBoot); err != nil {
		log.Fatal(err)
	}
	if err := profile.HandleHeartbeat(func(
		_ context.Context,
		session *csms.Session,
		_ v16.HeartbeatRequest,
	) (v16.HeartbeatConfirmation, error) {
		go reset(profile, session)
		return v16.HeartbeatConfirmation{CurrentTime: now()}, nil
	}); err != nil {
		log.Fatal(err)
	}
	server, err := csms.New(csms.Config{Router: router, Versions: []protocol.Version{protocol.OCPP16}})
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(":8080", server))
}

func acceptBoot(context.Context, *csms.Session, v16.BootNotificationRequest) (v16.BootNotificationConfirmation, error) {
	return v16.BootNotificationConfirmation{CurrentTime: now(), Interval: 300, Status: v16.BootNotificationConfirmationStatusAccepted}, nil
}

func reset(profile *ocpp16.Profile, session *csms.Session) {
	ctx, cancel := context.WithTimeout(session.Context(), 10*time.Second)
	defer cancel()
	confirmation, err := profile.CallReset(ctx, session, v16.ResetRequest{Type: v16.ResetRequestTypeSoft})
	if err != nil {
		log.Printf("Reset failed: %v", err)
		return
	}
	log.Printf("Reset status: %s", confirmation.Status)
}

func now() string { return time.Now().UTC().Format(time.RFC3339) }
