// Command structured-logger connects OCPP metadata to log/slog.
package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/seanlee0923/ocpp/csms"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	server, err := csms.New(csms.Config{
		Logger: csms.LoggerFunc(func(ctx context.Context, record csms.LogRecord) {
			logger.LogAttrs(ctx, slog.LevelInfo, string(record.Event),
				slog.String("identity", record.Identity),
				slog.String("version", string(record.Version)),
				slog.String("action", record.Action),
				slog.String("message_id", record.MessageID),
				slog.String("error_code", string(record.ErrorCode)),
				slog.String("reason", record.Reason),
			)
		}),
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(":8080", server))
}
