package queue

import (
	"time"

	"github.com/google/uuid"
)

type Message interface {
	GetData() (any, error)
	ID() string
	Timestamp() time.Time
	Meta() MessageData
}


// MessageData holds the meta data for the message.
// An instance of this struct is created for each message, and will
// be passed to all the steps in the pipeline.
type MessageData struct {
	id        string
	timestamp time.Time
}

func (m MessageData) Meta() MessageData {
	return m
}

func (m MessageData) ID() string {
	return m.id
}

func (m MessageData) Timestamp() time.Time {
	return m.timestamp
}


func NewMessageData() MessageData {
	return NewMessageWithId(uuid.NewString())
}

func NewMessageWithId(id string) MessageData {
	return MessageData{
		id:        id,
		timestamp: time.Now(),
	}
}