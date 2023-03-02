package main

import (
	"net/http"
	"os"

	"orchestrator/order"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/greg9702/go-cadence-example/pkg/cadence"
	"github.com/greg9702/go-cadence-example/pkg/cadence/client"
	uc "go.uber.org/cadence/client"
	"go.uber.org/cadence/worker"
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
	w.RegisterWorkflow(order.CreateOrderWorkflow)

	w.Start()

	logger.Log("msg", "HTTP", "addr", "8083")
	logger.Log("err", http.ListenAndServe(":8083", r))
}


