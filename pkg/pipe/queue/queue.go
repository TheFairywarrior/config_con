package queue


type Queue interface {
	Add(data Message) error
	Chan() <-chan Message
	Close()
}

// LocalQueue is a queue that is local to the current process.
// This ququ is going to be used as a local queue to pass messages between points
// within the current runtime.
type LocalQueue struct {
	queue chan Message
}

// Add adds a data to the queue.
func (q LocalQueue) Add(data Message) error {
	q.queue <- data
	return nil
}

// Chan returns the basic channel that the queue is using.
func (q LocalQueue) Chan() <-chan Message {
	return q.queue
}

// Close closes the queue.
func (q LocalQueue) Close() {
	close(q.queue)
}

func NewQueue(size int) LocalQueue {
	return LocalQueue{
		queue: make(chan Message, size),
	}
}
