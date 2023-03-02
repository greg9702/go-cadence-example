package payment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/greg9702/go-cadence-example/pkg/errors"
)

type Service interface {
	ProcessPayment(ctx context.Context, id string, totalCost uint32) error
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) ProcessPayment(ctx context.Context, id string, totalCost uint32) error {
	fmt.Printf("Processing payment for id: %s, cost: %d\n", id, totalCost)
	if totalCost > 100 {
		return &errors.ServiceError{
			Code:    http.StatusNotFound,
			Message: "Not enough funds",
		}
	}
	return nil
}
