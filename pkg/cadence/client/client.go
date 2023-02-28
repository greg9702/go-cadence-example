package client

import (
	"github.com/greg9702/go-cadence-example/pkg/cadence"
	"github.com/greg9702/go-cadence-example/pkg/cadence/builder"
	"go.uber.org/cadence/client"
	"go.uber.org/zap"
)


func SetupClient(config *cadence.CadenceConfig) (client.Client, error) {
	logger, _ := zap.NewDevelopment()
	b := builder.NewBuilder(logger, config.HostPort, config.Domain)
	client, err := b.BuildCadenceClient()
	if err != nil {
		return nil, err
	}

	return client, nil
}