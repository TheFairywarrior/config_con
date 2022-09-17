package publisher

import (
	"config_con/pkg/pipe/queue"
	"context"
	"fmt"
)

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


func (runner PublisherRunner) RunPublisher(cxt context.Context) {
	for {
		select {
		case <-cxt.Done():
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
