package main

import (
	"log/slog"
	"os"
)

type CardInfo struct {
	CardNumber string
	ExpiryDate string
	CVV        string
}

func main() {
	// TODO separate

	logLevel := &slog.LevelVar{} // It doesn't have any exported field
	// ```go
	//	type LevelVar struct {
	//		val atomic.Int64
	//	}
	// ```

	slog.Debug("Handle payment transaction1")
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     logLevel, // You can use slog.LevelInfo, slog.LevelDebug..
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == "password" {
				a.Value = slog.StringValue("*****")
			}
			for _, g := range groups {
				if g == "CardInfo" {
					a.Value = slog.StringValue("XXXXX")
				}
			}
			return a
		},
	}

	slog.Debug("Handle payment transaction2")
	slog.Info("Card infomation was wrong",
		slog.String("password", "we are gophers"),
		slog.Group("CardInfo",
			slog.String("CardNumber", "4242424242424242"),
			slog.String("ExpiryDate", "06/24"),
			slog.String("CVV", "000"),
		),
	)

	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, opts)))
	logLevel.Set(slog.LevelDebug)
	slog.Debug("Handle payment transaction3")

	// Now, LogLebel is debug and souce is shown up!
	// But... it might be
	slog.Info("Card infomation was wrong",
		slog.String("password", "we are gophers"),
		slog.Group("cardInfo",
			slog.String("cardNumber", "4242424242424242"),
			slog.String("expiryDate", "06/24"),
			slog.String("cvv", "000"),
		),
	)
}
