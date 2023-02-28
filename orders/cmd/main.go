package main

import (
	"context"
	"fmt"
	"net/http"
	"orders/dao"
	"os"

	"orders/order"

	"github.com/go-kit/kit/log"
	"github.com/greg9702/go-cadence-example/pkg/cadence"
	"github.com/greg9702/go-cadence-example/pkg/cadence/client"
	uc "go.uber.org/cadence/client"
	"go.uber.org/cadence/worker"
)

func main() {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "listen", "8082", "caller", log.DefaultCaller)

	orderDAO := dao.NewOrderDAO()

	config := cadence.SetupConfig("../config/development.yaml")

	var c client.CadenceAdapter
	c.Setup(config)

	r := order.NewHttpServer(order.NewService(orderDAO, c), logger)


	workerOptions := worker.Options{
		FeatureFlags: uc.FeatureFlags{
			WorkflowExecutionAlreadyCompletedErrorEnabled: true,
		},
	}

	w := worker.New(c.ServiceClient, config.Domain, "order", workerOptions)
	w.RegisterActivity(createOrder)

	w.Start()

	logger.Log("msg", "HTTP", "addr", "8082")
	logger.Log("err", http.ListenAndServe(":8082", r))
}


func createOrder(ctx context.Context) (string, error) {
	fmt.Println(fmt.Sprintf("order created"))
	return "order", nil
}
