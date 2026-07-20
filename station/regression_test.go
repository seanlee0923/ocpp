package station_test

import (
	"context"
	"encoding/json"
	"errors"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"syscall"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/seanlee0923/ocpp/csms"
	"github.com/seanlee0923/ocpp/protocol"
	"github.com/seanlee0923/ocpp/station"
	"github.com/seanlee0923/ocpp/v16"
	"github.com/seanlee0923/ocpp/v201"
)

// TestStationHandlerPanicIsRecovered proves a panicking Handle callback does
// not crash the Station: the CSMS observes an InternalError CALLERROR and
// the Station keeps working afterward.
func TestStationHandlerPanicIsRecovered(t *testing.T) {
	connected := make(chan *csms.Session, 1)
	server, err := csms.New(csms.Config{
		Versions:  []protocol.Version{protocol.OCPP16},
		OnConnect: func(session *csms.Session) { connected <- session },
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()

	st := dialStation(t, httpServer.URL, "CP-001", nil)
	if err := station.Handle(st, func(context.Context, v16.ResetRequest) (v16.ResetConfirmation, error) {
		panic("boom")
	}); err != nil {
		t.Fatal(err)
	}

	session := <-connected
	_, err = csms.Call[v16.ResetRequest, v16.ResetConfirmation](
		context.Background(), session, v16.ResetRequest{Type: v16.ResetRequestTypeSoft},
	)
	var remote *csms.RemoteCallError
	if !errors.As(err, &remote) || remote.Code != "InternalError" {
		t.Fatalf("error = %#v, want a remote InternalError CALLERROR", err)
	}

	if st.State() != station.Connected {
		t.Fatalf("state after handler panic = %v, want Connected (Station must survive a panicking handler)", st.State())
	}
}

// TestStationHandlerErrorDoesNotLeakDetails proves a plain (non-*CallError)
// handler error never reaches the CSMS as-is: only a fixed "internal
// error" description crosses the wire, matching how csms's handleCall
// normalizes an unrecognized handler error instead of relaying its
// Error() string (which could be a DB error, a file path, an internal
// hostname, ...).
func TestStationHandlerErrorDoesNotLeakDetails(t *testing.T) {
	connected := make(chan *csms.Session, 1)
	server, err := csms.New(csms.Config{
		Versions:  []protocol.Version{protocol.OCPP16},
		OnConnect: func(session *csms.Session) { connected <- session },
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()

	st := dialStation(t, httpServer.URL, "CP-001", nil)
	sensitive := "connect to postgres://internal-db.corp.example:5432 failed: password authentication failed for user \"admin\""
	if err := station.Handle(st, func(context.Context, v16.ResetRequest) (v16.ResetConfirmation, error) {
		return v16.ResetConfirmation{}, errors.New(sensitive)
	}); err != nil {
		t.Fatal(err)
	}

	session := <-connected
	_, err = csms.Call[v16.ResetRequest, v16.ResetConfirmation](
		context.Background(), session, v16.ResetRequest{Type: v16.ResetRequestTypeSoft},
	)
	var remote *csms.RemoteCallError
	if !errors.As(err, &remote) {
		t.Fatalf("error = %v, want a RemoteCallError", err)
	}
	if remote.Code != csms.InternalError || remote.Description != "internal error" {
		t.Fatalf("CALLERROR = %+v, want Code=InternalError Description=\"internal error\"", remote)
	}
	if strings.Contains(remote.Description, "postgres") || strings.Contains(remote.Description, "admin") {
		t.Fatalf("CALLERROR description leaked handler error detail: %q", remote.Description)
	}
}

// TestStationCallErrorIsNormalizedBeforeSend proves a handler-returned
// *CallError that violates OCPP-J's wire constraints (an empty code, an
// overlong description, non-object details) still reaches the CSMS as a
// well-formed CALLERROR instead of silently producing no response at all
// (protocol.Encode would otherwise reject it, and dispatch's conn.send
// call has nowhere else to report that failure).
func TestStationCallErrorIsNormalizedBeforeSend(t *testing.T) {
	connected := make(chan *csms.Session, 1)
	server, err := csms.New(csms.Config{
		Versions:  []protocol.Version{protocol.OCPP16},
		OnConnect: func(session *csms.Session) { connected <- session },
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()

	st := dialStation(t, httpServer.URL, "CP-001", nil)
	overlongDescription := strings.Repeat("x", 300) // over protocol.MaxErrorDescriptionLength (255)
	if err := station.Handle(st, func(context.Context, v16.ResetRequest) (v16.ResetConfirmation, error) {
		return v16.ResetConfirmation{}, &station.CallError{
			Code: "", Description: overlongDescription, Details: json.RawMessage(`[1,2,3]`),
		}
	}); err != nil {
		t.Fatal(err)
	}

	session := <-connected
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err = csms.Call[v16.ResetRequest, v16.ResetConfirmation](ctx, session, v16.ResetRequest{Type: v16.ResetRequestTypeSoft})
	var remote *csms.RemoteCallError
	if !errors.As(err, &remote) {
		t.Fatalf("error = %v, want a RemoteCallError (not a timeout from no response at all)", err)
	}
	if remote.Code != csms.InternalError {
		t.Fatalf("CALLERROR code = %q, want InternalError (empty code should fall back)", remote.Code)
	}
	if len(remote.Description) > 255 {
		t.Fatalf("CALLERROR description length = %d runes, want <= 255", len([]rune(remote.Description)))
	}
}

// TestStationCallErrorPassesThroughAlreadyValidFields proves
// normalizeCallError leaves an already wire-valid *CallError's fields
// untouched: a description short enough to need no truncating, and nil
// Details (which becomes "{}", not some other placeholder) — the
// complement of TestStationCallErrorIsNormalizedBeforeSend, which only
// exercises the fields that need fixing.
func TestStationCallErrorPassesThroughAlreadyValidFields(t *testing.T) {
	connected := make(chan *csms.Session, 1)
	server, err := csms.New(csms.Config{
		Versions:  []protocol.Version{protocol.OCPP16},
		OnConnect: func(session *csms.Session) { connected <- session },
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()

	st := dialStation(t, httpServer.URL, "CP-001", nil)
	if err := station.Handle(st, func(context.Context, v16.ResetRequest) (v16.ResetConfirmation, error) {
		return v16.ResetConfirmation{}, &station.CallError{Code: "NotSupported", Description: "short"}
	}); err != nil {
		t.Fatal(err)
	}

	session := <-connected
	_, err = csms.Call[v16.ResetRequest, v16.ResetConfirmation](
		context.Background(), session, v16.ResetRequest{Type: v16.ResetRequestTypeSoft},
	)
	var remote *csms.RemoteCallError
	if !errors.As(err, &remote) {
		t.Fatalf("error = %v, want a RemoteCallError", err)
	}
	if remote.Code != "NotSupported" || remote.Description != "short" {
		t.Fatalf("CALLERROR = %+v, want an already-valid Code/Description passed through unchanged", remote)
	}
	if string(remote.Details) != "{}" {
		t.Fatalf("CALLERROR details = %s, want {} for nil Details", remote.Details)
	}
}

func TestNewHandleRejectsDuplicateRegistration(t *testing.T) {
	st, err := station.New(station.Config{URL: "ws://localhost:8080", Identity: "CP-001", Version: protocol.OCPP16})
	if err != nil {
		t.Fatal(err)
	}
	ok := func(context.Context, v16.ResetRequest) (v16.ResetConfirmation, error) {
		return v16.ResetConfirmation{Status: v16.ResetConfirmationStatusAccepted}, nil
	}
	if err := station.Handle(st, ok); err != nil {
		t.Fatal(err)
	}
	if err := station.Handle(st, ok); !errors.Is(err, station.ErrHandlerAlreadyRegistered) {
		t.Fatalf("second Handle for the same action = %v, want ErrHandlerAlreadyRegistered", err)
	}
}

func TestCallAndHandleRejectNilStation(t *testing.T) {
	if _, err := station.Call[v16.ResetRequest, v16.ResetConfirmation](
		context.Background(), nil, v16.ResetRequest{Type: v16.ResetRequestTypeSoft},
	); err == nil {
		t.Fatal("Call with a nil Station succeeded, want an error")
	}
	if err := station.Handle[v16.ResetRequest, v16.ResetConfirmation](nil, func(context.Context, v16.ResetRequest) (v16.ResetConfirmation, error) {
		return v16.ResetConfirmation{}, nil
	}); err == nil {
		t.Fatal("Handle with a nil Station succeeded, want an error")
	}
}

func TestHandleRejectsNilHandler(t *testing.T) {
	st, err := station.New(station.Config{URL: "ws://localhost:8080", Identity: "CP-001", Version: protocol.OCPP16})
	if err != nil {
		t.Fatal(err)
	}
	if err := station.Handle[v16.ResetRequest, v16.ResetConfirmation](st, nil); err == nil {
		t.Fatal("Handle with a nil handler succeeded, want an error")
	}
}

// TestHandleRejectsConfirmationVersionMismatch proves Handle rejects a
// Request/Confirmation pair that share an action name ("Reset" exists on
// both OCPP 1.6 and 2.0.1) but belong to different OCPP versions.
func TestHandleRejectsConfirmationVersionMismatch(t *testing.T) {
	st, err := station.New(station.Config{URL: "ws://localhost:8080", Identity: "CP-001", Version: protocol.OCPP16})
	if err != nil {
		t.Fatal(err)
	}
	err = station.Handle[v16.ResetRequest, v201.ResetConfirmation](st, func(context.Context, v16.ResetRequest) (v201.ResetConfirmation, error) {
		return v201.ResetConfirmation{}, nil
	})
	if err == nil {
		t.Fatal("Handle with mismatched request/confirmation OCPP versions succeeded, want an error")
	}
}

// TestCallRejectsMismatchedPayloadDirection proves Call rejects a type
// parameter pair swapped the wrong way round (a confirmation type used as
// the request, and vice versa) instead of silently sending a confirmation
// payload as an outbound CALL.
func TestCallRejectsMismatchedPayloadDirection(t *testing.T) {
	st, err := station.New(station.Config{URL: "ws://localhost:8080", Identity: "CP-001", Version: protocol.OCPP16})
	if err != nil {
		t.Fatal(err)
	}
	_, err = station.Call[v16.ResetConfirmation, v16.ResetRequest](
		context.Background(), st, v16.ResetConfirmation{Status: v16.ResetConfirmationStatusAccepted},
	)
	if err == nil || !strings.Contains(err.Error(), "requires a request followed by a confirmation") {
		t.Fatalf("error = %v, want a direction-mismatch error", err)
	}
}

// TestCallRejectsVersionMismatchWithStation proves Call rejects a request
// whose OCPP version doesn't match the Station's configured version, rather
// than sending a CALL the CSMS never negotiated support for.
func TestCallRejectsVersionMismatchWithStation(t *testing.T) {
	st, err := station.New(station.Config{URL: "ws://localhost:8080", Identity: "CP-001", Version: protocol.OCPP16})
	if err != nil {
		t.Fatal(err)
	}
	_, err = station.Call[v201.ResetRequest, v201.ResetConfirmation](
		context.Background(), st, v201.ResetRequest{Type: v201.ResetRequestResetEnumImmediate},
	)
	if err == nil || !strings.Contains(err.Error(), "does not match station version") {
		t.Fatalf("error = %v, want a station-version-mismatch error", err)
	}
}

// TestCallRejectsMismatchedRequestConfirmationAction proves Call rejects a
// Request/Confirmation pair that share an OCPP version but belong to
// different actions (Reset vs Heartbeat), instead of waiting on a CALLRESULT
// the CSMS will never send under that action name.
func TestCallRejectsMismatchedRequestConfirmationAction(t *testing.T) {
	st, err := station.New(station.Config{URL: "ws://localhost:8080", Identity: "CP-001", Version: protocol.OCPP16})
	if err != nil {
		t.Fatal(err)
	}
	_, err = station.Call[v16.ResetRequest, v16.HeartbeatConfirmation](
		context.Background(), st, v16.ResetRequest{Type: v16.ResetRequestTypeSoft},
	)
	if err == nil || !strings.Contains(err.Error(), "metadata do not match") {
		t.Fatalf("error = %v, want a request/confirmation metadata mismatch error", err)
	}
}

// TestCallRejectsInvalidOutgoingRequest proves Call validates the outgoing
// request against its own Schema before ever touching the connection, so a
// caller-constructed request with an invalid enum value fails locally
// instead of being sent to the CSMS malformed.
func TestCallRejectsInvalidOutgoingRequest(t *testing.T) {
	st, err := station.New(station.Config{URL: "ws://localhost:8080", Identity: "CP-001", Version: protocol.OCPP16})
	if err != nil {
		t.Fatal(err)
	}
	_, err = station.Call[v16.ResetRequest, v16.ResetConfirmation](
		context.Background(), st, v16.ResetRequest{}, // zero-value Type is not "Hard" or "Soft"
	)
	if err == nil || !strings.Contains(err.Error(), "validate outgoing") {
		t.Fatalf("error = %v, want a request-validation error", err)
	}
}

// TestCallAppliesCallTimeoutWhenCtxHasNoDeadline proves Config.CallTimeout
// bounds a Call whose caller-supplied ctx carries no deadline of its own,
// not just a Call whose caller already set a shorter one.
func TestCallAppliesCallTimeoutWhenCtxHasNoDeadline(t *testing.T) {
	upgrader := websocket.Upgrader{Subprotocols: []string{string(protocol.OCPP16)}}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		// Deliberately no reply: the CSMS never answers, so only
		// Config.CallTimeout (not a caller ctx deadline, which this test
		// never supplies) can end the Call.
		_ = conn
	}))
	defer server.Close()

	st, err := station.New(station.Config{
		URL: wsURL(server.URL), Identity: "CP-001", Version: protocol.OCPP16,
		CallTimeout: 100 * time.Millisecond,
	})
	if err != nil {
		t.Fatal(err)
	}
	runInBackground(t, st)
	waitConnected(t, st)

	start := time.Now()
	_, err = station.Call[v16.ResetRequest, v16.ResetConfirmation](
		context.Background(), st, v16.ResetRequest{Type: v16.ResetRequestTypeSoft},
	)
	elapsed := time.Since(start)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("error = %v, want context.DeadlineExceeded", err)
	}
	if elapsed >= 2*time.Second {
		t.Fatalf("Call took %v to time out, want it bounded by CallTimeout (100ms)", elapsed)
	}
}

// TestCallReturnsRemoteCallError proves a CALLERROR the CSMS sends in
// response to a Station-initiated Call surfaces as a *station.RemoteCallError
// carrying the wire code/description, not a generic decode failure.
func TestCallReturnsRemoteCallError(t *testing.T) {
	router := csms.NewRouter()
	if err := csms.Handle(router, func(context.Context, *csms.Session, v16.ResetRequest) (v16.ResetConfirmation, error) {
		return v16.ResetConfirmation{}, &csms.CallError{Code: csms.NotSupported, Description: "reset is disabled"}
	}); err != nil {
		t.Fatal(err)
	}
	server, err := csms.New(csms.Config{Router: router, Versions: []protocol.Version{protocol.OCPP16}})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()

	st := dialStation(t, httpServer.URL, "CP-001", nil)
	_, err = station.Call[v16.ResetRequest, v16.ResetConfirmation](
		context.Background(), st, v16.ResetRequest{Type: v16.ResetRequestTypeSoft},
	)
	var remote *station.RemoteCallError
	if !errors.As(err, &remote) {
		t.Fatalf("error = %v, want a *station.RemoteCallError", err)
	}
	if remote.Code != string(csms.NotSupported) || remote.Description != "reset is disabled" {
		t.Fatalf("RemoteCallError = %+v, want Code=%q Description=%q", remote, csms.NotSupported, "reset is disabled")
	}
}

// TestCallRejectsInvalidConfirmationFromCSMS proves Call validates a
// CALLRESULT's payload against the expected confirmation's Schema before
// decoding it, so a CSMS that answers with a payload missing a required
// field surfaces as a clear validation error instead of a zero-value
// confirmation the caller could mistake for a real one.
func TestCallRejectsInvalidConfirmationFromCSMS(t *testing.T) {
	upgrader := websocket.Upgrader{Subprotocols: []string{string(protocol.OCPP16)}}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close()
		_, data, err := conn.ReadMessage()
		if err != nil {
			return
		}
		message, err := protocol.Decode(data)
		if err != nil {
			t.Error(err)
			return
		}
		call, ok := message.(protocol.Call)
		if !ok {
			t.Errorf("message = %#v, want a CALL", message)
			return
		}
		// ResetConfirmation requires "status"; this CALLRESULT omits it.
		encoded, err := protocol.Encode(protocol.CallResult{ID: call.ID, Payload: json.RawMessage(`{}`)})
		if err != nil {
			t.Error(err)
			return
		}
		_ = conn.WriteMessage(websocket.TextMessage, encoded)
	}))
	defer server.Close()

	st := dialStation(t, server.URL, "CP-001", nil)
	_, err := station.Call[v16.ResetRequest, v16.ResetConfirmation](
		context.Background(), st, v16.ResetRequest{Type: v16.ResetRequestTypeSoft},
	)
	if err == nil || !strings.Contains(err.Error(), "invalid") {
		t.Fatalf("error = %v, want an error naming an invalid confirmation", err)
	}
}

// TestHandleRejectsMismatchedPayloadDirection proves Handle rejects a type
// parameter pair swapped the wrong way round, mirroring
// TestCallRejectsMismatchedPayloadDirection for the inbound side.
func TestHandleRejectsMismatchedPayloadDirection(t *testing.T) {
	st, err := station.New(station.Config{URL: "ws://localhost:8080", Identity: "CP-001", Version: protocol.OCPP16})
	if err != nil {
		t.Fatal(err)
	}
	err = station.Handle[v16.ResetConfirmation, v16.ResetRequest](st, func(context.Context, v16.ResetConfirmation) (v16.ResetRequest, error) {
		return v16.ResetRequest{}, nil
	})
	if err == nil || !strings.Contains(err.Error(), "requires a request followed by a confirmation") {
		t.Fatalf("error = %v, want a direction-mismatch error", err)
	}
}

// TestHandleRejectsVersionMismatchWithStation proves Handle rejects a
// request whose OCPP version doesn't match the Station's configured
// version, instead of registering a handler for an action the CSMS can
// never legally send under this Station's negotiated subprotocol.
func TestHandleRejectsVersionMismatchWithStation(t *testing.T) {
	st, err := station.New(station.Config{URL: "ws://localhost:8080", Identity: "CP-001", Version: protocol.OCPP16})
	if err != nil {
		t.Fatal(err)
	}
	err = station.Handle[v201.ResetRequest, v201.ResetConfirmation](st, func(context.Context, v201.ResetRequest) (v201.ResetConfirmation, error) {
		return v201.ResetConfirmation{}, nil
	})
	if err == nil || !strings.Contains(err.Error(), "does not match station version") {
		t.Fatalf("error = %v, want a station-version-mismatch error", err)
	}
}

// TestHandleRejectsMismatchedRequestConfirmationAction proves Handle
// rejects a Request/Confirmation pair that share an OCPP version but
// belong to different actions.
func TestHandleRejectsMismatchedRequestConfirmationAction(t *testing.T) {
	st, err := station.New(station.Config{URL: "ws://localhost:8080", Identity: "CP-001", Version: protocol.OCPP16})
	if err != nil {
		t.Fatal(err)
	}
	err = station.Handle[v16.ResetRequest, v16.HeartbeatConfirmation](st, func(context.Context, v16.ResetRequest) (v16.HeartbeatConfirmation, error) {
		return v16.HeartbeatConfirmation{}, nil
	})
	if err == nil || !strings.Contains(err.Error(), "does not match confirmation action") {
		t.Fatalf("error = %v, want a request/confirmation action mismatch error", err)
	}
}

// TestStationRejectsMalformedInboundCallPayload proves Handle validates an
// inbound CALL's payload against the request's Schema before decoding it,
// so a CSMS that sends a payload missing a required field never reaches the
// registered handler and instead produces a CALLERROR.
func TestStationRejectsMalformedInboundCallPayload(t *testing.T) {
	upgrader := websocket.Upgrader{Subprotocols: []string{string(protocol.OCPP16)}}
	verified := make(chan struct{})
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		// Deliberately no conn.Close(): closing right after reading the
		// response would race Run's own return against the assertions
		// below, the same pitfall TestStationRejectsBinaryFrames avoids by
		// leaving the connection open past this handler returning.
		defer close(verified)
		// ResetRequest requires "type"; this CALL omits it.
		encoded, err := protocol.Encode(protocol.Call{ID: "1", Action: "Reset", Payload: json.RawMessage(`{}`)})
		if err != nil {
			t.Error(err)
			return
		}
		if err := conn.WriteMessage(websocket.TextMessage, encoded); err != nil {
			return
		}
		_, data, err := conn.ReadMessage()
		if err != nil {
			t.Errorf("reading the station's response: %v", err)
			return
		}
		message, err := protocol.Decode(data)
		if err != nil {
			t.Error(err)
			return
		}
		callError, ok := message.(protocol.CallError)
		if !ok || callError.ID != "1" {
			t.Errorf("response = %#v, want a CALLERROR for id 1", message)
		}
	}))
	defer server.Close()

	dispatched := make(chan struct{}, 1)
	st, err := station.New(station.Config{URL: wsURL(server.URL), Identity: "CP-001", Version: protocol.OCPP16})
	if err != nil {
		t.Fatal(err)
	}
	if err := station.Handle(st, func(context.Context, v16.ResetRequest) (v16.ResetConfirmation, error) {
		dispatched <- struct{}{}
		return v16.ResetConfirmation{Status: v16.ResetConfirmationStatusAccepted}, nil
	}); err != nil {
		t.Fatal(err)
	}
	runInBackground(t, st)
	waitConnected(t, st)

	select {
	case <-verified:
	case <-time.After(2 * time.Second):
		t.Fatal("server never received the station's CALLERROR response")
	}
	select {
	case <-dispatched:
		t.Fatal("station dispatched a CALL whose payload fails schema validation")
	case <-time.After(100 * time.Millisecond):
	}
}

// TestStationRejectsInvalidConfirmationFromHandler proves Handle validates
// a handler's returned confirmation against its Schema before sending it,
// so a handler bug that produces an invalid confirmation reaches the CSMS
// as an InternalError CALLERROR instead of a malformed CALLRESULT.
func TestStationRejectsInvalidConfirmationFromHandler(t *testing.T) {
	connected := make(chan *csms.Session, 1)
	server, err := csms.New(csms.Config{
		Versions:  []protocol.Version{protocol.OCPP16},
		OnConnect: func(session *csms.Session) { connected <- session },
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()

	st := dialStation(t, httpServer.URL, "CP-001", nil)
	if err := station.Handle(st, func(context.Context, v16.ResetRequest) (v16.ResetConfirmation, error) {
		return v16.ResetConfirmation{}, nil // zero-value Status is not "Accepted" or "Rejected"
	}); err != nil {
		t.Fatal(err)
	}

	session := <-connected
	_, err = csms.Call[v16.ResetRequest, v16.ResetConfirmation](
		context.Background(), session, v16.ResetRequest{Type: v16.ResetRequestTypeSoft},
	)
	var remote *csms.RemoteCallError
	if !errors.As(err, &remote) {
		t.Fatalf("error = %v, want a RemoteCallError", err)
	}
	if remote.Code != csms.InternalError {
		t.Fatalf("CALLERROR code = %q, want InternalError", remote.Code)
	}
}

func TestCallRejectsGeneratorPanicAndInvalidID(t *testing.T) {
	connected := make(chan *csms.Session, 1)
	server, err := csms.New(csms.Config{
		Versions:  []protocol.Version{protocol.OCPP16},
		OnConnect: func(session *csms.Session) { connected <- session },
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()

	st, err := station.New(station.Config{
		URL: wsURL(httpServer.URL), Identity: "CP-001", Version: protocol.OCPP16,
		UniqueIDGenerator: func() string { panic("generator failed") },
	})
	if err != nil {
		t.Fatal(err)
	}
	ctx := runInBackground(t, st)
	waitConnected(t, st)

	_, err = station.Call[v16.ResetRequest, v16.ResetConfirmation](ctx, st, v16.ResetRequest{Type: v16.ResetRequestTypeSoft})
	if !errors.Is(err, station.ErrUniqueIDGeneration) {
		t.Fatalf("panicking generator error = %v, want ErrUniqueIDGeneration", err)
	}

	empty, err := station.New(station.Config{
		URL: wsURL(httpServer.URL), Identity: "CP-002", Version: protocol.OCPP16,
		UniqueIDGenerator: func() string { return "" },
	})
	if err != nil {
		t.Fatal(err)
	}
	ctx = runInBackground(t, empty)
	waitConnected(t, empty)
	_, err = station.Call[v16.ResetRequest, v16.ResetConfirmation](ctx, empty, v16.ResetRequest{Type: v16.ResetRequestTypeSoft})
	if !errors.Is(err, station.ErrUniqueIDGeneration) {
		t.Fatalf("empty-string generator error = %v, want ErrUniqueIDGeneration", err)
	}
}

// TestStationHandlerContextCanceledOnDisconnect proves a Handle callback's
// ctx is canceled when the underlying connection drops, instead of running
// with context.Background() forever.
func TestStationHandlerContextCanceledOnDisconnect(t *testing.T) {
	connected := make(chan *csms.Session, 1)
	server, err := csms.New(csms.Config{
		Versions:  []protocol.Version{protocol.OCPP16},
		OnConnect: func(session *csms.Session) { connected <- session },
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()

	st := dialStation(t, httpServer.URL, "CP-001", nil)
	handlerEntered := make(chan struct{})
	canceled := make(chan struct{})
	if err := station.Handle(st, func(ctx context.Context, _ v16.ResetRequest) (v16.ResetConfirmation, error) {
		close(handlerEntered)
		<-ctx.Done()
		close(canceled)
		return v16.ResetConfirmation{Status: v16.ResetConfirmationStatusAccepted}, ctx.Err()
	}); err != nil {
		t.Fatal(err)
	}

	session := <-connected
	go func() {
		_, _ = csms.Call[v16.ResetRequest, v16.ResetConfirmation](
			context.Background(), session, v16.ResetRequest{Type: v16.ResetRequestTypeSoft},
		)
	}()

	select {
	case <-handlerEntered:
	case <-time.After(2 * time.Second):
		t.Fatal("handler was never invoked")
	}

	if err := session.Close(); err != nil {
		t.Fatal(err)
	}

	select {
	case <-canceled:
	case <-time.After(2 * time.Second):
		t.Fatal("handler ctx was not canceled after the connection closed")
	}
}

// TestStationHandlerContextCanceledOnCallTimeout proves an inbound
// handler's ctx is bounded by Config.CallTimeout on its own — a handler
// that never returns cannot occupy its handlerSlots slot for the whole
// connection's lifetime, matching csms.Session.startHandler's per-call
// bound (csms/session.go). Before this fix, dispatch only ever gave
// handlers conn.ctx, so nothing canceled a handler short of the whole
// connection closing.
func TestStationHandlerContextCanceledOnCallTimeout(t *testing.T) {
	connected := make(chan *csms.Session, 1)
	server, err := csms.New(csms.Config{
		Versions:  []protocol.Version{protocol.OCPP16},
		OnConnect: func(session *csms.Session) { connected <- session },
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()

	st, err := station.New(station.Config{
		URL: wsURL(httpServer.URL), Identity: "CP-001", Version: protocol.OCPP16,
		CallTimeout: 100 * time.Millisecond,
	})
	if err != nil {
		t.Fatal(err)
	}
	runInBackground(t, st)
	waitConnected(t, st)

	canceled := make(chan struct{})
	if err := station.Handle(st, func(ctx context.Context, _ v16.ResetRequest) (v16.ResetConfirmation, error) {
		<-ctx.Done()
		close(canceled)
		return v16.ResetConfirmation{Status: v16.ResetConfirmationStatusAccepted}, ctx.Err()
	}); err != nil {
		t.Fatal(err)
	}

	session := <-connected
	go func() {
		_, _ = csms.Call[v16.ResetRequest, v16.ResetConfirmation](
			context.Background(), session, v16.ResetRequest{Type: v16.ResetRequestTypeSoft},
		)
	}()

	select {
	case <-canceled:
	case <-time.After(2 * time.Second):
		t.Fatal("handler ctx was not canceled by CallTimeout")
	}
	// The connection itself must still be up: this was a per-call timeout,
	// not a connection-level one.
	if st.State() != station.Connected {
		t.Fatalf("state after CallTimeout = %v, want Connected", st.State())
	}
}

// TestStationRejectsCallsOnceHandlerQueueIsFull proves dispatch's admission
// bound (Config.MaxQueuedHandlers) rejects an overflow CALL with an
// immediate CALLERROR instead of piling up an unbounded number of
// goroutines waiting for a free handler slot.
func TestStationRejectsCallsOnceHandlerQueueIsFull(t *testing.T) {
	connected := make(chan *csms.Session, 1)
	server, err := csms.New(csms.Config{
		Versions:  []protocol.Version{protocol.OCPP16},
		OnConnect: func(session *csms.Session) { connected <- session },
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()

	st, err := station.New(station.Config{
		URL: wsURL(httpServer.URL), Identity: "CP-001", Version: protocol.OCPP16,
		MaxConcurrentHandlers: 1, MaxQueuedHandlers: 2,
	})
	if err != nil {
		t.Fatal(err)
	}
	runInBackground(t, st)
	waitConnected(t, st)

	block := make(chan struct{})
	defer close(block)
	if err := station.Handle(st, func(context.Context, v16.ResetRequest) (v16.ResetConfirmation, error) {
		<-block
		return v16.ResetConfirmation{Status: v16.ResetConfirmationStatusAccepted}, nil
	}); err != nil {
		t.Fatal(err)
	}

	session := <-connected
	// MaxQueuedHandlers=2 caps how many CALLs may be admitted at once,
	// running or waiting for the sole MaxConcurrentHandlers=1 slot. The 1st
	// call occupies that slot (blocked on <-block), the 2nd fills the
	// remaining admission slot (also blocked, waiting for the handler
	// slot), and the 3rd must be rejected immediately since admission is
	// already fully spoken for.
	results := make(chan error, 3)
	for range 3 {
		go func() {
			_, err := csms.Call[v16.ResetRequest, v16.ResetConfirmation](
				context.Background(), session, v16.ResetRequest{Type: v16.ResetRequestTypeSoft},
			)
			results <- err
		}()
	}

	select {
	case err := <-results:
		var remoteErr *csms.RemoteCallError
		if !errors.As(err, &remoteErr) {
			t.Fatalf("error = %v, want a RemoteCallError", err)
		}
		if remoteErr.Code != csms.GenericError {
			t.Fatalf("CALLERROR code = %q, want %q", remoteErr.Code, csms.GenericError)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("expected one call to be rejected immediately, but no result arrived in time")
	}

	// The other two calls must still be blocked (one running, one queued),
	// not also rejected or somehow completed early.
	select {
	case err := <-results:
		t.Fatalf("a second call resolved (err=%v) before the blocking handler was released", err)
	case <-time.After(100 * time.Millisecond):
	}
}

// TestStationCallWriteHonorsContextTimeout proves Call's outbound write is
// actually bounded by a deadline (Config.WriteTimeout, or ctx's own
// deadline if sooner), not just the ctx-cancellation check in Call's
// select — a context deadline alone does nothing to stop a blocking
// WriteMessage call, only an actual socket write deadline does. The CSMS
// side here completes the WebSocket handshake normally but then never
// reads another byte, and the station sends an oversized payload, so the
// write genuinely backs up at the TCP level instead of just simulating a
// slow handler.
func TestStationCallWriteHonorsContextTimeout(t *testing.T) {
	// SO_RCVBUF is set on the *listening* socket, before any handshake
	// happens, not after Accept: setting it post-accept (via
	// (*net.TCPConn).SetReadBuffer on the already-established connection,
	// tried first here) turned out not to reliably shrink the receive
	// window on every platform/kernel — the window scaling gorilla's TCP
	// handshake negotiates can already be locked in by then, so a large
	// enough write can still complete without ever blocking (observed:
	// reliably reproduced the write-blocks-until-WriteTimeout condition on
	// macOS, but not on Linux CI, where the write went through immediately
	// and the test only caught the outer 1s ctx deadline instead of the
	// inner 200ms WriteTimeout it's meant to prove). Setting it on the
	// listening socket via ListenConfig.Control applies before the
	// three-way handshake, so the small window is actually in effect from
	// the start on both platforms.
	lc := net.ListenConfig{Control: func(_, _ string, c syscall.RawConn) error {
		var sockErr error
		if err := c.Control(func(fd uintptr) {
			sockErr = syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_RCVBUF, 1024)
		}); err != nil {
			return err
		}
		return sockErr
	}}
	listener, err := lc.Listen(context.Background(), "tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}

	upgrader := websocket.Upgrader{}
	accepted := make(chan *websocket.Conn, 1)
	testServer := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		accepted <- conn
		// Deliberately no read loop: the connection accepts the handshake
		// and then goes silent, so nothing ever drains what the station
		// writes.
	}))
	testServer.Listener.Close()
	testServer.Listener = listener
	testServer.Start()
	defer testServer.Close()

	st, err := station.New(station.Config{
		URL: wsURL(testServer.URL), Identity: "CP-001", Version: protocol.OCPP16,
		WriteTimeout: 200 * time.Millisecond,
	})
	if err != nil {
		t.Fatal(err)
	}
	runInBackground(t, st)
	waitConnected(t, st)

	conn := <-accepted
	defer conn.Close()

	large := strings.Repeat("x", 8<<20) // 8MB: comfortably exceeds even a generously auto-tuned receive window.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	start := time.Now()
	_, err = station.Call[v16.DataTransferRequest, v16.DataTransferConfirmation](
		ctx, st, v16.DataTransferRequest{VendorID: "test", Data: &large},
	)
	elapsed := time.Since(start)
	if err == nil {
		t.Fatal("Call succeeded even though the CSMS never read anything off the wire")
	}
	// A generous margin above WriteTimeout (200ms) but comfortably below
	// ctx's own 10s: wide enough to absorb a slow/shared CI runner
	// (marshaling 8MB under -race in particular) without either shrinking
	// to the point of flaking on slower machines or growing so wide that
	// it stops actually distinguishing "bounded by WriteTimeout" from
	// "just eventually hit the outer ctx deadline".
	if elapsed >= 5*time.Second {
		t.Fatalf("Call took %v to return, want it bounded by WriteTimeout (200ms), well under ctx's own 10s deadline", elapsed)
	}
}

func TestStationBoundsConcurrentHandlers(t *testing.T) {
	const limit = 2

	connected := make(chan *csms.Session, 1)
	server, err := csms.New(csms.Config{
		Versions:  []protocol.Version{protocol.OCPP16},
		OnConnect: func(session *csms.Session) { connected <- session },
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()

	st, err := station.New(station.Config{
		URL: wsURL(httpServer.URL), Identity: "CP-001", Version: protocol.OCPP16, MaxConcurrentHandlers: limit,
	})
	if err != nil {
		t.Fatal(err)
	}
	runInBackground(t, st)
	waitConnected(t, st)

	var inFlight, maxObserved atomic.Int64
	block := make(chan struct{})
	defer close(block)
	if err := station.Handle(st, func(context.Context, v16.ResetRequest) (v16.ResetConfirmation, error) {
		current := inFlight.Add(1)
		for {
			observed := maxObserved.Load()
			if current <= observed || maxObserved.CompareAndSwap(observed, current) {
				break
			}
		}
		<-block
		inFlight.Add(-1)
		return v16.ResetConfirmation{Status: v16.ResetConfirmationStatusAccepted}, nil
	}); err != nil {
		t.Fatal(err)
	}

	session := <-connected
	const calls = limit + 3
	for range calls {
		go func() {
			_, _ = csms.Call[v16.ResetRequest, v16.ResetConfirmation](
				context.Background(), session, v16.ResetRequest{Type: v16.ResetRequestTypeSoft},
			)
		}()
	}

	deadline := time.Now().Add(2 * time.Second)
	for inFlight.Load() < limit && time.Now().Before(deadline) {
		time.Sleep(time.Millisecond)
	}
	// Give any over-limit dispatch a chance to (incorrectly) start before asserting.
	time.Sleep(50 * time.Millisecond)
	if got := maxObserved.Load(); got > limit {
		t.Fatalf("max concurrent handlers observed = %d, want <= %d", got, limit)
	}
}

// TestStationStopDuringBackoffReturnsErrStopped proves Stop() called while
// Run is blocked in its reconnect-backoff wait makes Run return ErrStopped,
// not nil (nil would happen if wait() only checked ctx.Err(), which is nil
// when only Stop, not ctx, ended the wait).
func TestStationStopDuringBackoffReturnsErrStopped(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			_ = conn.Close() // fail the handshake immediately so dial() errors fast
		}
	}()

	st, err := station.New(station.Config{
		URL: "ws://" + listener.Addr().String(), Identity: "CP-001", Version: protocol.OCPP16,
		ReconnectPolicy: &station.ReconnectPolicy{InitialDelay: time.Second, MaxDelay: time.Second, Multiplier: 1},
	})
	if err != nil {
		t.Fatal(err)
	}
	runDone := make(chan error, 1)
	go func() { runDone <- st.Run(context.Background()) }()

	// Give the first dial attempt time to fail and Run to enter its
	// InitialDelay=1s backoff wait before calling Stop.
	time.Sleep(100 * time.Millisecond)
	st.Stop()

	select {
	case err := <-runDone:
		if !errors.Is(err, station.ErrStopped) {
			t.Fatalf("Run after Stop during backoff wait = %v, want ErrStopped", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("Run did not return after Stop during backoff wait")
	}
}

// TestStationStopInterruptsInFlightDial proves Stop() unblocks a dial that
// is stuck in the WebSocket handshake (Config.HandshakeTimeout left at its
// zero-value default, i.e. no timeout), instead of leaving Run blocked
// until the caller separately cancels ctx or the OS-level TCP timeout
// eventually fires.
func TestStationStopInterruptsInFlightDial(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			// Accept the TCP connection but never write an HTTP response,
			// so the WebSocket handshake hangs indefinitely.
			_ = conn
		}
	}()

	st, err := station.New(station.Config{
		URL: "ws://" + listener.Addr().String(), Identity: "CP-001", Version: protocol.OCPP16,
	})
	if err != nil {
		t.Fatal(err)
	}
	runDone := make(chan error, 1)
	go func() { runDone <- st.Run(context.Background()) }()

	deadline := time.Now().Add(2 * time.Second)
	for st.State() != station.Connecting && time.Now().Before(deadline) {
		time.Sleep(time.Millisecond)
	}
	if st.State() != station.Connecting {
		t.Fatal("station never reached Connecting state")
	}
	// Give the handshake a moment to actually be in flight before stopping.
	time.Sleep(20 * time.Millisecond)

	st.Stop()
	select {
	case err := <-runDone:
		if err == nil {
			t.Fatal("Run after Stop during an in-flight dial succeeded, want an error")
		}
	case <-time.After(2 * time.Second):
		t.Fatal("Stop did not interrupt an in-flight dial; Run is still blocked")
	}
}

// TestStationEscapesIdentityInDialURL proves an Identity containing a
// character that is meaningful in a URL (a space here) round-trips
// correctly instead of producing a malformed dial request. The CSMS's
// default ValidateIdentity only accepts RFC 3986 unreserved characters, so
// this test supplies a permissive one to exercise a real-world CSMS that
// allows a broader identity charset.
func TestStationEscapesIdentityInDialURL(t *testing.T) {
	const identity = "CP 001"

	observed := make(chan string, 1)
	server, err := csms.New(csms.Config{
		Versions: []protocol.Version{protocol.OCPP16},
		Security: csms.SecurityConfig{ValidateIdentity: func(string) error { return nil }},
		OnConnect: func(session *csms.Session) {
			observed <- session.Identity()
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()

	dialStation(t, httpServer.URL, identity, nil)

	select {
	case got := <-observed:
		if got != identity {
			t.Fatalf("CSMS observed identity %q, want %q", got, identity)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("CSMS never observed a connection")
	}
}

// TestStationRejectsBinaryFrames proves the Station's read loop only
// accepts WebSocket text frames, matching csms.Server's own check
// (server.go's readLoop). Before this fix, readLoop discarded the frame
// type entirely, so a CALL delivered inside a binary frame — which OCPP-J
// never sends deliberately, but which the WebSocket protocol otherwise
// allows — would have been dispatched exactly like a normal text-framed
// one instead of being treated as a protocol violation.
func TestStationRejectsBinaryFrames(t *testing.T) {
	upgrader := websocket.Upgrader{Subprotocols: []string{string(protocol.OCPP16)}}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		// Deliberately no conn.Close() and no read loop of its own: a
		// hijacked WebSocket connection outlives this handler returning,
		// so the connection stays open on its own past this point,
		// letting the assertions below tell a real rejection (an error
		// naming text messages) apart from an unrelated closed-connection
		// error that any dropped connection would also produce.
		data, err := protocol.Encode(protocol.Call{ID: "1", Action: "Reset", Payload: json.RawMessage(`{"type":"Soft"}`)})
		if err != nil {
			t.Error(err)
			return
		}
		_ = conn.WriteMessage(websocket.BinaryMessage, data)
	}))
	defer server.Close()

	dispatched := make(chan struct{}, 1)
	st, err := station.New(station.Config{URL: wsURL(server.URL), Identity: "CP-001", Version: protocol.OCPP16})
	if err != nil {
		t.Fatal(err)
	}
	if err := station.Handle(st, func(context.Context, v16.ResetRequest) (v16.ResetConfirmation, error) {
		dispatched <- struct{}{}
		return v16.ResetConfirmation{Status: v16.ResetConfirmationStatusAccepted}, nil
	}); err != nil {
		t.Fatal(err)
	}
	runDone := make(chan error, 1)
	go func() { runDone <- st.Run(context.Background()) }()
	// Only st.Stop(), not also draining runDone here: whichever branch of
	// the select below fires already accounts for runDone's one buffered
	// slot, so a second read here would block forever if the "dispatched"
	// branch fired instead (Run may not have returned, or may return
	// after this cleanup already moved on).
	t.Cleanup(st.Stop)

	select {
	case <-dispatched:
		t.Fatal("station dispatched a CALL delivered inside a binary frame")
	case err := <-runDone:
		if err == nil || !strings.Contains(err.Error(), "text messages") {
			t.Fatalf("Run error = %v, want an error naming text messages (not just any connection failure)", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("station neither dispatched the CALL nor rejected the binary frame in time")
	}
}

// TestStationDisconnectsOnMalformedFrame proves a frame protocol.Decode
// can't parse at all closes the connection (surfacing through Run's return
// and OnDisconnect) instead of being silently skipped forever — matching
// csms.Server's read loop, which always treats this class of error as
// fatal to the connection regardless of OCPP version. Before this fix,
// readLoop's `continue` meant the CSMS got no signal anything was wrong
// and the station just sat there, connected, doing nothing useful.
func TestStationDisconnectsOnMalformedFrame(t *testing.T) {
	upgrader := websocket.Upgrader{Subprotocols: []string{string(protocol.OCPP16)}}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		// Deliberately no conn.Close(): see TestStationRejectsBinaryFrames
		// for why the connection is left open past this handler returning.
		_ = conn.WriteMessage(websocket.TextMessage, []byte(`not valid json at all`))
	}))
	defer server.Close()

	st, err := station.New(station.Config{URL: wsURL(server.URL), Identity: "CP-001", Version: protocol.OCPP16})
	if err != nil {
		t.Fatal(err)
	}
	runDone := make(chan error, 1)
	go func() { runDone <- st.Run(context.Background()) }()
	t.Cleanup(st.Stop)

	select {
	case err := <-runDone:
		if err == nil || !errors.Is(err, protocol.ErrInvalidMessage) {
			t.Fatalf("Run error = %v, want protocol.ErrInvalidMessage (not just any connection failure)", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("station did not disconnect after a malformed frame")
	}
}
