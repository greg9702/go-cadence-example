package logger

import (
	"context"
	"os"

	"github.com/go-kit/log"
	"github.com/google/uuid"
	"github.com/greg9702/go-cadence-example/pkg"
)

func GetTracedLog(ctx context.Context) log.Logger {
	logger, ok := ctx.Value(pkg.TracedLoggerKey).(log.Logger)
	if !ok {
		traceID := uuid.New()

		tracedLogger := log.NewLogfmtLogger(os.Stderr)
		tracedLogger = log.With(tracedLogger, pkg.TraceIDKey, traceID)

		return tracedLogger
	}

	return logger
}