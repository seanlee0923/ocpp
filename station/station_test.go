package station_test

import (
	"context"
	"errors"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/seanlee0923/ocpp/csms"
	"github.com/seanlee0923/ocpp/protocol"
	"github.com/seanlee0923/ocpp/station"
	"github.com/seanlee0923/ocpp/v16"
)

func TestStationCallRoundTrip(t *testing.T) {
	router := csms.NewRouter()
	if err := csms.Handle(router, func(_ context.Context, _ *csms.Session, req v16.BootNotificationRequest) (v16.BootNotificationConfirmation, error) {
		if req.ChargePointVendor == "" {
			t.Errorf("BootNotification vendor is empty")
		}
		return v16.BootNotificationConfirmation{
			Status: v16.BootNotificationConfirmationStatusAccepted, CurrentTime: "2026-07-19T00:00:00Z", Interval: 300,
		}, nil
	}); err != nil {
		t.Fatal(err)
	}
	server, err := csms.New(csms.Config{Router: router, Versions: []protocol.Version{protocol.OCPP16}})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()

	st := dialStation(t, httpServer.URL, "CP-001", nil)
	confirmation, err := station.Call[v16.BootNotificationRequest, v16.BootNotificationConfirmation](
		context.Background(), st, v16.BootNotificationRequest{ChargePointVendor: "Acme", ChargePointModel: "X1"},
	)
	if err != nil {
		t.Fatal(err)
	}
	if confirmation.Status != v16.BootNotificationConfirmationStatusAccepted {
		t.Fatalf("status = %q", confirmation.Status)
	}
}

func TestStationHandleInboundCall(t *testing.T) {
	connected := make(chan *csms.Session, 1)
	server, err := csms.New(csms.Config{
		Versions:  []protocol.Version{protocol.OCPP16},
		OnConnect: func(session *csms.Session) { connected <- session },
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()

	st := dialStation(t, httpServer.URL, "CP-001", nil)
	if err := station.Handle(st, func(_ context.Context, req v16.ResetRequest) (v16.ResetConfirmation, error) {
		if req.Type != v16.ResetRequestTypeSoft {
			t.Errorf("Reset type = %q, want Soft", req.Type)
		}
		return v16.ResetConfirmation{Status: v16.ResetConfirmationStatusAccepted}, nil
	}); err != nil {
		t.Fatal(err)
	}

	session := <-connected
	confirmation, err := csms.Call[v16.ResetRequest, v16.ResetConfirmation](
		context.Background(), session, v16.ResetRequest{Type: v16.ResetRequestTypeSoft},
	)
	if err != nil {
		t.Fatal(err)
	}
	if confirmation.Status != v16.ResetConfirmationStatusAccepted {
		t.Fatalf("status = %q", confirmation.Status)
	}
}

func TestStationReconnectsAfterDisconnect(t *testing.T) {
	connected := make(chan *csms.Session, 4)
	server, err := csms.New(csms.Config{
		Versions:  []protocol.Version{protocol.OCPP16},
		OnConnect: func(session *csms.Session) { connected <- session },
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()

	st := dialStation(t, httpServer.URL, "CP-001", &station.ReconnectPolicy{
		InitialDelay: 5 * time.Millisecond, MaxDelay: 20 * time.Millisecond, Multiplier: 2,
	})
	if err := station.Handle(st, func(_ context.Context, req v16.ResetRequest) (v16.ResetConfirmation, error) {
		return v16.ResetConfirmation{Status: v16.ResetConfirmationStatusAccepted}, nil
	}); err != nil {
		t.Fatal(err)
	}

	first := <-connected
	if err := first.Close(); err != nil {
		t.Fatal(err)
	}

	waitConnected(t, st)

	select {
	case second := <-connected:
		confirmation, err := csms.Call[v16.ResetRequest, v16.ResetConfirmation](
			context.Background(), second, v16.ResetRequest{Type: v16.ResetRequestTypeSoft},
		)
		if err != nil {
			t.Fatal(err)
		}
		if confirmation.Status != v16.ResetConfirmationStatusAccepted {
			t.Fatalf("status after reconnect = %q", confirmation.Status)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("CSMS did not observe a second connection after reconnect")
	}
}

func TestStationBasicAuth(t *testing.T) {
	authenticator := csms.AuthenticatorFunc(func(_ context.Context, request csms.AuthenticationRequest) (csms.Principal, error) {
		if request.Basic == nil || request.Basic.Username != "CP-AUTH" || string(request.Basic.Password) != "s3cret" {
			return csms.Principal{}, csms.ErrAuthenticationFailed
		}
		return csms.Principal{ID: request.Identity}, nil
	})
	server, err := csms.New(csms.Config{
		Versions: []protocol.Version{protocol.OCPP16},
		Security: csms.SecurityConfig{Profile: csms.SecurityProfileBasicAuth, Authenticator: authenticator},
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()

	connected := make(chan struct{})
	ok, err := station.New(station.Config{
		URL: wsURL(httpServer.URL), Identity: "CP-AUTH", Version: protocol.OCPP16,
		BasicAuth: &station.BasicCredentials{Username: "CP-AUTH", Password: "s3cret"},
		OnConnect: func(*station.Station) { close(connected) },
	})
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	runDone := make(chan error, 1)
	go func() { runDone <- ok.Run(ctx) }()
	select {
	case <-connected:
	case err := <-runDone:
		t.Fatalf("Run with correct credentials exited before connecting: %v", err)
	}
	if err := <-runDone; !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("Run with correct credentials = %v, want the connection to stay up until ctx times out", err)
	}

	bad, err := station.New(station.Config{
		URL: wsURL(httpServer.URL), Identity: "CP-AUTH", Version: protocol.OCPP16,
		BasicAuth: &station.BasicCredentials{Username: "CP-AUTH", Password: "wrong"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := bad.Run(context.Background()); err == nil {
		t.Fatal("Run with wrong credentials succeeded, want a dial error")
	}
}

func wsURL(serverURL string) string {
	return "ws" + strings.TrimPrefix(serverURL, "http")
}

func dialStation(t *testing.T, serverURL, identity string, reconnect *station.ReconnectPolicy) *station.Station {
	t.Helper()
	st, err := station.New(station.Config{
		URL: wsURL(serverURL), Identity: identity, Version: protocol.OCPP16, ReconnectPolicy: reconnect,
	})
	if err != nil {
		t.Fatal(err)
	}
	runInBackground(t, st)
	waitConnected(t, st)
	return st
}

// runInBackground starts st.Run in its own goroutine and registers a
// t.Cleanup that cancels it and waits for Run to actually return before the
// test finishes — a bare `go func() { _ = st.Run(ctx) }()` with no wait
// leaves that goroutine (and anything it's still doing, e.g. tearing down a
// connection) running past the end of the test, racing with later tests'
// state instead of being deterministically gone. Returns the ctx passed to
// Run, for tests that also need it for their own station.Call calls.
func runInBackground(t *testing.T, st *station.Station) context.Context {
	t.Helper()
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	go func() {
		defer close(done)
		_ = st.Run(ctx)
	}()
	t.Cleanup(func() {
		cancel()
		<-done
	})
	return ctx
}
