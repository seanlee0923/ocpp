package csms

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"sync"
	"testing"

	"ocpp-go/protocol"
	"ocpp-go/v16"
)

func TestRouterHandleRejectsInvalidRegistration(t *testing.T) {
	handler := func(context.Context, *Session, json.RawMessage) (any, error) { return nil, nil }
	tests := []struct {
		name    string
		router  *Router
		version protocol.Version
		action  string
		handler Handler
	}{
		{"nil router", nil, protocol.OCPP16, "Heartbeat", handler},
		{"unknown version", NewRouter(), protocol.Version("ocpp3.0"), "Heartbeat", handler},
		{"empty action", NewRouter(), protocol.OCPP16, "", handler},
		{"long action", NewRouter(), protocol.OCPP16, strings.Repeat("a", protocol.MaxActionLength+1), handler},
		{"nil handler", NewRouter(), protocol.OCPP16, "Heartbeat", nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := test.router.Handle(test.version, test.action, test.handler); err == nil {
				t.Fatal("Handle succeeded")
			}
		})
	}
}

func TestRouterHandleRejectsDuplicateWithoutReplacing(t *testing.T) {
	router := NewRouter()
	first := func(context.Context, *Session, json.RawMessage) (any, error) { return "first", nil }
	second := func(context.Context, *Session, json.RawMessage) (any, error) { return "second", nil }
	if err := router.Handle(protocol.OCPP16, "Heartbeat", first); err != nil {
		t.Fatal(err)
	}
	if err := router.Handle(protocol.OCPP16, "Heartbeat", second); !errors.Is(err, ErrHandlerAlreadyRegistered) {
		t.Fatalf("duplicate error = %v", err)
	}
	handler, ok := router.lookup(protocol.OCPP16, "Heartbeat")
	if !ok {
		t.Fatal("original handler not found")
	}
	result, err := handler(context.Background(), nil, json.RawMessage(`{}`))
	if err != nil || result != "first" {
		t.Fatalf("original handler result = %#v, error = %v", result, err)
	}
}

func TestRouterHandleAllowsSameActionForDifferentVersions(t *testing.T) {
	router := NewRouter()
	handler := func(context.Context, *Session, json.RawMessage) (any, error) { return nil, nil }
	if err := router.Handle(protocol.OCPP16, "Heartbeat", handler); err != nil {
		t.Fatal(err)
	}
	if err := router.Handle(protocol.OCPP201, "Heartbeat", handler); err != nil {
		t.Fatal(err)
	}
}

func TestRouterConcurrentDuplicateRegistrationHasSingleWinner(t *testing.T) {
	router := NewRouter()
	const attempts = 32
	errorsCh := make(chan error, attempts)
	var wg sync.WaitGroup
	for i := 0; i < attempts; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			errorsCh <- router.Handle(protocol.OCPP21, "Heartbeat", func(context.Context, *Session, json.RawMessage) (any, error) {
				return nil, nil
			})
		}()
	}
	wg.Wait()
	close(errorsCh)
	successes, duplicates := 0, 0
	for err := range errorsCh {
		switch {
		case err == nil:
			successes++
		case errors.Is(err, ErrHandlerAlreadyRegistered):
			duplicates++
		default:
			t.Fatalf("unexpected registration error: %v", err)
		}
	}
	if successes != 1 || duplicates != attempts-1 {
		t.Fatalf("successes = %d, duplicates = %d", successes, duplicates)
	}
}

func TestTypedHandlePropagatesDuplicateRegistration(t *testing.T) {
	router := NewRouter()
	handler := func(context.Context, *Session, v16.HeartbeatRequest) (v16.HeartbeatConfirmation, error) {
		return v16.HeartbeatConfirmation{}, nil
	}
	if err := Handle(router, handler); err != nil {
		t.Fatal(err)
	}
	if err := Handle(router, handler); !errors.Is(err, ErrHandlerAlreadyRegistered) {
		t.Fatalf("duplicate typed registration error = %v", err)
	}
}
