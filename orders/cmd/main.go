package main

import (
	"net/http"
	"orders/dao"
	"os"

	"orders/order"

	"github.com/go-kit/kit/log"
	"github.com/greg9702/go-cadence-example/pkg/cadence"
	"github.com/greg9702/go-cadence-example/pkg/cadence/client"
)

func main() {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "listen", "8082", "caller", log.DefaultCaller)


	config := cadence.SetupConfig("../config/development.yaml")

	var c client.CadenceAdapter
	c.Setup(config)

	orderDAO := dao.NewOrderDAO()
	svc := order.NewService(orderDAO, &c)

	r := order.NewHttpServer(svc, &c, logger)
	w := order.NewCadenceWorker(svc, &c, config, logger)

	w.Start()

	logger.Log("msg", "HTTP", "addr", "8082")
	logger.Log("err", http.ListenAndServe(":8082", r))
}



