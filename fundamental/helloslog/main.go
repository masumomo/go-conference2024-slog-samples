package main

import (
	"log"
	"log/slog"
	"os"
	"time"
)

func main() {
	message := "Hello"
	executedTime := time.Now()
	textLogger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	textLogger.Info("Message by textLogger", "message", message, "executedTime", executedTime)

	// Let's try to pass the same arguments to standard log
	log.Print("Message by log", "message", message, "executedTime", executedTime)
	slog.SetDefault(textLogger)

	// Let's try it again after changing the default
	log.Print("Message by log after SetDefault", "message", message, "executedTime", executedTime)

	jsonLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	jsonLogger.Info("Message by jsonLogger", "message", message, "executedTime", executedTime)
	slog.SetDefault(jsonLogger)

	// Let's try it with json logger
	log.Print("Message by default log", "message", message, "executedTime", executedTime)
	slog.Info("Message by default slog", "message", message, "executedTime", executedTime)

	// Meaningless to use standard log after introducing slog in your code,
	// But Go garanteed backforword compatability.
	// Any log can be default level. (We might be able to configure, but INFO it default)
	log.Fatal("Fatal Message by log after SetDefault", "message", message, "executedTime", executedTime)
}
