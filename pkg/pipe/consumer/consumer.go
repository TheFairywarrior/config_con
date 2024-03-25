package consumer

import (
	"context"
	"github.com/thefairywarrior/config_con/pkg/pipe/queue"
)

// Consumer interface is used in the pipeline to consume the data from multiple sources.
type Consumer interface {
	Consume(context.Context, queue.Queue) error
}
