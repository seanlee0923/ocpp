package csms

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/seanlee0923/ocpp/protocol"
)

// BenchmarkRouterLookup measures the per-dispatch cost of lookup, including
// the middleware slice copy and handler rewrap that happen on every inbound
// message, at increasing middleware counts.
func BenchmarkRouterLookup(b *testing.B) {
	for _, middlewareCount := range []int{0, 1, 5} {
		b.Run(middlewareCountLabel(middlewareCount), func(b *testing.B) {
			router := NewRouter()
			if err := router.Handle(protocol.OCPP16, "Heartbeat", func(context.Context, *Session, json.RawMessage) (any, error) {
				return nil, nil
			}); err != nil {
				b.Fatal(err)
			}
			for i := 0; i < middlewareCount; i++ {
				router.Use(func(_ protocol.Version, _ string, next Handler) Handler {
					return func(ctx context.Context, session *Session, raw json.RawMessage) (any, error) {
						return next(ctx, session, raw)
					}
				})
			}
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				if _, ok := router.lookup(protocol.OCPP16, "Heartbeat"); !ok {
					b.Fatal("handler not found")
				}
			}
		})
	}
}

func middlewareCountLabel(count int) string {
	switch count {
	case 0:
		return "no-middleware"
	case 1:
		return "1-middleware"
	default:
		return "5-middleware"
	}
}
