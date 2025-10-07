package telemetry

import "context"

type Logger interface {
	// Debug logs a message at debug level.
	Debug(msg string, args ...any)

	// DebugContext logs a message at debug level with a context.
	DebugContext(ctx context.Context, msg string, args ...any)

	// Info logs a message at info level.
	Info(msg string, args ...any)

	// InfoContext logs a message at info level with a context.
	InfoContext(ctx context.Context, msg string, args ...any)

	// Warn logs a message at warning level.
	Warn(msg string, args ...any)

	// WarnContext logs a message at warning level with a context.
	WarnContext(ctx context.Context, msg string, args ...any)

	// Error logs a message at error level.
	Error(msg string, args ...any)

	// ErrorContext logs a message at error level with a context.
	ErrorContext(ctx context.Context, msg string, args ...any)
}
