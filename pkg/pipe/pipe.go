package pipe

import (
	"config_con/pkg/pipe/consumer"
	"config_con/pkg/pipe/queue"
	"context"
	"fmt"
)

// Pipe is an instance of the full data pipeline.
// This is where the management for the connection between the stages is going to sit.
type Pipe struct {
	Context          context.Context
	Consumer         consumer.Consumer
	TransformerQueue queue.Queue
}

var currentPipes map[string]Pipe

func NewPipe(ctx context.Context, pipeName string, queue queue.LocalQueue, consumer consumer.Consumer) error {
	if _, ok := currentPipes[pipeName]; ok {
		return fmt.Errorf("pipe %s already exists", pipeName)
	}
	currentPipes[pipeName] = Pipe{
		Context:          ctx,
		Consumer:         consumer,
		TransformerQueue: queue,
	}
	return nil
}

// PipeConfig is the basic configuration for a pipe.
// And instance of this struct is built up from the yaml config file and will be used to
// create a Pipe instance.
type PipeConfig struct {
	Name     string `yaml:"name"`
	Consumer string `yaml:"consumer"`
}
