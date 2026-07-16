// Command tls-basic-auth demonstrates OCPP Security Profile 2.
package main

import (
	"context"
	"crypto/subtle"
	"crypto/tls"
	"log"
	"net/http"
	"os"

	"github.com/seanlee0923/ocpp/csms"
)

func main() {
	certFile, keyFile := required("TLS_CERT_FILE"), required("TLS_KEY_FILE")
	ocppServer, err := csms.New(csms.Config{Security: csms.SecurityConfig{
		Profile: csms.SecurityProfileTLSBasicAuth,
		Authenticator: csms.AuthenticatorFunc(func(
			_ context.Context,
			request csms.AuthenticationRequest,
		) (csms.Principal, error) {
			if request.Basic == nil || subtle.ConstantTimeCompare(request.Basic.Password, []byte("change-me")) != 1 {
				return csms.Principal{}, csms.ErrAuthenticationFailed
			}
			return csms.Principal{ID: request.Identity}, nil
		}),
	}})
	if err != nil {
		log.Fatal(err)
	}
	httpServer := &http.Server{
		Addr:      ":8443",
		Handler:   ocppServer,
		TLSConfig: &tls.Config{MinVersion: tls.VersionTLS12},
	}
	log.Fatal(httpServer.ListenAndServeTLS(certFile, keyFile))
}

func required(name string) string {
	value := os.Getenv(name)
	if value == "" {
		log.Fatalf("%s is required", name)
	}
	return value
}
