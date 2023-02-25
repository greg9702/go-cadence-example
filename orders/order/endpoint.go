package order

import (
	"context"
	"orders/dao"

	"github.com/go-kit/kit/endpoint"
)

type createOrderRequest struct {
	TotalCost uint32 `json:"total_cost"`
	VehicleNo string `json:"vehicle_no"`
}

type createOrderResponse struct {
	ID string `json:"id"`
}

func makeCreateOrderEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createOrderRequest)
		id, err := svc.CreateOrder(ctx, req.TotalCost, req.VehicleNo)
		return createOrderResponse{
			ID: id,
		}, err
	}
}

type getOrderRequest struct {
	ID string `json:"id"`
}

type getOrderResponse struct {
	Order *dao.Order `json:"order"`
}

func makeGetOrderEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getOrderRequest)
		order, err := svc.GetOrder(ctx, req.ID)
		if err != nil {
			return nil, err
		}
		return getOrderResponse{
			Order: order,
		}, nil
	}
}
