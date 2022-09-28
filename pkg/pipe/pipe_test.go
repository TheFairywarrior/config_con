package pipe

import (
	"config_con/pkg/pipe/queue"
	"config_con/pkg/pipe/transformer"
	"config_con/pkg/pipe/transformer/steps"
	"context"
	"testing"
)

type fakePublisher struct {
}

func (t fakePublisher) Publish(data []byte) error {
	return nil
}

type fakeConsumer struct {
}

func (fC fakeConsumer) Consume(ctx context.Context, queue queue.Queue) error {
	return nil
}

type fakeStep struct {

}


func TestPipe_Start(t *testing.T) {
	fP := fakePublisher{}
	fC := fakeConsumer{}
	trans := transformer.NewTransformer("name", []steps.Step{})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p := NewPipe(ctx, fC, trans, fP)

	p.Start()
}
