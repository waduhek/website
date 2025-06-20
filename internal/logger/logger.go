package logger

import (
	"context"
	"log/slog"
	"os"
)

type Logger interface {
	Debug(msg string, args ...any)

	DebugContext(ctx context.Context, msg string, args ...any)

	Info(msg string, args ...any)

	InfoContext(ctx context.Context, msg string, args ...any)

	Warn(msg string, args ...any)

	WarnContext(ctx context.Context, msg string, args ...any)

	Error(msg string, args ...any)

	ErrorContext(ctx context.Context, msg string, args ...any)
}

// NewLogger creates a new logger instance.
func NewLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, nil))
}
