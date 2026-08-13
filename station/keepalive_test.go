package station_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/seanlee0923/ocpp/protocol"
	"github.com/seanlee0923/ocpp/station"
	"github.com/seanlee0923/ocpp/v16"
)

// keepaliveServer upgrades one connection and drains it, invoking onPing for
// every inbound ping. Returning true from onPing sends the pong back;
// returning false deliberately stays silent, which is how a test simulates a
// CSMS that has stopped answering. The returned server keeps reading until
// the connection dies, since gorilla only processes control frames from
// inside a read call.
func keepaliveServer(t *testing.T, onPing func() bool, onConn func(*websocket.Conn)) *httptest.Server {
	t.Helper()
	upgrader := websocket.Upgrader{Subprotocols: []string{string(protocol.OCPP16)}}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close()
		if onPing != nil {
			conn.SetPingHandler(func(data string) error {
				if !onPing() {
					return nil
				}
				return conn.WriteControl(websocket.PongMessage, []byte(data), time.Now().Add(time.Second))
			})
		}
		if onConn != nil {
			go onConn(conn)
		}
		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				return
			}
		}
	}))
	t.Cleanup(server.Close)
	return server
}

// TestStationSendsPingsAtPingInterval proves Config.PingInterval makes the
// Station itself originate WebSocket pings, which is what a CSMS that never
// pings its clients (but drops connections it considers silent) requires.
func TestStationSendsPingsAtPingInterval(t *testing.T) {
	var pings atomic.Int64
	server := keepaliveServer(t, func() bool { pings.Add(1); return true }, nil)

	st, err := station.New(station.Config{
		URL: wsURL(server.URL), Identity: "CP-001", Version: protocol.OCPP16,
		PingInterval: 50 * time.Millisecond,
	})
	if err != nil {
		t.Fatal(err)
	}
	runInBackground(t, st)
	waitConnected(t, st)

	deadline := time.Now().Add(3 * time.Second)
	for time.Now().Before(deadline) {
		if pings.Load() >= 3 {
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
	t.Fatalf("station sent %d pings in 3s, want at least 3", pings.Load())
}

// TestStationPongTimeoutEndsConnection proves Config.PongTimeout detects a
// half-open peer: a CSMS that accepted the socket but has stopped answering
// entirely. Without it the Station reads forever against a dead peer.
func TestStationPongTimeoutEndsConnection(t *testing.T) {
	server := keepaliveServer(t, func() bool { return false }, nil)

	st, err := station.New(station.Config{
		URL: wsURL(server.URL), Identity: "CP-001", Version: protocol.OCPP16,
		PingInterval: 50 * time.Millisecond,
		PongTimeout:  200 * time.Millisecond,
	})
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	runErr := st.Run(ctx)
	if !errors.Is(runErr, station.ErrPongTimeout) {
		t.Fatalf("Run error = %v, want %v", runErr, station.ErrPongTimeout)
	}
}

// TestStationPongTimeoutReconnects proves a pong timeout is an ordinary
// connection failure that feeds the existing ReconnectPolicy, rather than a
// terminal error — a station whose CSMS went silent must come back on its
// own once the CSMS returns.
func TestStationPongTimeoutReconnects(t *testing.T) {
	connects := make(chan struct{}, 8)
	server := keepaliveServer(t, func() bool { return false }, nil)

	st, err := station.New(station.Config{
		URL: wsURL(server.URL), Identity: "CP-001", Version: protocol.OCPP16,
		PingInterval:    50 * time.Millisecond,
		PongTimeout:     200 * time.Millisecond,
		ReconnectPolicy: &station.ReconnectPolicy{InitialDelay: 10 * time.Millisecond, MaxDelay: 50 * time.Millisecond, Multiplier: 2},
		OnConnect: func(*station.Station) {
			select {
			case connects <- struct{}{}:
			default:
			}
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	runInBackground(t, st)

	for i := 0; i < 2; i++ {
		select {
		case <-connects:
		case <-time.After(5 * time.Second):
			t.Fatalf("only got %d connections, want the pong timeout to trigger a reconnect", i)
		}
	}
}

// TestStationStillAnswersPingsWithPongTimeoutSet guards the subtlest part of
// enabling PongTimeout: it installs a custom ping handler, which replaces the
// gorilla default that is what actually answers the CSMS's pings. If that
// replacement forgot to send the pong, turning on PongTimeout would silently
// stop the station from answering — breaking every CSMS that relies on
// pong-based liveness, which is the exact failure this feature exists to fix.
func TestStationStillAnswersPingsWithPongTimeoutSet(t *testing.T) {
	var pongs atomic.Int64
	server := keepaliveServer(t, nil, func(conn *websocket.Conn) {
		conn.SetPongHandler(func(string) error { pongs.Add(1); return nil })
		for {
			if err := conn.WriteControl(websocket.PingMessage, nil, time.Now().Add(time.Second)); err != nil {
				return
			}
			time.Sleep(50 * time.Millisecond)
		}
	})

	st, err := station.New(station.Config{
		URL: wsURL(server.URL), Identity: "CP-001", Version: protocol.OCPP16,
		PongTimeout: time.Second,
	})
	if err != nil {
		t.Fatal(err)
	}
	runInBackground(t, st)
	waitConnected(t, st)

	deadline := time.Now().Add(3 * time.Second)
	for time.Now().Before(deadline) {
		if pongs.Load() >= 3 {
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
	t.Fatalf("CSMS received %d pongs in 3s, want at least 3", pongs.Load())
}

// TestStationInboundTrafficHoldsOffPongTimeout proves PongTimeout measures
// silence from the CSMS, not specifically missing pongs: a CSMS that never
// pongs but keeps sending OCPP messages is demonstrably alive, and killing
// that connection would be a false positive. This is the case that makes
// PongTimeout usable on its own, without also enabling PingInterval.
func TestStationInboundTrafficHoldsOffPongTimeout(t *testing.T) {
	server := keepaliveServer(t, func() bool { return false }, func(conn *websocket.Conn) {
		for i := 0; ; i++ {
			frame := `[2,"srv-` + string(rune('a'+i%26)) + `","Reset",{"type":"Soft"}]`
			if err := conn.WriteMessage(websocket.TextMessage, []byte(frame)); err != nil {
				return
			}
			time.Sleep(50 * time.Millisecond)
		}
	})

	st, err := station.New(station.Config{
		URL: wsURL(server.URL), Identity: "CP-001", Version: protocol.OCPP16,
		PongTimeout: 250 * time.Millisecond,
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := station.Handle(st, func(context.Context, v16.ResetRequest) (v16.ResetConfirmation, error) {
		return v16.ResetConfirmation{Status: v16.ResetConfirmationStatusAccepted}, nil
	}); err != nil {
		t.Fatal(err)
	}
	runInBackground(t, st)
	waitConnected(t, st)

	time.Sleep(time.Second)
	if state := st.State(); state != station.Connected {
		t.Fatalf("state = %v after 1s of inbound traffic, want Connected", state)
	}
}
