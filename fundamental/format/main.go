package main

import (
	"log/slog"
	"os"
	"time"
)

func main() {

	userID := "user_id_1"
	executedTime := time.Now()

	slog.Info("Text message, which is default",
		slog.Time("executedTime", executedTime),
		slog.String("userID", userID),
	)
	// Outout text format log
	// 2024/05/10 12:02:07 INFO Text message, which is default executedTime=2024-05-10T12:02....

	jsonLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	jsonLogger.Info("Text message, which is default",
		slog.Time("executedTime", executedTime),
		slog.String("userID", userID),
	)
	// Output JSON format
	// {
	// 	"time":"2024-05-10T12:02:07.002055+03:00",
	// 	"level":"INFO",
	// 	"msg":"JSON message",
	// 	"executedTime":"200 userID=user_id_1024-05-10T12:02:06.999861+03:00",
	// 	"userID":"user_id_1"
	// }

}
