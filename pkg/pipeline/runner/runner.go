package runner

import (
	"context"

	"github.com/thefairywarrior/config_con/pkg/pipeline"
	"github.com/thefairywarrior/config_con/pkg/pipeline/consumer"
	"github.com/thefairywarrior/config_con/pkg/pipeline/publisher"
	"github.com/thefairywarrior/config_con/pkg/pipeline/transformer"
)

type Pipeline struct {
	consumer       consumer.Consumer
	transformSteps []transformer.TransformStep
	publishers     map[string]publisher.Publisher
}

func NewPipeline(consumer consumer.Consumer, transformSteps []transformer.TransformStep, publishers map[string]publisher.Publisher) *Pipeline {
	return &Pipeline{
		consumer:       consumer,
		transformSteps: transformSteps,
		publishers:     publishers,
	}
}

func (p *Pipeline) Transform(msg pipeline.Message) (transformer.TransformedData, error) {
	var out any = ""
	var err error
	for _, step := range p.transformSteps {
		out, err = step.Transform(msg)
		if err != nil {
			return nil, err
		}
	}
	return out.(transformer.TransformedData), nil
}

func (p *Pipeline) Publish(data transformer.TransformedData) error {
	for _, pubs := range data.List() {
		publisher := p.publishers[pubs]
		if err := publisher.Publish(data.Get(pubs).(map[string]any)); err != nil {
			return err
		}
	}
	return nil
}

func (p *Pipeline) Run(ctx context.Context) {
	c := make(chan pipeline.Message)
	go p.consumer.Consume(ctx, c)

	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-c:
			data, err := p.Transform(msg)
			if err != nil {
				//TODO ADD error handling here
				continue
			}
			if err := p.Publish(data); err != nil {
				//TODO ADD error handling here
				continue
			}
		}
	}
}