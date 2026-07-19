package station

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/seanlee0923/ocpp/protocol"
)

// Handle registers a typed handler for a CSMS-initiated CALL. It is the
// Station-side mirror of csms.Handle: the inbound payload is schema-
// validated before decode, and the handler's returned confirmation is
// schema-validated before it is sent back as a CALLRESULT. Handlers are
// registered on the Station and survive reconnects.
func Handle[Request protocol.Payload, Confirmation protocol.Payload](
	s *Station,
	handler func(context.Context, Request) (Confirmation, error),
) error {
	if s == nil {
		return fmt.Errorf("station: Station is nil")
	}
	if handler == nil {
		return fmt.Errorf("%w: handler is nil", ErrInvalidHandlerRegistration)
	}
	var request Request
	var confirmation Confirmation
	if isNilType(request) || isNilType(confirmation) {
		return fmt.Errorf("%w: request and confirmation type parameters must be non-pointer generated payloads", ErrInvalidHandlerRegistration)
	}
	if request.Direction() != protocol.RequestPayload || confirmation.Direction() != protocol.ConfirmationPayload {
		return fmt.Errorf("station: Handle requires a request followed by a confirmation")
	}
	if request.Version() != s.config.Version {
		return fmt.Errorf("station: request version %s does not match station version %s", request.Version(), s.config.Version)
	}
	if request.Version() != confirmation.Version() {
		return fmt.Errorf("station: request version %s does not match confirmation version %s", request.Version(), confirmation.Version())
	}
	if request.ActionName() != confirmation.ActionName() {
		return fmt.Errorf("station: request action %s does not match confirmation action %s", request.ActionName(), confirmation.ActionName())
	}

	action := request.ActionName()
	return s.handle(action, func(ctx context.Context, raw json.RawMessage) (json.RawMessage, error) {
		if err := request.ValidateJSON(raw); err != nil {
			return nil, fmt.Errorf("station: invalid %s request from CSMS: %w", action, err)
		}
		var decoded Request
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return nil, fmt.Errorf("station: decode %s request: %w", action, err)
		}
		response, err := handler(ctx, decoded)
		if err != nil {
			return nil, err
		}
		if err := response.Validate(); err != nil {
			return nil, fmt.Errorf("station: handler returned an invalid %s confirmation: %w", action, err)
		}
		return json.Marshal(response)
	})
}
