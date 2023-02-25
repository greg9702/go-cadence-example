package logger

import (
	"context"
	"github.com/go-kit/log"
	"github.com/greg9702/go-cadence-example/pkg"
)

func GetTracedLog(ctx context.Context) log.Logger {
	logger := ctx.Value(pkg.TracedLoggerKey).(log.Logger)
	return logger
}
