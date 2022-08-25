package queue


import (
	"config_con/pkg/utils/shortcuts"
	"testing"
)

func TestTransformerQueue_Crud(t *testing.T) {
	queue := TransformerQueue{
		queue: make(chan any, 1),
	}
	queue.Add(shortcuts.Map{})
	<- queue.Chan()
	queue.Add(shortcuts.Map{})
	<-queue.Chan()
	queue.Close()
}
