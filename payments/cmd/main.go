package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"payments/payment"

	"github.com/gorilla/mux"
	uc "go.uber.org/cadence/client"
	"go.uber.org/cadence/worker"

	"github.com/greg9702/go-cadence-example/pkg/cadence"
	"github.com/greg9702/go-cadence-example/pkg/cadence/client"

	"github.com/go-kit/kit/log"
)

func main() {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "listen", "8082", "caller", log.DefaultCaller)

	_ = payment.NewService()


	config := cadence.SetupConfig("../config/development.yaml")

	var c client.CadenceAdapter
	c.Setup(config)


	workerOptions := worker.Options{
		FeatureFlags: uc.FeatureFlags{
			WorkflowExecutionAlreadyCompletedErrorEnabled: true,
		},
	}

	w := worker.New(c.ServiceClient, config.Domain, "order", workerOptions)
	w.RegisterActivity(processPayment)

	w.Start()


	r := mux.NewRouter()

	logger.Log("msg", "HTTP", "addr", "8081")
	logger.Log("err", http.ListenAndServe(":8081", r))
}

func processPayment(ctx context.Context) (string, error) {
	fmt.Println(fmt.Sprintf("process payment activity trigerred"))
	return "foo", nil
}

