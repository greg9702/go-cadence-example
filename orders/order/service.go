package order

import (
	"context"
	"orders/dao"
)

type Service interface {
	CreateOrder(ctx context.Context, totalCost uint32, vehicleNo string) (string, error)
	GetOrder(ctx context.Context, id string) (*dao.Order, error)
}

type service struct {
	dao dao.OrderDAO
}

func NewService(dao dao.OrderDAO) Service {
	return &service{
		dao: dao,
	}
}

func (s *service) CreateOrder(ctx context.Context, totalCost uint32, vehicleNo string) (string, error) {
	id, err := s.dao.CreateOrder(ctx, totalCost, vehicleNo)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *service) GetOrder(ctx context.Context, id string) (*dao.Order, error) {
	order, err := s.dao.GetOrder(ctx, id)
	if err != nil {
		return nil, err
	}
	return order, nil
}
