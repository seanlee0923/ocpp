// Package ocpp201 implements the server-side OCPP 2.0.1 Core registration flow.
package ocpp201

import (
	"context"
	"encoding/json"
	"errors"
	"sync"
	"sync/atomic"

	"github.com/seanlee0923/ocpp/csms"
	"github.com/seanlee0923/ocpp/protocol"
	"github.com/seanlee0923/ocpp/v201"
)

var ErrNotBooted = errors.New("OCPP 2.0.1 session has not completed an accepted BootNotification")

type RegistrationState uint32

const (
	RegistrationUnknown RegistrationState = iota
	RegistrationAccepted
	RegistrationPending
	RegistrationRejected
)

type BootNotificationHandler func(context.Context, *csms.Session, v201.BootNotificationRequest) (v201.BootNotificationConfirmation, error)
type HeartbeatHandler func(context.Context, *csms.Session, v201.HeartbeatRequest) (v201.HeartbeatConfirmation, error)
type StatusNotificationHandler func(context.Context, *csms.Session, v201.StatusNotificationRequest) (v201.StatusNotificationConfirmation, error)
type AuthorizeHandler func(context.Context, *csms.Session, v201.AuthorizeRequest) (v201.AuthorizeConfirmation, error)
type TransactionEventHandler func(context.Context, *csms.Session, v201.TransactionEventRequest) (v201.TransactionEventConfirmation, error)
type MeterValuesHandler func(context.Context, *csms.Session, v201.MeterValuesRequest) (v201.MeterValuesConfirmation, error)
type NotifyReportHandler func(context.Context, *csms.Session, v201.NotifyReportRequest) (v201.NotifyReportConfirmation, error)

type sessionState struct{ registration atomic.Uint32 }

type Profile struct {
	router *csms.Router
	states sync.Map // map[*csms.Session]*sessionState
}

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
	return csms.Handle(profile.router, func(ctx context.Context, session *csms.Session, request v201.BootNotificationRequest) (v201.BootNotificationConfirmation, error) {
		confirmation, err := handler(ctx, session, request)
		if err != nil {
			return confirmation, err
		}
		if err := confirmation.Validate(); err != nil {
			return confirmation, err
		}
		state := profile.stateFor(session)
		switch confirmation.Status {
		case v201.BootNotificationConfirmationRegistrationStatusEnumAccepted:
			state.registration.Store(uint32(RegistrationAccepted))
		case v201.BootNotificationConfirmationRegistrationStatusEnumPending:
			state.registration.Store(uint32(RegistrationPending))
		case v201.BootNotificationConfirmationRegistrationStatusEnumRejected:
			state.registration.Store(uint32(RegistrationRejected))
		}
		return confirmation, nil
	})
}

func (profile *Profile) HandleHeartbeat(handler HeartbeatHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v201.HeartbeatRequest, v201.HeartbeatConfirmation](handler))
}
func (profile *Profile) HandleStatusNotification(handler StatusNotificationHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v201.StatusNotificationRequest, v201.StatusNotificationConfirmation](handler))
}
func (profile *Profile) HandleAuthorize(handler AuthorizeHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v201.AuthorizeRequest, v201.AuthorizeConfirmation](handler))
}
func (profile *Profile) HandleTransactionEvent(handler TransactionEventHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v201.TransactionEventRequest, v201.TransactionEventConfirmation](handler))
}
func (profile *Profile) HandleMeterValues(handler MeterValuesHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v201.MeterValuesRequest, v201.MeterValuesConfirmation](handler))
}
func (profile *Profile) HandleNotifyReport(handler NotifyReportHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v201.NotifyReportRequest, v201.NotifyReportConfirmation](handler))
}

func (profile *Profile) CallGetVariables(ctx context.Context, session *csms.Session, request v201.GetVariablesRequest) (v201.GetVariablesConfirmation, error) {
	if err := profile.requireBooted(session); err != nil {
		return v201.GetVariablesConfirmation{}, err
	}
	return csms.Call[v201.GetVariablesRequest, v201.GetVariablesConfirmation](ctx, session, request)
}
func (profile *Profile) CallSetVariables(ctx context.Context, session *csms.Session, request v201.SetVariablesRequest) (v201.SetVariablesConfirmation, error) {
	if err := profile.requireBooted(session); err != nil {
		return v201.SetVariablesConfirmation{}, err
	}
	return csms.Call[v201.SetVariablesRequest, v201.SetVariablesConfirmation](ctx, session, request)
}
func (profile *Profile) CallGetBaseReport(ctx context.Context, session *csms.Session, request v201.GetBaseReportRequest) (v201.GetBaseReportConfirmation, error) {
	if err := profile.requireBooted(session); err != nil {
		return v201.GetBaseReportConfirmation{}, err
	}
	return csms.Call[v201.GetBaseReportRequest, v201.GetBaseReportConfirmation](ctx, session, request)
}
func (profile *Profile) CallRequestStartTransaction(ctx context.Context, session *csms.Session, request v201.RequestStartTransactionRequest) (v201.RequestStartTransactionConfirmation, error) {
	if err := profile.requireBooted(session); err != nil {
		return v201.RequestStartTransactionConfirmation{}, err
	}
	return csms.Call[v201.RequestStartTransactionRequest, v201.RequestStartTransactionConfirmation](ctx, session, request)
}
func (profile *Profile) CallRequestStopTransaction(ctx context.Context, session *csms.Session, request v201.RequestStopTransactionRequest) (v201.RequestStopTransactionConfirmation, error) {
	if err := profile.requireBooted(session); err != nil {
		return v201.RequestStopTransactionConfirmation{}, err
	}
	return csms.Call[v201.RequestStopTransactionRequest, v201.RequestStopTransactionConfirmation](ctx, session, request)
}
func (profile *Profile) CallReset(ctx context.Context, session *csms.Session, request v201.ResetRequest) (v201.ResetConfirmation, error) {
	if err := profile.requireBooted(session); err != nil {
		return v201.ResetConfirmation{}, err
	}
	return csms.Call[v201.ResetRequest, v201.ResetConfirmation](ctx, session, request)
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
	if version != protocol.OCPP201 || action == "BootNotification" {
		return next
	}
	return func(ctx context.Context, session *csms.Session, payload json.RawMessage) (any, error) {
		if !profile.IsBooted(session) {
			return nil, &csms.CallError{Code: csms.ProtocolError, Description: "charging station has not completed BootNotification", Details: map[string]any{}}
		}
		return next(ctx, session, payload)
	}
}
