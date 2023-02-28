package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/greg9702/go-cadence-example/pkg/cadence"
	"github.com/greg9702/go-cadence-example/pkg/cadence/client"
	uc "go.uber.org/cadence/client"
	"go.uber.org/cadence/worker"
	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

func main() {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "listen", "8082", "caller", log.DefaultCaller)

	config := cadence.SetupConfig("../config/development.yaml")

	var c client.CadenceAdapter
	c.Setup(config)

	r := mux.NewRouter()

	workerOptions := worker.Options{
		FeatureFlags: uc.FeatureFlags{
			WorkflowExecutionAlreadyCompletedErrorEnabled: true,
		},
	}

	w := worker.New(c.ServiceClient, config.Domain, "order", workerOptions)
	w.RegisterWorkflow(CreateOrderWorkflow)

	w.Start()

	logger.Log("msg", "HTTP", "addr", "8083")
	logger.Log("err", http.ListenAndServe(":8083", r))
}

func CreateOrderWorkflow(ctx workflow.Context) (string, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("helloworld workflow started")

	fmt.Println("Starting workflow")
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
		TaskList: "order",
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	var order string
	err := workflow.ExecuteActivity(ctx, "main.createOrder").Get(ctx, &order)
	if err != nil {
		logger.Error("Activity failed.", zap.Error(err))
		return "", err
	}

	ao = workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
		TaskList: "payment",
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	var payment string
	err = workflow.ExecuteActivity(ctx, "main.processPayment").Get(ctx, &payment)
	if err != nil {
		logger.Error("Activity failed.", zap.Error(err))
		return "", err
	}
	return fmt.Sprintf("%s-%s", order, payment), nil
}
