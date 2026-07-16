// Command multi-version accepts OCPP 1.6, 2.0.1, and 2.1 on one CSMS server.
package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/seanlee0923/ocpp/csms"
	profile16 "github.com/seanlee0923/ocpp/profiles/ocpp16"
	profile201 "github.com/seanlee0923/ocpp/profiles/ocpp201"
	profile21 "github.com/seanlee0923/ocpp/profiles/ocpp21"
	"github.com/seanlee0923/ocpp/v16"
	"github.com/seanlee0923/ocpp/v201"
	"github.com/seanlee0923/ocpp/v21"
)

func main() {
	router := csms.NewRouter()
	p16, err := profile16.NewProfile(router)
	if err != nil {
		log.Fatal(err)
	}
	p201, err := profile201.NewProfile(router)
	if err != nil {
		log.Fatal(err)
	}
	p21, err := profile21.NewProfile(router)
	if err != nil {
		log.Fatal(err)
	}
	if err := p16.HandleBootNotification(boot16); err != nil {
		log.Fatal(err)
	}
	if err := p201.HandleBootNotification(boot201); err != nil {
		log.Fatal(err)
	}
	if err := p21.HandleBootNotification(boot21); err != nil {
		log.Fatal(err)
	}
	server, err := csms.New(csms.Config{Router: router}) // All supported versions.
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(":8080", server))
}

func boot16(context.Context, *csms.Session, v16.BootNotificationRequest) (v16.BootNotificationConfirmation, error) {
	return v16.BootNotificationConfirmation{CurrentTime: now(), Interval: 300, Status: v16.BootNotificationConfirmationStatusAccepted}, nil
}

func boot201(context.Context, *csms.Session, v201.BootNotificationRequest) (v201.BootNotificationConfirmation, error) {
	return v201.BootNotificationConfirmation{CurrentTime: now(), Interval: 300, Status: v201.BootNotificationConfirmationRegistrationStatusEnumAccepted}, nil
}

func boot21(context.Context, *csms.Session, v21.BootNotificationRequest) (v21.BootNotificationConfirmation, error) {
	return v21.BootNotificationConfirmation{CurrentTime: now(), Interval: 300, Status: v21.BootNotificationConfirmationRegistrationStatusEnumAccepted}, nil
}

func now() string { return time.Now().UTC().Format(time.RFC3339) }
