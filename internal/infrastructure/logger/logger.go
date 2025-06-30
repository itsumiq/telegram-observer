package logger

import (
	"log/slog"
	"os"
)

var (
	LevelDebug = slog.LevelDebug
	LevelInfo  = slog.LevelInfo
	LevelError = slog.LevelError
)

func New(level slog.Level) *slog.Logger {
	handler := slog.NewTextHandler(
		os.Stdout,
		&slog.HandlerOptions{Level: level, AddSource: true},
	)
	return slog.New(handler)
}
