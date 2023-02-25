package main

import (
	"net/http"
	"orders/dao"
	"os"

	"orders/order"

	"github.com/go-kit/kit/log"
)

func main() {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "listen", "8082", "caller", log.DefaultCaller)

	orderDAO := dao.NewOrderDAO()

	r := order.NewHttpServer(order.NewService(orderDAO), logger)

	logger.Log("msg", "HTTP", "addr", "8082")
	logger.Log("err", http.ListenAndServe(":8082", r))
}
