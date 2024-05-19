package main

import (
	"log/slog"
	"os"

	myhandler "github.com/masumomo/goconference2024_slog_demos/customhandler/step1"
)

func main() {
	handler := myhandler.New(os.Stdout, nil)
	logger := slog.New(handler).
		With("serviceName", "payment service"). // Output service name
		WithGroup("ProcessTxnFunc")             // Output within function name

	logger.Info("New payment",
		slog.Group("Transaction",
			slog.Int64("JPYAmount", 980),
			slog.String("Marchant", "Abema.tv"),
		),
	)
	logger.Warn("Unknown currency", slog.String("CurrenyCode", "???"))
}
