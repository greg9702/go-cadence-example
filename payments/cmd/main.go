package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"payments/payment"

	"github.com/greg9702/go-cadence-example/pkg/cadence"
	"github.com/greg9702/go-cadence-example/pkg/cadence/client"

	"github.com/go-kit/kit/log"
)

func main() {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "listen", "8082", "caller", log.DefaultCaller)

	_ = payment.NewService()



	config := cadence.SetupConfig("development.yaml")

	var c client.CadenceAdapter
	c.Setup(config)

	svc := payment.NewService()

	r := payment.NewHttpServer(svc, &c, logger)
	w := payment.NewCadenceWorker(svc, &c, config, logger)

	w.Start()



	logger.Log("msg", "HTTP", "addr", "8081")
	logger.Log("err", http.ListenAndServe(":8081", r))
}

func processPayment(ctx context.Context) (string, error) {
	fmt.Println(fmt.Sprintf("process payment activity trigerred"))
	return "payment", errors.New("foo")
}

