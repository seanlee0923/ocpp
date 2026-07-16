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
