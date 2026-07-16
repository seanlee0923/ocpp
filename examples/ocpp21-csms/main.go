// Command ocpp21-csms runs a minimal OCPP 2.1 CSMS server with a SEND handler.
package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/seanlee0923/ocpp/csms"
	"github.com/seanlee0923/ocpp/profiles/ocpp21"
	"github.com/seanlee0923/ocpp/protocol"
	"github.com/seanlee0923/ocpp/v21"
)

func main() {
	router := csms.NewRouter()
	profile, err := ocpp21.NewProfile(router)
	if err != nil {
		log.Fatal(err)
	}
	if err := profile.HandleBootNotification(handleBootNotification); err != nil {
		log.Fatal(err)
	}
	if err := profile.HandleNotifyPeriodicEventStream(handlePeriodicEvent); err != nil {
		log.Fatal(err)
	}
	server, err := csms.New(csms.Config{Router: router, Versions: []protocol.Version{protocol.OCPP21}})
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(":8080", server))
}

func handleBootNotification(
	_ context.Context,
	_ *csms.Session,
	request v21.BootNotificationRequest,
) (v21.BootNotificationConfirmation, error) {
	log.Printf("connected model=%s vendor=%s", request.ChargingStation.Model, request.ChargingStation.VendorName)
	return v21.BootNotificationConfirmation{
		CurrentTime: time.Now().UTC().Format(time.RFC3339),
		Interval:    300,
		Status:      v21.BootNotificationConfirmationRegistrationStatusEnumAccepted,
	}, nil
}

func handlePeriodicEvent(_ context.Context, session *csms.Session, _ v21.NotifyPeriodicEventStreamRequest) error {
	log.Printf("received periodic event stream from %s", session.Identity())
	return nil
}
