package order

import (
	"fmt"
	"time"

	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

// func StartCreateOrderWorker() {
// 	config := cadence.SetupConfig("../config/development.yaml")
// 	service, err := worker.SetupService(config)
// 	if err != nil {
// 		panic(err)
// 	}

// 	logger, _ := zap.NewDevelopment()

// 	workerOptions := w.Options{
// 		Logger:       logger,
// 	}

// 	taskList := "createOrder"

// 	cadenceWorker := w.New(service, config.Domain, taskList, workerOptions)
// 	err = cadenceWorker.Start()
// 	if err != nil {
// 		logger.Error("Failed to start workers.", zap.Error(err))
// 		panic("Failed to start workers")
// 	}

// 	cadenceWorker.RegisterWorkflow(helloWorldWorkflow)
// }

func HelloWorldWorkflow(ctx workflow.Context, name string) error {

	fmt.Println("XDDDDDD")
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("helloworld workflow started")
	var helloworldResult string
	err := workflow.ExecuteActivity(ctx, "createOrder", name).Get(ctx, &helloworldResult)
	if err != nil {
		logger.Error("Activity failed.", zap.Error(err))
		return err
	}
	return nil
}