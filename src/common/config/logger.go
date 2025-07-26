package config

import (
	"log/slog"
	"os"
)

func InitLogger(lvl string) *slog.Logger {
	level := parseLogLevel(lvl)
	logger := slog.New(slog.NewJSONHandler(os.Stdout,
		&slog.HandlerOptions{
			Level:     level,
			AddSource: true}))
	slog.SetDefault(logger)
	logger.Info("Logger initialized", slog.String("level", lvl))

	return logger
}

func parseLogLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelDebug

	}
}
