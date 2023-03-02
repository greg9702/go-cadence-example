package payment

import (
	"github.com/greg9702/go-cadence-example/pkg/cadence"
	"github.com/greg9702/go-cadence-example/pkg/cadence/client"
	"go.uber.org/cadence/worker"

	kitlog "github.com/go-kit/kit/log"
	"github.com/gorilla/mux"

	uc "go.uber.org/cadence/client"
)


func NewHttpServer(svc Service, client *client.CadenceAdapter, logger kitlog.Logger) *mux.Router {
	r := mux.NewRouter()

	return r
}

func NewCadenceWorker(svc Service, client *client.CadenceAdapter, config *cadence.CadenceConfig, logger kitlog.Logger) worker.Worker {
	workerOptions := worker.Options{
		FeatureFlags: uc.FeatureFlags{
			WorkflowExecutionAlreadyCompletedErrorEnabled: true,
		},
	}

	w := worker.New(client.ServiceClient, config.Domain, "payment", workerOptions)
	w.RegisterActivity(svc.ProcessPayment)

	return w
}