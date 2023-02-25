package order

import (
	"context"
	"github.com/greg9702/go-cadence-example/pkg/logger"

	"github.com/go-kit/kit/endpoint"
)

type createOrderRequest struct {
}

type createOrderResponse struct {
	Message string `json:"message"`
}

func makeCreateOrderEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		log := logger.GetTracedLog(ctx)
		log.Log("XDDDDDD")
		return createOrderResponse{"Hello"}, nil
	}
}
