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
	var transactions, reports, clearedLimits, schedulePulls, derStarts, tariffGets, batterySwaps atomic.Int32
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
	if err := profile.HandleClearedChargingLimit(func(_ context.Context, _ *csms.Session, request v21.ClearedChargingLimitRequest) (v21.ClearedChargingLimitConfirmation, error) {
		if request.ChargingLimitSource != "CSO" {
			t.Errorf("charging limit source = %q", request.ChargingLimitSource)
		}
		clearedLimits.Add(1)
		return v21.ClearedChargingLimitConfirmation{}, nil
	}); err != nil {
		t.Fatal(err)
	}
	if err := profile.HandlePullDynamicScheduleUpdate(func(_ context.Context, _ *csms.Session, request v21.PullDynamicScheduleUpdateRequest) (v21.PullDynamicScheduleUpdateConfirmation, error) {
		if request.ChargingProfileID != 210 {
			t.Errorf("charging profile ID = %d", request.ChargingProfileID)
		}
		schedulePulls.Add(1)
		return v21.PullDynamicScheduleUpdateConfirmation{Status: v21.PullDynamicScheduleUpdateConfirmationChargingProfileStatusEnumAccepted}, nil
	}); err != nil {
		t.Fatal(err)
	}
	if err := profile.HandleNotifyDERStartStop(func(_ context.Context, _ *csms.Session, request v21.NotifyDERStartStopRequest) (v21.NotifyDERStartStopConfirmation, error) {
		if request.ControlID != "DER-21" || !request.Started {
			t.Errorf("DER start/stop = %#v", request)
		}
		derStarts.Add(1)
		return v21.NotifyDERStartStopConfirmation{}, nil
	}); err != nil {
		t.Fatal(err)
	}
	if err := profile.HandleGetTariffs(func(_ context.Context, _ *csms.Session, request v21.GetTariffsRequest) (v21.GetTariffsConfirmation, error) {
		if request.EVSEID != 1 {
			t.Errorf("EVSE ID = %d", request.EVSEID)
		}
		tariffGets.Add(1)
		return v21.GetTariffsConfirmation{Status: v21.GetTariffsConfirmationTariffGetStatusEnumNoTariff}, nil
	}); err != nil {
		t.Fatal(err)
	}
	if err := profile.HandleBatterySwap(func(_ context.Context, _ *csms.Session, request v21.BatterySwapRequest) (v21.BatterySwapConfirmation, error) {
		if request.RequestID != 31 || len(request.BatteryData) != 1 {
			t.Errorf("BatterySwap = %#v", request)
		}
		batterySwaps.Add(1)
		return v21.BatterySwapConfirmation{}, nil
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
	sendCall(t, conn, "cleared-limit", v21.ClearedChargingLimitRequest{ChargingLimitSource: "CSO"})
	if response := readMessage(t, conn); response.Type() != protocol.CallResultType {
		t.Fatalf("ClearedChargingLimit response type = %d", response.Type())
	}
	sendCall(t, conn, "pull-schedule", v21.PullDynamicScheduleUpdateRequest{ChargingProfileID: 210})
	if response := readMessage(t, conn); response.Type() != protocol.CallResultType {
		t.Fatalf("PullDynamicScheduleUpdate response type = %d", response.Type())
	}
	sendCall(t, conn, "der-start", v21.NotifyDERStartStopRequest{ControlID: "DER-21", Started: true, Timestamp: "2026-07-16T00:02:30Z"})
	if response := readMessage(t, conn); response.Type() != protocol.CallResultType {
		t.Fatalf("NotifyDERStartStop response type = %d", response.Type())
	}
	sendCall(t, conn, "get-tariffs", v21.GetTariffsRequest{EVSEID: 1})
	if response := readMessage(t, conn); response.Type() != protocol.CallResultType {
		t.Fatalf("GetTariffs response type = %d", response.Type())
	}
	sendCall(t, conn, "battery-in", v21.BatterySwapRequest{
		RequestID: 31, EventType: v21.BatterySwapRequestBatterySwapEventEnumBatteryIn,
		IDToken:     v21.BatterySwapRequestIDToken{IDToken: "BATTERY-USER", Type: "Local"},
		BatteryData: []v21.BatterySwapRequestBatteryData{{EVSEID: 1, SerialNumber: "BAT-001", SoC: 80, SoH: 95}},
	})
	if response := readMessage(t, conn); response.Type() != protocol.CallResultType {
		t.Fatalf("BatterySwap response type = %d", response.Type())
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

	profilesResult := make(chan v21.GetChargingProfilesConfirmation, 1)
	profilesError := make(chan error, 1)
	go func() {
		confirmation, err := profile.CallGetChargingProfiles(context.Background(), session, v21.GetChargingProfilesRequest{
			RequestID: 22, ChargingProfile: v21.GetChargingProfilesRequestChargingProfileCriterion{},
		})
		if err != nil {
			profilesError <- err
			return
		}
		profilesResult <- confirmation
	}()
	outbound = readMessage(t, conn).(protocol.Call)
	if outbound.Action != "GetChargingProfiles" {
		t.Fatalf("outbound action = %q", outbound.Action)
	}
	sendCallResult(t, conn, outbound.ID, v21.GetChargingProfilesConfirmation{Status: v21.GetChargingProfilesConfirmationGetChargingProfileStatusEnumNoProfiles})
	select {
	case err := <-profilesError:
		t.Fatal(err)
	case confirmation := <-profilesResult:
		if confirmation.Status != v21.GetChargingProfilesConfirmationGetChargingProfileStatusEnumNoProfiles {
			t.Fatalf("status = %q", confirmation.Status)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for GetChargingProfiles confirmation")
	}

	priorityResult := make(chan v21.UsePriorityChargingConfirmation, 1)
	priorityError := make(chan error, 1)
	go func() {
		confirmation, err := profile.CallUsePriorityCharging(context.Background(), session, v21.UsePriorityChargingRequest{TransactionID: "TX-21", Activate: true})
		if err != nil {
			priorityError <- err
			return
		}
		priorityResult <- confirmation
	}()
	outbound = readMessage(t, conn).(protocol.Call)
	if outbound.Action != "UsePriorityCharging" {
		t.Fatalf("outbound action = %q", outbound.Action)
	}
	sendCallResult(t, conn, outbound.ID, v21.UsePriorityChargingConfirmation{Status: v21.UsePriorityChargingConfirmationPriorityChargingStatusEnumAccepted})
	select {
	case err := <-priorityError:
		t.Fatal(err)
	case confirmation := <-priorityResult:
		if confirmation.Status != v21.UsePriorityChargingConfirmationPriorityChargingStatusEnumAccepted {
			t.Fatalf("status = %q", confirmation.Status)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for UsePriorityCharging confirmation")
	}

	derResult := make(chan v21.ClearDERControlConfirmation, 1)
	derError := make(chan error, 1)
	go func() {
		confirmation, err := profile.CallClearDERControl(context.Background(), session, v21.ClearDERControlRequest{IsDefault: true})
		if err != nil {
			derError <- err
			return
		}
		derResult <- confirmation
	}()
	outbound = readMessage(t, conn).(protocol.Call)
	if outbound.Action != "ClearDERControl" {
		t.Fatalf("outbound action = %q", outbound.Action)
	}
	sendCallResult(t, conn, outbound.ID, v21.ClearDERControlConfirmation{Status: v21.ClearDERControlConfirmationDERControlStatusEnumAccepted})
	select {
	case err := <-derError:
		t.Fatal(err)
	case confirmation := <-derResult:
		if confirmation.Status != v21.ClearDERControlConfirmationDERControlStatusEnumAccepted {
			t.Fatalf("status = %q", confirmation.Status)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for ClearDERControl confirmation")
	}

	costResult := make(chan v21.CostUpdatedConfirmation, 1)
	costError := make(chan error, 1)
	go func() {
		confirmation, err := profile.CallCostUpdated(context.Background(), session, v21.CostUpdatedRequest{TotalCost: 12.5, TransactionID: "TX-21"})
		if err != nil {
			costError <- err
			return
		}
		costResult <- confirmation
	}()
	outbound = readMessage(t, conn).(protocol.Call)
	if outbound.Action != "CostUpdated" {
		t.Fatalf("outbound action = %q", outbound.Action)
	}
	sendCallResult(t, conn, outbound.ID, v21.CostUpdatedConfirmation{})
	select {
	case err := <-costError:
		t.Fatal(err)
	case <-costResult:
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for CostUpdated confirmation")
	}

	batteryResult := make(chan v21.RequestBatterySwapConfirmation, 1)
	batteryError := make(chan error, 1)
	go func() {
		confirmation, err := profile.CallRequestBatterySwap(context.Background(), session, v21.RequestBatterySwapRequest{
			RequestID: 32, IDToken: v21.RequestBatterySwapRequestIDToken{IDToken: "BATTERY-USER", Type: "Local"},
		})
		if err != nil {
			batteryError <- err
			return
		}
		batteryResult <- confirmation
	}()
	outbound = readMessage(t, conn).(protocol.Call)
	if outbound.Action != "RequestBatterySwap" {
		t.Fatalf("outbound action = %q", outbound.Action)
	}
	sendCallResult(t, conn, outbound.ID, v21.RequestBatterySwapConfirmation{Status: v21.RequestBatterySwapConfirmationGenericStatusEnumAccepted})
	select {
	case err := <-batteryError:
		t.Fatal(err)
	case confirmation := <-batteryResult:
		if confirmation.Status != v21.RequestBatterySwapConfirmationGenericStatusEnumAccepted {
			t.Fatalf("status = %q", confirmation.Status)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for RequestBatterySwap confirmation")
	}

	streamResult := make(chan v21.OpenPeriodicEventStreamConfirmation, 1)
	streamError := make(chan error, 1)
	go func() {
		confirmation, err := profile.CallOpenPeriodicEventStream(context.Background(), session, v21.OpenPeriodicEventStreamRequest{
			ConstantStreamData: v21.OpenPeriodicEventStreamRequestConstantStreamData{
				ID: 41, VariableMonitoringID: 410,
				Params: v21.OpenPeriodicEventStreamRequestPeriodicEventStreamParams{},
			},
		})
		if err != nil {
			streamError <- err
			return
		}
		streamResult <- confirmation
	}()
	outbound = readMessage(t, conn).(protocol.Call)
	if outbound.Action != "OpenPeriodicEventStream" {
		t.Fatalf("outbound action = %q", outbound.Action)
	}
	sendCallResult(t, conn, outbound.ID, v21.OpenPeriodicEventStreamConfirmation{Status: v21.OpenPeriodicEventStreamConfirmationGenericStatusEnumAccepted})
	select {
	case err := <-streamError:
		t.Fatal(err)
	case confirmation := <-streamResult:
		if confirmation.Status != v21.OpenPeriodicEventStreamConfirmationGenericStatusEnumAccepted {
			t.Fatalf("status = %q", confirmation.Status)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for OpenPeriodicEventStream confirmation")
	}
	if transactions.Load() != 3 || reports.Load() != 1 || clearedLimits.Load() != 1 || schedulePulls.Load() != 1 || derStarts.Load() != 1 || tariffGets.Load() != 1 || batterySwaps.Load() != 1 {
		t.Fatalf("transactions=%d reports=%d clearedLimits=%d schedulePulls=%d derStarts=%d tariffGets=%d batterySwaps=%d", transactions.Load(), reports.Load(), clearedLimits.Load(), schedulePulls.Load(), derStarts.Load(), tariffGets.Load(), batterySwaps.Load())
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
