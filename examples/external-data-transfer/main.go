// Command external-data-transfer demonstrates triggering an OCPP 1.6
// DataTransfer CALL from an HTTP API outside the OCPP handler flow.
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/seanlee0923/ocpp/csms"
	"github.com/seanlee0923/ocpp/profiles/ocpp16"
	"github.com/seanlee0923/ocpp/protocol"
	"github.com/seanlee0923/ocpp/v16"
)

const callTimeout = 10 * time.Second

type application struct {
	server  *csms.Server
	profile *ocpp16.Profile
}

type dataTransferRequest struct {
	VendorID  string  `json:"vendorId"`
	MessageID *string `json:"messageId,omitempty"`
	Data      *string `json:"data,omitempty"`
}

func main() {
	router := csms.NewRouter()
	profile, err := ocpp16.NewProfile(router)
	if err != nil {
		log.Fatal(err)
	}
	if err := profile.HandleBootNotification(acceptBoot); err != nil {
		log.Fatal(err)
	}

	server, err := csms.New(csms.Config{
		Router:   router,
		Versions: []protocol.Version{protocol.OCPP16},
	})
	if err != nil {
		log.Fatal(err)
	}
	app := &application{server: server, profile: profile}

	ocppHTTP := &http.Server{
		Addr:              ":8080",
		Handler:           server,
		ReadHeaderTimeout: 5 * time.Second,
	}
	apiMux := http.NewServeMux()
	apiMux.HandleFunc("POST /stations/{cpid}/data-transfer", app.dataTransfer)
	apiHTTP := &http.Server{
		Addr:              ":8081",
		Handler:           apiMux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	errCh := make(chan error, 2)
	go serve(errCh, "OCPP WebSocket", ocppHTTP)
	go serve(errCh, "external API", apiHTTP)
	log.Printf("OCPP endpoint: ws://localhost:8080/{cpid}")
	log.Printf("external API: POST http://localhost:8081/stations/{cpid}/data-transfer")

	signalCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	select {
	case <-signalCtx.Done():
	case err := <-errCh:
		log.Printf("server stopped: %v", err)
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("OCPP shutdown: %v", err)
	}
	if err := ocppHTTP.Shutdown(shutdownCtx); err != nil {
		log.Printf("OCPP HTTP shutdown: %v", err)
	}
	if err := apiHTTP.Shutdown(shutdownCtx); err != nil {
		log.Printf("API HTTP shutdown: %v", err)
	}
}

func serve(errCh chan<- error, name string, server *http.Server) {
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		errCh <- fmt.Errorf("%s server: %w", name, err)
	}
}

func acceptBoot(
	context.Context,
	*csms.Session,
	v16.BootNotificationRequest,
) (v16.BootNotificationConfirmation, error) {
	return v16.BootNotificationConfirmation{
		CurrentTime: time.Now().UTC().Format(time.RFC3339),
		Interval:    300,
		Status:      v16.BootNotificationConfirmationStatusAccepted,
	}, nil
}

func (app *application) dataTransfer(w http.ResponseWriter, r *http.Request) {
	var request dataTransferRequest
	decoder := json.NewDecoder(http.MaxBytesReader(w, r.Body, 1<<20))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&request); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if request.VendorID == "" {
		writeError(w, http.StatusBadRequest, "vendorId is required")
		return
	}

	cpid := r.PathValue("cpid")
	session, ok := app.server.Session(cpid)
	if !ok {
		writeError(w, http.StatusNotFound, "charging station is not connected")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), callTimeout)
	defer cancel()
	confirmation, err := app.profile.CallDataTransfer(ctx, session, v16.DataTransferRequest{
		VendorID: request.VendorID, MessageID: request.MessageID, Data: request.Data,
	})
	if err != nil {
		writeCallError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, confirmation)
}

func writeCallError(w http.ResponseWriter, err error) {
	status := http.StatusBadGateway
	switch {
	case errors.Is(err, context.DeadlineExceeded):
		status = http.StatusGatewayTimeout
	case errors.Is(err, context.Canceled):
		status = http.StatusRequestTimeout
	case errors.Is(err, ocpp16.ErrNotBooted),
		errors.Is(err, csms.ErrSessionClosed),
		errors.Is(err, csms.ErrSessionReplaced):
		status = http.StatusConflict
	case errors.Is(err, csms.ErrTooManyPendingCalls):
		status = http.StatusTooManyRequests
	}
	writeError(w, status, err.Error())
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(value); err != nil {
		log.Printf("encode HTTP response: %v", err)
	}
}
