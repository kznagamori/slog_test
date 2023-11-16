package main

import (
	"log"
	"os"

	"golang.org/x/exp/slog"
)

func main() {
	logLevel := new(slog.LevelVar)
	ops := slog.HandlerOptions{
		Level: logLevel,
	}

	logLevel.Set(slog.LevelDebug)

	logger := slog.New(slog.NewTextHandler(os.Stdout, &ops))
	slog.SetDefault(logger)
	slog.Info("slog Info message") // デフォルトロガーが適用されるため、プレーンテキストで出力される。
	logger.Debug("Debug message")
	logger.Info("Info message")
	logger.Warn("Warning message")
	logger.Error("Error message")

	slog.SetDefault(logger) // 以降、JSON形式で出力される。
	slog.Info("Info message")
	log.Println("log message by old logger")

	logLevel.Set(slog.LevelWarn)
	logger2 := slog.New(slog.NewJSONHandler(os.Stdout, &ops))
	slog.Info("slog Info message") // デフォルトロガーが適用されるため、プレーンテキストで出力される。
	logger2.Debug("Debug message")
	logger2.Info("Info message")
	logger2.Warn("Warning message")
	logger2.Error("Error message")

}
