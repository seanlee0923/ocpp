// This file intentionally uses package station rather than station_test: the
// regression must coordinate the private writeMu/send boundary directly.
// Exercising that boundary through the public API would require OS-dependent
// TCP backpressure and timing, which has already proved flaky on Linux CI.
package station

import (
	"context"
	"errors"
	"sync"
	"testing"

	"github.com/seanlee0923/ocpp/protocol"
)

// TestStationConnSendRechecksContextAfterWaitingForWriteLock is the
// station-side counterpart of csms's lock-wait cancellation test. conn is
// deliberately nil: the post-lock ctx check must return before any socket
// access.
func TestStationConnSendRechecksContextAfterWaitingForWriteLock(t *testing.T) {
	base, cancel := context.WithCancel(context.Background())
	ctx := &firstErrObservedContext{Context: base, observed: make(chan struct{})}
	conn := &stationConn{}
	conn.writeMu.Lock()

	result := make(chan error, 1)
	go func() {
		result <- conn.send(ctx, protocol.Call{ID: "1", Action: "Heartbeat", Payload: []byte(`{}`)})
	}()

	<-ctx.observed
	cancel()
	conn.writeMu.Unlock()
	if err := <-result; !errors.Is(err, context.Canceled) {
		t.Fatalf("send error = %v, want context.Canceled", err)
	}
}

type firstErrObservedContext struct {
	context.Context
	once     sync.Once
	observed chan struct{}
}

func (ctx *firstErrObservedContext) Err() error {
	err := ctx.Context.Err()
	ctx.once.Do(func() { close(ctx.observed) })
	return err
}
