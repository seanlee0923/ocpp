// Package ocpp16 implements the server-side OCPP 1.6 Core registration flow.
package ocpp16

import (
	"context"
	"encoding/json"
	"errors"
	"sync"
	"sync/atomic"

	"github.com/seanlee0923/ocpp/csms"
	"github.com/seanlee0923/ocpp/protocol"
	"github.com/seanlee0923/ocpp/v16"
)

var ErrNotBooted = errors.New("OCPP 1.6 session has not completed an accepted BootNotification")

type RegistrationState uint32

const (
	RegistrationUnknown RegistrationState = iota
	RegistrationAccepted
	RegistrationPending
	RegistrationRejected
)

type BootNotificationHandler func(context.Context, *csms.Session, v16.BootNotificationRequest) (v16.BootNotificationConfirmation, error)
type HeartbeatHandler func(context.Context, *csms.Session, v16.HeartbeatRequest) (v16.HeartbeatConfirmation, error)
type StatusNotificationHandler func(context.Context, *csms.Session, v16.StatusNotificationRequest) (v16.StatusNotificationConfirmation, error)
type AuthorizeHandler func(context.Context, *csms.Session, v16.AuthorizeRequest) (v16.AuthorizeConfirmation, error)
type StartTransactionHandler func(context.Context, *csms.Session, v16.StartTransactionRequest) (v16.StartTransactionConfirmation, error)
type MeterValuesHandler func(context.Context, *csms.Session, v16.MeterValuesRequest) (v16.MeterValuesConfirmation, error)
type StopTransactionHandler func(context.Context, *csms.Session, v16.StopTransactionRequest) (v16.StopTransactionConfirmation, error)
type DataTransferHandler func(context.Context, *csms.Session, v16.DataTransferRequest) (v16.DataTransferConfirmation, error)

type sessionState struct {
	registration atomic.Uint32
}

type Profile struct {
	router *csms.Router
	states sync.Map // map[*csms.Session]*sessionState
}

// NewProfile adds OCPP 1.6 protocol policy to router. Application handlers are
// registered independently with the Handle methods on Profile.
func NewProfile(router *csms.Router) (*Profile, error) {
	if router == nil {
		return nil, errors.New("router is nil")
	}
	profile := &Profile{router: router}
	router.Use(profile.registrationMiddleware)
	return profile, nil
}

func (profile *Profile) HandleBootNotification(handler BootNotificationHandler) error {
	if handler == nil {
		return errors.New("BootNotification handler is nil")
	}
	return csms.Handle(profile.router, func(
		ctx context.Context,
		session *csms.Session,
		request v16.BootNotificationRequest,
	) (v16.BootNotificationConfirmation, error) {
		confirmation, err := handler(ctx, session, request)
		if err != nil {
			return confirmation, err
		}
		if err := confirmation.Validate(); err != nil {
			return confirmation, err
		}
		state := profile.stateFor(session)
		switch confirmation.Status {
		case v16.BootNotificationConfirmationStatusAccepted:
			state.registration.Store(uint32(RegistrationAccepted))
		case v16.BootNotificationConfirmationStatusPending:
			state.registration.Store(uint32(RegistrationPending))
		case v16.BootNotificationConfirmationStatusRejected:
			state.registration.Store(uint32(RegistrationRejected))
		}
		return confirmation, nil
	})
}

func (profile *Profile) HandleHeartbeat(handler HeartbeatHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v16.HeartbeatRequest, v16.HeartbeatConfirmation](handler))
}

func (profile *Profile) HandleStatusNotification(handler StatusNotificationHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v16.StatusNotificationRequest, v16.StatusNotificationConfirmation](handler))
}

// HandleAuthorize registers application-owned idTag authorization logic. The
// profile only enforces that BootNotification was accepted and validates the
// request and confirmation against generated OCPP 1.6 constraints.
func (profile *Profile) HandleAuthorize(handler AuthorizeHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v16.AuthorizeRequest, v16.AuthorizeConfirmation](handler))
}

// HandleStartTransaction registers application-owned transaction creation.
// The application decides whether the idTag is accepted and allocates the
// OCPP 1.6 integer transactionId returned to the charging station.
func (profile *Profile) HandleStartTransaction(handler StartTransactionHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v16.StartTransactionRequest, v16.StartTransactionConfirmation](handler))
}

func (profile *Profile) HandleMeterValues(handler MeterValuesHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v16.MeterValuesRequest, v16.MeterValuesConfirmation](handler))
}

func (profile *Profile) HandleStopTransaction(handler StopTransactionHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v16.StopTransactionRequest, v16.StopTransactionConfirmation](handler))
}

// HandleDataTransfer registers the Charging Station to CSMS direction. Use
// CallDataTransfer for the opposite direction.
func (profile *Profile) HandleDataTransfer(handler DataTransferHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v16.DataTransferRequest, v16.DataTransferConfirmation](handler))
}

func (profile *Profile) CallDataTransfer(ctx context.Context, session *csms.Session, request v16.DataTransferRequest) (v16.DataTransferConfirmation, error) {
	return callBooted[v16.DataTransferRequest, v16.DataTransferConfirmation](profile, ctx, session, request)
}

func (profile *Profile) CallReset(ctx context.Context, session *csms.Session, request v16.ResetRequest) (v16.ResetConfirmation, error) {
	return callBooted[v16.ResetRequest, v16.ResetConfirmation](profile, ctx, session, request)
}

func (profile *Profile) CallRemoteStartTransaction(ctx context.Context, session *csms.Session, request v16.RemoteStartTransactionRequest) (v16.RemoteStartTransactionConfirmation, error) {
	return callBooted[v16.RemoteStartTransactionRequest, v16.RemoteStartTransactionConfirmation](profile, ctx, session, request)
}

func (profile *Profile) CallRemoteStopTransaction(ctx context.Context, session *csms.Session, request v16.RemoteStopTransactionRequest) (v16.RemoteStopTransactionConfirmation, error) {
	return callBooted[v16.RemoteStopTransactionRequest, v16.RemoteStopTransactionConfirmation](profile, ctx, session, request)
}

func (profile *Profile) CallChangeAvailability(ctx context.Context, session *csms.Session, request v16.ChangeAvailabilityRequest) (v16.ChangeAvailabilityConfirmation, error) {
	return callBooted[v16.ChangeAvailabilityRequest, v16.ChangeAvailabilityConfirmation](profile, ctx, session, request)
}

func (profile *Profile) CallUnlockConnector(ctx context.Context, session *csms.Session, request v16.UnlockConnectorRequest) (v16.UnlockConnectorConfirmation, error) {
	return callBooted[v16.UnlockConnectorRequest, v16.UnlockConnectorConfirmation](profile, ctx, session, request)
}

func (profile *Profile) CallTriggerMessage(ctx context.Context, session *csms.Session, request v16.TriggerMessageRequest) (v16.TriggerMessageConfirmation, error) {
	return callBooted[v16.TriggerMessageRequest, v16.TriggerMessageConfirmation](profile, ctx, session, request)
}

// callBooted requires an accepted BootNotification before delegating to
// csms.Call, matching the requireBooted guard every Call<Action> method uses.
func callBooted[Req, Conf protocol.Payload](profile *Profile, ctx context.Context, session *csms.Session, request Req) (Conf, error) {
	var zero Conf
	if err := profile.requireBooted(session); err != nil {
		return zero, err
	}
	return csms.Call[Req, Conf](ctx, session, request)
}

func (profile *Profile) requireBooted(session *csms.Session) error {
	if !profile.IsBooted(session) {
		return ErrNotBooted
	}
	return nil
}

func (profile *Profile) State(session *csms.Session) RegistrationState {
	if session == nil {
		return RegistrationUnknown
	}
	value, ok := profile.states.Load(session)
	if !ok {
		return RegistrationUnknown
	}
	return RegistrationState(value.(*sessionState).registration.Load())
}

func (profile *Profile) IsBooted(session *csms.Session) bool {
	return profile.State(session) == RegistrationAccepted
}

func (profile *Profile) stateFor(session *csms.Session) *sessionState {
	state := &sessionState{}
	actual, loaded := profile.states.LoadOrStore(session, state)
	if loaded {
		return actual.(*sessionState)
	}
	go func() {
		<-session.Done()
		profile.states.Delete(session)
	}()
	return state
}

func (profile *Profile) registrationMiddleware(version protocol.Version, action string, next csms.Handler) csms.Handler {
	if version != protocol.OCPP16 || action == "BootNotification" {
		return next
	}
	return func(ctx context.Context, session *csms.Session, payload json.RawMessage) (any, error) {
		if !profile.IsBooted(session) {
			return nil, &csms.CallError{
				Code:        csms.ProtocolError,
				Description: "charging station has not completed BootNotification",
				Details:     map[string]any{},
			}
		}
		return next(ctx, session, payload)
	}
}
