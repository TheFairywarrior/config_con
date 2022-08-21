package consumer

import (
	"config_con/pkg/pipe/consumer/api"
	"config_con/pkg/pipe/queue"
	"context"
)

// Consumer interface is used in the pipeline to consume the data from multiple sources.
type Consumer interface {
	Consume(context.Context, chan queue.TransformerQueue)
}

type ConsumerConfig struct {
	Name string               `yaml:"name"`
	Api  api.ApiConfiguration `yaml:"api"`
}
