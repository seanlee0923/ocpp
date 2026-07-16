// Package ocpp21 implements the server-side OCPP 2.1 Core registration flow.
package ocpp21

import (
	"context"
	"encoding/json"
	"errors"
	"sync"
	"sync/atomic"

	"ocpp-go/csms"
	"ocpp-go/protocol"
	"ocpp-go/v21"
)

var ErrNotBooted = errors.New("OCPP 2.1 session has not completed an accepted BootNotification")

type RegistrationState uint32

const (
	RegistrationUnknown RegistrationState = iota
	RegistrationAccepted
	RegistrationPending
	RegistrationRejected
)

type BootNotificationHandler func(context.Context, *csms.Session, v21.BootNotificationRequest) (v21.BootNotificationConfirmation, error)
type HeartbeatHandler func(context.Context, *csms.Session, v21.HeartbeatRequest) (v21.HeartbeatConfirmation, error)
type StatusNotificationHandler func(context.Context, *csms.Session, v21.StatusNotificationRequest) (v21.StatusNotificationConfirmation, error)
type AuthorizeHandler func(context.Context, *csms.Session, v21.AuthorizeRequest) (v21.AuthorizeConfirmation, error)
type TransactionEventHandler func(context.Context, *csms.Session, v21.TransactionEventRequest) (v21.TransactionEventConfirmation, error)
type MeterValuesHandler func(context.Context, *csms.Session, v21.MeterValuesRequest) (v21.MeterValuesConfirmation, error)
type NotifyReportHandler func(context.Context, *csms.Session, v21.NotifyReportRequest) (v21.NotifyReportConfirmation, error)
type NotifyEVChargingNeedsHandler func(context.Context, *csms.Session, v21.NotifyEVChargingNeedsRequest) (v21.NotifyEVChargingNeedsConfirmation, error)
type NotifyEVChargingScheduleHandler func(context.Context, *csms.Session, v21.NotifyEVChargingScheduleRequest) (v21.NotifyEVChargingScheduleConfirmation, error)
type NotifyAllowedEnergyTransferHandler func(context.Context, *csms.Session, v21.NotifyAllowedEnergyTransferRequest) (v21.NotifyAllowedEnergyTransferConfirmation, error)
type NotifyChargingLimitHandler func(context.Context, *csms.Session, v21.NotifyChargingLimitRequest) (v21.NotifyChargingLimitConfirmation, error)
type ClearedChargingLimitHandler func(context.Context, *csms.Session, v21.ClearedChargingLimitRequest) (v21.ClearedChargingLimitConfirmation, error)
type ReportChargingProfilesHandler func(context.Context, *csms.Session, v21.ReportChargingProfilesRequest) (v21.ReportChargingProfilesConfirmation, error)
type NotifyPriorityChargingHandler func(context.Context, *csms.Session, v21.NotifyPriorityChargingRequest) (v21.NotifyPriorityChargingConfirmation, error)
type PullDynamicScheduleUpdateHandler func(context.Context, *csms.Session, v21.PullDynamicScheduleUpdateRequest) (v21.PullDynamicScheduleUpdateConfirmation, error)
type ReportDERControlHandler func(context.Context, *csms.Session, v21.ReportDERControlRequest) (v21.ReportDERControlConfirmation, error)
type NotifyDERAlarmHandler func(context.Context, *csms.Session, v21.NotifyDERAlarmRequest) (v21.NotifyDERAlarmConfirmation, error)
type NotifyDERStartStopHandler func(context.Context, *csms.Session, v21.NotifyDERStartStopRequest) (v21.NotifyDERStartStopConfirmation, error)
type GetTariffsHandler func(context.Context, *csms.Session, v21.GetTariffsRequest) (v21.GetTariffsConfirmation, error)
type NotifySettlementHandler func(context.Context, *csms.Session, v21.NotifySettlementRequest) (v21.NotifySettlementConfirmation, error)
type NotifyWebPaymentStartedHandler func(context.Context, *csms.Session, v21.NotifyWebPaymentStartedRequest) (v21.NotifyWebPaymentStartedConfirmation, error)
type VatNumberValidationHandler func(context.Context, *csms.Session, v21.VatNumberValidationRequest) (v21.VatNumberValidationConfirmation, error)
type BatterySwapHandler func(context.Context, *csms.Session, v21.BatterySwapRequest) (v21.BatterySwapConfirmation, error)

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
	return csms.Handle(profile.router, func(ctx context.Context, session *csms.Session, request v21.BootNotificationRequest) (v21.BootNotificationConfirmation, error) {
		confirmation, err := handler(ctx, session, request)
		if err != nil {
			return confirmation, err
		}
		if err := confirmation.Validate(); err != nil {
			return confirmation, err
		}
		state := profile.stateFor(session)
		switch confirmation.Status {
		case v21.BootNotificationConfirmationRegistrationStatusEnumAccepted:
			state.registration.Store(uint32(RegistrationAccepted))
		case v21.BootNotificationConfirmationRegistrationStatusEnumPending:
			state.registration.Store(uint32(RegistrationPending))
		case v21.BootNotificationConfirmationRegistrationStatusEnumRejected:
			state.registration.Store(uint32(RegistrationRejected))
		}
		return confirmation, nil
	})
}

func (profile *Profile) HandleHeartbeat(handler HeartbeatHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v21.HeartbeatRequest, v21.HeartbeatConfirmation](handler))
}
func (profile *Profile) HandleStatusNotification(handler StatusNotificationHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v21.StatusNotificationRequest, v21.StatusNotificationConfirmation](handler))
}
func (profile *Profile) HandleAuthorize(handler AuthorizeHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v21.AuthorizeRequest, v21.AuthorizeConfirmation](handler))
}
func (profile *Profile) HandleTransactionEvent(handler TransactionEventHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v21.TransactionEventRequest, v21.TransactionEventConfirmation](handler))
}
func (profile *Profile) HandleMeterValues(handler MeterValuesHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v21.MeterValuesRequest, v21.MeterValuesConfirmation](handler))
}
func (profile *Profile) HandleNotifyReport(handler NotifyReportHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v21.NotifyReportRequest, v21.NotifyReportConfirmation](handler))
}
func (profile *Profile) HandleNotifyEVChargingNeeds(handler NotifyEVChargingNeedsHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v21.NotifyEVChargingNeedsRequest, v21.NotifyEVChargingNeedsConfirmation](handler))
}
func (profile *Profile) HandleNotifyEVChargingSchedule(handler NotifyEVChargingScheduleHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v21.NotifyEVChargingScheduleRequest, v21.NotifyEVChargingScheduleConfirmation](handler))
}
func (profile *Profile) HandleNotifyAllowedEnergyTransfer(handler NotifyAllowedEnergyTransferHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v21.NotifyAllowedEnergyTransferRequest, v21.NotifyAllowedEnergyTransferConfirmation](handler))
}
func (profile *Profile) HandleNotifyChargingLimit(handler NotifyChargingLimitHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v21.NotifyChargingLimitRequest, v21.NotifyChargingLimitConfirmation](handler))
}
func (profile *Profile) HandleClearedChargingLimit(handler ClearedChargingLimitHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v21.ClearedChargingLimitRequest, v21.ClearedChargingLimitConfirmation](handler))
}
func (profile *Profile) HandleReportChargingProfiles(handler ReportChargingProfilesHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v21.ReportChargingProfilesRequest, v21.ReportChargingProfilesConfirmation](handler))
}
func (profile *Profile) HandleNotifyPriorityCharging(handler NotifyPriorityChargingHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v21.NotifyPriorityChargingRequest, v21.NotifyPriorityChargingConfirmation](handler))
}
func (profile *Profile) HandlePullDynamicScheduleUpdate(handler PullDynamicScheduleUpdateHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v21.PullDynamicScheduleUpdateRequest, v21.PullDynamicScheduleUpdateConfirmation](handler))
}
func (profile *Profile) HandleReportDERControl(handler ReportDERControlHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v21.ReportDERControlRequest, v21.ReportDERControlConfirmation](handler))
}
func (profile *Profile) HandleNotifyDERAlarm(handler NotifyDERAlarmHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v21.NotifyDERAlarmRequest, v21.NotifyDERAlarmConfirmation](handler))
}
func (profile *Profile) HandleNotifyDERStartStop(handler NotifyDERStartStopHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v21.NotifyDERStartStopRequest, v21.NotifyDERStartStopConfirmation](handler))
}
func (profile *Profile) HandleGetTariffs(handler GetTariffsHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v21.GetTariffsRequest, v21.GetTariffsConfirmation](handler))
}
func (profile *Profile) HandleNotifySettlement(handler NotifySettlementHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v21.NotifySettlementRequest, v21.NotifySettlementConfirmation](handler))
}
func (profile *Profile) HandleNotifyWebPaymentStarted(handler NotifyWebPaymentStartedHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v21.NotifyWebPaymentStartedRequest, v21.NotifyWebPaymentStartedConfirmation](handler))
}
func (profile *Profile) HandleVatNumberValidation(handler VatNumberValidationHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v21.VatNumberValidationRequest, v21.VatNumberValidationConfirmation](handler))
}
func (profile *Profile) HandleBatterySwap(handler BatterySwapHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v21.BatterySwapRequest, v21.BatterySwapConfirmation](handler))
}

func (profile *Profile) CallGetVariables(ctx context.Context, session *csms.Session, request v21.GetVariablesRequest) (v21.GetVariablesConfirmation, error) {
	if err := profile.requireBooted(session); err != nil {
		return v21.GetVariablesConfirmation{}, err
	}
	return csms.Call[v21.GetVariablesRequest, v21.GetVariablesConfirmation](ctx, session, request)
}
func (profile *Profile) CallSetVariables(ctx context.Context, session *csms.Session, request v21.SetVariablesRequest) (v21.SetVariablesConfirmation, error) {
	if err := profile.requireBooted(session); err != nil {
		return v21.SetVariablesConfirmation{}, err
	}
	return csms.Call[v21.SetVariablesRequest, v21.SetVariablesConfirmation](ctx, session, request)
}
func (profile *Profile) CallGetBaseReport(ctx context.Context, session *csms.Session, request v21.GetBaseReportRequest) (v21.GetBaseReportConfirmation, error) {
	if err := profile.requireBooted(session); err != nil {
		return v21.GetBaseReportConfirmation{}, err
	}
	return csms.Call[v21.GetBaseReportRequest, v21.GetBaseReportConfirmation](ctx, session, request)
}
func (profile *Profile) CallRequestStartTransaction(ctx context.Context, session *csms.Session, request v21.RequestStartTransactionRequest) (v21.RequestStartTransactionConfirmation, error) {
	if err := profile.requireBooted(session); err != nil {
		return v21.RequestStartTransactionConfirmation{}, err
	}
	return csms.Call[v21.RequestStartTransactionRequest, v21.RequestStartTransactionConfirmation](ctx, session, request)
}
func (profile *Profile) CallRequestStopTransaction(ctx context.Context, session *csms.Session, request v21.RequestStopTransactionRequest) (v21.RequestStopTransactionConfirmation, error) {
	if err := profile.requireBooted(session); err != nil {
		return v21.RequestStopTransactionConfirmation{}, err
	}
	return csms.Call[v21.RequestStopTransactionRequest, v21.RequestStopTransactionConfirmation](ctx, session, request)
}
func (profile *Profile) CallReset(ctx context.Context, session *csms.Session, request v21.ResetRequest) (v21.ResetConfirmation, error) {
	if err := profile.requireBooted(session); err != nil {
		return v21.ResetConfirmation{}, err
	}
	return csms.Call[v21.ResetRequest, v21.ResetConfirmation](ctx, session, request)
}
func (profile *Profile) CallGetChargingProfiles(ctx context.Context, session *csms.Session, request v21.GetChargingProfilesRequest) (v21.GetChargingProfilesConfirmation, error) {
	if err := profile.requireBooted(session); err != nil {
		return v21.GetChargingProfilesConfirmation{}, err
	}
	return csms.Call[v21.GetChargingProfilesRequest, v21.GetChargingProfilesConfirmation](ctx, session, request)
}
func (profile *Profile) CallSetChargingProfile(ctx context.Context, session *csms.Session, request v21.SetChargingProfileRequest) (v21.SetChargingProfileConfirmation, error) {
	if err := profile.requireBooted(session); err != nil {
		return v21.SetChargingProfileConfirmation{}, err
	}
	return csms.Call[v21.SetChargingProfileRequest, v21.SetChargingProfileConfirmation](ctx, session, request)
}
func (profile *Profile) CallClearChargingProfile(ctx context.Context, session *csms.Session, request v21.ClearChargingProfileRequest) (v21.ClearChargingProfileConfirmation, error) {
	if err := profile.requireBooted(session); err != nil {
		return v21.ClearChargingProfileConfirmation{}, err
	}
	return csms.Call[v21.ClearChargingProfileRequest, v21.ClearChargingProfileConfirmation](ctx, session, request)
}
func (profile *Profile) CallUsePriorityCharging(ctx context.Context, session *csms.Session, request v21.UsePriorityChargingRequest) (v21.UsePriorityChargingConfirmation, error) {
	if err := profile.requireBooted(session); err != nil {
		return v21.UsePriorityChargingConfirmation{}, err
	}
	return csms.Call[v21.UsePriorityChargingRequest, v21.UsePriorityChargingConfirmation](ctx, session, request)
}
func (profile *Profile) CallUpdateDynamicSchedule(ctx context.Context, session *csms.Session, request v21.UpdateDynamicScheduleRequest) (v21.UpdateDynamicScheduleConfirmation, error) {
	if err := profile.requireBooted(session); err != nil {
		return v21.UpdateDynamicScheduleConfirmation{}, err
	}
	return csms.Call[v21.UpdateDynamicScheduleRequest, v21.UpdateDynamicScheduleConfirmation](ctx, session, request)
}
func (profile *Profile) CallAFRRSignal(ctx context.Context, session *csms.Session, request v21.AFRRSignalRequest) (v21.AFRRSignalConfirmation, error) {
	if err := profile.requireBooted(session); err != nil {
		return v21.AFRRSignalConfirmation{}, err
	}
	return csms.Call[v21.AFRRSignalRequest, v21.AFRRSignalConfirmation](ctx, session, request)
}
func (profile *Profile) CallSetDERControl(ctx context.Context, session *csms.Session, request v21.SetDERControlRequest) (v21.SetDERControlConfirmation, error) {
	if err := profile.requireBooted(session); err != nil {
		return v21.SetDERControlConfirmation{}, err
	}
	return csms.Call[v21.SetDERControlRequest, v21.SetDERControlConfirmation](ctx, session, request)
}
func (profile *Profile) CallGetDERControl(ctx context.Context, session *csms.Session, request v21.GetDERControlRequest) (v21.GetDERControlConfirmation, error) {
	if err := profile.requireBooted(session); err != nil {
		return v21.GetDERControlConfirmation{}, err
	}
	return csms.Call[v21.GetDERControlRequest, v21.GetDERControlConfirmation](ctx, session, request)
}
func (profile *Profile) CallClearDERControl(ctx context.Context, session *csms.Session, request v21.ClearDERControlRequest) (v21.ClearDERControlConfirmation, error) {
	if err := profile.requireBooted(session); err != nil {
		return v21.ClearDERControlConfirmation{}, err
	}
	return csms.Call[v21.ClearDERControlRequest, v21.ClearDERControlConfirmation](ctx, session, request)
}
func (profile *Profile) CallSetDefaultTariff(ctx context.Context, session *csms.Session, request v21.SetDefaultTariffRequest) (v21.SetDefaultTariffConfirmation, error) {
	if err := profile.requireBooted(session); err != nil {
		return v21.SetDefaultTariffConfirmation{}, err
	}
	return csms.Call[v21.SetDefaultTariffRequest, v21.SetDefaultTariffConfirmation](ctx, session, request)
}
func (profile *Profile) CallClearTariffs(ctx context.Context, session *csms.Session, request v21.ClearTariffsRequest) (v21.ClearTariffsConfirmation, error) {
	if err := profile.requireBooted(session); err != nil {
		return v21.ClearTariffsConfirmation{}, err
	}
	return csms.Call[v21.ClearTariffsRequest, v21.ClearTariffsConfirmation](ctx, session, request)
}
func (profile *Profile) CallChangeTransactionTariff(ctx context.Context, session *csms.Session, request v21.ChangeTransactionTariffRequest) (v21.ChangeTransactionTariffConfirmation, error) {
	if err := profile.requireBooted(session); err != nil {
		return v21.ChangeTransactionTariffConfirmation{}, err
	}
	return csms.Call[v21.ChangeTransactionTariffRequest, v21.ChangeTransactionTariffConfirmation](ctx, session, request)
}
func (profile *Profile) CallCostUpdated(ctx context.Context, session *csms.Session, request v21.CostUpdatedRequest) (v21.CostUpdatedConfirmation, error) {
	if err := profile.requireBooted(session); err != nil {
		return v21.CostUpdatedConfirmation{}, err
	}
	return csms.Call[v21.CostUpdatedRequest, v21.CostUpdatedConfirmation](ctx, session, request)
}
func (profile *Profile) CallRequestBatterySwap(ctx context.Context, session *csms.Session, request v21.RequestBatterySwapRequest) (v21.RequestBatterySwapConfirmation, error) {
	if err := profile.requireBooted(session); err != nil {
		return v21.RequestBatterySwapConfirmation{}, err
	}
	return csms.Call[v21.RequestBatterySwapRequest, v21.RequestBatterySwapConfirmation](ctx, session, request)
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
	if version != protocol.OCPP21 || action == "BootNotification" {
		return next
	}
	return func(ctx context.Context, session *csms.Session, payload json.RawMessage) (any, error) {
		if !profile.IsBooted(session) {
			return nil, &csms.CallError{Code: csms.ProtocolError, Description: "charging station has not completed BootNotification", Details: map[string]any{}}
		}
		return next(ctx, session, payload)
	}
}
