package worker

import (
	"github.com/greg9702/go-cadence-example/pkg/cadence"
	"github.com/greg9702/go-cadence-example/pkg/cadence/builder"
	"go.uber.org/cadence/.gen/go/cadence/workflowserviceclient"
	"go.uber.org/zap"
)

func SetupService(config *cadence.CadenceConfig) (workflowserviceclient.Interface, error) {
	logger, _ := zap.NewDevelopment()
	b := builder.NewBuilder(logger, config.HostPort, config.Domain)
	service, err := b.BuildServiceClient()
	if err != nil {
		return nil, err
	}
	
	return service, nil
}