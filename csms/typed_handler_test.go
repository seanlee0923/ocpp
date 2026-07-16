package csms

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"ocpp-go/protocol"
	"ocpp-go/v16"
	"ocpp-go/v201"
)

func TestTypedHandlerValidatesDecodesAndValidatesResponse(t *testing.T) {
	router := NewRouter()
	called := false
	err := Handle(router, func(_ context.Context, _ *Session, request v16.BootNotificationRequest) (v16.BootNotificationConfirmation, error) {
		called = true
		if request.ChargePointVendor != "Example" {
			t.Errorf("vendor = %q", request.ChargePointVendor)
		}
		return validV16BootConfirmation(), nil
	})
	if err != nil {
		t.Fatal(err)
	}
	handler, ok := router.lookup(protocol.OCPP16, "BootNotification")
	if !ok {
		t.Fatal("typed handler was not registered")
	}
	response, err := handler(context.Background(), nil, json.RawMessage(`{"chargePointVendor":"Example","chargePointModel":"AC-22K"}`))
	if err != nil {
		t.Fatal(err)
	}
	if !called {
		t.Fatal("typed handler was not called")
	}
	if _, ok := response.(v16.BootNotificationConfirmation); !ok {
		t.Fatalf("response type = %T", response)
	}
}

func TestTypedHandlerRejectsRawUnknownPropertyBeforeDecode(t *testing.T) {
	router := NewRouter()
	err := Handle(router, func(_ context.Context, _ *Session, _ v16.BootNotificationRequest) (v16.BootNotificationConfirmation, error) {
		t.Fatal("handler called for invalid payload")
		return validV16BootConfirmation(), nil
	})
	if err != nil {
		t.Fatal(err)
	}
	handler, _ := router.lookup(protocol.OCPP16, "BootNotification")
	_, err = handler(context.Background(), nil, json.RawMessage(`{"chargePointVendor":"Example","chargePointModel":"AC-22K","unknown":true}`))
	var callError *CallError
	if !errors.As(err, &callError) || callError.Code != FormationViolation {
		t.Fatalf("error = %#v, want FormationViolation", err)
	}
}

func TestTypedHandlerRejectsMismatchedTypes(t *testing.T) {
	err := Handle(NewRouter(), func(_ context.Context, _ *Session, _ v16.BootNotificationRequest) (v201.BootNotificationConfirmation, error) {
		return v201.BootNotificationConfirmation{}, nil
	})
	if err == nil {
		t.Fatal("Handle succeeded for mismatched versions")
	}
}

func TestTypedHandlerRejectsInvalidConfirmation(t *testing.T) {
	router := NewRouter()
	err := Handle(router, func(_ context.Context, _ *Session, _ v16.BootNotificationRequest) (v16.BootNotificationConfirmation, error) {
		return v16.BootNotificationConfirmation{}, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	handler, _ := router.lookup(protocol.OCPP16, "BootNotification")
	_, err = handler(context.Background(), nil, json.RawMessage(`{"chargePointVendor":"Example","chargePointModel":"AC-22K"}`))
	var callError *CallError
	if !errors.As(err, &callError) || callError.Code != InternalError {
		t.Fatalf("error = %#v, want InternalError", err)
	}
}

func validV16BootConfirmation() v16.BootNotificationConfirmation {
	return v16.BootNotificationConfirmation{
		CurrentTime: time.Now().UTC().Format(time.RFC3339),
		Interval:    300,
		Status:      v16.BootNotificationConfirmationStatusAccepted,
	}
}
