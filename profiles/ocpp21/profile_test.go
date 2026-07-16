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
	"github.com/seanlee0923/ocpp/csms"
	"github.com/seanlee0923/ocpp/protocol"
	"github.com/seanlee0923/ocpp/v21"
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
	var transactions, reports, clearedLimits, schedulePulls, derStarts, tariffGets, batterySwaps, firmwareStatuses, streamOpens atomic.Int32
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
	if err := profile.HandleGetCertificateChainStatus(func(_ context.Context, _ *csms.Session, request v21.GetCertificateChainStatusRequest) (v21.GetCertificateChainStatusConfirmation, error) {
		if len(request.CertificateStatusRequests) != 1 {
			t.Errorf("certificate status requests = %d", len(request.CertificateStatusRequests))
		}
		tariffGets.Add(1)
		return v21.GetCertificateChainStatusConfirmation{CertificateStatus: []v21.GetCertificateChainStatusConfirmationCertificateStatus{{
			Source:              v21.GetCertificateChainStatusConfirmationCertificateStatusSourceEnumOCSP,
			Status:              v21.GetCertificateChainStatusConfirmationCertificateStatusEnumGood,
			NextUpdate:          "2026-07-17T00:00:00Z",
			CertificateHashData: v21.GetCertificateChainStatusConfirmationCertificateHashData{HashAlgorithm: v21.GetCertificateChainStatusConfirmationHashAlgorithmEnumSHA256, IssuerNameHash: "name", IssuerKeyHash: "key", SerialNumber: "01"},
		}}}, nil
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
	if err := profile.HandleFirmwareStatusNotification(func(_ context.Context, _ *csms.Session, request v21.FirmwareStatusNotificationRequest) (v21.FirmwareStatusNotificationConfirmation, error) {
		if request.Status != v21.FirmwareStatusNotificationRequestFirmwareStatusEnumInstalled {
			t.Errorf("firmware status = %q", request.Status)
		}
		firmwareStatuses.Add(1)
		return v21.FirmwareStatusNotificationConfirmation{}, nil
	}); err != nil {
		t.Fatal(err)
	}
	if err := profile.HandleOpenPeriodicEventStream(func(_ context.Context, _ *csms.Session, request v21.OpenPeriodicEventStreamRequest) (v21.OpenPeriodicEventStreamConfirmation, error) {
		if request.ConstantStreamData.ID != 41 {
			t.Errorf("stream ID = %d", request.ConstantStreamData.ID)
		}
		streamOpens.Add(1)
		return v21.OpenPeriodicEventStreamConfirmation{Status: v21.OpenPeriodicEventStreamConfirmationGenericStatusEnumAccepted}, nil
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
	sendCall(t, conn, "certificate-chain-status", v21.GetCertificateChainStatusRequest{CertificateStatusRequests: []v21.GetCertificateChainStatusRequestCertificateStatusRequestInfo{{
		Source:              v21.GetCertificateChainStatusRequestCertificateStatusSourceEnumOCSP,
		Urls:                []string{"https://ocsp.example.com"},
		CertificateHashData: v21.GetCertificateChainStatusRequestCertificateHashData{HashAlgorithm: v21.GetCertificateChainStatusRequestHashAlgorithmEnumSHA256, IssuerNameHash: "name", IssuerKeyHash: "key", SerialNumber: "01"},
	}}})
	if response := readMessage(t, conn); response.Type() != protocol.CallResultType {
		t.Fatalf("GetCertificateChainStatus response type = %d", response.Type())
	}
	sendCall(t, conn, "battery-in", v21.BatterySwapRequest{
		RequestID: 31, EventType: v21.BatterySwapRequestBatterySwapEventEnumBatteryIn,
		IDToken:     v21.BatterySwapRequestIDToken{IDToken: "BATTERY-USER", Type: "Local"},
		BatteryData: []v21.BatterySwapRequestBatteryData{{EVSEID: 1, SerialNumber: "BAT-001", SoC: 80, SoH: 95}},
	})
	if response := readMessage(t, conn); response.Type() != protocol.CallResultType {
		t.Fatalf("BatterySwap response type = %d", response.Type())
	}
	sendCall(t, conn, "firmware-status", v21.FirmwareStatusNotificationRequest{Status: v21.FirmwareStatusNotificationRequestFirmwareStatusEnumInstalled})
	if response := readMessage(t, conn); response.Type() != protocol.CallResultType {
		t.Fatalf("FirmwareStatusNotification response type = %d", response.Type())
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

	sendCall(t, conn, "open-stream", v21.OpenPeriodicEventStreamRequest{ConstantStreamData: v21.OpenPeriodicEventStreamRequestConstantStreamData{
		ID: 41, VariableMonitoringID: 410, Params: v21.OpenPeriodicEventStreamRequestPeriodicEventStreamParams{},
	}})
	if response := readMessage(t, conn); response.Type() != protocol.CallResultType {
		t.Fatalf("OpenPeriodicEventStream response type = %d", response.Type())
	}
	interval := 10
	streamResult := make(chan v21.AdjustPeriodicEventStreamConfirmation, 1)
	streamError := make(chan error, 1)
	go func() {
		confirmation, err := profile.CallAdjustPeriodicEventStream(context.Background(), session, v21.AdjustPeriodicEventStreamRequest{
			ID: 41, Params: v21.AdjustPeriodicEventStreamRequestPeriodicEventStreamParams{Interval: &interval},
		})
		if err != nil {
			streamError <- err
			return
		}
		streamResult <- confirmation
	}()
	outbound = readMessage(t, conn).(protocol.Call)
	if outbound.Action != "AdjustPeriodicEventStream" {
		t.Fatalf("outbound action = %q", outbound.Action)
	}
	sendCallResult(t, conn, outbound.ID, v21.AdjustPeriodicEventStreamConfirmation{Status: v21.AdjustPeriodicEventStreamConfirmationGenericStatusEnumAccepted})
	select {
	case err := <-streamError:
		t.Fatal(err)
	case confirmation := <-streamResult:
		if confirmation.Status != v21.AdjustPeriodicEventStreamConfirmationGenericStatusEnumAccepted {
			t.Fatalf("status = %q", confirmation.Status)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for AdjustPeriodicEventStream confirmation")
	}

	tariffsResult := make(chan v21.GetTariffsConfirmation, 1)
	certificateError := make(chan error, 1)
	go func() {
		confirmation, err := profile.CallGetTariffs(context.Background(), session, v21.GetTariffsRequest{EVSEID: 1})
		if err != nil {
			certificateError <- err
			return
		}
		tariffsResult <- confirmation
	}()
	outbound = readMessage(t, conn).(protocol.Call)
	if outbound.Action != "GetTariffs" {
		t.Fatalf("outbound action = %q", outbound.Action)
	}
	sendCallResult(t, conn, outbound.ID, v21.GetTariffsConfirmation{Status: v21.GetTariffsConfirmationTariffGetStatusEnumNoTariff})
	select {
	case err := <-certificateError:
		t.Fatal(err)
	case confirmation := <-tariffsResult:
		if confirmation.Status != v21.GetTariffsConfirmationTariffGetStatusEnumNoTariff {
			t.Fatalf("tariff status = %#v", confirmation)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for GetTariffs confirmation")
	}
	if transactions.Load() != 3 || reports.Load() != 1 || clearedLimits.Load() != 1 || schedulePulls.Load() != 1 || derStarts.Load() != 1 || tariffGets.Load() != 1 || batterySwaps.Load() != 1 || firmwareStatuses.Load() != 1 || streamOpens.Load() != 1 {
		t.Fatalf("transactions=%d reports=%d clearedLimits=%d schedulePulls=%d derStarts=%d tariffGets=%d batterySwaps=%d firmwareStatuses=%d streamOpens=%d", transactions.Load(), reports.Load(), clearedLimits.Load(), schedulePulls.Load(), derStarts.Load(), tariffGets.Load(), batterySwaps.Load(), firmwareStatuses.Load(), streamOpens.Load())
	}
}

func TestNotifyPeriodicEventStreamUsesUnconfirmedSend(t *testing.T) {
	router := csms.NewRouter()
	profile, err := NewProfile(router)
	if err != nil {
		t.Fatal(err)
	}
	if err := profile.HandleBootNotification(func(context.Context, *csms.Session, v21.BootNotificationRequest) (v21.BootNotificationConfirmation, error) {
		return v21.BootNotificationConfirmation{CurrentTime: "2026-07-16T00:00:00Z", Interval: 300, Status: v21.BootNotificationConfirmationRegistrationStatusEnumAccepted}, nil
	}); err != nil {
		t.Fatal(err)
	}
	received := make(chan v21.NotifyPeriodicEventStreamRequest, 1)
	if err := profile.HandleNotifyPeriodicEventStream(func(_ context.Context, _ *csms.Session, request v21.NotifyPeriodicEventStreamRequest) error {
		received <- request
		return nil
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
	conn, _, err := dialer.Dial("ws"+strings.TrimPrefix(httpServer.URL, "http")+"/CP-STREAM", nil)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	sendCall(t, conn, "boot", v21.BootNotificationRequest{
		Reason:          v21.BootNotificationRequestBootReasonEnumPowerUp,
		ChargingStation: v21.BootNotificationRequestChargingStation{Model: "AC-22K", VendorName: "Example"},
	})
	_ = readMessage(t, conn)
	payload := v21.NotifyPeriodicEventStreamRequest{
		ID: 41, Pending: 0, Basetime: "2026-07-16T00:00:01Z",
		Data: []v21.NotifyPeriodicEventStreamRequestStreamDataElement{{T: 0, V: "230.4"}},
	}
	raw, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}
	data, err := protocol.Encode(protocol.Send{ID: "send-1", Action: payload.ActionName(), Payload: raw})
	if err != nil {
		t.Fatal(err)
	}
	if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
		t.Fatal(err)
	}
	select {
	case got := <-received:
		if got.ID != 41 || len(got.Data) != 1 || got.Data[0].V != "230.4" {
			t.Fatalf("stream = %#v", got)
		}
	case <-time.After(time.Second):
		t.Fatal("SEND handler was not called")
	}
	_ = conn.SetReadDeadline(time.Now().Add(50 * time.Millisecond))
	if _, _, err := conn.ReadMessage(); err == nil {
		t.Fatal("CSMS acknowledged an unconfirmed SEND")
	}
}

func TestCorrectedDirectionAndExtendedControlRoundTrips(t *testing.T) {
	router := csms.NewRouter()
	profile, err := NewProfile(router)
	if err != nil {
		t.Fatal(err)
	}
	if err := profile.HandleBootNotification(func(context.Context, *csms.Session, v21.BootNotificationRequest) (v21.BootNotificationConfirmation, error) {
		return v21.BootNotificationConfirmation{CurrentTime: "2026-07-16T00:00:00Z", Interval: 300, Status: v21.BootNotificationConfirmationRegistrationStatusEnumAccepted}, nil
	}); err != nil {
		t.Fatal(err)
	}
	closed := make(chan int, 1)
	if err := profile.HandleClosePeriodicEventStream(func(_ context.Context, _ *csms.Session, request v21.ClosePeriodicEventStreamRequest) (v21.ClosePeriodicEventStreamConfirmation, error) {
		closed <- request.ID
		return v21.ClosePeriodicEventStreamConfirmation{}, nil
	}); err != nil {
		t.Fatal(err)
	}
	derReports := make(chan int, 1)
	if err := profile.HandleReportDERControl(func(_ context.Context, _ *csms.Session, request v21.ReportDERControlRequest) (v21.ReportDERControlConfirmation, error) {
		derReports <- request.RequestID
		return v21.ReportDERControlConfirmation{}, nil
	}); err != nil {
		t.Fatal(err)
	}
	derAlarms := make(chan v21.NotifyDERAlarmRequestDERControlEnum, 1)
	if err := profile.HandleNotifyDERAlarm(func(_ context.Context, _ *csms.Session, request v21.NotifyDERAlarmRequest) (v21.NotifyDERAlarmConfirmation, error) {
		derAlarms <- request.ControlType
		return v21.NotifyDERAlarmConfirmation{}, nil
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
	conn, _, err := dialer.Dial("ws"+strings.TrimPrefix(httpServer.URL, "http")+"/CP-21-controls", nil)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	sendCall(t, conn, "boot", v21.BootNotificationRequest{
		Reason:          v21.BootNotificationRequestBootReasonEnumPowerUp,
		ChargingStation: v21.BootNotificationRequestChargingStation{Model: "AC-22K", VendorName: "Example"},
	})
	if response := readMessage(t, conn); response.Type() != protocol.CallResultType {
		t.Fatalf("BootNotification response type = %d", response.Type())
	}
	session, ok := server.Session("CP-21-controls")
	if !ok {
		t.Fatal("session not found")
	}

	assertOutboundRoundTrip(t, conn, "NotifyAllowedEnergyTransfer", func() (v21.NotifyAllowedEnergyTransferConfirmation, error) {
		return profile.CallNotifyAllowedEnergyTransfer(context.Background(), session, v21.NotifyAllowedEnergyTransferRequest{
			TransactionID: "TX-21", AllowedEnergyTransfer: []v21.NotifyAllowedEnergyTransferRequestEnergyTransferModeEnum{v21.NotifyAllowedEnergyTransferRequestEnergyTransferModeEnumACBPT},
		})
	}, v21.NotifyAllowedEnergyTransferConfirmation{Status: v21.NotifyAllowedEnergyTransferConfirmationNotifyAllowedEnergyTransferStatusEnumAccepted})

	assertOutboundRoundTrip(t, conn, "NotifyWebPaymentStarted", func() (v21.NotifyWebPaymentStartedConfirmation, error) {
		return profile.CallNotifyWebPaymentStarted(context.Background(), session, v21.NotifyWebPaymentStartedRequest{EVSEID: 1, Timeout: 60})
	}, v21.NotifyWebPaymentStartedConfirmation{})

	limit := 22.0
	assertOutboundRoundTrip(t, conn, "UpdateDynamicSchedule", func() (v21.UpdateDynamicScheduleConfirmation, error) {
		return profile.CallUpdateDynamicSchedule(context.Background(), session, v21.UpdateDynamicScheduleRequest{
			ChargingProfileID: 21, ScheduleUpdate: v21.UpdateDynamicScheduleRequestChargingScheduleUpdate{Limit: &limit},
		})
	}, v21.UpdateDynamicScheduleConfirmation{Status: v21.UpdateDynamicScheduleConfirmationChargingProfileStatusEnumAccepted})

	assertOutboundRoundTrip(t, conn, "AFRRSignal", func() (v21.AFRRSignalConfirmation, error) {
		return profile.CallAFRRSignal(context.Background(), session, v21.AFRRSignalRequest{Timestamp: "2026-07-16T00:01:00Z", Signal: 10})
	}, v21.AFRRSignalConfirmation{Status: v21.AFRRSignalConfirmationGenericStatusEnumAccepted})

	assertOutboundRoundTrip(t, conn, "SetDERControl", func() (v21.SetDERControlConfirmation, error) {
		return profile.CallSetDERControl(context.Background(), session, v21.SetDERControlRequest{
			ControlID: "DER-gradient", ControlType: v21.SetDERControlRequestDERControlEnumGradients,
			Gradient: &v21.SetDERControlRequestGradient{Priority: 1, Gradient: 2, SoftGradient: 1},
		})
	}, v21.SetDERControlConfirmation{Status: v21.SetDERControlConfirmationDERControlStatusEnumAccepted})

	assertOutboundRoundTrip(t, conn, "GetDERControl", func() (v21.GetDERControlConfirmation, error) {
		return profile.CallGetDERControl(context.Background(), session, v21.GetDERControlRequest{RequestID: 22})
	}, v21.GetDERControlConfirmation{Status: v21.GetDERControlConfirmationDERControlStatusEnumAccepted})

	assertOutboundRoundTrip(t, conn, "GetPeriodicEventStream", func() (v21.GetPeriodicEventStreamConfirmation, error) {
		return profile.CallGetPeriodicEventStream(context.Background(), session, v21.GetPeriodicEventStreamRequest{})
	}, v21.GetPeriodicEventStreamConfirmation{})

	sendCall(t, conn, "close-stream", v21.ClosePeriodicEventStreamRequest{ID: 41})
	if response := readMessage(t, conn); response.Type() != protocol.CallResultType {
		t.Fatalf("ClosePeriodicEventStream response type = %d", response.Type())
	}
	select {
	case id := <-closed:
		if id != 41 {
			t.Fatalf("closed stream ID = %d", id)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for ClosePeriodicEventStream handler")
	}

	sendCall(t, conn, "report-der", v21.ReportDERControlRequest{RequestID: 23})
	if response := readMessage(t, conn); response.Type() != protocol.CallResultType {
		t.Fatalf("ReportDERControl response type = %d", response.Type())
	}
	select {
	case requestID := <-derReports:
		if requestID != 23 {
			t.Fatalf("DER report request ID = %d", requestID)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for ReportDERControl handler")
	}

	sendCall(t, conn, "der-alarm", v21.NotifyDERAlarmRequest{
		ControlType: v21.NotifyDERAlarmRequestDERControlEnumGradients, Timestamp: "2026-07-16T00:02:00Z",
	})
	if response := readMessage(t, conn); response.Type() != protocol.CallResultType {
		t.Fatalf("NotifyDERAlarm response type = %d", response.Type())
	}
	select {
	case controlType := <-derAlarms:
		if controlType != v21.NotifyDERAlarmRequestDERControlEnumGradients {
			t.Fatalf("DER alarm control type = %q", controlType)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for NotifyDERAlarm handler")
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

func assertOutboundRoundTrip[Confirmation protocol.Payload](
	t *testing.T,
	conn *websocket.Conn,
	action string,
	invoke func() (Confirmation, error),
	response Confirmation,
) {
	t.Helper()
	result := make(chan Confirmation, 1)
	errResult := make(chan error, 1)
	go func() {
		confirmation, err := invoke()
		if err != nil {
			errResult <- err
			return
		}
		result <- confirmation
	}()
	outbound, ok := readMessage(t, conn).(protocol.Call)
	if !ok {
		t.Fatal("outbound message is not CALL")
	}
	if outbound.Action != action {
		t.Fatalf("outbound action = %q, want %q", outbound.Action, action)
	}
	sendCallResult(t, conn, outbound.ID, response)
	select {
	case err := <-errResult:
		t.Fatal(err)
	case <-result:
	case <-time.After(time.Second):
		t.Fatalf("timed out waiting for %s confirmation", action)
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
