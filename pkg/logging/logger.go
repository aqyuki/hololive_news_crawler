package logging

import (
	"context"
	"log/slog"
	"time"

	"github.com/m-mizutani/clog"
)

type contextKey string

const loggerKey = contextKey("logger")

// NewLogger creates a new logger
func NewLogger() *slog.Logger {
	var handler slog.Handler = clog.New(
		clog.WithTimeFmt(time.DateTime),
		clog.WithColor(true),
		clog.WithLevel(slog.LevelInfo),
		clog.WithPrinter(clog.PrettyPrinter),
	)
	return slog.New(handler)
}

// WithLogger returns a new context with the given logger
func WithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

// FromContext returns the logger from the given context
func FromContext(ctx context.Context) *slog.Logger {
	if logger, ok := ctx.Value(loggerKey).(*slog.Logger); ok {
		return logger
	}
	return NewLogger()
}
