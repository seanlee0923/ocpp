package ocpp21

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
	"ocpp-go/v21"
)

func TestRegistrationGateAndResetRoundTrip(t *testing.T) {
	router := csms.NewRouter()
	profile, err := NewProfile(router)
	if err != nil {
		t.Fatal(err)
	}
	var heartbeats atomic.Int32
	if err := profile.HandleBootNotification(func(context.Context, *csms.Session, v21.BootNotificationRequest) (v21.BootNotificationConfirmation, error) {
		return v21.BootNotificationConfirmation{CurrentTime: "2026-07-16T00:00:00Z", Interval: 300, Status: v21.BootNotificationConfirmationRegistrationStatusEnumAccepted}, nil
	}); err != nil {
		t.Fatal(err)
	}
	if err := profile.HandleHeartbeat(func(context.Context, *csms.Session, v21.HeartbeatRequest) (v21.HeartbeatConfirmation, error) {
		heartbeats.Add(1)
		return v21.HeartbeatConfirmation{CurrentTime: "2026-07-16T00:00:01Z"}, nil
	}); err != nil {
		t.Fatal(err)
	}

	server, err := csms.New(csms.Config{Router: router, Versions: []protocol.Version{protocol.OCPP21}})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()
	dialer := websocket.Dialer{Subprotocols: []string{string(protocol.OCPP21)}}
	conn, _, err := dialer.Dial("ws"+strings.TrimPrefix(httpServer.URL, "http")+"/CP-21", nil)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	sendCall(t, conn, "before", v21.HeartbeatRequest{})
	if response := readMessage(t, conn); response.Type() != protocol.CallErrorType {
		t.Fatalf("response type = %d", response.Type())
	}
	if heartbeats.Load() != 0 {
		t.Fatal("Heartbeat ran before BootNotification acceptance")
	}
	sendCall(t, conn, "boot", v21.BootNotificationRequest{
		Reason:          v21.BootNotificationRequestBootReasonEnumPowerUp,
		ChargingStation: v21.BootNotificationRequestChargingStation{Model: "AC-22K", VendorName: "Example"},
	})
	if response := readMessage(t, conn); response.Type() != protocol.CallResultType {
		t.Fatalf("response type = %d", response.Type())
	}
	session, ok := server.Session("CP-21")
	if !ok || !profile.IsBooted(session) {
		t.Fatal("session was not marked booted")
	}

	sendCall(t, conn, "heartbeat", v21.HeartbeatRequest{})
	if response := readMessage(t, conn); response.Type() != protocol.CallResultType {
		t.Fatalf("response type = %d", response.Type())
	}

	result := make(chan v21.ResetConfirmation, 1)
	errResult := make(chan error, 1)
	go func() {
		confirmation, err := profile.CallReset(context.Background(), session, v21.ResetRequest{Type: v21.ResetRequestResetEnumImmediate})
		if err != nil {
			errResult <- err
			return
		}
		result <- confirmation
	}()
	outbound := readMessage(t, conn).(protocol.Call)
	sendCallResult(t, conn, outbound.ID, v21.ResetConfirmation{Status: v21.ResetConfirmationResetStatusEnumAccepted})
	select {
	case err := <-errResult:
		t.Fatal(err)
	case confirmation := <-result:
		if confirmation.Status != v21.ResetConfirmationResetStatusEnumAccepted {
			t.Fatalf("status = %q", confirmation.Status)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for Reset confirmation")
	}
}

func TestTransactionAndDeviceModelRoundTrips(t *testing.T) {
	router := csms.NewRouter()
	profile, err := NewProfile(router)
	if err != nil {
		t.Fatal(err)
	}
	var transactions, reports atomic.Int32
	if err := profile.HandleBootNotification(func(context.Context, *csms.Session, v21.BootNotificationRequest) (v21.BootNotificationConfirmation, error) {
		return v21.BootNotificationConfirmation{CurrentTime: "2026-07-16T00:00:00Z", Interval: 300, Status: v21.BootNotificationConfirmationRegistrationStatusEnumAccepted}, nil
	}); err != nil {
		t.Fatal(err)
	}
	if err := profile.HandleTransactionEvent(func(_ context.Context, _ *csms.Session, request v21.TransactionEventRequest) (v21.TransactionEventConfirmation, error) {
		if request.TransactionInfo.TransactionID != "TX-21" {
			t.Errorf("transaction ID = %q", request.TransactionInfo.TransactionID)
		}
		transactions.Add(1)
		return v21.TransactionEventConfirmation{}, nil
	}); err != nil {
		t.Fatal(err)
	}
	if err := profile.HandleNotifyReport(func(context.Context, *csms.Session, v21.NotifyReportRequest) (v21.NotifyReportConfirmation, error) {
		reports.Add(1)
		return v21.NotifyReportConfirmation{}, nil
	}); err != nil {
		t.Fatal(err)
	}

	server, err := csms.New(csms.Config{Router: router, Versions: []protocol.Version{protocol.OCPP21}})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()
	dialer := websocket.Dialer{Subprotocols: []string{string(protocol.OCPP21)}}
	conn, _, err := dialer.Dial("ws"+strings.TrimPrefix(httpServer.URL, "http")+"/CP-21", nil)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	sendCall(t, conn, "boot", v21.BootNotificationRequest{
		Reason:          v21.BootNotificationRequestBootReasonEnumPowerUp,
		ChargingStation: v21.BootNotificationRequestChargingStation{Model: "AC-22K", VendorName: "Example"},
	})
	if response := readMessage(t, conn); response.Type() != protocol.CallResultType {
		t.Fatalf("Boot response type = %d", response.Type())
	}
	session, ok := server.Session("CP-21")
	if !ok {
		t.Fatal("session not found")
	}

	events := []v21.TransactionEventRequestTransactionEventEnum{
		v21.TransactionEventRequestTransactionEventEnumStarted,
		v21.TransactionEventRequestTransactionEventEnumUpdated,
		v21.TransactionEventRequestTransactionEventEnumEnded,
	}
	for seq, eventType := range events {
		sendCall(t, conn, "transaction", v21.TransactionEventRequest{
			EventType: eventType, Timestamp: "2026-07-16T00:01:00Z",
			TriggerReason: v21.TransactionEventRequestTriggerReasonEnumAuthorized, SeqNo: seq,
			TransactionInfo: v21.TransactionEventRequestTransaction{TransactionID: "TX-21"},
		})
		if response := readMessage(t, conn); response.Type() != protocol.CallResultType {
			t.Fatalf("TransactionEvent response type = %d", response.Type())
		}
	}

	sendCall(t, conn, "report", v21.NotifyReportRequest{RequestID: 21, GeneratedAt: "2026-07-16T00:02:00Z", SeqNo: 0})
	if response := readMessage(t, conn); response.Type() != protocol.CallResultType {
		t.Fatalf("NotifyReport response type = %d", response.Type())
	}

	variableResult := make(chan v21.GetVariablesConfirmation, 1)
	variableError := make(chan error, 1)
	go func() {
		confirmation, err := profile.CallGetVariables(context.Background(), session, v21.GetVariablesRequest{GetVariableData: []v21.GetVariablesRequestGetVariableData{{
			Component: v21.GetVariablesRequestComponent{Name: "ClockCtrlr"}, Variable: v21.GetVariablesRequestVariable{Name: "DateTime"},
		}}})
		if err != nil {
			variableError <- err
			return
		}
		variableResult <- confirmation
	}()
	outbound := readMessage(t, conn).(protocol.Call)
	value := "2026-07-16T00:03:00Z"
	sendCallResult(t, conn, outbound.ID, v21.GetVariablesConfirmation{GetVariableResult: []v21.GetVariablesConfirmationGetVariableResult{{
		AttributeStatus: v21.GetVariablesConfirmationGetVariableStatusEnumAccepted,
		AttributeValue:  &value,
		Component:       v21.GetVariablesConfirmationComponent{Name: "ClockCtrlr"},
		Variable:        v21.GetVariablesConfirmationVariable{Name: "DateTime"},
	}}})
	select {
	case err := <-variableError:
		t.Fatal(err)
	case confirmation := <-variableResult:
		if len(confirmation.GetVariableResult) != 1 || confirmation.GetVariableResult[0].AttributeValue == nil || *confirmation.GetVariableResult[0].AttributeValue != value {
			t.Fatalf("GetVariables confirmation = %#v", confirmation)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for GetVariables confirmation")
	}
	if transactions.Load() != 3 || reports.Load() != 1 {
		t.Fatalf("transactions=%d reports=%d", transactions.Load(), reports.Load())
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
		t.Fatal("nil handler succeeded")
	}
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
