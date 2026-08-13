package csms

import (
	"context"
	"errors"
	"net"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/seanlee0923/ocpp/protocol"
	"github.com/seanlee0923/ocpp/v16"
)

// dialStalledStation connects a station that completes the handshake and then
// never reads another byte, with a deliberately tiny receive buffer so the
// server's writes back up at the TCP level almost immediately.
//
// SO_RCVBUF is set via Dialer.Control, i.e. before connect, for the same
// reason TestStationCallWriteHonorsContextTimeout sets it on the listening
// socket rather than after Accept: the receive window is negotiated during
// the three-way handshake, so shrinking the buffer afterwards does not
// reliably shrink an already-negotiated window (it happened to work on macOS
// and not on Linux CI).
func dialStalledStation(t *testing.T, serverURL string, version protocol.Version) *websocket.Conn {
	t.Helper()
	netDialer := &net.Dialer{Control: func(_, _ string, c syscall.RawConn) error {
		var sockErr error
		if err := c.Control(func(fd uintptr) {
			sockErr = syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_RCVBUF, 1024)
		}); err != nil {
			return err
		}
		return sockErr
	}}
	dialer := websocket.Dialer{
		Subprotocols:   []string{string(version)},
		NetDialContext: netDialer.DialContext,
	}
	conn, _, err := dialer.Dial("ws"+strings.TrimPrefix(serverURL, "http")+"/CP-001", nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { conn.Close() })
	return conn
}

// TestCallReportsContextDeadlineWhenWriteTimesOut proves that a Call whose
// ctx deadline expires while the CALL frame is still being written reports
// that as context.DeadlineExceeded, not as the raw net timeout the socket
// write deadline produces.
//
// This matters twice over. Callers reasonably branch on
// errors.Is(err, context.DeadlineExceeded) to tell "my deadline was too
// short" apart from "the connection is broken", and got the wrong answer
// here. And the outbound-call metric is classified from callCtx.Err(), so
// the metric already reported MetricOutboundCallTimeout while the returned
// error disagreed — the two describing the same event differently is what
// made TestOutboundCallTimeoutEmitsMetric fail intermittently under load
// (its 20ms deadline is short enough to expire inside the write when the
// coverage-instrumented run is slow enough).
func TestCallReportsContextDeadlineWhenWriteTimesOut(t *testing.T) {
	connected := make(chan *Session, 1)
	server, err := New(Config{
		Versions:     []protocol.Version{protocol.OCPP16},
		WriteTimeout: 30 * time.Second, // far longer than ctx's, so ctx's deadline is the one that fires
		OnConnect:    func(session *Session) { connected <- session },
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := newTestHTTPServer(t, server)
	dialStalledStation(t, httpServer.URL, protocol.OCPP16)
	session := <-connected

	large := strings.Repeat("x", 1<<20) // 1MB: big enough to fill the shrunken receive window, small enough that marshaling it stays far under ctx's deadline
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err = Call[v16.DataTransferRequest, v16.DataTransferConfirmation](
		ctx, session, v16.DataTransferRequest{VendorID: "test", Data: &large},
	)
	if err == nil {
		t.Fatal("Call succeeded even though the station never read anything off the wire")
	}
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("error = %v, want it to wrap context.DeadlineExceeded", err)
	}
}

// TestCallReportsContextCancellationWhenWriteIsInterrupted is the
// cancellation half of the same contract. Canceling a ctx that carries no
// deadline cannot itself unblock a write already blocked in the kernel —
// only a socket deadline does that, here WriteTimeout — so what this pins
// down is the rule that once the write has failed, a ctx that ended for any
// reason is what gets reported, matching how the outbound-call metric is
// already classified from callCtx.Err().
func TestCallReportsContextCancellationWhenWriteIsInterrupted(t *testing.T) {
	connected := make(chan *Session, 1)
	server, err := New(Config{
		Versions:     []protocol.Version{protocol.OCPP16},
		WriteTimeout: 300 * time.Millisecond,
		OnConnect:    func(session *Session) { connected <- session },
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := newTestHTTPServer(t, server)
	dialStalledStation(t, httpServer.URL, protocol.OCPP16)
	session := <-connected

	large := strings.Repeat("x", 1<<20)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(100 * time.Millisecond) // while the write is blocked, before WriteTimeout fires
		cancel()
	}()
	defer cancel()
	_, err = Call[v16.DataTransferRequest, v16.DataTransferConfirmation](
		ctx, session, v16.DataTransferRequest{VendorID: "test", Data: &large},
	)
	if err == nil {
		t.Fatal("Call succeeded even though the station never read anything off the wire")
	}
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("error = %v, want it to wrap context.Canceled", err)
	}
}
