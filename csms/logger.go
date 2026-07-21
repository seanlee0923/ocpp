package csms

import (
	"context"

	"github.com/seanlee0923/ocpp/protocol"
)

type LogLevel uint8

const (
	LogDebug LogLevel = iota + 1
	LogInfo
	LogWarn
	LogError
)

type LogEvent string

const (
	LogSessionConnected    LogEvent = "session.connected"
	LogSessionDisconnected LogEvent = "session.disconnected"
	LogCallReceived        LogEvent = "call.received"
	LogCallCompleted       LogEvent = "call.completed"
	LogCallRejected        LogEvent = "call.rejected"
	LogCallAfterFailed     LogEvent = "call.after_failed"
	LogSendReceived        LogEvent = "send.received"
	LogSendCompleted       LogEvent = "send.completed"
	LogSendDropped         LogEvent = "send.dropped"
)

// LogRecord contains protocol metadata only. Payloads, credentials,
// certificates, idTokens and handler error text are intentionally omitted.
type LogRecord struct {
	Level       LogLevel
	Event       LogEvent
	Identity    string
	Version     protocol.Version
	MessageType protocol.MessageTypeID
	MessageID   string
	Action      string
	ErrorCode   ErrorCode
	Reason      string
}

type Logger interface {
	Log(context.Context, LogRecord)
}

type LoggerFunc func(context.Context, LogRecord)

func (f LoggerFunc) Log(ctx context.Context, record LogRecord) { f(ctx, record) }

func (s *Server) log(ctx context.Context, record LogRecord) {
	if s.config.Logger == nil {
		return
	}
	// A diagnostic hook must not be able to take down the protocol server.
	defer func() { _ = recover() }()
	s.config.Logger.Log(ctx, record)
}

func sessionLogRecord(session *Session, level LogLevel, event LogEvent) LogRecord {
	return LogRecord{Level: level, Event: event, Identity: session.Identity(), Version: session.Version()}
}
