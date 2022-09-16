package pipe

import (
	"config_con/pkg/pipe/consumer"
	"config_con/pkg/pipe/publisher"
	"config_con/pkg/pipe/transformer"
	"context"
)

// Pipe is an instance of the full data pipeline.
// This is where the management for the connection between the stages is going to sit.
type Pipe struct {
	context     context.Context
	consumer    consumer.Consumer
	transformer transformer.Transformer
	publisher   publisher.Publisher
}

func NewPipe(
	ctx context.Context,
	consumer consumer.Consumer,
	transformer transformer.Transformer,
	publisher publisher.Publisher,
) Pipe {
	return Pipe{
		context:     ctx,
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
