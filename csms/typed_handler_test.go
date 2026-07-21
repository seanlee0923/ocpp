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
	handler, ok := router.lookup(protocol.OCPP16, "BootNotification", callKind)
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
	handler, _ := router.lookup(protocol.OCPP16, "BootNotification", callKind)
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

func TestTypedAfterHandlerRejectsMismatchedTypes(t *testing.T) {
	err := HandleAfter(NewRouter(), func(context.Context, *Session, v16.BootNotificationRequest, v201.BootNotificationConfirmation) error {
		return nil
	})
	if !errors.Is(err, ErrInvalidHandlerRegistration) {
		t.Fatalf("HandleAfter error = %v, want ErrInvalidHandlerRegistration", err)
	}
}

func TestTypedAfterHandlerRejectsInvalidRegistration(t *testing.T) {
	handler := func(context.Context, *Session, v16.BootNotificationRequest, v16.BootNotificationConfirmation) error {
		return nil
	}
	if err := HandleAfter(nil, handler); !errors.Is(err, ErrInvalidHandlerRegistration) {
		t.Fatalf("nil router error = %v", err)
	}
	var nilHandler TypedAfterHandler[v16.BootNotificationRequest, v16.BootNotificationConfirmation]
	if err := HandleAfter(NewRouter(), nilHandler); !errors.Is(err, ErrInvalidHandlerRegistration) {
		t.Fatalf("nil handler error = %v", err)
	}
	if err := HandleAfter(NewRouter(), func(context.Context, *Session, *v16.BootNotificationRequest, v16.BootNotificationConfirmation) error {
		return nil
	}); !errors.Is(err, ErrInvalidHandlerRegistration) {
		t.Fatalf("pointer request error = %v", err)
	}
	if err := HandleAfter(NewRouter(), func(context.Context, *Session, v16.BootNotificationConfirmation, v16.BootNotificationRequest) error {
		return nil
	}); !errors.Is(err, ErrInvalidHandlerRegistration) {
		t.Fatalf("reversed directions error = %v", err)
	}
	if err := HandleAfter(NewRouter(), func(context.Context, *Session, v16.BootNotificationRequest, v16.HeartbeatConfirmation) error {
		return nil
	}); !errors.Is(err, ErrInvalidHandlerRegistration) {
		t.Fatalf("mismatched action error = %v", err)
	}
}

func TestTypedAfterHandlerDecodesAndPropagatesErrors(t *testing.T) {
	router := NewRouter()
	wantErr := errors.New("after failed")
	if err := HandleAfter(router, func(_ context.Context, _ *Session, request v16.BootNotificationRequest, confirmation v16.BootNotificationConfirmation) error {
		if request.ChargePointVendor != "Example" || confirmation.Status != v16.BootNotificationConfirmationStatusAccepted {
			t.Fatalf("request = %#v, confirmation = %#v", request, confirmation)
		}
		return wantErr
	}); err != nil {
		t.Fatal(err)
	}
	handler := router.lookupAfter(protocol.OCPP16, "BootNotification")[0]
	if err := handler(context.Background(), nil, json.RawMessage(`{"chargePointVendor":"Example","chargePointModel":"AC-22K"}`), validV16BootConfirmation()); !errors.Is(err, wantErr) {
		t.Fatalf("callback error = %v, want %v", err, wantErr)
	}
	if err := handler(context.Background(), nil, json.RawMessage(`{`), validV16BootConfirmation()); err == nil {
		t.Fatal("malformed request reached typed After handler")
	}
	if err := handler(context.Background(), nil, json.RawMessage(`{"chargePointVendor":"Example","chargePointModel":"AC-22K"}`), v16.HeartbeatConfirmation{}); err == nil {
		t.Fatal("wrong confirmation type reached typed After handler")
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
	handler, _ := router.lookup(protocol.OCPP16, "BootNotification", callKind)
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
