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
	return MessageData{
		id:        uuid.New().String(),
		timestamp: time.Now(),
	}
}
