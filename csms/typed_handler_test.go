package csms

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/seanlee0923/ocpp/protocol"
	"github.com/seanlee0923/ocpp/v16"
	"github.com/seanlee0923/ocpp/v201"
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

func TestTypedHandlerClassifiesPayloadConstraintViolations(t *testing.T) {
	router := NewRouter()
	err := Handle(router, func(_ context.Context, _ *Session, _ v16.BootNotificationRequest) (v16.BootNotificationConfirmation, error) {
		t.Fatal("handler called for invalid payload")
		return validV16BootConfirmation(), nil
	})
	if err != nil {
		t.Fatal(err)
	}
	handler, _ := router.lookup(protocol.OCPP16, "BootNotification")
	tests := []struct {
		name string
		raw  string
		code ErrorCode
	}{
		{"unknown property", `{"chargePointVendor":"Example","chargePointModel":"AC-22K","unknown":true}`, PropertyConstraintViolation},
		{"missing required property", `{"chargePointVendor":"Example"}`, OccurenceConstraintViolation},
		{"wrong property type", `{"chargePointVendor":42,"chargePointModel":"AC-22K"}`, TypeConstraintViolation},
		{"trailing JSON value", `{"chargePointVendor":"Example","chargePointModel":"AC-22K"} {}`, FormationViolation},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := handler(context.Background(), nil, json.RawMessage(test.raw))
			var callError *CallError
			if !errors.As(err, &callError) || callError.Code != test.code {
				t.Fatalf("error = %#v, want %s", err, test.code)
			}
		})
	}
}

func TestValidationErrorCodesAreVersionSpecific(t *testing.T) {
	missingRequired := (v16.BootNotificationRequest{}).ValidateJSON([]byte(`{"chargePointVendor":"Example"}`))
	if missingRequired == nil {
		t.Fatal("missing required property was accepted")
	}
	tests := []struct {
		version protocol.Version
		err     error
		want    ErrorCode
	}{
		{protocol.OCPP16, missingRequired, OccurenceConstraintViolation},
		{protocol.OCPP201, missingRequired, OccurrenceConstraintViolation},
		{protocol.OCPP21, missingRequired, OccurrenceConstraintViolation},
		{protocol.OCPP16, errors.New("malformed JSON"), FormationViolation},
		{protocol.OCPP201, errors.New("malformed JSON"), FormatViolation},
		{protocol.OCPP21, errors.New("malformed JSON"), FormatViolation},
	}
	for _, test := range tests {
		if got := validationErrorCode(test.version, test.err); got != test.want {
			t.Errorf("validationErrorCode(%s) = %q, want %q", test.version, got, test.want)
		}
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
