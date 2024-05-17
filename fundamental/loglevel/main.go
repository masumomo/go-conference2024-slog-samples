package main

import (
	"log"
	"log/slog"
)

func main() {
	slog.Info("Hello Info message!🖐️")
	// No output! because default level is Info
	slog.Debug("Hello Debug message?🤔")
	slog.Warn("Hello Warn message!😵")

	// Traditional log
	log.Print("Hello World!🖐️")
	log.Fatal("End of the world!😇 ")
}
