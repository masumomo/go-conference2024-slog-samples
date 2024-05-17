package main

import (
	"log/slog"
	"os"
	"time"
)

func main() {
	jsonLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	// Group is useful to chank of fields,
	jsonLogger.Info("Go Conference prolosal submitted!",
		slog.Time("submitAt", time.Date(2024, 6, 7, 13, 20, 20, 0, time.UTC)),
		slog.String("SubmittedBy", "Miki"),
		slog.Group("Session",
			slog.String("Title", "Hack everything!"),
			slog.Duration("Length", 20*time.Minute),
		),
	)
}
