package queue

import (
	"testing"
)

type TestMessage struct {
	MessageData
	data string
}

func (m TestMessage) GetData() any {
	return m.data
}

func TestTransformerQueue_Crud(t *testing.T) {
	queue := Queue{
		queue: make(chan Message, 1),
	}

	message := TestMessage{
		MessageData: NewMessageData(),
		data:        "test",
	}
	queue.Add(message)
	<-queue.Chan()
	queue.Add(message)
	<-queue.Chan()
	queue.Close()
}
