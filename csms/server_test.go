package csms

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"ocpp-go/protocol"
	"ocpp-go/v16"
	"ocpp-go/v201"
)

func TestServerNegotiatesAndRoutes(t *testing.T) {
	router := NewRouter()
	router.Handle(protocol.OCPP201, "BootNotification", func(_ context.Context, session *Session, payload json.RawMessage) (any, error) {
		return map[string]any{"status": "Accepted"}, nil
	})
	server, err := New(Config{Router: router, Versions: []protocol.Version{protocol.OCPP201}})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()

	dialer := websocket.Dialer{Subprotocols: []string{string(protocol.OCPP201)}}
	conn, _, err := dialer.Dial("ws"+strings.TrimPrefix(httpServer.URL, "http")+"/CP-001", nil)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	if conn.Subprotocol() != string(protocol.OCPP201) {
		t.Fatalf("subprotocol = %q", conn.Subprotocol())
	}
	if err := conn.WriteMessage(websocket.TextMessage, []byte(`[2,"1","BootNotification",{}]`)); err != nil {
		t.Fatal(err)
	}
	_, response, err := conn.ReadMessage()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(response), `"Accepted"`) {
		t.Fatalf("response = %s", response)
	}
}

func TestServerRejectsSendOnPre21Connections(t *testing.T) {
	t.Run(string(protocol.OCPP16), func(t *testing.T) {
		server, err := New(Config{Versions: []protocol.Version{protocol.OCPP16}})
		if err != nil {
			t.Fatal(err)
		}
		httpServer := httptest.NewServer(server)
		defer httpServer.Close()
		conn := dialTestStation(t, httpServer.URL, protocol.OCPP16)
		defer conn.Close()
		data, err := protocol.Encode(protocol.Send{ID: "send-1", Action: "NotifyPeriodicEventStream", Payload: json.RawMessage(`{}`)})
		if err != nil {
			t.Fatal(err)
		}
		if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
			t.Fatal(err)
		}
		if err := conn.SetReadDeadline(time.Now().Add(time.Second)); err != nil {
			t.Fatal(err)
		}
		if _, _, err := conn.ReadMessage(); err == nil {
			t.Fatal("OCPP 1.6 connection remained open after SEND")
		}
	})

	t.Run(string(protocol.OCPP201), func(t *testing.T) {
		server, err := New(Config{Versions: []protocol.Version{protocol.OCPP201}})
		if err != nil {
			t.Fatal(err)
		}
		httpServer := httptest.NewServer(server)
		defer httpServer.Close()
		conn := dialTestStation(t, httpServer.URL, protocol.OCPP201)
		defer conn.Close()
		data, err := protocol.Encode(protocol.Send{ID: "send-1", Action: "NotifyPeriodicEventStream", Payload: json.RawMessage(`{}`)})
		if err != nil {
			t.Fatal(err)
		}
		if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
			t.Fatal(err)
		}
		_, response, err := conn.ReadMessage()
		if err != nil {
			t.Fatal(err)
		}
		decoded, err := protocol.Decode(response)
		if err != nil {
			t.Fatal(err)
		}
		callError, ok := decoded.(protocol.CallError)
		if !ok || callError.ID != "send-1" || callError.Code != string(MessageTypeNotSupported) {
			t.Fatalf("response = %#v, want MessageTypeNotSupported CALLERROR", decoded)
		}
	})
}

func TestServerReturnsMessageTypeNotSupportedForUnknown20xMessage(t *testing.T) {
	for _, version := range []protocol.Version{protocol.OCPP201, protocol.OCPP21} {
		t.Run(string(version), func(t *testing.T) {
			server, err := New(Config{Versions: []protocol.Version{version}})
			if err != nil {
				t.Fatal(err)
			}
			httpServer := httptest.NewServer(server)
			defer httpServer.Close()
			conn := dialTestStation(t, httpServer.URL, version)
			defer conn.Close()
			if err := conn.WriteMessage(websocket.TextMessage, []byte(`[9,"future-1",{}]`)); err != nil {
				t.Fatal(err)
			}
			_, response, err := conn.ReadMessage()
			if err != nil {
				t.Fatal(err)
			}
			decoded, err := protocol.Decode(response)
			if err != nil {
				t.Fatal(err)
			}
			callError, ok := decoded.(protocol.CallError)
			if !ok || callError.ID != "future-1" || callError.Code != string(MessageTypeNotSupported) {
				t.Fatalf("response = %#v, want MessageTypeNotSupported CALLERROR", decoded)
			}
		})
	}
}

func TestServerWritesClassifiedConstraintCallErrors(t *testing.T) {
	router := NewRouter()
	if err := Handle(router, func(context.Context, *Session, v16.BootNotificationRequest) (v16.BootNotificationConfirmation, error) {
		return validV16BootConfirmation(), nil
	}); err != nil {
		t.Fatal(err)
	}
	server, err := New(Config{Router: router, Versions: []protocol.Version{protocol.OCPP16}})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()
	conn := dialTestStation(t, httpServer.URL, protocol.OCPP16)
	defer conn.Close()

	tests := []struct {
		id      string
		payload string
		code    ErrorCode
	}{
		{"property", `{"chargePointVendor":"Example","chargePointModel":"AC-22K","unknown":true}`, PropertyConstraintViolation},
		{"occurrence", `{"chargePointVendor":"Example"}`, OccurenceConstraintViolation},
		{"type", `{"chargePointVendor":42,"chargePointModel":"AC-22K"}`, TypeConstraintViolation},
	}
	for _, test := range tests {
		message := `[2,"` + test.id + `","BootNotification",` + test.payload + `]`
		if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			t.Fatal(err)
		}
		_, data, err := conn.ReadMessage()
		if err != nil {
			t.Fatal(err)
		}
		decoded, err := protocol.Decode(data)
		if err != nil {
			t.Fatal(err)
		}
		callError, ok := decoded.(protocol.CallError)
		if !ok {
			t.Fatalf("response = %T, want CALLERROR", decoded)
		}
		if callError.ID != test.id || callError.Code != string(test.code) {
			t.Fatalf("CALLERROR = %#v, want id=%q code=%q", callError, test.id, test.code)
		}
	}
}

func TestServerRejectsInvalidAuthorizeEnumBeforeHandler(t *testing.T) {
	router := NewRouter()
	called := false
	if err := Handle(router, func(context.Context, *Session, v201.AuthorizeRequest) (v201.AuthorizeConfirmation, error) {
		called = true
		return v201.AuthorizeConfirmation{}, nil
	}); err != nil {
		t.Fatal(err)
	}
	server, err := New(Config{Router: router, Versions: []protocol.Version{protocol.OCPP201}})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()
	conn := dialTestStation(t, httpServer.URL, protocol.OCPP201)
	defer conn.Close()

	message := `[2,"authorize-invalid","Authorize",{"idToken":{"idToken":"TAG-201","type":"Hello"}}]`
	if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		t.Fatal(err)
	}
	_, response, err := conn.ReadMessage()
	if err != nil {
		t.Fatal(err)
	}
	decoded, err := protocol.Decode(response)
	if err != nil {
		t.Fatal(err)
	}
	callError, ok := decoded.(protocol.CallError)
	if !ok || callError.ID != "authorize-invalid" || callError.Code != string(PropertyConstraintViolation) {
		t.Fatalf("response = %#v, want PropertyConstraintViolation CALLERROR", decoded)
	}
	if called {
		t.Fatal("Authorize handler ran for invalid idToken.type enum")
	}
}

func TestServerNormalizesApplicationCallErrors(t *testing.T) {
	router := NewRouter()
	router.Handle(protocol.OCPP201, "InvalidCode", func(context.Context, *Session, json.RawMessage) (any, error) {
		return nil, &CallError{Code: ErrorCode("VendorError"), Description: "must not leak", Details: map[string]any{"secret": true}}
	})
	router.Handle(protocol.OCPP201, "LongDescription", func(context.Context, *Session, json.RawMessage) (any, error) {
		return nil, &CallError{Code: GenericError, Description: strings.Repeat("가", protocol.MaxErrorDescriptionLength+1)}
	})
	server, err := New(Config{Router: router, Versions: []protocol.Version{protocol.OCPP201}})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()
	conn := dialTestStation(t, httpServer.URL, protocol.OCPP201)
	defer conn.Close()

	for _, test := range []struct {
		id, action, code string
		descriptionRunes int
	}{
		{"invalid", "InvalidCode", string(InternalError), len([]rune("internal error"))},
		{"long", "LongDescription", string(GenericError), protocol.MaxErrorDescriptionLength},
	} {
		message := `[2,"` + test.id + `","` + test.action + `",{}]`
		if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			t.Fatal(err)
		}
		_, response, err := conn.ReadMessage()
		if err != nil {
			t.Fatal(err)
		}
		decoded, err := protocol.Decode(response)
		if err != nil {
			t.Fatal(err)
		}
		callError := decoded.(protocol.CallError)
		if callError.Code != test.code || len([]rune(callError.Description)) != test.descriptionRunes {
			t.Fatalf("CALLERROR = %#v", callError)
		}
	}
}

func TestServerRejectsDuplicateSessionByDefault(t *testing.T) {
	server, err := New(Config{Versions: []protocol.Version{protocol.OCPP16}})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()
	first := dialTestStation(t, httpServer.URL, protocol.OCPP16)
	defer first.Close()

	dialer := websocket.Dialer{Subprotocols: []string{string(protocol.OCPP16)}}
	second, response, err := dialer.Dial("ws"+strings.TrimPrefix(httpServer.URL, "http")+"/CP-001", nil)
	if second != nil {
		second.Close()
	}
	if err == nil {
		t.Fatal("duplicate connection succeeded")
	}
	if response == nil || response.StatusCode != http.StatusConflict {
		t.Fatalf("response = %#v, want HTTP 409", response)
	}
	if server.SessionCount() != 1 {
		t.Fatalf("session count = %d", server.SessionCount())
	}
}

func TestServerCanReplaceDuplicateSession(t *testing.T) {
	duplicateCalls := 0
	server, err := New(Config{
		Versions: []protocol.Version{protocol.OCPP16},
		OnDuplicateSession: func(_ context.Context, attempt DuplicateSessionAttempt) (DuplicateSessionDecision, error) {
			duplicateCalls++
			if attempt.Identity != "CP-001" || attempt.Existing.Identity != "CP-001" {
				t.Errorf("unexpected attempt: %#v", attempt)
			}
			return ReplaceExistingSession, nil
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()
	firstConn := dialTestStation(t, httpServer.URL, protocol.OCPP16)
	defer firstConn.Close()
	firstSession, ok := server.Session("CP-001")
	if !ok {
		t.Fatal("first session not registered")
	}

	secondConn := dialTestStation(t, httpServer.URL, protocol.OCPP16)
	defer secondConn.Close()
	select {
	case <-firstSession.Done():
	case <-time.After(time.Second):
		t.Fatal("replaced session was not closed")
	}
	if !errors.Is(firstSession.Err(), ErrSessionReplaced) {
		t.Fatalf("first session error = %v", firstSession.Err())
	}
	current, ok := server.Session("CP-001")
	if !ok || current == firstSession {
		t.Fatal("replacement session not registered")
	}
	if duplicateCalls != 1 || server.SessionCount() != 1 {
		t.Fatalf("duplicate calls = %d, sessions = %d", duplicateCalls, server.SessionCount())
	}
}

func TestFailedReplacementUpgradeKeepsExistingSession(t *testing.T) {
	server, err := New(Config{
		Versions: []protocol.Version{protocol.OCPP16},
		OnDuplicateSession: func(context.Context, DuplicateSessionAttempt) (DuplicateSessionDecision, error) {
			return ReplaceExistingSession, nil
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()
	firstConn := dialTestStation(t, httpServer.URL, protocol.OCPP16)
	defer firstConn.Close()
	firstSession, _ := server.Session("CP-001")

	request, err := http.NewRequest(http.MethodGet, httpServer.URL+"/CP-001", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Sec-WebSocket-Protocol", string(protocol.OCPP16))
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Fatal(err)
	}
	response.Body.Close()
	if response.StatusCode != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", response.StatusCode)
	}
	current, ok := server.Session("CP-001")
	if !ok || current != firstSession {
		t.Fatal("failed replacement removed the existing session")
	}
}
