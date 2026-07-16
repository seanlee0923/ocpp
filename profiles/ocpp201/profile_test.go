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
	"github.com/seanlee0923/ocpp/csms"
	"github.com/seanlee0923/ocpp/protocol"
	"github.com/seanlee0923/ocpp/v201"
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

func TestTransactionAndDeviceModelRoundTrips(t *testing.T) {
	router := csms.NewRouter()
	profile, err := NewProfile(router)
	if err != nil {
		t.Fatal(err)
	}
	var authorizes, transactions, meterValues, reports atomic.Int32
	if err := profile.HandleBootNotification(acceptedBootHandler); err != nil {
		t.Fatal(err)
	}
	if err := profile.HandleAuthorize(func(_ context.Context, _ *csms.Session, request v201.AuthorizeRequest) (v201.AuthorizeConfirmation, error) {
		authorizes.Add(1)
		status := v201.AuthorizeConfirmationAuthorizationStatusEnumInvalid
		if request.IDToken.IDToken == "TAG-201" {
			status = v201.AuthorizeConfirmationAuthorizationStatusEnumAccepted
		}
		return v201.AuthorizeConfirmation{IDTokenInfo: v201.AuthorizeConfirmationIDTokenInfo{Status: status}}, nil
	}); err != nil {
		t.Fatal(err)
	}
	var expectedSeq atomic.Int32
	if err := profile.HandleTransactionEvent(func(_ context.Context, _ *csms.Session, request v201.TransactionEventRequest) (v201.TransactionEventConfirmation, error) {
		if request.TransactionInfo.TransactionID != "TX-201" || request.SeqNo != int(expectedSeq.Load()) {
			t.Errorf("TransactionEvent = %#v", request)
		}
		expectedSeq.Add(1)
		transactions.Add(1)
		return v201.TransactionEventConfirmation{}, nil
	}); err != nil {
		t.Fatal(err)
	}
	if err := profile.HandleMeterValues(func(context.Context, *csms.Session, v201.MeterValuesRequest) (v201.MeterValuesConfirmation, error) {
		meterValues.Add(1)
		return v201.MeterValuesConfirmation{}, nil
	}); err != nil {
		t.Fatal(err)
	}
	if err := profile.HandleNotifyReport(func(context.Context, *csms.Session, v201.NotifyReportRequest) (v201.NotifyReportConfirmation, error) {
		reports.Add(1)
		return v201.NotifyReportConfirmation{}, nil
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
	bootStation(t, conn)
	session, ok := server.Session("CP-201")
	if !ok {
		t.Fatal("session not found")
	}

	sendCall(t, conn, "authorize", v201.AuthorizeRequest{IDToken: v201.AuthorizeRequestIDToken{IDToken: "TAG-201", Type: v201.AuthorizeRequestIDTokenEnumLocal}})
	result := readMessage(t, conn).(protocol.CallResult)
	var authorization v201.AuthorizeConfirmation
	if err := json.Unmarshal(result.Payload, &authorization); err != nil {
		t.Fatal(err)
	}
	if authorization.IDTokenInfo.Status != v201.AuthorizeConfirmationAuthorizationStatusEnumAccepted {
		t.Fatalf("authorization status = %q", authorization.IDTokenInfo.Status)
	}

	events := []v201.TransactionEventRequestTransactionEventEnum{
		v201.TransactionEventRequestTransactionEventEnumStarted,
		v201.TransactionEventRequestTransactionEventEnumUpdated,
		v201.TransactionEventRequestTransactionEventEnumEnded,
	}
	for seq, eventType := range events {
		sendCall(t, conn, "transaction-"+string(rune('0'+seq)), v201.TransactionEventRequest{
			EventType: eventType, Timestamp: "2026-07-16T00:01:00Z",
			TriggerReason: v201.TransactionEventRequestTriggerReasonEnumAuthorized, SeqNo: seq,
			TransactionInfo: v201.TransactionEventRequestTransaction{TransactionID: "TX-201"},
		})
		if response := readMessage(t, conn); response.Type() != protocol.CallResultType {
			t.Fatalf("TransactionEvent response type = %d", response.Type())
		}
	}

	sendCall(t, conn, "meter", v201.MeterValuesRequest{EVSEID: 1, MeterValue: []v201.MeterValuesRequestMeterValue{{
		Timestamp: "2026-07-16T00:02:00Z", SampledValue: []v201.MeterValuesRequestSampledValue{{Value: 12.5}},
	}}})
	if response := readMessage(t, conn); response.Type() != protocol.CallResultType {
		t.Fatalf("MeterValues response type = %d", response.Type())
	}
	sendCall(t, conn, "report", v201.NotifyReportRequest{RequestID: 7, GeneratedAt: "2026-07-16T00:03:00Z", SeqNo: 0})
	if response := readMessage(t, conn); response.Type() != protocol.CallResultType {
		t.Fatalf("NotifyReport response type = %d", response.Type())
	}

	variableResult := make(chan v201.GetVariablesConfirmation, 1)
	variableError := make(chan error, 1)
	go func() {
		confirmation, err := profile.CallGetVariables(context.Background(), session, v201.GetVariablesRequest{GetVariableData: []v201.GetVariablesRequestGetVariableData{{
			Component: v201.GetVariablesRequestComponent{Name: "ClockCtrlr"}, Variable: v201.GetVariablesRequestVariable{Name: "DateTime"},
		}}})
		if err != nil {
			variableError <- err
			return
		}
		variableResult <- confirmation
	}()
	outbound := readMessage(t, conn).(protocol.Call)
	value := "2026-07-16T00:04:00Z"
	sendCallResult(t, conn, outbound.ID, v201.GetVariablesConfirmation{GetVariableResult: []v201.GetVariablesConfirmationGetVariableResult{{
		AttributeStatus: v201.GetVariablesConfirmationGetVariableStatusEnumAccepted,
		AttributeValue:  &value,
		Component:       v201.GetVariablesConfirmationComponent{Name: "ClockCtrlr"},
		Variable:        v201.GetVariablesConfirmationVariable{Name: "DateTime"},
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

	if authorizes.Load() != 1 || transactions.Load() != 3 || meterValues.Load() != 1 || reports.Load() != 1 {
		t.Fatalf("calls: authorize=%d transaction=%d meter=%d report=%d", authorizes.Load(), transactions.Load(), meterValues.Load(), reports.Load())
	}
}

func acceptedBootHandler(context.Context, *csms.Session, v201.BootNotificationRequest) (v201.BootNotificationConfirmation, error) {
	return v201.BootNotificationConfirmation{
		CurrentTime: "2026-07-16T00:00:00Z", Interval: 300,
		Status: v201.BootNotificationConfirmationRegistrationStatusEnumAccepted,
	}, nil
}

func bootStation(t *testing.T, conn *websocket.Conn) {
	t.Helper()
	sendCall(t, conn, "boot", v201.BootNotificationRequest{
		Reason:          v201.BootNotificationRequestBootReasonEnumPowerUp,
		ChargingStation: v201.BootNotificationRequestChargingStation{Model: "AC-22K", VendorName: "Example"},
	})
	if response := readMessage(t, conn); response.Type() != protocol.CallResultType {
		t.Fatalf("BootNotification response type = %d", response.Type())
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
