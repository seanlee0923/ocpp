package csms

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"ocpp-go/protocol"
)

func TestSessionPingPongUpdatesInfo(t *testing.T) {
	server, err := New(Config{
		Versions:     []protocol.Version{protocol.OCPP16},
		PingInterval: 10 * time.Millisecond,
		PongTimeout:  200 * time.Millisecond,
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()
	conn := dialTestStation(t, httpServer.URL, protocol.OCPP16)
	defer conn.Close()
	stationDone := readStationUntilClosed(conn)
	session, ok := server.Session("CP-001")
	if !ok {
		t.Fatal("session not registered")
	}

	deadline := time.Now().Add(time.Second)
	for session.Info().LastPongAt.IsZero() && time.Now().Before(deadline) {
		time.Sleep(5 * time.Millisecond)
	}
	info := session.Info()
	if info.LastPongAt.IsZero() {
		t.Fatal("pong was not observed")
	}
	if info.State != SessionActive || info.LastReceivedAt.IsZero() || info.ConnectedAt.IsZero() {
		t.Fatalf("unexpected session info: %#v", info)
	}
	_ = session.Close()
	<-stationDone
}

func TestIdleTimeoutIgnoresPongFrames(t *testing.T) {
	server, err := New(Config{
		Versions:     []protocol.Version{protocol.OCPP16},
		PingInterval: 10 * time.Millisecond,
		PongTimeout:  200 * time.Millisecond,
		IdleTimeout:  40 * time.Millisecond,
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()
	conn := dialTestStation(t, httpServer.URL, protocol.OCPP16)
	defer conn.Close()
	stationDone := readStationUntilClosed(conn)
	session, _ := server.Session("CP-001")
	select {
	case <-session.Done():
	case <-time.After(time.Second):
		t.Fatal("idle session was not closed")
	}
	if !errors.Is(session.Err(), ErrIdleTimeout) {
		t.Fatalf("session error = %v, want idle timeout", session.Err())
	}
	<-stationDone
}

func TestShutdownClosesSessionsAndRejectsConnections(t *testing.T) {
	server, err := New(Config{Versions: []protocol.Version{protocol.OCPP16}})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()
	conn := dialTestStation(t, httpServer.URL, protocol.OCPP16)
	defer conn.Close()
	session, _ := server.Session("CP-001")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		t.Fatal(err)
	}
	if !errors.Is(session.Err(), ErrServerShutdown) {
		t.Fatalf("session error = %v, want server shutdown", session.Err())
	}
	if session.Info().State != SessionClosed || server.SessionCount() != 0 {
		t.Fatalf("state = %d, session count = %d", session.Info().State, server.SessionCount())
	}

	dialer := websocket.Dialer{Subprotocols: []string{string(protocol.OCPP16)}}
	newConn, response, err := dialer.Dial("ws"+strings.TrimPrefix(httpServer.URL, "http")+"/CP-002", nil)
	if newConn != nil {
		newConn.Close()
	}
	if err == nil || response == nil || response.StatusCode != http.StatusServiceUnavailable {
		t.Fatalf("dial error = %v, response = %#v, want HTTP 503", err, response)
	}
}

func TestNewRejectsInvalidLifecycleTimeouts(t *testing.T) {
	if _, err := New(Config{PingInterval: time.Second, PongTimeout: time.Second}); err == nil {
		t.Fatal("New succeeded when pong timeout equals ping interval")
	}
	if _, err := New(Config{IdleTimeout: -time.Second}); err == nil {
		t.Fatal("New succeeded with negative idle timeout")
	}
}

func TestAbruptDisconnectAllowsSameIdentityToReconnect(t *testing.T) {
	server, err := New(Config{Versions: []protocol.Version{protocol.OCPP201}})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()
	first := dialTestStation(t, httpServer.URL, protocol.OCPP201)
	firstSession, ok := server.Session("CP-001")
	if !ok {
		t.Fatal("first session not registered")
	}
	if err := first.UnderlyingConn().Close(); err != nil {
		t.Fatal(err)
	}
	select {
	case <-firstSession.Done():
	case <-time.After(time.Second):
		t.Fatal("abruptly disconnected session was not closed")
	}
	deadline := time.Now().Add(time.Second)
	for server.SessionCount() != 0 && time.Now().Before(deadline) {
		time.Sleep(time.Millisecond)
	}
	if server.SessionCount() != 0 {
		t.Fatal("abruptly disconnected session was not unregistered")
	}

	second := dialTestStation(t, httpServer.URL, protocol.OCPP201)
	defer second.Close()
	secondSession, ok := server.Session("CP-001")
	if !ok || secondSession == firstSession {
		t.Fatal("same identity did not reconnect with a new session")
	}
}

func TestConcurrentSessionLifecycle(t *testing.T) {
	const stationCount = 32
	server, err := New(Config{Versions: []protocol.Version{protocol.OCPP21}})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()

	type dialResult struct {
		conn *websocket.Conn
		err  error
	}
	results := make(chan dialResult, stationCount)
	var wg sync.WaitGroup
	for i := 0; i < stationCount; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			dialer := websocket.Dialer{Subprotocols: []string{string(protocol.OCPP21)}}
			conn, _, err := dialer.Dial(
				"ws"+strings.TrimPrefix(httpServer.URL, "http")+fmt.Sprintf("/CP-%03d", index), nil,
			)
			results <- dialResult{conn: conn, err: err}
		}(i)
	}
	wg.Wait()
	close(results)
	connections := make([]*websocket.Conn, 0, stationCount)
	for result := range results {
		if result.err != nil {
			t.Fatal(result.err)
		}
		connections = append(connections, result.conn)
	}
	if server.SessionCount() != stationCount {
		t.Fatalf("session count = %d, want %d", server.SessionCount(), stationCount)
	}

	for _, conn := range connections {
		wg.Add(1)
		go func(conn *websocket.Conn) {
			defer wg.Done()
			_ = conn.Close()
		}(conn)
	}
	wg.Wait()
	deadline := time.Now().Add(2 * time.Second)
	for server.SessionCount() != 0 && time.Now().Before(deadline) {
		time.Sleep(time.Millisecond)
	}
	if server.SessionCount() != 0 {
		t.Fatalf("session count after close = %d", server.SessionCount())
	}
}

func readStationUntilClosed(conn *websocket.Conn) <-chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				return
			}
		}
	}()
	return done
}
