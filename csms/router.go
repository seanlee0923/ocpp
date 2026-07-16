package csms

import (
	"context"
	"encoding/json"
	"sync"

	"ocpp-go/protocol"
)

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

func (r *Router) Handle(version protocol.Version, action string, handler Handler) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.handlers[version] == nil {
		r.handlers[version] = make(map[string]Handler)
	}
	r.handlers[version][action] = handler
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
