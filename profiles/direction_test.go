package profiles_test

import (
	"reflect"
	"sort"
	"strings"
	"testing"

	"ocpp-go/profiles/ocpp16"
	"ocpp-go/profiles/ocpp201"
	"ocpp-go/profiles/ocpp21"
)

// This inventory follows the request initiator defined by the official OCA
// specifications. Handle means Charging Station -> CSMS; Call means CSMS ->
// Charging Station. Keeping the complete public wrapper set here makes a
// direction change or an accidentally added wrapper fail review visibly.
func TestProfileWrapperDirections(t *testing.T) {
	tests := []struct {
		name     string
		profile  any
		inbound  []string
		outbound []string
	}{
		{
			name: "OCPP 1.6", profile: (*ocpp16.Profile)(nil),
			inbound:  []string{"Authorize", "BootNotification", "DataTransfer", "Heartbeat", "MeterValues", "StartTransaction", "StatusNotification", "StopTransaction"},
			outbound: []string{"ChangeAvailability", "DataTransfer", "RemoteStartTransaction", "RemoteStopTransaction", "Reset", "TriggerMessage", "UnlockConnector"},
		},
		{
			name: "OCPP 2.0.1", profile: (*ocpp201.Profile)(nil),
			inbound:  []string{"Authorize", "BootNotification", "Heartbeat", "MeterValues", "NotifyReport", "StatusNotification", "TransactionEvent"},
			outbound: []string{"GetBaseReport", "GetVariables", "RequestStartTransaction", "RequestStopTransaction", "Reset", "SetVariables"},
		},
		{
			name: "OCPP 2.1", profile: (*ocpp21.Profile)(nil),
			inbound: []string{
				"Authorize", "BatterySwap", "BootNotification", "ClearedChargingLimit", "ClosePeriodicEventStream",
				"FirmwareStatusNotification", "Get15118EVCertificate", "GetCertificateChainStatus", "GetCertificateStatus", "Heartbeat",
				"MeterValues", "NotifyChargingLimit", "NotifyDERAlarm", "NotifyDERStartStop", "NotifyEVChargingNeeds",
				"NotifyEVChargingSchedule", "NotifyPeriodicEventStream", "NotifyPriorityCharging", "NotifyReport", "NotifySettlement",
				"OpenPeriodicEventStream", "PublishFirmwareStatusNotification", "PullDynamicScheduleUpdate", "ReportChargingProfiles",
				"ReportDERControl", "SignCertificate", "StatusNotification", "TransactionEvent", "VatNumberValidation",
			},
			outbound: []string{
				"AFRRSignal", "AdjustPeriodicEventStream", "CertificateSigned", "ChangeTransactionTariff", "ClearChargingProfile",
				"ClearDERControl", "ClearTariffs", "CostUpdated", "DeleteCertificate", "GetBaseReport", "GetChargingProfiles",
				"GetDERControl", "GetInstalledCertificateIds", "GetPeriodicEventStream", "GetTariffs", "GetVariables",
				"InstallCertificate", "NotifyAllowedEnergyTransfer", "NotifyWebPaymentStarted", "PublishFirmware", "RequestBatterySwap",
				"RequestStartTransaction", "RequestStopTransaction", "Reset", "SetChargingProfile", "SetDefaultTariff", "SetDERControl",
				"SetVariables", "UnpublishFirmware", "UpdateDynamicSchedule", "UpdateFirmware", "UsePriorityCharging",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotInbound, gotOutbound := wrapperNames(reflect.TypeOf(test.profile))
			assertNames(t, "inbound Handle wrappers", gotInbound, test.inbound)
			assertNames(t, "outbound Call wrappers", gotOutbound, test.outbound)
		})
	}
}

func wrapperNames(profile reflect.Type) (inbound, outbound []string) {
	for i := 0; i < profile.NumMethod(); i++ {
		name := profile.Method(i).Name
		switch {
		case strings.HasPrefix(name, "Handle"):
			inbound = append(inbound, strings.TrimPrefix(name, "Handle"))
		case strings.HasPrefix(name, "Call"):
			outbound = append(outbound, strings.TrimPrefix(name, "Call"))
		}
	}
	return inbound, outbound
}

func assertNames(t *testing.T, label string, got, want []string) {
	t.Helper()
	sort.Strings(got)
	sort.Strings(want)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("%s:\n got: %v\nwant: %v", label, got, want)
	}
}
