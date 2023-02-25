package encoder

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/greg9702/go-cadence-example/pkg/errors"
	"net/http"
)

type ErrResponse struct {
	Message string `json:"message"`
}

func ErrorEncoder(logger log.Logger) httptransport.ErrorEncoder {
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		httpCode := http.StatusInternalServerError
		if svcErr, ok := err.(*errors.ServiceError); ok {
			httpCode = svcErr.Code
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(httpCode)
		json.NewEncoder(w).Encode(ErrResponse{
			Message: err.Error(),
		})
	}
}

type LogErrorHandler struct {
	logger log.Logger
}

func NewLogErrorHandler(logger log.Logger) *LogErrorHandler {
	return &LogErrorHandler{
		logger: logger,
	}
}

func (h *LogErrorHandler) Handle(ctx context.Context, err error) {
	if svcErr, ok := err.(*errors.ServiceError); ok {
		h.logger.Log("err", err, "statusCode", svcErr.Code)
	} else {
		h.logger.Log("err", err)
	}
}
