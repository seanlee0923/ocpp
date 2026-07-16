// Command ocpp201-csms runs a minimal OCPP 2.0.1 CSMS server.
package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/seanlee0923/ocpp/csms"
	"github.com/seanlee0923/ocpp/profiles/ocpp201"
	"github.com/seanlee0923/ocpp/protocol"
	"github.com/seanlee0923/ocpp/v201"
)

func main() {
	router := csms.NewRouter()
	profile, err := ocpp201.NewProfile(router)
	if err != nil {
		log.Fatal(err)
	}
	if err := profile.HandleBootNotification(handleBootNotification); err != nil {
		log.Fatal(err)
	}
	server, err := csms.New(csms.Config{Router: router, Versions: []protocol.Version{protocol.OCPP201}})
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(":8080", server))
}

func handleBootNotification(
	_ context.Context,
	_ *csms.Session,
	request v201.BootNotificationRequest,
) (v201.BootNotificationConfirmation, error) {
	log.Printf("connected model=%s vendor=%s", request.ChargingStation.Model, request.ChargingStation.VendorName)
	return v201.BootNotificationConfirmation{
		CurrentTime: time.Now().UTC().Format(time.RFC3339),
		Interval:    300,
		Status:      v201.BootNotificationConfirmationRegistrationStatusEnumAccepted,
	}, nil
}
