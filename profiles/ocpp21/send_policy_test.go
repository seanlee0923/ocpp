package ocpp21

import (
	"context"
	"errors"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"ocpp-go/csms"
	"ocpp-go/protocol"
	"ocpp-go/v21"
)

func TestSendFailuresAreDroppedWithoutResponse(t *testing.T) {
	router := csms.NewRouter()
	profile, err := NewProfile(router)
	if err != nil {
		t.Fatal(err)
	}
	registerAcceptedBoot(t, profile)
	handled := make(chan struct{}, 1)
	if err := profile.HandleNotifyPeriodicEventStream(func(context.Context, *csms.Session, v21.NotifyPeriodicEventStreamRequest) error {
		handled <- struct{}{}
		return errors.New("application rejected stream data")
	}); err != nil {
		t.Fatal(err)
	}
	logs := make(chan csms.LogRecord, 16)
	server, err := csms.New(csms.Config{
		Router: router, Versions: []protocol.Version{protocol.OCPP21},
		Logger: csms.LoggerFunc(func(_ context.Context, record csms.LogRecord) { logs <- record }),
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()
	dialer := websocket.Dialer{Subprotocols: []string{string(protocol.OCPP21)}}
	conn, _, err := dialer.Dial("ws"+strings.TrimPrefix(httpServer.URL, "http")+"/CP-21-send-errors", nil)
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

	sendRawMessage(t, conn, `[6,"invalid","NotifyPeriodicEventStream",{}]`)
	waitForSendDrop(t, logs, "NotifyPeriodicEventStream", "handler_or_validation_error")
	select {
	case <-handled:
		t.Fatal("handler ran for schema-invalid SEND")
	default:
	}

	sendRawMessage(t, conn, `[6,"unknown","VendorSpecificSend",{}]`)
	waitForSendDrop(t, logs, "VendorSpecificSend", "action_not_registered")

	sendRawMessage(t, conn, `[6,"handler","NotifyPeriodicEventStream",{"id":1,"pending":0,"basetime":"2026-07-16T00:00:00Z","data":[{"t":0,"v":"230.4"}]}]`)
	waitForSendDrop(t, logs, "NotifyPeriodicEventStream", "handler_or_validation_error")
	select {
	case <-handled:
	case <-time.After(time.Second):
		t.Fatal("valid SEND did not reach handler")
	}

	// A regular CALL after all dropped SENDs proves that none produced a wire
	// response and the WebSocket remained usable.
	sendRawCall(t, conn, "after-send", "UnknownAction", `{}`)
	if response := readMessage(t, conn); response.Type() != protocol.CallErrorType {
		t.Fatalf("post-SEND response type = %d", response.Type())
	}
}

func sendRawMessage(t *testing.T, conn *websocket.Conn, message string) {
	t.Helper()
	if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		t.Fatal(err)
	}
}

func waitForSendDrop(t *testing.T, logs <-chan csms.LogRecord, action, reason string) {
	t.Helper()
	timer := time.NewTimer(time.Second)
	defer timer.Stop()
	for {
		select {
		case record := <-logs:
			if record.Event == csms.LogSendDropped && record.Action == action && record.Reason == reason {
				return
			}
		case <-timer.C:
			t.Fatalf("missing send.dropped log for %s (%s)", action, reason)
		}
	}
}
