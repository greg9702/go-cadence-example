package middleware

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/google/uuid"
	"github.com/greg9702/go-cadence-example/pkg"
	"time"
)

func LoggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			startTime := time.Now().UnixMilli()

			traceID := uuid.New()
			logger = log.With(logger, pkg.TraceIDKey, traceID)
			ctx = context.WithValue(ctx, pkg.TracedLoggerKey, logger)

			defer func() {
				endTime := time.Now().UnixMilli()
				logger.Log("Took", fmt.Sprintf("%dms", endTime-startTime))
			}()
			return next(ctx, request)
		}
	}
}
