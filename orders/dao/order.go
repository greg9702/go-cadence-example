package dao

import (
	"context"
	"github.com/google/uuid"
	"github.com/greg9702/go-cadence-example/pkg/errors"
	"github.com/greg9702/go-cadence-example/pkg/logger"
	"net/http"
)

type OrderDAO interface {
	CreateOrder(ctx context.Context, totalCost uint32, vehicleNo string) (string, error)
	GetOrder(ctx context.Context, id string) (*Order, error)
}

func NewOrderDAO() OrderDAO {
	return &orderInMemory{
		data: make(map[string]*Order),
	}
}

type orderInMemory struct {
	data map[string]*Order
}

func (o *orderInMemory) CreateOrder(ctx context.Context, totalCost uint32, vehicleNo string) (string, error) {
	log := logger.GetTracedLog(ctx)

	id := uuid.New().String()

	o.data[id] = &Order{
		ID:        id,
		TotalCost: totalCost,
		VehicleNo: vehicleNo,
		Status:    INITATED,
	}

	log.Log("msg", "order created", "id", id)

	return id, nil
}

func (o *orderInMemory) GetOrder(ctx context.Context, id string) (*Order, error) {
	log := logger.GetTracedLog(ctx)

	if order, ok := o.data[id]; ok {
		return order, nil
	}
	log.Log("msg", "order not found", "id", id)

	return nil, &errors.ServiceError{
		Code:    http.StatusNotFound,
		Message: "order not found",
	}
}
