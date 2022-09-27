package publisher

// Publisher interface used as a base for all publishers.
type Publisher interface {
	// Publish is used to publish the data to the publisher.
	Publish(data []byte) error
}
