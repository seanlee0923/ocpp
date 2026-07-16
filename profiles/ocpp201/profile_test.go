package ocpp201

import (
	"context"
	"encoding/json"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"ocpp-go/csms"
	"ocpp-go/protocol"
	"ocpp-go/v201"
)

func TestRegistrationGateAndResetRoundTrip(t *testing.T) {
	router := csms.NewRouter()
	profile, err := NewProfile(router)
	if err != nil {
		t.Fatal(err)
	}
	var heartbeats atomic.Int32
	if err := profile.HandleBootNotification(func(context.Context, *csms.Session, v201.BootNotificationRequest) (v201.BootNotificationConfirmation, error) {
		return v201.BootNotificationConfirmation{CurrentTime: "2026-07-16T00:00:00Z", Interval: 300, Status: v201.BootNotificationConfirmationRegistrationStatusEnumAccepted}, nil
	}); err != nil {
		t.Fatal(err)
	}
	if err := profile.HandleHeartbeat(func(context.Context, *csms.Session, v201.HeartbeatRequest) (v201.HeartbeatConfirmation, error) {
		heartbeats.Add(1)
		return v201.HeartbeatConfirmation{CurrentTime: "2026-07-16T00:00:01Z"}, nil
	}); err != nil {
		t.Fatal(err)
	}

	server, err := csms.New(csms.Config{Router: router, Versions: []protocol.Version{protocol.OCPP201}})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()
	conn := dialStation(t, httpServer.URL)
	defer conn.Close()

	sendCall(t, conn, "before", v201.HeartbeatRequest{})
	if response := readMessage(t, conn); response.Type() != protocol.CallErrorType {
		t.Fatalf("response type = %d", response.Type())
	}
	if heartbeats.Load() != 0 {
		t.Fatal("Heartbeat ran before BootNotification acceptance")
	}

	sendCall(t, conn, "boot", v201.BootNotificationRequest{
		Reason:          v201.BootNotificationRequestBootReasonEnumPowerUp,
		ChargingStation: v201.BootNotificationRequestChargingStation{Model: "AC-22K", VendorName: "Example"},
	})
	if response := readMessage(t, conn); response.Type() != protocol.CallResultType {
		t.Fatalf("response type = %d", response.Type())
	}
	session, ok := server.Session("CP-201")
	if !ok || !profile.IsBooted(session) {
		t.Fatal("session was not marked booted")
	}

	sendCall(t, conn, "heartbeat", v201.HeartbeatRequest{})
	if response := readMessage(t, conn); response.Type() != protocol.CallResultType {
		t.Fatalf("response type = %d", response.Type())
	}
	if heartbeats.Load() != 1 {
		t.Fatalf("heartbeat calls = %d", heartbeats.Load())
	}

	result := make(chan v201.ResetConfirmation, 1)
	errResult := make(chan error, 1)
	go func() {
		confirmation, err := profile.CallReset(context.Background(), session, v201.ResetRequest{Type: v201.ResetRequestResetEnumImmediate})
		if err != nil {
			errResult <- err
			return
		}
		result <- confirmation
	}()
	outbound := readMessage(t, conn).(protocol.Call)
	sendCallResult(t, conn, outbound.ID, v201.ResetConfirmation{Status: v201.ResetConfirmationResetStatusEnumAccepted})
	select {
	case err := <-errResult:
		t.Fatal(err)
	case confirmation := <-result:
		if confirmation.Status != v201.ResetConfirmationResetStatusEnumAccepted {
			t.Fatalf("status = %q", confirmation.Status)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for Reset confirmation")
	}
}

func TestNewProfileRejectsNilRouterAndHandler(t *testing.T) {
	if _, err := NewProfile(nil); err == nil {
		t.Fatal("NewProfile(nil) succeeded")
	}
	profile, err := NewProfile(csms.NewRouter())
	if err != nil {
		t.Fatal(err)
	}
	if err := profile.HandleBootNotification(nil); err == nil {
		t.Fatal("nil BootNotification handler succeeded")
	}
	if profile.State(nil) != RegistrationUnknown {
		t.Fatalf("nil state = %d", profile.State(nil))
	}
}

func dialStation(t *testing.T, serverURL string) *websocket.Conn {
	t.Helper()
	dialer := websocket.Dialer{Subprotocols: []string{string(protocol.OCPP201)}}
	conn, _, err := dialer.Dial("ws"+strings.TrimPrefix(serverURL, "http")+"/CP-201", nil)
	if err != nil {
		t.Fatal(err)
	}
	return conn
}

func sendCall(t *testing.T, conn *websocket.Conn, id string, payload protocol.Payload) {
	t.Helper()
	raw, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}
	data, err := protocol.Encode(protocol.Call{ID: id, Action: payload.ActionName(), Payload: raw})
	if err != nil {
		t.Fatal(err)
	}
	if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
		t.Fatal(err)
	}
}

func sendCallResult(t *testing.T, conn *websocket.Conn, id string, payload protocol.Payload) {
	t.Helper()
	raw, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}
	data, err := protocol.Encode(protocol.CallResult{ID: id, Payload: raw})
	if err != nil {
		t.Fatal(err)
	}
	if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
		t.Fatal(err)
	}
}

func readMessage(t *testing.T, conn *websocket.Conn) protocol.Message {
	t.Helper()
	_ = conn.SetReadDeadline(time.Now().Add(time.Second))
	_, data, err := conn.ReadMessage()
	if err != nil {
		t.Fatal(err)
	}
	message, err := protocol.Decode(data)
	if err != nil {
		t.Fatal(err)
	}
	return message
}
