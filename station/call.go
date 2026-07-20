package station

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/seanlee0923/ocpp/protocol"
)

// Call sends a generated request to the CSMS and waits for the matching
// generated confirmation or CALLERROR. It is the Station-initiated mirror
// of csms.Call.
func Call[Request protocol.Payload, Confirmation protocol.Payload](
	ctx context.Context,
	s *Station,
	request Request,
) (Confirmation, error) {
	var zero Confirmation
	if s == nil {
		return zero, fmt.Errorf("station: Station is nil")
	}
	if request.Direction() != protocol.RequestPayload || zero.Direction() != protocol.ConfirmationPayload {
		return zero, fmt.Errorf("station: Call requires a request followed by a confirmation")
	}
	if request.Version() != s.config.Version {
		return zero, fmt.Errorf("station: request version %s does not match station version %s", request.Version(), s.config.Version)
	}
	if zero.Version() != request.Version() || zero.ActionName() != request.ActionName() {
		return zero, fmt.Errorf("station: request and confirmation metadata do not match")
	}
	if err := request.Validate(); err != nil {
		return zero, fmt.Errorf("station: validate outgoing %s request: %w", request.ActionName(), err)
	}
	conn, err := s.currentConn()
	if err != nil {
		return zero, err
	}
	payload, err := json.Marshal(request)
	if err != nil {
		return zero, fmt.Errorf("station: marshal outgoing %s request: %w", request.ActionName(), err)
	}
	id, err := generateUniqueID(s.uniqueIDGen)
	if err != nil {
		return zero, err
	}
	response, err := conn.registerCall(id, s.config.MaxPendingCalls)
	if err != nil {
		return zero, err
	}
	defer conn.unregisterCall(id)

	callCtx := ctx
	cancel := func() {}
	if _, hasDeadline := ctx.Deadline(); !hasDeadline && s.config.CallTimeout > 0 {
		callCtx, cancel = context.WithTimeout(ctx, s.config.CallTimeout)
	}
	defer cancel()

	if err := conn.send(callCtx, protocol.Call{ID: id, Action: request.ActionName(), Payload: payload}); err != nil {
		return zero, err
	}

	select {
	case outcome := <-response:
		if outcome.err != nil {
			return zero, outcome.err
		}
		if err := zero.ValidateJSON(outcome.payload); err != nil {
			return zero, fmt.Errorf("station: invalid %s confirmation from CSMS: %w", request.ActionName(), err)
		}
		if err := json.Unmarshal(outcome.payload, &zero); err != nil {
			return zero, fmt.Errorf("station: decode %s confirmation: %w", request.ActionName(), err)
		}
		return zero, nil
	case <-callCtx.Done():
		return zero, callCtx.Err()
	case <-conn.closed:
		return zero, ErrNotConnected
	}
}
