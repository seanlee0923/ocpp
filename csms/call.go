package csms

import (
	"context"
	"encoding/json"
	"fmt"

	"ocpp-go/protocol"
)

// RemoteCallError is a CALLERROR returned by a Charging Station.
type RemoteCallError struct {
	Code        string
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
	id := session.uniqueIDGenerator()
	if id == "" {
		return zero, fmt.Errorf("unique ID generator returned an empty value")
	}
	response, err := session.registerCall(id)
	if err != nil {
		return zero, err
	}
	defer session.unregisterCall(id)

	callCtx := ctx
	cancel := func() {}
	if _, hasDeadline := ctx.Deadline(); !hasDeadline && session.callTimeout > 0 {
		callCtx, cancel = context.WithTimeout(ctx, session.callTimeout)
	}
	defer cancel()
	if err := session.Send(callCtx, protocol.Call{ID: id, Action: request.ActionName(), Payload: payload}); err != nil {
		return zero, err
	}

	select {
	case outcome := <-response:
		if outcome.err != nil {
			return zero, outcome.err
		}
		if err := zero.ValidateJSON(outcome.result.Payload); err != nil {
			return zero, fmt.Errorf("invalid %s confirmation from Charging Station: %w", request.ActionName(), err)
		}
		if err := json.Unmarshal(outcome.result.Payload, &zero); err != nil {
			return zero, fmt.Errorf("decode %s confirmation: %w", request.ActionName(), err)
		}
		return zero, nil
	case <-callCtx.Done():
		return zero, callCtx.Err()
	case <-session.Done():
		if err := session.Err(); err != nil {
			return zero, err
		}
		return zero, ErrSessionClosed
	}
}
