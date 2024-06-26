package publisher

import (
	"context"
	"fmt"

	"github.com/thefairywarrior/config_con/pkg/pipe/queue"
)

// PublisherRunner is the manager for running publisher instances.
// It holds the queue that will be read from and then passed into the publisher.
type PublisherRunner struct {
	publisher Publisher
	queue     queue.Queue
}

func (runner PublisherRunner) runPublisher(message queue.Message) {
	messageData, err := message.GetData()
	if err != nil {
		fmt.Println(err.Error())
	}
	err = runner.publisher.Publish(messageData.([]byte))
	if err != nil {
		fmt.Println(err.Error())
	}
}

// RunPublisher is the function that controls the publisher instance.
// It will run the publisher until the context is done.
func (runner PublisherRunner) RunPublisher(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case message := <-runner.queue.Chan():
			runner.runPublisher(message)
		}
	}
}

func NewPublisherRunner(publisher Publisher, queue queue.Queue) PublisherRunner {
	return PublisherRunner{
		publisher: publisher,
		queue:     queue,
	}
}
