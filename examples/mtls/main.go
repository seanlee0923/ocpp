// Command mtls demonstrates OCPP Security Profile 3.
package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/seanlee0923/ocpp/csms"
)

func main() {
	certFile, keyFile := required("TLS_CERT_FILE"), required("TLS_KEY_FILE")
	clientCAs := loadClientCAs(required("TLS_CLIENT_CA_FILE"))
	ocppServer, err := csms.New(csms.Config{Security: csms.SecurityConfig{
		Profile: csms.SecurityProfileTLSClientCertificate,
		Authenticator: csms.AuthenticatorFunc(func(
			_ context.Context,
			request csms.AuthenticationRequest,
		) (csms.Principal, error) {
			if request.TLS == nil || len(request.TLS.PeerCertificates) == 0 {
				return csms.Principal{}, csms.ErrAuthenticationFailed
			}
			certificate := request.TLS.PeerCertificates[0]
			if certificate.Subject.CommonName != request.Identity {
				return csms.Principal{}, csms.ErrAuthenticationFailed
			}
			return csms.Principal{ID: certificate.Subject.CommonName}, nil
		}),
	}})
	if err != nil {
		log.Fatal(err)
	}
	httpServer := &http.Server{
		Addr:    ":8443",
		Handler: ocppServer,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
			ClientAuth: tls.RequireAndVerifyClientCert,
			ClientCAs:  clientCAs,
		},
	}
	log.Fatal(httpServer.ListenAndServeTLS(certFile, keyFile))
}

func loadClientCAs(file string) *x509.CertPool {
	pem, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	pool := x509.NewCertPool()
	if !pool.AppendCertsFromPEM(pem) {
		log.Fatal(errors.New("TLS_CLIENT_CA_FILE contains no certificates"))
	}
	return pool
}

func required(name string) string {
	value := os.Getenv(name)
	if value == "" {
		log.Fatalf("%s is required", name)
	}
	return value
}
