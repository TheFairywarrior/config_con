package pipe

import (
	"context"
	"github.com/thefairywarrior/config_con/pkg/pipe/queue"
	"github.com/thefairywarrior/config_con/pkg/pipe/transformer"
	"github.com/thefairywarrior/config_con/pkg/pipe/transformer/steps"
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

func TestPipe_Start(t *testing.T) {
	fP := fakePublisher{}
	fC := fakeConsumer{}
	trans := transformer.NewTransformer("name", []steps.Step{})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p := NewPipe(ctx, fC, trans, fP)

	p.Start()
}
