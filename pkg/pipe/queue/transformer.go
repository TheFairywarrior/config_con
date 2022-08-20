package queue

import "config_con/pkg/utils/shortcuts"

// TransformerQueue is a queue of Transformer.
// This queue is going to be used by the consumer to push the data to its associated Transformer.
//
// The reason that this is a struct is for expandability later. The struct can be changed to a interface later
// and a new struct can be created for a simple queue like below or one that works with a cluster.
type TransformerQueue struct {
	queue chan shortcuts.Map
}

// Add adds a data to the queue.
func (q TransformerQueue) Add(data shortcuts.Map) {
	q.queue <- data
}

// Chan returns the basic channel that the queue is using.
func (q TransformerQueue) Chan() <-chan shortcuts.Map {
	return q.queue
}

// Close closes the queue.
func (q TransformerQueue) Close() {
	close(q.queue)
}
