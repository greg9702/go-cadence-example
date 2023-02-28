package main

import (
	"fmt"
	"go-cadence-example/orchestrator/order"

	"github.com/greg9702/go-cadence-example/pkg/cadence"
	"github.com/greg9702/go-cadence-example/pkg/cadence/client"
	c "go.uber.org/cadence/client"
	"go.uber.org/cadence/worker"
	"go.uber.org/cadence/workflow"
)


func main() {
	config := cadence.SetupConfig("../config/development.yaml")

	var client client.CadenceAdapter
	client.Setup(config)

	workerOptions := worker.Options{
		FeatureFlags: c.FeatureFlags{
			WorkflowExecutionAlreadyCompletedErrorEnabled: true,
		},
	}

	worker := worker.New(client.ServiceClient, config.Domain, "order", workerOptions)

	worker.RegisterWorkflowWithOptions(order.HelloWorldWorkflow, workflow.RegisterOptions{Name: "createOrderWorkflow"})
	fmt.Println("Registered workflow")

	worker.Start()
	select {}
}


