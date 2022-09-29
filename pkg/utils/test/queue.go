package test

import "config_con/pkg/pipe/queue"

type FakeMessage struct {
	queue.MessageData
	fakeData any
}

func NewFakeMessage(data any) FakeMessage {
	return FakeMessage{
		MessageData: queue.NewMessageData(),
		fakeData:    data,
	}
}

func (m FakeMessage) GetData() (any, error) {
	return m.fakeData, nil
}


