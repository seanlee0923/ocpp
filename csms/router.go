package csms

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"unicode/utf8"

	"ocpp-go/protocol"
)

var ErrHandlerAlreadyRegistered = errors.New("OCPP handler is already registered")

type Handler func(context.Context, *Session, json.RawMessage) (any, error)

type Middleware func(protocol.Version, string, Handler) Handler

type Router struct {
	mu         sync.RWMutex
	handlers   map[protocol.Version]map[string]Handler
	middleware []Middleware
}

// Use appends middleware that is applied to every registered handler at
// lookup time. Middleware is executed in registration order.
func (r *Router) Use(middleware Middleware) {
	if middleware == nil {
		return
	}
	r.mu.Lock()
	r.middleware = append(r.middleware, middleware)
	r.mu.Unlock()
}

func NewRouter() *Router {
	return &Router{handlers: make(map[protocol.Version]map[string]Handler)}
}

// Handle registers a new handler for version and action. A registration is
// immutable: duplicate keys are rejected instead of silently replacing the
// existing handler.
func (r *Router) Handle(version protocol.Version, action string, handler Handler) error {
	if r == nil {
		return fmt.Errorf("router is nil")
	}
	if !version.Valid() {
		return fmt.Errorf("unsupported OCPP version %q", version)
	}
	if utf8.RuneCountInString(action) == 0 || utf8.RuneCountInString(action) > protocol.MaxActionLength {
		return fmt.Errorf("action must contain 1 to %d characters", protocol.MaxActionLength)
	}
	if handler == nil {
		return fmt.Errorf("handler is nil")
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.handlers[version] == nil {
		r.handlers[version] = make(map[string]Handler)
	}
	if _, exists := r.handlers[version][action]; exists {
		return fmt.Errorf("%w for %s %s", ErrHandlerAlreadyRegistered, version, action)
	}
	r.handlers[version][action] = handler
	return nil
}

func (r *Router) lookup(version protocol.Version, action string) (Handler, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	handler, ok := r.handlers[version][action]
	if !ok {
		return nil, false
	}
	for index := len(r.middleware) - 1; index >= 0; index-- {
		handler = r.middleware[index](version, action, handler)
	}
	return handler, true
}
