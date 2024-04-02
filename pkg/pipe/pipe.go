package pipe

import (
	"context"

	"github.com/thefairywarrior/config_con/pkg/pipe/consumer"
	"github.com/thefairywarrior/config_con/pkg/pipe/publisher"
	"github.com/thefairywarrior/config_con/pkg/pipe/queue"
	"github.com/thefairywarrior/config_con/pkg/pipe/transformer"
)

// Pipe is an instance of the full data pipeline.
// This is where the management for the connection between the stages is going to sit.
type Pipe struct {
	ctx         context.Context
	consumer    consumer.Consumer
	transformer transformer.Transformer
	publisher   publisher.Publisher
}

func (pipe Pipe) Start() {
	transformerQueue := queue.NewLocalQueue(1)
	publisherQueue := queue.NewLocalQueue(1)

	go pipe.consumer.Consume(pipe.ctx, transformerQueue)
	go pipe.transformer.StartTransformer(pipe.ctx, transformerQueue, publisherQueue)

	publisherRunner := publisher.NewPublisherRunner(pipe.publisher, publisherQueue)
	go publisherRunner.RunPublisher(pipe.ctx)
	go func() {
		<-pipe.ctx.Done()
		transformerQueue.Close()
		publisherQueue.Close()
	}()
}

func NewPipe(
	ctx context.Context,
	consumer consumer.Consumer,
	transformer transformer.Transformer,
	publisher publisher.Publisher,
) Pipe {
	return Pipe{
		ctx:         ctx,
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
