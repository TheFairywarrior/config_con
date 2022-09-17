package pipe

import (
	"config_con/pkg/pipe/consumer"
	"config_con/pkg/pipe/publisher"
	"config_con/pkg/pipe/queue"
	"config_con/pkg/pipe/transformer"
	"context"
)

// Pipe is an instance of the full data pipeline.
// This is where the management for the connection between the stages is going to sit.
type Pipe struct {
	cxt         context.Context
	consumer    consumer.Consumer
	transformer transformer.Transformer
	publisher   publisher.Publisher
}

func (pipe Pipe) Start() {
	transformerQueue := queue.NewQueue(1)
	publisherQueue := queue.NewQueue(1)

	go pipe.consumer.Consume(pipe.cxt, transformerQueue)
	go pipe.transformer.StartTransformer(pipe.cxt, transformerQueue, publisherQueue)
	
	publisherRunner := publisher.NewPublisherRunner(pipe.publisher, publisherQueue)
	go publisherRunner.RunPublisher(pipe.cxt)
}

func NewPipe(
	ctx context.Context,
	consumer consumer.Consumer,
	transformer transformer.Transformer,
	publisher publisher.Publisher,
) Pipe {
	return Pipe{
		cxt:         ctx,
		consumer:    consumer,
		transformer: transformer,
		publisher:   publisher,
	}
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
