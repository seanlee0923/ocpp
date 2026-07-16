package ocpp16

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
	"github.com/seanlee0923/ocpp/v16"
)

type testHandlers struct {
	bootStatus    v16.BootNotificationConfirmationStatus
	heartbeats    atomic.Int32
	statuses      atomic.Int32
	authorizes    atomic.Int32
	starts        atomic.Int32
	meterValues   atomic.Int32
	stops         atomic.Int32
	dataTransfers atomic.Int32
}

func (service *testHandlers) BootNotification(context.Context, *csms.Session, v16.BootNotificationRequest) (v16.BootNotificationConfirmation, error) {
	return v16.BootNotificationConfirmation{
		Status: service.bootStatus, CurrentTime: "2026-07-16T00:00:00Z", Interval: 300,
	}, nil
}

func (service *testHandlers) Heartbeat(context.Context, *csms.Session, v16.HeartbeatRequest) (v16.HeartbeatConfirmation, error) {
	service.heartbeats.Add(1)
	return v16.HeartbeatConfirmation{CurrentTime: "2026-07-16T00:00:01Z"}, nil
}

func (service *testHandlers) StatusNotification(context.Context, *csms.Session, v16.StatusNotificationRequest) (v16.StatusNotificationConfirmation, error) {
	service.statuses.Add(1)
	return v16.StatusNotificationConfirmation{}, nil
}

func (service *testHandlers) Authorize(_ context.Context, _ *csms.Session, request v16.AuthorizeRequest) (v16.AuthorizeConfirmation, error) {
	service.authorizes.Add(1)
	status := v16.AuthorizeConfirmationIDTagInfoStatusInvalid
	if request.IDTag == "TAG-001" {
		status = v16.AuthorizeConfirmationIDTagInfoStatusAccepted
	}
	return v16.AuthorizeConfirmation{
		IDTagInfo: v16.AuthorizeConfirmationIDTagInfo{Status: status},
	}, nil
}

func (service *testHandlers) StartTransaction(_ context.Context, _ *csms.Session, request v16.StartTransactionRequest) (v16.StartTransactionConfirmation, error) {
	service.starts.Add(1)
	status := v16.StartTransactionConfirmationIDTagInfoStatusInvalid
	if request.IDTag == "TAG-001" && request.ConnectorID == 1 {
		status = v16.StartTransactionConfirmationIDTagInfoStatusAccepted
	}
	return v16.StartTransactionConfirmation{
		TransactionID: 7001,
		IDTagInfo:     v16.StartTransactionConfirmationIDTagInfo{Status: status},
	}, nil
}

func (service *testHandlers) MeterValues(_ context.Context, _ *csms.Session, request v16.MeterValuesRequest) (v16.MeterValuesConfirmation, error) {
	service.meterValues.Add(1)
	return v16.MeterValuesConfirmation{}, nil
}

func (service *testHandlers) StopTransaction(_ context.Context, _ *csms.Session, request v16.StopTransactionRequest) (v16.StopTransactionConfirmation, error) {
	service.stops.Add(1)
	return v16.StopTransactionConfirmation{
		IDTagInfo: &v16.StopTransactionConfirmationIDTagInfo{Status: v16.StopTransactionConfirmationIDTagInfoStatusAccepted},
	}, nil
}

func (service *testHandlers) DataTransfer(_ context.Context, _ *csms.Session, request v16.DataTransferRequest) (v16.DataTransferConfirmation, error) {
	service.dataTransfers.Add(1)
	return v16.DataTransferConfirmation{Status: v16.DataTransferConfirmationStatusAccepted, Data: request.Data}, nil
}

func TestRegistrationFlowGatesActionsUntilBootAccepted(t *testing.T) {
	router := csms.NewRouter()
	service := &testHandlers{bootStatus: v16.BootNotificationConfirmationStatusAccepted}
	profile, err := NewProfile(router)
	if err != nil {
		t.Fatal(err)
	}
	registerTestHandlers(t, profile, service)
	server, err := csms.New(csms.Config{Router: router, Versions: []protocol.Version{protocol.OCPP16}})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()
	conn := dialStation(t, httpServer.URL)
	defer conn.Close()

	sendCall(t, conn, "1", v16.HeartbeatRequest{})
	if callError := readMessage(t, conn).(protocol.CallError); callError.Code != string(csms.ProtocolError) {
		t.Fatalf("CALLERROR code = %q", callError.Code)
	}
	if service.heartbeats.Load() != 0 {
		t.Fatal("Heartbeat service ran before boot acceptance")
	}
	sendCall(t, conn, "authorize-before-boot", v16.AuthorizeRequest{IDTag: "TAG-001"})
	if callError := readMessage(t, conn).(protocol.CallError); callError.Code != string(csms.ProtocolError) {
		t.Fatalf("Authorize CALLERROR code = %q", callError.Code)
	}
	if service.authorizes.Load() != 0 {
		t.Fatal("Authorize handler ran before boot acceptance")
	}
	sendCall(t, conn, "start-before-boot", startTransactionRequest())
	if callError := readMessage(t, conn).(protocol.CallError); callError.Code != string(csms.ProtocolError) {
		t.Fatalf("StartTransaction CALLERROR code = %q", callError.Code)
	}
	if service.starts.Load() != 0 {
		t.Fatal("StartTransaction handler ran before boot acceptance")
	}

	sendCall(t, conn, "2", v16.BootNotificationRequest{
		ChargePointVendor: "Example", ChargePointModel: "AC-22K",
	})
	if _, ok := readMessage(t, conn).(protocol.CallResult); !ok {
		t.Fatal("BootNotification did not return CALLRESULT")
	}
	session, ok := server.Session("CP-001")
	if !ok || !profile.IsBooted(session) {
		t.Fatal("session was not marked booted")
	}

	sendCall(t, conn, "3", v16.HeartbeatRequest{})
	if _, ok := readMessage(t, conn).(protocol.CallResult); !ok {
		t.Fatal("Heartbeat did not return CALLRESULT")
	}
	sendCall(t, conn, "4", v16.StatusNotificationRequest{
		ConnectorID: 0,
		ErrorCode:   v16.StatusNotificationRequestErrorCodeNoError,
		Status:      v16.StatusNotificationRequestStatusAvailable,
	})
	if _, ok := readMessage(t, conn).(protocol.CallResult); !ok {
		t.Fatal("StatusNotification did not return CALLRESULT")
	}
	if service.heartbeats.Load() != 1 || service.statuses.Load() != 1 {
		t.Fatalf("heartbeat calls = %d, status calls = %d", service.heartbeats.Load(), service.statuses.Load())
	}

	sendCall(t, conn, "5", v16.AuthorizeRequest{IDTag: "TAG-001"})
	result, ok := readMessage(t, conn).(protocol.CallResult)
	if !ok {
		t.Fatal("Authorize did not return CALLRESULT")
	}
	var authorization v16.AuthorizeConfirmation
	if err := json.Unmarshal(result.Payload, &authorization); err != nil {
		t.Fatal(err)
	}
	if authorization.IDTagInfo.Status != v16.AuthorizeConfirmationIDTagInfoStatusAccepted {
		t.Fatalf("Authorize status = %q", authorization.IDTagInfo.Status)
	}
	if service.authorizes.Load() != 1 {
		t.Fatalf("Authorize calls = %d", service.authorizes.Load())
	}

	sendCall(t, conn, "6", startTransactionRequest())
	result, ok = readMessage(t, conn).(protocol.CallResult)
	if !ok {
		t.Fatal("StartTransaction did not return CALLRESULT")
	}
	var started v16.StartTransactionConfirmation
	if err := json.Unmarshal(result.Payload, &started); err != nil {
		t.Fatal(err)
	}
	if started.TransactionID != 7001 || started.IDTagInfo.Status != v16.StartTransactionConfirmationIDTagInfoStatusAccepted {
		t.Fatalf("StartTransaction confirmation = %#v", started)
	}
	if service.starts.Load() != 1 {
		t.Fatalf("StartTransaction calls = %d", service.starts.Load())
	}

	transactionID := 7001
	sendCall(t, conn, "7", v16.MeterValuesRequest{
		ConnectorID:   1,
		TransactionID: &transactionID,
		MeterValue: []v16.MeterValuesRequestMeterValueItem{{
			Timestamp:    "2026-07-16T00:01:00Z",
			SampledValue: []v16.MeterValuesRequestMeterValueItemSampledValueItem{{Value: "1300"}},
		}},
	})
	if _, ok := readMessage(t, conn).(protocol.CallResult); !ok {
		t.Fatal("MeterValues did not return CALLRESULT")
	}

	data := `{"vendor":"payload"}`
	sendCall(t, conn, "8", v16.DataTransferRequest{VendorID: "example.com", Data: &data})
	if _, ok := readMessage(t, conn).(protocol.CallResult); !ok {
		t.Fatal("DataTransfer did not return CALLRESULT")
	}

	sendCall(t, conn, "9", v16.StopTransactionRequest{
		MeterStop: 1400, Timestamp: "2026-07-16T00:02:00Z", TransactionID: transactionID,
	})
	result, ok = readMessage(t, conn).(protocol.CallResult)
	if !ok {
		t.Fatal("StopTransaction did not return CALLRESULT")
	}
	var stopped v16.StopTransactionConfirmation
	if err := json.Unmarshal(result.Payload, &stopped); err != nil {
		t.Fatal(err)
	}
	if stopped.IDTagInfo == nil || stopped.IDTagInfo.Status != v16.StopTransactionConfirmationIDTagInfoStatusAccepted {
		t.Fatalf("StopTransaction confirmation = %#v", stopped)
	}
	if service.meterValues.Load() != 1 || service.dataTransfers.Load() != 1 || service.stops.Load() != 1 {
		t.Fatalf("meter values = %d, data transfers = %d, stops = %d", service.meterValues.Load(), service.dataTransfers.Load(), service.stops.Load())
	}

	resetResult := make(chan v16.ResetConfirmation, 1)
	resetError := make(chan error, 1)
	go func() {
		confirmation, err := profile.CallReset(context.Background(), session, v16.ResetRequest{Type: v16.ResetRequestTypeSoft})
		if err != nil {
			resetError <- err
			return
		}
		resetResult <- confirmation
	}()
	outbound, ok := readMessage(t, conn).(protocol.Call)
	if !ok || outbound.Action != "Reset" {
		t.Fatalf("outbound message = %#v", outbound)
	}
	sendCallResult(t, conn, outbound.ID, v16.ResetConfirmation{Status: v16.ResetConfirmationStatusAccepted})
	select {
	case err := <-resetError:
		t.Fatal(err)
	case confirmation := <-resetResult:
		if confirmation.Status != v16.ResetConfirmationStatusAccepted {
			t.Fatalf("Reset status = %q", confirmation.Status)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for Reset confirmation")
	}
}

func TestPendingBootDoesNotUnlockOtherActions(t *testing.T) {
	router := csms.NewRouter()
	service := &testHandlers{bootStatus: v16.BootNotificationConfirmationStatusPending}
	profile, err := NewProfile(router)
	if err != nil {
		t.Fatal(err)
	}
	registerTestHandlers(t, profile, service)
	server, err := csms.New(csms.Config{Router: router, Versions: []protocol.Version{protocol.OCPP16}})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()
	conn := dialStation(t, httpServer.URL)
	defer conn.Close()
	sendCall(t, conn, "1", v16.BootNotificationRequest{ChargePointVendor: "Example", ChargePointModel: "AC-22K"})
	readMessage(t, conn)
	session, _ := server.Session("CP-001")
	if profile.State(session) != RegistrationPending {
		t.Fatalf("state = %d", profile.State(session))
	}
	sendCall(t, conn, "2", v16.HeartbeatRequest{})
	if message := readMessage(t, conn); message.Type() != protocol.CallErrorType {
		t.Fatalf("response type = %d", message.Type())
	}
}

func registerTestHandlers(t *testing.T, profile *Profile, handlers *testHandlers) {
	t.Helper()
	if err := profile.HandleBootNotification(handlers.BootNotification); err != nil {
		t.Fatal(err)
	}
	if err := profile.HandleHeartbeat(handlers.Heartbeat); err != nil {
		t.Fatal(err)
	}
	if err := profile.HandleStatusNotification(handlers.StatusNotification); err != nil {
		t.Fatal(err)
	}
	if err := profile.HandleAuthorize(handlers.Authorize); err != nil {
		t.Fatal(err)
	}
	if err := profile.HandleStartTransaction(handlers.StartTransaction); err != nil {
		t.Fatal(err)
	}
	if err := profile.HandleMeterValues(handlers.MeterValues); err != nil {
		t.Fatal(err)
	}
	if err := profile.HandleStopTransaction(handlers.StopTransaction); err != nil {
		t.Fatal(err)
	}
	if err := profile.HandleDataTransfer(handlers.DataTransfer); err != nil {
		t.Fatal(err)
	}
}

func startTransactionRequest() v16.StartTransactionRequest {
	return v16.StartTransactionRequest{
		ConnectorID: 1,
		IDTag:       "TAG-001",
		MeterStart:  1234,
		Timestamp:   "2026-07-16T00:00:02Z",
	}
}

func dialStation(t *testing.T, serverURL string) *websocket.Conn {
	t.Helper()
	dialer := websocket.Dialer{Subprotocols: []string{string(protocol.OCPP16)}}
	conn, _, err := dialer.Dial("ws"+strings.TrimPrefix(serverURL, "http")+"/CP-001", nil)
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
