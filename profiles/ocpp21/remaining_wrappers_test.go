package ocpp21

import (
	"context"
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/seanlee0923/ocpp/csms"
	"github.com/seanlee0923/ocpp/protocol"
	"github.com/seanlee0923/ocpp/v21"
)

func TestRemainingSmartChargingWrapperRoundTrips(t *testing.T) {
	router := csms.NewRouter()
	profile, err := NewProfile(router)
	if err != nil {
		t.Fatal(err)
	}
	registerAcceptedBoot(t, profile)
	received := make(chan string, 5)
	if err := profile.HandleNotifyEVChargingNeeds(func(context.Context, *csms.Session, v21.NotifyEVChargingNeedsRequest) (v21.NotifyEVChargingNeedsConfirmation, error) {
		received <- "NotifyEVChargingNeeds"
		return v21.NotifyEVChargingNeedsConfirmation{Status: v21.NotifyEVChargingNeedsConfirmationNotifyEVChargingNeedsStatusEnumAccepted}, nil
	}); err != nil {
		t.Fatal(err)
	}
	if err := profile.HandleNotifyEVChargingSchedule(func(context.Context, *csms.Session, v21.NotifyEVChargingScheduleRequest) (v21.NotifyEVChargingScheduleConfirmation, error) {
		received <- "NotifyEVChargingSchedule"
		return v21.NotifyEVChargingScheduleConfirmation{Status: v21.NotifyEVChargingScheduleConfirmationGenericStatusEnumAccepted}, nil
	}); err != nil {
		t.Fatal(err)
	}
	if err := profile.HandleNotifyChargingLimit(func(context.Context, *csms.Session, v21.NotifyChargingLimitRequest) (v21.NotifyChargingLimitConfirmation, error) {
		received <- "NotifyChargingLimit"
		return v21.NotifyChargingLimitConfirmation{}, nil
	}); err != nil {
		t.Fatal(err)
	}
	if err := profile.HandleReportChargingProfiles(func(context.Context, *csms.Session, v21.ReportChargingProfilesRequest) (v21.ReportChargingProfilesConfirmation, error) {
		received <- "ReportChargingProfiles"
		return v21.ReportChargingProfilesConfirmation{}, nil
	}); err != nil {
		t.Fatal(err)
	}
	if err := profile.HandleNotifyPriorityCharging(func(context.Context, *csms.Session, v21.NotifyPriorityChargingRequest) (v21.NotifyPriorityChargingConfirmation, error) {
		received <- "NotifyPriorityCharging"
		return v21.NotifyPriorityChargingConfirmation{}, nil
	}); err != nil {
		t.Fatal(err)
	}

	server, conn := startBootedProfile(t, router, profile, "CP-21-smart")
	session, ok := server.Session("CP-21-smart")
	if !ok {
		t.Fatal("session not found")
	}

	inbound := []struct {
		action  string
		payload string
	}{
		{"NotifyEVChargingNeeds", `{"evseId":1,"chargingNeeds":{"requestedEnergyTransfer":"AC_single_phase"}}`},
		{"NotifyEVChargingSchedule", `{"timeBase":"2026-07-16T00:00:00Z","chargingSchedule":{"id":1,"chargingRateUnit":"W","chargingSchedulePeriod":[{"startPeriod":0,"limit":11000}]} ,"evseId":1}`},
		{"NotifyChargingLimit", `{"chargingLimit":{"chargingLimitSource":"CSO"}}`},
		{"ReportChargingProfiles", `{"requestId":31,"chargingLimitSource":"CSO","chargingProfile":[{"id":1,"stackLevel":0,"chargingProfilePurpose":"TxDefaultProfile","chargingProfileKind":"Absolute","chargingSchedule":[{"id":1,"chargingRateUnit":"W","chargingSchedulePeriod":[{"startPeriod":0,"limit":11000}]}]}],"evseId":1}`},
		{"NotifyPriorityCharging", `{"transactionId":"TX-21","activated":true}`},
	}
	for i, test := range inbound {
		sendRawCall(t, conn, test.action+string(rune('0'+i)), test.action, test.payload)
		if response := readMessage(t, conn); response.Type() != protocol.CallResultType {
			t.Fatalf("%s response type = %d", test.action, response.Type())
		}
		select {
		case action := <-received:
			if action != test.action {
				t.Fatalf("handler action = %q, want %q", action, test.action)
			}
		case <-time.After(time.Second):
			t.Fatalf("timed out waiting for %s handler", test.action)
		}
	}

	setRequest := mustPayload[v21.SetChargingProfileRequest](t, `{"evseId":1,"chargingProfile":{"id":32,"stackLevel":0,"chargingProfilePurpose":"TxDefaultProfile","chargingProfileKind":"Absolute","chargingSchedule":[{"id":1,"chargingRateUnit":"W","chargingSchedulePeriod":[{"startPeriod":0,"limit":11000}]}]}}`)
	assertOutboundRoundTrip(t, conn, "SetChargingProfile", func() (v21.SetChargingProfileConfirmation, error) {
		return profile.CallSetChargingProfile(context.Background(), session, setRequest)
	}, v21.SetChargingProfileConfirmation{Status: v21.SetChargingProfileConfirmationChargingProfileStatusEnumAccepted})

	profileID := 32
	assertOutboundRoundTrip(t, conn, "ClearChargingProfile", func() (v21.ClearChargingProfileConfirmation, error) {
		return profile.CallClearChargingProfile(context.Background(), session, v21.ClearChargingProfileRequest{ChargingProfileID: &profileID})
	}, v21.ClearChargingProfileConfirmation{Status: v21.ClearChargingProfileConfirmationClearChargingProfileStatusEnumAccepted})
}

func TestRemainingPaymentCertificateAndFirmwareWrapperRoundTrips(t *testing.T) {
	router := csms.NewRouter()
	profile, err := NewProfile(router)
	if err != nil {
		t.Fatal(err)
	}
	registerAcceptedBoot(t, profile)
	received := make(chan string, 4)
	if err := profile.HandleNotifySettlement(func(context.Context, *csms.Session, v21.NotifySettlementRequest) (v21.NotifySettlementConfirmation, error) {
		received <- "NotifySettlement"
		return v21.NotifySettlementConfirmation{}, nil
	}); err != nil {
		t.Fatal(err)
	}
	if err := profile.HandleVatNumberValidation(func(_ context.Context, _ *csms.Session, request v21.VatNumberValidationRequest) (v21.VatNumberValidationConfirmation, error) {
		received <- "VatNumberValidation"
		return v21.VatNumberValidationConfirmation{VatNumber: request.VatNumber, Status: v21.VatNumberValidationConfirmationGenericStatusEnumAccepted}, nil
	}); err != nil {
		t.Fatal(err)
	}
	if err := profile.HandleSignCertificate(func(context.Context, *csms.Session, v21.SignCertificateRequest) (v21.SignCertificateConfirmation, error) {
		received <- "SignCertificate"
		return v21.SignCertificateConfirmation{Status: v21.SignCertificateConfirmationGenericStatusEnumAccepted}, nil
	}); err != nil {
		t.Fatal(err)
	}
	if err := profile.HandleGet15118EVCertificate(func(context.Context, *csms.Session, v21.Get15118EVCertificateRequest) (v21.Get15118EVCertificateConfirmation, error) {
		received <- "Get15118EVCertificate"
		return v21.Get15118EVCertificateConfirmation{Status: v21.Get15118EVCertificateConfirmationIso15118EVCertificateStatusEnumAccepted, ExiResponse: "response"}, nil
	}); err != nil {
		t.Fatal(err)
	}

	server, conn := startBootedProfile(t, router, profile, "CP-21-security")
	session, ok := server.Session("CP-21-security")
	if !ok {
		t.Fatal("session not found")
	}

	inbound := []struct {
		action  string
		payload string
	}{
		{"NotifySettlement", `{"pspRef":"PSP-1","status":"Settled","settlementAmount":12.5,"settlementTime":"2026-07-16T00:00:00Z"}`},
		{"VatNumberValidation", `{"vatNumber":"DE123456789"}`},
		{"SignCertificate", `{"csr":"certificate-request"}`},
		{"Get15118EVCertificate", `{"iso15118SchemaVersion":"urn:iso:15118:2:2013:MsgDef","action":"Install","exiRequest":"request"}`},
	}
	for i, test := range inbound {
		sendRawCall(t, conn, test.action+string(rune('0'+i)), test.action, test.payload)
		if response := readMessage(t, conn); response.Type() != protocol.CallResultType {
			t.Fatalf("%s response type = %d", test.action, response.Type())
		}
		select {
		case action := <-received:
			if action != test.action {
				t.Fatalf("handler action = %q, want %q", action, test.action)
			}
		case <-time.After(time.Second):
			t.Fatalf("timed out waiting for %s handler", test.action)
		}
	}

	assertOutboundRoundTrip(t, conn, "CertificateSigned", func() (v21.CertificateSignedConfirmation, error) {
		return profile.CallCertificateSigned(context.Background(), session, v21.CertificateSignedRequest{CertificateChain: "certificate-chain"})
	}, v21.CertificateSignedConfirmation{Status: v21.CertificateSignedConfirmationCertificateSignedStatusEnumAccepted})
	assertOutboundRoundTrip(t, conn, "InstallCertificate", func() (v21.InstallCertificateConfirmation, error) {
		return profile.CallInstallCertificate(context.Background(), session, v21.InstallCertificateRequest{CertificateType: v21.InstallCertificateRequestInstallCertificateUseEnumCSMSRootCertificate, Certificate: "certificate"})
	}, v21.InstallCertificateConfirmation{Status: v21.InstallCertificateConfirmationInstallCertificateStatusEnumAccepted})
	assertOutboundRoundTrip(t, conn, "DeleteCertificate", func() (v21.DeleteCertificateConfirmation, error) {
		return profile.CallDeleteCertificate(context.Background(), session, v21.DeleteCertificateRequest{CertificateHashData: v21.DeleteCertificateRequestCertificateHashData{
			HashAlgorithm: v21.DeleteCertificateRequestHashAlgorithmEnumSHA256, IssuerNameHash: "name", IssuerKeyHash: "key", SerialNumber: "01",
		}})
	}, v21.DeleteCertificateConfirmation{Status: v21.DeleteCertificateConfirmationDeleteCertificateStatusEnumAccepted})
	assertOutboundRoundTrip(t, conn, "GetInstalledCertificateIds", func() (v21.GetInstalledCertificateIdsConfirmation, error) {
		return profile.CallGetInstalledCertificateIds(context.Background(), session, v21.GetInstalledCertificateIdsRequest{})
	}, v21.GetInstalledCertificateIdsConfirmation{Status: v21.GetInstalledCertificateIdsConfirmationGetInstalledCertificateStatusEnumNotFound})
	assertOutboundRoundTrip(t, conn, "UpdateFirmware", func() (v21.UpdateFirmwareConfirmation, error) {
		return profile.CallUpdateFirmware(context.Background(), session, v21.UpdateFirmwareRequest{RequestID: 41, Firmware: v21.UpdateFirmwareRequestFirmware{
			Location: "https://example.com/firmware.bin", RetrieveDateTime: "2026-07-16T01:00:00Z",
		}})
	}, v21.UpdateFirmwareConfirmation{Status: v21.UpdateFirmwareConfirmationUpdateFirmwareStatusEnumAccepted})
	assertOutboundRoundTrip(t, conn, "PublishFirmware", func() (v21.PublishFirmwareConfirmation, error) {
		return profile.CallPublishFirmware(context.Background(), session, v21.PublishFirmwareRequest{Location: "https://example.com/firmware.bin", Checksum: "sha256", RequestID: 42})
	}, v21.PublishFirmwareConfirmation{Status: v21.PublishFirmwareConfirmationGenericStatusEnumAccepted})
	assertOutboundRoundTrip(t, conn, "UnpublishFirmware", func() (v21.UnpublishFirmwareConfirmation, error) {
		return profile.CallUnpublishFirmware(context.Background(), session, v21.UnpublishFirmwareRequest{Checksum: "sha256"})
	}, v21.UnpublishFirmwareConfirmation{Status: v21.UnpublishFirmwareConfirmationUnpublishFirmwareStatusEnumUnpublished})
}

func TestProfileRegistrationStateIsReleasedAfterDisconnect(t *testing.T) {
	router := csms.NewRouter()
	profile, err := NewProfile(router)
	if err != nil {
		t.Fatal(err)
	}
	registerAcceptedBoot(t, profile)
	server, conn := startBootedProfile(t, router, profile, "CP-21-state-release")
	session := mustSession(t, server, "CP-21-state-release")
	if _, ok := profile.states.Load(session); !ok {
		t.Fatal("registration state was not created")
	}
	if err := conn.UnderlyingConn().Close(); err != nil {
		t.Fatal(err)
	}
	select {
	case <-session.Done():
	case <-time.After(time.Second):
		t.Fatal("session did not close")
	}
	deadline := time.Now().Add(time.Second)
	for time.Now().Before(deadline) {
		if _, ok := profile.states.Load(session); !ok {
			return
		}
		time.Sleep(time.Millisecond)
	}
	t.Fatal("profile registration state retained the closed session")
}

func registerAcceptedBoot(t *testing.T, profile *Profile) {
	t.Helper()
	if err := profile.HandleBootNotification(func(context.Context, *csms.Session, v21.BootNotificationRequest) (v21.BootNotificationConfirmation, error) {
		return v21.BootNotificationConfirmation{CurrentTime: "2026-07-16T00:00:00Z", Interval: 300, Status: v21.BootNotificationConfirmationRegistrationStatusEnumAccepted}, nil
	}); err != nil {
		t.Fatal(err)
	}
}

func startBootedProfile(t *testing.T, router *csms.Router, profile *Profile, identity string) (*csms.Server, *websocket.Conn) {
	t.Helper()
	server, err := csms.New(csms.Config{Router: router, Versions: []protocol.Version{protocol.OCPP21}})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	t.Cleanup(httpServer.Close)
	dialer := websocket.Dialer{Subprotocols: []string{string(protocol.OCPP21)}}
	conn, _, err := dialer.Dial("ws"+strings.TrimPrefix(httpServer.URL, "http")+"/"+identity, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = conn.Close() })
	sendCall(t, conn, "boot", v21.BootNotificationRequest{
		Reason:          v21.BootNotificationRequestBootReasonEnumPowerUp,
		ChargingStation: v21.BootNotificationRequestChargingStation{Model: "AC-22K", VendorName: "Example"},
	})
	if response := readMessage(t, conn); response.Type() != protocol.CallResultType {
		t.Fatalf("BootNotification response type = %d", response.Type())
	}
	if !profile.IsBooted(mustSession(t, server, identity)) {
		t.Fatal("session was not marked booted")
	}
	return server, conn
}

func mustSession(t *testing.T, server *csms.Server, identity string) *csms.Session {
	t.Helper()
	session, ok := server.Session(identity)
	if !ok {
		t.Fatal("session not found")
	}
	return session
}

func sendRawCall(t *testing.T, conn *websocket.Conn, id, action, payload string) {
	t.Helper()
	data, err := protocol.Encode(protocol.Call{ID: id, Action: action, Payload: json.RawMessage(payload)})
	if err != nil {
		t.Fatal(err)
	}
	if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
		t.Fatal(err)
	}
}

func mustPayload[T any](t *testing.T, raw string) T {
	t.Helper()
	var payload T
	if err := json.Unmarshal([]byte(raw), &payload); err != nil {
		t.Fatal(err)
	}
	return payload
}
