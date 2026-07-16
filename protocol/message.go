package protocol

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"unicode/utf8"

	"ocpp-go/internal/validation"
)

// IsValidationError reports whether err is a generated OCPP payload
// constraint violation.
func IsValidationError(err error) bool { return validation.IsError(err) }

type MessageTypeID int

const (
	CallType       MessageTypeID = 2
	CallResultType MessageTypeID = 3
	CallErrorType  MessageTypeID = 4
	SendType       MessageTypeID = 6
)

var ErrInvalidMessage = errors.New("invalid OCPP message")

const (
	MaxUniqueIDLength         = 36
	MaxActionLength           = 64
	MaxErrorCodeLength        = 64
	MaxErrorDescriptionLength = 255
)

type Message interface {
	Type() MessageTypeID
	UniqueID() string
}

// Payload is implemented by every generated OCPP request and confirmation.
// It describes the wire action and version and validates itself against the
// corresponding official OCPP JSON Schema.
type Payload interface {
	ActionName() string
	Version() Version
	Direction() PayloadDirection
	SchemaName() string
	Validate() error
	ValidateJSON([]byte) error
}

type PayloadDirection uint8

const (
	RequestPayload PayloadDirection = iota + 1
	ConfirmationPayload
)

type Call struct {
	ID      string
	Action  string
	Payload json.RawMessage
}

// Send is an OCPP 2.1 unconfirmed message. The receiver must not return a
// CALLRESULT or CALLERROR.
type Send struct {
	ID      string
	Action  string
	Payload json.RawMessage
}

func (m Call) Type() MessageTypeID { return CallType }
func (m Call) UniqueID() string    { return m.ID }

func (m Send) Type() MessageTypeID { return SendType }
func (m Send) UniqueID() string    { return m.ID }

type CallResult struct {
	ID      string
	Payload json.RawMessage
}

func (m CallResult) Type() MessageTypeID { return CallResultType }
func (m CallResult) UniqueID() string    { return m.ID }

type CallError struct {
	ID          string
	Code        string
	Description string
	Details     json.RawMessage
}

func (m CallError) Type() MessageTypeID { return CallErrorType }
func (m CallError) UniqueID() string    { return m.ID }

func Decode(data []byte) (Message, error) {
	if err := validateJSONSyntax(data); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidMessage, err)
	}
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()

	var fields []json.RawMessage
	if err := decoder.Decode(&fields); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidMessage, err)
	}
	var trailing any
	if err := decoder.Decode(&trailing); err != io.EOF {
		return nil, fmt.Errorf("%w: trailing JSON value", ErrInvalidMessage)
	}
	if len(fields) == 0 {
		return nil, fmt.Errorf("%w: empty array", ErrInvalidMessage)
	}

	var messageType MessageTypeID
	if err := json.Unmarshal(fields[0], &messageType); err != nil {
		return nil, fmt.Errorf("%w: invalid message type", ErrInvalidMessage)
	}

	switch messageType {
	case CallType:
		if len(fields) != 4 {
			return nil, fmt.Errorf("%w: CALL requires 4 fields", ErrInvalidMessage)
		}
		var id, action string
		if err := json.Unmarshal(fields[1], &id); err != nil || !validLength(id, MaxUniqueIDLength) {
			return nil, fmt.Errorf("%w: invalid CALL unique ID", ErrInvalidMessage)
		}
		if err := json.Unmarshal(fields[2], &action); err != nil || !validLength(action, MaxActionLength) {
			return nil, fmt.Errorf("%w: invalid CALL action", ErrInvalidMessage)
		}
		if !isObject(fields[3]) {
			return nil, fmt.Errorf("%w: CALL payload must be an object", ErrInvalidMessage)
		}
		return Call{ID: id, Action: action, Payload: fields[3]}, nil
	case SendType:
		if len(fields) != 4 {
			return nil, fmt.Errorf("%w: SEND requires 4 fields", ErrInvalidMessage)
		}
		var id, action string
		if err := json.Unmarshal(fields[1], &id); err != nil || !validLength(id, MaxUniqueIDLength) {
			return nil, fmt.Errorf("%w: invalid SEND unique ID", ErrInvalidMessage)
		}
		if err := json.Unmarshal(fields[2], &action); err != nil || !validLength(action, MaxActionLength) {
			return nil, fmt.Errorf("%w: invalid SEND action", ErrInvalidMessage)
		}
		if !isObject(fields[3]) {
			return nil, fmt.Errorf("%w: SEND payload must be an object", ErrInvalidMessage)
		}
		return Send{ID: id, Action: action, Payload: fields[3]}, nil
	case CallResultType:
		if len(fields) != 3 {
			return nil, fmt.Errorf("%w: CALLRESULT requires 3 fields", ErrInvalidMessage)
		}
		var id string
		if err := json.Unmarshal(fields[1], &id); err != nil || !validLength(id, MaxUniqueIDLength) || !isObject(fields[2]) {
			return nil, fmt.Errorf("%w: invalid CALLRESULT", ErrInvalidMessage)
		}
		return CallResult{ID: id, Payload: fields[2]}, nil
	case CallErrorType:
		if len(fields) != 5 {
			return nil, fmt.Errorf("%w: CALLERROR requires 5 fields", ErrInvalidMessage)
		}
		var id, code, description string
		if json.Unmarshal(fields[1], &id) != nil || !validLength(id, MaxUniqueIDLength) || json.Unmarshal(fields[2], &code) != nil || !validLength(code, MaxErrorCodeLength) || json.Unmarshal(fields[3], &description) != nil || utf8.RuneCountInString(description) > MaxErrorDescriptionLength || !isObject(fields[4]) {
			return nil, fmt.Errorf("%w: invalid CALLERROR", ErrInvalidMessage)
		}
		return CallError{ID: id, Code: code, Description: description, Details: fields[4]}, nil
	default:
		return nil, fmt.Errorf("%w: unknown message type %d", ErrInvalidMessage, messageType)
	}
}

func Encode(message Message) ([]byte, error) {
	switch m := message.(type) {
	case Call:
		if !validLength(m.ID, MaxUniqueIDLength) || !validLength(m.Action, MaxActionLength) || !isObject(normalizedObject(m.Payload)) {
			return nil, fmt.Errorf("%w: invalid CALL", ErrInvalidMessage)
		}
		return json.Marshal([]any{CallType, m.ID, m.Action, normalizedObject(m.Payload)})
	case Send:
		if !validLength(m.ID, MaxUniqueIDLength) || !validLength(m.Action, MaxActionLength) || !isObject(normalizedObject(m.Payload)) {
			return nil, fmt.Errorf("%w: invalid SEND", ErrInvalidMessage)
		}
		return json.Marshal([]any{SendType, m.ID, m.Action, normalizedObject(m.Payload)})
	case CallResult:
		if !validLength(m.ID, MaxUniqueIDLength) || !isObject(normalizedObject(m.Payload)) {
			return nil, fmt.Errorf("%w: invalid CALLRESULT", ErrInvalidMessage)
		}
		return json.Marshal([]any{CallResultType, m.ID, normalizedObject(m.Payload)})
	case CallError:
		if !validLength(m.ID, MaxUniqueIDLength) || !validLength(m.Code, MaxErrorCodeLength) || utf8.RuneCountInString(m.Description) > MaxErrorDescriptionLength || !isObject(normalizedObject(m.Details)) {
			return nil, fmt.Errorf("%w: invalid CALLERROR", ErrInvalidMessage)
		}
		return json.Marshal([]any{CallErrorType, m.ID, m.Code, m.Description, normalizedObject(m.Details)})
	default:
		return nil, fmt.Errorf("unsupported message type %T", message)
	}
}

func validLength(value string, maximum int) bool {
	length := utf8.RuneCountInString(value)
	return length > 0 && length <= maximum
}

func validateJSONSyntax(data []byte) error {
	if !utf8.Valid(data) {
		return errors.New("JSON contains invalid UTF-8")
	}
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	if err := parseJSONValue(decoder); err != nil {
		return err
	}
	if _, err := decoder.Token(); err != io.EOF {
		if err == nil {
			return errors.New("JSON contains a trailing value")
		}
		return err
	}
	return nil
}

func parseJSONValue(decoder *json.Decoder) error {
	token, err := decoder.Token()
	if err != nil {
		return err
	}
	delimiter, ok := token.(json.Delim)
	if !ok {
		return nil
	}
	switch delimiter {
	case '{':
		seen := make(map[string]struct{})
		for decoder.More() {
			keyToken, err := decoder.Token()
			if err != nil {
				return err
			}
			key, ok := keyToken.(string)
			if !ok {
				return errors.New("JSON object key is not a string")
			}
			if _, duplicate := seen[key]; duplicate {
				return fmt.Errorf("duplicate JSON object key %q", key)
			}
			seen[key] = struct{}{}
			if err := parseJSONValue(decoder); err != nil {
				return err
			}
		}
		closing, err := decoder.Token()
		if err != nil {
			return err
		}
		if closing != json.Delim('}') {
			return errors.New("invalid JSON object closing delimiter")
		}
	case '[':
		for decoder.More() {
			if err := parseJSONValue(decoder); err != nil {
				return err
			}
		}
		closing, err := decoder.Token()
		if err != nil {
			return err
		}
		if closing != json.Delim(']') {
			return errors.New("invalid JSON array closing delimiter")
		}
	default:
		return fmt.Errorf("unexpected JSON delimiter %q", delimiter)
	}
	return nil
}

func isObject(raw json.RawMessage) bool {
	var object map[string]json.RawMessage
	return json.Unmarshal(raw, &object) == nil && object != nil
}

func normalizedObject(raw json.RawMessage) json.RawMessage {
	if len(raw) == 0 {
		return json.RawMessage(`{}`)
	}
	return raw
}
