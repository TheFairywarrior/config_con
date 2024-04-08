package consumer

import (
	"context"

	"github.com/thefairywarrior/config_con/pkg/pipeline"
)


type Consumer interface {
	Consume(ctx context.Context, c chan  pipeline.Message) error
}


