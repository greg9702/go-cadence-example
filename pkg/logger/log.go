package logger

import (
	"context"

	"github.com/go-kit/log"
	"github.com/google/uuid"
	"github.com/greg9702/go-cadence-example/pkg"
)

func GetTracedLog(ctx context.Context) log.Logger {
	logger, ok := ctx.Value(pkg.TracedLoggerKey).(log.Logger)
	if !ok {
		traceID := uuid.New()

		tracedLogger := log.With(logger, pkg.TraceIDKey, traceID)
		newCtx := context.WithValue(ctx, pkg.TracedLoggerKey, tracedLogger)
		ctx = newCtx

		return tracedLogger
	}

	return logger
}
