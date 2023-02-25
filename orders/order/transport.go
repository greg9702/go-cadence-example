package order

import (
	"context"
	"encoding/json"
	"github.com/greg9702/go-cadence-example/pkg/encoder"
	"github.com/greg9702/go-cadence-example/pkg/middleware"
	"github.com/greg9702/go-cadence-example/pkg/util"
	"net/http"

	"github.com/go-kit/kit/endpoint"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHttpServer(svc Service, logger kitlog.Logger) *mux.Router {
	options := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerFinalizer(util.NewServerFinalizer(logger)),
	}

	var createOrderEP endpoint.Endpoint
	createOrderEP = makeCreateOrderEndpoint(svc)
	createOrderEP = middleware.LoggingMiddleware(kitlog.With(logger, "method", "createOrder"))(createOrderEP)

	createOrderHandler := kithttp.NewServer(
		createOrderEP,
		decodeCreateOrderRequest,
		encoder.EncodeJSONResponse,
		options...,
	)

	r := mux.NewRouter()
	r.Methods("POST").Path("/v1/order").Handler(createOrderHandler)

	return r
}

func decodeCreateOrderRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request createOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
