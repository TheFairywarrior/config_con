package pipe

import (
	"config_con/pkg/pipe/consumer"
	"config_con/pkg/pipe/publisher"
	"config_con/pkg/pipe/queue"
	"config_con/pkg/pipe/transformer"
	"context"
	"fmt"
)

// Pipe is an instance of the full data pipeline.
// This is where the management for the connection between the stages is going to sit.
type Pipe struct {
	Context          context.Context
	Consumer         consumer.Consumer
	TransformerQueue queue.Queue
	Transformer      transformer.Transformer
	PublisherQueue   queue.Queue
	Publisher        publisher.Publisher
}

var currentPipes map[string]Pipe

func NewPipe(
	ctx context.Context,
	pipeName string,
	transformerQueue queue.Queue,
	consumer consumer.Consumer,
	transformer transformer.Transformer,
	publisherQueue queue.Queue,
	publisher publisher.Publisher,
) error {
	if _, ok := currentPipes[pipeName]; ok {
		return fmt.Errorf("pipe %s already exists", pipeName)
	}
	currentPipes[pipeName] = Pipe{
		Context:          ctx,
		Consumer:         consumer,
		TransformerQueue: transformerQueue,
		Transformer:      transformer,
		PublisherQueue:   publisherQueue,
		Publisher:        publisher,
	}
	return nil
}

// PipeConfig is the basic configuration for a pipe.
// And instance of this struct is built up from the yaml config file and will be used to
// create a Pipe instance.
type PipeConfig struct {
	Name        string `yaml:"name"`
	Consumer    string `yaml:"consumer"`
	Transformer string `yaml:"transformer"`
	Publisher   string `yaml:"publisher"`
}
