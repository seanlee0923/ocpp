package stationtest

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/seanlee0923/ocpp/protocol"
)

// Call sends a generated request to the CSMS and waits for the matching
// generated confirmation or CALLERROR. It is the Station-initiated mirror of
// csms.Call.
func Call[Request protocol.Payload, Confirmation protocol.Payload](
	ctx context.Context,
	station *Station,
	request Request,
) (Confirmation, error) {
	var zero Confirmation
	if request.Direction() != protocol.RequestPayload || zero.Direction() != protocol.ConfirmationPayload {
		return zero, fmt.Errorf("stationtest: Call requires a request followed by a confirmation")
	}
	if request.Version() != station.version {
		return zero, fmt.Errorf("stationtest: request version %s does not match station version %s", request.Version(), station.version)
	}
	if zero.Version() != request.Version() || zero.ActionName() != request.ActionName() {
		return zero, fmt.Errorf("stationtest: request and confirmation metadata do not match")
	}
	if err := request.Validate(); err != nil {
		return zero, fmt.Errorf("stationtest: validate outgoing %s request: %w", request.ActionName(), err)
	}
	payload, err := json.Marshal(request)
	if err != nil {
		return zero, fmt.Errorf("stationtest: marshal outgoing %s request: %w", request.ActionName(), err)
	}
	id := station.uniqueIDGenerator()
	response, err := station.registerCall(id)
	if err != nil {
		return zero, err
	}
	defer station.unregisterCall(id)
	if err := station.send(protocol.Call{ID: id, Action: request.ActionName(), Payload: payload}); err != nil {
		return zero, err
	}

	select {
	case outcome := <-response:
		if outcome.err != nil {
			return zero, outcome.err
		}
		if err := zero.ValidateJSON(outcome.payload); err != nil {
			return zero, fmt.Errorf("stationtest: invalid %s confirmation from CSMS: %w", request.ActionName(), err)
		}
		if err := json.Unmarshal(outcome.payload, &zero); err != nil {
			return zero, fmt.Errorf("stationtest: decode %s confirmation: %w", request.ActionName(), err)
		}
		return zero, nil
	case <-ctx.Done():
		return zero, ctx.Err()
	case <-station.closed:
		return zero, ErrStationClosed
	}
}

// Handle registers a typed handler for a CSMS-initiated CALL. It is the
// Station-side mirror of csms.Handle: the inbound payload is schema-
// validated before decode, and the handler's returned confirmation is
// schema-validated before it is sent back as a CALLRESULT.
func Handle[Request protocol.Payload, Confirmation protocol.Payload](
	station *Station,
	handler func(context.Context, Request) (Confirmation, error),
) error {
	var request Request
	var confirmation Confirmation
	if request.Direction() != protocol.RequestPayload || confirmation.Direction() != protocol.ConfirmationPayload {
		return fmt.Errorf("stationtest: Handle requires a request followed by a confirmation")
	}
	if request.Version() != station.version {
		return fmt.Errorf("stationtest: request version %s does not match station version %s", request.Version(), station.version)
	}
	if request.ActionName() != confirmation.ActionName() {
		return fmt.Errorf("stationtest: request action %s does not match confirmation action %s", request.ActionName(), confirmation.ActionName())
	}

	action := request.ActionName()
	station.handle(action, func(ctx context.Context, raw json.RawMessage) (json.RawMessage, error) {
		if err := request.ValidateJSON(raw); err != nil {
			return nil, fmt.Errorf("stationtest: invalid %s request from CSMS: %w", action, err)
		}
		var decoded Request
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return nil, fmt.Errorf("stationtest: decode %s request: %w", action, err)
		}
		response, err := handler(ctx, decoded)
		if err != nil {
			return nil, err
		}
		if err := response.Validate(); err != nil {
			return nil, fmt.Errorf("stationtest: handler returned an invalid %s confirmation: %w", action, err)
		}
		return json.Marshal(response)
	})
	return nil
}
