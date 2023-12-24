package logging

import (
	"log/slog"
	"time"

	"github.com/m-mizutani/clog"
)

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
