// Command ocpp16-csms runs a minimal OCPP 1.6 CSMS server.
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
	if err := profile.HandleBootNotification(handleBootNotification); err != nil {
		log.Fatal(err)
	}

	server, err := csms.New(csms.Config{
		Router:   router,
		Versions: []protocol.Version{protocol.OCPP16},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(":8080", server))
}

func handleBootNotification(
	_ context.Context,
	_ *csms.Session,
	_ v16.BootNotificationRequest,
) (v16.BootNotificationConfirmation, error) {
	return v16.BootNotificationConfirmation{
		CurrentTime: time.Now().UTC().Format(time.RFC3339),
		Interval:    300,
		Status:      v16.BootNotificationConfirmationStatusAccepted,
	}, nil
}
