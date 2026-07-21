package csms

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"unicode/utf8"

	"github.com/seanlee0923/ocpp/protocol"
)

var (
	ErrInvalidHandlerRegistration = errors.New("invalid OCPP handler registration")
	ErrHandlerAlreadyRegistered   = errors.New("OCPP handler is already registered")
)

type Handler func(context.Context, *Session, json.RawMessage) (any, error)

// AfterHandler runs after a CALLRESULT has been written successfully. The
// request is the original CALL payload and confirmation is the value returned
// by the main handler. Its error is diagnostic only because the peer has
// already received the successful response.
type AfterHandler func(context.Context, *Session, json.RawMessage, any) error

type Middleware func(protocol.Version, string, Handler) Handler

// messageKind distinguishes a CALL-answering registration from a SEND
// (fire-and-forget) one, so a CALL frame can never reach a SEND-only
// handler (whose result is meaningless as a CALLRESULT) and a SEND frame
// can never reach a CALL handler (silently discarding its confirmation).
type messageKind int

const (
	callKind messageKind = iota
	sendKind
)

type routerEntry struct {
	handler Handler
	kind    messageKind
}

type Router struct {
	mu         sync.RWMutex
	handlers   map[protocol.Version]map[string]routerEntry
	after      map[protocol.Version]map[string][]AfterHandler
	hasAfter   atomic.Bool
	middleware []Middleware
}

// Use appends middleware that is applied to every registered handler at
// lookup time. Middleware is executed in registration order.
func (r *Router) Use(middleware Middleware) {
	if r == nil || middleware == nil {
		return
	}
	r.mu.Lock()
	r.middleware = append(r.middleware, middleware)
	r.mu.Unlock()
}

func NewRouter() *Router {
	return &Router{handlers: make(map[protocol.Version]map[string]routerEntry)}
}

// Handle registers a new CALL-answering handler for version and action. A
// registration is immutable: duplicate keys are rejected instead of
// silently replacing the existing handler. Only a CALL frame can reach a
// handler registered this way — use HandleSend to register a handler for
// SEND (fire-and-forget) frames instead.
func (r *Router) Handle(version protocol.Version, action string, handler Handler) error {
	return r.register(version, action, handler, callKind)
}

// HandleSend registers a new SEND (fire-and-forget) handler for version and
// action, following the same immutability rule as Handle. Only a SEND frame
// can reach a handler registered this way.
func (r *Router) HandleSend(version protocol.Version, action string, handler Handler) error {
	return r.register(version, action, handler, sendKind)
}

// HandleAfter appends a callback that runs after a successful CALLRESULT write.
// An action may have multiple callbacks; they run in registration order.
func (r *Router) HandleAfter(version protocol.Version, action string, handler AfterHandler) error {
	if r == nil {
		return fmt.Errorf("%w: router is nil", ErrInvalidHandlerRegistration)
	}
	if !version.Valid() {
		return fmt.Errorf("%w: unsupported OCPP version %q", ErrInvalidHandlerRegistration, version)
	}
	if utf8.RuneCountInString(action) == 0 || utf8.RuneCountInString(action) > protocol.MaxActionLength {
		return fmt.Errorf("%w: action must contain 1 to %d characters", ErrInvalidHandlerRegistration, protocol.MaxActionLength)
	}
	if handler == nil {
		return fmt.Errorf("%w: handler is nil", ErrInvalidHandlerRegistration)
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.after == nil {
		r.after = make(map[protocol.Version]map[string][]AfterHandler)
	}
	if r.after[version] == nil {
		r.after[version] = make(map[string][]AfterHandler)
	}
	r.after[version][action] = append(r.after[version][action], handler)
	r.hasAfter.Store(true)
	return nil
}

func (r *Router) register(version protocol.Version, action string, handler Handler, kind messageKind) error {
	if r == nil {
		return fmt.Errorf("%w: router is nil", ErrInvalidHandlerRegistration)
	}
	if !version.Valid() {
		return fmt.Errorf("%w: unsupported OCPP version %q", ErrInvalidHandlerRegistration, version)
	}
	if utf8.RuneCountInString(action) == 0 || utf8.RuneCountInString(action) > protocol.MaxActionLength {
		return fmt.Errorf("%w: action must contain 1 to %d characters", ErrInvalidHandlerRegistration, protocol.MaxActionLength)
	}
	if handler == nil {
		return fmt.Errorf("%w: handler is nil", ErrInvalidHandlerRegistration)
	}
	r.mu.Lock()
	if r.handlers == nil {
		r.handlers = make(map[protocol.Version]map[string]routerEntry)
	}
	defer r.mu.Unlock()
	if r.handlers[version] == nil {
		r.handlers[version] = make(map[string]routerEntry)
	}
	if _, exists := r.handlers[version][action]; exists {
		return fmt.Errorf("%w for %s %s", ErrHandlerAlreadyRegistered, version, action)
	}
	r.handlers[version][action] = routerEntry{handler: handler, kind: kind}
	return nil
}

// lookup finds the handler registered for version and action, but only if
// it was registered for the requested kind (CALL vs SEND) — a CALL frame
// can never dispatch to a SEND-only handler and vice versa.
func (r *Router) lookup(version protocol.Version, action string, kind messageKind) (Handler, bool) {
	if r == nil {
		return nil, false
	}
	r.mu.RLock()
	entry, ok := r.handlers[version][action]
	if !ok || entry.kind != kind {
		r.mu.RUnlock()
		return nil, false
	}
	handler := entry.handler
	middleware := append([]Middleware(nil), r.middleware...)
	r.mu.RUnlock()
	for index := len(middleware) - 1; index >= 0; index-- {
		handler = middleware[index](version, action, handler)
		if handler == nil {
			return nil, false
		}
	}
	return handler, true
}

func (r *Router) lookupAfter(version protocol.Version, action string) []AfterHandler {
	if r == nil || !r.hasAfter.Load() {
		return nil
	}
	r.mu.RLock()
	handlers := append([]AfterHandler(nil), r.after[version][action]...)
	r.mu.RUnlock()
	return handlers
}
