// Package stationtest is a minimal, typed Charging Station test double used
// only by this module's own tests to drive round trips against a csms.Server
// without hand-written raw JSON frames. It is not a Charging Station client:
// it has no reconnect/backoff, no offline queueing, and no client-side
// security credential handling, all of which only matter for a real
// deployed charger.
package stationtest

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/seanlee0923/ocpp/protocol"
)

// ErrStationClosed is returned by Call when the Station connection is
// closed while a call is pending.
var ErrStationClosed = errors.New("stationtest: station is closed")

type callOutcome struct {
	payload json.RawMessage
	err     error
}

// Station is a single Charging-Station-side WebSocket connection to a
// csms.Server, with typed outbound Call and typed inbound Handle registered
// per action.
type Station struct {
	t                 *testing.T
	conn              *websocket.Conn
	version           protocol.Version
	writeMu           sync.Mutex
	pendingMu         sync.Mutex
	pending           map[string]chan callOutcome
	handlersMu        sync.RWMutex
	handlers          map[string]func(context.Context, json.RawMessage) (json.RawMessage, error)
	uniqueIDGenerator func() string
	closed            chan struct{}
	closeOnce         sync.Once
}

// Dial connects to serverURL as identity, negotiating version as the
// WebSocket subprotocol, and starts the read loop. It registers t.Cleanup to
// close the connection, so callers do not need an explicit defer/Close.
//
// On a dial failure it reports the error via t.Errorf and returns nil —
// never t.Fatal/FailNow, which testing requires to be called only from the
// goroutine running the test function, not from a station-per-goroutine
// load test. Callers must check for a nil return.
func Dial(t *testing.T, serverURL, identity string, version protocol.Version) *Station {
	t.Helper()
	dialer := websocket.Dialer{Subprotocols: []string{string(version)}}
	conn, _, err := dialer.Dial("ws"+strings.TrimPrefix(serverURL, "http")+"/"+identity, nil)
	if err != nil {
		t.Errorf("stationtest: dial %s: %v", identity, err)
		return nil
	}
	station := &Station{
		t:                 t,
		conn:              conn,
		version:           version,
		pending:           make(map[string]chan callOutcome),
		handlers:          make(map[string]func(context.Context, json.RawMessage) (json.RawMessage, error)),
		uniqueIDGenerator: uuid.NewString,
		closed:            make(chan struct{}),
	}
	go station.readLoop()
	t.Cleanup(station.Close)
	return station
}

// Close closes the underlying WebSocket connection. Safe to call more than
// once.
func (s *Station) Close() {
	s.closeOnce.Do(func() {
		close(s.closed)
		_ = s.conn.Close()
	})
}

func (s *Station) registerCall(id string) (<-chan callOutcome, error) {
	s.pendingMu.Lock()
	defer s.pendingMu.Unlock()
	select {
	case <-s.closed:
		return nil, ErrStationClosed
	default:
	}
	response := make(chan callOutcome, 1)
	s.pending[id] = response
	return response, nil
}

func (s *Station) unregisterCall(id string) {
	s.pendingMu.Lock()
	delete(s.pending, id)
	s.pendingMu.Unlock()
}

func (s *Station) resolve(message protocol.Message) bool {
	s.pendingMu.Lock()
	response, ok := s.pending[message.UniqueID()]
	if ok {
		delete(s.pending, message.UniqueID())
	}
	s.pendingMu.Unlock()
	if !ok {
		return false
	}
	switch value := message.(type) {
	case protocol.CallResult:
		response <- callOutcome{payload: value.Payload}
	case protocol.CallError:
		response <- callOutcome{err: &RemoteCallError{Code: value.Code, Description: value.Description, Details: value.Details}}
	default:
		response <- callOutcome{err: fmt.Errorf("stationtest: unexpected response type %d", message.Type())}
	}
	return true
}

// RemoteCallError is a CALLERROR returned by the CSMS in response to a
// Station-initiated Call.
type RemoteCallError struct {
	Code        string
	Description string
	Details     json.RawMessage
}

func (e *RemoteCallError) Error() string {
	return fmt.Sprintf("stationtest: CALLERROR %s: %s", e.Code, e.Description)
}

func (s *Station) send(message protocol.Message) error {
	data, err := protocol.Encode(message)
	if err != nil {
		return err
	}
	s.writeMu.Lock()
	defer s.writeMu.Unlock()
	return s.conn.WriteMessage(websocket.TextMessage, data)
}

func (s *Station) handle(action string, handler func(context.Context, json.RawMessage) (json.RawMessage, error)) {
	s.handlersMu.Lock()
	s.handlers[action] = handler
	s.handlersMu.Unlock()
}

func (s *Station) readLoop() {
	for {
		_, data, err := s.conn.ReadMessage()
		if err != nil {
			select {
			case <-s.closed:
			default:
				s.t.Errorf("stationtest: read: %v", err)
			}
			return
		}
		message, err := protocol.Decode(data)
		if err != nil {
			s.t.Errorf("stationtest: decode: %v", err)
			continue
		}
		switch value := message.(type) {
		case protocol.CallResult, protocol.CallError:
			if !s.resolve(message) {
				s.t.Errorf("stationtest: response %q has no matching pending call", message.UniqueID())
			}
		case protocol.Call:
			s.dispatch(value)
		default:
			s.t.Errorf("stationtest: unexpected inbound message type %d", message.Type())
		}
	}
}

func (s *Station) dispatch(call protocol.Call) {
	s.handlersMu.RLock()
	handler, ok := s.handlers[call.Action]
	s.handlersMu.RUnlock()
	if !ok {
		_ = s.send(protocol.CallError{ID: call.ID, Code: "NotImplemented", Description: "no stationtest handler registered for " + call.Action, Details: json.RawMessage(`{}`)})
		return
	}
	go func() {
		payload, err := handler(context.Background(), call.Payload)
		if err != nil {
			var callErr *CallError
			if errors.As(err, &callErr) {
				details := callErr.Details
				if details == nil {
					details = json.RawMessage(`{}`)
				}
				_ = s.send(protocol.CallError{ID: call.ID, Code: callErr.Code, Description: callErr.Description, Details: details})
				return
			}
			_ = s.send(protocol.CallError{ID: call.ID, Code: "InternalError", Description: err.Error(), Details: json.RawMessage(`{}`)})
			return
		}
		_ = s.send(protocol.CallResult{ID: call.ID, Payload: payload})
	}()
}

// CallError lets a Handle handler control the CALLERROR's code and
// description precisely, instead of the default "InternalError" wrapping any
// other returned error.
type CallError struct {
	Code        string
	Description string
	Details     json.RawMessage
}

func (e *CallError) Error() string { return e.Description }
