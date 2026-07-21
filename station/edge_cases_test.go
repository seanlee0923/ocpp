package station_test

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"math"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/seanlee0923/ocpp/csms"
	"github.com/seanlee0923/ocpp/protocol"
	"github.com/seanlee0923/ocpp/station"
	"github.com/seanlee0923/ocpp/v16"
)

func TestNewRejectsInvalidConfig(t *testing.T) {
	base := station.Config{URL: "ws://localhost:8080", Identity: "CP-001", Version: protocol.OCPP16}

	cases := map[string]func(station.Config) station.Config{
		"empty URL":                      func(c station.Config) station.Config { c.URL = ""; return c },
		"empty identity":                 func(c station.Config) station.Config { c.Identity = ""; return c },
		"empty version":                  func(c station.Config) station.Config { c.Version = ""; return c },
		"unsupported version":            func(c station.Config) station.Config { c.Version = protocol.Version("ocpp9.9"); return c },
		"negative CallTimeout":           func(c station.Config) station.Config { c.CallTimeout = -time.Second; return c },
		"negative MaxPendingCalls":       func(c station.Config) station.Config { c.MaxPendingCalls = -1; return c },
		"negative MaxConcurrentHandlers": func(c station.Config) station.Config { c.MaxConcurrentHandlers = -1; return c },
		"negative HandshakeTimeout":      func(c station.Config) station.Config { c.HandshakeTimeout = -time.Second; return c },
		"negative ReadLimit":             func(c station.Config) station.Config { c.ReadLimit = -1; return c },
		"negative reconnect delay": func(c station.Config) station.Config {
			c.ReconnectPolicy = &station.ReconnectPolicy{InitialDelay: -time.Second}
			return c
		},
		"NaN reconnect multiplier": func(c station.Config) station.Config {
			c.ReconnectPolicy = &station.ReconnectPolicy{InitialDelay: time.Second, Multiplier: math.NaN()}
			return c
		},
		"infinite reconnect multiplier": func(c station.Config) station.Config {
			c.ReconnectPolicy = &station.ReconnectPolicy{InitialDelay: time.Second, Multiplier: math.Inf(1)}
			return c
		},
	}
	for name, mutate := range cases {
		t.Run(name, func(t *testing.T) {
			config := mutate(base)
			if _, err := station.New(config); err == nil {
				t.Fatalf("New(%+v) succeeded, want an error", config)
			}
		})
	}

	if _, err := station.New(base); err != nil {
		t.Fatalf("New with a valid config failed: %v", err)
	}
}

func TestStationCallHonorsContextTimeout(t *testing.T) {
	httpServer, block, _ := newBlockingHeartbeatServer(t, nil)
	defer close(block)

	st := dialStation(t, httpServer.URL, "CP-001", nil)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()
	_, err := station.Call[v16.HeartbeatRequest, v16.HeartbeatConfirmation](ctx, st, v16.HeartbeatRequest{})
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("error = %v, want context deadline exceeded", err)
	}
}

// TestStationCallWithCanceledContextDoesNotWriteFrame is the station-side
// mirror of csms.TestCallWithCanceledContextDoesNotWriteFrame: cancellation
// must win before the CSMS can observe or execute the outbound CALL.
func TestStationCallWithCanceledContextDoesNotWriteFrame(t *testing.T) {
	received := make(chan []byte, 1)
	upgrader := websocket.Upgrader{Subprotocols: []string{string(protocol.OCPP16)}}
	httpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close()
		_, data, err := conn.ReadMessage()
		if err == nil {
			received <- data
		}
	}))
	t.Cleanup(httpServer.Close)

	st := dialStation(t, httpServer.URL, "CP-001", nil)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := station.Call[v16.HeartbeatRequest, v16.HeartbeatConfirmation](ctx, st, v16.HeartbeatRequest{})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("Call error = %v, want context.Canceled", err)
	}

	select {
	case data := <-received:
		t.Fatalf("CSMS received frame %s after Call context was already canceled", data)
	case <-time.After(200 * time.Millisecond):
	}
}

func TestStationOutboundPendingCallLimit(t *testing.T) {
	const pendingLimit = 4

	httpServer, block, inFlight := newBlockingHeartbeatServer(t, nil)

	st, err := station.New(station.Config{
		URL: wsURL(httpServer.URL), Identity: "CP-001", Version: protocol.OCPP16, MaxPendingCalls: pendingLimit,
	})
	if err != nil {
		t.Fatal(err)
	}
	ctx := runInBackground(t, st)
	waitConnected(t, st)

	results := make(chan error, pendingLimit)
	for range pendingLimit {
		go func() {
			_, err := station.Call[v16.HeartbeatRequest, v16.HeartbeatConfirmation](ctx, st, v16.HeartbeatRequest{})
			results <- err
		}()
	}
	deadline := time.Now().Add(2 * time.Second)
	for inFlight.Load() != pendingLimit && time.Now().Before(deadline) {
		time.Sleep(time.Millisecond)
	}
	if got := inFlight.Load(); got != pendingLimit {
		t.Fatalf("in-flight Heartbeats = %d, want %d", got, pendingLimit)
	}

	if _, err := station.Call[v16.HeartbeatRequest, v16.HeartbeatConfirmation](ctx, st, v16.HeartbeatRequest{}); !errors.Is(err, station.ErrTooManyPendingCalls) {
		t.Fatalf("overflow error = %v, want ErrTooManyPendingCalls", err)
	}

	close(block)
	for range pendingLimit {
		if err := <-results; err != nil {
			t.Fatal(err)
		}
	}
}

func TestStationStop(t *testing.T) {
	server, err := csms.New(csms.Config{Versions: []protocol.Version{protocol.OCPP16}})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()

	st, err := station.New(station.Config{URL: wsURL(httpServer.URL), Identity: "CP-001", Version: protocol.OCPP16})
	if err != nil {
		t.Fatal(err)
	}
	runDone := make(chan error, 1)
	go func() { runDone <- st.Run(context.Background()) }()
	waitConnected(t, st)

	st.Stop()
	select {
	case err := <-runDone:
		if !errors.Is(err, station.ErrStopped) {
			t.Fatalf("Run after Stop = %v, want ErrStopped", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("Run did not return after Stop")
	}
	if st.State() != station.Disconnected {
		t.Fatalf("state after Stop = %v, want Disconnected", st.State())
	}
	if _, err := station.Call[v16.HeartbeatRequest, v16.HeartbeatConfirmation](context.Background(), st, v16.HeartbeatRequest{}); !errors.Is(err, station.ErrNotConnected) {
		t.Fatalf("Call after Stop = %v, want ErrNotConnected", err)
	}

	// Stop is safe to call again, and safe to call before Run ever ran.
	st.Stop()
	unstarted, err := station.New(station.Config{URL: wsURL(httpServer.URL), Identity: "CP-002", Version: protocol.OCPP16})
	if err != nil {
		t.Fatal(err)
	}
	unstarted.Stop()
}

func TestStationPendingCallFailsOnReconnect(t *testing.T) {
	connected := make(chan *csms.Session, 2)
	httpServer, block, inFlight := newBlockingHeartbeatServer(t, func(session *csms.Session) { connected <- session })
	defer close(block)

	st := dialStation(t, httpServer.URL, "CP-001", nil)
	session := <-connected

	callErr := make(chan error, 1)
	go func() {
		_, err := station.Call[v16.HeartbeatRequest, v16.HeartbeatConfirmation](context.Background(), st, v16.HeartbeatRequest{})
		callErr <- err
	}()
	deadline := time.Now().Add(2 * time.Second)
	for inFlight.Load() == 0 && time.Now().Before(deadline) {
		time.Sleep(time.Millisecond)
	}
	if inFlight.Load() == 0 {
		t.Fatal("Heartbeat never reached the CSMS handler")
	}

	if err := session.Close(); err != nil {
		t.Fatal(err)
	}

	select {
	case err := <-callErr:
		if err == nil {
			t.Fatal("Call succeeded after the underlying connection was closed, want an error")
		}
	case <-time.After(2 * time.Second):
		t.Fatal("pending Call did not fail promptly after disconnect")
	}
}

// newBlockingHeartbeatServer starts a CSMS whose Heartbeat handler blocks
// until block is closed, incrementing inFlight for each Heartbeat currently
// blocked — shared by tests that need a Call to stay pending on the CSMS
// side. onConnect, if non-nil, is passed through to csms.Config.OnConnect.
func newBlockingHeartbeatServer(t *testing.T, onConnect func(*csms.Session)) (httpServer *httptest.Server, block chan struct{}, inFlight *atomic.Int64) {
	t.Helper()
	router := csms.NewRouter()
	block = make(chan struct{})
	inFlight = new(atomic.Int64)
	if err := router.Handle(protocol.OCPP16, "Heartbeat", func(context.Context, *csms.Session, json.RawMessage) (any, error) {
		inFlight.Add(1)
		<-block
		return map[string]any{"currentTime": "2026-07-19T00:00:00Z"}, nil
	}); err != nil {
		t.Fatal(err)
	}
	server, err := csms.New(csms.Config{Router: router, Versions: []protocol.Version{protocol.OCPP16}, OnConnect: onConnect})
	if err != nil {
		t.Fatal(err)
	}
	httpServer = httptest.NewServer(server)
	t.Cleanup(httpServer.Close)
	return httpServer, block, inFlight
}

func TestStationTLS(t *testing.T) {
	server, err := csms.New(csms.Config{Versions: []protocol.Version{protocol.OCPP16}})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewTLSServer(server)
	defer httpServer.Close()

	st, err := station.New(station.Config{
		URL:       "wss" + strings.TrimPrefix(httpServer.URL, "https"),
		Identity:  "CP-001",
		Version:   protocol.OCPP16,
		TLSConfig: &tls.Config{InsecureSkipVerify: true}, //nolint:gosec // test-only, httptest.NewTLSServer's cert is self-signed
	})
	if err != nil {
		t.Fatal(err)
	}
	runInBackground(t, st)
	waitConnected(t, st)
}

func waitConnected(t *testing.T, st *station.Station) {
	t.Helper()
	deadline := time.Now().Add(2 * time.Second)
	for st.State() != station.Connected && time.Now().Before(deadline) {
		time.Sleep(time.Millisecond)
	}
	if st.State() != station.Connected {
		t.Fatal("station did not connect")
	}
}
