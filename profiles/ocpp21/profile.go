// Package ocpp21 implements the server-side OCPP 2.1 Core registration flow.
package ocpp21

import (
	"context"
	"encoding/json"
	"errors"
	"sync"
	"sync/atomic"

	"github.com/seanlee0923/ocpp/csms"
	"github.com/seanlee0923/ocpp/protocol"
	"github.com/seanlee0923/ocpp/v21"
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
type NotifyChargingLimitHandler func(context.Context, *csms.Session, v21.NotifyChargingLimitRequest) (v21.NotifyChargingLimitConfirmation, error)
type ClearedChargingLimitHandler func(context.Context, *csms.Session, v21.ClearedChargingLimitRequest) (v21.ClearedChargingLimitConfirmation, error)
type ReportChargingProfilesHandler func(context.Context, *csms.Session, v21.ReportChargingProfilesRequest) (v21.ReportChargingProfilesConfirmation, error)
type NotifyPriorityChargingHandler func(context.Context, *csms.Session, v21.NotifyPriorityChargingRequest) (v21.NotifyPriorityChargingConfirmation, error)
type PullDynamicScheduleUpdateHandler func(context.Context, *csms.Session, v21.PullDynamicScheduleUpdateRequest) (v21.PullDynamicScheduleUpdateConfirmation, error)
type ReportDERControlHandler func(context.Context, *csms.Session, v21.ReportDERControlRequest) (v21.ReportDERControlConfirmation, error)
type NotifyDERAlarmHandler func(context.Context, *csms.Session, v21.NotifyDERAlarmRequest) (v21.NotifyDERAlarmConfirmation, error)
type NotifyDERStartStopHandler func(context.Context, *csms.Session, v21.NotifyDERStartStopRequest) (v21.NotifyDERStartStopConfirmation, error)
type NotifySettlementHandler func(context.Context, *csms.Session, v21.NotifySettlementRequest) (v21.NotifySettlementConfirmation, error)
type VatNumberValidationHandler func(context.Context, *csms.Session, v21.VatNumberValidationRequest) (v21.VatNumberValidationConfirmation, error)
type BatterySwapHandler func(context.Context, *csms.Session, v21.BatterySwapRequest) (v21.BatterySwapConfirmation, error)
type NotifyPeriodicEventStreamHandler func(context.Context, *csms.Session, v21.NotifyPeriodicEventStreamRequest) error
type OpenPeriodicEventStreamHandler func(context.Context, *csms.Session, v21.OpenPeriodicEventStreamRequest) (v21.OpenPeriodicEventStreamConfirmation, error)
type ClosePeriodicEventStreamHandler func(context.Context, *csms.Session, v21.ClosePeriodicEventStreamRequest) (v21.ClosePeriodicEventStreamConfirmation, error)
type SignCertificateHandler func(context.Context, *csms.Session, v21.SignCertificateRequest) (v21.SignCertificateConfirmation, error)
type Get15118EVCertificateHandler func(context.Context, *csms.Session, v21.Get15118EVCertificateRequest) (v21.Get15118EVCertificateConfirmation, error)
type GetCertificateStatusHandler func(context.Context, *csms.Session, v21.GetCertificateStatusRequest) (v21.GetCertificateStatusConfirmation, error)
type GetCertificateChainStatusHandler func(context.Context, *csms.Session, v21.GetCertificateChainStatusRequest) (v21.GetCertificateChainStatusConfirmation, error)
type FirmwareStatusNotificationHandler func(context.Context, *csms.Session, v21.FirmwareStatusNotificationRequest) (v21.FirmwareStatusNotificationConfirmation, error)
type PublishFirmwareStatusNotificationHandler func(context.Context, *csms.Session, v21.PublishFirmwareStatusNotificationRequest) (v21.PublishFirmwareStatusNotificationConfirmation, error)

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
func (profile *Profile) HandleNotifySettlement(handler NotifySettlementHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v21.NotifySettlementRequest, v21.NotifySettlementConfirmation](handler))
}
func (profile *Profile) HandleVatNumberValidation(handler VatNumberValidationHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v21.VatNumberValidationRequest, v21.VatNumberValidationConfirmation](handler))
}
func (profile *Profile) HandleBatterySwap(handler BatterySwapHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v21.BatterySwapRequest, v21.BatterySwapConfirmation](handler))
}
func (profile *Profile) HandleNotifyPeriodicEventStream(handler NotifyPeriodicEventStreamHandler) error {
	return csms.HandleSend(profile.router, csms.TypedSendHandler[v21.NotifyPeriodicEventStreamRequest](handler))
}
func (profile *Profile) HandleOpenPeriodicEventStream(handler OpenPeriodicEventStreamHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v21.OpenPeriodicEventStreamRequest, v21.OpenPeriodicEventStreamConfirmation](handler))
}
func (profile *Profile) HandleClosePeriodicEventStream(handler ClosePeriodicEventStreamHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v21.ClosePeriodicEventStreamRequest, v21.ClosePeriodicEventStreamConfirmation](handler))
}
func (profile *Profile) HandleSignCertificate(handler SignCertificateHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v21.SignCertificateRequest, v21.SignCertificateConfirmation](handler))
}
func (profile *Profile) HandleGet15118EVCertificate(handler Get15118EVCertificateHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v21.Get15118EVCertificateRequest, v21.Get15118EVCertificateConfirmation](handler))
}
func (profile *Profile) HandleGetCertificateStatus(handler GetCertificateStatusHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v21.GetCertificateStatusRequest, v21.GetCertificateStatusConfirmation](handler))
}
func (profile *Profile) HandleGetCertificateChainStatus(handler GetCertificateChainStatusHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v21.GetCertificateChainStatusRequest, v21.GetCertificateChainStatusConfirmation](handler))
}
func (profile *Profile) HandleFirmwareStatusNotification(handler FirmwareStatusNotificationHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v21.FirmwareStatusNotificationRequest, v21.FirmwareStatusNotificationConfirmation](handler))
}
func (profile *Profile) HandlePublishFirmwareStatusNotification(handler PublishFirmwareStatusNotificationHandler) error {
	return csms.Handle(profile.router, csms.TypedHandler[v21.PublishFirmwareStatusNotificationRequest, v21.PublishFirmwareStatusNotificationConfirmation](handler))
}

func (profile *Profile) CallGetVariables(ctx context.Context, session *csms.Session, request v21.GetVariablesRequest) (v21.GetVariablesConfirmation, error) {
	return callBooted[v21.GetVariablesRequest, v21.GetVariablesConfirmation](profile, ctx, session, request)
}
func (profile *Profile) CallSetVariables(ctx context.Context, session *csms.Session, request v21.SetVariablesRequest) (v21.SetVariablesConfirmation, error) {
	return callBooted[v21.SetVariablesRequest, v21.SetVariablesConfirmation](profile, ctx, session, request)
}
func (profile *Profile) CallGetBaseReport(ctx context.Context, session *csms.Session, request v21.GetBaseReportRequest) (v21.GetBaseReportConfirmation, error) {
	return callBooted[v21.GetBaseReportRequest, v21.GetBaseReportConfirmation](profile, ctx, session, request)
}
func (profile *Profile) CallRequestStartTransaction(ctx context.Context, session *csms.Session, request v21.RequestStartTransactionRequest) (v21.RequestStartTransactionConfirmation, error) {
	return callBooted[v21.RequestStartTransactionRequest, v21.RequestStartTransactionConfirmation](profile, ctx, session, request)
}
func (profile *Profile) CallRequestStopTransaction(ctx context.Context, session *csms.Session, request v21.RequestStopTransactionRequest) (v21.RequestStopTransactionConfirmation, error) {
	return callBooted[v21.RequestStopTransactionRequest, v21.RequestStopTransactionConfirmation](profile, ctx, session, request)
}
func (profile *Profile) CallReset(ctx context.Context, session *csms.Session, request v21.ResetRequest) (v21.ResetConfirmation, error) {
	return callBooted[v21.ResetRequest, v21.ResetConfirmation](profile, ctx, session, request)
}
func (profile *Profile) CallGetChargingProfiles(ctx context.Context, session *csms.Session, request v21.GetChargingProfilesRequest) (v21.GetChargingProfilesConfirmation, error) {
	return callBooted[v21.GetChargingProfilesRequest, v21.GetChargingProfilesConfirmation](profile, ctx, session, request)
}
func (profile *Profile) CallSetChargingProfile(ctx context.Context, session *csms.Session, request v21.SetChargingProfileRequest) (v21.SetChargingProfileConfirmation, error) {
	return callBooted[v21.SetChargingProfileRequest, v21.SetChargingProfileConfirmation](profile, ctx, session, request)
}
func (profile *Profile) CallClearChargingProfile(ctx context.Context, session *csms.Session, request v21.ClearChargingProfileRequest) (v21.ClearChargingProfileConfirmation, error) {
	return callBooted[v21.ClearChargingProfileRequest, v21.ClearChargingProfileConfirmation](profile, ctx, session, request)
}
func (profile *Profile) CallNotifyAllowedEnergyTransfer(ctx context.Context, session *csms.Session, request v21.NotifyAllowedEnergyTransferRequest) (v21.NotifyAllowedEnergyTransferConfirmation, error) {
	return callBooted[v21.NotifyAllowedEnergyTransferRequest, v21.NotifyAllowedEnergyTransferConfirmation](profile, ctx, session, request)
}
func (profile *Profile) CallUsePriorityCharging(ctx context.Context, session *csms.Session, request v21.UsePriorityChargingRequest) (v21.UsePriorityChargingConfirmation, error) {
	return callBooted[v21.UsePriorityChargingRequest, v21.UsePriorityChargingConfirmation](profile, ctx, session, request)
}
func (profile *Profile) CallUpdateDynamicSchedule(ctx context.Context, session *csms.Session, request v21.UpdateDynamicScheduleRequest) (v21.UpdateDynamicScheduleConfirmation, error) {
	return callBooted[v21.UpdateDynamicScheduleRequest, v21.UpdateDynamicScheduleConfirmation](profile, ctx, session, request)
}
func (profile *Profile) CallAFRRSignal(ctx context.Context, session *csms.Session, request v21.AFRRSignalRequest) (v21.AFRRSignalConfirmation, error) {
	return callBooted[v21.AFRRSignalRequest, v21.AFRRSignalConfirmation](profile, ctx, session, request)
}
func (profile *Profile) CallSetDERControl(ctx context.Context, session *csms.Session, request v21.SetDERControlRequest) (v21.SetDERControlConfirmation, error) {
	return callBooted[v21.SetDERControlRequest, v21.SetDERControlConfirmation](profile, ctx, session, request)
}
func (profile *Profile) CallGetDERControl(ctx context.Context, session *csms.Session, request v21.GetDERControlRequest) (v21.GetDERControlConfirmation, error) {
	return callBooted[v21.GetDERControlRequest, v21.GetDERControlConfirmation](profile, ctx, session, request)
}
func (profile *Profile) CallClearDERControl(ctx context.Context, session *csms.Session, request v21.ClearDERControlRequest) (v21.ClearDERControlConfirmation, error) {
	return callBooted[v21.ClearDERControlRequest, v21.ClearDERControlConfirmation](profile, ctx, session, request)
}
func (profile *Profile) CallSetDefaultTariff(ctx context.Context, session *csms.Session, request v21.SetDefaultTariffRequest) (v21.SetDefaultTariffConfirmation, error) {
	return callBooted[v21.SetDefaultTariffRequest, v21.SetDefaultTariffConfirmation](profile, ctx, session, request)
}
func (profile *Profile) CallClearTariffs(ctx context.Context, session *csms.Session, request v21.ClearTariffsRequest) (v21.ClearTariffsConfirmation, error) {
	return callBooted[v21.ClearTariffsRequest, v21.ClearTariffsConfirmation](profile, ctx, session, request)
}
func (profile *Profile) CallChangeTransactionTariff(ctx context.Context, session *csms.Session, request v21.ChangeTransactionTariffRequest) (v21.ChangeTransactionTariffConfirmation, error) {
	return callBooted[v21.ChangeTransactionTariffRequest, v21.ChangeTransactionTariffConfirmation](profile, ctx, session, request)
}
func (profile *Profile) CallCostUpdated(ctx context.Context, session *csms.Session, request v21.CostUpdatedRequest) (v21.CostUpdatedConfirmation, error) {
	return callBooted[v21.CostUpdatedRequest, v21.CostUpdatedConfirmation](profile, ctx, session, request)
}
func (profile *Profile) CallGetTariffs(ctx context.Context, session *csms.Session, request v21.GetTariffsRequest) (v21.GetTariffsConfirmation, error) {
	return callBooted[v21.GetTariffsRequest, v21.GetTariffsConfirmation](profile, ctx, session, request)
}
func (profile *Profile) CallNotifyWebPaymentStarted(ctx context.Context, session *csms.Session, request v21.NotifyWebPaymentStartedRequest) (v21.NotifyWebPaymentStartedConfirmation, error) {
	return callBooted[v21.NotifyWebPaymentStartedRequest, v21.NotifyWebPaymentStartedConfirmation](profile, ctx, session, request)
}
func (profile *Profile) CallRequestBatterySwap(ctx context.Context, session *csms.Session, request v21.RequestBatterySwapRequest) (v21.RequestBatterySwapConfirmation, error) {
	return callBooted[v21.RequestBatterySwapRequest, v21.RequestBatterySwapConfirmation](profile, ctx, session, request)
}
func (profile *Profile) CallGetPeriodicEventStream(ctx context.Context, session *csms.Session, request v21.GetPeriodicEventStreamRequest) (v21.GetPeriodicEventStreamConfirmation, error) {
	return callBooted[v21.GetPeriodicEventStreamRequest, v21.GetPeriodicEventStreamConfirmation](profile, ctx, session, request)
}
func (profile *Profile) CallAdjustPeriodicEventStream(ctx context.Context, session *csms.Session, request v21.AdjustPeriodicEventStreamRequest) (v21.AdjustPeriodicEventStreamConfirmation, error) {
	return callBooted[v21.AdjustPeriodicEventStreamRequest, v21.AdjustPeriodicEventStreamConfirmation](profile, ctx, session, request)
}
func (profile *Profile) CallCertificateSigned(ctx context.Context, session *csms.Session, request v21.CertificateSignedRequest) (v21.CertificateSignedConfirmation, error) {
	return callBooted[v21.CertificateSignedRequest, v21.CertificateSignedConfirmation](profile, ctx, session, request)
}
func (profile *Profile) CallInstallCertificate(ctx context.Context, session *csms.Session, request v21.InstallCertificateRequest) (v21.InstallCertificateConfirmation, error) {
	return callBooted[v21.InstallCertificateRequest, v21.InstallCertificateConfirmation](profile, ctx, session, request)
}
func (profile *Profile) CallDeleteCertificate(ctx context.Context, session *csms.Session, request v21.DeleteCertificateRequest) (v21.DeleteCertificateConfirmation, error) {
	return callBooted[v21.DeleteCertificateRequest, v21.DeleteCertificateConfirmation](profile, ctx, session, request)
}
func (profile *Profile) CallGetInstalledCertificateIds(ctx context.Context, session *csms.Session, request v21.GetInstalledCertificateIdsRequest) (v21.GetInstalledCertificateIdsConfirmation, error) {
	return callBooted[v21.GetInstalledCertificateIdsRequest, v21.GetInstalledCertificateIdsConfirmation](profile, ctx, session, request)
}
func (profile *Profile) CallUpdateFirmware(ctx context.Context, session *csms.Session, request v21.UpdateFirmwareRequest) (v21.UpdateFirmwareConfirmation, error) {
	return callBooted[v21.UpdateFirmwareRequest, v21.UpdateFirmwareConfirmation](profile, ctx, session, request)
}
func (profile *Profile) CallPublishFirmware(ctx context.Context, session *csms.Session, request v21.PublishFirmwareRequest) (v21.PublishFirmwareConfirmation, error) {
	return callBooted[v21.PublishFirmwareRequest, v21.PublishFirmwareConfirmation](profile, ctx, session, request)
}
func (profile *Profile) CallUnpublishFirmware(ctx context.Context, session *csms.Session, request v21.UnpublishFirmwareRequest) (v21.UnpublishFirmwareConfirmation, error) {
	return callBooted[v21.UnpublishFirmwareRequest, v21.UnpublishFirmwareConfirmation](profile, ctx, session, request)
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
