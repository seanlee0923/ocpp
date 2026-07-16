// Command duplicate-session demonstrates an explicit replace-existing policy.
package main

import (
	"context"
	"log"
	"net/http"

	"github.com/seanlee0923/ocpp/csms"
)

func main() {
	server, err := csms.New(csms.Config{
		OnDuplicateSession: func(
			_ context.Context,
			attempt csms.DuplicateSessionAttempt,
		) (csms.DuplicateSessionDecision, error) {
			log.Printf("replacing identity=%s connected_at=%s", attempt.Identity, attempt.Existing.ConnectedAt)
			return csms.ReplaceExistingSession, nil
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(":8080", server))
}
