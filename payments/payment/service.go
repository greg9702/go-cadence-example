package payment

import (
	"context"
	"github.com/greg9702/go-cadence-example/pkg/errors"
	"net/http"
)

type Service interface {
	ProcessPayment(ctx context.Context, totalCost uint32) error
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) ProcessPayment(ctx context.Context, totalCost uint32) error {
	if totalCost > 100 {
		return &errors.ServiceError{
			Code:    http.StatusNotFound,
			Message: "Not enough funds",
		}
	}
	return nil
}
