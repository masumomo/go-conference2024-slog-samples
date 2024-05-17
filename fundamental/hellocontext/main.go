package main

import (
	"log"
	"log/slog"
)

func main() {
	slog.Info("Hello Info message!🖐️")
	slog.Debug("Hello Debug message?🤔")
	slog.Warn("Hello Warn message!😵")

	// Meaningless to use standard log after introducing slog in your code,
	// But Go garanteed backforword compatability.
	// There is no log level for standard log
	// but just Timestamp is added.
	log.Print("Hello World!🖐️")
	log.Fatal("End of the world!😇 ")
}
