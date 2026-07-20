package protocol

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestDecodeCall(t *testing.T) {
	message, err := Decode([]byte(`[2,"abc","BootNotification",{"reason":"PowerUp"}]`))
	if err != nil {
		t.Fatal(err)
	}
	call, ok := message.(Call)
	if !ok || call.ID != "abc" || call.Action != "BootNotification" {
		t.Fatalf("unexpected message: %#v", message)
	}
}

func TestSendRoundTrip(t *testing.T) {
	want := Send{ID: "stream-1", Action: "NotifyPeriodicEventStream", Payload: json.RawMessage(`{"id":1}`)}
	data, err := Encode(want)
	if err != nil {
		t.Fatal(err)
	}
	message, err := Decode(data)
	if err != nil {
		t.Fatal(err)
	}
	got, ok := message.(Send)
	if !ok || got.ID != want.ID || got.Action != want.Action || string(got.Payload) != string(want.Payload) {
		t.Fatalf("SEND = %#v", message)
	}
}

func TestDecodeRejectsInvalidMessages(t *testing.T) {
	tests := []string{
		`{}`,
		`[]`,
		`[2,"id","Action"]`,
		`[2,"id","Action",[]]`,
		`[9,"id",{}]`,
		`[2,"id","Action",{"value":1,"value":2}]`,
		`[2,"1234567890123456789012345678901234567","Action",{}]`,
	}
	for _, input := range tests {
		if _, err := Decode([]byte(input)); !errors.Is(err, ErrInvalidMessage) {
			t.Errorf("Decode(%s) error = %v", input, err)
		}
	}
}

func TestDecodeRejectsInvalidUTF8(t *testing.T) {
	input := append([]byte(`[2,"id","Action",{"value":"`), 0xff)
	input = append(input, []byte(`"}]`)...)
	if _, err := Decode(input); !errors.Is(err, ErrInvalidMessage) {
		t.Fatalf("Decode error = %v", err)
	}
}

func TestDecodePreservesIDForUnsupportedMessageType(t *testing.T) {
	_, err := Decode([]byte(`[9,"future-1",{}]`))
	var unsupported *UnsupportedMessageTypeError
	if !errors.As(err, &unsupported) {
		t.Fatalf("error = %T %v, want UnsupportedMessageTypeError", err, err)
	}
	if unsupported.ID != "future-1" || unsupported.Type != 9 {
		t.Fatalf("unsupported message = %#v", unsupported)
	}

	_, err = Decode([]byte(`[9,42,{}]`))
	if !errors.As(err, &unsupported) || unsupported.ID != "-1" {
		t.Fatalf("invalid ID fallback error = %#v", err)
	}
}

func TestEncodeCallErrorUsesEmptyDetailsObject(t *testing.T) {
	data, err := Encode(CallError{ID: "abc", Code: "NotImplemented", Description: "missing"})
	if err != nil {
		t.Fatal(err)
	}
	var fields []json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		t.Fatal(err)
	}
	if string(fields[4]) != `{}` {
		t.Fatalf("details = %s", fields[4])
	}
}

// TestEncodeRejectsInvalidMessages exercises Encode's error-return branch
// for each of the four Message types — unlike Decode's negative paths
// (already covered by TestDecodeRejectsInvalidMessages), none of these
// were exercised by any prior test.
func TestEncodeRejectsInvalidMessages(t *testing.T) {
	overlong := func(n int) string {
		b := make([]byte, n)
		for i := range b {
			b[i] = 'a'
		}
		return string(b)
	}
	tests := []struct {
		name    string
		message Message
	}{
		{"Call empty ID", Call{ID: "", Action: "Reset", Payload: json.RawMessage(`{}`)}},
		{"Call overlong action", Call{ID: "abc", Action: overlong(MaxActionLength + 1), Payload: json.RawMessage(`{}`)}},
		{"Call non-object payload", Call{ID: "abc", Action: "Reset", Payload: json.RawMessage(`[1,2]`)}},
		{"Send empty ID", Send{ID: "", Action: "NotifyPeriodicEventStream", Payload: json.RawMessage(`{}`)}},
		{"Send non-object payload", Send{ID: "abc", Action: "NotifyPeriodicEventStream", Payload: json.RawMessage(`"x"`)}},
		{"CallResult empty ID", CallResult{ID: "", Payload: json.RawMessage(`{}`)}},
		{"CallResult non-object payload", CallResult{ID: "abc", Payload: json.RawMessage(`null`)}},
		{"CallError empty ID", CallError{ID: "", Code: "NotImplemented", Description: "x"}},
		{"CallError overlong code", CallError{ID: "abc", Code: overlong(MaxErrorCodeLength + 1), Description: "x"}},
		{"CallError overlong description", CallError{ID: "abc", Code: "NotImplemented", Description: overlong(MaxErrorDescriptionLength + 1)}},
		{"CallError non-object details", CallError{ID: "abc", Code: "NotImplemented", Description: "x", Details: json.RawMessage(`[1]`)}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if _, err := Encode(test.message); !errors.Is(err, ErrInvalidMessage) {
				t.Errorf("Encode error = %v, want ErrInvalidMessage", err)
			}
		})
	}
}

// unsupportedMessage is a Message implementation Encode has never seen,
// exercising its default case.
type unsupportedMessage struct{}

func (unsupportedMessage) Type() MessageTypeID { return 0 }
func (unsupportedMessage) UniqueID() string    { return "x" }
func (unsupportedMessage) ocppMessage()        {}

func TestEncodeRejectsUnsupportedMessageType(t *testing.T) {
	if _, err := Encode(unsupportedMessage{}); err == nil {
		t.Fatal("Encode did not reject an unsupported Message implementation")
	}
}
