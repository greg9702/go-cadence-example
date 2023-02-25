package order

import (
	"context"
	"encoding/json"
	"github.com/greg9702/go-cadence-example/pkg/encoder"
	"github.com/greg9702/go-cadence-example/pkg/errors"
	"github.com/greg9702/go-cadence-example/pkg/middleware"
	"net/http"

	"github.com/go-kit/kit/endpoint"

	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHttpServer(svc Service, logger kitlog.Logger) *mux.Router {
	options := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(encoder.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encoder.ErrorEncoder(logger)),
	}

	var createOrderEP endpoint.Endpoint
	createOrderEP = makeCreateOrderEndpoint(svc)
	createOrderEP = middleware.LoggingMiddleware(kitlog.With(logger, "method", "createOrder"))(createOrderEP)

	var getOrderEP endpoint.Endpoint
	getOrderEP = makeGetOrderEndpoint(svc)
	getOrderEP = middleware.LoggingMiddleware(kitlog.With(logger, "method", "getOrder"))(getOrderEP)

	createOrderHandler := kithttp.NewServer(
		createOrderEP,
		decodeCreateOrderRequest,
		encoder.EncodeJSONResponse,
		options...,
	)

	getOrderHandler := kithttp.NewServer(
		getOrderEP,
		decodeGetOrderRequest,
		encoder.EncodeJSONResponse,
		options...,
	)

	r := mux.NewRouter()
	r.Methods("POST").Path("/v1/order").Handler(createOrderHandler)
	r.Methods("GET").Path("/v1/order/{id}").Handler(getOrderHandler)

	return r
}

func decodeCreateOrderRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request createOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	if request.TotalCost == 0 {
		return nil, &errors.ServiceError{
			Code:    http.StatusBadRequest,
			Message: "missing total cost",
		}
	}
	if request.VehicleNo == "" {
		return nil, &errors.ServiceError{
			Code:    http.StatusBadRequest,
			Message: "missing vehicle number",
		}
	}

	return request, nil
}

func decodeGetOrderRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request getOrderRequest
	vars := mux.Vars(r)

	if id, ok := vars["id"]; !ok {
		return nil, &errors.ServiceError{
			Code:    http.StatusBadRequest,
			Message: "missing order id",
		}
	} else {
		request.ID = id
	}

	return request, nil
}
