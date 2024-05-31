package main

import (
	"log/slog"
	"os"

	"github.com/masumomo/goconference2024_slog_demos/myhandler/final"
	"github.com/masumomo/goconference2024_slog_demos/myhandler/step1"
	"github.com/masumomo/goconference2024_slog_demos/myhandler/step2"
	"github.com/masumomo/goconference2024_slog_demos/myhandler/step3"
	"github.com/masumomo/goconference2024_slog_demos/myhandler/step4"
	"github.com/masumomo/goconference2024_slog_demos/myhandler/step5"
)

func main() {
	handler := final.New(os.Stdout, nil)
	logger := slog.New(handler).
		With("serviceName", "payment service"). // Output service name
		WithGroup("ProcessTxnFunc")             // Output within function name

	if len(os.Args) > 1 { // For demo, we have switch the handler if version is provided
		switch os.Args[1] {
		case "step1":
			handler := step1.New(os.Stdout, nil)
			logger = slog.New(handler).
				With("serviceName", "payment service"). // Output service name
				WithGroup("ProcessTxnFunc")             // Output within function name

		case "step2":
			handler := step2.New(os.Stdout, nil)
			logger = slog.New(handler).
				With("serviceName", "payment service"). // Output service name
				WithGroup("ProcessTxnFunc")             // Output within function name

		case "step3":
			handler := step3.New(os.Stdout, nil)
			logger = slog.New(handler).
				With("serviceName", "payment service"). // Output service name
				WithGroup("ProcessTxnFunc")             // Output within function name

		case "step4":
			handler := step4.New(os.Stdout, nil)
			logger = slog.New(handler).
				With("serviceName", "payment service"). // Output service name
				WithGroup("ProcessTxnFunc")             // Output within function name

		case "step5":
			handler := step5.New(os.Stdout, nil)
			logger = slog.New(handler).
				With("serviceName", "payment service"). // Output service name
				WithGroup("ProcessTxnFunc")             // Output within function name

		}
	}
	slog.SetDefault(logger)

	slog.Info("New payment",
		slog.Group("Transaction",
			slog.Int64("JPYAmount", 980),
			slog.String("Marchant", "Abema.tv"),
		),
	)
	// Some logic
	slog.Warn("Unknown currency", slog.String("CurrenyCode", "???"))
}
