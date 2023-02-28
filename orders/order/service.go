package order

import (
	"context"
	"fmt"
	"orders/dao"
	"time"

	"github.com/google/uuid"
	"github.com/greg9702/go-cadence-example/pkg/cadence/client"
	uc "go.uber.org/cadence/client"
)

type Service interface {
	CreateOrder(ctx context.Context, totalCost uint32, vehicleNo string) (string, error)
	GetOrder(ctx context.Context, id string) (*dao.Order, error)
	ApproveOrder(ctx context.Context, id string) error
	RejectOrder(ctx context.Context, id string) error
}

type service struct {
	dao dao.OrderDAO
	client client.CadenceAdapter
}

func NewService(dao dao.OrderDAO, c client.CadenceAdapter) Service {
	return &service{
		dao: dao,
		client: c,
	}
}

func (s *service) CreateOrder(ctx context.Context, totalCost uint32, vehicleNo string) (string, error) {
	workflowOptions := uc.StartWorkflowOptions{
		ID:                              "createOrder_" + uuid.New().String(),
		TaskList:                        "order",
		ExecutionStartToCloseTimeout:    time.Minute,
		DecisionTaskStartToCloseTimeout: time.Minute,
	}

	fmt.Println("Starting createOrder workflow")
	resp, err := s.client.CadenceClient.StartWorkflow(ctx, workflowOptions, "main.CreateOrderWorkflow")
	if err != nil {
		fmt.Println("XDDDDD ERRR", err)
	}
	fmt.Println(resp)
	return "", nil

	// id, err := s.dao.CreateOrder(ctx, totalCost, vehicleNo)
	// if err != nil {
	// 	return "", err
	// }
	// return id, nil
}

func (s *service) GetOrder(ctx context.Context, id string) (*dao.Order, error) {
	order, err := s.dao.GetOrder(ctx, id)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *service) ApproveOrder(ctx context.Context, id string) error {
	err := s.dao.UpdateOrderStatus(ctx, id, dao.CREATED)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) RejectOrder(ctx context.Context, id string) error {
	err := s.dao.UpdateOrderStatus(ctx, id, dao.FAILED)
	if err != nil {
		return err
	}
	return nil
}
