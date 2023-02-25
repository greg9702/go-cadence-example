package util

import (
	"context"
	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"net/http"
)

func NewServerFinalizer(logger kitlog.Logger) kithttp.ServerFinalizerFunc {
	return func(ctx context.Context, code int, r *http.Request) {
		logger.Log("status", code, "path", r.RequestURI, "method", r.Method)
	}
}
