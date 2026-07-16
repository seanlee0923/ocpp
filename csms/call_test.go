package csms

import (
	"context"
	"errors"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
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
		Versions:          []protocol.Version{protocol.OCPP16},
		UniqueIDGenerator: func() string { return "rejected-call" },
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

	go func() {
		_, _, _ = conn.ReadMessage()
		_ = conn.WriteMessage(websocket.TextMessage, []byte(`[4,"rejected-call","NotSupported","cannot reset",{"reason":"busy"}]`))
	}()
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
	conn := dialTestStation(t, httpServer.URL, protocol.OCPP16)
	defer conn.Close()
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
	conn := dialTestStation(t, httpServer.URL, protocol.OCPP16)
	defer conn.Close()
	_, err = Call[v16.ResetRequest, v16.ResetConfirmation](
		context.Background(), <-connected, v16.ResetRequest{Type: v16.ResetRequestTypeSoft},
	)
	if !errors.Is(err, ErrUniqueIDGeneration) {
		t.Fatalf("error = %v, want ErrUniqueIDGeneration", err)
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
