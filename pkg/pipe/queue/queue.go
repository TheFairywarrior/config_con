package queue

// Queue is a queue of Transformer.
// This queue is going to be used by the consumer to push the data to its associated Transformer.
//
// The reason that this is a struct is for expandability later. The struct can be changed to a interface later
// and a new struct can be created for a simple queue like below or one that works with a cluster.
type Queue struct {
	queue chan Message
}

// Add adds a data to the queue.
func (q Queue) Add(data Message) {
	q.queue <- data
}

// Chan returns the basic channel that the queue is using.
func (q Queue) Chan() <-chan Message {
	return q.queue
}

// Close closes the queue.
func (q Queue) Close() {
	close(q.queue)
}

func NewQueue(size int) Queue {
	return Queue{
		queue: make(chan Message, size),
	}
}
