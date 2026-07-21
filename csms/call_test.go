package csms

import (
	"context"
	"errors"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/seanlee0923/ocpp/internal/stationtest"
	"github.com/seanlee0923/ocpp/protocol"
	"github.com/seanlee0923/ocpp/v16"
)

func TestCallRoundTripUsesConfiguredUniqueID(t *testing.T) {
	connected := make(chan *Session, 1)
	server, err := New(Config{
		Versions:          []protocol.Version{protocol.OCPP16},
		UniqueIDGenerator: func() string { return "custom-call-id" },
		OnConnect:         func(session *Session) { connected <- session },
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()
	conn := dialTestStation(t, httpServer.URL, protocol.OCPP16)
	defer conn.Close()
	session := <-connected

	stationErr := make(chan error, 1)
	go func() {
		_, data, err := conn.ReadMessage()
		if err != nil {
			stationErr <- err
			return
		}
		message, err := protocol.Decode(data)
		if err != nil {
			stationErr <- err
			return
		}
		call, ok := message.(protocol.Call)
		if !ok || call.ID != "custom-call-id" || call.Action != "Reset" {
			stationErr <- errors.New("unexpected outbound CALL")
			return
		}
		stationErr <- conn.WriteMessage(websocket.TextMessage, []byte(`[3,"custom-call-id",{"status":"Accepted"}]`))
	}()

	confirmation, err := Call[v16.ResetRequest, v16.ResetConfirmation](
		context.Background(), session, v16.ResetRequest{Type: v16.ResetRequestTypeSoft},
	)
	if err != nil {
		t.Fatal(err)
	}
	if confirmation.Status != v16.ResetConfirmationStatusAccepted {
		t.Fatalf("status = %q", confirmation.Status)
	}
	if err := <-stationErr; err != nil {
		t.Fatal(err)
	}
}

func TestCallReturnsRemoteCallError(t *testing.T) {
	connected := make(chan *Session, 1)
	server, err := New(Config{
		Versions:  []protocol.Version{protocol.OCPP16},
		OnConnect: func(session *Session) { connected <- session },
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()
	station := stationtest.Dial(t, httpServer.URL, "CP-001", protocol.OCPP16)
	if station == nil {
		return
	}
	handleReset(t, station, func(context.Context, v16.ResetRequest) (v16.ResetConfirmation, error) {
		return v16.ResetConfirmation{}, &stationtest.CallError{Code: "NotSupported", Description: "cannot reset"}
	})
	session := <-connected

	_, err = Call[v16.ResetRequest, v16.ResetConfirmation](
		context.Background(), session, v16.ResetRequest{Type: v16.ResetRequestTypeSoft},
	)
	var remote *RemoteCallError
	if !errors.As(err, &remote) || remote.Code != "NotSupported" {
		t.Fatalf("error = %#v, want remote NotSupported CALLERROR", err)
	}
}

func TestCallHonorsContextTimeout(t *testing.T) {
	connected := make(chan *Session, 1)
	server, err := New(Config{
		Versions:  []protocol.Version{protocol.OCPP16},
		OnConnect: func(session *Session) { connected <- session },
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()
	station := stationtest.Dial(t, httpServer.URL, "CP-001", protocol.OCPP16)
	if station == nil {
		return
	}
	handleReset(t, station, func(context.Context, v16.ResetRequest) (v16.ResetConfirmation, error) {
		time.Sleep(200 * time.Millisecond)
		return v16.ResetConfirmation{Status: v16.ResetConfirmationStatusAccepted}, nil
	})
	session := <-connected

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()
	_, err = Call[v16.ResetRequest, v16.ResetConfirmation](
		ctx, session, v16.ResetRequest{Type: v16.ResetRequestTypeSoft},
	)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("error = %v, want context deadline exceeded", err)
	}
}

// TestCallWithCanceledContextDoesNotWriteFrame proves an already-canceled
// context stops a CALL before any wire side effect. Returning
// context.Canceled after transmitting the frame would let the remote station
// execute a non-idempotent action while the caller believes it never started.
func TestCallWithCanceledContextDoesNotWriteFrame(t *testing.T) {
	connected := make(chan *Session, 1)
	server, err := New(Config{
		Versions:  []protocol.Version{protocol.OCPP16},
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
	cancel()
	_, err = Call[v16.ResetRequest, v16.ResetConfirmation](
		ctx, session, v16.ResetRequest{Type: v16.ResetRequestTypeSoft},
	)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("Call error = %v, want context.Canceled", err)
	}

	if err := conn.SetReadDeadline(time.Now().Add(200 * time.Millisecond)); err != nil {
		t.Fatal(err)
	}
	if _, data, err := conn.ReadMessage(); err == nil {
		t.Fatalf("peer received frame %s after Call context was already canceled", data)
	}
}

// TestSendRechecksContextAfterWaitingForWriteLock covers cancellation that
// happens after send's initial ctx check but while another writer owns the
// WebSocket write lock. The second check must return before touching conn.
func TestSendRechecksContextAfterWaitingForWriteLock(t *testing.T) {
	base, cancel := context.WithCancel(context.Background())
	ctx := &firstErrObservedContext{Context: base, observed: make(chan struct{})}
	session := &Session{}
	session.writeMu.Lock()

	result := make(chan error, 1)
	go func() {
		result <- session.send(ctx, protocol.Call{ID: "1", Action: "Reset", Payload: []byte(`{"type":"Soft"}`)})
	}()

	<-ctx.observed
	cancel()
	session.writeMu.Unlock()
	if err := <-result; !errors.Is(err, context.Canceled) {
		t.Fatalf("send error = %v, want context.Canceled", err)
	}
}

type firstErrObservedContext struct {
	context.Context
	once     sync.Once
	observed chan struct{}
}

func (ctx *firstErrObservedContext) Err() error {
	err := ctx.Context.Err()
	ctx.once.Do(func() { close(ctx.observed) })
	return err
}

func TestCallConvertsUniqueIDGeneratorPanicToError(t *testing.T) {
	connected := make(chan *Session, 1)
	server, err := New(Config{
		Versions: []protocol.Version{protocol.OCPP16},
		UniqueIDGenerator: func() string {
			panic("generator failed")
		},
		OnConnect: func(session *Session) { connected <- session },
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()
	stationtest.Dial(t, httpServer.URL, "CP-001", protocol.OCPP16)
	_, err = Call[v16.ResetRequest, v16.ResetConfirmation](
		context.Background(), <-connected, v16.ResetRequest{Type: v16.ResetRequestTypeSoft},
	)
	if !errors.Is(err, ErrUniqueIDGeneration) {
		t.Fatalf("error = %v, want ErrUniqueIDGeneration", err)
	}
}

func handleReset(t *testing.T, station *stationtest.Station, handler func(context.Context, v16.ResetRequest) (v16.ResetConfirmation, error)) {
	t.Helper()
	if err := stationtest.Handle(station, handler); err != nil {
		t.Fatal(err)
	}
}

func TestGenerateUniqueIDRejectsInvalidLength(t *testing.T) {
	for _, id := range []string{"", strings.Repeat("x", protocol.MaxUniqueIDLength+1)} {
		if _, err := generateUniqueID(func() string { return id }); !errors.Is(err, ErrUniqueIDGeneration) {
			t.Fatalf("ID length %d error = %v", len(id), err)
		}
	}
}

func dialTestStation(t *testing.T, serverURL string, version protocol.Version) *websocket.Conn {
	t.Helper()
	dialer := websocket.Dialer{Subprotocols: []string{string(version)}}
	conn, _, err := dialer.Dial("ws"+strings.TrimPrefix(serverURL, "http")+"/CP-001", nil)
	if err != nil {
		t.Fatal(err)
	}
	return conn
}
