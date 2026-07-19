package station_test

import (
	"context"
	"errors"
	"net"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

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
