package csms

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

// TypedHandler receives a schema-validated request and returns a confirmation
// of the same OCPP version and action.
type TypedHandler[Request protocol.Payload, Confirmation protocol.Payload] func(
	context.Context,
	*Session,
	Request,
) (Confirmation, error)

type TypedSendHandler[Request protocol.Payload] func(context.Context, *Session, Request) error

// HandleSend registers an OCPP 2.1 SEND handler. SEND is unconfirmed, so an
// error returned by the handler is not written back to the Charging Station.
func HandleSend[Request protocol.Payload](router *Router, handler TypedSendHandler[Request]) error {
	if router == nil {
		return fmt.Errorf("router is nil")
	}
	if handler == nil {
		return fmt.Errorf("handler is nil")
	}
	var request Request
	if isNilType(request) || request.Direction() != protocol.RequestPayload {
		return fmt.Errorf("SEND handler requires a non-pointer generated request payload")
	}
	version, action := request.Version(), request.ActionName()
	return router.Handle(version, action, func(ctx context.Context, session *Session, raw json.RawMessage) (any, error) {
		if err := request.ValidateJSON(raw); err != nil {
			return nil, err
		}
		var decoded Request
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return nil, err
		}
		return nil, handler(ctx, session, decoded)
	})
}

// Handle registers a type-safe OCPP handler. Request and Confirmation must be
// generated non-pointer payload types for the same action and version.
func Handle[Request protocol.Payload, Confirmation protocol.Payload](
	router *Router,
	handler TypedHandler[Request, Confirmation],
) error {
	if router == nil {
		return fmt.Errorf("router is nil")
	}
	if handler == nil {
		return fmt.Errorf("handler is nil")
	}
	var request Request
	var confirmation Confirmation
	if isNilType(request) || isNilType(confirmation) {
		return fmt.Errorf("request and confirmation type parameters must be non-pointer generated payloads")
	}
	if request.Direction() != protocol.RequestPayload || confirmation.Direction() != protocol.ConfirmationPayload {
		return fmt.Errorf("typed handler requires a request followed by a confirmation")
	}
	if request.Version() != confirmation.Version() {
		return fmt.Errorf("request version %s does not match confirmation version %s", request.Version(), confirmation.Version())
	}
	if request.ActionName() != confirmation.ActionName() {
		return fmt.Errorf("request action %s does not match confirmation action %s", request.ActionName(), confirmation.ActionName())
	}

	version := request.Version()
	action := request.ActionName()
	return router.Handle(version, action, func(ctx context.Context, session *Session, raw json.RawMessage) (any, error) {
		if err := request.ValidateJSON(raw); err != nil {
			return nil, &CallError{
				Code:        validationErrorCode(version, err),
				Description: "request payload does not conform to the OCPP schema",
				Details:     map[string]any{"error": err.Error()},
			}
		}

		var decoded Request
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return nil, &CallError{
				Code:        formatViolationCode(version),
				Description: "request payload cannot be decoded",
				Details:     map[string]any{"error": err.Error()},
			}
		}
		response, err := handler(ctx, session, decoded)
		if err != nil {
			return nil, err
		}
		if err := response.Validate(); err != nil {
			return nil, &CallError{
				Code:        InternalError,
				Description: "handler returned an invalid OCPP confirmation",
				Details:     map[string]any{"error": err.Error()},
			}
		}
		return response, nil
	})
}

func validationErrorCode(version protocol.Version, err error) ErrorCode {
	kind, ok := validation.KindOf(err)
	if !ok {
		return formatViolationCode(version)
	}
	switch kind {
	case validation.PropertyConstraintError:
		return PropertyConstraintViolation
	case validation.OccurrenceConstraintError:
		if version == protocol.OCPP16 {
			return OccurenceConstraintViolation
		}
		return OccurrenceConstraintViolation
	case validation.TypeConstraintError:
		return TypeConstraintViolation
	default:
		return formatViolationCode(version)
	}
}

func formatViolationCode(version protocol.Version) ErrorCode {
	if version == protocol.OCPP16 {
		return FormationViolation
	}
	return FormatViolation
}

func isNilType[T any](value T) bool {
	typeOf := reflect.TypeOf(value)
	return typeOf == nil || typeOf.Kind() == reflect.Pointer
}
