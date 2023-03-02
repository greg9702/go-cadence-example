package order

import (
	"fmt"
	"time"

	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

func CreateOrderWorkflow(ctx workflow.Context, id string, totalCost uint32) (string, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("helloworld workflow started")

	fmt.Println("Starting workflow")
	err := processPayment(ctx, id, totalCost)
	if err != nil {
		logger.Error("Activity failed.", zap.Error(err))

		err = cancelOrder(ctx, id)
		if err != nil {
			return "failed", nil
		}

		return "cancelled", nil
	}

	err = approveOrder(ctx, id)
	if err != nil {
		logger.Error("Activity failed.", zap.Error(err))
		return "failed", err
	}


	return "completed", nil
}

func processPayment(ctx workflow.Context, id string, totalCost uint32) error {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
		TaskList: "payment",
	}

	ctx = workflow.WithActivityOptions(ctx, ao)
	err := workflow.ExecuteActivity(ctx, "payments/payment.Service.ProcessPayment", id, totalCost).Get(ctx, nil)
	if err != nil {
		return err
	}
	return nil
}

func cancelOrder(ctx workflow.Context, id string) error {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
		TaskList: "order",
	}

	ctx = workflow.WithActivityOptions(ctx, ao)
	err := workflow.ExecuteActivity(ctx, "orders/order.Service.RejectOrder", id).Get(ctx, nil)
	if err != nil {
		return err
	}
	return nil
}


func approveOrder(ctx workflow.Context, id string) error {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
		TaskList: "order",
	}

	ctx = workflow.WithActivityOptions(ctx, ao)
	err := workflow.ExecuteActivity(ctx, "orders/order.Service.ApproveOrder", id).Get(ctx, nil)
	if err != nil {
		return err
	}
	return nil
}