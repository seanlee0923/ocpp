// Command basic-auth demonstrates OCPP Security Profile 1.
package main

import (
	"context"
	"crypto/subtle"
	"log"
	"net/http"

	"github.com/seanlee0923/ocpp/csms"
)

func main() {
	server, err := csms.New(csms.Config{
		Security: csms.SecurityConfig{
			Profile:       csms.SecurityProfileBasicAuth,
			Authenticator: csms.AuthenticatorFunc(authenticate),
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(":8080", server))
}

func authenticate(_ context.Context, request csms.AuthenticationRequest) (csms.Principal, error) {
	// Replace this fixed value with a credential store lookup. Do not retain
	// request.Basic.Password; the library clears it after this callback.
	if request.Basic == nil || subtle.ConstantTimeCompare(request.Basic.Password, []byte("change-me")) != 1 {
		return csms.Principal{}, csms.ErrAuthenticationFailed
	}
	return csms.Principal{ID: request.Identity}, nil
}
