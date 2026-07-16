package csms

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http/httptest"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"ocpp-go/protocol"
	"ocpp-go/v16"
)

func TestBoundedConcurrentInboundCallLoad(t *testing.T) {
	const (
		stationCount    = 16
		callsPerStation = 50
	)
	var handled atomic.Int64
	router := NewRouter()
	if err := router.Handle(protocol.OCPP16, "Heartbeat", func(context.Context, *Session, json.RawMessage) (any, error) {
		handled.Add(1)
		return map[string]any{"currentTime": "2026-07-16T00:00:00Z"}, nil
	}); err != nil {
		t.Fatal(err)
	}
	server, err := New(Config{
		Router: router, Versions: []protocol.Version{protocol.OCPP16}, MaxConcurrentHandlers: 8,
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()

	errorsCh := make(chan error, stationCount)
	var wg sync.WaitGroup
	for station := 0; station < stationCount; station++ {
		wg.Add(1)
		go func(station int) {
			defer wg.Done()
			dialer := websocket.Dialer{Subprotocols: []string{string(protocol.OCPP16)}}
			conn, _, err := dialer.Dial(
				"ws"+strings.TrimPrefix(httpServer.URL, "http")+fmt.Sprintf("/LOAD-%02d", station), nil,
			)
			if err != nil {
				errorsCh <- err
				return
			}
			defer conn.Close()
			for call := 0; call < callsPerStation; call++ {
				id := fmt.Sprintf("%02d-%03d", station, call)
				if err := conn.WriteMessage(websocket.TextMessage, []byte(`[2,"`+id+`","Heartbeat",{}]`)); err != nil {
					errorsCh <- err
					return
				}
				_, data, err := conn.ReadMessage()
				if err != nil {
					errorsCh <- err
					return
				}
				message, err := protocol.Decode(data)
				if err != nil {
					errorsCh <- err
					return
				}
				result, ok := message.(protocol.CallResult)
				if !ok || result.ID != id {
					errorsCh <- fmt.Errorf("response = %#v, want CALLRESULT %s", message, id)
					return
				}
			}
		}(station)
	}
	wg.Wait()
	close(errorsCh)
	for err := range errorsCh {
		t.Error(err)
	}
	want := int64(stationCount * callsPerStation)
	if got := handled.Load(); got != want {
		t.Fatalf("handled calls = %d, want %d", got, want)
	}
}

func TestOutboundPendingCallLimitUnderConcurrency(t *testing.T) {
	const pendingLimit = 4
	connected := make(chan *Session, 1)
	var sequence atomic.Int64
	server, err := New(Config{
		Versions:        []protocol.Version{protocol.OCPP16},
		MaxPendingCalls: pendingLimit,
		UniqueIDGenerator: func() string {
			return fmt.Sprintf("load-%d", sequence.Add(1))
		},
		OnConnect: func(session *Session) { connected <- session },
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()
	conn := dialTestStation(t, httpServer.URL, protocol.OCPP16)
	defer conn.Close()
	session := <-connected

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	results := make(chan error, pendingLimit)
	for range pendingLimit {
		go func() {
			_, err := Call[v16.ResetRequest, v16.ResetConfirmation](ctx, session, v16.ResetRequest{Type: v16.ResetRequestTypeSoft})
			results <- err
		}()
	}
	deadline := time.Now().Add(time.Second)
	for pendingCallCount(session) != pendingLimit && time.Now().Before(deadline) {
		time.Sleep(time.Millisecond)
	}
	if got := pendingCallCount(session); got != pendingLimit {
		t.Fatalf("pending calls = %d, want %d", got, pendingLimit)
	}

	_, err = Call[v16.ResetRequest, v16.ResetConfirmation](context.Background(), session, v16.ResetRequest{Type: v16.ResetRequestTypeSoft})
	if !errors.Is(err, ErrTooManyPendingCalls) {
		t.Fatalf("overflow error = %v, want ErrTooManyPendingCalls", err)
	}
	cancel()
	for range pendingLimit {
		if err := <-results; !errors.Is(err, context.Canceled) {
			t.Fatalf("pending call error = %v, want context canceled", err)
		}
	}
	if got := pendingCallCount(session); got != 0 {
		t.Fatalf("pending calls after cancellation = %d", got)
	}
}

func TestSessionSoak(t *testing.T) {
	durationText := os.Getenv("OCPP_SOAK_DURATION")
	if durationText == "" {
		t.Skip("set OCPP_SOAK_DURATION (for example 30m) to run the opt-in soak test")
	}
	duration, err := time.ParseDuration(durationText)
	if err != nil || duration <= 0 {
		t.Fatalf("invalid OCPP_SOAK_DURATION %q", durationText)
	}
	const stationCount = 8
	var handled atomic.Int64
	router := NewRouter()
	if err := router.Handle(protocol.OCPP16, "Heartbeat", func(context.Context, *Session, json.RawMessage) (any, error) {
		handled.Add(1)
		return map[string]any{"currentTime": time.Now().UTC().Format(time.RFC3339)}, nil
	}); err != nil {
		t.Fatal(err)
	}
	server, err := New(Config{Router: router, Versions: []protocol.Version{protocol.OCPP16}})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()

	stopAt := time.Now().Add(duration)
	errorsCh := make(chan error, stationCount)
	var wg sync.WaitGroup
	for station := 0; station < stationCount; station++ {
		wg.Add(1)
		go func(station int) {
			defer wg.Done()
			dialer := websocket.Dialer{Subprotocols: []string{string(protocol.OCPP16)}}
			conn, _, err := dialer.Dial(
				"ws"+strings.TrimPrefix(httpServer.URL, "http")+fmt.Sprintf("/SOAK-%02d", station), nil,
			)
			if err != nil {
				errorsCh <- err
				return
			}
			defer conn.Close()
			for call := 0; time.Now().Before(stopAt); call++ {
				id := fmt.Sprintf("%02d-%d", station, call)
				if err := conn.WriteMessage(websocket.TextMessage, []byte(`[2,"`+id+`","Heartbeat",{}]`)); err != nil {
					errorsCh <- err
					return
				}
				_, data, err := conn.ReadMessage()
				if err != nil {
					errorsCh <- err
					return
				}
				message, err := protocol.Decode(data)
				if err != nil {
					errorsCh <- err
					return
				}
				if result, ok := message.(protocol.CallResult); !ok || result.ID != id {
					errorsCh <- fmt.Errorf("response = %#v, want CALLRESULT %s", message, id)
					return
				}
			}
		}(station)
	}
	wg.Wait()
	close(errorsCh)
	for err := range errorsCh {
		t.Error(err)
	}
	if handled.Load() == 0 {
		t.Fatal("soak test handled no calls")
	}
	t.Logf("handled %d calls across %d sessions in %s", handled.Load(), stationCount, duration)
}

func pendingCallCount(session *Session) int {
	session.pendingMu.Lock()
	defer session.pendingMu.Unlock()
	return len(session.pending)
}
