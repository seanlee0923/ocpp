package csms

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"
	"unicode/utf8"

	"github.com/seanlee0923/ocpp/protocol"
)

// RemoteCallError is a CALLERROR returned by a Charging Station.
type RemoteCallError struct {
	Code        ErrorCode
	Description string
	Details     json.RawMessage
}

func (e *RemoteCallError) Error() string {
	return fmt.Sprintf("OCPP CALLERROR %s: %s", e.Code, e.Description)
}

// Call sends a generated request to a Charging Station and waits for the
// matching generated confirmation or CALLERROR.
func Call[Request protocol.Payload, Confirmation protocol.Payload](
	ctx context.Context,
	session *Session,
	request Request,
) (Confirmation, error) {
	var zero Confirmation
	if session == nil {
		return zero, fmt.Errorf("session is nil")
	}
	if request.Direction() != protocol.RequestPayload || zero.Direction() != protocol.ConfirmationPayload {
		return zero, fmt.Errorf("Call requires a request followed by a confirmation")
	}
	if request.Version() != session.Version() {
		return zero, fmt.Errorf("request version %s does not match session version %s", request.Version(), session.Version())
	}
	if zero.Version() != request.Version() || zero.ActionName() != request.ActionName() {
		return zero, fmt.Errorf("request and confirmation metadata do not match")
	}
	if err := request.Validate(); err != nil {
		return zero, fmt.Errorf("validate outgoing %s request: %w", request.ActionName(), err)
	}
	payload, err := json.Marshal(request)
	if err != nil {
		return zero, fmt.Errorf("marshal outgoing %s request: %w", request.ActionName(), err)
	}
	id, err := generateUniqueID(session.uniqueIDGenerator)
	if err != nil {
		return zero, err
	}
	response, err := session.registerCall(id)
	if err != nil {
		if errors.Is(err, ErrTooManyPendingCalls) {
			session.recordMetric(ctx, MetricEvent{
				Type: MetricOutboundCallRejected, Identity: session.identity, Version: session.version,
				MessageType: protocol.CallType, Action: request.ActionName(),
			})
		}
		return zero, err
	}
	defer session.unregisterCall(id)
	start := time.Now()
	metric := MetricEvent{
		Identity: session.identity, Version: session.version,
		MessageType: protocol.CallType, Action: request.ActionName(),
	}
	// finish records a terminal outcome for this call. Duration always
	// measures from just after registerCall admitted the call into the
	// pending-call table (so it includes CALL marshal/write time), not
	// strictly from MetricOutboundCallSent.
	finish := func(ctx context.Context, eventType MetricEventType) {
		metric.Type, metric.Duration = eventType, time.Since(start)
		session.recordMetric(ctx, metric)
	}

	callCtx := ctx
	cancel := func() {}
	if _, hasDeadline := ctx.Deadline(); !hasDeadline && session.callTimeout > 0 {
		callCtx, cancel = context.WithTimeout(ctx, session.callTimeout)
	}
	defer cancel()
	if err := session.send(callCtx, protocol.Call{ID: id, Action: request.ActionName(), Payload: payload}); err != nil {
		// context.Background(): this event reports why the call ended, and
		// callCtx (or even ctx itself) may already be done here, so a
		// Metrics implementation that defensively no-ops on an
		// already-canceled ctx must not silently lose it.
		finish(context.Background(), classifyContextOutcome(callCtx.Err()))
		return zero, err
	}
	metric.Type = MetricOutboundCallSent
	session.recordMetric(ctx, metric)

	select {
	case outcome := <-response:
		if outcome.err != nil {
			var remoteErr *RemoteCallError
			if errors.As(outcome.err, &remoteErr) {
				metric.ErrorCode = remoteErr.Code
			}
			finish(ctx, MetricOutboundCallFailed)
			return zero, outcome.err
		}
		if err := zero.ValidateJSON(outcome.result.Payload); err != nil {
			finish(ctx, MetricOutboundCallFailed)
			return zero, fmt.Errorf("invalid %s confirmation from Charging Station: %w", request.ActionName(), err)
		}
		if err := json.Unmarshal(outcome.result.Payload, &zero); err != nil {
			finish(ctx, MetricOutboundCallFailed)
			return zero, fmt.Errorf("decode %s confirmation: %w", request.ActionName(), err)
		}
		finish(ctx, MetricOutboundCallCompleted)
		return zero, nil
	case <-callCtx.Done():
		finish(context.Background(), classifyContextOutcome(callCtx.Err()))
		return zero, callCtx.Err()
	case <-session.Done():
		finish(context.Background(), MetricOutboundCallFailed)
		if err := session.Err(); err != nil {
			return zero, err
		}
		return zero, ErrSessionClosed
	}
}

// classifyContextOutcome maps a context's terminal error to the matching
// outbound-call metric type. A nil or non-context error (for example a
// transport write failure unrelated to ctx) yields MetricOutboundCallFailed.
func classifyContextOutcome(err error) MetricEventType {
	switch {
	case errors.Is(err, context.DeadlineExceeded):
		return MetricOutboundCallTimeout
	case errors.Is(err, context.Canceled):
		return MetricOutboundCallCanceled
	default:
		return MetricOutboundCallFailed
	}
}

func generateUniqueID(generator UniqueIDGenerator) (id string, err error) {
	defer func() {
		if recovered := recover(); recovered != nil {
			id = ""
			err = fmt.Errorf("%w: generator panic: %v", ErrUniqueIDGeneration, recovered)
		}
	}()
	if generator == nil {
		return "", fmt.Errorf("%w: generator is nil", ErrUniqueIDGeneration)
	}
	id = generator()
	if length := utf8.RuneCountInString(id); length == 0 || length > protocol.MaxUniqueIDLength {
		return "", fmt.Errorf("%w: ID must contain 1 to %d characters", ErrUniqueIDGeneration, protocol.MaxUniqueIDLength)
	}
	return id, nil
}
