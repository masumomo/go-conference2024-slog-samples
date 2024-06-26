package main

import (
	"log/slog"
	"os"
	"time"
)

func main() {
	// Add context for this logger
	srvlogger := slog.New(slog.NewJSONHandler(os.Stdin, nil)).
		With(slog.String("appVer", "1.21")).
		WithGroup("app-log")

	srvlogger.Info("Start processing")
	srvlogger.Info("Go Conference prolosal submitted!",
		slog.Time("submitAt", time.Date(2024, 6, 7, 13, 20, 20, 0, time.UTC)),
		slog.String("SubmittedBy", "Miki"),
		slog.Group("Session",
			slog.String("Title", "Hack everything!"),
			slog.Duration("Length", 20*time.Minute),
		),
	)
}
