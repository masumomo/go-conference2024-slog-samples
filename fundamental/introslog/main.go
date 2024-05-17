package main

import (
	"log/slog"
)

func main() {
	slog.Info("Hello world!", slog.String("From", "Japan"), slog.Float64("Go ver", 1.21))
}
